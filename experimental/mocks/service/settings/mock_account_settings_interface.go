// Code generated by mockery v2.43.0. DO NOT EDIT.

package settings

import (
	settings "github.com/databricks/databricks-sdk-go/service/settings"
	mock "github.com/stretchr/testify/mock"
)

// MockAccountSettingsInterface is an autogenerated mock type for the AccountSettingsInterface type
type MockAccountSettingsInterface struct {
	mock.Mock
}

type MockAccountSettingsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAccountSettingsInterface) EXPECT() *MockAccountSettingsInterface_Expecter {
	return &MockAccountSettingsInterface_Expecter{mock: &_m.Mock}
}

// CspEnablementAccount provides a mock function with given fields:
func (_m *MockAccountSettingsInterface) CspEnablementAccount() settings.CspEnablementAccountInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CspEnablementAccount")
	}

	var r0 settings.CspEnablementAccountInterface
	if rf, ok := ret.Get(0).(func() settings.CspEnablementAccountInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.CspEnablementAccountInterface)
		}
	}

	return r0
}

// MockAccountSettingsInterface_CspEnablementAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CspEnablementAccount'
type MockAccountSettingsInterface_CspEnablementAccount_Call struct {
	*mock.Call
}

// CspEnablementAccount is a helper method to define mock.On call
func (_e *MockAccountSettingsInterface_Expecter) CspEnablementAccount() *MockAccountSettingsInterface_CspEnablementAccount_Call {
	return &MockAccountSettingsInterface_CspEnablementAccount_Call{Call: _e.mock.On("CspEnablementAccount")}
}

func (_c *MockAccountSettingsInterface_CspEnablementAccount_Call) Run(run func()) *MockAccountSettingsInterface_CspEnablementAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAccountSettingsInterface_CspEnablementAccount_Call) Return(_a0 settings.CspEnablementAccountInterface) *MockAccountSettingsInterface_CspEnablementAccount_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountSettingsInterface_CspEnablementAccount_Call) RunAndReturn(run func() settings.CspEnablementAccountInterface) *MockAccountSettingsInterface_CspEnablementAccount_Call {
	_c.Call.Return(run)
	return _c
}

// EsmEnablementAccount provides a mock function with given fields:
func (_m *MockAccountSettingsInterface) EsmEnablementAccount() settings.EsmEnablementAccountInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EsmEnablementAccount")
	}

	var r0 settings.EsmEnablementAccountInterface
	if rf, ok := ret.Get(0).(func() settings.EsmEnablementAccountInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.EsmEnablementAccountInterface)
		}
	}

	return r0
}

// MockAccountSettingsInterface_EsmEnablementAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EsmEnablementAccount'
type MockAccountSettingsInterface_EsmEnablementAccount_Call struct {
	*mock.Call
}

// EsmEnablementAccount is a helper method to define mock.On call
func (_e *MockAccountSettingsInterface_Expecter) EsmEnablementAccount() *MockAccountSettingsInterface_EsmEnablementAccount_Call {
	return &MockAccountSettingsInterface_EsmEnablementAccount_Call{Call: _e.mock.On("EsmEnablementAccount")}
}

func (_c *MockAccountSettingsInterface_EsmEnablementAccount_Call) Run(run func()) *MockAccountSettingsInterface_EsmEnablementAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAccountSettingsInterface_EsmEnablementAccount_Call) Return(_a0 settings.EsmEnablementAccountInterface) *MockAccountSettingsInterface_EsmEnablementAccount_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountSettingsInterface_EsmEnablementAccount_Call) RunAndReturn(run func() settings.EsmEnablementAccountInterface) *MockAccountSettingsInterface_EsmEnablementAccount_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockAccountSettingsInterface) Impl() settings.AccountSettingsService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 settings.AccountSettingsService
	if rf, ok := ret.Get(0).(func() settings.AccountSettingsService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.AccountSettingsService)
		}
	}

	return r0
}

// MockAccountSettingsInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockAccountSettingsInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockAccountSettingsInterface_Expecter) Impl() *MockAccountSettingsInterface_Impl_Call {
	return &MockAccountSettingsInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockAccountSettingsInterface_Impl_Call) Run(run func()) *MockAccountSettingsInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAccountSettingsInterface_Impl_Call) Return(_a0 settings.AccountSettingsService) *MockAccountSettingsInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountSettingsInterface_Impl_Call) RunAndReturn(run func() settings.AccountSettingsService) *MockAccountSettingsInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// PersonalCompute provides a mock function with given fields:
func (_m *MockAccountSettingsInterface) PersonalCompute() settings.PersonalComputeInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for PersonalCompute")
	}

	var r0 settings.PersonalComputeInterface
	if rf, ok := ret.Get(0).(func() settings.PersonalComputeInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.PersonalComputeInterface)
		}
	}

	return r0
}

// MockAccountSettingsInterface_PersonalCompute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PersonalCompute'
type MockAccountSettingsInterface_PersonalCompute_Call struct {
	*mock.Call
}

// PersonalCompute is a helper method to define mock.On call
func (_e *MockAccountSettingsInterface_Expecter) PersonalCompute() *MockAccountSettingsInterface_PersonalCompute_Call {
	return &MockAccountSettingsInterface_PersonalCompute_Call{Call: _e.mock.On("PersonalCompute")}
}

func (_c *MockAccountSettingsInterface_PersonalCompute_Call) Run(run func()) *MockAccountSettingsInterface_PersonalCompute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAccountSettingsInterface_PersonalCompute_Call) Return(_a0 settings.PersonalComputeInterface) *MockAccountSettingsInterface_PersonalCompute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountSettingsInterface_PersonalCompute_Call) RunAndReturn(run func() settings.PersonalComputeInterface) *MockAccountSettingsInterface_PersonalCompute_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockAccountSettingsInterface) WithImpl(impl settings.AccountSettingsService) settings.AccountSettingsInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 settings.AccountSettingsInterface
	if rf, ok := ret.Get(0).(func(settings.AccountSettingsService) settings.AccountSettingsInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.AccountSettingsInterface)
		}
	}

	return r0
}

// MockAccountSettingsInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockAccountSettingsInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl settings.AccountSettingsService
func (_e *MockAccountSettingsInterface_Expecter) WithImpl(impl interface{}) *MockAccountSettingsInterface_WithImpl_Call {
	return &MockAccountSettingsInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockAccountSettingsInterface_WithImpl_Call) Run(run func(impl settings.AccountSettingsService)) *MockAccountSettingsInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(settings.AccountSettingsService))
	})
	return _c
}

func (_c *MockAccountSettingsInterface_WithImpl_Call) Return(_a0 settings.AccountSettingsInterface) *MockAccountSettingsInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountSettingsInterface_WithImpl_Call) RunAndReturn(run func(settings.AccountSettingsService) settings.AccountSettingsInterface) *MockAccountSettingsInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAccountSettingsInterface creates a new instance of MockAccountSettingsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAccountSettingsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAccountSettingsInterface {
	mock := &MockAccountSettingsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
