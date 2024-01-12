// Code generated by mockery v2.39.1. DO NOT EDIT.

package catalog

import (
	context "context"

	catalog "github.com/databricks/databricks-sdk-go/service/catalog"

	listing "github.com/databricks/databricks-sdk-go/listing"

	mock "github.com/stretchr/testify/mock"
)

// MockRegisteredModelsInterface is an autogenerated mock type for the RegisteredModelsInterface type
type MockRegisteredModelsInterface struct {
	mock.Mock
}

type MockRegisteredModelsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRegisteredModelsInterface) EXPECT() *MockRegisteredModelsInterface_Expecter {
	return &MockRegisteredModelsInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockRegisteredModelsInterface) Create(ctx context.Context, request catalog.CreateRegisteredModelRequest) (*catalog.RegisteredModelInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *catalog.RegisteredModelInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateRegisteredModelRequest) (*catalog.RegisteredModelInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateRegisteredModelRequest) *catalog.RegisteredModelInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.RegisteredModelInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.CreateRegisteredModelRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRegisteredModelsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockRegisteredModelsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.CreateRegisteredModelRequest
func (_e *MockRegisteredModelsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockRegisteredModelsInterface_Create_Call {
	return &MockRegisteredModelsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockRegisteredModelsInterface_Create_Call) Run(run func(ctx context.Context, request catalog.CreateRegisteredModelRequest)) *MockRegisteredModelsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.CreateRegisteredModelRequest))
	})
	return _c
}

func (_c *MockRegisteredModelsInterface_Create_Call) Return(_a0 *catalog.RegisteredModelInfo, _a1 error) *MockRegisteredModelsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRegisteredModelsInterface_Create_Call) RunAndReturn(run func(context.Context, catalog.CreateRegisteredModelRequest) (*catalog.RegisteredModelInfo, error)) *MockRegisteredModelsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockRegisteredModelsInterface) Delete(ctx context.Context, request catalog.DeleteRegisteredModelRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.DeleteRegisteredModelRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRegisteredModelsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockRegisteredModelsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.DeleteRegisteredModelRequest
func (_e *MockRegisteredModelsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockRegisteredModelsInterface_Delete_Call {
	return &MockRegisteredModelsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockRegisteredModelsInterface_Delete_Call) Run(run func(ctx context.Context, request catalog.DeleteRegisteredModelRequest)) *MockRegisteredModelsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.DeleteRegisteredModelRequest))
	})
	return _c
}

func (_c *MockRegisteredModelsInterface_Delete_Call) Return(_a0 error) *MockRegisteredModelsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRegisteredModelsInterface_Delete_Call) RunAndReturn(run func(context.Context, catalog.DeleteRegisteredModelRequest) error) *MockRegisteredModelsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAlias provides a mock function with given fields: ctx, request
func (_m *MockRegisteredModelsInterface) DeleteAlias(ctx context.Context, request catalog.DeleteAliasRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAlias")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.DeleteAliasRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRegisteredModelsInterface_DeleteAlias_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAlias'
type MockRegisteredModelsInterface_DeleteAlias_Call struct {
	*mock.Call
}

// DeleteAlias is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.DeleteAliasRequest
func (_e *MockRegisteredModelsInterface_Expecter) DeleteAlias(ctx interface{}, request interface{}) *MockRegisteredModelsInterface_DeleteAlias_Call {
	return &MockRegisteredModelsInterface_DeleteAlias_Call{Call: _e.mock.On("DeleteAlias", ctx, request)}
}

func (_c *MockRegisteredModelsInterface_DeleteAlias_Call) Run(run func(ctx context.Context, request catalog.DeleteAliasRequest)) *MockRegisteredModelsInterface_DeleteAlias_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.DeleteAliasRequest))
	})
	return _c
}

