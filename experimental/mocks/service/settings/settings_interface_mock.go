// Code generated by mockery v2.38.0. DO NOT EDIT.

package settings

import (
	context "context"

	settings "github.com/databricks/databricks-sdk-go/service/settings"
	mock "github.com/stretchr/testify/mock"
)

// MockSettingsInterface is an autogenerated mock type for the SettingsInterface type
type MockSettingsInterface struct {
	mock.Mock
}

type MockSettingsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSettingsInterface) EXPECT() *MockSettingsInterface_Expecter {
	return &MockSettingsInterface_Expecter{mock: &_m.Mock}
}

// DeleteDefaultWorkspaceNamespace provides a mock function with given fields: ctx, request
func (_m *MockSettingsInterface) DeleteDefaultWorkspaceNamespace(ctx context.Context, request settings.DeleteDefaultWorkspaceNamespaceRequest) (*settings.DeleteDefaultWorkspaceNamespaceResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDefaultWorkspaceNamespace")
	}

	var r0 *settings.DeleteDefaultWorkspaceNamespaceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.DeleteDefaultWorkspaceNamespaceRequest) (*settings.DeleteDefaultWorkspaceNamespaceResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.DeleteDefaultWorkspaceNamespaceRequest) *settings.DeleteDefaultWorkspaceNamespaceResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.DeleteDefaultWorkspaceNamespaceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.DeleteDefaultWorkspaceNamespaceRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSettingsInterface_DeleteDefaultWorkspaceNamespace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteDefaultWorkspaceNamespace'
type MockSettingsInterface_DeleteDefaultWorkspaceNamespace_Call struct {
	*mock.Call
}

// DeleteDefaultWorkspaceNamespace is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.DeleteDefaultWorkspaceNamespaceRequest
func (_e *MockSettingsInterface_Expecter) DeleteDefaultWorkspaceNamespace(ctx interface{}, request interface{}) *MockSettingsInterface_DeleteDefaultWorkspaceNamespace_Call {
	return &MockSettingsInterface_DeleteDefaultWorkspaceNamespace_Call{Call: _e.mock.On("DeleteDefaultWorkspaceNamespace", ctx, request)}
}

func (_c *MockSettingsInterface_DeleteDefaultWorkspaceNamespace_Call) Run(run func(ctx context.Context, request settings.DeleteDefaultWorkspaceNamespaceRequest)) *MockSettingsInterface_DeleteDefaultWorkspaceNamespace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.DeleteDefaultWorkspaceNamespaceRequest))
	})
	return _c
}

func (_c *MockSettingsInterface_DeleteDefaultWorkspaceNamespace_Call) Return(_a0 *settings.DeleteDefaultWorkspaceNamespaceResponse, _a1 error) *MockSettingsInterface_DeleteDefaultWorkspaceNamespace_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSettingsInterface_DeleteDefaultWorkspaceNamespace_Call) RunAndReturn(run func(context.Context, settings.DeleteDefaultWorkspaceNamespaceRequest) (*settings.DeleteDefaultWorkspaceNamespaceResponse, error)) *MockSettingsInterface_DeleteDefaultWorkspaceNamespace_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockSettingsInterface) Impl() settings.SettingsService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 settings.SettingsService
	if rf, ok := ret.Get(0).(func() settings.SettingsService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.SettingsService)
		}
	}

	return r0
}

// MockSettingsInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockSettingsInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockSettingsInterface_Expecter) Impl() *MockSettingsInterface_Impl_Call {
	return &MockSettingsInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockSettingsInterface_Impl_Call) Run(run func()) *MockSettingsInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSettingsInterface_Impl_Call) Return(_a0 settings.SettingsService) *MockSettingsInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSettingsInterface_Impl_Call) RunAndReturn(run func() settings.SettingsService) *MockSettingsInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// ReadDefaultWorkspaceNamespace provides a mock function with given fields: ctx, request
func (_m *MockSettingsInterface) ReadDefaultWorkspaceNamespace(ctx context.Context, request settings.ReadDefaultWorkspaceNamespaceRequest) (*settings.DefaultNamespaceSetting, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ReadDefaultWorkspaceNamespace")
	}

	var r0 *settings.DefaultNamespaceSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.ReadDefaultWorkspaceNamespaceRequest) (*settings.DefaultNamespaceSetting, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.ReadDefaultWorkspaceNamespaceRequest) *settings.DefaultNamespaceSetting); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.DefaultNamespaceSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.ReadDefaultWorkspaceNamespaceRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSettingsInterface_ReadDefaultWorkspaceNamespace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadDefaultWorkspaceNamespace'
type MockSettingsInterface_ReadDefaultWorkspaceNamespace_Call struct {
	*mock.Call
}

