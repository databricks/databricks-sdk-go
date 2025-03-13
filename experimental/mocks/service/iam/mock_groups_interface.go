// Code generated by mockery v2.53.2. DO NOT EDIT.

package iam

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	iam "github.com/databricks/databricks-sdk-go/service/iam"

	mock "github.com/stretchr/testify/mock"
)

// MockGroupsInterface is an autogenerated mock type for the GroupsInterface type
type MockGroupsInterface struct {
	mock.Mock
}

type MockGroupsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGroupsInterface) EXPECT() *MockGroupsInterface_Expecter {
	return &MockGroupsInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockGroupsInterface) Create(ctx context.Context, request iam.Group) (*iam.Group, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *iam.Group
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.Group) (*iam.Group, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.Group) *iam.Group); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.Group) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGroupsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockGroupsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.Group
func (_e *MockGroupsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockGroupsInterface_Create_Call {
	return &MockGroupsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockGroupsInterface_Create_Call) Run(run func(ctx context.Context, request iam.Group)) *MockGroupsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.Group))
	})
	return _c
}

func (_c *MockGroupsInterface_Create_Call) Return(_a0 *iam.Group, _a1 error) *MockGroupsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGroupsInterface_Create_Call) RunAndReturn(run func(context.Context, iam.Group) (*iam.Group, error)) *MockGroupsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockGroupsInterface) Delete(ctx context.Context, request iam.DeleteGroupRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.DeleteGroupRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockGroupsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockGroupsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.DeleteGroupRequest
func (_e *MockGroupsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockGroupsInterface_Delete_Call {
	return &MockGroupsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockGroupsInterface_Delete_Call) Run(run func(ctx context.Context, request iam.DeleteGroupRequest)) *MockGroupsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.DeleteGroupRequest))
	})
	return _c
}

func (_c *MockGroupsInterface_Delete_Call) Return(_a0 error) *MockGroupsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGroupsInterface_Delete_Call) RunAndReturn(run func(context.Context, iam.DeleteGroupRequest) error) *MockGroupsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *MockGroupsInterface) DeleteById(ctx context.Context, id string) error {
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

// MockGroupsInterface_DeleteById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteById'
type MockGroupsInterface_DeleteById_Call struct {
	*mock.Call
}

// DeleteById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockGroupsInterface_Expecter) DeleteById(ctx interface{}, id interface{}) *MockGroupsInterface_DeleteById_Call {
	return &MockGroupsInterface_DeleteById_Call{Call: _e.mock.On("DeleteById", ctx, id)}
}

func (_c *MockGroupsInterface_DeleteById_Call) Run(run func(ctx context.Context, id string)) *MockGroupsInterface_DeleteById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockGroupsInterface_DeleteById_Call) Return(_a0 error) *MockGroupsInterface_DeleteById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGroupsInterface_DeleteById_Call) RunAndReturn(run func(context.Context, string) error) *MockGroupsInterface_DeleteById_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockGroupsInterface) Get(ctx context.Context, request iam.GetGroupRequest) (*iam.Group, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *iam.Group
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.GetGroupRequest) (*iam.Group, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.GetGroupRequest) *iam.Group); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.GetGroupRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGroupsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockGroupsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.GetGroupRequest
func (_e *MockGroupsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockGroupsInterface_Get_Call {
	return &MockGroupsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockGroupsInterface_Get_Call) Run(run func(ctx context.Context, request iam.GetGroupRequest)) *MockGroupsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.GetGroupRequest))
	})
	return _c
}

func (_c *MockGroupsInterface_Get_Call) Return(_a0 *iam.Group, _a1 error) *MockGroupsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGroupsInterface_Get_Call) RunAndReturn(run func(context.Context, iam.GetGroupRequest) (*iam.Group, error)) *MockGroupsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByDisplayName provides a mock function with given fields: ctx, name
func (_m *MockGroupsInterface) GetByDisplayName(ctx context.Context, name string) (*iam.Group, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByDisplayName")
	}

	var r0 *iam.Group
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*iam.Group, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *iam.Group); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGroupsInterface_GetByDisplayName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByDisplayName'
type MockGroupsInterface_GetByDisplayName_Call struct {
	*mock.Call
}

// GetByDisplayName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockGroupsInterface_Expecter) GetByDisplayName(ctx interface{}, name interface{}) *MockGroupsInterface_GetByDisplayName_Call {
	return &MockGroupsInterface_GetByDisplayName_Call{Call: _e.mock.On("GetByDisplayName", ctx, name)}
}

func (_c *MockGroupsInterface_GetByDisplayName_Call) Run(run func(ctx context.Context, name string)) *MockGroupsInterface_GetByDisplayName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockGroupsInterface_GetByDisplayName_Call) Return(_a0 *iam.Group, _a1 error) *MockGroupsInterface_GetByDisplayName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGroupsInterface_GetByDisplayName_Call) RunAndReturn(run func(context.Context, string) (*iam.Group, error)) *MockGroupsInterface_GetByDisplayName_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: ctx, id
func (_m *MockGroupsInterface) GetById(ctx context.Context, id string) (*iam.Group, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *iam.Group
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*iam.Group, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *iam.Group); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGroupsInterface_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type MockGroupsInterface_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockGroupsInterface_Expecter) GetById(ctx interface{}, id interface{}) *MockGroupsInterface_GetById_Call {
	return &MockGroupsInterface_GetById_Call{Call: _e.mock.On("GetById", ctx, id)}
}

