// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tus/tusd/pkg/azurestore (interfaces: AzService,AzBlob,BlockBlob,AppendBlob)

// Package azurestore_test is a generated GoMock package.
package azurestore_test

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	azurestore "github.com/tus/tusd/pkg/azurestore"
	io "io"
	reflect "reflect"
)

// MockAzService is a mock of AzService interface
type MockAzService struct {
	ctrl     *gomock.Controller
	recorder *MockAzServiceMockRecorder
}

// MockAzServiceMockRecorder is the mock recorder for MockAzService
type MockAzServiceMockRecorder struct {
	mock *MockAzService
}

// NewMockAzService creates a new mock instance
func NewMockAzService(ctrl *gomock.Controller) *MockAzService {
	mock := &MockAzService{ctrl: ctrl}
	mock.recorder = &MockAzServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAzService) EXPECT() *MockAzServiceMockRecorder {
	return m.recorder
}

// ContainerURL mocks base method
func (m *MockAzService) ContainerURL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerURL")
	ret0, _ := ret[0].(string)
	return ret0
}

// ContainerURL indicates an expected call of ContainerURL
func (mr *MockAzServiceMockRecorder) ContainerURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerURL", reflect.TypeOf((*MockAzService)(nil).ContainerURL))
}

// GetFileBlob mocks base method
func (m *MockAzService) GetFileBlob(arg0 string) (azurestore.AzBlob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileBlob", arg0)
	ret0, _ := ret[0].(azurestore.AzBlob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileBlob indicates an expected call of GetFileBlob
func (mr *MockAzServiceMockRecorder) GetFileBlob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileBlob", reflect.TypeOf((*MockAzService)(nil).GetFileBlob), arg0)
}

// NewContainer mocks base method
func (m *MockAzService) NewContainer(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewContainer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewContainer indicates an expected call of NewContainer
func (mr *MockAzServiceMockRecorder) NewContainer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewContainer", reflect.TypeOf((*MockAzService)(nil).NewContainer), arg0)
}

// NewFileBlob mocks base method
func (m *MockAzService) NewFileBlob(arg0 context.Context, arg1 string, arg2 ...azurestore.OptionFileBlob) (azurestore.AzBlob, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewFileBlob", varargs...)
	ret0, _ := ret[0].(azurestore.AzBlob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewFileBlob indicates an expected call of NewFileBlob
func (mr *MockAzServiceMockRecorder) NewFileBlob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewFileBlob", reflect.TypeOf((*MockAzService)(nil).NewFileBlob), varargs...)
}

// MockAzBlob is a mock of AzBlob interface
type MockAzBlob struct {
	ctrl     *gomock.Controller
	recorder *MockAzBlobMockRecorder
}

// MockAzBlobMockRecorder is the mock recorder for MockAzBlob
type MockAzBlobMockRecorder struct {
	mock *MockAzBlob
}

// NewMockAzBlob creates a new mock instance
func NewMockAzBlob(ctrl *gomock.Controller) *MockAzBlob {
	mock := &MockAzBlob{ctrl: ctrl}
	mock.recorder = &MockAzBlobMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAzBlob) EXPECT() *MockAzBlobMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockAzBlob) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockAzBlobMockRecorder) Close(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAzBlob)(nil).Close), arg0)
}

// Delete mocks base method
func (m *MockAzBlob) Delete(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockAzBlobMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAzBlob)(nil).Delete), arg0)
}

// Download mocks base method
func (m *MockAzBlob) Download(arg0 context.Context) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download
func (mr *MockAzBlobMockRecorder) Download(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockAzBlob)(nil).Download), arg0)
}

// Exists mocks base method
func (m *MockAzBlob) Exists(arg0 context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists
func (mr *MockAzBlobMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockAzBlob)(nil).Exists), arg0)
}

// MaxChunkSizeLimit mocks base method
func (m *MockAzBlob) MaxChunkSizeLimit() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxChunkSizeLimit")
	ret0, _ := ret[0].(int64)
	return ret0
}

