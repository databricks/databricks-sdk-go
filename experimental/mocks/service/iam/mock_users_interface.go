// Code generated by mockery v2.39.1. DO NOT EDIT.

package iam

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	iam "github.com/databricks/databricks-sdk-go/service/iam"

	mock "github.com/stretchr/testify/mock"
)

// MockUsersInterface is an autogenerated mock type for the UsersInterface type
type MockUsersInterface struct {
	mock.Mock
}

type MockUsersInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUsersInterface) EXPECT() *MockUsersInterface_Expecter {
	return &MockUsersInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockUsersInterface) Create(ctx context.Context, request iam.User) (*iam.User, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *iam.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.User) (*iam.User, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.User) *iam.User); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.User) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUsersInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockUsersInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.User
func (_e *MockUsersInterface_Expecter) Create(ctx interface{}, request interface{}) *MockUsersInterface_Create_Call {
	return &MockUsersInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockUsersInterface_Create_Call) Run(run func(ctx context.Context, request iam.User)) *MockUsersInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.User))
	})
	return _c
}

func (_c *MockUsersInterface_Create_Call) Return(_a0 *iam.User, _a1 error) *MockUsersInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUsersInterface_Create_Call) RunAndReturn(run func(context.Context, iam.User) (*iam.User, error)) *MockUsersInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockUsersInterface) Delete(ctx context.Context, request iam.DeleteUserRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.DeleteUserRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUsersInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockUsersInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.DeleteUserRequest
func (_e *MockUsersInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockUsersInterface_Delete_Call {
	return &MockUsersInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockUsersInterface_Delete_Call) Run(run func(ctx context.Context, request iam.DeleteUserRequest)) *MockUsersInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.DeleteUserRequest))
	})
	return _c
}

func (_c *MockUsersInterface_Delete_Call) Return(_a0 error) *MockUsersInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUsersInterface_Delete_Call) RunAndReturn(run func(context.Context, iam.DeleteUserRequest) error) *MockUsersInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *MockUsersInterface) DeleteById(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteById")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUsersInterface_DeleteById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteById'
type MockUsersInterface_DeleteById_Call struct {
	*mock.Call
}

// DeleteById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockUsersInterface_Expecter) DeleteById(ctx interface{}, id interface{}) *MockUsersInterface_DeleteById_Call {
	return &MockUsersInterface_DeleteById_Call{Call: _e.mock.On("DeleteById", ctx, id)}
}

func (_c *MockUsersInterface_DeleteById_Call) Run(run func(ctx context.Context, id string)) *MockUsersInterface_DeleteById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockUsersInterface_DeleteById_Call) Return(_a0 error) *MockUsersInterface_DeleteById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUsersInterface_DeleteById_Call) RunAndReturn(run func(context.Context, string) error) *MockUsersInterface_DeleteById_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockUsersInterface) Get(ctx context.Context, request iam.GetUserRequest) (*iam.User, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *iam.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.GetUserRequest) (*iam.User, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.GetUserRequest) *iam.User); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.GetUserRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUsersInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockUsersInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.GetUserRequest
func (_e *MockUsersInterface_Expecter) Get(ctx interface{}, request interface{}) *MockUsersInterface_Get_Call {
	return &MockUsersInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockUsersInterface_Get_Call) Run(run func(ctx context.Context, request iam.GetUserRequest)) *MockUsersInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.GetUserRequest))
	})
	return _c
}

func (_c *MockUsersInterface_Get_Call) Return(_a0 *iam.User, _a1 error) *MockUsersInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUsersInterface_Get_Call) RunAndReturn(run func(context.Context, iam.GetUserRequest) (*iam.User, error)) *MockUsersInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: ctx, id
func (_m *MockUsersInterface) GetById(ctx context.Context, id string) (*iam.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *iam.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*iam.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *iam.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUsersInterface_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type MockUsersInterface_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockUsersInterface_Expecter) GetById(ctx interface{}, id interface{}) *MockUsersInterface_GetById_Call {
	return &MockUsersInterface_GetById_Call{Call: _e.mock.On("GetById", ctx, id)}
}

func (_c *MockUsersInterface_GetById_Call) Run(run func(ctx context.Context, id string)) *MockUsersInterface_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockUsersInterface_GetById_Call) Return(_a0 *iam.User, _a1 error) *MockUsersInterface_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUsersInterface_GetById_Call) RunAndReturn(run func(context.Context, string) (*iam.User, error)) *MockUsersInterface_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// GetByUserName provides a mock function with given fields: ctx, name
func (_m *MockUsersInterface) GetByUserName(ctx context.Context, name string) (*iam.User, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByUserName")
	}

	var r0 *iam.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*iam.User, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *iam.User); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUsersInterface_GetByUserName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByUserName'
