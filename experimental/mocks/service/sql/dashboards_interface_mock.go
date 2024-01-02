// Code generated by mockery v2.39.1. DO NOT EDIT.

package sql

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	mock "github.com/stretchr/testify/mock"

	sql "github.com/databricks/databricks-sdk-go/service/sql"
)

// MockDashboardsInterface is an autogenerated mock type for the DashboardsInterface type
type MockDashboardsInterface struct {
	mock.Mock
}

type MockDashboardsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDashboardsInterface) EXPECT() *MockDashboardsInterface_Expecter {
	return &MockDashboardsInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockDashboardsInterface) Create(ctx context.Context, request sql.CreateDashboardRequest) (*sql.Dashboard, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *sql.Dashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.CreateDashboardRequest) (*sql.Dashboard, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sql.CreateDashboardRequest) *sql.Dashboard); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Dashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sql.CreateDashboardRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDashboardsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockDashboardsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.CreateDashboardRequest
func (_e *MockDashboardsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockDashboardsInterface_Create_Call {
	return &MockDashboardsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockDashboardsInterface_Create_Call) Run(run func(ctx context.Context, request sql.CreateDashboardRequest)) *MockDashboardsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.CreateDashboardRequest))
	})
	return _c
}

func (_c *MockDashboardsInterface_Create_Call) Return(_a0 *sql.Dashboard, _a1 error) *MockDashboardsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDashboardsInterface_Create_Call) RunAndReturn(run func(context.Context, sql.CreateDashboardRequest) (*sql.Dashboard, error)) *MockDashboardsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// DashboardNameToIdMap provides a mock function with given fields: ctx, request
func (_m *MockDashboardsInterface) DashboardNameToIdMap(ctx context.Context, request sql.ListDashboardsRequest) (map[string]string, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for DashboardNameToIdMap")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.ListDashboardsRequest) (map[string]string, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sql.ListDashboardsRequest) map[string]string); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sql.ListDashboardsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDashboardsInterface_DashboardNameToIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DashboardNameToIdMap'
type MockDashboardsInterface_DashboardNameToIdMap_Call struct {
	*mock.Call
}

// DashboardNameToIdMap is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.ListDashboardsRequest
func (_e *MockDashboardsInterface_Expecter) DashboardNameToIdMap(ctx interface{}, request interface{}) *MockDashboardsInterface_DashboardNameToIdMap_Call {
	return &MockDashboardsInterface_DashboardNameToIdMap_Call{Call: _e.mock.On("DashboardNameToIdMap", ctx, request)}
}

func (_c *MockDashboardsInterface_DashboardNameToIdMap_Call) Run(run func(ctx context.Context, request sql.ListDashboardsRequest)) *MockDashboardsInterface_DashboardNameToIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.ListDashboardsRequest))
	})
	return _c
}

func (_c *MockDashboardsInterface_DashboardNameToIdMap_Call) Return(_a0 map[string]string, _a1 error) *MockDashboardsInterface_DashboardNameToIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDashboardsInterface_DashboardNameToIdMap_Call) RunAndReturn(run func(context.Context, sql.ListDashboardsRequest) (map[string]string, error)) *MockDashboardsInterface_DashboardNameToIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockDashboardsInterface) Delete(ctx context.Context, request sql.DeleteDashboardRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.DeleteDashboardRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDashboardsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockDashboardsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.DeleteDashboardRequest
func (_e *MockDashboardsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockDashboardsInterface_Delete_Call {
	return &MockDashboardsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockDashboardsInterface_Delete_Call) Run(run func(ctx context.Context, request sql.DeleteDashboardRequest)) *MockDashboardsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.DeleteDashboardRequest))
	})
	return _c
}

