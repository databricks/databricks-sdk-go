// Code generated by mockery v2.43.0. DO NOT EDIT.

package settings

import (
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

// AutomaticClusterUpdate provides a mock function with given fields:
func (_m *MockSettingsInterface) AutomaticClusterUpdate() settings.AutomaticClusterUpdateInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AutomaticClusterUpdate")
	}

	var r0 settings.AutomaticClusterUpdateInterface
	if rf, ok := ret.Get(0).(func() settings.AutomaticClusterUpdateInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.AutomaticClusterUpdateInterface)
		}
	}

	return r0
}

// MockSettingsInterface_AutomaticClusterUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AutomaticClusterUpdate'
type MockSettingsInterface_AutomaticClusterUpdate_Call struct {
	*mock.Call
}

// AutomaticClusterUpdate is a helper method to define mock.On call
func (_e *MockSettingsInterface_Expecter) AutomaticClusterUpdate() *MockSettingsInterface_AutomaticClusterUpdate_Call {
	return &MockSettingsInterface_AutomaticClusterUpdate_Call{Call: _e.mock.On("AutomaticClusterUpdate")}
}

func (_c *MockSettingsInterface_AutomaticClusterUpdate_Call) Run(run func()) *MockSettingsInterface_AutomaticClusterUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSettingsInterface_AutomaticClusterUpdate_Call) Return(_a0 settings.AutomaticClusterUpdateInterface) *MockSettingsInterface_AutomaticClusterUpdate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSettingsInterface_AutomaticClusterUpdate_Call) RunAndReturn(run func() settings.AutomaticClusterUpdateInterface) *MockSettingsInterface_AutomaticClusterUpdate_Call {
	_c.Call.Return(run)
	return _c
}

// ComplianceSecurityProfile provides a mock function with given fields:
func (_m *MockSettingsInterface) ComplianceSecurityProfile() settings.ComplianceSecurityProfileInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ComplianceSecurityProfile")
	}

	var r0 settings.ComplianceSecurityProfileInterface
	if rf, ok := ret.Get(0).(func() settings.ComplianceSecurityProfileInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.ComplianceSecurityProfileInterface)
		}
	}

	return r0
}

// MockSettingsInterface_ComplianceSecurityProfile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ComplianceSecurityProfile'
type MockSettingsInterface_ComplianceSecurityProfile_Call struct {
	*mock.Call
}

// ComplianceSecurityProfile is a helper method to define mock.On call
func (_e *MockSettingsInterface_Expecter) ComplianceSecurityProfile() *MockSettingsInterface_ComplianceSecurityProfile_Call {
	return &MockSettingsInterface_ComplianceSecurityProfile_Call{Call: _e.mock.On("ComplianceSecurityProfile")}
}

func (_c *MockSettingsInterface_ComplianceSecurityProfile_Call) Run(run func()) *MockSettingsInterface_ComplianceSecurityProfile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSettingsInterface_ComplianceSecurityProfile_Call) Return(_a0 settings.ComplianceSecurityProfileInterface) *MockSettingsInterface_ComplianceSecurityProfile_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSettingsInterface_ComplianceSecurityProfile_Call) RunAndReturn(run func() settings.ComplianceSecurityProfileInterface) *MockSettingsInterface_ComplianceSecurityProfile_Call {
	_c.Call.Return(run)
	return _c
}

// DefaultNamespace provides a mock function with given fields:
func (_m *MockSettingsInterface) DefaultNamespace() settings.DefaultNamespaceInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DefaultNamespace")
	}

	var r0 settings.DefaultNamespaceInterface
	if rf, ok := ret.Get(0).(func() settings.DefaultNamespaceInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.DefaultNamespaceInterface)
		}
	}

	return r0
}

// MockSettingsInterface_DefaultNamespace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DefaultNamespace'
type MockSettingsInterface_DefaultNamespace_Call struct {
	*mock.Call
}

// DefaultNamespace is a helper method to define mock.On call
func (_e *MockSettingsInterface_Expecter) DefaultNamespace() *MockSettingsInterface_DefaultNamespace_Call {
	return &MockSettingsInterface_DefaultNamespace_Call{Call: _e.mock.On("DefaultNamespace")}
}