type MockUsersInterface_GetByUserName_Call struct {
	*mock.Call
}

// GetByUserName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockUsersInterface_Expecter) GetByUserName(ctx interface{}, name interface{}) *MockUsersInterface_GetByUserName_Call {
	return &MockUsersInterface_GetByUserName_Call{Call: _e.mock.On("GetByUserName", ctx, name)}
}

func (_c *MockUsersInterface_GetByUserName_Call) Run(run func(ctx context.Context, name string)) *MockUsersInterface_GetByUserName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockUsersInterface_GetByUserName_Call) Return(_a0 *iam.User, _a1 error) *MockUsersInterface_GetByUserName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUsersInterface_GetByUserName_Call) RunAndReturn(run func(context.Context, string) (*iam.User, error)) *MockUsersInterface_GetByUserName_Call {
	_c.Call.Return(run)
	return _c
}

// GetPermissionLevels provides a mock function with given fields: ctx
func (_m *MockUsersInterface) GetPermissionLevels(ctx context.Context) (*iam.GetPasswordPermissionLevelsResponse, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetPermissionLevels")
	}

	var r0 *iam.GetPasswordPermissionLevelsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*iam.GetPasswordPermissionLevelsResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *iam.GetPasswordPermissionLevelsResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.GetPasswordPermissionLevelsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUsersInterface_GetPermissionLevels_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPermissionLevels'
type MockUsersInterface_GetPermissionLevels_Call struct {
	*mock.Call
}

// GetPermissionLevels is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockUsersInterface_Expecter) GetPermissionLevels(ctx interface{}) *MockUsersInterface_GetPermissionLevels_Call {
	return &MockUsersInterface_GetPermissionLevels_Call{Call: _e.mock.On("GetPermissionLevels", ctx)}
}

func (_c *MockUsersInterface_GetPermissionLevels_Call) Run(run func(ctx context.Context)) *MockUsersInterface_GetPermissionLevels_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockUsersInterface_GetPermissionLevels_Call) Return(_a0 *iam.GetPasswordPermissionLevelsResponse, _a1 error) *MockUsersInterface_GetPermissionLevels_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUsersInterface_GetPermissionLevels_Call) RunAndReturn(run func(context.Context) (*iam.GetPasswordPermissionLevelsResponse, error)) *MockUsersInterface_GetPermissionLevels_Call {
	_c.Call.Return(run)
	return _c
}

// GetPermissions provides a mock function with given fields: ctx
func (_m *MockUsersInterface) GetPermissions(ctx context.Context) (*iam.PasswordPermissions, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetPermissions")
	}

	var r0 *iam.PasswordPermissions
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*iam.PasswordPermissions, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *iam.PasswordPermissions); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.PasswordPermissions)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUsersInterface_GetPermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPermissions'
type MockUsersInterface_GetPermissions_Call struct {
	*mock.Call
}

// GetPermissions is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockUsersInterface_Expecter) GetPermissions(ctx interface{}) *MockUsersInterface_GetPermissions_Call {
	return &MockUsersInterface_GetPermissions_Call{Call: _e.mock.On("GetPermissions", ctx)}
}

func (_c *MockUsersInterface_GetPermissions_Call) Run(run func(ctx context.Context)) *MockUsersInterface_GetPermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockUsersInterface_GetPermissions_Call) Return(_a0 *iam.PasswordPermissions, _a1 error) *MockUsersInterface_GetPermissions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUsersInterface_GetPermissions_Call) RunAndReturn(run func(context.Context) (*iam.PasswordPermissions, error)) *MockUsersInterface_GetPermissions_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockUsersInterface) Impl() iam.UsersService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 iam.UsersService
	if rf, ok := ret.Get(0).(func() iam.UsersService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(iam.UsersService)
		}
	}

	return r0
}

// MockUsersInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockUsersInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockUsersInterface_Expecter) Impl() *MockUsersInterface_Impl_Call {
	return &MockUsersInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockUsersInterface_Impl_Call) Run(run func()) *MockUsersInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUsersInterface_Impl_Call) Return(_a0 iam.UsersService) *MockUsersInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUsersInterface_Impl_Call) RunAndReturn(run func() iam.UsersService) *MockUsersInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockUsersInterface) List(ctx context.Context, request iam.ListUsersRequest) *listing.DeduplicatingIterator[iam.User, string] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *listing.DeduplicatingIterator[iam.User, string]
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListUsersRequest) *listing.DeduplicatingIterator[iam.User, string]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*listing.DeduplicatingIterator[iam.User, string])
		}
	}

	return r0
}

// MockUsersInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockUsersInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.ListUsersRequest
func (_e *MockUsersInterface_Expecter) List(ctx interface{}, request interface{}) *MockUsersInterface_List_Call {
	return &MockUsersInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockUsersInterface_List_Call) Run(run func(ctx context.Context, request iam.ListUsersRequest)) *MockUsersInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.ListUsersRequest))
	})
	return _c
}

func (_c *MockUsersInterface_List_Call) Return(_a0 *listing.DeduplicatingIterator[iam.User, string]) *MockUsersInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUsersInterface_List_Call) RunAndReturn(run func(context.Context, iam.ListUsersRequest) *listing.DeduplicatingIterator[iam.User, string]) *MockUsersInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockUsersInterface) ListAll(ctx context.Context, request iam.ListUsersRequest) ([]iam.User, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []iam.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListUsersRequest) ([]iam.User, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListUsersRequest) []iam.User); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]iam.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.ListUsersRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUsersInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockUsersInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.ListUsersRequest
func (_e *MockUsersInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockUsersInterface_ListAll_Call {
	return &MockUsersInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockUsersInterface_ListAll_Call) Run(run func(ctx context.Context, request iam.ListUsersRequest)) *MockUsersInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.ListUsersRequest))
	})
	return _c
}

func (_c *MockUsersInterface_ListAll_Call) Return(_a0 []iam.User, _a1 error) *MockUsersInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUsersInterface_ListAll_Call) RunAndReturn(run func(context.Context, iam.ListUsersRequest) ([]iam.User, error)) *MockUsersInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, request
func (_m *MockUsersInterface) Patch(ctx context.Context, request iam.PartialUpdate) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Patch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.PartialUpdate) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUsersInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type MockUsersInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.PartialUpdate
func (_e *MockUsersInterface_Expecter) Patch(ctx interface{}, request interface{}) *MockUsersInterface_Patch_Call {
	return &MockUsersInterface_Patch_Call{Call: _e.mock.On("Patch", ctx, request)}
}

func (_c *MockUsersInterface_Patch_Call) Run(run func(ctx context.Context, request iam.PartialUpdate)) *MockUsersInterface_Patch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.PartialUpdate))
	})
	return _c
}

func (_c *MockUsersInterface_Patch_Call) Return(_a0 error) *MockUsersInterface_Patch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUsersInterface_Patch_Call) RunAndReturn(run func(context.Context, iam.PartialUpdate) error) *MockUsersInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// SetPermissions provides a mock function with given fields: ctx, request
func (_m *MockUsersInterface) SetPermissions(ctx context.Context, request iam.PasswordPermissionsRequest) (*iam.PasswordPermissions, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for SetPermissions")
	}

	var r0 *iam.PasswordPermissions
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.PasswordPermissionsRequest) (*iam.PasswordPermissions, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.PasswordPermissionsRequest) *iam.PasswordPermissions); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.PasswordPermissions)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.PasswordPermissionsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUsersInterface_SetPermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPermissions'
type MockUsersInterface_SetPermissions_Call struct {
	*mock.Call
}

// SetPermissions is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.PasswordPermissionsRequest
func (_e *MockUsersInterface_Expecter) SetPermissions(ctx interface{}, request interface{}) *MockUsersInterface_SetPermissions_Call {
	return &MockUsersInterface_SetPermissions_Call{Call: _e.mock.On("SetPermissions", ctx, request)}
}

func (_c *MockUsersInterface_SetPermissions_Call) Run(run func(ctx context.Context, request iam.PasswordPermissionsRequest)) *MockUsersInterface_SetPermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.PasswordPermissionsRequest))
	})
	return _c
}

func (_c *MockUsersInterface_SetPermissions_Call) Return(_a0 *iam.PasswordPermissions, _a1 error) *MockUsersInterface_SetPermissions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUsersInterface_SetPermissions_Call) RunAndReturn(run func(context.Context, iam.PasswordPermissionsRequest) (*iam.PasswordPermissions, error)) *MockUsersInterface_SetPermissions_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockUsersInterface) Update(ctx context.Context, request iam.User) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.User) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUsersInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockUsersInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.User
func (_e *MockUsersInterface_Expecter) Update(ctx interface{}, request interface{}) *MockUsersInterface_Update_Call {
	return &MockUsersInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockUsersInterface_Update_Call) Run(run func(ctx context.Context, request iam.User)) *MockUsersInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.User))
	})
	return _c
}

