// Code generated by mockery v2.53.2. DO NOT EDIT.

package settings

import (
	context "context"

	settings "github.com/databricks/databricks-sdk-go/service/settings"
	mock "github.com/stretchr/testify/mock"
)

// MockEnableResultsDownloadingInterface is an autogenerated mock type for the EnableResultsDownloadingInterface type
type MockEnableResultsDownloadingInterface struct {
	mock.Mock
}

type MockEnableResultsDownloadingInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEnableResultsDownloadingInterface) EXPECT() *MockEnableResultsDownloadingInterface_Expecter {
	return &MockEnableResultsDownloadingInterface_Expecter{mock: &_m.Mock}
}

// GetEnableResultsDownloading provides a mock function with given fields: ctx
func (_m *MockEnableResultsDownloadingInterface) GetEnableResultsDownloading(ctx context.Context) (*settings.EnableResultsDownloading, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetEnableResultsDownloading")
	}

	var r0 *settings.EnableResultsDownloading
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*settings.EnableResultsDownloading, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *settings.EnableResultsDownloading); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.EnableResultsDownloading)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEnableResultsDownloadingInterface_GetEnableResultsDownloading_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEnableResultsDownloading'
type MockEnableResultsDownloadingInterface_GetEnableResultsDownloading_Call struct {
	*mock.Call
}

// GetEnableResultsDownloading is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockEnableResultsDownloadingInterface_Expecter) GetEnableResultsDownloading(ctx interface{}) *MockEnableResultsDownloadingInterface_GetEnableResultsDownloading_Call {
	return &MockEnableResultsDownloadingInterface_GetEnableResultsDownloading_Call{Call: _e.mock.On("GetEnableResultsDownloading", ctx)}
}

func (_c *MockEnableResultsDownloadingInterface_GetEnableResultsDownloading_Call) Run(run func(ctx context.Context)) *MockEnableResultsDownloadingInterface_GetEnableResultsDownloading_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockEnableResultsDownloadingInterface_GetEnableResultsDownloading_Call) Return(_a0 *settings.EnableResultsDownloading, _a1 error) *MockEnableResultsDownloadingInterface_GetEnableResultsDownloading_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEnableResultsDownloadingInterface_GetEnableResultsDownloading_Call) RunAndReturn(run func(context.Context) (*settings.EnableResultsDownloading, error)) *MockEnableResultsDownloadingInterface_GetEnableResultsDownloading_Call {
	_c.Call.Return(run)
	return _c
}

// PatchEnableResultsDownloading provides a mock function with given fields: ctx, request
func (_m *MockEnableResultsDownloadingInterface) PatchEnableResultsDownloading(ctx context.Context, request settings.UpdateEnableResultsDownloadingRequest) (*settings.EnableResultsDownloading, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for PatchEnableResultsDownloading")
	}

	var r0 *settings.EnableResultsDownloading
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateEnableResultsDownloadingRequest) (*settings.EnableResultsDownloading, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateEnableResultsDownloadingRequest) *settings.EnableResultsDownloading); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.EnableResultsDownloading)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.UpdateEnableResultsDownloadingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEnableResultsDownloadingInterface_PatchEnableResultsDownloading_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PatchEnableResultsDownloading'
type MockEnableResultsDownloadingInterface_PatchEnableResultsDownloading_Call struct {
	*mock.Call
}

// PatchEnableResultsDownloading is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.UpdateEnableResultsDownloadingRequest
func (_e *MockEnableResultsDownloadingInterface_Expecter) PatchEnableResultsDownloading(ctx interface{}, request interface{}) *MockEnableResultsDownloadingInterface_PatchEnableResultsDownloading_Call {
	return &MockEnableResultsDownloadingInterface_PatchEnableResultsDownloading_Call{Call: _e.mock.On("PatchEnableResultsDownloading", ctx, request)}
}

func (_c *MockEnableResultsDownloadingInterface_PatchEnableResultsDownloading_Call) Run(run func(ctx context.Context, request settings.UpdateEnableResultsDownloadingRequest)) *MockEnableResultsDownloadingInterface_PatchEnableResultsDownloading_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.UpdateEnableResultsDownloadingRequest))
	})
	return _c
}

func (_c *MockEnableResultsDownloadingInterface_PatchEnableResultsDownloading_Call) Return(_a0 *settings.EnableResultsDownloading, _a1 error) *MockEnableResultsDownloadingInterface_PatchEnableResultsDownloading_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEnableResultsDownloadingInterface_PatchEnableResultsDownloading_Call) RunAndReturn(run func(context.Context, settings.UpdateEnableResultsDownloadingRequest) (*settings.EnableResultsDownloading, error)) *MockEnableResultsDownloadingInterface_PatchEnableResultsDownloading_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockEnableResultsDownloadingInterface creates a new instance of MockEnableResultsDownloadingInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEnableResultsDownloadingInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEnableResultsDownloadingInterface {
	mock := &MockEnableResultsDownloadingInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
