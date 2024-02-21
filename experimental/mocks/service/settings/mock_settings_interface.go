// Code generated by mockery v2.39.1. DO NOT EDIT.

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

// DeleteDefaultNamespaceSetting provides a mock function with given fields: ctx, request
func (_m *MockSettingsInterface) DeleteDefaultNamespaceSetting(ctx context.Context, request settings.DeleteDefaultNamespaceSettingRequest) (*settings.DeleteDefaultNamespaceSettingResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDefaultNamespaceSetting")
	}

	var r0 *settings.DeleteDefaultNamespaceSettingResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.DeleteDefaultNamespaceSettingRequest) (*settings.DeleteDefaultNamespaceSettingResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.DeleteDefaultNamespaceSettingRequest) *settings.DeleteDefaultNamespaceSettingResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.DeleteDefaultNamespaceSettingResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.DeleteDefaultNamespaceSettingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSettingsInterface_DeleteDefaultNamespaceSetting_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteDefaultNamespaceSetting'
type MockSettingsInterface_DeleteDefaultNamespaceSetting_Call struct {
	*mock.Call
}

// DeleteDefaultNamespaceSetting is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.DeleteDefaultNamespaceSettingRequest
func (_e *MockSettingsInterface_Expecter) DeleteDefaultNamespaceSetting(ctx interface{}, request interface{}) *MockSettingsInterface_DeleteDefaultNamespaceSetting_Call {
	return &MockSettingsInterface_DeleteDefaultNamespaceSetting_Call{Call: _e.mock.On("DeleteDefaultNamespaceSetting", ctx, request)}
}

func (_c *MockSettingsInterface_DeleteDefaultNamespaceSetting_Call) Run(run func(ctx context.Context, request settings.DeleteDefaultNamespaceSettingRequest)) *MockSettingsInterface_DeleteDefaultNamespaceSetting_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.DeleteDefaultNamespaceSettingRequest))
	})
	return _c
}

func (_c *MockSettingsInterface_DeleteDefaultNamespaceSetting_Call) Return(_a0 *settings.DeleteDefaultNamespaceSettingResponse, _a1 error) *MockSettingsInterface_DeleteDefaultNamespaceSetting_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSettingsInterface_DeleteDefaultNamespaceSetting_Call) RunAndReturn(run func(context.Context, settings.DeleteDefaultNamespaceSettingRequest) (*settings.DeleteDefaultNamespaceSettingResponse, error)) *MockSettingsInterface_DeleteDefaultNamespaceSetting_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteRestrictWorkspaceAdminsSetting provides a mock function with given fields: ctx, request
func (_m *MockSettingsInterface) DeleteRestrictWorkspaceAdminsSetting(ctx context.Context, request settings.DeleteRestrictWorkspaceAdminsSettingRequest) (*settings.DeleteRestrictWorkspaceAdminsSettingResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRestrictWorkspaceAdminsSetting")
	}

	var r0 *settings.DeleteRestrictWorkspaceAdminsSettingResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.DeleteRestrictWorkspaceAdminsSettingRequest) (*settings.DeleteRestrictWorkspaceAdminsSettingResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.DeleteRestrictWorkspaceAdminsSettingRequest) *settings.DeleteRestrictWorkspaceAdminsSettingResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.DeleteRestrictWorkspaceAdminsSettingResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.DeleteRestrictWorkspaceAdminsSettingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSettingsInterface_DeleteRestrictWorkspaceAdminsSetting_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteRestrictWorkspaceAdminsSetting'
type MockSettingsInterface_DeleteRestrictWorkspaceAdminsSetting_Call struct {
	*mock.Call
}

// DeleteRestrictWorkspaceAdminsSetting is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.DeleteRestrictWorkspaceAdminsSettingRequest
func (_e *MockSettingsInterface_Expecter) DeleteRestrictWorkspaceAdminsSetting(ctx interface{}, request interface{}) *MockSettingsInterface_DeleteRestrictWorkspaceAdminsSetting_Call {
	return &MockSettingsInterface_DeleteRestrictWorkspaceAdminsSetting_Call{Call: _e.mock.On("DeleteRestrictWorkspaceAdminsSetting", ctx, request)}
}

func (_c *MockSettingsInterface_DeleteRestrictWorkspaceAdminsSetting_Call) Run(run func(ctx context.Context, request settings.DeleteRestrictWorkspaceAdminsSettingRequest)) *MockSettingsInterface_DeleteRestrictWorkspaceAdminsSetting_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.DeleteRestrictWorkspaceAdminsSettingRequest))
	})
	return _c
}