func (_c *MockSettingsInterface_DefaultNamespace_Call) Run(run func()) *MockSettingsInterface_DefaultNamespace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSettingsInterface_DefaultNamespace_Call) Return(_a0 settings.DefaultNamespaceInterface) *MockSettingsInterface_DefaultNamespace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSettingsInterface_DefaultNamespace_Call) RunAndReturn(run func() settings.DefaultNamespaceInterface) *MockSettingsInterface_DefaultNamespace_Call {
	_c.Call.Return(run)
	return _c
}

// EnhancedSecurityMonitoring provides a mock function with given fields:
func (_m *MockSettingsInterface) EnhancedSecurityMonitoring() settings.EnhancedSecurityMonitoringInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EnhancedSecurityMonitoring")
	}

	var r0 settings.EnhancedSecurityMonitoringInterface
	if rf, ok := ret.Get(0).(func() settings.EnhancedSecurityMonitoringInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.EnhancedSecurityMonitoringInterface)
		}
	}

	return r0
}

// MockSettingsInterface_EnhancedSecurityMonitoring_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnhancedSecurityMonitoring'
type MockSettingsInterface_EnhancedSecurityMonitoring_Call struct {
	*mock.Call
}

// EnhancedSecurityMonitoring is a helper method to define mock.On call
func (_e *MockSettingsInterface_Expecter) EnhancedSecurityMonitoring() *MockSettingsInterface_EnhancedSecurityMonitoring_Call {
	return &MockSettingsInterface_EnhancedSecurityMonitoring_Call{Call: _e.mock.On("EnhancedSecurityMonitoring")}
}

func (_c *MockSettingsInterface_EnhancedSecurityMonitoring_Call) Run(run func()) *MockSettingsInterface_EnhancedSecurityMonitoring_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSettingsInterface_EnhancedSecurityMonitoring_Call) Return(_a0 settings.EnhancedSecurityMonitoringInterface) *MockSettingsInterface_EnhancedSecurityMonitoring_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSettingsInterface_EnhancedSecurityMonitoring_Call) RunAndReturn(run func() settings.EnhancedSecurityMonitoringInterface) *MockSettingsInterface_EnhancedSecurityMonitoring_Call {
	_c.Call.Return(run)
	return _c
}

// RestrictWorkspaceAdmins provides a mock function with given fields:
func (_m *MockSettingsInterface) RestrictWorkspaceAdmins() settings.RestrictWorkspaceAdminsInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RestrictWorkspaceAdmins")
	}

	var r0 settings.RestrictWorkspaceAdminsInterface
	if rf, ok := ret.Get(0).(func() settings.RestrictWorkspaceAdminsInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.RestrictWorkspaceAdminsInterface)
		}
	}

	return r0
}

// MockSettingsInterface_RestrictWorkspaceAdmins_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RestrictWorkspaceAdmins'
type MockSettingsInterface_RestrictWorkspaceAdmins_Call struct {
	*mock.Call
}

// RestrictWorkspaceAdmins is a helper method to define mock.On call
func (_e *MockSettingsInterface_Expecter) RestrictWorkspaceAdmins() *MockSettingsInterface_RestrictWorkspaceAdmins_Call {
	return &MockSettingsInterface_RestrictWorkspaceAdmins_Call{Call: _e.mock.On("RestrictWorkspaceAdmins")}
}

func (_c *MockSettingsInterface_RestrictWorkspaceAdmins_Call) Run(run func()) *MockSettingsInterface_RestrictWorkspaceAdmins_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSettingsInterface_RestrictWorkspaceAdmins_Call) Return(_a0 settings.RestrictWorkspaceAdminsInterface) *MockSettingsInterface_RestrictWorkspaceAdmins_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSettingsInterface_RestrictWorkspaceAdmins_Call) RunAndReturn(run func() settings.RestrictWorkspaceAdminsInterface) *MockSettingsInterface_RestrictWorkspaceAdmins_Call {
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
