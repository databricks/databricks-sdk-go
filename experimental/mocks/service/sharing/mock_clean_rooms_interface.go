// Code generated by mockery v2.39.1. DO NOT EDIT.

package sharing

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	mock "github.com/stretchr/testify/mock"

	sharing "github.com/databricks/databricks-sdk-go/service/sharing"
)

// MockCleanRoomsInterface is an autogenerated mock type for the CleanRoomsInterface type
type MockCleanRoomsInterface struct {
	mock.Mock
}

type MockCleanRoomsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCleanRoomsInterface) EXPECT() *MockCleanRoomsInterface_Expecter {
	return &MockCleanRoomsInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockCleanRoomsInterface) Create(ctx context.Context, request sharing.CreateCleanRoom) (*sharing.CleanRoomInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *sharing.CleanRoomInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sharing.CreateCleanRoom) (*sharing.CleanRoomInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sharing.CreateCleanRoom) *sharing.CleanRoomInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sharing.CleanRoomInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sharing.CreateCleanRoom) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCleanRoomsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockCleanRoomsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.CreateCleanRoom
func (_e *MockCleanRoomsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockCleanRoomsInterface_Create_Call {
	return &MockCleanRoomsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockCleanRoomsInterface_Create_Call) Run(run func(ctx context.Context, request sharing.CreateCleanRoom)) *MockCleanRoomsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.CreateCleanRoom))
	})
	return _c
}

func (_c *MockCleanRoomsInterface_Create_Call) Return(_a0 *sharing.CleanRoomInfo, _a1 error) *MockCleanRoomsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCleanRoomsInterface_Create_Call) RunAndReturn(run func(context.Context, sharing.CreateCleanRoom) (*sharing.CleanRoomInfo, error)) *MockCleanRoomsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockCleanRoomsInterface) Delete(ctx context.Context, request sharing.DeleteCleanRoomRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, sharing.DeleteCleanRoomRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCleanRoomsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockCleanRoomsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.DeleteCleanRoomRequest
func (_e *MockCleanRoomsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockCleanRoomsInterface_Delete_Call {
	return &MockCleanRoomsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockCleanRoomsInterface_Delete_Call) Run(run func(ctx context.Context, request sharing.DeleteCleanRoomRequest)) *MockCleanRoomsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.DeleteCleanRoomRequest))
	})
	return _c
}

func (_c *MockCleanRoomsInterface_Delete_Call) Return(_a0 error) *MockCleanRoomsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCleanRoomsInterface_Delete_Call) RunAndReturn(run func(context.Context, sharing.DeleteCleanRoomRequest) error) *MockCleanRoomsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByNameArg provides a mock function with given fields: ctx, nameArg
func (_m *MockCleanRoomsInterface) DeleteByNameArg(ctx context.Context, nameArg string) error {
	ret := _m.Called(ctx, nameArg)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByNameArg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, nameArg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCleanRoomsInterface_DeleteByNameArg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByNameArg'
type MockCleanRoomsInterface_DeleteByNameArg_Call struct {
	*mock.Call
}

// DeleteByNameArg is a helper method to define mock.On call
//   - ctx context.Context
//   - nameArg string
func (_e *MockCleanRoomsInterface_Expecter) DeleteByNameArg(ctx interface{}, nameArg interface{}) *MockCleanRoomsInterface_DeleteByNameArg_Call {
	return &MockCleanRoomsInterface_DeleteByNameArg_Call{Call: _e.mock.On("DeleteByNameArg", ctx, nameArg)}
}

func (_c *MockCleanRoomsInterface_DeleteByNameArg_Call) Run(run func(ctx context.Context, nameArg string)) *MockCleanRoomsInterface_DeleteByNameArg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockCleanRoomsInterface_DeleteByNameArg_Call) Return(_a0 error) *MockCleanRoomsInterface_DeleteByNameArg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCleanRoomsInterface_DeleteByNameArg_Call) RunAndReturn(run func(context.Context, string) error) *MockCleanRoomsInterface_DeleteByNameArg_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockCleanRoomsInterface) Get(ctx context.Context, request sharing.GetCleanRoomRequest) (*sharing.CleanRoomInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *sharing.CleanRoomInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sharing.GetCleanRoomRequest) (*sharing.CleanRoomInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sharing.GetCleanRoomRequest) *sharing.CleanRoomInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sharing.CleanRoomInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sharing.GetCleanRoomRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCleanRoomsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockCleanRoomsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.GetCleanRoomRequest
func (_e *MockCleanRoomsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockCleanRoomsInterface_Get_Call {
	return &MockCleanRoomsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockCleanRoomsInterface_Get_Call) Run(run func(ctx context.Context, request sharing.GetCleanRoomRequest)) *MockCleanRoomsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.GetCleanRoomRequest))
	})
	return _c
}