// ReadDefaultWorkspaceNamespace is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.ReadDefaultWorkspaceNamespaceRequest
func (_e *MockSettingsInterface_Expecter) ReadDefaultWorkspaceNamespace(ctx interface{}, request interface{}) *MockSettingsInterface_ReadDefaultWorkspaceNamespace_Call {
	return &MockSettingsInterface_ReadDefaultWorkspaceNamespace_Call{Call: _e.mock.On("ReadDefaultWorkspaceNamespace", ctx, request)}
}

func (_c *MockSettingsInterface_ReadDefaultWorkspaceNamespace_Call) Run(run func(ctx context.Context, request settings.ReadDefaultWorkspaceNamespaceRequest)) *MockSettingsInterface_ReadDefaultWorkspaceNamespace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.ReadDefaultWorkspaceNamespaceRequest))
	})
	return _c
}

func (_c *MockSettingsInterface_ReadDefaultWorkspaceNamespace_Call) Return(_a0 *settings.DefaultNamespaceSetting, _a1 error) *MockSettingsInterface_ReadDefaultWorkspaceNamespace_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSettingsInterface_ReadDefaultWorkspaceNamespace_Call) RunAndReturn(run func(context.Context, settings.ReadDefaultWorkspaceNamespaceRequest) (*settings.DefaultNamespaceSetting, error)) *MockSettingsInterface_ReadDefaultWorkspaceNamespace_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateDefaultWorkspaceNamespace provides a mock function with given fields: ctx, request
func (_m *MockSettingsInterface) UpdateDefaultWorkspaceNamespace(ctx context.Context, request settings.UpdateDefaultWorkspaceNamespaceRequest) (*settings.DefaultNamespaceSetting, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDefaultWorkspaceNamespace")
	}

	var r0 *settings.DefaultNamespaceSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateDefaultWorkspaceNamespaceRequest) (*settings.DefaultNamespaceSetting, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateDefaultWorkspaceNamespaceRequest) *settings.DefaultNamespaceSetting); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.DefaultNamespaceSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.UpdateDefaultWorkspaceNamespaceRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSettingsInterface_UpdateDefaultWorkspaceNamespace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateDefaultWorkspaceNamespace'
type MockSettingsInterface_UpdateDefaultWorkspaceNamespace_Call struct {
	*mock.Call
}

// UpdateDefaultWorkspaceNamespace is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.UpdateDefaultWorkspaceNamespaceRequest
func (_e *MockSettingsInterface_Expecter) UpdateDefaultWorkspaceNamespace(ctx interface{}, request interface{}) *MockSettingsInterface_UpdateDefaultWorkspaceNamespace_Call {
	return &MockSettingsInterface_UpdateDefaultWorkspaceNamespace_Call{Call: _e.mock.On("UpdateDefaultWorkspaceNamespace", ctx, request)}
}

func (_c *MockSettingsInterface_UpdateDefaultWorkspaceNamespace_Call) Run(run func(ctx context.Context, request settings.UpdateDefaultWorkspaceNamespaceRequest)) *MockSettingsInterface_UpdateDefaultWorkspaceNamespace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.UpdateDefaultWorkspaceNamespaceRequest))
	})
	return _c
}

func (_c *MockSettingsInterface_UpdateDefaultWorkspaceNamespace_Call) Return(_a0 *settings.DefaultNamespaceSetting, _a1 error) *MockSettingsInterface_UpdateDefaultWorkspaceNamespace_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSettingsInterface_UpdateDefaultWorkspaceNamespace_Call) RunAndReturn(run func(context.Context, settings.UpdateDefaultWorkspaceNamespaceRequest) (*settings.DefaultNamespaceSetting, error)) *MockSettingsInterface_UpdateDefaultWorkspaceNamespace_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockSettingsInterface) WithImpl(impl settings.SettingsService) settings.SettingsInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 settings.SettingsInterface
	if rf, ok := ret.Get(0).(func(settings.SettingsService) settings.SettingsInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.SettingsInterface)
		}
	}

	return r0
}

// MockSettingsInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockSettingsInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl settings.SettingsService
func (_e *MockSettingsInterface_Expecter) WithImpl(impl interface{}) *MockSettingsInterface_WithImpl_Call {
	return &MockSettingsInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockSettingsInterface_WithImpl_Call) Run(run func(impl settings.SettingsService)) *MockSettingsInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(settings.SettingsService))
	})
	return _c
}

func (_c *MockSettingsInterface_WithImpl_Call) Return(_a0 settings.SettingsInterface) *MockSettingsInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSettingsInterface_WithImpl_Call) RunAndReturn(run func(settings.SettingsService) settings.SettingsInterface) *MockSettingsInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSettingsInterface creates a new instance of MockSettingsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSettingsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSettingsInterface {
	mock := &MockSettingsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
