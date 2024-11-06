// Code generated by mockery v2.43.0. DO NOT EDIT.

package sharing

import (
	context "context"

	sharing "github.com/databricks/databricks-sdk-go/service/sharing"
	mock "github.com/stretchr/testify/mock"
)

// MockRecipientActivationInterface is an autogenerated mock type for the RecipientActivationInterface type
type MockRecipientActivationInterface struct {
	mock.Mock
}

type MockRecipientActivationInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRecipientActivationInterface) EXPECT() *MockRecipientActivationInterface_Expecter {
	return &MockRecipientActivationInterface_Expecter{mock: &_m.Mock}
}

// GetActivationUrlInfo provides a mock function with given fields: ctx, request
func (_m *MockRecipientActivationInterface) GetActivationUrlInfo(ctx context.Context, request sharing.GetActivationUrlInfoRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetActivationUrlInfo")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, sharing.GetActivationUrlInfoRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRecipientActivationInterface_GetActivationUrlInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetActivationUrlInfo'
type MockRecipientActivationInterface_GetActivationUrlInfo_Call struct {
	*mock.Call
}

// GetActivationUrlInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.GetActivationUrlInfoRequest
func (_e *MockRecipientActivationInterface_Expecter) GetActivationUrlInfo(ctx interface{}, request interface{}) *MockRecipientActivationInterface_GetActivationUrlInfo_Call {
	return &MockRecipientActivationInterface_GetActivationUrlInfo_Call{Call: _e.mock.On("GetActivationUrlInfo", ctx, request)}
}

func (_c *MockRecipientActivationInterface_GetActivationUrlInfo_Call) Run(run func(ctx context.Context, request sharing.GetActivationUrlInfoRequest)) *MockRecipientActivationInterface_GetActivationUrlInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.GetActivationUrlInfoRequest))
	})
	return _c
}

func (_c *MockRecipientActivationInterface_GetActivationUrlInfo_Call) Return(_a0 error) *MockRecipientActivationInterface_GetActivationUrlInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRecipientActivationInterface_GetActivationUrlInfo_Call) RunAndReturn(run func(context.Context, sharing.GetActivationUrlInfoRequest) error) *MockRecipientActivationInterface_GetActivationUrlInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetActivationUrlInfoByActivationUrl provides a mock function with given fields: ctx, activationUrl
func (_m *MockRecipientActivationInterface) GetActivationUrlInfoByActivationUrl(ctx context.Context, activationUrl string) error {
	ret := _m.Called(ctx, activationUrl)

	if len(ret) == 0 {
		panic("no return value specified for GetActivationUrlInfoByActivationUrl")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, activationUrl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRecipientActivationInterface_GetActivationUrlInfoByActivationUrl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetActivationUrlInfoByActivationUrl'
type MockRecipientActivationInterface_GetActivationUrlInfoByActivationUrl_Call struct {
	*mock.Call
}

// GetActivationUrlInfoByActivationUrl is a helper method to define mock.On call
//   - ctx context.Context
//   - activationUrl string
func (_e *MockRecipientActivationInterface_Expecter) GetActivationUrlInfoByActivationUrl(ctx interface{}, activationUrl interface{}) *MockRecipientActivationInterface_GetActivationUrlInfoByActivationUrl_Call {
	return &MockRecipientActivationInterface_GetActivationUrlInfoByActivationUrl_Call{Call: _e.mock.On("GetActivationUrlInfoByActivationUrl", ctx, activationUrl)}
}

func (_c *MockRecipientActivationInterface_GetActivationUrlInfoByActivationUrl_Call) Run(run func(ctx context.Context, activationUrl string)) *MockRecipientActivationInterface_GetActivationUrlInfoByActivationUrl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRecipientActivationInterface_GetActivationUrlInfoByActivationUrl_Call) Return(_a0 error) *MockRecipientActivationInterface_GetActivationUrlInfoByActivationUrl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRecipientActivationInterface_GetActivationUrlInfoByActivationUrl_Call) RunAndReturn(run func(context.Context, string) error) *MockRecipientActivationInterface_GetActivationUrlInfoByActivationUrl_Call {
	_c.Call.Return(run)
	return _c
}