func (_c *MockCleanRoomsInterface_Get_Call) Return(_a0 *sharing.CleanRoomInfo, _a1 error) *MockCleanRoomsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCleanRoomsInterface_Get_Call) RunAndReturn(run func(context.Context, sharing.GetCleanRoomRequest) (*sharing.CleanRoomInfo, error)) *MockCleanRoomsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByNameArg provides a mock function with given fields: ctx, nameArg
func (_m *MockCleanRoomsInterface) GetByNameArg(ctx context.Context, nameArg string) (*sharing.CleanRoomInfo, error) {
	ret := _m.Called(ctx, nameArg)

	if len(ret) == 0 {
		panic("no return value specified for GetByNameArg")
	}

	var r0 *sharing.CleanRoomInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*sharing.CleanRoomInfo, error)); ok {
		return rf(ctx, nameArg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *sharing.CleanRoomInfo); ok {
		r0 = rf(ctx, nameArg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sharing.CleanRoomInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, nameArg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCleanRoomsInterface_GetByNameArg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByNameArg'
type MockCleanRoomsInterface_GetByNameArg_Call struct {
	*mock.Call
}

// GetByNameArg is a helper method to define mock.On call
//   - ctx context.Context
//   - nameArg string
func (_e *MockCleanRoomsInterface_Expecter) GetByNameArg(ctx interface{}, nameArg interface{}) *MockCleanRoomsInterface_GetByNameArg_Call {
	return &MockCleanRoomsInterface_GetByNameArg_Call{Call: _e.mock.On("GetByNameArg", ctx, nameArg)}
}

func (_c *MockCleanRoomsInterface_GetByNameArg_Call) Run(run func(ctx context.Context, nameArg string)) *MockCleanRoomsInterface_GetByNameArg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockCleanRoomsInterface_GetByNameArg_Call) Return(_a0 *sharing.CleanRoomInfo, _a1 error) *MockCleanRoomsInterface_GetByNameArg_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCleanRoomsInterface_GetByNameArg_Call) RunAndReturn(run func(context.Context, string) (*sharing.CleanRoomInfo, error)) *MockCleanRoomsInterface_GetByNameArg_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockCleanRoomsInterface) Impl() sharing.CleanRoomsService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 sharing.CleanRoomsService
	if rf, ok := ret.Get(0).(func() sharing.CleanRoomsService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sharing.CleanRoomsService)
		}
	}

	return r0
}

// MockCleanRoomsInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockCleanRoomsInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockCleanRoomsInterface_Expecter) Impl() *MockCleanRoomsInterface_Impl_Call {
	return &MockCleanRoomsInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockCleanRoomsInterface_Impl_Call) Run(run func()) *MockCleanRoomsInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCleanRoomsInterface_Impl_Call) Return(_a0 sharing.CleanRoomsService) *MockCleanRoomsInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCleanRoomsInterface_Impl_Call) RunAndReturn(run func() sharing.CleanRoomsService) *MockCleanRoomsInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockCleanRoomsInterface) List(ctx context.Context, request sharing.ListCleanRoomsRequest) *listing.PaginatingIterator[sharing.ListCleanRoomsRequest, *sharing.ListCleanRoomsResponse, sharing.CleanRoomInfo] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *listing.PaginatingIterator[sharing.ListCleanRoomsRequest, *sharing.ListCleanRoomsResponse, sharing.CleanRoomInfo]
	if rf, ok := ret.Get(0).(func(context.Context, sharing.ListCleanRoomsRequest) *listing.PaginatingIterator[sharing.ListCleanRoomsRequest, *sharing.ListCleanRoomsResponse, sharing.CleanRoomInfo]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*listing.PaginatingIterator[sharing.ListCleanRoomsRequest, *sharing.ListCleanRoomsResponse, sharing.CleanRoomInfo])
		}
	}

	return r0
}

// MockCleanRoomsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockCleanRoomsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.ListCleanRoomsRequest
func (_e *MockCleanRoomsInterface_Expecter) List(ctx interface{}, request interface{}) *MockCleanRoomsInterface_List_Call {
	return &MockCleanRoomsInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockCleanRoomsInterface_List_Call) Run(run func(ctx context.Context, request sharing.ListCleanRoomsRequest)) *MockCleanRoomsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.ListCleanRoomsRequest))
	})
	return _c
}