func (_c *MockUsersInterface_Update_Call) Return(_a0 error) *MockUsersInterface_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUsersInterface_Update_Call) RunAndReturn(run func(context.Context, iam.User) error) *MockUsersInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdatePermissions provides a mock function with given fields: ctx, request
func (_m *MockUsersInterface) UpdatePermissions(ctx context.Context, request iam.PasswordPermissionsRequest) (*iam.PasswordPermissions, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePermissions")
	}

	var r0 *iam.PasswordPermissions
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.PasswordPermissionsRequest) (*iam.PasswordPermissions, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.PasswordPermissionsRequest) *iam.PasswordPermissions); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.PasswordPermissions)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.PasswordPermissionsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUsersInterface_UpdatePermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdatePermissions'
type MockUsersInterface_UpdatePermissions_Call struct {
	*mock.Call
}

// UpdatePermissions is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.PasswordPermissionsRequest
func (_e *MockUsersInterface_Expecter) UpdatePermissions(ctx interface{}, request interface{}) *MockUsersInterface_UpdatePermissions_Call {
	return &MockUsersInterface_UpdatePermissions_Call{Call: _e.mock.On("UpdatePermissions", ctx, request)}
}

func (_c *MockUsersInterface_UpdatePermissions_Call) Run(run func(ctx context.Context, request iam.PasswordPermissionsRequest)) *MockUsersInterface_UpdatePermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.PasswordPermissionsRequest))
	})
	return _c
}

func (_c *MockUsersInterface_UpdatePermissions_Call) Return(_a0 *iam.PasswordPermissions, _a1 error) *MockUsersInterface_UpdatePermissions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUsersInterface_UpdatePermissions_Call) RunAndReturn(run func(context.Context, iam.PasswordPermissionsRequest) (*iam.PasswordPermissions, error)) *MockUsersInterface_UpdatePermissions_Call {
	_c.Call.Return(run)
	return _c
}

// UserUserNameToIdMap provides a mock function with given fields: ctx, request
func (_m *MockUsersInterface) UserUserNameToIdMap(ctx context.Context, request iam.ListUsersRequest) (map[string]string, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for UserUserNameToIdMap")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListUsersRequest) (map[string]string, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListUsersRequest) map[string]string); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.ListUsersRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUsersInterface_UserUserNameToIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserUserNameToIdMap'
type MockUsersInterface_UserUserNameToIdMap_Call struct {
	*mock.Call
}

// UserUserNameToIdMap is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.ListUsersRequest
func (_e *MockUsersInterface_Expecter) UserUserNameToIdMap(ctx interface{}, request interface{}) *MockUsersInterface_UserUserNameToIdMap_Call {
	return &MockUsersInterface_UserUserNameToIdMap_Call{Call: _e.mock.On("UserUserNameToIdMap", ctx, request)}
}

func (_c *MockUsersInterface_UserUserNameToIdMap_Call) Run(run func(ctx context.Context, request iam.ListUsersRequest)) *MockUsersInterface_UserUserNameToIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.ListUsersRequest))
	})
	return _c
}

func (_c *MockUsersInterface_UserUserNameToIdMap_Call) Return(_a0 map[string]string, _a1 error) *MockUsersInterface_UserUserNameToIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUsersInterface_UserUserNameToIdMap_Call) RunAndReturn(run func(context.Context, iam.ListUsersRequest) (map[string]string, error)) *MockUsersInterface_UserUserNameToIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockUsersInterface) WithImpl(impl iam.UsersService) iam.UsersInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 iam.UsersInterface
	if rf, ok := ret.Get(0).(func(iam.UsersService) iam.UsersInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(iam.UsersInterface)
		}
	}

	return r0
}

// MockUsersInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockUsersInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl iam.UsersService
func (_e *MockUsersInterface_Expecter) WithImpl(impl interface{}) *MockUsersInterface_WithImpl_Call {
	return &MockUsersInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockUsersInterface_WithImpl_Call) Run(run func(impl iam.UsersService)) *MockUsersInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(iam.UsersService))
	})
	return _c
}

func (_c *MockUsersInterface_WithImpl_Call) Return(_a0 iam.UsersInterface) *MockUsersInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUsersInterface_WithImpl_Call) RunAndReturn(run func(iam.UsersService) iam.UsersInterface) *MockUsersInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUsersInterface creates a new instance of MockUsersInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUsersInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUsersInterface {
	mock := &MockUsersInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