// MaxChunkSizeLimit indicates an expected call of MaxChunkSizeLimit
func (mr *MockAzBlobMockRecorder) MaxChunkSizeLimit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxChunkSizeLimit", reflect.TypeOf((*MockAzBlob)(nil).MaxChunkSizeLimit))
}

// MaxSizeLimit mocks base method
func (m *MockAzBlob) MaxSizeLimit() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxSizeLimit")
	ret0, _ := ret[0].(int64)
	return ret0
}

// MaxSizeLimit indicates an expected call of MaxSizeLimit
func (mr *MockAzBlobMockRecorder) MaxSizeLimit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxSizeLimit", reflect.TypeOf((*MockAzBlob)(nil).MaxSizeLimit))
}

// Offset mocks base method
func (m *MockAzBlob) Offset(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Offset", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Offset indicates an expected call of Offset
func (mr *MockAzBlobMockRecorder) Offset(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Offset", reflect.TypeOf((*MockAzBlob)(nil).Offset), arg0)
}

// Upload mocks base method
func (m *MockAzBlob) Upload(arg0 context.Context, arg1 io.ReadSeeker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload
func (mr *MockAzBlobMockRecorder) Upload(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockAzBlob)(nil).Upload), arg0, arg1)
}

// MockBlockBlob is a mock of BlockBlob interface
type MockBlockBlob struct {
	ctrl     *gomock.Controller
	recorder *MockBlockBlobMockRecorder
}

// MockBlockBlobMockRecorder is the mock recorder for MockBlockBlob
type MockBlockBlobMockRecorder struct {
	mock *MockBlockBlob
}

// NewMockBlockBlob creates a new mock instance
func NewMockBlockBlob(ctrl *gomock.Controller) *MockBlockBlob {
	mock := &MockBlockBlob{ctrl: ctrl}
	mock.recorder = &MockBlockBlobMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBlockBlob) EXPECT() *MockBlockBlobMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockBlockBlob) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockBlockBlobMockRecorder) Close(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockBlockBlob)(nil).Close), arg0)
}

// Delete mocks base method
func (m *MockBlockBlob) Delete(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockBlockBlobMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBlockBlob)(nil).Delete), arg0)
}

// Download mocks base method
func (m *MockBlockBlob) Download(arg0 context.Context) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download
func (mr *MockBlockBlobMockRecorder) Download(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockBlockBlob)(nil).Download), arg0)
}

// Exists mocks base method
func (m *MockBlockBlob) Exists(arg0 context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists
func (mr *MockBlockBlobMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockBlockBlob)(nil).Exists), arg0)
}

// GetUncommittedIndexes mocks base method
func (m *MockBlockBlob) GetUncommittedIndexes(arg0 context.Context) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUncommittedIndexes", arg0)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUncommittedIndexes indicates an expected call of GetUncommittedIndexes
func (mr *MockBlockBlobMockRecorder) GetUncommittedIndexes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUncommittedIndexes", reflect.TypeOf((*MockBlockBlob)(nil).GetUncommittedIndexes), arg0)
}

// MaxChunkSizeLimit mocks base method
func (m *MockBlockBlob) MaxChunkSizeLimit() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxChunkSizeLimit")
	ret0, _ := ret[0].(int64)
	return ret0
}

// MaxChunkSizeLimit indicates an expected call of MaxChunkSizeLimit
func (mr *MockBlockBlobMockRecorder) MaxChunkSizeLimit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxChunkSizeLimit", reflect.TypeOf((*MockBlockBlob)(nil).MaxChunkSizeLimit))
}

// MaxSizeLimit mocks base method
func (m *MockBlockBlob) MaxSizeLimit() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxSizeLimit")
	ret0, _ := ret[0].(int64)
	return ret0
}