func (_c *MockRegisteredModelsInterface_DeleteAlias_Call) Return(_a0 error) *MockRegisteredModelsInterface_DeleteAlias_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRegisteredModelsInterface_DeleteAlias_Call) RunAndReturn(run func(context.Context, catalog.DeleteAliasRequest) error) *MockRegisteredModelsInterface_DeleteAlias_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAliasByFullNameAndAlias provides a mock function with given fields: ctx, fullName, alias
func (_m *MockRegisteredModelsInterface) DeleteAliasByFullNameAndAlias(ctx context.Context, fullName string, alias string) error {
	ret := _m.Called(ctx, fullName, alias)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAliasByFullNameAndAlias")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, fullName, alias)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRegisteredModelsInterface_DeleteAliasByFullNameAndAlias_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAliasByFullNameAndAlias'
type MockRegisteredModelsInterface_DeleteAliasByFullNameAndAlias_Call struct {
	*mock.Call
}

// DeleteAliasByFullNameAndAlias is a helper method to define mock.On call
//   - ctx context.Context
//   - fullName string
//   - alias string
func (_e *MockRegisteredModelsInterface_Expecter) DeleteAliasByFullNameAndAlias(ctx interface{}, fullName interface{}, alias interface{}) *MockRegisteredModelsInterface_DeleteAliasByFullNameAndAlias_Call {
	return &MockRegisteredModelsInterface_DeleteAliasByFullNameAndAlias_Call{Call: _e.mock.On("DeleteAliasByFullNameAndAlias", ctx, fullName, alias)}
}

func (_c *MockRegisteredModelsInterface_DeleteAliasByFullNameAndAlias_Call) Run(run func(ctx context.Context, fullName string, alias string)) *MockRegisteredModelsInterface_DeleteAliasByFullNameAndAlias_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockRegisteredModelsInterface_DeleteAliasByFullNameAndAlias_Call) Return(_a0 error) *MockRegisteredModelsInterface_DeleteAliasByFullNameAndAlias_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRegisteredModelsInterface_DeleteAliasByFullNameAndAlias_Call) RunAndReturn(run func(context.Context, string, string) error) *MockRegisteredModelsInterface_DeleteAliasByFullNameAndAlias_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByFullName provides a mock function with given fields: ctx, fullName
func (_m *MockRegisteredModelsInterface) DeleteByFullName(ctx context.Context, fullName string) error {
	ret := _m.Called(ctx, fullName)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByFullName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, fullName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRegisteredModelsInterface_DeleteByFullName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByFullName'
type MockRegisteredModelsInterface_DeleteByFullName_Call struct {
	*mock.Call
}

// DeleteByFullName is a helper method to define mock.On call
//   - ctx context.Context
//   - fullName string
func (_e *MockRegisteredModelsInterface_Expecter) DeleteByFullName(ctx interface{}, fullName interface{}) *MockRegisteredModelsInterface_DeleteByFullName_Call {
	return &MockRegisteredModelsInterface_DeleteByFullName_Call{Call: _e.mock.On("DeleteByFullName", ctx, fullName)}
}

func (_c *MockRegisteredModelsInterface_DeleteByFullName_Call) Run(run func(ctx context.Context, fullName string)) *MockRegisteredModelsInterface_DeleteByFullName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRegisteredModelsInterface_DeleteByFullName_Call) Return(_a0 error) *MockRegisteredModelsInterface_DeleteByFullName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRegisteredModelsInterface_DeleteByFullName_Call) RunAndReturn(run func(context.Context, string) error) *MockRegisteredModelsInterface_DeleteByFullName_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockRegisteredModelsInterface) Get(ctx context.Context, request catalog.GetRegisteredModelRequest) (*catalog.RegisteredModelInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *catalog.RegisteredModelInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetRegisteredModelRequest) (*catalog.RegisteredModelInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetRegisteredModelRequest) *catalog.RegisteredModelInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.RegisteredModelInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.GetRegisteredModelRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRegisteredModelsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockRegisteredModelsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.GetRegisteredModelRequest
func (_e *MockRegisteredModelsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockRegisteredModelsInterface_Get_Call {
	return &MockRegisteredModelsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockRegisteredModelsInterface_Get_Call) Run(run func(ctx context.Context, request catalog.GetRegisteredModelRequest)) *MockRegisteredModelsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.GetRegisteredModelRequest))
	})
	return _c
}

