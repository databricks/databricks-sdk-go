// Code generated by mockery v2.39.1. DO NOT EDIT.

package catalog

import (
	context "context"

	catalog "github.com/databricks/databricks-sdk-go/service/catalog"

	mock "github.com/stretchr/testify/mock"
)

// MockOnlineTablesInterface is an autogenerated mock type for the OnlineTablesInterface type
type MockOnlineTablesInterface struct {
	mock.Mock
}

type MockOnlineTablesInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockOnlineTablesInterface) EXPECT() *MockOnlineTablesInterface_Expecter {
	return &MockOnlineTablesInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockOnlineTablesInterface) Create(ctx context.Context, request catalog.ViewData) (*catalog.OnlineTable, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *catalog.OnlineTable
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ViewData) (*catalog.OnlineTable, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ViewData) *catalog.OnlineTable); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.OnlineTable)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.ViewData) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOnlineTablesInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockOnlineTablesInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ViewData
func (_e *MockOnlineTablesInterface_Expecter) Create(ctx interface{}, request interface{}) *MockOnlineTablesInterface_Create_Call {
	return &MockOnlineTablesInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockOnlineTablesInterface_Create_Call) Run(run func(ctx context.Context, request catalog.ViewData)) *MockOnlineTablesInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ViewData))
	})
	return _c
}

func (_c *MockOnlineTablesInterface_Create_Call) Return(_a0 *catalog.OnlineTable, _a1 error) *MockOnlineTablesInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockOnlineTablesInterface_Create_Call) RunAndReturn(run func(context.Context, catalog.ViewData) (*catalog.OnlineTable, error)) *MockOnlineTablesInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockOnlineTablesInterface) Delete(ctx context.Context, request catalog.DeleteOnlineTableRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.DeleteOnlineTableRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockOnlineTablesInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockOnlineTablesInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.DeleteOnlineTableRequest
func (_e *MockOnlineTablesInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockOnlineTablesInterface_Delete_Call {
	return &MockOnlineTablesInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockOnlineTablesInterface_Delete_Call) Run(run func(ctx context.Context, request catalog.DeleteOnlineTableRequest)) *MockOnlineTablesInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.DeleteOnlineTableRequest))
	})
	return _c
}

func (_c *MockOnlineTablesInterface_Delete_Call) Return(_a0 error) *MockOnlineTablesInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOnlineTablesInterface_Delete_Call) RunAndReturn(run func(context.Context, catalog.DeleteOnlineTableRequest) error) *MockOnlineTablesInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByName provides a mock function with given fields: ctx, name
func (_m *MockOnlineTablesInterface) DeleteByName(ctx context.Context, name string) error {
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

// MockOnlineTablesInterface_DeleteByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByName'
type MockOnlineTablesInterface_DeleteByName_Call struct {
	*mock.Call
}

// DeleteByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockOnlineTablesInterface_Expecter) DeleteByName(ctx interface{}, name interface{}) *MockOnlineTablesInterface_DeleteByName_Call {
	return &MockOnlineTablesInterface_DeleteByName_Call{Call: _e.mock.On("DeleteByName", ctx, name)}
}

func (_c *MockOnlineTablesInterface_DeleteByName_Call) Run(run func(ctx context.Context, name string)) *MockOnlineTablesInterface_DeleteByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockOnlineTablesInterface_DeleteByName_Call) Return(_a0 error) *MockOnlineTablesInterface_DeleteByName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOnlineTablesInterface_DeleteByName_Call) RunAndReturn(run func(context.Context, string) error) *MockOnlineTablesInterface_DeleteByName_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockOnlineTablesInterface) Get(ctx context.Context, request catalog.GetOnlineTableRequest) (*catalog.OnlineTable, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *catalog.OnlineTable
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetOnlineTableRequest) (*catalog.OnlineTable, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetOnlineTableRequest) *catalog.OnlineTable); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.OnlineTable)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.GetOnlineTableRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOnlineTablesInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockOnlineTablesInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.GetOnlineTableRequest
func (_e *MockOnlineTablesInterface_Expecter) Get(ctx interface{}, request interface{}) *MockOnlineTablesInterface_Get_Call {
	return &MockOnlineTablesInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockOnlineTablesInterface_Get_Call) Run(run func(ctx context.Context, request catalog.GetOnlineTableRequest)) *MockOnlineTablesInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.GetOnlineTableRequest))
	})
	return _c
}

