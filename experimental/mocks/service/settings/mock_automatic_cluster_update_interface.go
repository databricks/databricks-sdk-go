// Code generated by mockery v2.43.0. DO NOT EDIT.

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
func (_m *MockAutomaticClusterUpdateInterface) Get(ctx context.Context, request settings.GetAutomaticClusterUpdateSettingRequest) (*settings.AutomaticClusterUpdateSetting, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *settings.AutomaticClusterUpdateSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetAutomaticClusterUpdateSettingRequest) (*settings.AutomaticClusterUpdateSetting, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetAutomaticClusterUpdateSettingRequest) *settings.AutomaticClusterUpdateSetting); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.AutomaticClusterUpdateSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.GetAutomaticClusterUpdateSettingRequest) error); ok {
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
//   - request settings.GetAutomaticClusterUpdateSettingRequest
func (_e *MockAutomaticClusterUpdateInterface_Expecter) Get(ctx interface{}, request interface{}) *MockAutomaticClusterUpdateInterface_Get_Call {
	return &MockAutomaticClusterUpdateInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockAutomaticClusterUpdateInterface_Get_Call) Run(run func(ctx context.Context, request settings.GetAutomaticClusterUpdateSettingRequest)) *MockAutomaticClusterUpdateInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.GetAutomaticClusterUpdateSettingRequest))
	})
	return _c
}

func (_c *MockAutomaticClusterUpdateInterface_Get_Call) Return(_a0 *settings.AutomaticClusterUpdateSetting, _a1 error) *MockAutomaticClusterUpdateInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAutomaticClusterUpdateInterface_Get_Call) RunAndReturn(run func(context.Context, settings.GetAutomaticClusterUpdateSettingRequest) (*settings.AutomaticClusterUpdateSetting, error)) *MockAutomaticClusterUpdateInterface_Get_Call {
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