func (_c *MockGroupsInterface_GetById_Call) Run(run func(ctx context.Context, id string)) *MockGroupsInterface_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockGroupsInterface_GetById_Call) Return(_a0 *iam.Group, _a1 error) *MockGroupsInterface_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGroupsInterface_GetById_Call) RunAndReturn(run func(context.Context, string) (*iam.Group, error)) *MockGroupsInterface_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// GroupDisplayNameToIdMap provides a mock function with given fields: ctx, request
func (_m *MockGroupsInterface) GroupDisplayNameToIdMap(ctx context.Context, request iam.ListGroupsRequest) (map[string]string, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GroupDisplayNameToIdMap")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListGroupsRequest) (map[string]string, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListGroupsRequest) map[string]string); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.ListGroupsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGroupsInterface_GroupDisplayNameToIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GroupDisplayNameToIdMap'
type MockGroupsInterface_GroupDisplayNameToIdMap_Call struct {
	*mock.Call
}

// GroupDisplayNameToIdMap is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.ListGroupsRequest
func (_e *MockGroupsInterface_Expecter) GroupDisplayNameToIdMap(ctx interface{}, request interface{}) *MockGroupsInterface_GroupDisplayNameToIdMap_Call {
	return &MockGroupsInterface_GroupDisplayNameToIdMap_Call{Call: _e.mock.On("GroupDisplayNameToIdMap", ctx, request)}
}

func (_c *MockGroupsInterface_GroupDisplayNameToIdMap_Call) Run(run func(ctx context.Context, request iam.ListGroupsRequest)) *MockGroupsInterface_GroupDisplayNameToIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.ListGroupsRequest))
	})
	return _c
}

func (_c *MockGroupsInterface_GroupDisplayNameToIdMap_Call) Return(_a0 map[string]string, _a1 error) *MockGroupsInterface_GroupDisplayNameToIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGroupsInterface_GroupDisplayNameToIdMap_Call) RunAndReturn(run func(context.Context, iam.ListGroupsRequest) (map[string]string, error)) *MockGroupsInterface_GroupDisplayNameToIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockGroupsInterface) List(ctx context.Context, request iam.ListGroupsRequest) listing.Iterator[iam.Group] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[iam.Group]
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListGroupsRequest) listing.Iterator[iam.Group]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[iam.Group])
		}
	}

	return r0
}

// MockGroupsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockGroupsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.ListGroupsRequest
func (_e *MockGroupsInterface_Expecter) List(ctx interface{}, request interface{}) *MockGroupsInterface_List_Call {
	return &MockGroupsInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockGroupsInterface_List_Call) Run(run func(ctx context.Context, request iam.ListGroupsRequest)) *MockGroupsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.ListGroupsRequest))
	})
	return _c
}

func (_c *MockGroupsInterface_List_Call) Return(_a0 listing.Iterator[iam.Group]) *MockGroupsInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGroupsInterface_List_Call) RunAndReturn(run func(context.Context, iam.ListGroupsRequest) listing.Iterator[iam.Group]) *MockGroupsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockGroupsInterface) ListAll(ctx context.Context, request iam.ListGroupsRequest) ([]iam.Group, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []iam.Group
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListGroupsRequest) ([]iam.Group, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListGroupsRequest) []iam.Group); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]iam.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.ListGroupsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGroupsInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockGroupsInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.ListGroupsRequest
func (_e *MockGroupsInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockGroupsInterface_ListAll_Call {
	return &MockGroupsInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockGroupsInterface_ListAll_Call) Run(run func(ctx context.Context, request iam.ListGroupsRequest)) *MockGroupsInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.ListGroupsRequest))
	})
	return _c
}

func (_c *MockGroupsInterface_ListAll_Call) Return(_a0 []iam.Group, _a1 error) *MockGroupsInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGroupsInterface_ListAll_Call) RunAndReturn(run func(context.Context, iam.ListGroupsRequest) ([]iam.Group, error)) *MockGroupsInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, request
func (_m *MockGroupsInterface) Patch(ctx context.Context, request iam.PartialUpdate) error {
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

// MockGroupsInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type MockGroupsInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.PartialUpdate
func (_e *MockGroupsInterface_Expecter) Patch(ctx interface{}, request interface{}) *MockGroupsInterface_Patch_Call {
	return &MockGroupsInterface_Patch_Call{Call: _e.mock.On("Patch", ctx, request)}
}

func (_c *MockGroupsInterface_Patch_Call) Run(run func(ctx context.Context, request iam.PartialUpdate)) *MockGroupsInterface_Patch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.PartialUpdate))
	})
	return _c
}

func (_c *MockGroupsInterface_Patch_Call) Return(_a0 error) *MockGroupsInterface_Patch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGroupsInterface_Patch_Call) RunAndReturn(run func(context.Context, iam.PartialUpdate) error) *MockGroupsInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockGroupsInterface) Update(ctx context.Context, request iam.Group) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.Group) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockGroupsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockGroupsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.Group
func (_e *MockGroupsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockGroupsInterface_Update_Call {
	return &MockGroupsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockGroupsInterface_Update_Call) Run(run func(ctx context.Context, request iam.Group)) *MockGroupsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.Group))
	})
	return _c
}

func (_c *MockGroupsInterface_Update_Call) Return(_a0 error) *MockGroupsInterface_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGroupsInterface_Update_Call) RunAndReturn(run func(context.Context, iam.Group) error) *MockGroupsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockGroupsInterface creates a new instance of MockGroupsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGroupsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGroupsInterface {
	mock := &MockGroupsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
