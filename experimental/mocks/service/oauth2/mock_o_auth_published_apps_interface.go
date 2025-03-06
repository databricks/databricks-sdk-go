// Code generated by mockery v2.53.0. DO NOT EDIT.

package oauth2

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	mock "github.com/stretchr/testify/mock"

	oauth2 "github.com/databricks/databricks-sdk-go/service/oauth2"
)

// MockOAuthPublishedAppsInterface is an autogenerated mock type for the OAuthPublishedAppsInterface type
type MockOAuthPublishedAppsInterface struct {
	mock.Mock
}

type MockOAuthPublishedAppsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockOAuthPublishedAppsInterface) EXPECT() *MockOAuthPublishedAppsInterface_Expecter {
	return &MockOAuthPublishedAppsInterface_Expecter{mock: &_m.Mock}
}

// List provides a mock function with given fields: ctx, request
func (_m *MockOAuthPublishedAppsInterface) List(ctx context.Context, request oauth2.ListOAuthPublishedAppsRequest) listing.Iterator[oauth2.PublishedAppOutput] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[oauth2.PublishedAppOutput]
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.ListOAuthPublishedAppsRequest) listing.Iterator[oauth2.PublishedAppOutput]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[oauth2.PublishedAppOutput])
		}
	}

	return r0
}

// MockOAuthPublishedAppsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockOAuthPublishedAppsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request oauth2.ListOAuthPublishedAppsRequest
func (_e *MockOAuthPublishedAppsInterface_Expecter) List(ctx interface{}, request interface{}) *MockOAuthPublishedAppsInterface_List_Call {
	return &MockOAuthPublishedAppsInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockOAuthPublishedAppsInterface_List_Call) Run(run func(ctx context.Context, request oauth2.ListOAuthPublishedAppsRequest)) *MockOAuthPublishedAppsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oauth2.ListOAuthPublishedAppsRequest))
	})
	return _c
}

func (_c *MockOAuthPublishedAppsInterface_List_Call) Return(_a0 listing.Iterator[oauth2.PublishedAppOutput]) *MockOAuthPublishedAppsInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOAuthPublishedAppsInterface_List_Call) RunAndReturn(run func(context.Context, oauth2.ListOAuthPublishedAppsRequest) listing.Iterator[oauth2.PublishedAppOutput]) *MockOAuthPublishedAppsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockOAuthPublishedAppsInterface) ListAll(ctx context.Context, request oauth2.ListOAuthPublishedAppsRequest) ([]oauth2.PublishedAppOutput, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []oauth2.PublishedAppOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.ListOAuthPublishedAppsRequest) ([]oauth2.PublishedAppOutput, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.ListOAuthPublishedAppsRequest) []oauth2.PublishedAppOutput); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]oauth2.PublishedAppOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, oauth2.ListOAuthPublishedAppsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOAuthPublishedAppsInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockOAuthPublishedAppsInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request oauth2.ListOAuthPublishedAppsRequest
func (_e *MockOAuthPublishedAppsInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockOAuthPublishedAppsInterface_ListAll_Call {
	return &MockOAuthPublishedAppsInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockOAuthPublishedAppsInterface_ListAll_Call) Run(run func(ctx context.Context, request oauth2.ListOAuthPublishedAppsRequest)) *MockOAuthPublishedAppsInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oauth2.ListOAuthPublishedAppsRequest))
	})
	return _c
}

func (_c *MockOAuthPublishedAppsInterface_ListAll_Call) Return(_a0 []oauth2.PublishedAppOutput, _a1 error) *MockOAuthPublishedAppsInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockOAuthPublishedAppsInterface_ListAll_Call) RunAndReturn(run func(context.Context, oauth2.ListOAuthPublishedAppsRequest) ([]oauth2.PublishedAppOutput, error)) *MockOAuthPublishedAppsInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockOAuthPublishedAppsInterface creates a new instance of MockOAuthPublishedAppsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOAuthPublishedAppsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOAuthPublishedAppsInterface {
	mock := &MockOAuthPublishedAppsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