func (_c *MockRegisteredModelsInterface_Get_Call) Return(_a0 *catalog.RegisteredModelInfo, _a1 error) *MockRegisteredModelsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRegisteredModelsInterface_Get_Call) RunAndReturn(run func(context.Context, catalog.GetRegisteredModelRequest) (*catalog.RegisteredModelInfo, error)) *MockRegisteredModelsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByFullName provides a mock function with given fields: ctx, fullName
func (_m *MockRegisteredModelsInterface) GetByFullName(ctx context.Context, fullName string) (*catalog.RegisteredModelInfo, error) {
	ret := _m.Called(ctx, fullName)

	if len(ret) == 0 {
		panic("no return value specified for GetByFullName")
	}

	var r0 *catalog.RegisteredModelInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*catalog.RegisteredModelInfo, error)); ok {
		return rf(ctx, fullName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *catalog.RegisteredModelInfo); ok {
		r0 = rf(ctx, fullName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.RegisteredModelInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, fullName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRegisteredModelsInterface_GetByFullName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByFullName'
type MockRegisteredModelsInterface_GetByFullName_Call struct {
	*mock.Call
}

// GetByFullName is a helper method to define mock.On call
//   - ctx context.Context
//   - fullName string
func (_e *MockRegisteredModelsInterface_Expecter) GetByFullName(ctx interface{}, fullName interface{}) *MockRegisteredModelsInterface_GetByFullName_Call {
	return &MockRegisteredModelsInterface_GetByFullName_Call{Call: _e.mock.On("GetByFullName", ctx, fullName)}
}

func (_c *MockRegisteredModelsInterface_GetByFullName_Call) Run(run func(ctx context.Context, fullName string)) *MockRegisteredModelsInterface_GetByFullName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRegisteredModelsInterface_GetByFullName_Call) Return(_a0 *catalog.RegisteredModelInfo, _a1 error) *MockRegisteredModelsInterface_GetByFullName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRegisteredModelsInterface_GetByFullName_Call) RunAndReturn(run func(context.Context, string) (*catalog.RegisteredModelInfo, error)) *MockRegisteredModelsInterface_GetByFullName_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *MockRegisteredModelsInterface) GetByName(ctx context.Context, name string) (*catalog.RegisteredModelInfo, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *catalog.RegisteredModelInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*catalog.RegisteredModelInfo, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *catalog.RegisteredModelInfo); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.RegisteredModelInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRegisteredModelsInterface_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type MockRegisteredModelsInterface_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockRegisteredModelsInterface_Expecter) GetByName(ctx interface{}, name interface{}) *MockRegisteredModelsInterface_GetByName_Call {
	return &MockRegisteredModelsInterface_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *MockRegisteredModelsInterface_GetByName_Call) Run(run func(ctx context.Context, name string)) *MockRegisteredModelsInterface_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRegisteredModelsInterface_GetByName_Call) Return(_a0 *catalog.RegisteredModelInfo, _a1 error) *MockRegisteredModelsInterface_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRegisteredModelsInterface_GetByName_Call) RunAndReturn(run func(context.Context, string) (*catalog.RegisteredModelInfo, error)) *MockRegisteredModelsInterface_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockRegisteredModelsInterface) Impl() catalog.RegisteredModelsService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 catalog.RegisteredModelsService
	if rf, ok := ret.Get(0).(func() catalog.RegisteredModelsService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.RegisteredModelsService)
		}
	}

	return r0
}

// MockRegisteredModelsInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockRegisteredModelsInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockRegisteredModelsInterface_Expecter) Impl() *MockRegisteredModelsInterface_Impl_Call {
	return &MockRegisteredModelsInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockRegisteredModelsInterface_Impl_Call) Run(run func()) *MockRegisteredModelsInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRegisteredModelsInterface_Impl_Call) Return(_a0 catalog.RegisteredModelsService) *MockRegisteredModelsInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRegisteredModelsInterface_Impl_Call) RunAndReturn(run func() catalog.RegisteredModelsService) *MockRegisteredModelsInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockRegisteredModelsInterface) List(ctx context.Context, request catalog.ListRegisteredModelsRequest) listing.Iterator[catalog.RegisteredModelInfo] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[catalog.RegisteredModelInfo]
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListRegisteredModelsRequest) listing.Iterator[catalog.RegisteredModelInfo]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[catalog.RegisteredModelInfo])
		}
	}

	return r0
}

// MockRegisteredModelsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockRegisteredModelsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListRegisteredModelsRequest
func (_e *MockRegisteredModelsInterface_Expecter) List(ctx interface{}, request interface{}) *MockRegisteredModelsInterface_List_Call {
	return &MockRegisteredModelsInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockRegisteredModelsInterface_List_Call) Run(run func(ctx context.Context, request catalog.ListRegisteredModelsRequest)) *MockRegisteredModelsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListRegisteredModelsRequest))
	})
	return _c
}

func (_c *MockRegisteredModelsInterface_List_Call) Return(_a0 listing.Iterator[catalog.RegisteredModelInfo]) *MockRegisteredModelsInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRegisteredModelsInterface_List_Call) RunAndReturn(run func(context.Context, catalog.ListRegisteredModelsRequest) listing.Iterator[catalog.RegisteredModelInfo]) *MockRegisteredModelsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockRegisteredModelsInterface) ListAll(ctx context.Context, request catalog.ListRegisteredModelsRequest) ([]catalog.RegisteredModelInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []catalog.RegisteredModelInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListRegisteredModelsRequest) ([]catalog.RegisteredModelInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListRegisteredModelsRequest) []catalog.RegisteredModelInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]catalog.RegisteredModelInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.ListRegisteredModelsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRegisteredModelsInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockRegisteredModelsInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListRegisteredModelsRequest
func (_e *MockRegisteredModelsInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockRegisteredModelsInterface_ListAll_Call {
	return &MockRegisteredModelsInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockRegisteredModelsInterface_ListAll_Call) Run(run func(ctx context.Context, request catalog.ListRegisteredModelsRequest)) *MockRegisteredModelsInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListRegisteredModelsRequest))
	})
	return _c
}

func (_c *MockRegisteredModelsInterface_ListAll_Call) Return(_a0 []catalog.RegisteredModelInfo, _a1 error) *MockRegisteredModelsInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRegisteredModelsInterface_ListAll_Call) RunAndReturn(run func(context.Context, catalog.ListRegisteredModelsRequest) ([]catalog.RegisteredModelInfo, error)) *MockRegisteredModelsInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// RegisteredModelInfoNameToFullNameMap provides a mock function with given fields: ctx, request
func (_m *MockRegisteredModelsInterface) RegisteredModelInfoNameToFullNameMap(ctx context.Context, request catalog.ListRegisteredModelsRequest) (map[string]string, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for RegisteredModelInfoNameToFullNameMap")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListRegisteredModelsRequest) (map[string]string, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListRegisteredModelsRequest) map[string]string); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.ListRegisteredModelsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRegisteredModelsInterface_RegisteredModelInfoNameToFullNameMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisteredModelInfoNameToFullNameMap'
type MockRegisteredModelsInterface_RegisteredModelInfoNameToFullNameMap_Call struct {
	*mock.Call
}

// RegisteredModelInfoNameToFullNameMap is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListRegisteredModelsRequest
func (_e *MockRegisteredModelsInterface_Expecter) RegisteredModelInfoNameToFullNameMap(ctx interface{}, request interface{}) *MockRegisteredModelsInterface_RegisteredModelInfoNameToFullNameMap_Call {
	return &MockRegisteredModelsInterface_RegisteredModelInfoNameToFullNameMap_Call{Call: _e.mock.On("RegisteredModelInfoNameToFullNameMap", ctx, request)}
}

func (_c *MockRegisteredModelsInterface_RegisteredModelInfoNameToFullNameMap_Call) Run(run func(ctx context.Context, request catalog.ListRegisteredModelsRequest)) *MockRegisteredModelsInterface_RegisteredModelInfoNameToFullNameMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListRegisteredModelsRequest))
	})
	return _c
}

func (_c *MockRegisteredModelsInterface_RegisteredModelInfoNameToFullNameMap_Call) Return(_a0 map[string]string, _a1 error) *MockRegisteredModelsInterface_RegisteredModelInfoNameToFullNameMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRegisteredModelsInterface_RegisteredModelInfoNameToFullNameMap_Call) RunAndReturn(run func(context.Context, catalog.ListRegisteredModelsRequest) (map[string]string, error)) *MockRegisteredModelsInterface_RegisteredModelInfoNameToFullNameMap_Call {
	_c.Call.Return(run)
	return _c
}

