// Code generated by mockery v2.39.1. DO NOT EDIT.

package iam

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	iam "github.com/databricks/databricks-sdk-go/service/iam"

	mock "github.com/stretchr/testify/mock"
)

// MockAccountUsersInterface is an autogenerated mock type for the AccountUsersInterface type
type MockAccountUsersInterface struct {
	mock.Mock
}

type MockAccountUsersInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAccountUsersInterface) EXPECT() *MockAccountUsersInterface_Expecter {
	return &MockAccountUsersInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockAccountUsersInterface) Create(ctx context.Context, request iam.User) (*iam.User, error) {
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

// MockAccountUsersInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockAccountUsersInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.User
func (_e *MockAccountUsersInterface_Expecter) Create(ctx interface{}, request interface{}) *MockAccountUsersInterface_Create_Call {
	return &MockAccountUsersInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockAccountUsersInterface_Create_Call) Run(run func(ctx context.Context, request iam.User)) *MockAccountUsersInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.User))
	})
	return _c
}

func (_c *MockAccountUsersInterface_Create_Call) Return(_a0 *iam.User, _a1 error) *MockAccountUsersInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountUsersInterface_Create_Call) RunAndReturn(run func(context.Context, iam.User) (*iam.User, error)) *MockAccountUsersInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockAccountUsersInterface) Delete(ctx context.Context, request iam.DeleteAccountUserRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.DeleteAccountUserRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAccountUsersInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockAccountUsersInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.DeleteAccountUserRequest
func (_e *MockAccountUsersInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockAccountUsersInterface_Delete_Call {
	return &MockAccountUsersInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockAccountUsersInterface_Delete_Call) Run(run func(ctx context.Context, request iam.DeleteAccountUserRequest)) *MockAccountUsersInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.DeleteAccountUserRequest))
	})
	return _c
}

func (_c *MockAccountUsersInterface_Delete_Call) Return(_a0 error) *MockAccountUsersInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountUsersInterface_Delete_Call) RunAndReturn(run func(context.Context, iam.DeleteAccountUserRequest) error) *MockAccountUsersInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *MockAccountUsersInterface) DeleteById(ctx context.Context, id string) error {
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

// MockAccountUsersInterface_DeleteById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteById'
type MockAccountUsersInterface_DeleteById_Call struct {
	*mock.Call
}

// DeleteById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockAccountUsersInterface_Expecter) DeleteById(ctx interface{}, id interface{}) *MockAccountUsersInterface_DeleteById_Call {
	return &MockAccountUsersInterface_DeleteById_Call{Call: _e.mock.On("DeleteById", ctx, id)}
}

func (_c *MockAccountUsersInterface_DeleteById_Call) Run(run func(ctx context.Context, id string)) *MockAccountUsersInterface_DeleteById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAccountUsersInterface_DeleteById_Call) Return(_a0 error) *MockAccountUsersInterface_DeleteById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountUsersInterface_DeleteById_Call) RunAndReturn(run func(context.Context, string) error) *MockAccountUsersInterface_DeleteById_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockAccountUsersInterface) Get(ctx context.Context, request iam.GetAccountUserRequest) (*iam.User, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *iam.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.GetAccountUserRequest) (*iam.User, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.GetAccountUserRequest) *iam.User); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.GetAccountUserRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountUsersInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockAccountUsersInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.GetAccountUserRequest
func (_e *MockAccountUsersInterface_Expecter) Get(ctx interface{}, request interface{}) *MockAccountUsersInterface_Get_Call {
	return &MockAccountUsersInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockAccountUsersInterface_Get_Call) Run(run func(ctx context.Context, request iam.GetAccountUserRequest)) *MockAccountUsersInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.GetAccountUserRequest))
	})
	return _c
}

func (_c *MockAccountUsersInterface_Get_Call) Return(_a0 *iam.User, _a1 error) *MockAccountUsersInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountUsersInterface_Get_Call) RunAndReturn(run func(context.Context, iam.GetAccountUserRequest) (*iam.User, error)) *MockAccountUsersInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: ctx, id
func (_m *MockAccountUsersInterface) GetById(ctx context.Context, id string) (*iam.User, error) {
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

// MockAccountUsersInterface_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type MockAccountUsersInterface_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockAccountUsersInterface_Expecter) GetById(ctx interface{}, id interface{}) *MockAccountUsersInterface_GetById_Call {
	return &MockAccountUsersInterface_GetById_Call{Call: _e.mock.On("GetById", ctx, id)}
}

