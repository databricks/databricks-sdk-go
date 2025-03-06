// Code generated by mockery v2.53.0. DO NOT EDIT.

package catalog

import (
	context "context"

	catalog "github.com/databricks/databricks-sdk-go/service/catalog"

	listing "github.com/databricks/databricks-sdk-go/listing"

	mock "github.com/stretchr/testify/mock"
)

// MockFunctionsInterface is an autogenerated mock type for the FunctionsInterface type
type MockFunctionsInterface struct {
	mock.Mock
}

type MockFunctionsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFunctionsInterface) EXPECT() *MockFunctionsInterface_Expecter {
	return &MockFunctionsInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockFunctionsInterface) Create(ctx context.Context, request catalog.CreateFunctionRequest) (*catalog.FunctionInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *catalog.FunctionInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateFunctionRequest) (*catalog.FunctionInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateFunctionRequest) *catalog.FunctionInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.FunctionInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.CreateFunctionRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFunctionsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockFunctionsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.CreateFunctionRequest
func (_e *MockFunctionsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockFunctionsInterface_Create_Call {
	return &MockFunctionsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockFunctionsInterface_Create_Call) Run(run func(ctx context.Context, request catalog.CreateFunctionRequest)) *MockFunctionsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.CreateFunctionRequest))
	})
	return _c
}

func (_c *MockFunctionsInterface_Create_Call) Return(_a0 *catalog.FunctionInfo, _a1 error) *MockFunctionsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFunctionsInterface_Create_Call) RunAndReturn(run func(context.Context, catalog.CreateFunctionRequest) (*catalog.FunctionInfo, error)) *MockFunctionsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockFunctionsInterface) Delete(ctx context.Context, request catalog.DeleteFunctionRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.DeleteFunctionRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFunctionsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockFunctionsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.DeleteFunctionRequest
func (_e *MockFunctionsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockFunctionsInterface_Delete_Call {
	return &MockFunctionsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockFunctionsInterface_Delete_Call) Run(run func(ctx context.Context, request catalog.DeleteFunctionRequest)) *MockFunctionsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.DeleteFunctionRequest))
	})
	return _c
}

func (_c *MockFunctionsInterface_Delete_Call) Return(_a0 error) *MockFunctionsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFunctionsInterface_Delete_Call) RunAndReturn(run func(context.Context, catalog.DeleteFunctionRequest) error) *MockFunctionsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByName provides a mock function with given fields: ctx, name
func (_m *MockFunctionsInterface) DeleteByName(ctx context.Context, name string) error {
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

// MockFunctionsInterface_DeleteByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByName'
type MockFunctionsInterface_DeleteByName_Call struct {
	*mock.Call
}

// DeleteByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockFunctionsInterface_Expecter) DeleteByName(ctx interface{}, name interface{}) *MockFunctionsInterface_DeleteByName_Call {
	return &MockFunctionsInterface_DeleteByName_Call{Call: _e.mock.On("DeleteByName", ctx, name)}
}

func (_c *MockFunctionsInterface_DeleteByName_Call) Run(run func(ctx context.Context, name string)) *MockFunctionsInterface_DeleteByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockFunctionsInterface_DeleteByName_Call) Return(_a0 error) *MockFunctionsInterface_DeleteByName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFunctionsInterface_DeleteByName_Call) RunAndReturn(run func(context.Context, string) error) *MockFunctionsInterface_DeleteByName_Call {
	_c.Call.Return(run)
	return _c
}

// FunctionInfoNameToFullNameMap provides a mock function with given fields: ctx, request
func (_m *MockFunctionsInterface) FunctionInfoNameToFullNameMap(ctx context.Context, request catalog.ListFunctionsRequest) (map[string]string, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for FunctionInfoNameToFullNameMap")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListFunctionsRequest) (map[string]string, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListFunctionsRequest) map[string]string); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.ListFunctionsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFunctionsInterface_FunctionInfoNameToFullNameMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FunctionInfoNameToFullNameMap'
type MockFunctionsInterface_FunctionInfoNameToFullNameMap_Call struct {
	*mock.Call
}

// FunctionInfoNameToFullNameMap is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListFunctionsRequest
func (_e *MockFunctionsInterface_Expecter) FunctionInfoNameToFullNameMap(ctx interface{}, request interface{}) *MockFunctionsInterface_FunctionInfoNameToFullNameMap_Call {
	return &MockFunctionsInterface_FunctionInfoNameToFullNameMap_Call{Call: _e.mock.On("FunctionInfoNameToFullNameMap", ctx, request)}
}

func (_c *MockFunctionsInterface_FunctionInfoNameToFullNameMap_Call) Run(run func(ctx context.Context, request catalog.ListFunctionsRequest)) *MockFunctionsInterface_FunctionInfoNameToFullNameMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListFunctionsRequest))
	})
	return _c
}

