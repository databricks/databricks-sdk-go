// Code generated by mockery v2.39.1. DO NOT EDIT.

package files

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	files "github.com/databricks/databricks-sdk-go/service/files"

	mock "github.com/stretchr/testify/mock"
)

// MockFilesInterface is an autogenerated mock type for the FilesInterface type
type MockFilesInterface struct {
	mock.Mock
}

type MockFilesInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFilesInterface) EXPECT() *MockFilesInterface_Expecter {
	return &MockFilesInterface_Expecter{mock: &_m.Mock}
}

// CreateDirectory provides a mock function with given fields: ctx, request
func (_m *MockFilesInterface) CreateDirectory(ctx context.Context, request files.CreateDirectoryRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for CreateDirectory")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, files.CreateDirectoryRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFilesInterface_CreateDirectory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateDirectory'
type MockFilesInterface_CreateDirectory_Call struct {
	*mock.Call
}

// CreateDirectory is a helper method to define mock.On call
//   - ctx context.Context
//   - request files.CreateDirectoryRequest
func (_e *MockFilesInterface_Expecter) CreateDirectory(ctx interface{}, request interface{}) *MockFilesInterface_CreateDirectory_Call {
	return &MockFilesInterface_CreateDirectory_Call{Call: _e.mock.On("CreateDirectory", ctx, request)}
}

func (_c *MockFilesInterface_CreateDirectory_Call) Run(run func(ctx context.Context, request files.CreateDirectoryRequest)) *MockFilesInterface_CreateDirectory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(files.CreateDirectoryRequest))
	})
	return _c
}

func (_c *MockFilesInterface_CreateDirectory_Call) Return(_a0 error) *MockFilesInterface_CreateDirectory_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilesInterface_CreateDirectory_Call) RunAndReturn(run func(context.Context, files.CreateDirectoryRequest) error) *MockFilesInterface_CreateDirectory_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockFilesInterface) Delete(ctx context.Context, request files.DeleteFileRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, files.DeleteFileRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFilesInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockFilesInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request files.DeleteFileRequest
func (_e *MockFilesInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockFilesInterface_Delete_Call {
	return &MockFilesInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockFilesInterface_Delete_Call) Run(run func(ctx context.Context, request files.DeleteFileRequest)) *MockFilesInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(files.DeleteFileRequest))
	})
	return _c
}

func (_c *MockFilesInterface_Delete_Call) Return(_a0 error) *MockFilesInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilesInterface_Delete_Call) RunAndReturn(run func(context.Context, files.DeleteFileRequest) error) *MockFilesInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByFilePath provides a mock function with given fields: ctx, filePath
func (_m *MockFilesInterface) DeleteByFilePath(ctx context.Context, filePath string) error {
	ret := _m.Called(ctx, filePath)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByFilePath")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, filePath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFilesInterface_DeleteByFilePath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByFilePath'
type MockFilesInterface_DeleteByFilePath_Call struct {
	*mock.Call
}

// DeleteByFilePath is a helper method to define mock.On call
//   - ctx context.Context
//   - filePath string
func (_e *MockFilesInterface_Expecter) DeleteByFilePath(ctx interface{}, filePath interface{}) *MockFilesInterface_DeleteByFilePath_Call {
	return &MockFilesInterface_DeleteByFilePath_Call{Call: _e.mock.On("DeleteByFilePath", ctx, filePath)}
}

func (_c *MockFilesInterface_DeleteByFilePath_Call) Run(run func(ctx context.Context, filePath string)) *MockFilesInterface_DeleteByFilePath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockFilesInterface_DeleteByFilePath_Call) Return(_a0 error) *MockFilesInterface_DeleteByFilePath_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilesInterface_DeleteByFilePath_Call) RunAndReturn(run func(context.Context, string) error) *MockFilesInterface_DeleteByFilePath_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteDirectory provides a mock function with given fields: ctx, request
func (_m *MockFilesInterface) DeleteDirectory(ctx context.Context, request files.DeleteDirectoryRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDirectory")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, files.DeleteDirectoryRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFilesInterface_DeleteDirectory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteDirectory'
type MockFilesInterface_DeleteDirectory_Call struct {
	*mock.Call
}

// DeleteDirectory is a helper method to define mock.On call
//   - ctx context.Context
//   - request files.DeleteDirectoryRequest
func (_e *MockFilesInterface_Expecter) DeleteDirectory(ctx interface{}, request interface{}) *MockFilesInterface_DeleteDirectory_Call {
	return &MockFilesInterface_DeleteDirectory_Call{Call: _e.mock.On("DeleteDirectory", ctx, request)}
}

func (_c *MockFilesInterface_DeleteDirectory_Call) Run(run func(ctx context.Context, request files.DeleteDirectoryRequest)) *MockFilesInterface_DeleteDirectory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(files.DeleteDirectoryRequest))
	})
	return _c
}

