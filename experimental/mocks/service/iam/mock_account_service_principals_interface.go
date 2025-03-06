// Code generated by mockery v2.53.0. DO NOT EDIT.

package iam

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	iam "github.com/databricks/databricks-sdk-go/service/iam"

	mock "github.com/stretchr/testify/mock"
)

// MockAccountServicePrincipalsInterface is an autogenerated mock type for the AccountServicePrincipalsInterface type
type MockAccountServicePrincipalsInterface struct {
	mock.Mock
}

type MockAccountServicePrincipalsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAccountServicePrincipalsInterface) EXPECT() *MockAccountServicePrincipalsInterface_Expecter {
	return &MockAccountServicePrincipalsInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockAccountServicePrincipalsInterface) Create(ctx context.Context, request iam.ServicePrincipal) (*iam.ServicePrincipal, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *iam.ServicePrincipal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.ServicePrincipal) (*iam.ServicePrincipal, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.ServicePrincipal) *iam.ServicePrincipal); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.ServicePrincipal)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.ServicePrincipal) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountServicePrincipalsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockAccountServicePrincipalsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.ServicePrincipal
func (_e *MockAccountServicePrincipalsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockAccountServicePrincipalsInterface_Create_Call {
	return &MockAccountServicePrincipalsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockAccountServicePrincipalsInterface_Create_Call) Run(run func(ctx context.Context, request iam.ServicePrincipal)) *MockAccountServicePrincipalsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.ServicePrincipal))
	})
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_Create_Call) Return(_a0 *iam.ServicePrincipal, _a1 error) *MockAccountServicePrincipalsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_Create_Call) RunAndReturn(run func(context.Context, iam.ServicePrincipal) (*iam.ServicePrincipal, error)) *MockAccountServicePrincipalsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockAccountServicePrincipalsInterface) Delete(ctx context.Context, request iam.DeleteAccountServicePrincipalRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.DeleteAccountServicePrincipalRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAccountServicePrincipalsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockAccountServicePrincipalsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.DeleteAccountServicePrincipalRequest
func (_e *MockAccountServicePrincipalsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockAccountServicePrincipalsInterface_Delete_Call {
	return &MockAccountServicePrincipalsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockAccountServicePrincipalsInterface_Delete_Call) Run(run func(ctx context.Context, request iam.DeleteAccountServicePrincipalRequest)) *MockAccountServicePrincipalsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.DeleteAccountServicePrincipalRequest))
	})
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_Delete_Call) Return(_a0 error) *MockAccountServicePrincipalsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_Delete_Call) RunAndReturn(run func(context.Context, iam.DeleteAccountServicePrincipalRequest) error) *MockAccountServicePrincipalsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *MockAccountServicePrincipalsInterface) DeleteById(ctx context.Context, id string) error {
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

// MockAccountServicePrincipalsInterface_DeleteById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteById'
type MockAccountServicePrincipalsInterface_DeleteById_Call struct {
	*mock.Call
}

// DeleteById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockAccountServicePrincipalsInterface_Expecter) DeleteById(ctx interface{}, id interface{}) *MockAccountServicePrincipalsInterface_DeleteById_Call {
	return &MockAccountServicePrincipalsInterface_DeleteById_Call{Call: _e.mock.On("DeleteById", ctx, id)}
}

func (_c *MockAccountServicePrincipalsInterface_DeleteById_Call) Run(run func(ctx context.Context, id string)) *MockAccountServicePrincipalsInterface_DeleteById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_DeleteById_Call) Return(_a0 error) *MockAccountServicePrincipalsInterface_DeleteById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_DeleteById_Call) RunAndReturn(run func(context.Context, string) error) *MockAccountServicePrincipalsInterface_DeleteById_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockAccountServicePrincipalsInterface) Get(ctx context.Context, request iam.GetAccountServicePrincipalRequest) (*iam.ServicePrincipal, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *iam.ServicePrincipal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.GetAccountServicePrincipalRequest) (*iam.ServicePrincipal, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.GetAccountServicePrincipalRequest) *iam.ServicePrincipal); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.ServicePrincipal)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.GetAccountServicePrincipalRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountServicePrincipalsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockAccountServicePrincipalsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.GetAccountServicePrincipalRequest
func (_e *MockAccountServicePrincipalsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockAccountServicePrincipalsInterface_Get_Call {
	return &MockAccountServicePrincipalsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockAccountServicePrincipalsInterface_Get_Call) Run(run func(ctx context.Context, request iam.GetAccountServicePrincipalRequest)) *MockAccountServicePrincipalsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.GetAccountServicePrincipalRequest))
	})
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_Get_Call) Return(_a0 *iam.ServicePrincipal, _a1 error) *MockAccountServicePrincipalsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_Get_Call) RunAndReturn(run func(context.Context, iam.GetAccountServicePrincipalRequest) (*iam.ServicePrincipal, error)) *MockAccountServicePrincipalsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByDisplayName provides a mock function with given fields: ctx, name
func (_m *MockAccountServicePrincipalsInterface) GetByDisplayName(ctx context.Context, name string) (*iam.ServicePrincipal, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByDisplayName")
	}

	var r0 *iam.ServicePrincipal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*iam.ServicePrincipal, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *iam.ServicePrincipal); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.ServicePrincipal)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountServicePrincipalsInterface_GetByDisplayName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByDisplayName'
