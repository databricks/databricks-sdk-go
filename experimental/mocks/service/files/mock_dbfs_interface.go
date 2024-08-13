// Code generated by mockery v2.43.0. DO NOT EDIT.

package files

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	files "github.com/databricks/databricks-sdk-go/service/files"

	mock "github.com/stretchr/testify/mock"
)

// MockDbfsInterface is an autogenerated mock type for the DbfsInterface type
type MockDbfsInterface struct {
	mock.Mock
}

type MockDbfsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDbfsInterface) EXPECT() *MockDbfsInterface_Expecter {
	return &MockDbfsInterface_Expecter{mock: &_m.Mock}
}

// AddBlock provides a mock function with given fields: ctx, request
func (_m *MockDbfsInterface) AddBlock(ctx context.Context, request files.AddBlock) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for AddBlock")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, files.AddBlock) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDbfsInterface_AddBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddBlock'
type MockDbfsInterface_AddBlock_Call struct {
	*mock.Call
}

// AddBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - request files.AddBlock
func (_e *MockDbfsInterface_Expecter) AddBlock(ctx interface{}, request interface{}) *MockDbfsInterface_AddBlock_Call {
	return &MockDbfsInterface_AddBlock_Call{Call: _e.mock.On("AddBlock", ctx, request)}
}

func (_c *MockDbfsInterface_AddBlock_Call) Run(run func(ctx context.Context, request files.AddBlock)) *MockDbfsInterface_AddBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(files.AddBlock))
	})
	return _c
}

func (_c *MockDbfsInterface_AddBlock_Call) Return(_a0 error) *MockDbfsInterface_AddBlock_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDbfsInterface_AddBlock_Call) RunAndReturn(run func(context.Context, files.AddBlock) error) *MockDbfsInterface_AddBlock_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields: ctx, request
func (_m *MockDbfsInterface) Close(ctx context.Context, request files.Close) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, files.Close) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDbfsInterface_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockDbfsInterface_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - ctx context.Context
//   - request files.Close
func (_e *MockDbfsInterface_Expecter) Close(ctx interface{}, request interface{}) *MockDbfsInterface_Close_Call {
	return &MockDbfsInterface_Close_Call{Call: _e.mock.On("Close", ctx, request)}
}

func (_c *MockDbfsInterface_Close_Call) Run(run func(ctx context.Context, request files.Close)) *MockDbfsInterface_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(files.Close))
	})
	return _c
}

func (_c *MockDbfsInterface_Close_Call) Return(_a0 error) *MockDbfsInterface_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDbfsInterface_Close_Call) RunAndReturn(run func(context.Context, files.Close) error) *MockDbfsInterface_Close_Call {
	_c.Call.Return(run)
	return _c
}

// CloseByHandle provides a mock function with given fields: ctx, handle
func (_m *MockDbfsInterface) CloseByHandle(ctx context.Context, handle int64) error {
	ret := _m.Called(ctx, handle)

	if len(ret) == 0 {
		panic("no return value specified for CloseByHandle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, handle)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDbfsInterface_CloseByHandle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseByHandle'
type MockDbfsInterface_CloseByHandle_Call struct {
	*mock.Call
}

// CloseByHandle is a helper method to define mock.On call
//   - ctx context.Context
//   - handle int64
func (_e *MockDbfsInterface_Expecter) CloseByHandle(ctx interface{}, handle interface{}) *MockDbfsInterface_CloseByHandle_Call {
	return &MockDbfsInterface_CloseByHandle_Call{Call: _e.mock.On("CloseByHandle", ctx, handle)}
}

func (_c *MockDbfsInterface_CloseByHandle_Call) Run(run func(ctx context.Context, handle int64)) *MockDbfsInterface_CloseByHandle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockDbfsInterface_CloseByHandle_Call) Return(_a0 error) *MockDbfsInterface_CloseByHandle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDbfsInterface_CloseByHandle_Call) RunAndReturn(run func(context.Context, int64) error) *MockDbfsInterface_CloseByHandle_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockDbfsInterface) Create(ctx context.Context, request files.Create) (*files.CreateResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *files.CreateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, files.Create) (*files.CreateResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, files.Create) *files.CreateResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*files.CreateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, files.Create) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDbfsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockDbfsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request files.Create
func (_e *MockDbfsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockDbfsInterface_Create_Call {
	return &MockDbfsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockDbfsInterface_Create_Call) Run(run func(ctx context.Context, request files.Create)) *MockDbfsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(files.Create))
	})
	return _c
}

