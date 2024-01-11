// Code generated by mockery v2.39.1. DO NOT EDIT.

package sharing

import (
	context "context"

	catalog "github.com/databricks/databricks-sdk-go/service/catalog"

	listing "github.com/databricks/databricks-sdk-go/listing"

	mock "github.com/stretchr/testify/mock"

	sharing "github.com/databricks/databricks-sdk-go/service/sharing"
)

// MockSharesInterface is an autogenerated mock type for the SharesInterface type
type MockSharesInterface struct {
	mock.Mock
}

type MockSharesInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSharesInterface) EXPECT() *MockSharesInterface_Expecter {
	return &MockSharesInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockSharesInterface) Create(ctx context.Context, request sharing.CreateShare) (*sharing.ShareInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *sharing.ShareInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sharing.CreateShare) (*sharing.ShareInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sharing.CreateShare) *sharing.ShareInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sharing.ShareInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sharing.CreateShare) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSharesInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockSharesInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.CreateShare
func (_e *MockSharesInterface_Expecter) Create(ctx interface{}, request interface{}) *MockSharesInterface_Create_Call {
	return &MockSharesInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockSharesInterface_Create_Call) Run(run func(ctx context.Context, request sharing.CreateShare)) *MockSharesInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.CreateShare))
	})
	return _c
}

func (_c *MockSharesInterface_Create_Call) Return(_a0 *sharing.ShareInfo, _a1 error) *MockSharesInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSharesInterface_Create_Call) RunAndReturn(run func(context.Context, sharing.CreateShare) (*sharing.ShareInfo, error)) *MockSharesInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockSharesInterface) Delete(ctx context.Context, request sharing.DeleteShareRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, sharing.DeleteShareRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSharesInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockSharesInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.DeleteShareRequest
func (_e *MockSharesInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockSharesInterface_Delete_Call {
	return &MockSharesInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockSharesInterface_Delete_Call) Run(run func(ctx context.Context, request sharing.DeleteShareRequest)) *MockSharesInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.DeleteShareRequest))
	})
	return _c
}

func (_c *MockSharesInterface_Delete_Call) Return(_a0 error) *MockSharesInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSharesInterface_Delete_Call) RunAndReturn(run func(context.Context, sharing.DeleteShareRequest) error) *MockSharesInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByName provides a mock function with given fields: ctx, name
func (_m *MockSharesInterface) DeleteByName(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSharesInterface_DeleteByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByName'
type MockSharesInterface_DeleteByName_Call struct {
	*mock.Call
}

// DeleteByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockSharesInterface_Expecter) DeleteByName(ctx interface{}, name interface{}) *MockSharesInterface_DeleteByName_Call {
	return &MockSharesInterface_DeleteByName_Call{Call: _e.mock.On("DeleteByName", ctx, name)}
}

func (_c *MockSharesInterface_DeleteByName_Call) Run(run func(ctx context.Context, name string)) *MockSharesInterface_DeleteByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSharesInterface_DeleteByName_Call) Return(_a0 error) *MockSharesInterface_DeleteByName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSharesInterface_DeleteByName_Call) RunAndReturn(run func(context.Context, string) error) *MockSharesInterface_DeleteByName_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockSharesInterface) Get(ctx context.Context, request sharing.GetShareRequest) (*sharing.ShareInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *sharing.ShareInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sharing.GetShareRequest) (*sharing.ShareInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sharing.GetShareRequest) *sharing.ShareInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sharing.ShareInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sharing.GetShareRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSharesInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockSharesInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.GetShareRequest
func (_e *MockSharesInterface_Expecter) Get(ctx interface{}, request interface{}) *MockSharesInterface_Get_Call {
	return &MockSharesInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockSharesInterface_Get_Call) Run(run func(ctx context.Context, request sharing.GetShareRequest)) *MockSharesInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.GetShareRequest))
	})
	return _c
}

func (_c *MockSharesInterface_Get_Call) Return(_a0 *sharing.ShareInfo, _a1 error) *MockSharesInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSharesInterface_Get_Call) RunAndReturn(run func(context.Context, sharing.GetShareRequest) (*sharing.ShareInfo, error)) *MockSharesInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *MockSharesInterface) GetByName(ctx context.Context, name string) (*sharing.ShareInfo, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *sharing.ShareInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*sharing.ShareInfo, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *sharing.ShareInfo); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sharing.ShareInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSharesInterface_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type MockSharesInterface_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockSharesInterface_Expecter) GetByName(ctx interface{}, name interface{}) *MockSharesInterface_GetByName_Call {
	return &MockSharesInterface_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *MockSharesInterface_GetByName_Call) Run(run func(ctx context.Context, name string)) *MockSharesInterface_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSharesInterface_GetByName_Call) Return(_a0 *sharing.ShareInfo, _a1 error) *MockSharesInterface_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSharesInterface_GetByName_Call) RunAndReturn(run func(context.Context, string) (*sharing.ShareInfo, error)) *MockSharesInterface_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockSharesInterface) Impl() sharing.SharesService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 sharing.SharesService
	if rf, ok := ret.Get(0).(func() sharing.SharesService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sharing.SharesService)
		}
	}

	return r0
}

// MockSharesInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockSharesInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockSharesInterface_Expecter) Impl() *MockSharesInterface_Impl_Call {
	return &MockSharesInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockSharesInterface_Impl_Call) Run(run func()) *MockSharesInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSharesInterface_Impl_Call) Return(_a0 sharing.SharesService) *MockSharesInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSharesInterface_Impl_Call) RunAndReturn(run func() sharing.SharesService) *MockSharesInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx
func (_m *MockSharesInterface) List(ctx context.Context) *listing.PaginatingIterator[struct{}, *sharing.ListSharesResponse, sharing.ShareInfo] {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *listing.PaginatingIterator[struct{}, *sharing.ListSharesResponse, sharing.ShareInfo]
	if rf, ok := ret.Get(0).(func(context.Context) *listing.PaginatingIterator[struct{}, *sharing.ListSharesResponse, sharing.ShareInfo]); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*listing.PaginatingIterator[struct{}, *sharing.ListSharesResponse, sharing.ShareInfo])
		}
	}

	return r0
}

// MockSharesInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockSharesInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockSharesInterface_Expecter) List(ctx interface{}) *MockSharesInterface_List_Call {
	return &MockSharesInterface_List_Call{Call: _e.mock.On("List", ctx)}
}

func (_c *MockSharesInterface_List_Call) Run(run func(ctx context.Context)) *MockSharesInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSharesInterface_List_Call) Return(_a0 *listing.PaginatingIterator[struct{}, *sharing.ListSharesResponse, sharing.ShareInfo]) *MockSharesInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSharesInterface_List_Call) RunAndReturn(run func(context.Context) *listing.PaginatingIterator[struct{}, *sharing.ListSharesResponse, sharing.ShareInfo]) *MockSharesInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx
func (_m *MockSharesInterface) ListAll(ctx context.Context) ([]sharing.ShareInfo, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []sharing.ShareInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]sharing.ShareInfo, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []sharing.ShareInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sharing.ShareInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSharesInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockSharesInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockSharesInterface_Expecter) ListAll(ctx interface{}) *MockSharesInterface_ListAll_Call {
	return &MockSharesInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx)}
}

func (_c *MockSharesInterface_ListAll_Call) Run(run func(ctx context.Context)) *MockSharesInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSharesInterface_ListAll_Call) Return(_a0 []sharing.ShareInfo, _a1 error) *MockSharesInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSharesInterface_ListAll_Call) RunAndReturn(run func(context.Context) ([]sharing.ShareInfo, error)) *MockSharesInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// SharePermissions provides a mock function with given fields: ctx, request
func (_m *MockSharesInterface) SharePermissions(ctx context.Context, request sharing.SharePermissionsRequest) (*catalog.PermissionsList, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for SharePermissions")
	}

	var r0 *catalog.PermissionsList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sharing.SharePermissionsRequest) (*catalog.PermissionsList, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sharing.SharePermissionsRequest) *catalog.PermissionsList); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.PermissionsList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sharing.SharePermissionsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSharesInterface_SharePermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SharePermissions'
type MockSharesInterface_SharePermissions_Call struct {
	*mock.Call
}

// SharePermissions is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.SharePermissionsRequest
func (_e *MockSharesInterface_Expecter) SharePermissions(ctx interface{}, request interface{}) *MockSharesInterface_SharePermissions_Call {
	return &MockSharesInterface_SharePermissions_Call{Call: _e.mock.On("SharePermissions", ctx, request)}
}

func (_c *MockSharesInterface_SharePermissions_Call) Run(run func(ctx context.Context, request sharing.SharePermissionsRequest)) *MockSharesInterface_SharePermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.SharePermissionsRequest))
	})
	return _c
}

func (_c *MockSharesInterface_SharePermissions_Call) Return(_a0 *catalog.PermissionsList, _a1 error) *MockSharesInterface_SharePermissions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSharesInterface_SharePermissions_Call) RunAndReturn(run func(context.Context, sharing.SharePermissionsRequest) (*catalog.PermissionsList, error)) *MockSharesInterface_SharePermissions_Call {
	_c.Call.Return(run)
	return _c
}

