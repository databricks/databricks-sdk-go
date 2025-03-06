// Code generated by mockery v2.53.0. DO NOT EDIT.

package dashboards

import (
	context "context"

	dashboards "github.com/databricks/databricks-sdk-go/service/dashboards"
	mock "github.com/stretchr/testify/mock"
)

// MockLakeviewEmbeddedInterface is an autogenerated mock type for the LakeviewEmbeddedInterface type
type MockLakeviewEmbeddedInterface struct {
	mock.Mock
}

type MockLakeviewEmbeddedInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLakeviewEmbeddedInterface) EXPECT() *MockLakeviewEmbeddedInterface_Expecter {
	return &MockLakeviewEmbeddedInterface_Expecter{mock: &_m.Mock}
}

// GetPublishedDashboardEmbedded provides a mock function with given fields: ctx, request
func (_m *MockLakeviewEmbeddedInterface) GetPublishedDashboardEmbedded(ctx context.Context, request dashboards.GetPublishedDashboardEmbeddedRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetPublishedDashboardEmbedded")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, dashboards.GetPublishedDashboardEmbeddedRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLakeviewEmbeddedInterface_GetPublishedDashboardEmbedded_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPublishedDashboardEmbedded'
type MockLakeviewEmbeddedInterface_GetPublishedDashboardEmbedded_Call struct {
	*mock.Call
}

// GetPublishedDashboardEmbedded is a helper method to define mock.On call
//   - ctx context.Context
//   - request dashboards.GetPublishedDashboardEmbeddedRequest
func (_e *MockLakeviewEmbeddedInterface_Expecter) GetPublishedDashboardEmbedded(ctx interface{}, request interface{}) *MockLakeviewEmbeddedInterface_GetPublishedDashboardEmbedded_Call {
	return &MockLakeviewEmbeddedInterface_GetPublishedDashboardEmbedded_Call{Call: _e.mock.On("GetPublishedDashboardEmbedded", ctx, request)}
}

func (_c *MockLakeviewEmbeddedInterface_GetPublishedDashboardEmbedded_Call) Run(run func(ctx context.Context, request dashboards.GetPublishedDashboardEmbeddedRequest)) *MockLakeviewEmbeddedInterface_GetPublishedDashboardEmbedded_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dashboards.GetPublishedDashboardEmbeddedRequest))
	})
	return _c
}

func (_c *MockLakeviewEmbeddedInterface_GetPublishedDashboardEmbedded_Call) Return(_a0 error) *MockLakeviewEmbeddedInterface_GetPublishedDashboardEmbedded_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLakeviewEmbeddedInterface_GetPublishedDashboardEmbedded_Call) RunAndReturn(run func(context.Context, dashboards.GetPublishedDashboardEmbeddedRequest) error) *MockLakeviewEmbeddedInterface_GetPublishedDashboardEmbedded_Call {
	_c.Call.Return(run)
	return _c
}

// GetPublishedDashboardEmbeddedByDashboardId provides a mock function with given fields: ctx, dashboardId
func (_m *MockLakeviewEmbeddedInterface) GetPublishedDashboardEmbeddedByDashboardId(ctx context.Context, dashboardId string) error {
	ret := _m.Called(ctx, dashboardId)

	if len(ret) == 0 {
		panic("no return value specified for GetPublishedDashboardEmbeddedByDashboardId")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, dashboardId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLakeviewEmbeddedInterface_GetPublishedDashboardEmbeddedByDashboardId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPublishedDashboardEmbeddedByDashboardId'
type MockLakeviewEmbeddedInterface_GetPublishedDashboardEmbeddedByDashboardId_Call struct {
	*mock.Call
}

// GetPublishedDashboardEmbeddedByDashboardId is a helper method to define mock.On call
//   - ctx context.Context
//   - dashboardId string
func (_e *MockLakeviewEmbeddedInterface_Expecter) GetPublishedDashboardEmbeddedByDashboardId(ctx interface{}, dashboardId interface{}) *MockLakeviewEmbeddedInterface_GetPublishedDashboardEmbeddedByDashboardId_Call {
	return &MockLakeviewEmbeddedInterface_GetPublishedDashboardEmbeddedByDashboardId_Call{Call: _e.mock.On("GetPublishedDashboardEmbeddedByDashboardId", ctx, dashboardId)}
}

func (_c *MockLakeviewEmbeddedInterface_GetPublishedDashboardEmbeddedByDashboardId_Call) Run(run func(ctx context.Context, dashboardId string)) *MockLakeviewEmbeddedInterface_GetPublishedDashboardEmbeddedByDashboardId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockLakeviewEmbeddedInterface_GetPublishedDashboardEmbeddedByDashboardId_Call) Return(_a0 error) *MockLakeviewEmbeddedInterface_GetPublishedDashboardEmbeddedByDashboardId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLakeviewEmbeddedInterface_GetPublishedDashboardEmbeddedByDashboardId_Call) RunAndReturn(run func(context.Context, string) error) *MockLakeviewEmbeddedInterface_GetPublishedDashboardEmbeddedByDashboardId_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLakeviewEmbeddedInterface creates a new instance of MockLakeviewEmbeddedInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLakeviewEmbeddedInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLakeviewEmbeddedInterface {
	mock := &MockLakeviewEmbeddedInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