func (_c *MockFilesInterface_DeleteDirectory_Call) Return(_a0 error) *MockFilesInterface_DeleteDirectory_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilesInterface_DeleteDirectory_Call) RunAndReturn(run func(context.Context, files.DeleteDirectoryRequest) error) *MockFilesInterface_DeleteDirectory_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteDirectoryByDirectoryPath provides a mock function with given fields: ctx, directoryPath
func (_m *MockFilesInterface) DeleteDirectoryByDirectoryPath(ctx context.Context, directoryPath string) error {
	ret := _m.Called(ctx, directoryPath)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDirectoryByDirectoryPath")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, directoryPath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFilesInterface_DeleteDirectoryByDirectoryPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteDirectoryByDirectoryPath'
type MockFilesInterface_DeleteDirectoryByDirectoryPath_Call struct {
	*mock.Call
}

// DeleteDirectoryByDirectoryPath is a helper method to define mock.On call
//   - ctx context.Context
//   - directoryPath string
func (_e *MockFilesInterface_Expecter) DeleteDirectoryByDirectoryPath(ctx interface{}, directoryPath interface{}) *MockFilesInterface_DeleteDirectoryByDirectoryPath_Call {
	return &MockFilesInterface_DeleteDirectoryByDirectoryPath_Call{Call: _e.mock.On("DeleteDirectoryByDirectoryPath", ctx, directoryPath)}
}

func (_c *MockFilesInterface_DeleteDirectoryByDirectoryPath_Call) Run(run func(ctx context.Context, directoryPath string)) *MockFilesInterface_DeleteDirectoryByDirectoryPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockFilesInterface_DeleteDirectoryByDirectoryPath_Call) Return(_a0 error) *MockFilesInterface_DeleteDirectoryByDirectoryPath_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilesInterface_DeleteDirectoryByDirectoryPath_Call) RunAndReturn(run func(context.Context, string) error) *MockFilesInterface_DeleteDirectoryByDirectoryPath_Call {
	_c.Call.Return(run)
	return _c
}

// Download provides a mock function with given fields: ctx, request
func (_m *MockFilesInterface) Download(ctx context.Context, request files.DownloadRequest) (*files.DownloadResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Download")
	}

	var r0 *files.DownloadResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, files.DownloadRequest) (*files.DownloadResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, files.DownloadRequest) *files.DownloadResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*files.DownloadResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, files.DownloadRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFilesInterface_Download_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Download'
type MockFilesInterface_Download_Call struct {
	*mock.Call
}

// Download is a helper method to define mock.On call
//   - ctx context.Context
//   - request files.DownloadRequest
func (_e *MockFilesInterface_Expecter) Download(ctx interface{}, request interface{}) *MockFilesInterface_Download_Call {
	return &MockFilesInterface_Download_Call{Call: _e.mock.On("Download", ctx, request)}
}

func (_c *MockFilesInterface_Download_Call) Run(run func(ctx context.Context, request files.DownloadRequest)) *MockFilesInterface_Download_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(files.DownloadRequest))
	})
	return _c
}

