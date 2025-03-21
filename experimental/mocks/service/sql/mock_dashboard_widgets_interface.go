// Code generated by mockery v2.53.2. DO NOT EDIT.

package sql

import (
	context "context"

	sql "github.com/databricks/databricks-sdk-go/service/sql"
	mock "github.com/stretchr/testify/mock"
)

// MockDashboardWidgetsInterface is an autogenerated mock type for the DashboardWidgetsInterface type
type MockDashboardWidgetsInterface struct {
	mock.Mock
}

type MockDashboardWidgetsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDashboardWidgetsInterface) EXPECT() *MockDashboardWidgetsInterface_Expecter {
	return &MockDashboardWidgetsInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockDashboardWidgetsInterface) Create(ctx context.Context, request sql.CreateWidget) (*sql.Widget, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *sql.Widget
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.CreateWidget) (*sql.Widget, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sql.CreateWidget) *sql.Widget); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Widget)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sql.CreateWidget) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDashboardWidgetsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockDashboardWidgetsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.CreateWidget
func (_e *MockDashboardWidgetsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockDashboardWidgetsInterface_Create_Call {
	return &MockDashboardWidgetsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockDashboardWidgetsInterface_Create_Call) Run(run func(ctx context.Context, request sql.CreateWidget)) *MockDashboardWidgetsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.CreateWidget))
	})
	return _c
}

func (_c *MockDashboardWidgetsInterface_Create_Call) Return(_a0 *sql.Widget, _a1 error) *MockDashboardWidgetsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDashboardWidgetsInterface_Create_Call) RunAndReturn(run func(context.Context, sql.CreateWidget) (*sql.Widget, error)) *MockDashboardWidgetsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockDashboardWidgetsInterface) Delete(ctx context.Context, request sql.DeleteDashboardWidgetRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.DeleteDashboardWidgetRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDashboardWidgetsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockDashboardWidgetsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.DeleteDashboardWidgetRequest
func (_e *MockDashboardWidgetsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockDashboardWidgetsInterface_Delete_Call {
	return &MockDashboardWidgetsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockDashboardWidgetsInterface_Delete_Call) Run(run func(ctx context.Context, request sql.DeleteDashboardWidgetRequest)) *MockDashboardWidgetsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.DeleteDashboardWidgetRequest))
	})
	return _c
}

func (_c *MockDashboardWidgetsInterface_Delete_Call) Return(_a0 error) *MockDashboardWidgetsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDashboardWidgetsInterface_Delete_Call) RunAndReturn(run func(context.Context, sql.DeleteDashboardWidgetRequest) error) *MockDashboardWidgetsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *MockDashboardWidgetsInterface) DeleteById(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteById")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDashboardWidgetsInterface_DeleteById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteById'
type MockDashboardWidgetsInterface_DeleteById_Call struct {
	*mock.Call
}

// DeleteById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockDashboardWidgetsInterface_Expecter) DeleteById(ctx interface{}, id interface{}) *MockDashboardWidgetsInterface_DeleteById_Call {
	return &MockDashboardWidgetsInterface_DeleteById_Call{Call: _e.mock.On("DeleteById", ctx, id)}
}

func (_c *MockDashboardWidgetsInterface_DeleteById_Call) Run(run func(ctx context.Context, id string)) *MockDashboardWidgetsInterface_DeleteById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockDashboardWidgetsInterface_DeleteById_Call) Return(_a0 error) *MockDashboardWidgetsInterface_DeleteById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDashboardWidgetsInterface_DeleteById_Call) RunAndReturn(run func(context.Context, string) error) *MockDashboardWidgetsInterface_DeleteById_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockDashboardWidgetsInterface) Update(ctx context.Context, request sql.CreateWidget) (*sql.Widget, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *sql.Widget
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.CreateWidget) (*sql.Widget, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sql.CreateWidget) *sql.Widget); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Widget)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sql.CreateWidget) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDashboardWidgetsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockDashboardWidgetsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.CreateWidget
func (_e *MockDashboardWidgetsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockDashboardWidgetsInterface_Update_Call {
	return &MockDashboardWidgetsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockDashboardWidgetsInterface_Update_Call) Run(run func(ctx context.Context, request sql.CreateWidget)) *MockDashboardWidgetsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.CreateWidget))
	})
	return _c
}

func (_c *MockDashboardWidgetsInterface_Update_Call) Return(_a0 *sql.Widget, _a1 error) *MockDashboardWidgetsInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDashboardWidgetsInterface_Update_Call) RunAndReturn(run func(context.Context, sql.CreateWidget) (*sql.Widget, error)) *MockDashboardWidgetsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDashboardWidgetsInterface creates a new instance of MockDashboardWidgetsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDashboardWidgetsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDashboardWidgetsInterface {
	mock := &MockDashboardWidgetsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
