// Code generated by mockery v2.39.1. DO NOT EDIT.

package marketplace

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	marketplace "github.com/databricks/databricks-sdk-go/service/marketplace"

	mock "github.com/stretchr/testify/mock"
)

// MockProviderExchangeFiltersInterface is an autogenerated mock type for the ProviderExchangeFiltersInterface type
type MockProviderExchangeFiltersInterface struct {
	mock.Mock
}

type MockProviderExchangeFiltersInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockProviderExchangeFiltersInterface) EXPECT() *MockProviderExchangeFiltersInterface_Expecter {
	return &MockProviderExchangeFiltersInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockProviderExchangeFiltersInterface) Create(ctx context.Context, request marketplace.CreateExchangeFilterRequest) (*marketplace.CreateExchangeFilterResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *marketplace.CreateExchangeFilterResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.CreateExchangeFilterRequest) (*marketplace.CreateExchangeFilterResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.CreateExchangeFilterRequest) *marketplace.CreateExchangeFilterResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplace.CreateExchangeFilterResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, marketplace.CreateExchangeFilterRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderExchangeFiltersInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockProviderExchangeFiltersInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.CreateExchangeFilterRequest
func (_e *MockProviderExchangeFiltersInterface_Expecter) Create(ctx interface{}, request interface{}) *MockProviderExchangeFiltersInterface_Create_Call {
	return &MockProviderExchangeFiltersInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockProviderExchangeFiltersInterface_Create_Call) Run(run func(ctx context.Context, request marketplace.CreateExchangeFilterRequest)) *MockProviderExchangeFiltersInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.CreateExchangeFilterRequest))
	})
	return _c
}

func (_c *MockProviderExchangeFiltersInterface_Create_Call) Return(_a0 *marketplace.CreateExchangeFilterResponse, _a1 error) *MockProviderExchangeFiltersInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderExchangeFiltersInterface_Create_Call) RunAndReturn(run func(context.Context, marketplace.CreateExchangeFilterRequest) (*marketplace.CreateExchangeFilterResponse, error)) *MockProviderExchangeFiltersInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockProviderExchangeFiltersInterface) Delete(ctx context.Context, request marketplace.DeleteExchangeFilterRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.DeleteExchangeFilterRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockProviderExchangeFiltersInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockProviderExchangeFiltersInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.DeleteExchangeFilterRequest
func (_e *MockProviderExchangeFiltersInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockProviderExchangeFiltersInterface_Delete_Call {
	return &MockProviderExchangeFiltersInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockProviderExchangeFiltersInterface_Delete_Call) Run(run func(ctx context.Context, request marketplace.DeleteExchangeFilterRequest)) *MockProviderExchangeFiltersInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.DeleteExchangeFilterRequest))
	})
	return _c
}

func (_c *MockProviderExchangeFiltersInterface_Delete_Call) Return(_a0 error) *MockProviderExchangeFiltersInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProviderExchangeFiltersInterface_Delete_Call) RunAndReturn(run func(context.Context, marketplace.DeleteExchangeFilterRequest) error) *MockProviderExchangeFiltersInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *MockProviderExchangeFiltersInterface) DeleteById(ctx context.Context, id string) error {
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

// MockProviderExchangeFiltersInterface_DeleteById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteById'
type MockProviderExchangeFiltersInterface_DeleteById_Call struct {
	*mock.Call
}

// DeleteById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockProviderExchangeFiltersInterface_Expecter) DeleteById(ctx interface{}, id interface{}) *MockProviderExchangeFiltersInterface_DeleteById_Call {
	return &MockProviderExchangeFiltersInterface_DeleteById_Call{Call: _e.mock.On("DeleteById", ctx, id)}
}

func (_c *MockProviderExchangeFiltersInterface_DeleteById_Call) Run(run func(ctx context.Context, id string)) *MockProviderExchangeFiltersInterface_DeleteById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockProviderExchangeFiltersInterface_DeleteById_Call) Return(_a0 error) *MockProviderExchangeFiltersInterface_DeleteById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProviderExchangeFiltersInterface_DeleteById_Call) RunAndReturn(run func(context.Context, string) error) *MockProviderExchangeFiltersInterface_DeleteById_Call {
	_c.Call.Return(run)
	return _c
}

