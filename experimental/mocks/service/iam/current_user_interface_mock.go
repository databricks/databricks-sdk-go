// Code generated by mockery v2.38.0. DO NOT EDIT.

package iam

import (
	context "context"

	iam "github.com/databricks/databricks-sdk-go/service/iam"
	mock "github.com/stretchr/testify/mock"
)

// MockCurrentUserInterface is an autogenerated mock type for the CurrentUserInterface type
type MockCurrentUserInterface struct {
	mock.Mock
}

type MockCurrentUserInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCurrentUserInterface) EXPECT() *MockCurrentUserInterface_Expecter {
	return &MockCurrentUserInterface_Expecter{mock: &_m.Mock}
}

// Impl provides a mock function with given fields:
func (_m *MockCurrentUserInterface) Impl() iam.CurrentUserService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 iam.CurrentUserService
	if rf, ok := ret.Get(0).(func() iam.CurrentUserService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(iam.CurrentUserService)
		}
	}

	return r0
}

// MockCurrentUserInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockCurrentUserInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockCurrentUserInterface_Expecter) Impl() *MockCurrentUserInterface_Impl_Call {
	return &MockCurrentUserInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockCurrentUserInterface_Impl_Call) Run(run func()) *MockCurrentUserInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCurrentUserInterface_Impl_Call) Return(_a0 iam.CurrentUserService) *MockCurrentUserInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCurrentUserInterface_Impl_Call) RunAndReturn(run func() iam.CurrentUserService) *MockCurrentUserInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// Me provides a mock function with given fields: ctx
func (_m *MockCurrentUserInterface) Me(ctx context.Context) (*iam.User, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Me")
	}

	var r0 *iam.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*iam.User, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *iam.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCurrentUserInterface_Me_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Me'
type MockCurrentUserInterface_Me_Call struct {
	*mock.Call
}

// Me is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockCurrentUserInterface_Expecter) Me(ctx interface{}) *MockCurrentUserInterface_Me_Call {
	return &MockCurrentUserInterface_Me_Call{Call: _e.mock.On("Me", ctx)}
}

func (_c *MockCurrentUserInterface_Me_Call) Run(run func(ctx context.Context)) *MockCurrentUserInterface_Me_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockCurrentUserInterface_Me_Call) Return(_a0 *iam.User, _a1 error) *MockCurrentUserInterface_Me_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCurrentUserInterface_Me_Call) RunAndReturn(run func(context.Context) (*iam.User, error)) *MockCurrentUserInterface_Me_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockCurrentUserInterface) WithImpl(impl iam.CurrentUserService) iam.CurrentUserInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 iam.CurrentUserInterface
	if rf, ok := ret.Get(0).(func(iam.CurrentUserService) iam.CurrentUserInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(iam.CurrentUserInterface)
		}
	}

	return r0
}

// MockCurrentUserInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockCurrentUserInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl iam.CurrentUserService
func (_e *MockCurrentUserInterface_Expecter) WithImpl(impl interface{}) *MockCurrentUserInterface_WithImpl_Call {
	return &MockCurrentUserInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockCurrentUserInterface_WithImpl_Call) Run(run func(impl iam.CurrentUserService)) *MockCurrentUserInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(iam.CurrentUserService))
	})
	return _c
}

func (_c *MockCurrentUserInterface_WithImpl_Call) Return(_a0 iam.CurrentUserInterface) *MockCurrentUserInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCurrentUserInterface_WithImpl_Call) RunAndReturn(run func(iam.CurrentUserService) iam.CurrentUserInterface) *MockCurrentUserInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCurrentUserInterface creates a new instance of MockCurrentUserInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCurrentUserInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCurrentUserInterface {
	mock := &MockCurrentUserInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