// MaxSizeLimit indicates an expected call of MaxSizeLimit
func (mr *MockBlockBlobMockRecorder) MaxSizeLimit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxSizeLimit", reflect.TypeOf((*MockBlockBlob)(nil).MaxSizeLimit))
}

// Offset mocks base method
func (m *MockBlockBlob) Offset(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Offset", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Offset indicates an expected call of Offset
func (mr *MockBlockBlobMockRecorder) Offset(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Offset", reflect.TypeOf((*MockBlockBlob)(nil).Offset), arg0)
}

// SetIndexes mocks base method
func (m *MockBlockBlob) SetIndexes(arg0 []int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetIndexes", arg0)
}

// SetIndexes indicates an expected call of SetIndexes
func (mr *MockBlockBlobMockRecorder) SetIndexes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIndexes", reflect.TypeOf((*MockBlockBlob)(nil).SetIndexes), arg0)
}

// Upload mocks base method
func (m *MockBlockBlob) Upload(arg0 context.Context, arg1 io.ReadSeeker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload
func (mr *MockBlockBlobMockRecorder) Upload(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockBlockBlob)(nil).Upload), arg0, arg1)
}

// MockAppendBlob is a mock of AppendBlob interface
type MockAppendBlob struct {
	ctrl     *gomock.Controller
	recorder *MockAppendBlobMockRecorder
}

// MockAppendBlobMockRecorder is the mock recorder for MockAppendBlob
type MockAppendBlobMockRecorder struct {
	mock *MockAppendBlob
}

// NewMockAppendBlob creates a new mock instance
func NewMockAppendBlob(ctrl *gomock.Controller) *MockAppendBlob {
	mock := &MockAppendBlob{ctrl: ctrl}
	mock.recorder = &MockAppendBlobMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppendBlob) EXPECT() *MockAppendBlobMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockAppendBlob) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockAppendBlobMockRecorder) Close(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAppendBlob)(nil).Close), arg0)
}

// Delete mocks base method
func (m *MockAppendBlob) Delete(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockAppendBlobMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAppendBlob)(nil).Delete), arg0)
}

// Download mocks base method
func (m *MockAppendBlob) Download(arg0 context.Context) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download
func (mr *MockAppendBlobMockRecorder) Download(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockAppendBlob)(nil).Download), arg0)
}

// Exists mocks base method
func (m *MockAppendBlob) Exists(arg0 context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists
func (mr *MockAppendBlobMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockAppendBlob)(nil).Exists), arg0)
}

// MaxChunkSizeLimit mocks base method
func (m *MockAppendBlob) MaxChunkSizeLimit() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxChunkSizeLimit")
	ret0, _ := ret[0].(int64)
	return ret0
}

// MaxChunkSizeLimit indicates an expected call of MaxChunkSizeLimit
func (mr *MockAppendBlobMockRecorder) MaxChunkSizeLimit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxChunkSizeLimit", reflect.TypeOf((*MockAppendBlob)(nil).MaxChunkSizeLimit))
}

// MaxSizeLimit mocks base method
func (m *MockAppendBlob) MaxSizeLimit() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxSizeLimit")
	ret0, _ := ret[0].(int64)
	return ret0
}

// MaxSizeLimit indicates an expected call of MaxSizeLimit
func (mr *MockAppendBlobMockRecorder) MaxSizeLimit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxSizeLimit", reflect.TypeOf((*MockAppendBlob)(nil).MaxSizeLimit))
}

// Offset mocks base method
func (m *MockAppendBlob) Offset(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Offset", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Offset indicates an expected call of Offset
func (mr *MockAppendBlobMockRecorder) Offset(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Offset", reflect.TypeOf((*MockAppendBlob)(nil).Offset), arg0)
}

// Upload mocks base method
func (m *MockAppendBlob) Upload(arg0 context.Context, arg1 io.ReadSeeker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload
func (mr *MockAppendBlobMockRecorder) Upload(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockAppendBlob)(nil).Upload), arg0, arg1)
}