func (_c *MockDbfsInterface_Create_Call) Return(_a0 *files.CreateResponse, _a1 error) *MockDbfsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDbfsInterface_Create_Call) RunAndReturn(run func(context.Context, files.Create) (*files.CreateResponse, error)) *MockDbfsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockDbfsInterface) Delete(ctx context.Context, request files.Delete) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, files.Delete) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDbfsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockDbfsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request files.Delete
func (_e *MockDbfsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockDbfsInterface_Delete_Call {
	return &MockDbfsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockDbfsInterface_Delete_Call) Run(run func(ctx context.Context, request files.Delete)) *MockDbfsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(files.Delete))
	})
	return _c
}

func (_c *MockDbfsInterface_Delete_Call) Return(_a0 error) *MockDbfsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDbfsInterface_Delete_Call) RunAndReturn(run func(context.Context, files.Delete) error) *MockDbfsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetStatus provides a mock function with given fields: ctx, request
func (_m *MockDbfsInterface) GetStatus(ctx context.Context, request files.GetStatusRequest) (*files.FileInfo, error) {
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

// MockDbfsInterface_GetStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatus'
type MockDbfsInterface_GetStatus_Call struct {
	*mock.Call
}

// GetStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - request files.GetStatusRequest
func (_e *MockDbfsInterface_Expecter) GetStatus(ctx interface{}, request interface{}) *MockDbfsInterface_GetStatus_Call {
	return &MockDbfsInterface_GetStatus_Call{Call: _e.mock.On("GetStatus", ctx, request)}
}

func (_c *MockDbfsInterface_GetStatus_Call) Run(run func(ctx context.Context, request files.GetStatusRequest)) *MockDbfsInterface_GetStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(files.GetStatusRequest))
	})
	return _c
}

func (_c *MockDbfsInterface_GetStatus_Call) Return(_a0 *files.FileInfo, _a1 error) *MockDbfsInterface_GetStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDbfsInterface_GetStatus_Call) RunAndReturn(run func(context.Context, files.GetStatusRequest) (*files.FileInfo, error)) *MockDbfsInterface_GetStatus_Call {
	_c.Call.Return(run)
	return _c
}

// GetStatusByPath provides a mock function with given fields: ctx, path
func (_m *MockDbfsInterface) GetStatusByPath(ctx context.Context, path string) (*files.FileInfo, error) {
	ret := _m.Called(ctx, path)

	if len(ret) == 0 {
		panic("no return value specified for GetStatusByPath")
	}

	var r0 *files.FileInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*files.FileInfo, error)); ok {
		return rf(ctx, path)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *files.FileInfo); ok {
		r0 = rf(ctx, path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*files.FileInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDbfsInterface_GetStatusByPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatusByPath'
type MockDbfsInterface_GetStatusByPath_Call struct {
	*mock.Call
}

// GetStatusByPath is a helper method to define mock.On call
//   - ctx context.Context
//   - path string
func (_e *MockDbfsInterface_Expecter) GetStatusByPath(ctx interface{}, path interface{}) *MockDbfsInterface_GetStatusByPath_Call {
	return &MockDbfsInterface_GetStatusByPath_Call{Call: _e.mock.On("GetStatusByPath", ctx, path)}
}

func (_c *MockDbfsInterface_GetStatusByPath_Call) Run(run func(ctx context.Context, path string)) *MockDbfsInterface_GetStatusByPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockDbfsInterface_GetStatusByPath_Call) Return(_a0 *files.FileInfo, _a1 error) *MockDbfsInterface_GetStatusByPath_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDbfsInterface_GetStatusByPath_Call) RunAndReturn(run func(context.Context, string) (*files.FileInfo, error)) *MockDbfsInterface_GetStatusByPath_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockDbfsInterface) List(ctx context.Context, request files.ListDbfsRequest) listing.Iterator[files.FileInfo] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[files.FileInfo]
	if rf, ok := ret.Get(0).(func(context.Context, files.ListDbfsRequest) listing.Iterator[files.FileInfo]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[files.FileInfo])
		}
	}

	return r0
}

// MockDbfsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockDbfsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request files.ListDbfsRequest
func (_e *MockDbfsInterface_Expecter) List(ctx interface{}, request interface{}) *MockDbfsInterface_List_Call {
	return &MockDbfsInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockDbfsInterface_List_Call) Run(run func(ctx context.Context, request files.ListDbfsRequest)) *MockDbfsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(files.ListDbfsRequest))
	})
	return _c
}

func (_c *MockDbfsInterface_List_Call) Return(_a0 listing.Iterator[files.FileInfo]) *MockDbfsInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDbfsInterface_List_Call) RunAndReturn(run func(context.Context, files.ListDbfsRequest) listing.Iterator[files.FileInfo]) *MockDbfsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockDbfsInterface) ListAll(ctx context.Context, request files.ListDbfsRequest) ([]files.FileInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []files.FileInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, files.ListDbfsRequest) ([]files.FileInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, files.ListDbfsRequest) []files.FileInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]files.FileInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, files.ListDbfsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDbfsInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockDbfsInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request files.ListDbfsRequest
func (_e *MockDbfsInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockDbfsInterface_ListAll_Call {
	return &MockDbfsInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockDbfsInterface_ListAll_Call) Run(run func(ctx context.Context, request files.ListDbfsRequest)) *MockDbfsInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(files.ListDbfsRequest))
	})
	return _c
}

func (_c *MockDbfsInterface_ListAll_Call) Return(_a0 []files.FileInfo, _a1 error) *MockDbfsInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDbfsInterface_ListAll_Call) RunAndReturn(run func(context.Context, files.ListDbfsRequest) ([]files.FileInfo, error)) *MockDbfsInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// ListByPath provides a mock function with given fields: ctx, path
func (_m *MockDbfsInterface) ListByPath(ctx context.Context, path string) (*files.ListStatusResponse, error) {
	ret := _m.Called(ctx, path)

	if len(ret) == 0 {
		panic("no return value specified for ListByPath")
	}

	var r0 *files.ListStatusResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*files.ListStatusResponse, error)); ok {
		return rf(ctx, path)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *files.ListStatusResponse); ok {
		r0 = rf(ctx, path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*files.ListStatusResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDbfsInterface_ListByPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListByPath'
type MockDbfsInterface_ListByPath_Call struct {
	*mock.Call
}

// ListByPath is a helper method to define mock.On call
//   - ctx context.Context
//   - path string
func (_e *MockDbfsInterface_Expecter) ListByPath(ctx interface{}, path interface{}) *MockDbfsInterface_ListByPath_Call {
	return &MockDbfsInterface_ListByPath_Call{Call: _e.mock.On("ListByPath", ctx, path)}
}

func (_c *MockDbfsInterface_ListByPath_Call) Run(run func(ctx context.Context, path string)) *MockDbfsInterface_ListByPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockDbfsInterface_ListByPath_Call) Return(_a0 *files.ListStatusResponse, _a1 error) *MockDbfsInterface_ListByPath_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDbfsInterface_ListByPath_Call) RunAndReturn(run func(context.Context, string) (*files.ListStatusResponse, error)) *MockDbfsInterface_ListByPath_Call {
	_c.Call.Return(run)
	return _c
}

