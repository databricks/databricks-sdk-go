// Code generated by mockery v2.43.0. DO NOT EDIT.

package marketplace

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	marketplace "github.com/databricks/databricks-sdk-go/service/marketplace"

	mock "github.com/stretchr/testify/mock"
)

// MockProviderProvidersInterface is an autogenerated mock type for the ProviderProvidersInterface type
type MockProviderProvidersInterface struct {
	mock.Mock
}

type MockProviderProvidersInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockProviderProvidersInterface) EXPECT() *MockProviderProvidersInterface_Expecter {
	return &MockProviderProvidersInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockProviderProvidersInterface) Create(ctx context.Context, request marketplace.CreateProviderRequest) (*marketplace.CreateProviderResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *marketplace.CreateProviderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.CreateProviderRequest) (*marketplace.CreateProviderResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.CreateProviderRequest) *marketplace.CreateProviderResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplace.CreateProviderResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, marketplace.CreateProviderRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderProvidersInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockProviderProvidersInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.CreateProviderRequest
func (_e *MockProviderProvidersInterface_Expecter) Create(ctx interface{}, request interface{}) *MockProviderProvidersInterface_Create_Call {
	return &MockProviderProvidersInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockProviderProvidersInterface_Create_Call) Run(run func(ctx context.Context, request marketplace.CreateProviderRequest)) *MockProviderProvidersInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.CreateProviderRequest))
	})
	return _c
}

func (_c *MockProviderProvidersInterface_Create_Call) Return(_a0 *marketplace.CreateProviderResponse, _a1 error) *MockProviderProvidersInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderProvidersInterface_Create_Call) RunAndReturn(run func(context.Context, marketplace.CreateProviderRequest) (*marketplace.CreateProviderResponse, error)) *MockProviderProvidersInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockProviderProvidersInterface) Delete(ctx context.Context, request marketplace.DeleteProviderRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.DeleteProviderRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockProviderProvidersInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockProviderProvidersInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.DeleteProviderRequest
func (_e *MockProviderProvidersInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockProviderProvidersInterface_Delete_Call {
	return &MockProviderProvidersInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockProviderProvidersInterface_Delete_Call) Run(run func(ctx context.Context, request marketplace.DeleteProviderRequest)) *MockProviderProvidersInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.DeleteProviderRequest))
	})
	return _c
}

func (_c *MockProviderProvidersInterface_Delete_Call) Return(_a0 error) *MockProviderProvidersInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProviderProvidersInterface_Delete_Call) RunAndReturn(run func(context.Context, marketplace.DeleteProviderRequest) error) *MockProviderProvidersInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *MockProviderProvidersInterface) DeleteById(ctx context.Context, id string) error {
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

// MockProviderProvidersInterface_DeleteById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteById'
type MockProviderProvidersInterface_DeleteById_Call struct {
	*mock.Call
}

// DeleteById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockProviderProvidersInterface_Expecter) DeleteById(ctx interface{}, id interface{}) *MockProviderProvidersInterface_DeleteById_Call {
	return &MockProviderProvidersInterface_DeleteById_Call{Call: _e.mock.On("DeleteById", ctx, id)}
}

func (_c *MockProviderProvidersInterface_DeleteById_Call) Run(run func(ctx context.Context, id string)) *MockProviderProvidersInterface_DeleteById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockProviderProvidersInterface_DeleteById_Call) Return(_a0 error) *MockProviderProvidersInterface_DeleteById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProviderProvidersInterface_DeleteById_Call) RunAndReturn(run func(context.Context, string) error) *MockProviderProvidersInterface_DeleteById_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockProviderProvidersInterface) Get(ctx context.Context, request marketplace.GetProviderRequest) (*marketplace.GetProviderResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *marketplace.GetProviderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.GetProviderRequest) (*marketplace.GetProviderResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.GetProviderRequest) *marketplace.GetProviderResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplace.GetProviderResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, marketplace.GetProviderRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderProvidersInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockProviderProvidersInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.GetProviderRequest
func (_e *MockProviderProvidersInterface_Expecter) Get(ctx interface{}, request interface{}) *MockProviderProvidersInterface_Get_Call {
	return &MockProviderProvidersInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockProviderProvidersInterface_Get_Call) Run(run func(ctx context.Context, request marketplace.GetProviderRequest)) *MockProviderProvidersInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.GetProviderRequest))
	})
	return _c
}