func (_c *MockOnlineTablesInterface_Get_Call) Return(_a0 *catalog.OnlineTable, _a1 error) *MockOnlineTablesInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockOnlineTablesInterface_Get_Call) RunAndReturn(run func(context.Context, catalog.GetOnlineTableRequest) (*catalog.OnlineTable, error)) *MockOnlineTablesInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *MockOnlineTablesInterface) GetByName(ctx context.Context, name string) (*catalog.OnlineTable, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *catalog.OnlineTable
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*catalog.OnlineTable, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *catalog.OnlineTable); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.OnlineTable)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOnlineTablesInterface_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type MockOnlineTablesInterface_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockOnlineTablesInterface_Expecter) GetByName(ctx interface{}, name interface{}) *MockOnlineTablesInterface_GetByName_Call {
	return &MockOnlineTablesInterface_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *MockOnlineTablesInterface_GetByName_Call) Run(run func(ctx context.Context, name string)) *MockOnlineTablesInterface_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockOnlineTablesInterface_GetByName_Call) Return(_a0 *catalog.OnlineTable, _a1 error) *MockOnlineTablesInterface_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockOnlineTablesInterface_GetByName_Call) RunAndReturn(run func(context.Context, string) (*catalog.OnlineTable, error)) *MockOnlineTablesInterface_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockOnlineTablesInterface) Impl() catalog.OnlineTablesService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 catalog.OnlineTablesService
	if rf, ok := ret.Get(0).(func() catalog.OnlineTablesService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.OnlineTablesService)
		}
	}

	return r0
}

// MockOnlineTablesInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockOnlineTablesInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockOnlineTablesInterface_Expecter) Impl() *MockOnlineTablesInterface_Impl_Call {
	return &MockOnlineTablesInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockOnlineTablesInterface_Impl_Call) Run(run func()) *MockOnlineTablesInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockOnlineTablesInterface_Impl_Call) Return(_a0 catalog.OnlineTablesService) *MockOnlineTablesInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOnlineTablesInterface_Impl_Call) RunAndReturn(run func() catalog.OnlineTablesService) *MockOnlineTablesInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockOnlineTablesInterface) WithImpl(impl catalog.OnlineTablesService) catalog.OnlineTablesInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 catalog.OnlineTablesInterface
	if rf, ok := ret.Get(0).(func(catalog.OnlineTablesService) catalog.OnlineTablesInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.OnlineTablesInterface)
		}
	}

	return r0
}

// MockOnlineTablesInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockOnlineTablesInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl catalog.OnlineTablesService
func (_e *MockOnlineTablesInterface_Expecter) WithImpl(impl interface{}) *MockOnlineTablesInterface_WithImpl_Call {
	return &MockOnlineTablesInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockOnlineTablesInterface_WithImpl_Call) Run(run func(impl catalog.OnlineTablesService)) *MockOnlineTablesInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(catalog.OnlineTablesService))
	})
	return _c
}

func (_c *MockOnlineTablesInterface_WithImpl_Call) Return(_a0 catalog.OnlineTablesInterface) *MockOnlineTablesInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOnlineTablesInterface_WithImpl_Call) RunAndReturn(run func(catalog.OnlineTablesService) catalog.OnlineTablesInterface) *MockOnlineTablesInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockOnlineTablesInterface creates a new instance of MockOnlineTablesInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOnlineTablesInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOnlineTablesInterface {
	mock := &MockOnlineTablesInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