// ExchangeFilterNameToIdMap provides a mock function with given fields: ctx, request
func (_m *MockProviderExchangeFiltersInterface) ExchangeFilterNameToIdMap(ctx context.Context, request marketplace.ListExchangeFiltersRequest) (map[string]string, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ExchangeFilterNameToIdMap")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.ListExchangeFiltersRequest) (map[string]string, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.ListExchangeFiltersRequest) map[string]string); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, marketplace.ListExchangeFiltersRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderExchangeFiltersInterface_ExchangeFilterNameToIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExchangeFilterNameToIdMap'
type MockProviderExchangeFiltersInterface_ExchangeFilterNameToIdMap_Call struct {
	*mock.Call
}

// ExchangeFilterNameToIdMap is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.ListExchangeFiltersRequest
func (_e *MockProviderExchangeFiltersInterface_Expecter) ExchangeFilterNameToIdMap(ctx interface{}, request interface{}) *MockProviderExchangeFiltersInterface_ExchangeFilterNameToIdMap_Call {
	return &MockProviderExchangeFiltersInterface_ExchangeFilterNameToIdMap_Call{Call: _e.mock.On("ExchangeFilterNameToIdMap", ctx, request)}
}

func (_c *MockProviderExchangeFiltersInterface_ExchangeFilterNameToIdMap_Call) Run(run func(ctx context.Context, request marketplace.ListExchangeFiltersRequest)) *MockProviderExchangeFiltersInterface_ExchangeFilterNameToIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.ListExchangeFiltersRequest))
	})
	return _c
}

func (_c *MockProviderExchangeFiltersInterface_ExchangeFilterNameToIdMap_Call) Return(_a0 map[string]string, _a1 error) *MockProviderExchangeFiltersInterface_ExchangeFilterNameToIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderExchangeFiltersInterface_ExchangeFilterNameToIdMap_Call) RunAndReturn(run func(context.Context, marketplace.ListExchangeFiltersRequest) (map[string]string, error)) *MockProviderExchangeFiltersInterface_ExchangeFilterNameToIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *MockProviderExchangeFiltersInterface) GetByName(ctx context.Context, name string) (*marketplace.ExchangeFilter, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *marketplace.ExchangeFilter
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*marketplace.ExchangeFilter, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *marketplace.ExchangeFilter); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplace.ExchangeFilter)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderExchangeFiltersInterface_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type MockProviderExchangeFiltersInterface_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockProviderExchangeFiltersInterface_Expecter) GetByName(ctx interface{}, name interface{}) *MockProviderExchangeFiltersInterface_GetByName_Call {
	return &MockProviderExchangeFiltersInterface_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *MockProviderExchangeFiltersInterface_GetByName_Call) Run(run func(ctx context.Context, name string)) *MockProviderExchangeFiltersInterface_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockProviderExchangeFiltersInterface_GetByName_Call) Return(_a0 *marketplace.ExchangeFilter, _a1 error) *MockProviderExchangeFiltersInterface_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderExchangeFiltersInterface_GetByName_Call) RunAndReturn(run func(context.Context, string) (*marketplace.ExchangeFilter, error)) *MockProviderExchangeFiltersInterface_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockProviderExchangeFiltersInterface) Impl() marketplace.ProviderExchangeFiltersService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 marketplace.ProviderExchangeFiltersService
	if rf, ok := ret.Get(0).(func() marketplace.ProviderExchangeFiltersService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(marketplace.ProviderExchangeFiltersService)
		}
	}

	return r0
}

// MockProviderExchangeFiltersInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockProviderExchangeFiltersInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockProviderExchangeFiltersInterface_Expecter) Impl() *MockProviderExchangeFiltersInterface_Impl_Call {
	return &MockProviderExchangeFiltersInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockProviderExchangeFiltersInterface_Impl_Call) Run(run func()) *MockProviderExchangeFiltersInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProviderExchangeFiltersInterface_Impl_Call) Return(_a0 marketplace.ProviderExchangeFiltersService) *MockProviderExchangeFiltersInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProviderExchangeFiltersInterface_Impl_Call) RunAndReturn(run func() marketplace.ProviderExchangeFiltersService) *MockProviderExchangeFiltersInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockProviderExchangeFiltersInterface) List(ctx context.Context, request marketplace.ListExchangeFiltersRequest) listing.Iterator[marketplace.ExchangeFilter] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[marketplace.ExchangeFilter]
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.ListExchangeFiltersRequest) listing.Iterator[marketplace.ExchangeFilter]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[marketplace.ExchangeFilter])
		}
	}

	return r0
}

// MockProviderExchangeFiltersInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockProviderExchangeFiltersInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.ListExchangeFiltersRequest
func (_e *MockProviderExchangeFiltersInterface_Expecter) List(ctx interface{}, request interface{}) *MockProviderExchangeFiltersInterface_List_Call {
	return &MockProviderExchangeFiltersInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockProviderExchangeFiltersInterface_List_Call) Run(run func(ctx context.Context, request marketplace.ListExchangeFiltersRequest)) *MockProviderExchangeFiltersInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.ListExchangeFiltersRequest))
	})
	return _c
}

