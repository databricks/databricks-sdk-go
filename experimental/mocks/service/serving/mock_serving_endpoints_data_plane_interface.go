// Code generated by mockery v2.43.0. DO NOT EDIT.

package serving

import (
	context "context"

	serving "github.com/databricks/databricks-sdk-go/service/serving"
	mock "github.com/stretchr/testify/mock"
)

// MockServingEndpointsDataPlaneInterface is an autogenerated mock type for the ServingEndpointsDataPlaneInterface type
type MockServingEndpointsDataPlaneInterface struct {
	mock.Mock
}

type MockServingEndpointsDataPlaneInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockServingEndpointsDataPlaneInterface) EXPECT() *MockServingEndpointsDataPlaneInterface_Expecter {
	return &MockServingEndpointsDataPlaneInterface_Expecter{mock: &_m.Mock}
}

// Impl provides a mock function with given fields:
func (_m *MockServingEndpointsDataPlaneInterface) Impl() serving.ServingEndpointsDataPlaneService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 serving.ServingEndpointsDataPlaneService
	if rf, ok := ret.Get(0).(func() serving.ServingEndpointsDataPlaneService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(serving.ServingEndpointsDataPlaneService)
		}
	}

	return r0
}

// MockServingEndpointsDataPlaneInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockServingEndpointsDataPlaneInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockServingEndpointsDataPlaneInterface_Expecter) Impl() *MockServingEndpointsDataPlaneInterface_Impl_Call {
	return &MockServingEndpointsDataPlaneInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockServingEndpointsDataPlaneInterface_Impl_Call) Run(run func()) *MockServingEndpointsDataPlaneInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockServingEndpointsDataPlaneInterface_Impl_Call) Return(_a0 serving.ServingEndpointsDataPlaneService) *MockServingEndpointsDataPlaneInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServingEndpointsDataPlaneInterface_Impl_Call) RunAndReturn(run func() serving.ServingEndpointsDataPlaneService) *MockServingEndpointsDataPlaneInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// Query provides a mock function with given fields: ctx, request
func (_m *MockServingEndpointsDataPlaneInterface) Query(ctx context.Context, request serving.QueryEndpointInput) (*serving.QueryEndpointResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 *serving.QueryEndpointResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, serving.QueryEndpointInput) (*serving.QueryEndpointResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, serving.QueryEndpointInput) *serving.QueryEndpointResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*serving.QueryEndpointResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, serving.QueryEndpointInput) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockServingEndpointsDataPlaneInterface_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type MockServingEndpointsDataPlaneInterface_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//   - ctx context.Context
//   - request serving.QueryEndpointInput
func (_e *MockServingEndpointsDataPlaneInterface_Expecter) Query(ctx interface{}, request interface{}) *MockServingEndpointsDataPlaneInterface_Query_Call {
	return &MockServingEndpointsDataPlaneInterface_Query_Call{Call: _e.mock.On("Query", ctx, request)}
}

func (_c *MockServingEndpointsDataPlaneInterface_Query_Call) Run(run func(ctx context.Context, request serving.QueryEndpointInput)) *MockServingEndpointsDataPlaneInterface_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(serving.QueryEndpointInput))
	})
	return _c
}

func (_c *MockServingEndpointsDataPlaneInterface_Query_Call) Return(_a0 *serving.QueryEndpointResponse, _a1 error) *MockServingEndpointsDataPlaneInterface_Query_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServingEndpointsDataPlaneInterface_Query_Call) RunAndReturn(run func(context.Context, serving.QueryEndpointInput) (*serving.QueryEndpointResponse, error)) *MockServingEndpointsDataPlaneInterface_Query_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockServingEndpointsDataPlaneInterface) WithImpl(impl serving.ServingEndpointsDataPlaneService) serving.ServingEndpointsDataPlaneInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 serving.ServingEndpointsDataPlaneInterface
	if rf, ok := ret.Get(0).(func(serving.ServingEndpointsDataPlaneService) serving.ServingEndpointsDataPlaneInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(serving.ServingEndpointsDataPlaneInterface)
		}
	}

	return r0
}

// MockServingEndpointsDataPlaneInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockServingEndpointsDataPlaneInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl serving.ServingEndpointsDataPlaneService
func (_e *MockServingEndpointsDataPlaneInterface_Expecter) WithImpl(impl interface{}) *MockServingEndpointsDataPlaneInterface_WithImpl_Call {
	return &MockServingEndpointsDataPlaneInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockServingEndpointsDataPlaneInterface_WithImpl_Call) Run(run func(impl serving.ServingEndpointsDataPlaneService)) *MockServingEndpointsDataPlaneInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(serving.ServingEndpointsDataPlaneService))
	})
	return _c
}

func (_c *MockServingEndpointsDataPlaneInterface_WithImpl_Call) Return(_a0 serving.ServingEndpointsDataPlaneInterface) *MockServingEndpointsDataPlaneInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServingEndpointsDataPlaneInterface_WithImpl_Call) RunAndReturn(run func(serving.ServingEndpointsDataPlaneService) serving.ServingEndpointsDataPlaneInterface) *MockServingEndpointsDataPlaneInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockServingEndpointsDataPlaneInterface creates a new instance of MockServingEndpointsDataPlaneInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockServingEndpointsDataPlaneInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockServingEndpointsDataPlaneInterface {
	mock := &MockServingEndpointsDataPlaneInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
