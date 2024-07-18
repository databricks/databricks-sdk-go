// Code generated by mockery v2.43.0. DO NOT EDIT.

package sql

import (
	context "context"

	sql "github.com/databricks/databricks-sdk-go/service/sql"
	mock "github.com/stretchr/testify/mock"
)

// MockQueryVisualizationsInterface is an autogenerated mock type for the QueryVisualizationsInterface type
type MockQueryVisualizationsInterface struct {
	mock.Mock
}

type MockQueryVisualizationsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockQueryVisualizationsInterface) EXPECT() *MockQueryVisualizationsInterface_Expecter {
	return &MockQueryVisualizationsInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockQueryVisualizationsInterface) Create(ctx context.Context, request sql.CreateVisualizationRequest) (*sql.Visualization, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *sql.Visualization
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.CreateVisualizationRequest) (*sql.Visualization, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sql.CreateVisualizationRequest) *sql.Visualization); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Visualization)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sql.CreateVisualizationRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryVisualizationsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockQueryVisualizationsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.CreateVisualizationRequest
func (_e *MockQueryVisualizationsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockQueryVisualizationsInterface_Create_Call {
	return &MockQueryVisualizationsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockQueryVisualizationsInterface_Create_Call) Run(run func(ctx context.Context, request sql.CreateVisualizationRequest)) *MockQueryVisualizationsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.CreateVisualizationRequest))
	})
	return _c
}

func (_c *MockQueryVisualizationsInterface_Create_Call) Return(_a0 *sql.Visualization, _a1 error) *MockQueryVisualizationsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQueryVisualizationsInterface_Create_Call) RunAndReturn(run func(context.Context, sql.CreateVisualizationRequest) (*sql.Visualization, error)) *MockQueryVisualizationsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockQueryVisualizationsInterface) Delete(ctx context.Context, request sql.DeleteVisualizationRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.DeleteVisualizationRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQueryVisualizationsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockQueryVisualizationsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.DeleteVisualizationRequest
func (_e *MockQueryVisualizationsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockQueryVisualizationsInterface_Delete_Call {
	return &MockQueryVisualizationsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockQueryVisualizationsInterface_Delete_Call) Run(run func(ctx context.Context, request sql.DeleteVisualizationRequest)) *MockQueryVisualizationsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.DeleteVisualizationRequest))
	})
	return _c
}

func (_c *MockQueryVisualizationsInterface_Delete_Call) Return(_a0 error) *MockQueryVisualizationsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQueryVisualizationsInterface_Delete_Call) RunAndReturn(run func(context.Context, sql.DeleteVisualizationRequest) error) *MockQueryVisualizationsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *MockQueryVisualizationsInterface) DeleteById(ctx context.Context, id string) error {
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

// MockQueryVisualizationsInterface_DeleteById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteById'
type MockQueryVisualizationsInterface_DeleteById_Call struct {
	*mock.Call
}

// DeleteById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockQueryVisualizationsInterface_Expecter) DeleteById(ctx interface{}, id interface{}) *MockQueryVisualizationsInterface_DeleteById_Call {
	return &MockQueryVisualizationsInterface_DeleteById_Call{Call: _e.mock.On("DeleteById", ctx, id)}
}

func (_c *MockQueryVisualizationsInterface_DeleteById_Call) Run(run func(ctx context.Context, id string)) *MockQueryVisualizationsInterface_DeleteById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockQueryVisualizationsInterface_DeleteById_Call) Return(_a0 error) *MockQueryVisualizationsInterface_DeleteById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQueryVisualizationsInterface_DeleteById_Call) RunAndReturn(run func(context.Context, string) error) *MockQueryVisualizationsInterface_DeleteById_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockQueryVisualizationsInterface) Impl() sql.QueryVisualizationsService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 sql.QueryVisualizationsService
	if rf, ok := ret.Get(0).(func() sql.QueryVisualizationsService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.QueryVisualizationsService)
		}
	}

	return r0
}

// MockQueryVisualizationsInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockQueryVisualizationsInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockQueryVisualizationsInterface_Expecter) Impl() *MockQueryVisualizationsInterface_Impl_Call {
	return &MockQueryVisualizationsInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockQueryVisualizationsInterface_Impl_Call) Run(run func()) *MockQueryVisualizationsInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockQueryVisualizationsInterface_Impl_Call) Return(_a0 sql.QueryVisualizationsService) *MockQueryVisualizationsInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQueryVisualizationsInterface_Impl_Call) RunAndReturn(run func() sql.QueryVisualizationsService) *MockQueryVisualizationsInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockQueryVisualizationsInterface) Update(ctx context.Context, request sql.UpdateVisualizationRequest) (*sql.Visualization, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *sql.Visualization
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.UpdateVisualizationRequest) (*sql.Visualization, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sql.UpdateVisualizationRequest) *sql.Visualization); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Visualization)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sql.UpdateVisualizationRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryVisualizationsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockQueryVisualizationsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.UpdateVisualizationRequest
func (_e *MockQueryVisualizationsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockQueryVisualizationsInterface_Update_Call {
	return &MockQueryVisualizationsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockQueryVisualizationsInterface_Update_Call) Run(run func(ctx context.Context, request sql.UpdateVisualizationRequest)) *MockQueryVisualizationsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.UpdateVisualizationRequest))
	})
	return _c
}

func (_c *MockQueryVisualizationsInterface_Update_Call) Return(_a0 *sql.Visualization, _a1 error) *MockQueryVisualizationsInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQueryVisualizationsInterface_Update_Call) RunAndReturn(run func(context.Context, sql.UpdateVisualizationRequest) (*sql.Visualization, error)) *MockQueryVisualizationsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockQueryVisualizationsInterface) WithImpl(impl sql.QueryVisualizationsService) sql.QueryVisualizationsInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 sql.QueryVisualizationsInterface
	if rf, ok := ret.Get(0).(func(sql.QueryVisualizationsService) sql.QueryVisualizationsInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.QueryVisualizationsInterface)
		}
	}

	return r0
}

// MockQueryVisualizationsInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockQueryVisualizationsInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl sql.QueryVisualizationsService
func (_e *MockQueryVisualizationsInterface_Expecter) WithImpl(impl interface{}) *MockQueryVisualizationsInterface_WithImpl_Call {
	return &MockQueryVisualizationsInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockQueryVisualizationsInterface_WithImpl_Call) Run(run func(impl sql.QueryVisualizationsService)) *MockQueryVisualizationsInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(sql.QueryVisualizationsService))
	})
	return _c
}

func (_c *MockQueryVisualizationsInterface_WithImpl_Call) Return(_a0 sql.QueryVisualizationsInterface) *MockQueryVisualizationsInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQueryVisualizationsInterface_WithImpl_Call) RunAndReturn(run func(sql.QueryVisualizationsService) sql.QueryVisualizationsInterface) *MockQueryVisualizationsInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockQueryVisualizationsInterface creates a new instance of MockQueryVisualizationsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockQueryVisualizationsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockQueryVisualizationsInterface {
	mock := &MockQueryVisualizationsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
