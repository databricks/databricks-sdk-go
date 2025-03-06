// Code generated by mockery v2.53.0. DO NOT EDIT.

package settings

import (
	context "context"

	settings "github.com/databricks/databricks-sdk-go/service/settings"
	mock "github.com/stretchr/testify/mock"
)

// MockDisableLegacyDbfsInterface is an autogenerated mock type for the DisableLegacyDbfsInterface type
type MockDisableLegacyDbfsInterface struct {
	mock.Mock
}

type MockDisableLegacyDbfsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDisableLegacyDbfsInterface) EXPECT() *MockDisableLegacyDbfsInterface_Expecter {
	return &MockDisableLegacyDbfsInterface_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockDisableLegacyDbfsInterface) Delete(ctx context.Context, request settings.DeleteDisableLegacyDbfsRequest) (*settings.DeleteDisableLegacyDbfsResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 *settings.DeleteDisableLegacyDbfsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.DeleteDisableLegacyDbfsRequest) (*settings.DeleteDisableLegacyDbfsResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.DeleteDisableLegacyDbfsRequest) *settings.DeleteDisableLegacyDbfsResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.DeleteDisableLegacyDbfsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.DeleteDisableLegacyDbfsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDisableLegacyDbfsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockDisableLegacyDbfsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.DeleteDisableLegacyDbfsRequest
func (_e *MockDisableLegacyDbfsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockDisableLegacyDbfsInterface_Delete_Call {
	return &MockDisableLegacyDbfsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockDisableLegacyDbfsInterface_Delete_Call) Run(run func(ctx context.Context, request settings.DeleteDisableLegacyDbfsRequest)) *MockDisableLegacyDbfsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.DeleteDisableLegacyDbfsRequest))
	})
	return _c
}

func (_c *MockDisableLegacyDbfsInterface_Delete_Call) Return(_a0 *settings.DeleteDisableLegacyDbfsResponse, _a1 error) *MockDisableLegacyDbfsInterface_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDisableLegacyDbfsInterface_Delete_Call) RunAndReturn(run func(context.Context, settings.DeleteDisableLegacyDbfsRequest) (*settings.DeleteDisableLegacyDbfsResponse, error)) *MockDisableLegacyDbfsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockDisableLegacyDbfsInterface) Get(ctx context.Context, request settings.GetDisableLegacyDbfsRequest) (*settings.DisableLegacyDbfs, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *settings.DisableLegacyDbfs
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetDisableLegacyDbfsRequest) (*settings.DisableLegacyDbfs, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetDisableLegacyDbfsRequest) *settings.DisableLegacyDbfs); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.DisableLegacyDbfs)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.GetDisableLegacyDbfsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDisableLegacyDbfsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockDisableLegacyDbfsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.GetDisableLegacyDbfsRequest
func (_e *MockDisableLegacyDbfsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockDisableLegacyDbfsInterface_Get_Call {
	return &MockDisableLegacyDbfsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockDisableLegacyDbfsInterface_Get_Call) Run(run func(ctx context.Context, request settings.GetDisableLegacyDbfsRequest)) *MockDisableLegacyDbfsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.GetDisableLegacyDbfsRequest))
	})
	return _c
}

func (_c *MockDisableLegacyDbfsInterface_Get_Call) Return(_a0 *settings.DisableLegacyDbfs, _a1 error) *MockDisableLegacyDbfsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDisableLegacyDbfsInterface_Get_Call) RunAndReturn(run func(context.Context, settings.GetDisableLegacyDbfsRequest) (*settings.DisableLegacyDbfs, error)) *MockDisableLegacyDbfsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockDisableLegacyDbfsInterface) Update(ctx context.Context, request settings.UpdateDisableLegacyDbfsRequest) (*settings.DisableLegacyDbfs, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *settings.DisableLegacyDbfs
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateDisableLegacyDbfsRequest) (*settings.DisableLegacyDbfs, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateDisableLegacyDbfsRequest) *settings.DisableLegacyDbfs); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.DisableLegacyDbfs)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.UpdateDisableLegacyDbfsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDisableLegacyDbfsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockDisableLegacyDbfsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.UpdateDisableLegacyDbfsRequest
func (_e *MockDisableLegacyDbfsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockDisableLegacyDbfsInterface_Update_Call {
	return &MockDisableLegacyDbfsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockDisableLegacyDbfsInterface_Update_Call) Run(run func(ctx context.Context, request settings.UpdateDisableLegacyDbfsRequest)) *MockDisableLegacyDbfsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.UpdateDisableLegacyDbfsRequest))
	})
	return _c
}

func (_c *MockDisableLegacyDbfsInterface_Update_Call) Return(_a0 *settings.DisableLegacyDbfs, _a1 error) *MockDisableLegacyDbfsInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDisableLegacyDbfsInterface_Update_Call) RunAndReturn(run func(context.Context, settings.UpdateDisableLegacyDbfsRequest) (*settings.DisableLegacyDbfs, error)) *MockDisableLegacyDbfsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDisableLegacyDbfsInterface creates a new instance of MockDisableLegacyDbfsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDisableLegacyDbfsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDisableLegacyDbfsInterface {
	mock := &MockDisableLegacyDbfsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