// Mkdirs provides a mock function with given fields: ctx, request
func (_m *MockDbfsInterface) Mkdirs(ctx context.Context, request files.MkDirs) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Mkdirs")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, files.MkDirs) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDbfsInterface_Mkdirs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Mkdirs'
type MockDbfsInterface_Mkdirs_Call struct {
	*mock.Call
}

// Mkdirs is a helper method to define mock.On call
//   - ctx context.Context
//   - request files.MkDirs
func (_e *MockDbfsInterface_Expecter) Mkdirs(ctx interface{}, request interface{}) *MockDbfsInterface_Mkdirs_Call {
	return &MockDbfsInterface_Mkdirs_Call{Call: _e.mock.On("Mkdirs", ctx, request)}
}

func (_c *MockDbfsInterface_Mkdirs_Call) Run(run func(ctx context.Context, request files.MkDirs)) *MockDbfsInterface_Mkdirs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(files.MkDirs))
	})
	return _c
}

func (_c *MockDbfsInterface_Mkdirs_Call) Return(_a0 error) *MockDbfsInterface_Mkdirs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDbfsInterface_Mkdirs_Call) RunAndReturn(run func(context.Context, files.MkDirs) error) *MockDbfsInterface_Mkdirs_Call {
	_c.Call.Return(run)
	return _c
}

// MkdirsByPath provides a mock function with given fields: ctx, path
func (_m *MockDbfsInterface) MkdirsByPath(ctx context.Context, path string) error {
	ret := _m.Called(ctx, path)

	if len(ret) == 0 {
		panic("no return value specified for MkdirsByPath")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDbfsInterface_MkdirsByPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MkdirsByPath'
type MockDbfsInterface_MkdirsByPath_Call struct {
	*mock.Call
}

// MkdirsByPath is a helper method to define mock.On call
//   - ctx context.Context
//   - path string
func (_e *MockDbfsInterface_Expecter) MkdirsByPath(ctx interface{}, path interface{}) *MockDbfsInterface_MkdirsByPath_Call {
	return &MockDbfsInterface_MkdirsByPath_Call{Call: _e.mock.On("MkdirsByPath", ctx, path)}
}

func (_c *MockDbfsInterface_MkdirsByPath_Call) Run(run func(ctx context.Context, path string)) *MockDbfsInterface_MkdirsByPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockDbfsInterface_MkdirsByPath_Call) Return(_a0 error) *MockDbfsInterface_MkdirsByPath_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDbfsInterface_MkdirsByPath_Call) RunAndReturn(run func(context.Context, string) error) *MockDbfsInterface_MkdirsByPath_Call {
	_c.Call.Return(run)
	return _c
}

// Move provides a mock function with given fields: ctx, request
func (_m *MockDbfsInterface) Move(ctx context.Context, request files.Move) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Move")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, files.Move) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDbfsInterface_Move_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Move'
type MockDbfsInterface_Move_Call struct {
	*mock.Call
}

// Move is a helper method to define mock.On call
//   - ctx context.Context
//   - request files.Move
func (_e *MockDbfsInterface_Expecter) Move(ctx interface{}, request interface{}) *MockDbfsInterface_Move_Call {
	return &MockDbfsInterface_Move_Call{Call: _e.mock.On("Move", ctx, request)}
}

func (_c *MockDbfsInterface_Move_Call) Run(run func(ctx context.Context, request files.Move)) *MockDbfsInterface_Move_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(files.Move))
	})
	return _c
}

func (_c *MockDbfsInterface_Move_Call) Return(_a0 error) *MockDbfsInterface_Move_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDbfsInterface_Move_Call) RunAndReturn(run func(context.Context, files.Move) error) *MockDbfsInterface_Move_Call {
	_c.Call.Return(run)
	return _c
}