type MockAccountServicePrincipalsInterface_GetByDisplayName_Call struct {
	*mock.Call
}

// GetByDisplayName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockAccountServicePrincipalsInterface_Expecter) GetByDisplayName(ctx interface{}, name interface{}) *MockAccountServicePrincipalsInterface_GetByDisplayName_Call {
	return &MockAccountServicePrincipalsInterface_GetByDisplayName_Call{Call: _e.mock.On("GetByDisplayName", ctx, name)}
}

func (_c *MockAccountServicePrincipalsInterface_GetByDisplayName_Call) Run(run func(ctx context.Context, name string)) *MockAccountServicePrincipalsInterface_GetByDisplayName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_GetByDisplayName_Call) Return(_a0 *iam.ServicePrincipal, _a1 error) *MockAccountServicePrincipalsInterface_GetByDisplayName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_GetByDisplayName_Call) RunAndReturn(run func(context.Context, string) (*iam.ServicePrincipal, error)) *MockAccountServicePrincipalsInterface_GetByDisplayName_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: ctx, id
func (_m *MockAccountServicePrincipalsInterface) GetById(ctx context.Context, id string) (*iam.ServicePrincipal, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *iam.ServicePrincipal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*iam.ServicePrincipal, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *iam.ServicePrincipal); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.ServicePrincipal)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountServicePrincipalsInterface_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type MockAccountServicePrincipalsInterface_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockAccountServicePrincipalsInterface_Expecter) GetById(ctx interface{}, id interface{}) *MockAccountServicePrincipalsInterface_GetById_Call {
	return &MockAccountServicePrincipalsInterface_GetById_Call{Call: _e.mock.On("GetById", ctx, id)}
}

func (_c *MockAccountServicePrincipalsInterface_GetById_Call) Run(run func(ctx context.Context, id string)) *MockAccountServicePrincipalsInterface_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_GetById_Call) Return(_a0 *iam.ServicePrincipal, _a1 error) *MockAccountServicePrincipalsInterface_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_GetById_Call) RunAndReturn(run func(context.Context, string) (*iam.ServicePrincipal, error)) *MockAccountServicePrincipalsInterface_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockAccountServicePrincipalsInterface) List(ctx context.Context, request iam.ListAccountServicePrincipalsRequest) listing.Iterator[iam.ServicePrincipal] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[iam.ServicePrincipal]
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListAccountServicePrincipalsRequest) listing.Iterator[iam.ServicePrincipal]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[iam.ServicePrincipal])
		}
	}

	return r0
}

// MockAccountServicePrincipalsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockAccountServicePrincipalsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.ListAccountServicePrincipalsRequest
func (_e *MockAccountServicePrincipalsInterface_Expecter) List(ctx interface{}, request interface{}) *MockAccountServicePrincipalsInterface_List_Call {
	return &MockAccountServicePrincipalsInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockAccountServicePrincipalsInterface_List_Call) Run(run func(ctx context.Context, request iam.ListAccountServicePrincipalsRequest)) *MockAccountServicePrincipalsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.ListAccountServicePrincipalsRequest))
	})
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_List_Call) Return(_a0 listing.Iterator[iam.ServicePrincipal]) *MockAccountServicePrincipalsInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_List_Call) RunAndReturn(run func(context.Context, iam.ListAccountServicePrincipalsRequest) listing.Iterator[iam.ServicePrincipal]) *MockAccountServicePrincipalsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockAccountServicePrincipalsInterface) ListAll(ctx context.Context, request iam.ListAccountServicePrincipalsRequest) ([]iam.ServicePrincipal, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []iam.ServicePrincipal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListAccountServicePrincipalsRequest) ([]iam.ServicePrincipal, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListAccountServicePrincipalsRequest) []iam.ServicePrincipal); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]iam.ServicePrincipal)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.ListAccountServicePrincipalsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountServicePrincipalsInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockAccountServicePrincipalsInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.ListAccountServicePrincipalsRequest
func (_e *MockAccountServicePrincipalsInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockAccountServicePrincipalsInterface_ListAll_Call {
	return &MockAccountServicePrincipalsInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockAccountServicePrincipalsInterface_ListAll_Call) Run(run func(ctx context.Context, request iam.ListAccountServicePrincipalsRequest)) *MockAccountServicePrincipalsInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.ListAccountServicePrincipalsRequest))
	})
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_ListAll_Call) Return(_a0 []iam.ServicePrincipal, _a1 error) *MockAccountServicePrincipalsInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_ListAll_Call) RunAndReturn(run func(context.Context, iam.ListAccountServicePrincipalsRequest) ([]iam.ServicePrincipal, error)) *MockAccountServicePrincipalsInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, request
func (_m *MockAccountServicePrincipalsInterface) Patch(ctx context.Context, request iam.PartialUpdate) error {
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

// MockAccountServicePrincipalsInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type MockAccountServicePrincipalsInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.PartialUpdate
func (_e *MockAccountServicePrincipalsInterface_Expecter) Patch(ctx interface{}, request interface{}) *MockAccountServicePrincipalsInterface_Patch_Call {
	return &MockAccountServicePrincipalsInterface_Patch_Call{Call: _e.mock.On("Patch", ctx, request)}
}

func (_c *MockAccountServicePrincipalsInterface_Patch_Call) Run(run func(ctx context.Context, request iam.PartialUpdate)) *MockAccountServicePrincipalsInterface_Patch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.PartialUpdate))
	})
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_Patch_Call) Return(_a0 error) *MockAccountServicePrincipalsInterface_Patch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_Patch_Call) RunAndReturn(run func(context.Context, iam.PartialUpdate) error) *MockAccountServicePrincipalsInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// ServicePrincipalDisplayNameToIdMap provides a mock function with given fields: ctx, request
func (_m *MockAccountServicePrincipalsInterface) ServicePrincipalDisplayNameToIdMap(ctx context.Context, request iam.ListAccountServicePrincipalsRequest) (map[string]string, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ServicePrincipalDisplayNameToIdMap")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListAccountServicePrincipalsRequest) (map[string]string, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListAccountServicePrincipalsRequest) map[string]string); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.ListAccountServicePrincipalsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountServicePrincipalsInterface_ServicePrincipalDisplayNameToIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ServicePrincipalDisplayNameToIdMap'
type MockAccountServicePrincipalsInterface_ServicePrincipalDisplayNameToIdMap_Call struct {
	*mock.Call
}

// ServicePrincipalDisplayNameToIdMap is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.ListAccountServicePrincipalsRequest
func (_e *MockAccountServicePrincipalsInterface_Expecter) ServicePrincipalDisplayNameToIdMap(ctx interface{}, request interface{}) *MockAccountServicePrincipalsInterface_ServicePrincipalDisplayNameToIdMap_Call {
	return &MockAccountServicePrincipalsInterface_ServicePrincipalDisplayNameToIdMap_Call{Call: _e.mock.On("ServicePrincipalDisplayNameToIdMap", ctx, request)}
}

func (_c *MockAccountServicePrincipalsInterface_ServicePrincipalDisplayNameToIdMap_Call) Run(run func(ctx context.Context, request iam.ListAccountServicePrincipalsRequest)) *MockAccountServicePrincipalsInterface_ServicePrincipalDisplayNameToIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.ListAccountServicePrincipalsRequest))
	})
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_ServicePrincipalDisplayNameToIdMap_Call) Return(_a0 map[string]string, _a1 error) *MockAccountServicePrincipalsInterface_ServicePrincipalDisplayNameToIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_ServicePrincipalDisplayNameToIdMap_Call) RunAndReturn(run func(context.Context, iam.ListAccountServicePrincipalsRequest) (map[string]string, error)) *MockAccountServicePrincipalsInterface_ServicePrincipalDisplayNameToIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockAccountServicePrincipalsInterface) Update(ctx context.Context, request iam.ServicePrincipal) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.ServicePrincipal) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAccountServicePrincipalsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockAccountServicePrincipalsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.ServicePrincipal
func (_e *MockAccountServicePrincipalsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockAccountServicePrincipalsInterface_Update_Call {
	return &MockAccountServicePrincipalsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockAccountServicePrincipalsInterface_Update_Call) Run(run func(ctx context.Context, request iam.ServicePrincipal)) *MockAccountServicePrincipalsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.ServicePrincipal))
	})
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_Update_Call) Return(_a0 error) *MockAccountServicePrincipalsInterface_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountServicePrincipalsInterface_Update_Call) RunAndReturn(run func(context.Context, iam.ServicePrincipal) error) *MockAccountServicePrincipalsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAccountServicePrincipalsInterface creates a new instance of MockAccountServicePrincipalsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAccountServicePrincipalsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAccountServicePrincipalsInterface {
	mock := &MockAccountServicePrincipalsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