func (_c *MockSettingsInterface_DeleteRestrictWorkspaceAdminsSetting_Call) Return(_a0 *settings.DeleteRestrictWorkspaceAdminsSettingResponse, _a1 error) *MockSettingsInterface_DeleteRestrictWorkspaceAdminsSetting_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSettingsInterface_DeleteRestrictWorkspaceAdminsSetting_Call) RunAndReturn(run func(context.Context, settings.DeleteRestrictWorkspaceAdminsSettingRequest) (*settings.DeleteRestrictWorkspaceAdminsSettingResponse, error)) *MockSettingsInterface_DeleteRestrictWorkspaceAdminsSetting_Call {
	_c.Call.Return(run)
	return _c
}

// GetDefaultNamespaceSetting provides a mock function with given fields: ctx, request
func (_m *MockSettingsInterface) GetDefaultNamespaceSetting(ctx context.Context, request settings.GetDefaultNamespaceSettingRequest) (*settings.DefaultNamespaceSetting, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetDefaultNamespaceSetting")
	}

	var r0 *settings.DefaultNamespaceSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetDefaultNamespaceSettingRequest) (*settings.DefaultNamespaceSetting, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetDefaultNamespaceSettingRequest) *settings.DefaultNamespaceSetting); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.DefaultNamespaceSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.GetDefaultNamespaceSettingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSettingsInterface_GetDefaultNamespaceSetting_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDefaultNamespaceSetting'
type MockSettingsInterface_GetDefaultNamespaceSetting_Call struct {
	*mock.Call
}

// GetDefaultNamespaceSetting is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.GetDefaultNamespaceSettingRequest
func (_e *MockSettingsInterface_Expecter) GetDefaultNamespaceSetting(ctx interface{}, request interface{}) *MockSettingsInterface_GetDefaultNamespaceSetting_Call {
	return &MockSettingsInterface_GetDefaultNamespaceSetting_Call{Call: _e.mock.On("GetDefaultNamespaceSetting", ctx, request)}
}

func (_c *MockSettingsInterface_GetDefaultNamespaceSetting_Call) Run(run func(ctx context.Context, request settings.GetDefaultNamespaceSettingRequest)) *MockSettingsInterface_GetDefaultNamespaceSetting_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.GetDefaultNamespaceSettingRequest))
	})
	return _c
}

func (_c *MockSettingsInterface_GetDefaultNamespaceSetting_Call) Return(_a0 *settings.DefaultNamespaceSetting, _a1 error) *MockSettingsInterface_GetDefaultNamespaceSetting_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSettingsInterface_GetDefaultNamespaceSetting_Call) RunAndReturn(run func(context.Context, settings.GetDefaultNamespaceSettingRequest) (*settings.DefaultNamespaceSetting, error)) *MockSettingsInterface_GetDefaultNamespaceSetting_Call {
	_c.Call.Return(run)
	return _c
}

// GetRestrictWorkspaceAdminsSetting provides a mock function with given fields: ctx, request
func (_m *MockSettingsInterface) GetRestrictWorkspaceAdminsSetting(ctx context.Context, request settings.GetRestrictWorkspaceAdminsSettingRequest) (*settings.RestrictWorkspaceAdminsSetting, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetRestrictWorkspaceAdminsSetting")
	}

	var r0 *settings.RestrictWorkspaceAdminsSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetRestrictWorkspaceAdminsSettingRequest) (*settings.RestrictWorkspaceAdminsSetting, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetRestrictWorkspaceAdminsSettingRequest) *settings.RestrictWorkspaceAdminsSetting); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.RestrictWorkspaceAdminsSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.GetRestrictWorkspaceAdminsSettingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSettingsInterface_GetRestrictWorkspaceAdminsSetting_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRestrictWorkspaceAdminsSetting'
type MockSettingsInterface_GetRestrictWorkspaceAdminsSetting_Call struct {
	*mock.Call
}

// GetRestrictWorkspaceAdminsSetting is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.GetRestrictWorkspaceAdminsSettingRequest
func (_e *MockSettingsInterface_Expecter) GetRestrictWorkspaceAdminsSetting(ctx interface{}, request interface{}) *MockSettingsInterface_GetRestrictWorkspaceAdminsSetting_Call {
	return &MockSettingsInterface_GetRestrictWorkspaceAdminsSetting_Call{Call: _e.mock.On("GetRestrictWorkspaceAdminsSetting", ctx, request)}
}

func (_c *MockSettingsInterface_GetRestrictWorkspaceAdminsSetting_Call) Run(run func(ctx context.Context, request settings.GetRestrictWorkspaceAdminsSettingRequest)) *MockSettingsInterface_GetRestrictWorkspaceAdminsSetting_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.GetRestrictWorkspaceAdminsSettingRequest))
	})
	return _c
}