func (_c *MockFunctionsInterface_FunctionInfoNameToFullNameMap_Call) Return(_a0 map[string]string, _a1 error) *MockFunctionsInterface_FunctionInfoNameToFullNameMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFunctionsInterface_FunctionInfoNameToFullNameMap_Call) RunAndReturn(run func(context.Context, catalog.ListFunctionsRequest) (map[string]string, error)) *MockFunctionsInterface_FunctionInfoNameToFullNameMap_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockFunctionsInterface) Get(ctx context.Context, request catalog.GetFunctionRequest) (*catalog.FunctionInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *catalog.FunctionInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetFunctionRequest) (*catalog.FunctionInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetFunctionRequest) *catalog.FunctionInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.FunctionInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.GetFunctionRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFunctionsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockFunctionsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.GetFunctionRequest
func (_e *MockFunctionsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockFunctionsInterface_Get_Call {
	return &MockFunctionsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockFunctionsInterface_Get_Call) Run(run func(ctx context.Context, request catalog.GetFunctionRequest)) *MockFunctionsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.GetFunctionRequest))
	})
	return _c
}

func (_c *MockFunctionsInterface_Get_Call) Return(_a0 *catalog.FunctionInfo, _a1 error) *MockFunctionsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFunctionsInterface_Get_Call) RunAndReturn(run func(context.Context, catalog.GetFunctionRequest) (*catalog.FunctionInfo, error)) *MockFunctionsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *MockFunctionsInterface) GetByName(ctx context.Context, name string) (*catalog.FunctionInfo, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *catalog.FunctionInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*catalog.FunctionInfo, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *catalog.FunctionInfo); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.FunctionInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFunctionsInterface_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type MockFunctionsInterface_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockFunctionsInterface_Expecter) GetByName(ctx interface{}, name interface{}) *MockFunctionsInterface_GetByName_Call {
	return &MockFunctionsInterface_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *MockFunctionsInterface_GetByName_Call) Run(run func(ctx context.Context, name string)) *MockFunctionsInterface_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockFunctionsInterface_GetByName_Call) Return(_a0 *catalog.FunctionInfo, _a1 error) *MockFunctionsInterface_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFunctionsInterface_GetByName_Call) RunAndReturn(run func(context.Context, string) (*catalog.FunctionInfo, error)) *MockFunctionsInterface_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockFunctionsInterface) List(ctx context.Context, request catalog.ListFunctionsRequest) listing.Iterator[catalog.FunctionInfo] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[catalog.FunctionInfo]
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListFunctionsRequest) listing.Iterator[catalog.FunctionInfo]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[catalog.FunctionInfo])
		}
	}

	return r0
}

// MockFunctionsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockFunctionsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListFunctionsRequest
func (_e *MockFunctionsInterface_Expecter) List(ctx interface{}, request interface{}) *MockFunctionsInterface_List_Call {
	return &MockFunctionsInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockFunctionsInterface_List_Call) Run(run func(ctx context.Context, request catalog.ListFunctionsRequest)) *MockFunctionsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListFunctionsRequest))
	})
	return _c
}

func (_c *MockFunctionsInterface_List_Call) Return(_a0 listing.Iterator[catalog.FunctionInfo]) *MockFunctionsInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFunctionsInterface_List_Call) RunAndReturn(run func(context.Context, catalog.ListFunctionsRequest) listing.Iterator[catalog.FunctionInfo]) *MockFunctionsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockFunctionsInterface) ListAll(ctx context.Context, request catalog.ListFunctionsRequest) ([]catalog.FunctionInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []catalog.FunctionInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListFunctionsRequest) ([]catalog.FunctionInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListFunctionsRequest) []catalog.FunctionInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]catalog.FunctionInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.ListFunctionsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFunctionsInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockFunctionsInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListFunctionsRequest
func (_e *MockFunctionsInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockFunctionsInterface_ListAll_Call {
	return &MockFunctionsInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockFunctionsInterface_ListAll_Call) Run(run func(ctx context.Context, request catalog.ListFunctionsRequest)) *MockFunctionsInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListFunctionsRequest))
	})
	return _c
}

func (_c *MockFunctionsInterface_ListAll_Call) Return(_a0 []catalog.FunctionInfo, _a1 error) *MockFunctionsInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFunctionsInterface_ListAll_Call) RunAndReturn(run func(context.Context, catalog.ListFunctionsRequest) ([]catalog.FunctionInfo, error)) *MockFunctionsInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockFunctionsInterface) Update(ctx context.Context, request catalog.UpdateFunction) (*catalog.FunctionInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *catalog.FunctionInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdateFunction) (*catalog.FunctionInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdateFunction) *catalog.FunctionInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.FunctionInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.UpdateFunction) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFunctionsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockFunctionsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.UpdateFunction
func (_e *MockFunctionsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockFunctionsInterface_Update_Call {
	return &MockFunctionsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockFunctionsInterface_Update_Call) Run(run func(ctx context.Context, request catalog.UpdateFunction)) *MockFunctionsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.UpdateFunction))
	})
	return _c
}

func (_c *MockFunctionsInterface_Update_Call) Return(_a0 *catalog.FunctionInfo, _a1 error) *MockFunctionsInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFunctionsInterface_Update_Call) RunAndReturn(run func(context.Context, catalog.UpdateFunction) (*catalog.FunctionInfo, error)) *MockFunctionsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFunctionsInterface creates a new instance of MockFunctionsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFunctionsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFunctionsInterface {
	mock := &MockFunctionsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
