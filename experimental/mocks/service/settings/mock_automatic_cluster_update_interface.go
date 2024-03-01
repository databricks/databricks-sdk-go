// Code generated by mockery v2.39.1. DO NOT EDIT.

package settings

import (
	context "context"

	settings "github.com/databricks/databricks-sdk-go/service/settings"
	mock "github.com/stretchr/testify/mock"
)

// MockAutomaticClusterUpdateInterface is an autogenerated mock type for the AutomaticClusterUpdateInterface type
type MockAutomaticClusterUpdateInterface struct {
	mock.Mock
}

type MockAutomaticClusterUpdateInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAutomaticClusterUpdateInterface) EXPECT() *MockAutomaticClusterUpdateInterface_Expecter {
	return &MockAutomaticClusterUpdateInterface_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockAutomaticClusterUpdateInterface) Get(ctx context.Context, request settings.GetAutomaticClusterUpdateRequest) (*settings.AutomaticClusterUpdateSetting, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *settings.AutomaticClusterUpdateSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetAutomaticClusterUpdateRequest) (*settings.AutomaticClusterUpdateSetting, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetAutomaticClusterUpdateRequest) *settings.AutomaticClusterUpdateSetting); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.AutomaticClusterUpdateSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.GetAutomaticClusterUpdateRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAutomaticClusterUpdateInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockAutomaticClusterUpdateInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.GetAutomaticClusterUpdateRequest
func (_e *MockAutomaticClusterUpdateInterface_Expecter) Get(ctx interface{}, request interface{}) *MockAutomaticClusterUpdateInterface_Get_Call {
	return &MockAutomaticClusterUpdateInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockAutomaticClusterUpdateInterface_Get_Call) Run(run func(ctx context.Context, request settings.GetAutomaticClusterUpdateRequest)) *MockAutomaticClusterUpdateInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.GetAutomaticClusterUpdateRequest))
	})
	return _c
}

func (_c *MockAutomaticClusterUpdateInterface_Get_Call) Return(_a0 *settings.AutomaticClusterUpdateSetting, _a1 error) *MockAutomaticClusterUpdateInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAutomaticClusterUpdateInterface_Get_Call) RunAndReturn(run func(context.Context, settings.GetAutomaticClusterUpdateRequest) (*settings.AutomaticClusterUpdateSetting, error)) *MockAutomaticClusterUpdateInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockAutomaticClusterUpdateInterface) Impl() settings.AutomaticClusterUpdateService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 settings.AutomaticClusterUpdateService
	if rf, ok := ret.Get(0).(func() settings.AutomaticClusterUpdateService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.AutomaticClusterUpdateService)
		}
	}

	return r0
}

// MockAutomaticClusterUpdateInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockAutomaticClusterUpdateInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockAutomaticClusterUpdateInterface_Expecter) Impl() *MockAutomaticClusterUpdateInterface_Impl_Call {
	return &MockAutomaticClusterUpdateInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockAutomaticClusterUpdateInterface_Impl_Call) Run(run func()) *MockAutomaticClusterUpdateInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAutomaticClusterUpdateInterface_Impl_Call) Return(_a0 settings.AutomaticClusterUpdateService) *MockAutomaticClusterUpdateInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAutomaticClusterUpdateInterface_Impl_Call) RunAndReturn(run func() settings.AutomaticClusterUpdateService) *MockAutomaticClusterUpdateInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockAutomaticClusterUpdateInterface) Update(ctx context.Context, request settings.UpdateAutomaticClusterUpdateSettingRequest) (*settings.AutomaticClusterUpdateSetting, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *settings.AutomaticClusterUpdateSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateAutomaticClusterUpdateSettingRequest) (*settings.AutomaticClusterUpdateSetting, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateAutomaticClusterUpdateSettingRequest) *settings.AutomaticClusterUpdateSetting); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.AutomaticClusterUpdateSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.UpdateAutomaticClusterUpdateSettingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAutomaticClusterUpdateInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockAutomaticClusterUpdateInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.UpdateAutomaticClusterUpdateSettingRequest
func (_e *MockAutomaticClusterUpdateInterface_Expecter) Update(ctx interface{}, request interface{}) *MockAutomaticClusterUpdateInterface_Update_Call {
	return &MockAutomaticClusterUpdateInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockAutomaticClusterUpdateInterface_Update_Call) Run(run func(ctx context.Context, request settings.UpdateAutomaticClusterUpdateSettingRequest)) *MockAutomaticClusterUpdateInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.UpdateAutomaticClusterUpdateSettingRequest))
	})
	return _c
}

func (_c *MockAutomaticClusterUpdateInterface_Update_Call) Return(_a0 *settings.AutomaticClusterUpdateSetting, _a1 error) *MockAutomaticClusterUpdateInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAutomaticClusterUpdateInterface_Update_Call) RunAndReturn(run func(context.Context, settings.UpdateAutomaticClusterUpdateSettingRequest) (*settings.AutomaticClusterUpdateSetting, error)) *MockAutomaticClusterUpdateInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockAutomaticClusterUpdateInterface) WithImpl(impl settings.AutomaticClusterUpdateService) settings.AutomaticClusterUpdateInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 settings.AutomaticClusterUpdateInterface
	if rf, ok := ret.Get(0).(func(settings.AutomaticClusterUpdateService) settings.AutomaticClusterUpdateInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.AutomaticClusterUpdateInterface)
		}
	}

	return r0
}

// MockAutomaticClusterUpdateInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockAutomaticClusterUpdateInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl settings.AutomaticClusterUpdateService
func (_e *MockAutomaticClusterUpdateInterface_Expecter) WithImpl(impl interface{}) *MockAutomaticClusterUpdateInterface_WithImpl_Call {
	return &MockAutomaticClusterUpdateInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockAutomaticClusterUpdateInterface_WithImpl_Call) Run(run func(impl settings.AutomaticClusterUpdateService)) *MockAutomaticClusterUpdateInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(settings.AutomaticClusterUpdateService))
	})
	return _c
}

func (_c *MockAutomaticClusterUpdateInterface_WithImpl_Call) Return(_a0 settings.AutomaticClusterUpdateInterface) *MockAutomaticClusterUpdateInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAutomaticClusterUpdateInterface_WithImpl_Call) RunAndReturn(run func(settings.AutomaticClusterUpdateService) settings.AutomaticClusterUpdateInterface) *MockAutomaticClusterUpdateInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAutomaticClusterUpdateInterface creates a new instance of MockAutomaticClusterUpdateInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAutomaticClusterUpdateInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAutomaticClusterUpdateInterface {
	mock := &MockAutomaticClusterUpdateInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