func (_c *MockFilesInterface_Download_Call) Return(_a0 *files.DownloadResponse, _a1 error) *MockFilesInterface_Download_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFilesInterface_Download_Call) RunAndReturn(run func(context.Context, files.DownloadRequest) (*files.DownloadResponse, error)) *MockFilesInterface_Download_Call {
	_c.Call.Return(run)
	return _c
}

// DownloadByFilePath provides a mock function with given fields: ctx, filePath
func (_m *MockFilesInterface) DownloadByFilePath(ctx context.Context, filePath string) (*files.DownloadResponse, error) {
	ret := _m.Called(ctx, filePath)

	if len(ret) == 0 {
		panic("no return value specified for DownloadByFilePath")
	}

	var r0 *files.DownloadResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*files.DownloadResponse, error)); ok {
		return rf(ctx, filePath)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *files.DownloadResponse); ok {
		r0 = rf(ctx, filePath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*files.DownloadResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFilesInterface_DownloadByFilePath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadByFilePath'
type MockFilesInterface_DownloadByFilePath_Call struct {
	*mock.Call
}

// DownloadByFilePath is a helper method to define mock.On call
//   - ctx context.Context
//   - filePath string
func (_e *MockFilesInterface_Expecter) DownloadByFilePath(ctx interface{}, filePath interface{}) *MockFilesInterface_DownloadByFilePath_Call {
	return &MockFilesInterface_DownloadByFilePath_Call{Call: _e.mock.On("DownloadByFilePath", ctx, filePath)}
}

func (_c *MockFilesInterface_DownloadByFilePath_Call) Run(run func(ctx context.Context, filePath string)) *MockFilesInterface_DownloadByFilePath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockFilesInterface_DownloadByFilePath_Call) Return(_a0 *files.DownloadResponse, _a1 error) *MockFilesInterface_DownloadByFilePath_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFilesInterface_DownloadByFilePath_Call) RunAndReturn(run func(context.Context, string) (*files.DownloadResponse, error)) *MockFilesInterface_DownloadByFilePath_Call {
	_c.Call.Return(run)
	return _c
}

// GetStatus provides a mock function with given fields: ctx, request
func (_m *MockFilesInterface) GetStatus(ctx context.Context, request files.GetStatusRequest) (*files.FileInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetStatus")
	}

	var r0 *files.FileInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, files.GetStatusRequest) (*files.FileInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, files.GetStatusRequest) *files.FileInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*files.FileInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, files.GetStatusRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFilesInterface_GetStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatus'
type MockFilesInterface_GetStatus_Call struct {
	*mock.Call
}

// GetStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - request files.GetStatusRequest
func (_e *MockFilesInterface_Expecter) GetStatus(ctx interface{}, request interface{}) *MockFilesInterface_GetStatus_Call {
	return &MockFilesInterface_GetStatus_Call{Call: _e.mock.On("GetStatus", ctx, request)}
}

func (_c *MockFilesInterface_GetStatus_Call) Run(run func(ctx context.Context, request files.GetStatusRequest)) *MockFilesInterface_GetStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(files.GetStatusRequest))
	})
	return _c
}

func (_c *MockFilesInterface_GetStatus_Call) Return(_a0 *files.FileInfo, _a1 error) *MockFilesInterface_GetStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFilesInterface_GetStatus_Call) RunAndReturn(run func(context.Context, files.GetStatusRequest) (*files.FileInfo, error)) *MockFilesInterface_GetStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockFilesInterface) Impl() files.FilesService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 files.FilesService
	if rf, ok := ret.Get(0).(func() files.FilesService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(files.FilesService)
		}
	}

	return r0
}

// MockFilesInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockFilesInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockFilesInterface_Expecter) Impl() *MockFilesInterface_Impl_Call {
	return &MockFilesInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockFilesInterface_Impl_Call) Run(run func()) *MockFilesInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFilesInterface_Impl_Call) Return(_a0 files.FilesService) *MockFilesInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilesInterface_Impl_Call) RunAndReturn(run func() files.FilesService) *MockFilesInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// ListDirectoryContents provides a mock function with given fields: ctx, request
