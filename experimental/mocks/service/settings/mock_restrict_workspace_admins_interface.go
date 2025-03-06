// Code generated by mockery v2.53.0. DO NOT EDIT.

package settings

import (
	context "context"

	settings "github.com/databricks/databricks-sdk-go/service/settings"
	mock "github.com/stretchr/testify/mock"
)

// MockRestrictWorkspaceAdminsInterface is an autogenerated mock type for the RestrictWorkspaceAdminsInterface type
type MockRestrictWorkspaceAdminsInterface struct {
	mock.Mock
}

type MockRestrictWorkspaceAdminsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRestrictWorkspaceAdminsInterface) EXPECT() *MockRestrictWorkspaceAdminsInterface_Expecter {
	return &MockRestrictWorkspaceAdminsInterface_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockRestrictWorkspaceAdminsInterface) Delete(ctx context.Context, request settings.DeleteRestrictWorkspaceAdminsSettingRequest) (*settings.DeleteRestrictWorkspaceAdminsSettingResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
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

// MockRestrictWorkspaceAdminsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockRestrictWorkspaceAdminsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.DeleteRestrictWorkspaceAdminsSettingRequest
func (_e *MockRestrictWorkspaceAdminsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockRestrictWorkspaceAdminsInterface_Delete_Call {
	return &MockRestrictWorkspaceAdminsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockRestrictWorkspaceAdminsInterface_Delete_Call) Run(run func(ctx context.Context, request settings.DeleteRestrictWorkspaceAdminsSettingRequest)) *MockRestrictWorkspaceAdminsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.DeleteRestrictWorkspaceAdminsSettingRequest))
	})
	return _c
}

func (_c *MockRestrictWorkspaceAdminsInterface_Delete_Call) Return(_a0 *settings.DeleteRestrictWorkspaceAdminsSettingResponse, _a1 error) *MockRestrictWorkspaceAdminsInterface_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRestrictWorkspaceAdminsInterface_Delete_Call) RunAndReturn(run func(context.Context, settings.DeleteRestrictWorkspaceAdminsSettingRequest) (*settings.DeleteRestrictWorkspaceAdminsSettingResponse, error)) *MockRestrictWorkspaceAdminsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockRestrictWorkspaceAdminsInterface) Get(ctx context.Context, request settings.GetRestrictWorkspaceAdminsSettingRequest) (*settings.RestrictWorkspaceAdminsSetting, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
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

// MockRestrictWorkspaceAdminsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockRestrictWorkspaceAdminsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.GetRestrictWorkspaceAdminsSettingRequest
func (_e *MockRestrictWorkspaceAdminsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockRestrictWorkspaceAdminsInterface_Get_Call {
	return &MockRestrictWorkspaceAdminsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockRestrictWorkspaceAdminsInterface_Get_Call) Run(run func(ctx context.Context, request settings.GetRestrictWorkspaceAdminsSettingRequest)) *MockRestrictWorkspaceAdminsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.GetRestrictWorkspaceAdminsSettingRequest))
	})
	return _c
}

func (_c *MockRestrictWorkspaceAdminsInterface_Get_Call) Return(_a0 *settings.RestrictWorkspaceAdminsSetting, _a1 error) *MockRestrictWorkspaceAdminsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRestrictWorkspaceAdminsInterface_Get_Call) RunAndReturn(run func(context.Context, settings.GetRestrictWorkspaceAdminsSettingRequest) (*settings.RestrictWorkspaceAdminsSetting, error)) *MockRestrictWorkspaceAdminsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockRestrictWorkspaceAdminsInterface) Update(ctx context.Context, request settings.UpdateRestrictWorkspaceAdminsSettingRequest) (*settings.RestrictWorkspaceAdminsSetting, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
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

// MockRestrictWorkspaceAdminsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockRestrictWorkspaceAdminsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.UpdateRestrictWorkspaceAdminsSettingRequest
func (_e *MockRestrictWorkspaceAdminsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockRestrictWorkspaceAdminsInterface_Update_Call {
	return &MockRestrictWorkspaceAdminsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockRestrictWorkspaceAdminsInterface_Update_Call) Run(run func(ctx context.Context, request settings.UpdateRestrictWorkspaceAdminsSettingRequest)) *MockRestrictWorkspaceAdminsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.UpdateRestrictWorkspaceAdminsSettingRequest))
	})
	return _c
}

func (_c *MockRestrictWorkspaceAdminsInterface_Update_Call) Return(_a0 *settings.RestrictWorkspaceAdminsSetting, _a1 error) *MockRestrictWorkspaceAdminsInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRestrictWorkspaceAdminsInterface_Update_Call) RunAndReturn(run func(context.Context, settings.UpdateRestrictWorkspaceAdminsSettingRequest) (*settings.RestrictWorkspaceAdminsSetting, error)) *MockRestrictWorkspaceAdminsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRestrictWorkspaceAdminsInterface creates a new instance of MockRestrictWorkspaceAdminsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRestrictWorkspaceAdminsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRestrictWorkspaceAdminsInterface {
	mock := &MockRestrictWorkspaceAdminsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