func (_c *MockAccountUsersInterface_GetById_Call) Run(run func(ctx context.Context, id string)) *MockAccountUsersInterface_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAccountUsersInterface_GetById_Call) Return(_a0 *iam.User, _a1 error) *MockAccountUsersInterface_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountUsersInterface_GetById_Call) RunAndReturn(run func(context.Context, string) (*iam.User, error)) *MockAccountUsersInterface_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// GetByUserName provides a mock function with given fields: ctx, name
func (_m *MockAccountUsersInterface) GetByUserName(ctx context.Context, name string) (*iam.User, error) {
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

// MockAccountUsersInterface_GetByUserName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByUserName'
type MockAccountUsersInterface_GetByUserName_Call struct {
	*mock.Call
}

// GetByUserName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockAccountUsersInterface_Expecter) GetByUserName(ctx interface{}, name interface{}) *MockAccountUsersInterface_GetByUserName_Call {
	return &MockAccountUsersInterface_GetByUserName_Call{Call: _e.mock.On("GetByUserName", ctx, name)}
}

func (_c *MockAccountUsersInterface_GetByUserName_Call) Run(run func(ctx context.Context, name string)) *MockAccountUsersInterface_GetByUserName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAccountUsersInterface_GetByUserName_Call) Return(_a0 *iam.User, _a1 error) *MockAccountUsersInterface_GetByUserName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountUsersInterface_GetByUserName_Call) RunAndReturn(run func(context.Context, string) (*iam.User, error)) *MockAccountUsersInterface_GetByUserName_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockAccountUsersInterface) Impl() iam.AccountUsersService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 iam.AccountUsersService
	if rf, ok := ret.Get(0).(func() iam.AccountUsersService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(iam.AccountUsersService)
		}
	}

	return r0
}

// MockAccountUsersInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockAccountUsersInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockAccountUsersInterface_Expecter) Impl() *MockAccountUsersInterface_Impl_Call {
	return &MockAccountUsersInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockAccountUsersInterface_Impl_Call) Run(run func()) *MockAccountUsersInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAccountUsersInterface_Impl_Call) Return(_a0 iam.AccountUsersService) *MockAccountUsersInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountUsersInterface_Impl_Call) RunAndReturn(run func() iam.AccountUsersService) *MockAccountUsersInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockAccountUsersInterface) List(ctx context.Context, request iam.ListAccountUsersRequest) *listing.DeduplicatingIterator[iam.User, string] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *listing.DeduplicatingIterator[iam.User, string]
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListAccountUsersRequest) *listing.DeduplicatingIterator[iam.User, string]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*listing.DeduplicatingIterator[iam.User, string])
		}
	}

	return r0
}

// MockAccountUsersInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockAccountUsersInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.ListAccountUsersRequest
func (_e *MockAccountUsersInterface_Expecter) List(ctx interface{}, request interface{}) *MockAccountUsersInterface_List_Call {
	return &MockAccountUsersInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockAccountUsersInterface_List_Call) Run(run func(ctx context.Context, request iam.ListAccountUsersRequest)) *MockAccountUsersInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.ListAccountUsersRequest))
	})
	return _c
}

func (_c *MockAccountUsersInterface_List_Call) Return(_a0 *listing.DeduplicatingIterator[iam.User, string]) *MockAccountUsersInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountUsersInterface_List_Call) RunAndReturn(run func(context.Context, iam.ListAccountUsersRequest) *listing.DeduplicatingIterator[iam.User, string]) *MockAccountUsersInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockAccountUsersInterface) ListAll(ctx context.Context, request iam.ListAccountUsersRequest) ([]iam.User, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []iam.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListAccountUsersRequest) ([]iam.User, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListAccountUsersRequest) []iam.User); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]iam.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.ListAccountUsersRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountUsersInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockAccountUsersInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.ListAccountUsersRequest
func (_e *MockAccountUsersInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockAccountUsersInterface_ListAll_Call {
	return &MockAccountUsersInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockAccountUsersInterface_ListAll_Call) Run(run func(ctx context.Context, request iam.ListAccountUsersRequest)) *MockAccountUsersInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.ListAccountUsersRequest))
	})
	return _c
}

