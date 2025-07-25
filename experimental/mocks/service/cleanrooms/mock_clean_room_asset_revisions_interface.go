// Code generated by mockery v2.53.2. DO NOT EDIT.

package cleanrooms

import (
	context "context"

	cleanrooms "github.com/databricks/databricks-sdk-go/service/cleanrooms"

	listing "github.com/databricks/databricks-sdk-go/listing"

	mock "github.com/stretchr/testify/mock"
)

// MockCleanRoomAssetRevisionsInterface is an autogenerated mock type for the CleanRoomAssetRevisionsInterface type
type MockCleanRoomAssetRevisionsInterface struct {
	mock.Mock
}

type MockCleanRoomAssetRevisionsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCleanRoomAssetRevisionsInterface) EXPECT() *MockCleanRoomAssetRevisionsInterface_Expecter {
	return &MockCleanRoomAssetRevisionsInterface_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockCleanRoomAssetRevisionsInterface) Get(ctx context.Context, request cleanrooms.GetCleanRoomAssetRevisionRequest) (*cleanrooms.CleanRoomAsset, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *cleanrooms.CleanRoomAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cleanrooms.GetCleanRoomAssetRevisionRequest) (*cleanrooms.CleanRoomAsset, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cleanrooms.GetCleanRoomAssetRevisionRequest) *cleanrooms.CleanRoomAsset); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cleanrooms.CleanRoomAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, cleanrooms.GetCleanRoomAssetRevisionRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCleanRoomAssetRevisionsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockCleanRoomAssetRevisionsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request cleanrooms.GetCleanRoomAssetRevisionRequest
func (_e *MockCleanRoomAssetRevisionsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockCleanRoomAssetRevisionsInterface_Get_Call {
	return &MockCleanRoomAssetRevisionsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockCleanRoomAssetRevisionsInterface_Get_Call) Run(run func(ctx context.Context, request cleanrooms.GetCleanRoomAssetRevisionRequest)) *MockCleanRoomAssetRevisionsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(cleanrooms.GetCleanRoomAssetRevisionRequest))
	})
	return _c
}

func (_c *MockCleanRoomAssetRevisionsInterface_Get_Call) Return(_a0 *cleanrooms.CleanRoomAsset, _a1 error) *MockCleanRoomAssetRevisionsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCleanRoomAssetRevisionsInterface_Get_Call) RunAndReturn(run func(context.Context, cleanrooms.GetCleanRoomAssetRevisionRequest) (*cleanrooms.CleanRoomAsset, error)) *MockCleanRoomAssetRevisionsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockCleanRoomAssetRevisionsInterface) List(ctx context.Context, request cleanrooms.ListCleanRoomAssetRevisionsRequest) listing.Iterator[cleanrooms.CleanRoomAsset] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[cleanrooms.CleanRoomAsset]
	if rf, ok := ret.Get(0).(func(context.Context, cleanrooms.ListCleanRoomAssetRevisionsRequest) listing.Iterator[cleanrooms.CleanRoomAsset]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[cleanrooms.CleanRoomAsset])
		}
	}

	return r0
}

// MockCleanRoomAssetRevisionsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockCleanRoomAssetRevisionsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request cleanrooms.ListCleanRoomAssetRevisionsRequest
func (_e *MockCleanRoomAssetRevisionsInterface_Expecter) List(ctx interface{}, request interface{}) *MockCleanRoomAssetRevisionsInterface_List_Call {
	return &MockCleanRoomAssetRevisionsInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockCleanRoomAssetRevisionsInterface_List_Call) Run(run func(ctx context.Context, request cleanrooms.ListCleanRoomAssetRevisionsRequest)) *MockCleanRoomAssetRevisionsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(cleanrooms.ListCleanRoomAssetRevisionsRequest))
	})
	return _c
}

func (_c *MockCleanRoomAssetRevisionsInterface_List_Call) Return(_a0 listing.Iterator[cleanrooms.CleanRoomAsset]) *MockCleanRoomAssetRevisionsInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCleanRoomAssetRevisionsInterface_List_Call) RunAndReturn(run func(context.Context, cleanrooms.ListCleanRoomAssetRevisionsRequest) listing.Iterator[cleanrooms.CleanRoomAsset]) *MockCleanRoomAssetRevisionsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockCleanRoomAssetRevisionsInterface) ListAll(ctx context.Context, request cleanrooms.ListCleanRoomAssetRevisionsRequest) ([]cleanrooms.CleanRoomAsset, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []cleanrooms.CleanRoomAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cleanrooms.ListCleanRoomAssetRevisionsRequest) ([]cleanrooms.CleanRoomAsset, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cleanrooms.ListCleanRoomAssetRevisionsRequest) []cleanrooms.CleanRoomAsset); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cleanrooms.CleanRoomAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, cleanrooms.ListCleanRoomAssetRevisionsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCleanRoomAssetRevisionsInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockCleanRoomAssetRevisionsInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request cleanrooms.ListCleanRoomAssetRevisionsRequest
func (_e *MockCleanRoomAssetRevisionsInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockCleanRoomAssetRevisionsInterface_ListAll_Call {
	return &MockCleanRoomAssetRevisionsInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockCleanRoomAssetRevisionsInterface_ListAll_Call) Run(run func(ctx context.Context, request cleanrooms.ListCleanRoomAssetRevisionsRequest)) *MockCleanRoomAssetRevisionsInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(cleanrooms.ListCleanRoomAssetRevisionsRequest))
	})
	return _c
}

func (_c *MockCleanRoomAssetRevisionsInterface_ListAll_Call) Return(_a0 []cleanrooms.CleanRoomAsset, _a1 error) *MockCleanRoomAssetRevisionsInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCleanRoomAssetRevisionsInterface_ListAll_Call) RunAndReturn(run func(context.Context, cleanrooms.ListCleanRoomAssetRevisionsRequest) ([]cleanrooms.CleanRoomAsset, error)) *MockCleanRoomAssetRevisionsInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCleanRoomAssetRevisionsInterface creates a new instance of MockCleanRoomAssetRevisionsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCleanRoomAssetRevisionsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCleanRoomAssetRevisionsInterface {
	mock := &MockCleanRoomAssetRevisionsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
