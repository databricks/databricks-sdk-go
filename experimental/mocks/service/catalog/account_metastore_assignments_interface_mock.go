// Code generated by mockery v2.38.0. DO NOT EDIT.

package catalog

import (
	context "context"

	catalog "github.com/databricks/databricks-sdk-go/service/catalog"

	listing "github.com/databricks/databricks-sdk-go/listing"

	mock "github.com/stretchr/testify/mock"
)

// MockAccountMetastoreAssignmentsInterface is an autogenerated mock type for the AccountMetastoreAssignmentsInterface type
type MockAccountMetastoreAssignmentsInterface struct {
	mock.Mock
}

type MockAccountMetastoreAssignmentsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAccountMetastoreAssignmentsInterface) EXPECT() *MockAccountMetastoreAssignmentsInterface_Expecter {
	return &MockAccountMetastoreAssignmentsInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockAccountMetastoreAssignmentsInterface) Create(ctx context.Context, request catalog.AccountsCreateMetastoreAssignment) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.AccountsCreateMetastoreAssignment) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAccountMetastoreAssignmentsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockAccountMetastoreAssignmentsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.AccountsCreateMetastoreAssignment
func (_e *MockAccountMetastoreAssignmentsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockAccountMetastoreAssignmentsInterface_Create_Call {
	return &MockAccountMetastoreAssignmentsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockAccountMetastoreAssignmentsInterface_Create_Call) Run(run func(ctx context.Context, request catalog.AccountsCreateMetastoreAssignment)) *MockAccountMetastoreAssignmentsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.AccountsCreateMetastoreAssignment))
	})
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_Create_Call) Return(_a0 error) *MockAccountMetastoreAssignmentsInterface_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_Create_Call) RunAndReturn(run func(context.Context, catalog.AccountsCreateMetastoreAssignment) error) *MockAccountMetastoreAssignmentsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockAccountMetastoreAssignmentsInterface) Delete(ctx context.Context, request catalog.DeleteAccountMetastoreAssignmentRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.DeleteAccountMetastoreAssignmentRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAccountMetastoreAssignmentsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockAccountMetastoreAssignmentsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.DeleteAccountMetastoreAssignmentRequest
func (_e *MockAccountMetastoreAssignmentsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockAccountMetastoreAssignmentsInterface_Delete_Call {
	return &MockAccountMetastoreAssignmentsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockAccountMetastoreAssignmentsInterface_Delete_Call) Run(run func(ctx context.Context, request catalog.DeleteAccountMetastoreAssignmentRequest)) *MockAccountMetastoreAssignmentsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.DeleteAccountMetastoreAssignmentRequest))
	})
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_Delete_Call) Return(_a0 error) *MockAccountMetastoreAssignmentsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_Delete_Call) RunAndReturn(run func(context.Context, catalog.DeleteAccountMetastoreAssignmentRequest) error) *MockAccountMetastoreAssignmentsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByWorkspaceIdAndMetastoreId provides a mock function with given fields: ctx, workspaceId, metastoreId
func (_m *MockAccountMetastoreAssignmentsInterface) DeleteByWorkspaceIdAndMetastoreId(ctx context.Context, workspaceId int64, metastoreId string) error {
	ret := _m.Called(ctx, workspaceId, metastoreId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByWorkspaceIdAndMetastoreId")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) error); ok {
		r0 = rf(ctx, workspaceId, metastoreId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAccountMetastoreAssignmentsInterface_DeleteByWorkspaceIdAndMetastoreId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByWorkspaceIdAndMetastoreId'
type MockAccountMetastoreAssignmentsInterface_DeleteByWorkspaceIdAndMetastoreId_Call struct {
	*mock.Call
}

// DeleteByWorkspaceIdAndMetastoreId is a helper method to define mock.On call
//   - ctx context.Context
//   - workspaceId int64
//   - metastoreId string
func (_e *MockAccountMetastoreAssignmentsInterface_Expecter) DeleteByWorkspaceIdAndMetastoreId(ctx interface{}, workspaceId interface{}, metastoreId interface{}) *MockAccountMetastoreAssignmentsInterface_DeleteByWorkspaceIdAndMetastoreId_Call {
	return &MockAccountMetastoreAssignmentsInterface_DeleteByWorkspaceIdAndMetastoreId_Call{Call: _e.mock.On("DeleteByWorkspaceIdAndMetastoreId", ctx, workspaceId, metastoreId)}
}

func (_c *MockAccountMetastoreAssignmentsInterface_DeleteByWorkspaceIdAndMetastoreId_Call) Run(run func(ctx context.Context, workspaceId int64, metastoreId string)) *MockAccountMetastoreAssignmentsInterface_DeleteByWorkspaceIdAndMetastoreId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(string))
	})
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_DeleteByWorkspaceIdAndMetastoreId_Call) Return(_a0 error) *MockAccountMetastoreAssignmentsInterface_DeleteByWorkspaceIdAndMetastoreId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_DeleteByWorkspaceIdAndMetastoreId_Call) RunAndReturn(run func(context.Context, int64, string) error) *MockAccountMetastoreAssignmentsInterface_DeleteByWorkspaceIdAndMetastoreId_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockAccountMetastoreAssignmentsInterface) Get(ctx context.Context, request catalog.GetAccountMetastoreAssignmentRequest) (*catalog.AccountsMetastoreAssignment, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *catalog.AccountsMetastoreAssignment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetAccountMetastoreAssignmentRequest) (*catalog.AccountsMetastoreAssignment, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetAccountMetastoreAssignmentRequest) *catalog.AccountsMetastoreAssignment); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.AccountsMetastoreAssignment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.GetAccountMetastoreAssignmentRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountMetastoreAssignmentsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockAccountMetastoreAssignmentsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.GetAccountMetastoreAssignmentRequest
func (_e *MockAccountMetastoreAssignmentsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockAccountMetastoreAssignmentsInterface_Get_Call {
	return &MockAccountMetastoreAssignmentsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockAccountMetastoreAssignmentsInterface_Get_Call) Run(run func(ctx context.Context, request catalog.GetAccountMetastoreAssignmentRequest)) *MockAccountMetastoreAssignmentsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.GetAccountMetastoreAssignmentRequest))
	})
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_Get_Call) Return(_a0 *catalog.AccountsMetastoreAssignment, _a1 error) *MockAccountMetastoreAssignmentsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_Get_Call) RunAndReturn(run func(context.Context, catalog.GetAccountMetastoreAssignmentRequest) (*catalog.AccountsMetastoreAssignment, error)) *MockAccountMetastoreAssignmentsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByWorkspaceId provides a mock function with given fields: ctx, workspaceId
func (_m *MockAccountMetastoreAssignmentsInterface) GetByWorkspaceId(ctx context.Context, workspaceId int64) (*catalog.AccountsMetastoreAssignment, error) {
	ret := _m.Called(ctx, workspaceId)

	if len(ret) == 0 {
		panic("no return value specified for GetByWorkspaceId")
	}

	var r0 *catalog.AccountsMetastoreAssignment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*catalog.AccountsMetastoreAssignment, error)); ok {
		return rf(ctx, workspaceId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *catalog.AccountsMetastoreAssignment); ok {
		r0 = rf(ctx, workspaceId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.AccountsMetastoreAssignment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, workspaceId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountMetastoreAssignmentsInterface_GetByWorkspaceId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByWorkspaceId'
type MockAccountMetastoreAssignmentsInterface_GetByWorkspaceId_Call struct {
	*mock.Call
}

// GetByWorkspaceId is a helper method to define mock.On call
//   - ctx context.Context
//   - workspaceId int64
func (_e *MockAccountMetastoreAssignmentsInterface_Expecter) GetByWorkspaceId(ctx interface{}, workspaceId interface{}) *MockAccountMetastoreAssignmentsInterface_GetByWorkspaceId_Call {
	return &MockAccountMetastoreAssignmentsInterface_GetByWorkspaceId_Call{Call: _e.mock.On("GetByWorkspaceId", ctx, workspaceId)}
}

func (_c *MockAccountMetastoreAssignmentsInterface_GetByWorkspaceId_Call) Run(run func(ctx context.Context, workspaceId int64)) *MockAccountMetastoreAssignmentsInterface_GetByWorkspaceId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_GetByWorkspaceId_Call) Return(_a0 *catalog.AccountsMetastoreAssignment, _a1 error) *MockAccountMetastoreAssignmentsInterface_GetByWorkspaceId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_GetByWorkspaceId_Call) RunAndReturn(run func(context.Context, int64) (*catalog.AccountsMetastoreAssignment, error)) *MockAccountMetastoreAssignmentsInterface_GetByWorkspaceId_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockAccountMetastoreAssignmentsInterface) Impl() catalog.AccountMetastoreAssignmentsService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 catalog.AccountMetastoreAssignmentsService
	if rf, ok := ret.Get(0).(func() catalog.AccountMetastoreAssignmentsService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.AccountMetastoreAssignmentsService)
		}
	}

	return r0
}

// MockAccountMetastoreAssignmentsInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockAccountMetastoreAssignmentsInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockAccountMetastoreAssignmentsInterface_Expecter) Impl() *MockAccountMetastoreAssignmentsInterface_Impl_Call {
	return &MockAccountMetastoreAssignmentsInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockAccountMetastoreAssignmentsInterface_Impl_Call) Run(run func()) *MockAccountMetastoreAssignmentsInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_Impl_Call) Return(_a0 catalog.AccountMetastoreAssignmentsService) *MockAccountMetastoreAssignmentsInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_Impl_Call) RunAndReturn(run func() catalog.AccountMetastoreAssignmentsService) *MockAccountMetastoreAssignmentsInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockAccountMetastoreAssignmentsInterface) List(ctx context.Context, request catalog.ListAccountMetastoreAssignmentsRequest) *listing.PaginatingIterator[catalog.ListAccountMetastoreAssignmentsRequest, *catalog.ListAccountMetastoreAssignmentsResponse, int64] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *listing.PaginatingIterator[catalog.ListAccountMetastoreAssignmentsRequest, *catalog.ListAccountMetastoreAssignmentsResponse, int64]
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListAccountMetastoreAssignmentsRequest) *listing.PaginatingIterator[catalog.ListAccountMetastoreAssignmentsRequest, *catalog.ListAccountMetastoreAssignmentsResponse, int64]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*listing.PaginatingIterator[catalog.ListAccountMetastoreAssignmentsRequest, *catalog.ListAccountMetastoreAssignmentsResponse, int64])
		}
	}

	return r0
}

// MockAccountMetastoreAssignmentsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockAccountMetastoreAssignmentsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListAccountMetastoreAssignmentsRequest
func (_e *MockAccountMetastoreAssignmentsInterface_Expecter) List(ctx interface{}, request interface{}) *MockAccountMetastoreAssignmentsInterface_List_Call {
	return &MockAccountMetastoreAssignmentsInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockAccountMetastoreAssignmentsInterface_List_Call) Run(run func(ctx context.Context, request catalog.ListAccountMetastoreAssignmentsRequest)) *MockAccountMetastoreAssignmentsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListAccountMetastoreAssignmentsRequest))
	})
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_List_Call) Return(_a0 *listing.PaginatingIterator[catalog.ListAccountMetastoreAssignmentsRequest, *catalog.ListAccountMetastoreAssignmentsResponse, int64]) *MockAccountMetastoreAssignmentsInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_List_Call) RunAndReturn(run func(context.Context, catalog.ListAccountMetastoreAssignmentsRequest) *listing.PaginatingIterator[catalog.ListAccountMetastoreAssignmentsRequest, *catalog.ListAccountMetastoreAssignmentsResponse, int64]) *MockAccountMetastoreAssignmentsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockAccountMetastoreAssignmentsInterface) ListAll(ctx context.Context, request catalog.ListAccountMetastoreAssignmentsRequest) ([]int64, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListAccountMetastoreAssignmentsRequest) ([]int64, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListAccountMetastoreAssignmentsRequest) []int64); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.ListAccountMetastoreAssignmentsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountMetastoreAssignmentsInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockAccountMetastoreAssignmentsInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListAccountMetastoreAssignmentsRequest
func (_e *MockAccountMetastoreAssignmentsInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockAccountMetastoreAssignmentsInterface_ListAll_Call {
	return &MockAccountMetastoreAssignmentsInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockAccountMetastoreAssignmentsInterface_ListAll_Call) Run(run func(ctx context.Context, request catalog.ListAccountMetastoreAssignmentsRequest)) *MockAccountMetastoreAssignmentsInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListAccountMetastoreAssignmentsRequest))
	})
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_ListAll_Call) Return(_a0 []int64, _a1 error) *MockAccountMetastoreAssignmentsInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_ListAll_Call) RunAndReturn(run func(context.Context, catalog.ListAccountMetastoreAssignmentsRequest) ([]int64, error)) *MockAccountMetastoreAssignmentsInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// ListByMetastoreId provides a mock function with given fields: ctx, metastoreId
func (_m *MockAccountMetastoreAssignmentsInterface) ListByMetastoreId(ctx context.Context, metastoreId string) (*catalog.ListAccountMetastoreAssignmentsResponse, error) {
	ret := _m.Called(ctx, metastoreId)

	if len(ret) == 0 {
		panic("no return value specified for ListByMetastoreId")
	}

	var r0 *catalog.ListAccountMetastoreAssignmentsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*catalog.ListAccountMetastoreAssignmentsResponse, error)); ok {
		return rf(ctx, metastoreId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *catalog.ListAccountMetastoreAssignmentsResponse); ok {
		r0 = rf(ctx, metastoreId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.ListAccountMetastoreAssignmentsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, metastoreId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountMetastoreAssignmentsInterface_ListByMetastoreId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListByMetastoreId'
type MockAccountMetastoreAssignmentsInterface_ListByMetastoreId_Call struct {
	*mock.Call
}

// ListByMetastoreId is a helper method to define mock.On call
//   - ctx context.Context
//   - metastoreId string
func (_e *MockAccountMetastoreAssignmentsInterface_Expecter) ListByMetastoreId(ctx interface{}, metastoreId interface{}) *MockAccountMetastoreAssignmentsInterface_ListByMetastoreId_Call {
	return &MockAccountMetastoreAssignmentsInterface_ListByMetastoreId_Call{Call: _e.mock.On("ListByMetastoreId", ctx, metastoreId)}
}

func (_c *MockAccountMetastoreAssignmentsInterface_ListByMetastoreId_Call) Run(run func(ctx context.Context, metastoreId string)) *MockAccountMetastoreAssignmentsInterface_ListByMetastoreId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_ListByMetastoreId_Call) Return(_a0 *catalog.ListAccountMetastoreAssignmentsResponse, _a1 error) *MockAccountMetastoreAssignmentsInterface_ListByMetastoreId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_ListByMetastoreId_Call) RunAndReturn(run func(context.Context, string) (*catalog.ListAccountMetastoreAssignmentsResponse, error)) *MockAccountMetastoreAssignmentsInterface_ListByMetastoreId_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockAccountMetastoreAssignmentsInterface) Update(ctx context.Context, request catalog.AccountsUpdateMetastoreAssignment) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.AccountsUpdateMetastoreAssignment) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAccountMetastoreAssignmentsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockAccountMetastoreAssignmentsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.AccountsUpdateMetastoreAssignment
func (_e *MockAccountMetastoreAssignmentsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockAccountMetastoreAssignmentsInterface_Update_Call {
	return &MockAccountMetastoreAssignmentsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockAccountMetastoreAssignmentsInterface_Update_Call) Run(run func(ctx context.Context, request catalog.AccountsUpdateMetastoreAssignment)) *MockAccountMetastoreAssignmentsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.AccountsUpdateMetastoreAssignment))
	})
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_Update_Call) Return(_a0 error) *MockAccountMetastoreAssignmentsInterface_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_Update_Call) RunAndReturn(run func(context.Context, catalog.AccountsUpdateMetastoreAssignment) error) *MockAccountMetastoreAssignmentsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockAccountMetastoreAssignmentsInterface) WithImpl(impl catalog.AccountMetastoreAssignmentsService) catalog.AccountMetastoreAssignmentsInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 catalog.AccountMetastoreAssignmentsInterface
	if rf, ok := ret.Get(0).(func(catalog.AccountMetastoreAssignmentsService) catalog.AccountMetastoreAssignmentsInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.AccountMetastoreAssignmentsInterface)
		}
	}

	return r0
}

// MockAccountMetastoreAssignmentsInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockAccountMetastoreAssignmentsInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl catalog.AccountMetastoreAssignmentsService
func (_e *MockAccountMetastoreAssignmentsInterface_Expecter) WithImpl(impl interface{}) *MockAccountMetastoreAssignmentsInterface_WithImpl_Call {
	return &MockAccountMetastoreAssignmentsInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockAccountMetastoreAssignmentsInterface_WithImpl_Call) Run(run func(impl catalog.AccountMetastoreAssignmentsService)) *MockAccountMetastoreAssignmentsInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(catalog.AccountMetastoreAssignmentsService))
	})
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_WithImpl_Call) Return(_a0 catalog.AccountMetastoreAssignmentsInterface) *MockAccountMetastoreAssignmentsInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountMetastoreAssignmentsInterface_WithImpl_Call) RunAndReturn(run func(catalog.AccountMetastoreAssignmentsService) catalog.AccountMetastoreAssignmentsInterface) *MockAccountMetastoreAssignmentsInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAccountMetastoreAssignmentsInterface creates a new instance of MockAccountMetastoreAssignmentsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAccountMetastoreAssignmentsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAccountMetastoreAssignmentsInterface {
	mock := &MockAccountMetastoreAssignmentsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
