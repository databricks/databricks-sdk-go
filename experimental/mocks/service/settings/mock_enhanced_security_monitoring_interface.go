// Code generated by mockery v2.53.2. DO NOT EDIT.

package settings

import (
	context "context"

	settings "github.com/databricks/databricks-sdk-go/service/settings"
	mock "github.com/stretchr/testify/mock"
)

// MockEnhancedSecurityMonitoringInterface is an autogenerated mock type for the EnhancedSecurityMonitoringInterface type
type MockEnhancedSecurityMonitoringInterface struct {
	mock.Mock
}

type MockEnhancedSecurityMonitoringInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEnhancedSecurityMonitoringInterface) EXPECT() *MockEnhancedSecurityMonitoringInterface_Expecter {
	return &MockEnhancedSecurityMonitoringInterface_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockEnhancedSecurityMonitoringInterface) Get(ctx context.Context, request settings.GetEnhancedSecurityMonitoringSettingRequest) (*settings.EnhancedSecurityMonitoringSetting, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *settings.EnhancedSecurityMonitoringSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetEnhancedSecurityMonitoringSettingRequest) (*settings.EnhancedSecurityMonitoringSetting, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetEnhancedSecurityMonitoringSettingRequest) *settings.EnhancedSecurityMonitoringSetting); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.EnhancedSecurityMonitoringSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.GetEnhancedSecurityMonitoringSettingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEnhancedSecurityMonitoringInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockEnhancedSecurityMonitoringInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.GetEnhancedSecurityMonitoringSettingRequest
func (_e *MockEnhancedSecurityMonitoringInterface_Expecter) Get(ctx interface{}, request interface{}) *MockEnhancedSecurityMonitoringInterface_Get_Call {
	return &MockEnhancedSecurityMonitoringInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockEnhancedSecurityMonitoringInterface_Get_Call) Run(run func(ctx context.Context, request settings.GetEnhancedSecurityMonitoringSettingRequest)) *MockEnhancedSecurityMonitoringInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.GetEnhancedSecurityMonitoringSettingRequest))
	})
	return _c
}

func (_c *MockEnhancedSecurityMonitoringInterface_Get_Call) Return(_a0 *settings.EnhancedSecurityMonitoringSetting, _a1 error) *MockEnhancedSecurityMonitoringInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEnhancedSecurityMonitoringInterface_Get_Call) RunAndReturn(run func(context.Context, settings.GetEnhancedSecurityMonitoringSettingRequest) (*settings.EnhancedSecurityMonitoringSetting, error)) *MockEnhancedSecurityMonitoringInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockEnhancedSecurityMonitoringInterface) Update(ctx context.Context, request settings.UpdateEnhancedSecurityMonitoringSettingRequest) (*settings.EnhancedSecurityMonitoringSetting, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *settings.EnhancedSecurityMonitoringSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateEnhancedSecurityMonitoringSettingRequest) (*settings.EnhancedSecurityMonitoringSetting, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateEnhancedSecurityMonitoringSettingRequest) *settings.EnhancedSecurityMonitoringSetting); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.EnhancedSecurityMonitoringSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.UpdateEnhancedSecurityMonitoringSettingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEnhancedSecurityMonitoringInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockEnhancedSecurityMonitoringInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.UpdateEnhancedSecurityMonitoringSettingRequest
func (_e *MockEnhancedSecurityMonitoringInterface_Expecter) Update(ctx interface{}, request interface{}) *MockEnhancedSecurityMonitoringInterface_Update_Call {
	return &MockEnhancedSecurityMonitoringInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockEnhancedSecurityMonitoringInterface_Update_Call) Run(run func(ctx context.Context, request settings.UpdateEnhancedSecurityMonitoringSettingRequest)) *MockEnhancedSecurityMonitoringInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.UpdateEnhancedSecurityMonitoringSettingRequest))
	})
	return _c
}

func (_c *MockEnhancedSecurityMonitoringInterface_Update_Call) Return(_a0 *settings.EnhancedSecurityMonitoringSetting, _a1 error) *MockEnhancedSecurityMonitoringInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEnhancedSecurityMonitoringInterface_Update_Call) RunAndReturn(run func(context.Context, settings.UpdateEnhancedSecurityMonitoringSettingRequest) (*settings.EnhancedSecurityMonitoringSetting, error)) *MockEnhancedSecurityMonitoringInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockEnhancedSecurityMonitoringInterface creates a new instance of MockEnhancedSecurityMonitoringInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEnhancedSecurityMonitoringInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEnhancedSecurityMonitoringInterface {
	mock := &MockEnhancedSecurityMonitoringInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
