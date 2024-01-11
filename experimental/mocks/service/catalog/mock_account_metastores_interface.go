// Code generated by mockery v2.39.1. DO NOT EDIT.

package catalog

import (
	context "context"

	catalog "github.com/databricks/databricks-sdk-go/service/catalog"

	listing "github.com/databricks/databricks-sdk-go/listing"

	mock "github.com/stretchr/testify/mock"
)

// MockAccountMetastoresInterface is an autogenerated mock type for the AccountMetastoresInterface type
type MockAccountMetastoresInterface struct {
	mock.Mock
}

type MockAccountMetastoresInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAccountMetastoresInterface) EXPECT() *MockAccountMetastoresInterface_Expecter {
	return &MockAccountMetastoresInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockAccountMetastoresInterface) Create(ctx context.Context, request catalog.AccountsCreateMetastore) (*catalog.AccountsMetastoreInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *catalog.AccountsMetastoreInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.AccountsCreateMetastore) (*catalog.AccountsMetastoreInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.AccountsCreateMetastore) *catalog.AccountsMetastoreInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.AccountsMetastoreInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.AccountsCreateMetastore) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountMetastoresInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockAccountMetastoresInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.AccountsCreateMetastore
func (_e *MockAccountMetastoresInterface_Expecter) Create(ctx interface{}, request interface{}) *MockAccountMetastoresInterface_Create_Call {
	return &MockAccountMetastoresInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockAccountMetastoresInterface_Create_Call) Run(run func(ctx context.Context, request catalog.AccountsCreateMetastore)) *MockAccountMetastoresInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.AccountsCreateMetastore))
	})
	return _c
}

func (_c *MockAccountMetastoresInterface_Create_Call) Return(_a0 *catalog.AccountsMetastoreInfo, _a1 error) *MockAccountMetastoresInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountMetastoresInterface_Create_Call) RunAndReturn(run func(context.Context, catalog.AccountsCreateMetastore) (*catalog.AccountsMetastoreInfo, error)) *MockAccountMetastoresInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockAccountMetastoresInterface) Delete(ctx context.Context, request catalog.DeleteAccountMetastoreRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.DeleteAccountMetastoreRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAccountMetastoresInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockAccountMetastoresInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.DeleteAccountMetastoreRequest
func (_e *MockAccountMetastoresInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockAccountMetastoresInterface_Delete_Call {
	return &MockAccountMetastoresInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockAccountMetastoresInterface_Delete_Call) Run(run func(ctx context.Context, request catalog.DeleteAccountMetastoreRequest)) *MockAccountMetastoresInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.DeleteAccountMetastoreRequest))
	})
	return _c
}

func (_c *MockAccountMetastoresInterface_Delete_Call) Return(_a0 error) *MockAccountMetastoresInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountMetastoresInterface_Delete_Call) RunAndReturn(run func(context.Context, catalog.DeleteAccountMetastoreRequest) error) *MockAccountMetastoresInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByMetastoreId provides a mock function with given fields: ctx, metastoreId
func (_m *MockAccountMetastoresInterface) DeleteByMetastoreId(ctx context.Context, metastoreId string) error {
	ret := _m.Called(ctx, metastoreId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByMetastoreId")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, metastoreId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAccountMetastoresInterface_DeleteByMetastoreId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByMetastoreId'
type MockAccountMetastoresInterface_DeleteByMetastoreId_Call struct {
	*mock.Call
}

// DeleteByMetastoreId is a helper method to define mock.On call
//   - ctx context.Context
//   - metastoreId string
func (_e *MockAccountMetastoresInterface_Expecter) DeleteByMetastoreId(ctx interface{}, metastoreId interface{}) *MockAccountMetastoresInterface_DeleteByMetastoreId_Call {
	return &MockAccountMetastoresInterface_DeleteByMetastoreId_Call{Call: _e.mock.On("DeleteByMetastoreId", ctx, metastoreId)}
}

func (_c *MockAccountMetastoresInterface_DeleteByMetastoreId_Call) Run(run func(ctx context.Context, metastoreId string)) *MockAccountMetastoresInterface_DeleteByMetastoreId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAccountMetastoresInterface_DeleteByMetastoreId_Call) Return(_a0 error) *MockAccountMetastoresInterface_DeleteByMetastoreId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountMetastoresInterface_DeleteByMetastoreId_Call) RunAndReturn(run func(context.Context, string) error) *MockAccountMetastoresInterface_DeleteByMetastoreId_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockAccountMetastoresInterface) Get(ctx context.Context, request catalog.GetAccountMetastoreRequest) (*catalog.AccountsMetastoreInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *catalog.AccountsMetastoreInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetAccountMetastoreRequest) (*catalog.AccountsMetastoreInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetAccountMetastoreRequest) *catalog.AccountsMetastoreInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.AccountsMetastoreInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.GetAccountMetastoreRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountMetastoresInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockAccountMetastoresInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.GetAccountMetastoreRequest
func (_e *MockAccountMetastoresInterface_Expecter) Get(ctx interface{}, request interface{}) *MockAccountMetastoresInterface_Get_Call {
	return &MockAccountMetastoresInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockAccountMetastoresInterface_Get_Call) Run(run func(ctx context.Context, request catalog.GetAccountMetastoreRequest)) *MockAccountMetastoresInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.GetAccountMetastoreRequest))
	})
	return _c
}