func (_c *MockProviderProvidersInterface_Get_Call) Return(_a0 *marketplace.GetProviderResponse, _a1 error) *MockProviderProvidersInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderProvidersInterface_Get_Call) RunAndReturn(run func(context.Context, marketplace.GetProviderRequest) (*marketplace.GetProviderResponse, error)) *MockProviderProvidersInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: ctx, id
func (_m *MockProviderProvidersInterface) GetById(ctx context.Context, id string) (*marketplace.GetProviderResponse, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *marketplace.GetProviderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*marketplace.GetProviderResponse, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *marketplace.GetProviderResponse); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplace.GetProviderResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderProvidersInterface_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type MockProviderProvidersInterface_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockProviderProvidersInterface_Expecter) GetById(ctx interface{}, id interface{}) *MockProviderProvidersInterface_GetById_Call {
	return &MockProviderProvidersInterface_GetById_Call{Call: _e.mock.On("GetById", ctx, id)}
}

func (_c *MockProviderProvidersInterface_GetById_Call) Run(run func(ctx context.Context, id string)) *MockProviderProvidersInterface_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockProviderProvidersInterface_GetById_Call) Return(_a0 *marketplace.GetProviderResponse, _a1 error) *MockProviderProvidersInterface_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderProvidersInterface_GetById_Call) RunAndReturn(run func(context.Context, string) (*marketplace.GetProviderResponse, error)) *MockProviderProvidersInterface_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *MockProviderProvidersInterface) GetByName(ctx context.Context, name string) (*marketplace.ProviderInfo, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *marketplace.ProviderInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*marketplace.ProviderInfo, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *marketplace.ProviderInfo); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplace.ProviderInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderProvidersInterface_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type MockProviderProvidersInterface_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockProviderProvidersInterface_Expecter) GetByName(ctx interface{}, name interface{}) *MockProviderProvidersInterface_GetByName_Call {
	return &MockProviderProvidersInterface_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *MockProviderProvidersInterface_GetByName_Call) Run(run func(ctx context.Context, name string)) *MockProviderProvidersInterface_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockProviderProvidersInterface_GetByName_Call) Return(_a0 *marketplace.ProviderInfo, _a1 error) *MockProviderProvidersInterface_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderProvidersInterface_GetByName_Call) RunAndReturn(run func(context.Context, string) (*marketplace.ProviderInfo, error)) *MockProviderProvidersInterface_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockProviderProvidersInterface) Impl() marketplace.ProviderProvidersService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 marketplace.ProviderProvidersService
	if rf, ok := ret.Get(0).(func() marketplace.ProviderProvidersService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(marketplace.ProviderProvidersService)
		}
	}

	return r0
}

// MockProviderProvidersInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockProviderProvidersInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockProviderProvidersInterface_Expecter) Impl() *MockProviderProvidersInterface_Impl_Call {
	return &MockProviderProvidersInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockProviderProvidersInterface_Impl_Call) Run(run func()) *MockProviderProvidersInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProviderProvidersInterface_Impl_Call) Return(_a0 marketplace.ProviderProvidersService) *MockProviderProvidersInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProviderProvidersInterface_Impl_Call) RunAndReturn(run func() marketplace.ProviderProvidersService) *MockProviderProvidersInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockProviderProvidersInterface) List(ctx context.Context, request marketplace.ListProvidersRequest) listing.Iterator[marketplace.ProviderInfo] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[marketplace.ProviderInfo]
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.ListProvidersRequest) listing.Iterator[marketplace.ProviderInfo]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[marketplace.ProviderInfo])
		}
	}

	return r0
}

// MockProviderProvidersInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockProviderProvidersInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.ListProvidersRequest
func (_e *MockProviderProvidersInterface_Expecter) List(ctx interface{}, request interface{}) *MockProviderProvidersInterface_List_Call {
	return &MockProviderProvidersInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockProviderProvidersInterface_List_Call) Run(run func(ctx context.Context, request marketplace.ListProvidersRequest)) *MockProviderProvidersInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.ListProvidersRequest))
	})
	return _c
}

func (_c *MockProviderProvidersInterface_List_Call) Return(_a0 listing.Iterator[marketplace.ProviderInfo]) *MockProviderProvidersInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProviderProvidersInterface_List_Call) RunAndReturn(run func(context.Context, marketplace.ListProvidersRequest) listing.Iterator[marketplace.ProviderInfo]) *MockProviderProvidersInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockProviderProvidersInterface) ListAll(ctx context.Context, request marketplace.ListProvidersRequest) ([]marketplace.ProviderInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []marketplace.ProviderInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.ListProvidersRequest) ([]marketplace.ProviderInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.ListProvidersRequest) []marketplace.ProviderInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]marketplace.ProviderInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, marketplace.ListProvidersRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderProvidersInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockProviderProvidersInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.ListProvidersRequest
func (_e *MockProviderProvidersInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockProviderProvidersInterface_ListAll_Call {
	return &MockProviderProvidersInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockProviderProvidersInterface_ListAll_Call) Run(run func(ctx context.Context, request marketplace.ListProvidersRequest)) *MockProviderProvidersInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.ListProvidersRequest))
	})
	return _c
}