// RetrieveToken provides a mock function with given fields: ctx, request
func (_m *MockRecipientActivationInterface) RetrieveToken(ctx context.Context, request sharing.RetrieveTokenRequest) (*sharing.RetrieveTokenResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveToken")
	}

	var r0 *sharing.RetrieveTokenResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sharing.RetrieveTokenRequest) (*sharing.RetrieveTokenResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sharing.RetrieveTokenRequest) *sharing.RetrieveTokenResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sharing.RetrieveTokenResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sharing.RetrieveTokenRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRecipientActivationInterface_RetrieveToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RetrieveToken'
type MockRecipientActivationInterface_RetrieveToken_Call struct {
	*mock.Call
}

// RetrieveToken is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.RetrieveTokenRequest
func (_e *MockRecipientActivationInterface_Expecter) RetrieveToken(ctx interface{}, request interface{}) *MockRecipientActivationInterface_RetrieveToken_Call {
	return &MockRecipientActivationInterface_RetrieveToken_Call{Call: _e.mock.On("RetrieveToken", ctx, request)}
}

func (_c *MockRecipientActivationInterface_RetrieveToken_Call) Run(run func(ctx context.Context, request sharing.RetrieveTokenRequest)) *MockRecipientActivationInterface_RetrieveToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.RetrieveTokenRequest))
	})
	return _c
}

func (_c *MockRecipientActivationInterface_RetrieveToken_Call) Return(_a0 *sharing.RetrieveTokenResponse, _a1 error) *MockRecipientActivationInterface_RetrieveToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRecipientActivationInterface_RetrieveToken_Call) RunAndReturn(run func(context.Context, sharing.RetrieveTokenRequest) (*sharing.RetrieveTokenResponse, error)) *MockRecipientActivationInterface_RetrieveToken_Call {
	_c.Call.Return(run)
	return _c
}

// RetrieveTokenByActivationUrl provides a mock function with given fields: ctx, activationUrl
func (_m *MockRecipientActivationInterface) RetrieveTokenByActivationUrl(ctx context.Context, activationUrl string) (*sharing.RetrieveTokenResponse, error) {
	ret := _m.Called(ctx, activationUrl)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveTokenByActivationUrl")
	}

	var r0 *sharing.RetrieveTokenResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*sharing.RetrieveTokenResponse, error)); ok {
		return rf(ctx, activationUrl)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *sharing.RetrieveTokenResponse); ok {
		r0 = rf(ctx, activationUrl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sharing.RetrieveTokenResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, activationUrl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRecipientActivationInterface_RetrieveTokenByActivationUrl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RetrieveTokenByActivationUrl'
type MockRecipientActivationInterface_RetrieveTokenByActivationUrl_Call struct {
	*mock.Call
}

// RetrieveTokenByActivationUrl is a helper method to define mock.On call
//   - ctx context.Context
//   - activationUrl string
func (_e *MockRecipientActivationInterface_Expecter) RetrieveTokenByActivationUrl(ctx interface{}, activationUrl interface{}) *MockRecipientActivationInterface_RetrieveTokenByActivationUrl_Call {
	return &MockRecipientActivationInterface_RetrieveTokenByActivationUrl_Call{Call: _e.mock.On("RetrieveTokenByActivationUrl", ctx, activationUrl)}
}

func (_c *MockRecipientActivationInterface_RetrieveTokenByActivationUrl_Call) Run(run func(ctx context.Context, activationUrl string)) *MockRecipientActivationInterface_RetrieveTokenByActivationUrl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRecipientActivationInterface_RetrieveTokenByActivationUrl_Call) Return(_a0 *sharing.RetrieveTokenResponse, _a1 error) *MockRecipientActivationInterface_RetrieveTokenByActivationUrl_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRecipientActivationInterface_RetrieveTokenByActivationUrl_Call) RunAndReturn(run func(context.Context, string) (*sharing.RetrieveTokenResponse, error)) *MockRecipientActivationInterface_RetrieveTokenByActivationUrl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRecipientActivationInterface creates a new instance of MockRecipientActivationInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRecipientActivationInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRecipientActivationInterface {
	mock := &MockRecipientActivationInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
