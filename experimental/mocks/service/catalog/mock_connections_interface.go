// Code generated by mockery v2.43.0. DO NOT EDIT.

package catalog

import (
	context "context"

	catalog "github.com/databricks/databricks-sdk-go/service/catalog"

	listing "github.com/databricks/databricks-sdk-go/listing"

	mock "github.com/stretchr/testify/mock"
)

// MockConnectionsInterface is an autogenerated mock type for the ConnectionsInterface type
type MockConnectionsInterface struct {
	mock.Mock
}

type MockConnectionsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConnectionsInterface) EXPECT() *MockConnectionsInterface_Expecter {
	return &MockConnectionsInterface_Expecter{mock: &_m.Mock}
}

// ConnectionInfoNameToFullNameMap provides a mock function with given fields: ctx, request
func (_m *MockConnectionsInterface) ConnectionInfoNameToFullNameMap(ctx context.Context, request catalog.ListConnectionsRequest) (map[string]string, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ConnectionInfoNameToFullNameMap")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListConnectionsRequest) (map[string]string, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListConnectionsRequest) map[string]string); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.ListConnectionsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionsInterface_ConnectionInfoNameToFullNameMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConnectionInfoNameToFullNameMap'
type MockConnectionsInterface_ConnectionInfoNameToFullNameMap_Call struct {
	*mock.Call
}

// ConnectionInfoNameToFullNameMap is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListConnectionsRequest
func (_e *MockConnectionsInterface_Expecter) ConnectionInfoNameToFullNameMap(ctx interface{}, request interface{}) *MockConnectionsInterface_ConnectionInfoNameToFullNameMap_Call {
	return &MockConnectionsInterface_ConnectionInfoNameToFullNameMap_Call{Call: _e.mock.On("ConnectionInfoNameToFullNameMap", ctx, request)}
}

func (_c *MockConnectionsInterface_ConnectionInfoNameToFullNameMap_Call) Run(run func(ctx context.Context, request catalog.ListConnectionsRequest)) *MockConnectionsInterface_ConnectionInfoNameToFullNameMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListConnectionsRequest))
	})
	return _c
}

func (_c *MockConnectionsInterface_ConnectionInfoNameToFullNameMap_Call) Return(_a0 map[string]string, _a1 error) *MockConnectionsInterface_ConnectionInfoNameToFullNameMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionsInterface_ConnectionInfoNameToFullNameMap_Call) RunAndReturn(run func(context.Context, catalog.ListConnectionsRequest) (map[string]string, error)) *MockConnectionsInterface_ConnectionInfoNameToFullNameMap_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockConnectionsInterface) Create(ctx context.Context, request catalog.CreateConnection) (*catalog.ConnectionInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *catalog.ConnectionInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateConnection) (*catalog.ConnectionInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateConnection) *catalog.ConnectionInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.ConnectionInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.CreateConnection) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockConnectionsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.CreateConnection
func (_e *MockConnectionsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockConnectionsInterface_Create_Call {
	return &MockConnectionsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockConnectionsInterface_Create_Call) Run(run func(ctx context.Context, request catalog.CreateConnection)) *MockConnectionsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.CreateConnection))
	})
	return _c
}

func (_c *MockConnectionsInterface_Create_Call) Return(_a0 *catalog.ConnectionInfo, _a1 error) *MockConnectionsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionsInterface_Create_Call) RunAndReturn(run func(context.Context, catalog.CreateConnection) (*catalog.ConnectionInfo, error)) *MockConnectionsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockConnectionsInterface) Delete(ctx context.Context, request catalog.DeleteConnectionRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.DeleteConnectionRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockConnectionsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockConnectionsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.DeleteConnectionRequest
func (_e *MockConnectionsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockConnectionsInterface_Delete_Call {
	return &MockConnectionsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockConnectionsInterface_Delete_Call) Run(run func(ctx context.Context, request catalog.DeleteConnectionRequest)) *MockConnectionsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.DeleteConnectionRequest))
	})
	return _c
}

