// Code generated by mockery v2.43.0. DO NOT EDIT.

package settings

import (
	context "context"

	settings "github.com/databricks/databricks-sdk-go/service/settings"
	mock "github.com/stretchr/testify/mock"
)

// MockCredentialsManagerInterface is an autogenerated mock type for the CredentialsManagerInterface type
type MockCredentialsManagerInterface struct {
	mock.Mock
}

type MockCredentialsManagerInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCredentialsManagerInterface) EXPECT() *MockCredentialsManagerInterface_Expecter {
	return &MockCredentialsManagerInterface_Expecter{mock: &_m.Mock}
}

// ExchangeToken provides a mock function with given fields: ctx, request
func (_m *MockCredentialsManagerInterface) ExchangeToken(ctx context.Context, request settings.ExchangeTokenRequest) (*settings.ExchangeTokenResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ExchangeToken")
	}

	var r0 *settings.ExchangeTokenResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.ExchangeTokenRequest) (*settings.ExchangeTokenResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.ExchangeTokenRequest) *settings.ExchangeTokenResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.ExchangeTokenResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.ExchangeTokenRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCredentialsManagerInterface_ExchangeToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExchangeToken'
type MockCredentialsManagerInterface_ExchangeToken_Call struct {
	*mock.Call
}

// ExchangeToken is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.ExchangeTokenRequest
func (_e *MockCredentialsManagerInterface_Expecter) ExchangeToken(ctx interface{}, request interface{}) *MockCredentialsManagerInterface_ExchangeToken_Call {
	return &MockCredentialsManagerInterface_ExchangeToken_Call{Call: _e.mock.On("ExchangeToken", ctx, request)}
}

func (_c *MockCredentialsManagerInterface_ExchangeToken_Call) Run(run func(ctx context.Context, request settings.ExchangeTokenRequest)) *MockCredentialsManagerInterface_ExchangeToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.ExchangeTokenRequest))
	})
	return _c
}

func (_c *MockCredentialsManagerInterface_ExchangeToken_Call) Return(_a0 *settings.ExchangeTokenResponse, _a1 error) *MockCredentialsManagerInterface_ExchangeToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCredentialsManagerInterface_ExchangeToken_Call) RunAndReturn(run func(context.Context, settings.ExchangeTokenRequest) (*settings.ExchangeTokenResponse, error)) *MockCredentialsManagerInterface_ExchangeToken_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockCredentialsManagerInterface) Impl() settings.CredentialsManagerService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 settings.CredentialsManagerService
	if rf, ok := ret.Get(0).(func() settings.CredentialsManagerService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.CredentialsManagerService)
		}
	}

	return r0
}

// MockCredentialsManagerInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockCredentialsManagerInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockCredentialsManagerInterface_Expecter) Impl() *MockCredentialsManagerInterface_Impl_Call {
	return &MockCredentialsManagerInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockCredentialsManagerInterface_Impl_Call) Run(run func()) *MockCredentialsManagerInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCredentialsManagerInterface_Impl_Call) Return(_a0 settings.CredentialsManagerService) *MockCredentialsManagerInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCredentialsManagerInterface_Impl_Call) RunAndReturn(run func() settings.CredentialsManagerService) *MockCredentialsManagerInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockCredentialsManagerInterface) WithImpl(impl settings.CredentialsManagerService) settings.CredentialsManagerInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 settings.CredentialsManagerInterface
	if rf, ok := ret.Get(0).(func(settings.CredentialsManagerService) settings.CredentialsManagerInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.CredentialsManagerInterface)
		}
	}

	return r0
}

// MockCredentialsManagerInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockCredentialsManagerInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl settings.CredentialsManagerService
func (_e *MockCredentialsManagerInterface_Expecter) WithImpl(impl interface{}) *MockCredentialsManagerInterface_WithImpl_Call {
	return &MockCredentialsManagerInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockCredentialsManagerInterface_WithImpl_Call) Run(run func(impl settings.CredentialsManagerService)) *MockCredentialsManagerInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(settings.CredentialsManagerService))
	})
	return _c
}

func (_c *MockCredentialsManagerInterface_WithImpl_Call) Return(_a0 settings.CredentialsManagerInterface) *MockCredentialsManagerInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCredentialsManagerInterface_WithImpl_Call) RunAndReturn(run func(settings.CredentialsManagerService) settings.CredentialsManagerInterface) *MockCredentialsManagerInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCredentialsManagerInterface creates a new instance of MockCredentialsManagerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCredentialsManagerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCredentialsManagerInterface {
	mock := &MockCredentialsManagerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