// Open provides a mock function with given fields: ctx, path, mode
func (_m *MockDbfsInterface) Open(ctx context.Context, path string, mode files.FileMode) (files.Handle, error) {
	ret := _m.Called(ctx, path, mode)

	if len(ret) == 0 {
		panic("no return value specified for Open")
	}

	var r0 files.Handle
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, files.FileMode) (files.Handle, error)); ok {
		return rf(ctx, path, mode)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, files.FileMode) files.Handle); ok {
		r0 = rf(ctx, path, mode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(files.Handle)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, files.FileMode) error); ok {
		r1 = rf(ctx, path, mode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDbfsInterface_Open_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Open'
type MockDbfsInterface_Open_Call struct {
	*mock.Call
}

// Open is a helper method to define mock.On call
//   - ctx context.Context
//   - path string
//   - mode files.FileMode
func (_e *MockDbfsInterface_Expecter) Open(ctx interface{}, path interface{}, mode interface{}) *MockDbfsInterface_Open_Call {
	return &MockDbfsInterface_Open_Call{Call: _e.mock.On("Open", ctx, path, mode)}
}

func (_c *MockDbfsInterface_Open_Call) Run(run func(ctx context.Context, path string, mode files.FileMode)) *MockDbfsInterface_Open_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(files.FileMode))
	})
	return _c
}

func (_c *MockDbfsInterface_Open_Call) Return(_a0 files.Handle, _a1 error) *MockDbfsInterface_Open_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDbfsInterface_Open_Call) RunAndReturn(run func(context.Context, string, files.FileMode) (files.Handle, error)) *MockDbfsInterface_Open_Call {
	_c.Call.Return(run)
	return _c
}

// Put provides a mock function with given fields: ctx, request
func (_m *MockDbfsInterface) Put(ctx context.Context, request files.Put) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Put")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, files.Put) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDbfsInterface_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type MockDbfsInterface_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - ctx context.Context
//   - request files.Put
func (_e *MockDbfsInterface_Expecter) Put(ctx interface{}, request interface{}) *MockDbfsInterface_Put_Call {
	return &MockDbfsInterface_Put_Call{Call: _e.mock.On("Put", ctx, request)}
}

func (_c *MockDbfsInterface_Put_Call) Run(run func(ctx context.Context, request files.Put)) *MockDbfsInterface_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(files.Put))
	})
	return _c
}

func (_c *MockDbfsInterface_Put_Call) Return(_a0 error) *MockDbfsInterface_Put_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDbfsInterface_Put_Call) RunAndReturn(run func(context.Context, files.Put) error) *MockDbfsInterface_Put_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: ctx, request
func (_m *MockDbfsInterface) Read(ctx context.Context, request files.ReadDbfsRequest) (*files.ReadResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 *files.ReadResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, files.ReadDbfsRequest) (*files.ReadResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, files.ReadDbfsRequest) *files.ReadResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*files.ReadResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, files.ReadDbfsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDbfsInterface_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type MockDbfsInterface_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - ctx context.Context
//   - request files.ReadDbfsRequest
func (_e *MockDbfsInterface_Expecter) Read(ctx interface{}, request interface{}) *MockDbfsInterface_Read_Call {
	return &MockDbfsInterface_Read_Call{Call: _e.mock.On("Read", ctx, request)}
}

func (_c *MockDbfsInterface_Read_Call) Run(run func(ctx context.Context, request files.ReadDbfsRequest)) *MockDbfsInterface_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(files.ReadDbfsRequest))
	})
	return _c
}

func (_c *MockDbfsInterface_Read_Call) Return(_a0 *files.ReadResponse, _a1 error) *MockDbfsInterface_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDbfsInterface_Read_Call) RunAndReturn(run func(context.Context, files.ReadDbfsRequest) (*files.ReadResponse, error)) *MockDbfsInterface_Read_Call {
	_c.Call.Return(run)
	return _c
}

