// Code generated by mockery v2.43.0. DO NOT EDIT.

package settings

import (
	context "context"

	settings "github.com/databricks/databricks-sdk-go/service/settings"
	mock "github.com/stretchr/testify/mock"
)

// MockEnableIpAccessListsInterface is an autogenerated mock type for the EnableIpAccessListsInterface type
type MockEnableIpAccessListsInterface struct {
	mock.Mock
}

type MockEnableIpAccessListsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEnableIpAccessListsInterface) EXPECT() *MockEnableIpAccessListsInterface_Expecter {
	return &MockEnableIpAccessListsInterface_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockEnableIpAccessListsInterface) Delete(ctx context.Context, request settings.DeleteAccountIpAccessEnableRequest) (*settings.DeleteAccountIpAccessEnableResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 *settings.DeleteAccountIpAccessEnableResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.DeleteAccountIpAccessEnableRequest) (*settings.DeleteAccountIpAccessEnableResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.DeleteAccountIpAccessEnableRequest) *settings.DeleteAccountIpAccessEnableResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.DeleteAccountIpAccessEnableResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.DeleteAccountIpAccessEnableRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEnableIpAccessListsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockEnableIpAccessListsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.DeleteAccountIpAccessEnableRequest
func (_e *MockEnableIpAccessListsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockEnableIpAccessListsInterface_Delete_Call {
	return &MockEnableIpAccessListsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockEnableIpAccessListsInterface_Delete_Call) Run(run func(ctx context.Context, request settings.DeleteAccountIpAccessEnableRequest)) *MockEnableIpAccessListsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.DeleteAccountIpAccessEnableRequest))
	})
	return _c
}

func (_c *MockEnableIpAccessListsInterface_Delete_Call) Return(_a0 *settings.DeleteAccountIpAccessEnableResponse, _a1 error) *MockEnableIpAccessListsInterface_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEnableIpAccessListsInterface_Delete_Call) RunAndReturn(run func(context.Context, settings.DeleteAccountIpAccessEnableRequest) (*settings.DeleteAccountIpAccessEnableResponse, error)) *MockEnableIpAccessListsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockEnableIpAccessListsInterface) Get(ctx context.Context, request settings.GetAccountIpAccessEnableRequest) (*settings.AccountIpAccessEnable, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *settings.AccountIpAccessEnable
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetAccountIpAccessEnableRequest) (*settings.AccountIpAccessEnable, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetAccountIpAccessEnableRequest) *settings.AccountIpAccessEnable); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.AccountIpAccessEnable)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.GetAccountIpAccessEnableRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEnableIpAccessListsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockEnableIpAccessListsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.GetAccountIpAccessEnableRequest
func (_e *MockEnableIpAccessListsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockEnableIpAccessListsInterface_Get_Call {
	return &MockEnableIpAccessListsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockEnableIpAccessListsInterface_Get_Call) Run(run func(ctx context.Context, request settings.GetAccountIpAccessEnableRequest)) *MockEnableIpAccessListsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.GetAccountIpAccessEnableRequest))
	})
	return _c
}

func (_c *MockEnableIpAccessListsInterface_Get_Call) Return(_a0 *settings.AccountIpAccessEnable, _a1 error) *MockEnableIpAccessListsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEnableIpAccessListsInterface_Get_Call) RunAndReturn(run func(context.Context, settings.GetAccountIpAccessEnableRequest) (*settings.AccountIpAccessEnable, error)) *MockEnableIpAccessListsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockEnableIpAccessListsInterface) Update(ctx context.Context, request settings.UpdateAccountIpAccessEnableRequest) (*settings.AccountIpAccessEnable, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *settings.AccountIpAccessEnable
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateAccountIpAccessEnableRequest) (*settings.AccountIpAccessEnable, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateAccountIpAccessEnableRequest) *settings.AccountIpAccessEnable); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.AccountIpAccessEnable)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.UpdateAccountIpAccessEnableRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEnableIpAccessListsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockEnableIpAccessListsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.UpdateAccountIpAccessEnableRequest
func (_e *MockEnableIpAccessListsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockEnableIpAccessListsInterface_Update_Call {
	return &MockEnableIpAccessListsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockEnableIpAccessListsInterface_Update_Call) Run(run func(ctx context.Context, request settings.UpdateAccountIpAccessEnableRequest)) *MockEnableIpAccessListsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.UpdateAccountIpAccessEnableRequest))
	})
	return _c
}

func (_c *MockEnableIpAccessListsInterface_Update_Call) Return(_a0 *settings.AccountIpAccessEnable, _a1 error) *MockEnableIpAccessListsInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEnableIpAccessListsInterface_Update_Call) RunAndReturn(run func(context.Context, settings.UpdateAccountIpAccessEnableRequest) (*settings.AccountIpAccessEnable, error)) *MockEnableIpAccessListsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockEnableIpAccessListsInterface creates a new instance of MockEnableIpAccessListsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEnableIpAccessListsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEnableIpAccessListsInterface {
	mock := &MockEnableIpAccessListsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