func (_m *MockFilesInterface) ListDirectoryContents(ctx context.Context, request files.ListDirectoryContentsRequest) listing.Iterator[files.DirectoryEntry] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListDirectoryContents")
	}

	var r0 listing.Iterator[files.DirectoryEntry]
	if rf, ok := ret.Get(0).(func(context.Context, files.ListDirectoryContentsRequest) listing.Iterator[files.DirectoryEntry]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[files.DirectoryEntry])
		}
	}

	return r0
}

// MockFilesInterface_ListDirectoryContents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListDirectoryContents'
type MockFilesInterface_ListDirectoryContents_Call struct {
	*mock.Call
}

// ListDirectoryContents is a helper method to define mock.On call
//   - ctx context.Context
//   - request files.ListDirectoryContentsRequest
func (_e *MockFilesInterface_Expecter) ListDirectoryContents(ctx interface{}, request interface{}) *MockFilesInterface_ListDirectoryContents_Call {
	return &MockFilesInterface_ListDirectoryContents_Call{Call: _e.mock.On("ListDirectoryContents", ctx, request)}
}

func (_c *MockFilesInterface_ListDirectoryContents_Call) Run(run func(ctx context.Context, request files.ListDirectoryContentsRequest)) *MockFilesInterface_ListDirectoryContents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(files.ListDirectoryContentsRequest))
	})
	return _c
}

func (_c *MockFilesInterface_ListDirectoryContents_Call) Return(_a0 listing.Iterator[files.DirectoryEntry]) *MockFilesInterface_ListDirectoryContents_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilesInterface_ListDirectoryContents_Call) RunAndReturn(run func(context.Context, files.ListDirectoryContentsRequest) listing.Iterator[files.DirectoryEntry]) *MockFilesInterface_ListDirectoryContents_Call {
	_c.Call.Return(run)
	return _c
}

// ListDirectoryContentsAll provides a mock function with given fields: ctx, request
func (_m *MockFilesInterface) ListDirectoryContentsAll(ctx context.Context, request files.ListDirectoryContentsRequest) ([]files.DirectoryEntry, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListDirectoryContentsAll")
	}

	var r0 []files.DirectoryEntry
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, files.ListDirectoryContentsRequest) ([]files.DirectoryEntry, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, files.ListDirectoryContentsRequest) []files.DirectoryEntry); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]files.DirectoryEntry)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, files.ListDirectoryContentsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFilesInterface_ListDirectoryContentsAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListDirectoryContentsAll'
type MockFilesInterface_ListDirectoryContentsAll_Call struct {
	*mock.Call
}

// ListDirectoryContentsAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request files.ListDirectoryContentsRequest
func (_e *MockFilesInterface_Expecter) ListDirectoryContentsAll(ctx interface{}, request interface{}) *MockFilesInterface_ListDirectoryContentsAll_Call {
	return &MockFilesInterface_ListDirectoryContentsAll_Call{Call: _e.mock.On("ListDirectoryContentsAll", ctx, request)}
}

func (_c *MockFilesInterface_ListDirectoryContentsAll_Call) Run(run func(ctx context.Context, request files.ListDirectoryContentsRequest)) *MockFilesInterface_ListDirectoryContentsAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(files.ListDirectoryContentsRequest))
	})
	return _c
}

func (_c *MockFilesInterface_ListDirectoryContentsAll_Call) Return(_a0 []files.DirectoryEntry, _a1 error) *MockFilesInterface_ListDirectoryContentsAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFilesInterface_ListDirectoryContentsAll_Call) RunAndReturn(run func(context.Context, files.ListDirectoryContentsRequest) ([]files.DirectoryEntry, error)) *MockFilesInterface_ListDirectoryContentsAll_Call {
	_c.Call.Return(run)
	return _c
}