// SharePermissionsByName provides a mock function with given fields: ctx, name
func (_m *MockSharesInterface) SharePermissionsByName(ctx context.Context, name string) (*catalog.PermissionsList, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for SharePermissionsByName")
	}

	var r0 *catalog.PermissionsList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*catalog.PermissionsList, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *catalog.PermissionsList); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.PermissionsList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSharesInterface_SharePermissionsByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SharePermissionsByName'
type MockSharesInterface_SharePermissionsByName_Call struct {
	*mock.Call
}

// SharePermissionsByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockSharesInterface_Expecter) SharePermissionsByName(ctx interface{}, name interface{}) *MockSharesInterface_SharePermissionsByName_Call {
	return &MockSharesInterface_SharePermissionsByName_Call{Call: _e.mock.On("SharePermissionsByName", ctx, name)}
}

func (_c *MockSharesInterface_SharePermissionsByName_Call) Run(run func(ctx context.Context, name string)) *MockSharesInterface_SharePermissionsByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSharesInterface_SharePermissionsByName_Call) Return(_a0 *catalog.PermissionsList, _a1 error) *MockSharesInterface_SharePermissionsByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSharesInterface_SharePermissionsByName_Call) RunAndReturn(run func(context.Context, string) (*catalog.PermissionsList, error)) *MockSharesInterface_SharePermissionsByName_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockSharesInterface) Update(ctx context.Context, request sharing.UpdateShare) (*sharing.ShareInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *sharing.ShareInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sharing.UpdateShare) (*sharing.ShareInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sharing.UpdateShare) *sharing.ShareInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sharing.ShareInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sharing.UpdateShare) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSharesInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockSharesInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.UpdateShare
func (_e *MockSharesInterface_Expecter) Update(ctx interface{}, request interface{}) *MockSharesInterface_Update_Call {
	return &MockSharesInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockSharesInterface_Update_Call) Run(run func(ctx context.Context, request sharing.UpdateShare)) *MockSharesInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.UpdateShare))
	})
	return _c
}

func (_c *MockSharesInterface_Update_Call) Return(_a0 *sharing.ShareInfo, _a1 error) *MockSharesInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSharesInterface_Update_Call) RunAndReturn(run func(context.Context, sharing.UpdateShare) (*sharing.ShareInfo, error)) *MockSharesInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdatePermissions provides a mock function with given fields: ctx, request
func (_m *MockSharesInterface) UpdatePermissions(ctx context.Context, request sharing.UpdateSharePermissions) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePermissions")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, sharing.UpdateSharePermissions) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSharesInterface_UpdatePermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdatePermissions'
type MockSharesInterface_UpdatePermissions_Call struct {
	*mock.Call
}

// UpdatePermissions is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.UpdateSharePermissions
func (_e *MockSharesInterface_Expecter) UpdatePermissions(ctx interface{}, request interface{}) *MockSharesInterface_UpdatePermissions_Call {
	return &MockSharesInterface_UpdatePermissions_Call{Call: _e.mock.On("UpdatePermissions", ctx, request)}
}

func (_c *MockSharesInterface_UpdatePermissions_Call) Run(run func(ctx context.Context, request sharing.UpdateSharePermissions)) *MockSharesInterface_UpdatePermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.UpdateSharePermissions))
	})
	return _c
}

func (_c *MockSharesInterface_UpdatePermissions_Call) Return(_a0 error) *MockSharesInterface_UpdatePermissions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSharesInterface_UpdatePermissions_Call) RunAndReturn(run func(context.Context, sharing.UpdateSharePermissions) error) *MockSharesInterface_UpdatePermissions_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockSharesInterface) WithImpl(impl sharing.SharesService) sharing.SharesInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 sharing.SharesInterface
	if rf, ok := ret.Get(0).(func(sharing.SharesService) sharing.SharesInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sharing.SharesInterface)
		}
	}

	return r0
}

// MockSharesInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockSharesInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl sharing.SharesService
func (_e *MockSharesInterface_Expecter) WithImpl(impl interface{}) *MockSharesInterface_WithImpl_Call {
	return &MockSharesInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockSharesInterface_WithImpl_Call) Run(run func(impl sharing.SharesService)) *MockSharesInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(sharing.SharesService))
	})
	return _c
}

func (_c *MockSharesInterface_WithImpl_Call) Return(_a0 sharing.SharesInterface) *MockSharesInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSharesInterface_WithImpl_Call) RunAndReturn(run func(sharing.SharesService) sharing.SharesInterface) *MockSharesInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSharesInterface creates a new instance of MockSharesInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSharesInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSharesInterface {
	mock := &MockSharesInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
