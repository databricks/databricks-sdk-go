// Code generated by mockery v2.53.2. DO NOT EDIT.

package settings

import (
	context "context"

	settings "github.com/databricks/databricks-sdk-go/service/settings"
	mock "github.com/stretchr/testify/mock"
)

// MockAibiDashboardEmbeddingAccessPolicyInterface is an autogenerated mock type for the AibiDashboardEmbeddingAccessPolicyInterface type
type MockAibiDashboardEmbeddingAccessPolicyInterface struct {
	mock.Mock
}

type MockAibiDashboardEmbeddingAccessPolicyInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAibiDashboardEmbeddingAccessPolicyInterface) EXPECT() *MockAibiDashboardEmbeddingAccessPolicyInterface_Expecter {
	return &MockAibiDashboardEmbeddingAccessPolicyInterface_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockAibiDashboardEmbeddingAccessPolicyInterface) Delete(ctx context.Context, request settings.DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) (*settings.DeleteAibiDashboardEmbeddingAccessPolicySettingResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 *settings.DeleteAibiDashboardEmbeddingAccessPolicySettingResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) (*settings.DeleteAibiDashboardEmbeddingAccessPolicySettingResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) *settings.DeleteAibiDashboardEmbeddingAccessPolicySettingResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.DeleteAibiDashboardEmbeddingAccessPolicySettingResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAibiDashboardEmbeddingAccessPolicyInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockAibiDashboardEmbeddingAccessPolicyInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.DeleteAibiDashboardEmbeddingAccessPolicySettingRequest
func (_e *MockAibiDashboardEmbeddingAccessPolicyInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockAibiDashboardEmbeddingAccessPolicyInterface_Delete_Call {
	return &MockAibiDashboardEmbeddingAccessPolicyInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockAibiDashboardEmbeddingAccessPolicyInterface_Delete_Call) Run(run func(ctx context.Context, request settings.DeleteAibiDashboardEmbeddingAccessPolicySettingRequest)) *MockAibiDashboardEmbeddingAccessPolicyInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.DeleteAibiDashboardEmbeddingAccessPolicySettingRequest))
	})
	return _c
}

func (_c *MockAibiDashboardEmbeddingAccessPolicyInterface_Delete_Call) Return(_a0 *settings.DeleteAibiDashboardEmbeddingAccessPolicySettingResponse, _a1 error) *MockAibiDashboardEmbeddingAccessPolicyInterface_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAibiDashboardEmbeddingAccessPolicyInterface_Delete_Call) RunAndReturn(run func(context.Context, settings.DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) (*settings.DeleteAibiDashboardEmbeddingAccessPolicySettingResponse, error)) *MockAibiDashboardEmbeddingAccessPolicyInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockAibiDashboardEmbeddingAccessPolicyInterface) Get(ctx context.Context, request settings.GetAibiDashboardEmbeddingAccessPolicySettingRequest) (*settings.AibiDashboardEmbeddingAccessPolicySetting, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *settings.AibiDashboardEmbeddingAccessPolicySetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetAibiDashboardEmbeddingAccessPolicySettingRequest) (*settings.AibiDashboardEmbeddingAccessPolicySetting, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetAibiDashboardEmbeddingAccessPolicySettingRequest) *settings.AibiDashboardEmbeddingAccessPolicySetting); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.AibiDashboardEmbeddingAccessPolicySetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.GetAibiDashboardEmbeddingAccessPolicySettingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAibiDashboardEmbeddingAccessPolicyInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockAibiDashboardEmbeddingAccessPolicyInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.GetAibiDashboardEmbeddingAccessPolicySettingRequest
func (_e *MockAibiDashboardEmbeddingAccessPolicyInterface_Expecter) Get(ctx interface{}, request interface{}) *MockAibiDashboardEmbeddingAccessPolicyInterface_Get_Call {
	return &MockAibiDashboardEmbeddingAccessPolicyInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockAibiDashboardEmbeddingAccessPolicyInterface_Get_Call) Run(run func(ctx context.Context, request settings.GetAibiDashboardEmbeddingAccessPolicySettingRequest)) *MockAibiDashboardEmbeddingAccessPolicyInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.GetAibiDashboardEmbeddingAccessPolicySettingRequest))
	})
	return _c
}

func (_c *MockAibiDashboardEmbeddingAccessPolicyInterface_Get_Call) Return(_a0 *settings.AibiDashboardEmbeddingAccessPolicySetting, _a1 error) *MockAibiDashboardEmbeddingAccessPolicyInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAibiDashboardEmbeddingAccessPolicyInterface_Get_Call) RunAndReturn(run func(context.Context, settings.GetAibiDashboardEmbeddingAccessPolicySettingRequest) (*settings.AibiDashboardEmbeddingAccessPolicySetting, error)) *MockAibiDashboardEmbeddingAccessPolicyInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockAibiDashboardEmbeddingAccessPolicyInterface) Update(ctx context.Context, request settings.UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) (*settings.AibiDashboardEmbeddingAccessPolicySetting, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *settings.AibiDashboardEmbeddingAccessPolicySetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) (*settings.AibiDashboardEmbeddingAccessPolicySetting, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) *settings.AibiDashboardEmbeddingAccessPolicySetting); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.AibiDashboardEmbeddingAccessPolicySetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAibiDashboardEmbeddingAccessPolicyInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockAibiDashboardEmbeddingAccessPolicyInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.UpdateAibiDashboardEmbeddingAccessPolicySettingRequest
func (_e *MockAibiDashboardEmbeddingAccessPolicyInterface_Expecter) Update(ctx interface{}, request interface{}) *MockAibiDashboardEmbeddingAccessPolicyInterface_Update_Call {
	return &MockAibiDashboardEmbeddingAccessPolicyInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockAibiDashboardEmbeddingAccessPolicyInterface_Update_Call) Run(run func(ctx context.Context, request settings.UpdateAibiDashboardEmbeddingAccessPolicySettingRequest)) *MockAibiDashboardEmbeddingAccessPolicyInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.UpdateAibiDashboardEmbeddingAccessPolicySettingRequest))
	})
	return _c
}

func (_c *MockAibiDashboardEmbeddingAccessPolicyInterface_Update_Call) Return(_a0 *settings.AibiDashboardEmbeddingAccessPolicySetting, _a1 error) *MockAibiDashboardEmbeddingAccessPolicyInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAibiDashboardEmbeddingAccessPolicyInterface_Update_Call) RunAndReturn(run func(context.Context, settings.UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) (*settings.AibiDashboardEmbeddingAccessPolicySetting, error)) *MockAibiDashboardEmbeddingAccessPolicyInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAibiDashboardEmbeddingAccessPolicyInterface creates a new instance of MockAibiDashboardEmbeddingAccessPolicyInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAibiDashboardEmbeddingAccessPolicyInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAibiDashboardEmbeddingAccessPolicyInterface {
	mock := &MockAibiDashboardEmbeddingAccessPolicyInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