func (_c *MockCleanRoomsInterface_List_Call) Return(_a0 *listing.PaginatingIterator[sharing.ListCleanRoomsRequest, *sharing.ListCleanRoomsResponse, sharing.CleanRoomInfo]) *MockCleanRoomsInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCleanRoomsInterface_List_Call) RunAndReturn(run func(context.Context, sharing.ListCleanRoomsRequest) *listing.PaginatingIterator[sharing.ListCleanRoomsRequest, *sharing.ListCleanRoomsResponse, sharing.CleanRoomInfo]) *MockCleanRoomsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockCleanRoomsInterface) ListAll(ctx context.Context, request sharing.ListCleanRoomsRequest) ([]sharing.CleanRoomInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []sharing.CleanRoomInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sharing.ListCleanRoomsRequest) ([]sharing.CleanRoomInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sharing.ListCleanRoomsRequest) []sharing.CleanRoomInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sharing.CleanRoomInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sharing.ListCleanRoomsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCleanRoomsInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockCleanRoomsInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.ListCleanRoomsRequest
func (_e *MockCleanRoomsInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockCleanRoomsInterface_ListAll_Call {
	return &MockCleanRoomsInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockCleanRoomsInterface_ListAll_Call) Run(run func(ctx context.Context, request sharing.ListCleanRoomsRequest)) *MockCleanRoomsInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.ListCleanRoomsRequest))
	})
	return _c
}

func (_c *MockCleanRoomsInterface_ListAll_Call) Return(_a0 []sharing.CleanRoomInfo, _a1 error) *MockCleanRoomsInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCleanRoomsInterface_ListAll_Call) RunAndReturn(run func(context.Context, sharing.ListCleanRoomsRequest) ([]sharing.CleanRoomInfo, error)) *MockCleanRoomsInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockCleanRoomsInterface) Update(ctx context.Context, request sharing.UpdateCleanRoom) (*sharing.CleanRoomInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *sharing.CleanRoomInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sharing.UpdateCleanRoom) (*sharing.CleanRoomInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sharing.UpdateCleanRoom) *sharing.CleanRoomInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sharing.CleanRoomInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sharing.UpdateCleanRoom) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCleanRoomsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockCleanRoomsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.UpdateCleanRoom
func (_e *MockCleanRoomsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockCleanRoomsInterface_Update_Call {
	return &MockCleanRoomsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockCleanRoomsInterface_Update_Call) Run(run func(ctx context.Context, request sharing.UpdateCleanRoom)) *MockCleanRoomsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.UpdateCleanRoom))
	})
	return _c
}

func (_c *MockCleanRoomsInterface_Update_Call) Return(_a0 *sharing.CleanRoomInfo, _a1 error) *MockCleanRoomsInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCleanRoomsInterface_Update_Call) RunAndReturn(run func(context.Context, sharing.UpdateCleanRoom) (*sharing.CleanRoomInfo, error)) *MockCleanRoomsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockCleanRoomsInterface) WithImpl(impl sharing.CleanRoomsService) sharing.CleanRoomsInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 sharing.CleanRoomsInterface
	if rf, ok := ret.Get(0).(func(sharing.CleanRoomsService) sharing.CleanRoomsInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sharing.CleanRoomsInterface)
		}
	}

	return r0
}

// MockCleanRoomsInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockCleanRoomsInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl sharing.CleanRoomsService
func (_e *MockCleanRoomsInterface_Expecter) WithImpl(impl interface{}) *MockCleanRoomsInterface_WithImpl_Call {
	return &MockCleanRoomsInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockCleanRoomsInterface_WithImpl_Call) Run(run func(impl sharing.CleanRoomsService)) *MockCleanRoomsInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(sharing.CleanRoomsService))
	})
	return _c
}

func (_c *MockCleanRoomsInterface_WithImpl_Call) Return(_a0 sharing.CleanRoomsInterface) *MockCleanRoomsInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCleanRoomsInterface_WithImpl_Call) RunAndReturn(run func(sharing.CleanRoomsService) sharing.CleanRoomsInterface) *MockCleanRoomsInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCleanRoomsInterface creates a new instance of MockCleanRoomsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCleanRoomsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCleanRoomsInterface {
	mock := &MockCleanRoomsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
