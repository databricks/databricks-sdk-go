// Code generated by mockery v2.53.2. DO NOT EDIT.

package sql

import (
	context "context"

	sql "github.com/databricks/databricks-sdk-go/service/sql"
	mock "github.com/stretchr/testify/mock"
)

// MockQueryHistoryInterface is an autogenerated mock type for the QueryHistoryInterface type
type MockQueryHistoryInterface struct {
	mock.Mock
}

type MockQueryHistoryInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockQueryHistoryInterface) EXPECT() *MockQueryHistoryInterface_Expecter {
	return &MockQueryHistoryInterface_Expecter{mock: &_m.Mock}
}

// List provides a mock function with given fields: ctx, request
func (_m *MockQueryHistoryInterface) List(ctx context.Context, request sql.ListQueryHistoryRequest) (*sql.ListQueriesResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *sql.ListQueriesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.ListQueryHistoryRequest) (*sql.ListQueriesResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sql.ListQueryHistoryRequest) *sql.ListQueriesResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.ListQueriesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sql.ListQueryHistoryRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryHistoryInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockQueryHistoryInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.ListQueryHistoryRequest
func (_e *MockQueryHistoryInterface_Expecter) List(ctx interface{}, request interface{}) *MockQueryHistoryInterface_List_Call {
	return &MockQueryHistoryInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockQueryHistoryInterface_List_Call) Run(run func(ctx context.Context, request sql.ListQueryHistoryRequest)) *MockQueryHistoryInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.ListQueryHistoryRequest))
	})
	return _c
}

func (_c *MockQueryHistoryInterface_List_Call) Return(_a0 *sql.ListQueriesResponse, _a1 error) *MockQueryHistoryInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQueryHistoryInterface_List_Call) RunAndReturn(run func(context.Context, sql.ListQueryHistoryRequest) (*sql.ListQueriesResponse, error)) *MockQueryHistoryInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockQueryHistoryInterface creates a new instance of MockQueryHistoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockQueryHistoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockQueryHistoryInterface {
	mock := &MockQueryHistoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