// SetAlias provides a mock function with given fields: ctx, request
func (_m *MockRegisteredModelsInterface) SetAlias(ctx context.Context, request catalog.SetRegisteredModelAliasRequest) (*catalog.RegisteredModelAlias, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for SetAlias")
	}

	var r0 *catalog.RegisteredModelAlias
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.SetRegisteredModelAliasRequest) (*catalog.RegisteredModelAlias, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.SetRegisteredModelAliasRequest) *catalog.RegisteredModelAlias); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.RegisteredModelAlias)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.SetRegisteredModelAliasRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRegisteredModelsInterface_SetAlias_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetAlias'
type MockRegisteredModelsInterface_SetAlias_Call struct {
	*mock.Call
}

// SetAlias is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.SetRegisteredModelAliasRequest
func (_e *MockRegisteredModelsInterface_Expecter) SetAlias(ctx interface{}, request interface{}) *MockRegisteredModelsInterface_SetAlias_Call {
	return &MockRegisteredModelsInterface_SetAlias_Call{Call: _e.mock.On("SetAlias", ctx, request)}
}

func (_c *MockRegisteredModelsInterface_SetAlias_Call) Run(run func(ctx context.Context, request catalog.SetRegisteredModelAliasRequest)) *MockRegisteredModelsInterface_SetAlias_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.SetRegisteredModelAliasRequest))
	})
	return _c
}

func (_c *MockRegisteredModelsInterface_SetAlias_Call) Return(_a0 *catalog.RegisteredModelAlias, _a1 error) *MockRegisteredModelsInterface_SetAlias_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRegisteredModelsInterface_SetAlias_Call) RunAndReturn(run func(context.Context, catalog.SetRegisteredModelAliasRequest) (*catalog.RegisteredModelAlias, error)) *MockRegisteredModelsInterface_SetAlias_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockRegisteredModelsInterface) Update(ctx context.Context, request catalog.UpdateRegisteredModelRequest) (*catalog.RegisteredModelInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *catalog.RegisteredModelInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdateRegisteredModelRequest) (*catalog.RegisteredModelInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdateRegisteredModelRequest) *catalog.RegisteredModelInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.RegisteredModelInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.UpdateRegisteredModelRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRegisteredModelsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockRegisteredModelsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.UpdateRegisteredModelRequest
func (_e *MockRegisteredModelsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockRegisteredModelsInterface_Update_Call {
	return &MockRegisteredModelsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockRegisteredModelsInterface_Update_Call) Run(run func(ctx context.Context, request catalog.UpdateRegisteredModelRequest)) *MockRegisteredModelsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.UpdateRegisteredModelRequest))
	})
	return _c
}

func (_c *MockRegisteredModelsInterface_Update_Call) Return(_a0 *catalog.RegisteredModelInfo, _a1 error) *MockRegisteredModelsInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRegisteredModelsInterface_Update_Call) RunAndReturn(run func(context.Context, catalog.UpdateRegisteredModelRequest) (*catalog.RegisteredModelInfo, error)) *MockRegisteredModelsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockRegisteredModelsInterface) WithImpl(impl catalog.RegisteredModelsService) catalog.RegisteredModelsInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 catalog.RegisteredModelsInterface
	if rf, ok := ret.Get(0).(func(catalog.RegisteredModelsService) catalog.RegisteredModelsInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.RegisteredModelsInterface)
		}
	}

	return r0
}

// MockRegisteredModelsInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockRegisteredModelsInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl catalog.RegisteredModelsService
func (_e *MockRegisteredModelsInterface_Expecter) WithImpl(impl interface{}) *MockRegisteredModelsInterface_WithImpl_Call {
	return &MockRegisteredModelsInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockRegisteredModelsInterface_WithImpl_Call) Run(run func(impl catalog.RegisteredModelsService)) *MockRegisteredModelsInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(catalog.RegisteredModelsService))
	})
	return _c
}

func (_c *MockRegisteredModelsInterface_WithImpl_Call) Return(_a0 catalog.RegisteredModelsInterface) *MockRegisteredModelsInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRegisteredModelsInterface_WithImpl_Call) RunAndReturn(run func(catalog.RegisteredModelsService) catalog.RegisteredModelsInterface) *MockRegisteredModelsInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRegisteredModelsInterface creates a new instance of MockRegisteredModelsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRegisteredModelsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRegisteredModelsInterface {
	mock := &MockRegisteredModelsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