func (_c *MockDashboardsInterface_Delete_Call) Return(_a0 error) *MockDashboardsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDashboardsInterface_Delete_Call) RunAndReturn(run func(context.Context, sql.DeleteDashboardRequest) error) *MockDashboardsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByDashboardId provides a mock function with given fields: ctx, dashboardId
func (_m *MockDashboardsInterface) DeleteByDashboardId(ctx context.Context, dashboardId string) error {
	ret := _m.Called(ctx, dashboardId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByDashboardId")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, dashboardId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDashboardsInterface_DeleteByDashboardId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByDashboardId'
type MockDashboardsInterface_DeleteByDashboardId_Call struct {
	*mock.Call
}

// DeleteByDashboardId is a helper method to define mock.On call
//   - ctx context.Context
//   - dashboardId string
func (_e *MockDashboardsInterface_Expecter) DeleteByDashboardId(ctx interface{}, dashboardId interface{}) *MockDashboardsInterface_DeleteByDashboardId_Call {
	return &MockDashboardsInterface_DeleteByDashboardId_Call{Call: _e.mock.On("DeleteByDashboardId", ctx, dashboardId)}
}

func (_c *MockDashboardsInterface_DeleteByDashboardId_Call) Run(run func(ctx context.Context, dashboardId string)) *MockDashboardsInterface_DeleteByDashboardId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockDashboardsInterface_DeleteByDashboardId_Call) Return(_a0 error) *MockDashboardsInterface_DeleteByDashboardId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDashboardsInterface_DeleteByDashboardId_Call) RunAndReturn(run func(context.Context, string) error) *MockDashboardsInterface_DeleteByDashboardId_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockDashboardsInterface) Get(ctx context.Context, request sql.GetDashboardRequest) (*sql.Dashboard, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *sql.Dashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.GetDashboardRequest) (*sql.Dashboard, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sql.GetDashboardRequest) *sql.Dashboard); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Dashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sql.GetDashboardRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDashboardsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockDashboardsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.GetDashboardRequest
func (_e *MockDashboardsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockDashboardsInterface_Get_Call {
	return &MockDashboardsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockDashboardsInterface_Get_Call) Run(run func(ctx context.Context, request sql.GetDashboardRequest)) *MockDashboardsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.GetDashboardRequest))
	})
	return _c
}

func (_c *MockDashboardsInterface_Get_Call) Return(_a0 *sql.Dashboard, _a1 error) *MockDashboardsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDashboardsInterface_Get_Call) RunAndReturn(run func(context.Context, sql.GetDashboardRequest) (*sql.Dashboard, error)) *MockDashboardsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByDashboardId provides a mock function with given fields: ctx, dashboardId
func (_m *MockDashboardsInterface) GetByDashboardId(ctx context.Context, dashboardId string) (*sql.Dashboard, error) {
	ret := _m.Called(ctx, dashboardId)

	if len(ret) == 0 {
		panic("no return value specified for GetByDashboardId")
	}

	var r0 *sql.Dashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*sql.Dashboard, error)); ok {
		return rf(ctx, dashboardId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *sql.Dashboard); ok {
		r0 = rf(ctx, dashboardId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Dashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, dashboardId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDashboardsInterface_GetByDashboardId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByDashboardId'
type MockDashboardsInterface_GetByDashboardId_Call struct {
	*mock.Call
}

// GetByDashboardId is a helper method to define mock.On call
//   - ctx context.Context
//   - dashboardId string
func (_e *MockDashboardsInterface_Expecter) GetByDashboardId(ctx interface{}, dashboardId interface{}) *MockDashboardsInterface_GetByDashboardId_Call {
	return &MockDashboardsInterface_GetByDashboardId_Call{Call: _e.mock.On("GetByDashboardId", ctx, dashboardId)}
}

func (_c *MockDashboardsInterface_GetByDashboardId_Call) Run(run func(ctx context.Context, dashboardId string)) *MockDashboardsInterface_GetByDashboardId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockDashboardsInterface_GetByDashboardId_Call) Return(_a0 *sql.Dashboard, _a1 error) *MockDashboardsInterface_GetByDashboardId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDashboardsInterface_GetByDashboardId_Call) RunAndReturn(run func(context.Context, string) (*sql.Dashboard, error)) *MockDashboardsInterface_GetByDashboardId_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *MockDashboardsInterface) GetByName(ctx context.Context, name string) (*sql.Dashboard, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *sql.Dashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*sql.Dashboard, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *sql.Dashboard); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Dashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDashboardsInterface_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type MockDashboardsInterface_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockDashboardsInterface_Expecter) GetByName(ctx interface{}, name interface{}) *MockDashboardsInterface_GetByName_Call {
	return &MockDashboardsInterface_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *MockDashboardsInterface_GetByName_Call) Run(run func(ctx context.Context, name string)) *MockDashboardsInterface_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockDashboardsInterface_GetByName_Call) Return(_a0 *sql.Dashboard, _a1 error) *MockDashboardsInterface_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDashboardsInterface_GetByName_Call) RunAndReturn(run func(context.Context, string) (*sql.Dashboard, error)) *MockDashboardsInterface_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockDashboardsInterface) Impl() sql.DashboardsService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 sql.DashboardsService
	if rf, ok := ret.Get(0).(func() sql.DashboardsService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.DashboardsService)
		}
	}

	return r0
}

// MockDashboardsInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockDashboardsInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockDashboardsInterface_Expecter) Impl() *MockDashboardsInterface_Impl_Call {
	return &MockDashboardsInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockDashboardsInterface_Impl_Call) Run(run func()) *MockDashboardsInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDashboardsInterface_Impl_Call) Return(_a0 sql.DashboardsService) *MockDashboardsInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDashboardsInterface_Impl_Call) RunAndReturn(run func() sql.DashboardsService) *MockDashboardsInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockDashboardsInterface) List(ctx context.Context, request sql.ListDashboardsRequest) *listing.DeduplicatingIterator[sql.Dashboard, string] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *listing.DeduplicatingIterator[sql.Dashboard, string]
	if rf, ok := ret.Get(0).(func(context.Context, sql.ListDashboardsRequest) *listing.DeduplicatingIterator[sql.Dashboard, string]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*listing.DeduplicatingIterator[sql.Dashboard, string])
		}
	}

	return r0
}

// MockDashboardsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockDashboardsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.ListDashboardsRequest
func (_e *MockDashboardsInterface_Expecter) List(ctx interface{}, request interface{}) *MockDashboardsInterface_List_Call {
	return &MockDashboardsInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockDashboardsInterface_List_Call) Run(run func(ctx context.Context, request sql.ListDashboardsRequest)) *MockDashboardsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.ListDashboardsRequest))
	})
	return _c
}

func (_c *MockDashboardsInterface_List_Call) Return(_a0 *listing.DeduplicatingIterator[sql.Dashboard, string]) *MockDashboardsInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDashboardsInterface_List_Call) RunAndReturn(run func(context.Context, sql.ListDashboardsRequest) *listing.DeduplicatingIterator[sql.Dashboard, string]) *MockDashboardsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockDashboardsInterface) ListAll(ctx context.Context, request sql.ListDashboardsRequest) ([]sql.Dashboard, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []sql.Dashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.ListDashboardsRequest) ([]sql.Dashboard, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sql.ListDashboardsRequest) []sql.Dashboard); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sql.Dashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sql.ListDashboardsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDashboardsInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockDashboardsInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.ListDashboardsRequest
func (_e *MockDashboardsInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockDashboardsInterface_ListAll_Call {
	return &MockDashboardsInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockDashboardsInterface_ListAll_Call) Run(run func(ctx context.Context, request sql.ListDashboardsRequest)) *MockDashboardsInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.ListDashboardsRequest))
	})
	return _c
}

func (_c *MockDashboardsInterface_ListAll_Call) Return(_a0 []sql.Dashboard, _a1 error) *MockDashboardsInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDashboardsInterface_ListAll_Call) RunAndReturn(run func(context.Context, sql.ListDashboardsRequest) ([]sql.Dashboard, error)) *MockDashboardsInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// Restore provides a mock function with given fields: ctx, request
func (_m *MockDashboardsInterface) Restore(ctx context.Context, request sql.RestoreDashboardRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Restore")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.RestoreDashboardRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDashboardsInterface_Restore_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Restore'
type MockDashboardsInterface_Restore_Call struct {
	*mock.Call
}

// Restore is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.RestoreDashboardRequest
func (_e *MockDashboardsInterface_Expecter) Restore(ctx interface{}, request interface{}) *MockDashboardsInterface_Restore_Call {
	return &MockDashboardsInterface_Restore_Call{Call: _e.mock.On("Restore", ctx, request)}
}

func (_c *MockDashboardsInterface_Restore_Call) Run(run func(ctx context.Context, request sql.RestoreDashboardRequest)) *MockDashboardsInterface_Restore_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.RestoreDashboardRequest))
	})
	return _c
}

func (_c *MockDashboardsInterface_Restore_Call) Return(_a0 error) *MockDashboardsInterface_Restore_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDashboardsInterface_Restore_Call) RunAndReturn(run func(context.Context, sql.RestoreDashboardRequest) error) *MockDashboardsInterface_Restore_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockDashboardsInterface) WithImpl(impl sql.DashboardsService) sql.DashboardsInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 sql.DashboardsInterface
	if rf, ok := ret.Get(0).(func(sql.DashboardsService) sql.DashboardsInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.DashboardsInterface)
		}
	}

	return r0
}

// MockDashboardsInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockDashboardsInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl sql.DashboardsService
func (_e *MockDashboardsInterface_Expecter) WithImpl(impl interface{}) *MockDashboardsInterface_WithImpl_Call {
	return &MockDashboardsInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockDashboardsInterface_WithImpl_Call) Run(run func(impl sql.DashboardsService)) *MockDashboardsInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(sql.DashboardsService))
	})
	return _c
}

func (_c *MockDashboardsInterface_WithImpl_Call) Return(_a0 sql.DashboardsInterface) *MockDashboardsInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDashboardsInterface_WithImpl_Call) RunAndReturn(run func(sql.DashboardsService) sql.DashboardsInterface) *MockDashboardsInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDashboardsInterface creates a new instance of MockDashboardsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDashboardsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDashboardsInterface {
	mock := &MockDashboardsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