// ListDirectoryContentsByDirectoryPath provides a mock function with given fields: ctx, directoryPath
func (_m *MockFilesInterface) ListDirectoryContentsByDirectoryPath(ctx context.Context, directoryPath string) (*files.ListDirectoryResponse, error) {
	ret := _m.Called(ctx, directoryPath)

	if len(ret) == 0 {
		panic("no return value specified for ListDirectoryContentsByDirectoryPath")
	}

	var r0 *files.ListDirectoryResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*files.ListDirectoryResponse, error)); ok {
		return rf(ctx, directoryPath)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *files.ListDirectoryResponse); ok {
		r0 = rf(ctx, directoryPath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*files.ListDirectoryResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, directoryPath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFilesInterface_ListDirectoryContentsByDirectoryPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListDirectoryContentsByDirectoryPath'
type MockFilesInterface_ListDirectoryContentsByDirectoryPath_Call struct {
	*mock.Call
}

// ListDirectoryContentsByDirectoryPath is a helper method to define mock.On call
//   - ctx context.Context
//   - directoryPath string
func (_e *MockFilesInterface_Expecter) ListDirectoryContentsByDirectoryPath(ctx interface{}, directoryPath interface{}) *MockFilesInterface_ListDirectoryContentsByDirectoryPath_Call {
	return &MockFilesInterface_ListDirectoryContentsByDirectoryPath_Call{Call: _e.mock.On("ListDirectoryContentsByDirectoryPath", ctx, directoryPath)}
}

func (_c *MockFilesInterface_ListDirectoryContentsByDirectoryPath_Call) Run(run func(ctx context.Context, directoryPath string)) *MockFilesInterface_ListDirectoryContentsByDirectoryPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockFilesInterface_ListDirectoryContentsByDirectoryPath_Call) Return(_a0 *files.ListDirectoryResponse, _a1 error) *MockFilesInterface_ListDirectoryContentsByDirectoryPath_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFilesInterface_ListDirectoryContentsByDirectoryPath_Call) RunAndReturn(run func(context.Context, string) (*files.ListDirectoryResponse, error)) *MockFilesInterface_ListDirectoryContentsByDirectoryPath_Call {
	_c.Call.Return(run)
	return _c
}

// Upload provides a mock function with given fields: ctx, request
func (_m *MockFilesInterface) Upload(ctx context.Context, request files.UploadRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Upload")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, files.UploadRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFilesInterface_Upload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Upload'
type MockFilesInterface_Upload_Call struct {
	*mock.Call
}

// Upload is a helper method to define mock.On call
//   - ctx context.Context
//   - request files.UploadRequest
func (_e *MockFilesInterface_Expecter) Upload(ctx interface{}, request interface{}) *MockFilesInterface_Upload_Call {
	return &MockFilesInterface_Upload_Call{Call: _e.mock.On("Upload", ctx, request)}
}

func (_c *MockFilesInterface_Upload_Call) Run(run func(ctx context.Context, request files.UploadRequest)) *MockFilesInterface_Upload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(files.UploadRequest))
	})
	return _c
}

func (_c *MockFilesInterface_Upload_Call) Return(_a0 error) *MockFilesInterface_Upload_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilesInterface_Upload_Call) RunAndReturn(run func(context.Context, files.UploadRequest) error) *MockFilesInterface_Upload_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockFilesInterface) WithImpl(impl files.FilesService) files.FilesInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 files.FilesInterface
	if rf, ok := ret.Get(0).(func(files.FilesService) files.FilesInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(files.FilesInterface)
		}
	}

	return r0
}

// MockFilesInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockFilesInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl files.FilesService
func (_e *MockFilesInterface_Expecter) WithImpl(impl interface{}) *MockFilesInterface_WithImpl_Call {
	return &MockFilesInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockFilesInterface_WithImpl_Call) Run(run func(impl files.FilesService)) *MockFilesInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(files.FilesService))
	})
	return _c
}

func (_c *MockFilesInterface_WithImpl_Call) Return(_a0 files.FilesInterface) *MockFilesInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFilesInterface_WithImpl_Call) RunAndReturn(run func(files.FilesService) files.FilesInterface) *MockFilesInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFilesInterface creates a new instance of MockFilesInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFilesInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFilesInterface {
	mock := &MockFilesInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
