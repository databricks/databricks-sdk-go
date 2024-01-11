// Code generated by mockery v2.39.1. DO NOT EDIT.

package billing

import (
	context "context"

	billing "github.com/databricks/databricks-sdk-go/service/billing"

	mock "github.com/stretchr/testify/mock"
)

// MockBillableUsageInterface is an autogenerated mock type for the BillableUsageInterface type
type MockBillableUsageInterface struct {
	mock.Mock
}

type MockBillableUsageInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBillableUsageInterface) EXPECT() *MockBillableUsageInterface_Expecter {
	return &MockBillableUsageInterface_Expecter{mock: &_m.Mock}
}

// Download provides a mock function with given fields: ctx, request
func (_m *MockBillableUsageInterface) Download(ctx context.Context, request billing.DownloadRequest) (*billing.DownloadResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Download")
	}

	var r0 *billing.DownloadResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, billing.DownloadRequest) (*billing.DownloadResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, billing.DownloadRequest) *billing.DownloadResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billing.DownloadResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, billing.DownloadRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBillableUsageInterface_Download_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Download'
type MockBillableUsageInterface_Download_Call struct {
	*mock.Call
}

// Download is a helper method to define mock.On call
//   - ctx context.Context
//   - request billing.DownloadRequest
func (_e *MockBillableUsageInterface_Expecter) Download(ctx interface{}, request interface{}) *MockBillableUsageInterface_Download_Call {
	return &MockBillableUsageInterface_Download_Call{Call: _e.mock.On("Download", ctx, request)}
}

func (_c *MockBillableUsageInterface_Download_Call) Run(run func(ctx context.Context, request billing.DownloadRequest)) *MockBillableUsageInterface_Download_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(billing.DownloadRequest))
	})
	return _c
}

func (_c *MockBillableUsageInterface_Download_Call) Return(_a0 *billing.DownloadResponse, _a1 error) *MockBillableUsageInterface_Download_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBillableUsageInterface_Download_Call) RunAndReturn(run func(context.Context, billing.DownloadRequest) (*billing.DownloadResponse, error)) *MockBillableUsageInterface_Download_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockBillableUsageInterface) Impl() billing.BillableUsageService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 billing.BillableUsageService
	if rf, ok := ret.Get(0).(func() billing.BillableUsageService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(billing.BillableUsageService)
		}
	}

	return r0
}

// MockBillableUsageInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockBillableUsageInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockBillableUsageInterface_Expecter) Impl() *MockBillableUsageInterface_Impl_Call {
	return &MockBillableUsageInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockBillableUsageInterface_Impl_Call) Run(run func()) *MockBillableUsageInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBillableUsageInterface_Impl_Call) Return(_a0 billing.BillableUsageService) *MockBillableUsageInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBillableUsageInterface_Impl_Call) RunAndReturn(run func() billing.BillableUsageService) *MockBillableUsageInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockBillableUsageInterface) WithImpl(impl billing.BillableUsageService) billing.BillableUsageInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 billing.BillableUsageInterface
	if rf, ok := ret.Get(0).(func(billing.BillableUsageService) billing.BillableUsageInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(billing.BillableUsageInterface)
		}
	}

	return r0
}

// MockBillableUsageInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockBillableUsageInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl billing.BillableUsageService
func (_e *MockBillableUsageInterface_Expecter) WithImpl(impl interface{}) *MockBillableUsageInterface_WithImpl_Call {
	return &MockBillableUsageInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockBillableUsageInterface_WithImpl_Call) Run(run func(impl billing.BillableUsageService)) *MockBillableUsageInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(billing.BillableUsageService))
	})
	return _c
}

func (_c *MockBillableUsageInterface_WithImpl_Call) Return(_a0 billing.BillableUsageInterface) *MockBillableUsageInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBillableUsageInterface_WithImpl_Call) RunAndReturn(run func(billing.BillableUsageService) billing.BillableUsageInterface) *MockBillableUsageInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBillableUsageInterface creates a new instance of MockBillableUsageInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBillableUsageInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBillableUsageInterface {
	mock := &MockBillableUsageInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