func (_c *MockProviderProvidersInterface_ListAll_Call) Return(_a0 []marketplace.ProviderInfo, _a1 error) *MockProviderProvidersInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderProvidersInterface_ListAll_Call) RunAndReturn(run func(context.Context, marketplace.ListProvidersRequest) ([]marketplace.ProviderInfo, error)) *MockProviderProvidersInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// ProviderInfoNameToIdMap provides a mock function with given fields: ctx, request
func (_m *MockProviderProvidersInterface) ProviderInfoNameToIdMap(ctx context.Context, request marketplace.ListProvidersRequest) (map[string]string, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ProviderInfoNameToIdMap")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.ListProvidersRequest) (map[string]string, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.ListProvidersRequest) map[string]string); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, marketplace.ListProvidersRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderProvidersInterface_ProviderInfoNameToIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProviderInfoNameToIdMap'
type MockProviderProvidersInterface_ProviderInfoNameToIdMap_Call struct {
	*mock.Call
}

// ProviderInfoNameToIdMap is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.ListProvidersRequest
func (_e *MockProviderProvidersInterface_Expecter) ProviderInfoNameToIdMap(ctx interface{}, request interface{}) *MockProviderProvidersInterface_ProviderInfoNameToIdMap_Call {
	return &MockProviderProvidersInterface_ProviderInfoNameToIdMap_Call{Call: _e.mock.On("ProviderInfoNameToIdMap", ctx, request)}
}

func (_c *MockProviderProvidersInterface_ProviderInfoNameToIdMap_Call) Run(run func(ctx context.Context, request marketplace.ListProvidersRequest)) *MockProviderProvidersInterface_ProviderInfoNameToIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.ListProvidersRequest))
	})
	return _c
}

func (_c *MockProviderProvidersInterface_ProviderInfoNameToIdMap_Call) Return(_a0 map[string]string, _a1 error) *MockProviderProvidersInterface_ProviderInfoNameToIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderProvidersInterface_ProviderInfoNameToIdMap_Call) RunAndReturn(run func(context.Context, marketplace.ListProvidersRequest) (map[string]string, error)) *MockProviderProvidersInterface_ProviderInfoNameToIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockProviderProvidersInterface) Update(ctx context.Context, request marketplace.UpdateProviderRequest) (*marketplace.UpdateProviderResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *marketplace.UpdateProviderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.UpdateProviderRequest) (*marketplace.UpdateProviderResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.UpdateProviderRequest) *marketplace.UpdateProviderResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplace.UpdateProviderResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, marketplace.UpdateProviderRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderProvidersInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockProviderProvidersInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.UpdateProviderRequest
func (_e *MockProviderProvidersInterface_Expecter) Update(ctx interface{}, request interface{}) *MockProviderProvidersInterface_Update_Call {
	return &MockProviderProvidersInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockProviderProvidersInterface_Update_Call) Run(run func(ctx context.Context, request marketplace.UpdateProviderRequest)) *MockProviderProvidersInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.UpdateProviderRequest))
	})
	return _c
}

func (_c *MockProviderProvidersInterface_Update_Call) Return(_a0 *marketplace.UpdateProviderResponse, _a1 error) *MockProviderProvidersInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderProvidersInterface_Update_Call) RunAndReturn(run func(context.Context, marketplace.UpdateProviderRequest) (*marketplace.UpdateProviderResponse, error)) *MockProviderProvidersInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockProviderProvidersInterface) WithImpl(impl marketplace.ProviderProvidersService) marketplace.ProviderProvidersInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 marketplace.ProviderProvidersInterface
	if rf, ok := ret.Get(0).(func(marketplace.ProviderProvidersService) marketplace.ProviderProvidersInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(marketplace.ProviderProvidersInterface)
		}
	}

	return r0
}

// MockProviderProvidersInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockProviderProvidersInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl marketplace.ProviderProvidersService
func (_e *MockProviderProvidersInterface_Expecter) WithImpl(impl interface{}) *MockProviderProvidersInterface_WithImpl_Call {
	return &MockProviderProvidersInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockProviderProvidersInterface_WithImpl_Call) Run(run func(impl marketplace.ProviderProvidersService)) *MockProviderProvidersInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(marketplace.ProviderProvidersService))
	})
	return _c
}

func (_c *MockProviderProvidersInterface_WithImpl_Call) Return(_a0 marketplace.ProviderProvidersInterface) *MockProviderProvidersInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProviderProvidersInterface_WithImpl_Call) RunAndReturn(run func(marketplace.ProviderProvidersService) marketplace.ProviderProvidersInterface) *MockProviderProvidersInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockProviderProvidersInterface creates a new instance of MockProviderProvidersInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockProviderProvidersInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockProviderProvidersInterface {
	mock := &MockProviderProvidersInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