func (_c *MockSettingsInterface_GetRestrictWorkspaceAdminsSetting_Call) Return(_a0 *settings.RestrictWorkspaceAdminsSetting, _a1 error) *MockSettingsInterface_GetRestrictWorkspaceAdminsSetting_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSettingsInterface_GetRestrictWorkspaceAdminsSetting_Call) RunAndReturn(run func(context.Context, settings.GetRestrictWorkspaceAdminsSettingRequest) (*settings.RestrictWorkspaceAdminsSetting, error)) *MockSettingsInterface_GetRestrictWorkspaceAdminsSetting_Call {
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

// UpdateDefaultNamespaceSetting provides a mock function with given fields: ctx, request
func (_m *MockSettingsInterface) UpdateDefaultNamespaceSetting(ctx context.Context, request settings.UpdateDefaultNamespaceSettingRequest) (*settings.DefaultNamespaceSetting, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDefaultNamespaceSetting")
	}

	var r0 *settings.DefaultNamespaceSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateDefaultNamespaceSettingRequest) (*settings.DefaultNamespaceSetting, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateDefaultNamespaceSettingRequest) *settings.DefaultNamespaceSetting); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.DefaultNamespaceSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.UpdateDefaultNamespaceSettingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSettingsInterface_UpdateDefaultNamespaceSetting_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateDefaultNamespaceSetting'
type MockSettingsInterface_UpdateDefaultNamespaceSetting_Call struct {
	*mock.Call
}

// UpdateDefaultNamespaceSetting is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.UpdateDefaultNamespaceSettingRequest
func (_e *MockSettingsInterface_Expecter) UpdateDefaultNamespaceSetting(ctx interface{}, request interface{}) *MockSettingsInterface_UpdateDefaultNamespaceSetting_Call {
	return &MockSettingsInterface_UpdateDefaultNamespaceSetting_Call{Call: _e.mock.On("UpdateDefaultNamespaceSetting", ctx, request)}
}

func (_c *MockSettingsInterface_UpdateDefaultNamespaceSetting_Call) Run(run func(ctx context.Context, request settings.UpdateDefaultNamespaceSettingRequest)) *MockSettingsInterface_UpdateDefaultNamespaceSetting_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.UpdateDefaultNamespaceSettingRequest))
	})
	return _c
}

func (_c *MockSettingsInterface_UpdateDefaultNamespaceSetting_Call) Return(_a0 *settings.DefaultNamespaceSetting, _a1 error) *MockSettingsInterface_UpdateDefaultNamespaceSetting_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSettingsInterface_UpdateDefaultNamespaceSetting_Call) RunAndReturn(run func(context.Context, settings.UpdateDefaultNamespaceSettingRequest) (*settings.DefaultNamespaceSetting, error)) *MockSettingsInterface_UpdateDefaultNamespaceSetting_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateRestrictWorkspaceAdminsSetting provides a mock function with given fields: ctx, request
func (_m *MockSettingsInterface) UpdateRestrictWorkspaceAdminsSetting(ctx context.Context, request settings.UpdateRestrictWorkspaceAdminsSettingRequest) (*settings.RestrictWorkspaceAdminsSetting, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRestrictWorkspaceAdminsSetting")
	}

	var r0 *settings.RestrictWorkspaceAdminsSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateRestrictWorkspaceAdminsSettingRequest) (*settings.RestrictWorkspaceAdminsSetting, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateRestrictWorkspaceAdminsSettingRequest) *settings.RestrictWorkspaceAdminsSetting); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.RestrictWorkspaceAdminsSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.UpdateRestrictWorkspaceAdminsSettingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSettingsInterface_UpdateRestrictWorkspaceAdminsSetting_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateRestrictWorkspaceAdminsSetting'
type MockSettingsInterface_UpdateRestrictWorkspaceAdminsSetting_Call struct {
	*mock.Call
}

// UpdateRestrictWorkspaceAdminsSetting is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.UpdateRestrictWorkspaceAdminsSettingRequest
func (_e *MockSettingsInterface_Expecter) UpdateRestrictWorkspaceAdminsSetting(ctx interface{}, request interface{}) *MockSettingsInterface_UpdateRestrictWorkspaceAdminsSetting_Call {
	return &MockSettingsInterface_UpdateRestrictWorkspaceAdminsSetting_Call{Call: _e.mock.On("UpdateRestrictWorkspaceAdminsSetting", ctx, request)}
}

func (_c *MockSettingsInterface_UpdateRestrictWorkspaceAdminsSetting_Call) Run(run func(ctx context.Context, request settings.UpdateRestrictWorkspaceAdminsSettingRequest)) *MockSettingsInterface_UpdateRestrictWorkspaceAdminsSetting_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.UpdateRestrictWorkspaceAdminsSettingRequest))
	})
	return _c
}

func (_c *MockSettingsInterface_UpdateRestrictWorkspaceAdminsSetting_Call) Return(_a0 *settings.RestrictWorkspaceAdminsSetting, _a1 error) *MockSettingsInterface_UpdateRestrictWorkspaceAdminsSetting_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSettingsInterface_UpdateRestrictWorkspaceAdminsSetting_Call) RunAndReturn(run func(context.Context, settings.UpdateRestrictWorkspaceAdminsSettingRequest) (*settings.RestrictWorkspaceAdminsSetting, error)) *MockSettingsInterface_UpdateRestrictWorkspaceAdminsSetting_Call {
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