func (_c *MockAccountUsersInterface_ListAll_Call) Return(_a0 []iam.User, _a1 error) *MockAccountUsersInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountUsersInterface_ListAll_Call) RunAndReturn(run func(context.Context, iam.ListAccountUsersRequest) ([]iam.User, error)) *MockAccountUsersInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, request
func (_m *MockAccountUsersInterface) Patch(ctx context.Context, request iam.PartialUpdate) error {
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

// MockAccountUsersInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type MockAccountUsersInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.PartialUpdate
func (_e *MockAccountUsersInterface_Expecter) Patch(ctx interface{}, request interface{}) *MockAccountUsersInterface_Patch_Call {
	return &MockAccountUsersInterface_Patch_Call{Call: _e.mock.On("Patch", ctx, request)}
}

func (_c *MockAccountUsersInterface_Patch_Call) Run(run func(ctx context.Context, request iam.PartialUpdate)) *MockAccountUsersInterface_Patch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.PartialUpdate))
	})
	return _c
}

func (_c *MockAccountUsersInterface_Patch_Call) Return(_a0 error) *MockAccountUsersInterface_Patch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountUsersInterface_Patch_Call) RunAndReturn(run func(context.Context, iam.PartialUpdate) error) *MockAccountUsersInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockAccountUsersInterface) Update(ctx context.Context, request iam.User) error {
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

// MockAccountUsersInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockAccountUsersInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.User
func (_e *MockAccountUsersInterface_Expecter) Update(ctx interface{}, request interface{}) *MockAccountUsersInterface_Update_Call {
	return &MockAccountUsersInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockAccountUsersInterface_Update_Call) Run(run func(ctx context.Context, request iam.User)) *MockAccountUsersInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.User))
	})
	return _c
}

func (_c *MockAccountUsersInterface_Update_Call) Return(_a0 error) *MockAccountUsersInterface_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountUsersInterface_Update_Call) RunAndReturn(run func(context.Context, iam.User) error) *MockAccountUsersInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UserUserNameToIdMap provides a mock function with given fields: ctx, request
func (_m *MockAccountUsersInterface) UserUserNameToIdMap(ctx context.Context, request iam.ListAccountUsersRequest) (map[string]string, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for UserUserNameToIdMap")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListAccountUsersRequest) (map[string]string, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListAccountUsersRequest) map[string]string); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.ListAccountUsersRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountUsersInterface_UserUserNameToIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserUserNameToIdMap'
type MockAccountUsersInterface_UserUserNameToIdMap_Call struct {
	*mock.Call
}

// UserUserNameToIdMap is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.ListAccountUsersRequest
func (_e *MockAccountUsersInterface_Expecter) UserUserNameToIdMap(ctx interface{}, request interface{}) *MockAccountUsersInterface_UserUserNameToIdMap_Call {
	return &MockAccountUsersInterface_UserUserNameToIdMap_Call{Call: _e.mock.On("UserUserNameToIdMap", ctx, request)}
}

func (_c *MockAccountUsersInterface_UserUserNameToIdMap_Call) Run(run func(ctx context.Context, request iam.ListAccountUsersRequest)) *MockAccountUsersInterface_UserUserNameToIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.ListAccountUsersRequest))
	})
	return _c
}

func (_c *MockAccountUsersInterface_UserUserNameToIdMap_Call) Return(_a0 map[string]string, _a1 error) *MockAccountUsersInterface_UserUserNameToIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountUsersInterface_UserUserNameToIdMap_Call) RunAndReturn(run func(context.Context, iam.ListAccountUsersRequest) (map[string]string, error)) *MockAccountUsersInterface_UserUserNameToIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockAccountUsersInterface) WithImpl(impl iam.AccountUsersService) iam.AccountUsersInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 iam.AccountUsersInterface
	if rf, ok := ret.Get(0).(func(iam.AccountUsersService) iam.AccountUsersInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(iam.AccountUsersInterface)
		}
	}

	return r0
}

// MockAccountUsersInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockAccountUsersInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl iam.AccountUsersService
func (_e *MockAccountUsersInterface_Expecter) WithImpl(impl interface{}) *MockAccountUsersInterface_WithImpl_Call {
	return &MockAccountUsersInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockAccountUsersInterface_WithImpl_Call) Run(run func(impl iam.AccountUsersService)) *MockAccountUsersInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(iam.AccountUsersService))
	})
	return _c
}

func (_c *MockAccountUsersInterface_WithImpl_Call) Return(_a0 iam.AccountUsersInterface) *MockAccountUsersInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountUsersInterface_WithImpl_Call) RunAndReturn(run func(iam.AccountUsersService) iam.AccountUsersInterface) *MockAccountUsersInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAccountUsersInterface creates a new instance of MockAccountUsersInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAccountUsersInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAccountUsersInterface {
	mock := &MockAccountUsersInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