func (_c *MockProviderExchangeFiltersInterface_List_Call) Return(_a0 listing.Iterator[marketplace.ExchangeFilter]) *MockProviderExchangeFiltersInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProviderExchangeFiltersInterface_List_Call) RunAndReturn(run func(context.Context, marketplace.ListExchangeFiltersRequest) listing.Iterator[marketplace.ExchangeFilter]) *MockProviderExchangeFiltersInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockProviderExchangeFiltersInterface) ListAll(ctx context.Context, request marketplace.ListExchangeFiltersRequest) ([]marketplace.ExchangeFilter, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []marketplace.ExchangeFilter
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.ListExchangeFiltersRequest) ([]marketplace.ExchangeFilter, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.ListExchangeFiltersRequest) []marketplace.ExchangeFilter); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]marketplace.ExchangeFilter)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, marketplace.ListExchangeFiltersRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderExchangeFiltersInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockProviderExchangeFiltersInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.ListExchangeFiltersRequest
func (_e *MockProviderExchangeFiltersInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockProviderExchangeFiltersInterface_ListAll_Call {
	return &MockProviderExchangeFiltersInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockProviderExchangeFiltersInterface_ListAll_Call) Run(run func(ctx context.Context, request marketplace.ListExchangeFiltersRequest)) *MockProviderExchangeFiltersInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.ListExchangeFiltersRequest))
	})
	return _c
}

func (_c *MockProviderExchangeFiltersInterface_ListAll_Call) Return(_a0 []marketplace.ExchangeFilter, _a1 error) *MockProviderExchangeFiltersInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderExchangeFiltersInterface_ListAll_Call) RunAndReturn(run func(context.Context, marketplace.ListExchangeFiltersRequest) ([]marketplace.ExchangeFilter, error)) *MockProviderExchangeFiltersInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockProviderExchangeFiltersInterface) Update(ctx context.Context, request marketplace.UpdateExchangeFilterRequest) (*marketplace.UpdateExchangeFilterResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *marketplace.UpdateExchangeFilterResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.UpdateExchangeFilterRequest) (*marketplace.UpdateExchangeFilterResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.UpdateExchangeFilterRequest) *marketplace.UpdateExchangeFilterResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplace.UpdateExchangeFilterResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, marketplace.UpdateExchangeFilterRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderExchangeFiltersInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockProviderExchangeFiltersInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.UpdateExchangeFilterRequest
func (_e *MockProviderExchangeFiltersInterface_Expecter) Update(ctx interface{}, request interface{}) *MockProviderExchangeFiltersInterface_Update_Call {
	return &MockProviderExchangeFiltersInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockProviderExchangeFiltersInterface_Update_Call) Run(run func(ctx context.Context, request marketplace.UpdateExchangeFilterRequest)) *MockProviderExchangeFiltersInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.UpdateExchangeFilterRequest))
	})
	return _c
}

func (_c *MockProviderExchangeFiltersInterface_Update_Call) Return(_a0 *marketplace.UpdateExchangeFilterResponse, _a1 error) *MockProviderExchangeFiltersInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderExchangeFiltersInterface_Update_Call) RunAndReturn(run func(context.Context, marketplace.UpdateExchangeFilterRequest) (*marketplace.UpdateExchangeFilterResponse, error)) *MockProviderExchangeFiltersInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockProviderExchangeFiltersInterface) WithImpl(impl marketplace.ProviderExchangeFiltersService) marketplace.ProviderExchangeFiltersInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 marketplace.ProviderExchangeFiltersInterface
	if rf, ok := ret.Get(0).(func(marketplace.ProviderExchangeFiltersService) marketplace.ProviderExchangeFiltersInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(marketplace.ProviderExchangeFiltersInterface)
		}
	}

	return r0
}

// MockProviderExchangeFiltersInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockProviderExchangeFiltersInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl marketplace.ProviderExchangeFiltersService
func (_e *MockProviderExchangeFiltersInterface_Expecter) WithImpl(impl interface{}) *MockProviderExchangeFiltersInterface_WithImpl_Call {
	return &MockProviderExchangeFiltersInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockProviderExchangeFiltersInterface_WithImpl_Call) Run(run func(impl marketplace.ProviderExchangeFiltersService)) *MockProviderExchangeFiltersInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(marketplace.ProviderExchangeFiltersService))
	})
	return _c
}

func (_c *MockProviderExchangeFiltersInterface_WithImpl_Call) Return(_a0 marketplace.ProviderExchangeFiltersInterface) *MockProviderExchangeFiltersInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProviderExchangeFiltersInterface_WithImpl_Call) RunAndReturn(run func(marketplace.ProviderExchangeFiltersService) marketplace.ProviderExchangeFiltersInterface) *MockProviderExchangeFiltersInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockProviderExchangeFiltersInterface creates a new instance of MockProviderExchangeFiltersInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockProviderExchangeFiltersInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockProviderExchangeFiltersInterface {
	mock := &MockProviderExchangeFiltersInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