func (_c *MockConnectionsInterface_Delete_Call) Return(_a0 error) *MockConnectionsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConnectionsInterface_Delete_Call) RunAndReturn(run func(context.Context, catalog.DeleteConnectionRequest) error) *MockConnectionsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByName provides a mock function with given fields: ctx, name
func (_m *MockConnectionsInterface) DeleteByName(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockConnectionsInterface_DeleteByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByName'
type MockConnectionsInterface_DeleteByName_Call struct {
	*mock.Call
}

// DeleteByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockConnectionsInterface_Expecter) DeleteByName(ctx interface{}, name interface{}) *MockConnectionsInterface_DeleteByName_Call {
	return &MockConnectionsInterface_DeleteByName_Call{Call: _e.mock.On("DeleteByName", ctx, name)}
}

func (_c *MockConnectionsInterface_DeleteByName_Call) Run(run func(ctx context.Context, name string)) *MockConnectionsInterface_DeleteByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockConnectionsInterface_DeleteByName_Call) Return(_a0 error) *MockConnectionsInterface_DeleteByName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConnectionsInterface_DeleteByName_Call) RunAndReturn(run func(context.Context, string) error) *MockConnectionsInterface_DeleteByName_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockConnectionsInterface) Get(ctx context.Context, request catalog.GetConnectionRequest) (*catalog.ConnectionInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *catalog.ConnectionInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetConnectionRequest) (*catalog.ConnectionInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetConnectionRequest) *catalog.ConnectionInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.ConnectionInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.GetConnectionRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockConnectionsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.GetConnectionRequest
func (_e *MockConnectionsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockConnectionsInterface_Get_Call {
	return &MockConnectionsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockConnectionsInterface_Get_Call) Run(run func(ctx context.Context, request catalog.GetConnectionRequest)) *MockConnectionsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.GetConnectionRequest))
	})
	return _c
}

func (_c *MockConnectionsInterface_Get_Call) Return(_a0 *catalog.ConnectionInfo, _a1 error) *MockConnectionsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionsInterface_Get_Call) RunAndReturn(run func(context.Context, catalog.GetConnectionRequest) (*catalog.ConnectionInfo, error)) *MockConnectionsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *MockConnectionsInterface) GetByName(ctx context.Context, name string) (*catalog.ConnectionInfo, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *catalog.ConnectionInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*catalog.ConnectionInfo, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *catalog.ConnectionInfo); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.ConnectionInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionsInterface_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type MockConnectionsInterface_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockConnectionsInterface_Expecter) GetByName(ctx interface{}, name interface{}) *MockConnectionsInterface_GetByName_Call {
	return &MockConnectionsInterface_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *MockConnectionsInterface_GetByName_Call) Run(run func(ctx context.Context, name string)) *MockConnectionsInterface_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockConnectionsInterface_GetByName_Call) Return(_a0 *catalog.ConnectionInfo, _a1 error) *MockConnectionsInterface_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionsInterface_GetByName_Call) RunAndReturn(run func(context.Context, string) (*catalog.ConnectionInfo, error)) *MockConnectionsInterface_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockConnectionsInterface) Impl() catalog.ConnectionsService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 catalog.ConnectionsService
	if rf, ok := ret.Get(0).(func() catalog.ConnectionsService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.ConnectionsService)
		}
	}

	return r0
}

// MockConnectionsInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockConnectionsInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockConnectionsInterface_Expecter) Impl() *MockConnectionsInterface_Impl_Call {
	return &MockConnectionsInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockConnectionsInterface_Impl_Call) Run(run func()) *MockConnectionsInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConnectionsInterface_Impl_Call) Return(_a0 catalog.ConnectionsService) *MockConnectionsInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConnectionsInterface_Impl_Call) RunAndReturn(run func() catalog.ConnectionsService) *MockConnectionsInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockConnectionsInterface) List(ctx context.Context, request catalog.ListConnectionsRequest) listing.Iterator[catalog.ConnectionInfo] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[catalog.ConnectionInfo]
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListConnectionsRequest) listing.Iterator[catalog.ConnectionInfo]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[catalog.ConnectionInfo])
		}
	}

	return r0
}

// MockConnectionsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockConnectionsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListConnectionsRequest
func (_e *MockConnectionsInterface_Expecter) List(ctx interface{}, request interface{}) *MockConnectionsInterface_List_Call {
	return &MockConnectionsInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockConnectionsInterface_List_Call) Run(run func(ctx context.Context, request catalog.ListConnectionsRequest)) *MockConnectionsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListConnectionsRequest))
	})
	return _c
}

func (_c *MockConnectionsInterface_List_Call) Return(_a0 listing.Iterator[catalog.ConnectionInfo]) *MockConnectionsInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConnectionsInterface_List_Call) RunAndReturn(run func(context.Context, catalog.ListConnectionsRequest) listing.Iterator[catalog.ConnectionInfo]) *MockConnectionsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockConnectionsInterface) ListAll(ctx context.Context, request catalog.ListConnectionsRequest) ([]catalog.ConnectionInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []catalog.ConnectionInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListConnectionsRequest) ([]catalog.ConnectionInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListConnectionsRequest) []catalog.ConnectionInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]catalog.ConnectionInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.ListConnectionsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionsInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockConnectionsInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListConnectionsRequest
func (_e *MockConnectionsInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockConnectionsInterface_ListAll_Call {
	return &MockConnectionsInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockConnectionsInterface_ListAll_Call) Run(run func(ctx context.Context, request catalog.ListConnectionsRequest)) *MockConnectionsInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListConnectionsRequest))
	})
	return _c
}

func (_c *MockConnectionsInterface_ListAll_Call) Return(_a0 []catalog.ConnectionInfo, _a1 error) *MockConnectionsInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionsInterface_ListAll_Call) RunAndReturn(run func(context.Context, catalog.ListConnectionsRequest) ([]catalog.ConnectionInfo, error)) *MockConnectionsInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockConnectionsInterface) Update(ctx context.Context, request catalog.UpdateConnection) (*catalog.ConnectionInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *catalog.ConnectionInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdateConnection) (*catalog.ConnectionInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdateConnection) *catalog.ConnectionInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.ConnectionInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.UpdateConnection) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockConnectionsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.UpdateConnection
func (_e *MockConnectionsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockConnectionsInterface_Update_Call {
	return &MockConnectionsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockConnectionsInterface_Update_Call) Run(run func(ctx context.Context, request catalog.UpdateConnection)) *MockConnectionsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.UpdateConnection))
	})
	return _c
}

func (_c *MockConnectionsInterface_Update_Call) Return(_a0 *catalog.ConnectionInfo, _a1 error) *MockConnectionsInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionsInterface_Update_Call) RunAndReturn(run func(context.Context, catalog.UpdateConnection) (*catalog.ConnectionInfo, error)) *MockConnectionsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockConnectionsInterface) WithImpl(impl catalog.ConnectionsService) catalog.ConnectionsInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 catalog.ConnectionsInterface
	if rf, ok := ret.Get(0).(func(catalog.ConnectionsService) catalog.ConnectionsInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.ConnectionsInterface)
		}
	}

	return r0
}

// MockConnectionsInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockConnectionsInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl catalog.ConnectionsService
func (_e *MockConnectionsInterface_Expecter) WithImpl(impl interface{}) *MockConnectionsInterface_WithImpl_Call {
	return &MockConnectionsInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockConnectionsInterface_WithImpl_Call) Run(run func(impl catalog.ConnectionsService)) *MockConnectionsInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(catalog.ConnectionsService))
	})
	return _c
}

func (_c *MockConnectionsInterface_WithImpl_Call) Return(_a0 catalog.ConnectionsInterface) *MockConnectionsInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConnectionsInterface_WithImpl_Call) RunAndReturn(run func(catalog.ConnectionsService) catalog.ConnectionsInterface) *MockConnectionsInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockConnectionsInterface creates a new instance of MockConnectionsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConnectionsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConnectionsInterface {
	mock := &MockConnectionsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
