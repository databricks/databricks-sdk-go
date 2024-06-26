// Code generated by mockery v2.43.0. DO NOT EDIT.

package sql

import (
	context "context"

	sql "github.com/databricks/databricks-sdk-go/service/sql"
	mock "github.com/stretchr/testify/mock"
)

// MockAlertsInterface is an autogenerated mock type for the AlertsInterface type
type MockAlertsInterface struct {
	mock.Mock
}

type MockAlertsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAlertsInterface) EXPECT() *MockAlertsInterface_Expecter {
	return &MockAlertsInterface_Expecter{mock: &_m.Mock}
}

// AlertNameToIdMap provides a mock function with given fields: ctx
func (_m *MockAlertsInterface) AlertNameToIdMap(ctx context.Context) (map[string]string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for AlertNameToIdMap")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (map[string]string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) map[string]string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAlertsInterface_AlertNameToIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AlertNameToIdMap'
type MockAlertsInterface_AlertNameToIdMap_Call struct {
	*mock.Call
}

// AlertNameToIdMap is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockAlertsInterface_Expecter) AlertNameToIdMap(ctx interface{}) *MockAlertsInterface_AlertNameToIdMap_Call {
	return &MockAlertsInterface_AlertNameToIdMap_Call{Call: _e.mock.On("AlertNameToIdMap", ctx)}
}

func (_c *MockAlertsInterface_AlertNameToIdMap_Call) Run(run func(ctx context.Context)) *MockAlertsInterface_AlertNameToIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockAlertsInterface_AlertNameToIdMap_Call) Return(_a0 map[string]string, _a1 error) *MockAlertsInterface_AlertNameToIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAlertsInterface_AlertNameToIdMap_Call) RunAndReturn(run func(context.Context) (map[string]string, error)) *MockAlertsInterface_AlertNameToIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockAlertsInterface) Create(ctx context.Context, request sql.CreateAlert) (*sql.Alert, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *sql.Alert
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.CreateAlert) (*sql.Alert, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sql.CreateAlert) *sql.Alert); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Alert)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sql.CreateAlert) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAlertsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockAlertsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.CreateAlert
func (_e *MockAlertsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockAlertsInterface_Create_Call {
	return &MockAlertsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockAlertsInterface_Create_Call) Run(run func(ctx context.Context, request sql.CreateAlert)) *MockAlertsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.CreateAlert))
	})
	return _c
}

func (_c *MockAlertsInterface_Create_Call) Return(_a0 *sql.Alert, _a1 error) *MockAlertsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAlertsInterface_Create_Call) RunAndReturn(run func(context.Context, sql.CreateAlert) (*sql.Alert, error)) *MockAlertsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockAlertsInterface) Delete(ctx context.Context, request sql.DeleteAlertRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.DeleteAlertRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAlertsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockAlertsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.DeleteAlertRequest
func (_e *MockAlertsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockAlertsInterface_Delete_Call {
	return &MockAlertsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockAlertsInterface_Delete_Call) Run(run func(ctx context.Context, request sql.DeleteAlertRequest)) *MockAlertsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.DeleteAlertRequest))
	})
	return _c
}

func (_c *MockAlertsInterface_Delete_Call) Return(_a0 error) *MockAlertsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAlertsInterface_Delete_Call) RunAndReturn(run func(context.Context, sql.DeleteAlertRequest) error) *MockAlertsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByAlertId provides a mock function with given fields: ctx, alertId
func (_m *MockAlertsInterface) DeleteByAlertId(ctx context.Context, alertId string) error {
	ret := _m.Called(ctx, alertId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByAlertId")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, alertId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAlertsInterface_DeleteByAlertId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByAlertId'
type MockAlertsInterface_DeleteByAlertId_Call struct {
	*mock.Call
}

// DeleteByAlertId is a helper method to define mock.On call
//   - ctx context.Context
//   - alertId string
func (_e *MockAlertsInterface_Expecter) DeleteByAlertId(ctx interface{}, alertId interface{}) *MockAlertsInterface_DeleteByAlertId_Call {
	return &MockAlertsInterface_DeleteByAlertId_Call{Call: _e.mock.On("DeleteByAlertId", ctx, alertId)}
}

func (_c *MockAlertsInterface_DeleteByAlertId_Call) Run(run func(ctx context.Context, alertId string)) *MockAlertsInterface_DeleteByAlertId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAlertsInterface_DeleteByAlertId_Call) Return(_a0 error) *MockAlertsInterface_DeleteByAlertId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAlertsInterface_DeleteByAlertId_Call) RunAndReturn(run func(context.Context, string) error) *MockAlertsInterface_DeleteByAlertId_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockAlertsInterface) Get(ctx context.Context, request sql.GetAlertRequest) (*sql.Alert, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *sql.Alert
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.GetAlertRequest) (*sql.Alert, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sql.GetAlertRequest) *sql.Alert); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Alert)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sql.GetAlertRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAlertsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockAlertsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.GetAlertRequest
func (_e *MockAlertsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockAlertsInterface_Get_Call {
	return &MockAlertsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockAlertsInterface_Get_Call) Run(run func(ctx context.Context, request sql.GetAlertRequest)) *MockAlertsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.GetAlertRequest))
	})
	return _c
}

