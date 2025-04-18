// Code generated by mockery v2.53.2. DO NOT EDIT.

package settings

import (
	context "context"

	settings "github.com/databricks/databricks-sdk-go/service/settings"
	mock "github.com/stretchr/testify/mock"
)

// MockAibiDashboardEmbeddingApprovedDomainsInterface is an autogenerated mock type for the AibiDashboardEmbeddingApprovedDomainsInterface type
type MockAibiDashboardEmbeddingApprovedDomainsInterface struct {
	mock.Mock
}

type MockAibiDashboardEmbeddingApprovedDomainsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAibiDashboardEmbeddingApprovedDomainsInterface) EXPECT() *MockAibiDashboardEmbeddingApprovedDomainsInterface_Expecter {
	return &MockAibiDashboardEmbeddingApprovedDomainsInterface_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockAibiDashboardEmbeddingApprovedDomainsInterface) Delete(ctx context.Context, request settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 *settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) *settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAibiDashboardEmbeddingApprovedDomainsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockAibiDashboardEmbeddingApprovedDomainsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest
func (_e *MockAibiDashboardEmbeddingApprovedDomainsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockAibiDashboardEmbeddingApprovedDomainsInterface_Delete_Call {
	return &MockAibiDashboardEmbeddingApprovedDomainsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockAibiDashboardEmbeddingApprovedDomainsInterface_Delete_Call) Run(run func(ctx context.Context, request settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest)) *MockAibiDashboardEmbeddingApprovedDomainsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest))
	})
	return _c
}

func (_c *MockAibiDashboardEmbeddingApprovedDomainsInterface_Delete_Call) Return(_a0 *settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse, _a1 error) *MockAibiDashboardEmbeddingApprovedDomainsInterface_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAibiDashboardEmbeddingApprovedDomainsInterface_Delete_Call) RunAndReturn(run func(context.Context, settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse, error)) *MockAibiDashboardEmbeddingApprovedDomainsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockAibiDashboardEmbeddingApprovedDomainsInterface) Get(ctx context.Context, request settings.GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*settings.AibiDashboardEmbeddingApprovedDomainsSetting, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *settings.AibiDashboardEmbeddingApprovedDomainsSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*settings.AibiDashboardEmbeddingApprovedDomainsSetting, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) *settings.AibiDashboardEmbeddingApprovedDomainsSetting); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.AibiDashboardEmbeddingApprovedDomainsSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAibiDashboardEmbeddingApprovedDomainsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockAibiDashboardEmbeddingApprovedDomainsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.GetAibiDashboardEmbeddingApprovedDomainsSettingRequest
func (_e *MockAibiDashboardEmbeddingApprovedDomainsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockAibiDashboardEmbeddingApprovedDomainsInterface_Get_Call {
	return &MockAibiDashboardEmbeddingApprovedDomainsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockAibiDashboardEmbeddingApprovedDomainsInterface_Get_Call) Run(run func(ctx context.Context, request settings.GetAibiDashboardEmbeddingApprovedDomainsSettingRequest)) *MockAibiDashboardEmbeddingApprovedDomainsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.GetAibiDashboardEmbeddingApprovedDomainsSettingRequest))
	})
	return _c
}

func (_c *MockAibiDashboardEmbeddingApprovedDomainsInterface_Get_Call) Return(_a0 *settings.AibiDashboardEmbeddingApprovedDomainsSetting, _a1 error) *MockAibiDashboardEmbeddingApprovedDomainsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAibiDashboardEmbeddingApprovedDomainsInterface_Get_Call) RunAndReturn(run func(context.Context, settings.GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*settings.AibiDashboardEmbeddingApprovedDomainsSetting, error)) *MockAibiDashboardEmbeddingApprovedDomainsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockAibiDashboardEmbeddingApprovedDomainsInterface) Update(ctx context.Context, request settings.UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*settings.AibiDashboardEmbeddingApprovedDomainsSetting, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *settings.AibiDashboardEmbeddingApprovedDomainsSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*settings.AibiDashboardEmbeddingApprovedDomainsSetting, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) *settings.AibiDashboardEmbeddingApprovedDomainsSetting); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.AibiDashboardEmbeddingApprovedDomainsSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAibiDashboardEmbeddingApprovedDomainsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockAibiDashboardEmbeddingApprovedDomainsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest
func (_e *MockAibiDashboardEmbeddingApprovedDomainsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockAibiDashboardEmbeddingApprovedDomainsInterface_Update_Call {
	return &MockAibiDashboardEmbeddingApprovedDomainsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockAibiDashboardEmbeddingApprovedDomainsInterface_Update_Call) Run(run func(ctx context.Context, request settings.UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest)) *MockAibiDashboardEmbeddingApprovedDomainsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest))
	})
	return _c
}

func (_c *MockAibiDashboardEmbeddingApprovedDomainsInterface_Update_Call) Return(_a0 *settings.AibiDashboardEmbeddingApprovedDomainsSetting, _a1 error) *MockAibiDashboardEmbeddingApprovedDomainsInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAibiDashboardEmbeddingApprovedDomainsInterface_Update_Call) RunAndReturn(run func(context.Context, settings.UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*settings.AibiDashboardEmbeddingApprovedDomainsSetting, error)) *MockAibiDashboardEmbeddingApprovedDomainsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAibiDashboardEmbeddingApprovedDomainsInterface creates a new instance of MockAibiDashboardEmbeddingApprovedDomainsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAibiDashboardEmbeddingApprovedDomainsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAibiDashboardEmbeddingApprovedDomainsInterface {
	mock := &MockAibiDashboardEmbeddingApprovedDomainsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
