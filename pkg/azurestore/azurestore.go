package azurestore

import (
	"bufio"
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/tus/tusd/internal/uid"
	"github.com/tus/tusd/pkg/handler"
	"io"
	"path/filepath"
	"strings"
)

type AzureStore struct {
	Service      *AzService
	ObjectPrefix string
	Container    string
}

type AzureUpload struct {
	ID          string
	InfoBlob    FileBlob
	FileBlob    FileBlob
	InfoHandler *handler.FileInfo
}

func New(service *AzService) *AzureStore {
	return &AzureStore{
		Service: service,
	}
}

// UseIn sets this store as the core data store in the passed composer and adds
// all possible extension to it.
func (store AzureStore) UseIn(composer *handler.StoreComposer) {
	composer.UseCore(store)
	composer.UseTerminater(store)
	composer.UseLengthDeferrer(store)
}

// Create new upload InfoHandler file and ID
func (store AzureStore) NewUpload(ctx context.Context, info handler.FileInfo) (handler.Upload, error) {

	var id string
	if info.ID == "" {
		id = uid.Uid()
		info.ID = id
	} else {
		id = info.ID
	}

	idInfo := store.infoPath(id)

	// create the info file
	infoBlob, err := store.Service.NewFileBlob(ctx, idInfo)

	if err != nil {
		return nil, err
	}

	info.Storage = map[string]string{
		"Type":      "azurestore",
		"Container": store.Container,
		"Key":       store.keyWithPrefix(id),
	}

	var blobType BlobType

	if info.Size > int64(MaxBlockBlobSize) {
		return nil, fmt.Errorf("azurestore: max upload of %d bytes exceeded MaxAppendBlobSize of %d and"+
			" MaxBlockBlobSize of"+
			" %d bytes",
			info.Size, int64(MaxAppendBlobSize), int64(MaxBlockBlobSize))
	} else {
		if info.Size > int64(MaxAppendBlobSize) {
			blobType = BlockBlobType
		} else {
			blobType = AppendBlobType
		}
	}

	fileBlob, err := store.Service.NewFileBlob(ctx, id, SetBlobType(blobType))

	if err != nil {
		return nil, err
	}

	azureUpload := &AzureUpload{
		ID:          info.ID,
		InfoHandler: &info,
		InfoBlob:    infoBlob,
		FileBlob:    fileBlob,
	}

	// write the info file
	err = azureUpload.writeInfo(ctx)
	if err != nil {
		return nil, fmt.Errorf("azurestore: unable to create InfoHandler file:\n%s", err)
	}

	return azureUpload, nil
}

// Get the file info and file offset from Azure Storage
func (store AzureStore) GetUpload(ctx context.Context, id string) (handle handler.Upload, err error) {

	info := handler.FileInfo{}

	infoFileName := store.infoPath(id)

	infoBlob, err := store.Service.NewFileBlob(ctx, infoFileName)

	if err != nil {
		return nil, err
	}

	// Download info file from Azure storage
	data, err := infoBlob.Download(ctx)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &info); err != nil {
		return nil, err
	}

	fileBlob, err := store.Service.NewFileBlob(ctx, id)

	if err != nil {
		return nil, err
	}

	offset, err := fileBlob.Offset(ctx)

	if err != nil {
		return nil, err
	}

	// Set the offset
	info.Offset = offset

	return &AzureUpload{
		ID:          id,
		InfoBlob:    infoBlob,
		FileBlob:    fileBlob,
		InfoHandler: &info,
	}, nil
}

func (store AzureStore) AsTerminatableUpload(upload handler.Upload) handler.TerminatableUpload {
	return upload.(*AzureUpload)
}

func (store AzureStore) AsLengthDeclarableUpload(upload handler.Upload) handler.LengthDeclarableUpload {
	return upload.(*AzureUpload)
}

func (upload *AzureUpload) WriteChunk(ctx context.Context, offset int64, src io.Reader) (int64, error) {

	r := bufio.NewReader(src)
	buf := new(bytes.Buffer)
	n, err := r.WriteTo(buf)
	if err != nil {
		return 0, err
	}

	// Get the max chunk size for this specific blob type (append / block)
	maxChunkSize := upload.FileBlob.MaxChunkSizeLimit()

	chunkSize := int64(binary.Size(buf.Bytes()))
	if chunkSize > maxChunkSize {
		return 0, fmt.Errorf("azurestore: Chunk of size %v too large. Max chunk size is %v", chunkSize, maxChunkSize)
	}

	re := bytes.NewReader(buf.Bytes())
	err = upload.FileBlob.Upload(ctx, re)
	if err != nil {
		return 0, err
	}

	upload.InfoHandler.Offset += n

	return n, nil
}

func (upload *AzureUpload) GetInfo(ctx context.Context) (handler.FileInfo, error) {
	info := handler.FileInfo{}

	if upload.InfoHandler != nil {
		return *upload.InfoHandler, nil
	}

	data, err := upload.InfoBlob.Download(ctx)
	if err != nil {
		return info, err
	}

	if err := json.Unmarshal(data, &info); err != nil {
		return info, err
	}

	upload.InfoHandler = &info
	return info, nil
}

// Get the upload file from the Azure storage
func (upload *AzureUpload) GetReader(ctx context.Context) (io.Reader, error) {
	b, err := upload.FileBlob.Download(ctx)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

// Finish the file upload
func (upload *AzureUpload) FinishUpload(ctx context.Context) error {
	return upload.FileBlob.Close(ctx)
}

// Delete files
func (upload *AzureUpload) Terminate(ctx context.Context) error {
	// Delete InfoHandler
	err := upload.InfoBlob.Delete(ctx)
	if err != nil {
		return err
	}

	// Delete file
	err = upload.FileBlob.Delete(ctx)
	return err
}

func (store AzureStore) binPath(id string) string {
	return filepath.Join(store.Service.ContainerURL.String(), id)
}

func (store AzureStore) infoPath(id string) string {
	return filepath.Join(store.Service.ContainerURL.String(), id+".info")
}

func (upload *AzureUpload) writeInfo(ctx context.Context) (err error) {
	data, err := json.Marshal(upload.InfoHandler)
	if err != nil {
		return err
	}

	r := bytes.NewReader(data)

	err = upload.InfoBlob.Upload(ctx, r)

	return err
}

func (upload *AzureUpload) DeclareLength(ctx context.Context, length int64) error {
	upload.InfoHandler.Size = length
	upload.InfoHandler.SizeIsDeferred = false
	return upload.writeInfo(ctx)
}

func (store *AzureStore) keyWithPrefix(key string) string {
	prefix := store.ObjectPrefix
	if prefix != "" && !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}

	return prefix + key
}