// ReadFile provides a mock function with given fields: ctx, name
func (_m *MockDbfsInterface) ReadFile(ctx context.Context, name string) ([]byte, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for ReadFile")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]byte, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []byte); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDbfsInterface_ReadFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadFile'
type MockDbfsInterface_ReadFile_Call struct {
	*mock.Call
}

// ReadFile is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockDbfsInterface_Expecter) ReadFile(ctx interface{}, name interface{}) *MockDbfsInterface_ReadFile_Call {
	return &MockDbfsInterface_ReadFile_Call{Call: _e.mock.On("ReadFile", ctx, name)}
}

func (_c *MockDbfsInterface_ReadFile_Call) Run(run func(ctx context.Context, name string)) *MockDbfsInterface_ReadFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockDbfsInterface_ReadFile_Call) Return(_a0 []byte, _a1 error) *MockDbfsInterface_ReadFile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDbfsInterface_ReadFile_Call) RunAndReturn(run func(context.Context, string) ([]byte, error)) *MockDbfsInterface_ReadFile_Call {
	_c.Call.Return(run)
	return _c
}

// RecursiveList provides a mock function with given fields: ctx, path
func (_m *MockDbfsInterface) RecursiveList(ctx context.Context, path string) ([]files.FileInfo, error) {
	ret := _m.Called(ctx, path)

	if len(ret) == 0 {
		panic("no return value specified for RecursiveList")
	}

	var r0 []files.FileInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]files.FileInfo, error)); ok {
		return rf(ctx, path)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []files.FileInfo); ok {
		r0 = rf(ctx, path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]files.FileInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDbfsInterface_RecursiveList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecursiveList'
type MockDbfsInterface_RecursiveList_Call struct {
	*mock.Call
}

// RecursiveList is a helper method to define mock.On call
//   - ctx context.Context
//   - path string
func (_e *MockDbfsInterface_Expecter) RecursiveList(ctx interface{}, path interface{}) *MockDbfsInterface_RecursiveList_Call {
	return &MockDbfsInterface_RecursiveList_Call{Call: _e.mock.On("RecursiveList", ctx, path)}
}

func (_c *MockDbfsInterface_RecursiveList_Call) Run(run func(ctx context.Context, path string)) *MockDbfsInterface_RecursiveList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockDbfsInterface_RecursiveList_Call) Return(_a0 []files.FileInfo, _a1 error) *MockDbfsInterface_RecursiveList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDbfsInterface_RecursiveList_Call) RunAndReturn(run func(context.Context, string) ([]files.FileInfo, error)) *MockDbfsInterface_RecursiveList_Call {
	_c.Call.Return(run)
	return _c
}

// WriteFile provides a mock function with given fields: ctx, name, data
func (_m *MockDbfsInterface) WriteFile(ctx context.Context, name string, data []byte) error {
	ret := _m.Called(ctx, name, data)

	if len(ret) == 0 {
		panic("no return value specified for WriteFile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) error); ok {
		r0 = rf(ctx, name, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDbfsInterface_WriteFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteFile'
type MockDbfsInterface_WriteFile_Call struct {
	*mock.Call
}

// WriteFile is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - data []byte
func (_e *MockDbfsInterface_Expecter) WriteFile(ctx interface{}, name interface{}, data interface{}) *MockDbfsInterface_WriteFile_Call {
	return &MockDbfsInterface_WriteFile_Call{Call: _e.mock.On("WriteFile", ctx, name, data)}
}

func (_c *MockDbfsInterface_WriteFile_Call) Run(run func(ctx context.Context, name string, data []byte)) *MockDbfsInterface_WriteFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]byte))
	})
	return _c
}

func (_c *MockDbfsInterface_WriteFile_Call) Return(_a0 error) *MockDbfsInterface_WriteFile_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDbfsInterface_WriteFile_Call) RunAndReturn(run func(context.Context, string, []byte) error) *MockDbfsInterface_WriteFile_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDbfsInterface creates a new instance of MockDbfsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDbfsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDbfsInterface {
	mock := &MockDbfsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