func (_c *MockAccountMetastoresInterface_Get_Call) Return(_a0 *catalog.AccountsMetastoreInfo, _a1 error) *MockAccountMetastoresInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountMetastoresInterface_Get_Call) RunAndReturn(run func(context.Context, catalog.GetAccountMetastoreRequest) (*catalog.AccountsMetastoreInfo, error)) *MockAccountMetastoresInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByMetastoreId provides a mock function with given fields: ctx, metastoreId
func (_m *MockAccountMetastoresInterface) GetByMetastoreId(ctx context.Context, metastoreId string) (*catalog.AccountsMetastoreInfo, error) {
	ret := _m.Called(ctx, metastoreId)

	if len(ret) == 0 {
		panic("no return value specified for GetByMetastoreId")
	}

	var r0 *catalog.AccountsMetastoreInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*catalog.AccountsMetastoreInfo, error)); ok {
		return rf(ctx, metastoreId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *catalog.AccountsMetastoreInfo); ok {
		r0 = rf(ctx, metastoreId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.AccountsMetastoreInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, metastoreId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountMetastoresInterface_GetByMetastoreId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByMetastoreId'
type MockAccountMetastoresInterface_GetByMetastoreId_Call struct {
	*mock.Call
}

// GetByMetastoreId is a helper method to define mock.On call
//   - ctx context.Context
//   - metastoreId string
func (_e *MockAccountMetastoresInterface_Expecter) GetByMetastoreId(ctx interface{}, metastoreId interface{}) *MockAccountMetastoresInterface_GetByMetastoreId_Call {
	return &MockAccountMetastoresInterface_GetByMetastoreId_Call{Call: _e.mock.On("GetByMetastoreId", ctx, metastoreId)}
}

func (_c *MockAccountMetastoresInterface_GetByMetastoreId_Call) Run(run func(ctx context.Context, metastoreId string)) *MockAccountMetastoresInterface_GetByMetastoreId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAccountMetastoresInterface_GetByMetastoreId_Call) Return(_a0 *catalog.AccountsMetastoreInfo, _a1 error) *MockAccountMetastoresInterface_GetByMetastoreId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountMetastoresInterface_GetByMetastoreId_Call) RunAndReturn(run func(context.Context, string) (*catalog.AccountsMetastoreInfo, error)) *MockAccountMetastoresInterface_GetByMetastoreId_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockAccountMetastoresInterface) Impl() catalog.AccountMetastoresService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 catalog.AccountMetastoresService
	if rf, ok := ret.Get(0).(func() catalog.AccountMetastoresService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.AccountMetastoresService)
		}
	}

	return r0
}

// MockAccountMetastoresInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockAccountMetastoresInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockAccountMetastoresInterface_Expecter) Impl() *MockAccountMetastoresInterface_Impl_Call {
	return &MockAccountMetastoresInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockAccountMetastoresInterface_Impl_Call) Run(run func()) *MockAccountMetastoresInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAccountMetastoresInterface_Impl_Call) Return(_a0 catalog.AccountMetastoresService) *MockAccountMetastoresInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountMetastoresInterface_Impl_Call) RunAndReturn(run func() catalog.AccountMetastoresService) *MockAccountMetastoresInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx
func (_m *MockAccountMetastoresInterface) List(ctx context.Context) *listing.PaginatingIterator[struct{}, *catalog.ListMetastoresResponse, catalog.MetastoreInfo] {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *listing.PaginatingIterator[struct{}, *catalog.ListMetastoresResponse, catalog.MetastoreInfo]
	if rf, ok := ret.Get(0).(func(context.Context) *listing.PaginatingIterator[struct{}, *catalog.ListMetastoresResponse, catalog.MetastoreInfo]); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*listing.PaginatingIterator[struct{}, *catalog.ListMetastoresResponse, catalog.MetastoreInfo])
		}
	}

	return r0
}

// MockAccountMetastoresInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockAccountMetastoresInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockAccountMetastoresInterface_Expecter) List(ctx interface{}) *MockAccountMetastoresInterface_List_Call {
	return &MockAccountMetastoresInterface_List_Call{Call: _e.mock.On("List", ctx)}
}

func (_c *MockAccountMetastoresInterface_List_Call) Run(run func(ctx context.Context)) *MockAccountMetastoresInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockAccountMetastoresInterface_List_Call) Return(_a0 *listing.PaginatingIterator[struct{}, *catalog.ListMetastoresResponse, catalog.MetastoreInfo]) *MockAccountMetastoresInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountMetastoresInterface_List_Call) RunAndReturn(run func(context.Context) *listing.PaginatingIterator[struct{}, *catalog.ListMetastoresResponse, catalog.MetastoreInfo]) *MockAccountMetastoresInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx
func (_m *MockAccountMetastoresInterface) ListAll(ctx context.Context) ([]catalog.MetastoreInfo, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []catalog.MetastoreInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]catalog.MetastoreInfo, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []catalog.MetastoreInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]catalog.MetastoreInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountMetastoresInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockAccountMetastoresInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockAccountMetastoresInterface_Expecter) ListAll(ctx interface{}) *MockAccountMetastoresInterface_ListAll_Call {
	return &MockAccountMetastoresInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx)}
}

func (_c *MockAccountMetastoresInterface_ListAll_Call) Run(run func(ctx context.Context)) *MockAccountMetastoresInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockAccountMetastoresInterface_ListAll_Call) Return(_a0 []catalog.MetastoreInfo, _a1 error) *MockAccountMetastoresInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountMetastoresInterface_ListAll_Call) RunAndReturn(run func(context.Context) ([]catalog.MetastoreInfo, error)) *MockAccountMetastoresInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockAccountMetastoresInterface) Update(ctx context.Context, request catalog.AccountsUpdateMetastore) (*catalog.AccountsMetastoreInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *catalog.AccountsMetastoreInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.AccountsUpdateMetastore) (*catalog.AccountsMetastoreInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.AccountsUpdateMetastore) *catalog.AccountsMetastoreInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.AccountsMetastoreInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.AccountsUpdateMetastore) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountMetastoresInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockAccountMetastoresInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.AccountsUpdateMetastore
func (_e *MockAccountMetastoresInterface_Expecter) Update(ctx interface{}, request interface{}) *MockAccountMetastoresInterface_Update_Call {
	return &MockAccountMetastoresInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockAccountMetastoresInterface_Update_Call) Run(run func(ctx context.Context, request catalog.AccountsUpdateMetastore)) *MockAccountMetastoresInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.AccountsUpdateMetastore))
	})
	return _c
}

func (_c *MockAccountMetastoresInterface_Update_Call) Return(_a0 *catalog.AccountsMetastoreInfo, _a1 error) *MockAccountMetastoresInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountMetastoresInterface_Update_Call) RunAndReturn(run func(context.Context, catalog.AccountsUpdateMetastore) (*catalog.AccountsMetastoreInfo, error)) *MockAccountMetastoresInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockAccountMetastoresInterface) WithImpl(impl catalog.AccountMetastoresService) catalog.AccountMetastoresInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 catalog.AccountMetastoresInterface
	if rf, ok := ret.Get(0).(func(catalog.AccountMetastoresService) catalog.AccountMetastoresInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.AccountMetastoresInterface)
		}
	}

	return r0
}

// MockAccountMetastoresInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockAccountMetastoresInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl catalog.AccountMetastoresService
func (_e *MockAccountMetastoresInterface_Expecter) WithImpl(impl interface{}) *MockAccountMetastoresInterface_WithImpl_Call {
	return &MockAccountMetastoresInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockAccountMetastoresInterface_WithImpl_Call) Run(run func(impl catalog.AccountMetastoresService)) *MockAccountMetastoresInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(catalog.AccountMetastoresService))
	})
	return _c
}

func (_c *MockAccountMetastoresInterface_WithImpl_Call) Return(_a0 catalog.AccountMetastoresInterface) *MockAccountMetastoresInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountMetastoresInterface_WithImpl_Call) RunAndReturn(run func(catalog.AccountMetastoresService) catalog.AccountMetastoresInterface) *MockAccountMetastoresInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAccountMetastoresInterface creates a new instance of MockAccountMetastoresInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAccountMetastoresInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAccountMetastoresInterface {
	mock := &MockAccountMetastoresInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