func (_c *MockAlertsInterface_Get_Call) Return(_a0 *sql.Alert, _a1 error) *MockAlertsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAlertsInterface_Get_Call) RunAndReturn(run func(context.Context, sql.GetAlertRequest) (*sql.Alert, error)) *MockAlertsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByAlertId provides a mock function with given fields: ctx, alertId
func (_m *MockAlertsInterface) GetByAlertId(ctx context.Context, alertId string) (*sql.Alert, error) {
	ret := _m.Called(ctx, alertId)

	if len(ret) == 0 {
		panic("no return value specified for GetByAlertId")
	}

	var r0 *sql.Alert
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*sql.Alert, error)); ok {
		return rf(ctx, alertId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *sql.Alert); ok {
		r0 = rf(ctx, alertId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Alert)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, alertId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAlertsInterface_GetByAlertId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByAlertId'
type MockAlertsInterface_GetByAlertId_Call struct {
	*mock.Call
}

// GetByAlertId is a helper method to define mock.On call
//   - ctx context.Context
//   - alertId string
func (_e *MockAlertsInterface_Expecter) GetByAlertId(ctx interface{}, alertId interface{}) *MockAlertsInterface_GetByAlertId_Call {
	return &MockAlertsInterface_GetByAlertId_Call{Call: _e.mock.On("GetByAlertId", ctx, alertId)}
}

func (_c *MockAlertsInterface_GetByAlertId_Call) Run(run func(ctx context.Context, alertId string)) *MockAlertsInterface_GetByAlertId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAlertsInterface_GetByAlertId_Call) Return(_a0 *sql.Alert, _a1 error) *MockAlertsInterface_GetByAlertId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAlertsInterface_GetByAlertId_Call) RunAndReturn(run func(context.Context, string) (*sql.Alert, error)) *MockAlertsInterface_GetByAlertId_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *MockAlertsInterface) GetByName(ctx context.Context, name string) (*sql.Alert, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *sql.Alert
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*sql.Alert, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *sql.Alert); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Alert)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAlertsInterface_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type MockAlertsInterface_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockAlertsInterface_Expecter) GetByName(ctx interface{}, name interface{}) *MockAlertsInterface_GetByName_Call {
	return &MockAlertsInterface_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *MockAlertsInterface_GetByName_Call) Run(run func(ctx context.Context, name string)) *MockAlertsInterface_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAlertsInterface_GetByName_Call) Return(_a0 *sql.Alert, _a1 error) *MockAlertsInterface_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAlertsInterface_GetByName_Call) RunAndReturn(run func(context.Context, string) (*sql.Alert, error)) *MockAlertsInterface_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockAlertsInterface) Impl() sql.AlertsService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 sql.AlertsService
	if rf, ok := ret.Get(0).(func() sql.AlertsService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.AlertsService)
		}
	}

	return r0
}

// MockAlertsInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockAlertsInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockAlertsInterface_Expecter) Impl() *MockAlertsInterface_Impl_Call {
	return &MockAlertsInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockAlertsInterface_Impl_Call) Run(run func()) *MockAlertsInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAlertsInterface_Impl_Call) Return(_a0 sql.AlertsService) *MockAlertsInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAlertsInterface_Impl_Call) RunAndReturn(run func() sql.AlertsService) *MockAlertsInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx
func (_m *MockAlertsInterface) List(ctx context.Context) ([]sql.Alert, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []sql.Alert
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]sql.Alert, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []sql.Alert); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sql.Alert)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAlertsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockAlertsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockAlertsInterface_Expecter) List(ctx interface{}) *MockAlertsInterface_List_Call {
	return &MockAlertsInterface_List_Call{Call: _e.mock.On("List", ctx)}
}

func (_c *MockAlertsInterface_List_Call) Run(run func(ctx context.Context)) *MockAlertsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockAlertsInterface_List_Call) Return(_a0 []sql.Alert, _a1 error) *MockAlertsInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAlertsInterface_List_Call) RunAndReturn(run func(context.Context) ([]sql.Alert, error)) *MockAlertsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockAlertsInterface) Update(ctx context.Context, request sql.EditAlert) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.EditAlert) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAlertsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockAlertsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.EditAlert
func (_e *MockAlertsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockAlertsInterface_Update_Call {
	return &MockAlertsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockAlertsInterface_Update_Call) Run(run func(ctx context.Context, request sql.EditAlert)) *MockAlertsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.EditAlert))
	})
	return _c
}

func (_c *MockAlertsInterface_Update_Call) Return(_a0 error) *MockAlertsInterface_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAlertsInterface_Update_Call) RunAndReturn(run func(context.Context, sql.EditAlert) error) *MockAlertsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockAlertsInterface) WithImpl(impl sql.AlertsService) sql.AlertsInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 sql.AlertsInterface
	if rf, ok := ret.Get(0).(func(sql.AlertsService) sql.AlertsInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.AlertsInterface)
		}
	}

	return r0
}

// MockAlertsInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockAlertsInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl sql.AlertsService
func (_e *MockAlertsInterface_Expecter) WithImpl(impl interface{}) *MockAlertsInterface_WithImpl_Call {
	return &MockAlertsInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockAlertsInterface_WithImpl_Call) Run(run func(impl sql.AlertsService)) *MockAlertsInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(sql.AlertsService))
	})
	return _c
}

func (_c *MockAlertsInterface_WithImpl_Call) Return(_a0 sql.AlertsInterface) *MockAlertsInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAlertsInterface_WithImpl_Call) RunAndReturn(run func(sql.AlertsService) sql.AlertsInterface) *MockAlertsInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAlertsInterface creates a new instance of MockAlertsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAlertsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAlertsInterface {
	mock := &MockAlertsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
