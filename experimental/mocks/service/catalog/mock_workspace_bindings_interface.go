// Code generated by mockery v2.53.2. DO NOT EDIT.

package catalog

import (
	context "context"

	catalog "github.com/databricks/databricks-sdk-go/service/catalog"

	listing "github.com/databricks/databricks-sdk-go/listing"

	mock "github.com/stretchr/testify/mock"
)

// MockWorkspaceBindingsInterface is an autogenerated mock type for the WorkspaceBindingsInterface type
type MockWorkspaceBindingsInterface struct {
	mock.Mock
}

type MockWorkspaceBindingsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWorkspaceBindingsInterface) EXPECT() *MockWorkspaceBindingsInterface_Expecter {
	return &MockWorkspaceBindingsInterface_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockWorkspaceBindingsInterface) Get(ctx context.Context, request catalog.GetWorkspaceBindingRequest) (*catalog.GetCatalogWorkspaceBindingsResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *catalog.GetCatalogWorkspaceBindingsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetWorkspaceBindingRequest) (*catalog.GetCatalogWorkspaceBindingsResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetWorkspaceBindingRequest) *catalog.GetCatalogWorkspaceBindingsResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.GetCatalogWorkspaceBindingsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.GetWorkspaceBindingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkspaceBindingsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockWorkspaceBindingsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.GetWorkspaceBindingRequest
func (_e *MockWorkspaceBindingsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockWorkspaceBindingsInterface_Get_Call {
	return &MockWorkspaceBindingsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockWorkspaceBindingsInterface_Get_Call) Run(run func(ctx context.Context, request catalog.GetWorkspaceBindingRequest)) *MockWorkspaceBindingsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.GetWorkspaceBindingRequest))
	})
	return _c
}

func (_c *MockWorkspaceBindingsInterface_Get_Call) Return(_a0 *catalog.GetCatalogWorkspaceBindingsResponse, _a1 error) *MockWorkspaceBindingsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkspaceBindingsInterface_Get_Call) RunAndReturn(run func(context.Context, catalog.GetWorkspaceBindingRequest) (*catalog.GetCatalogWorkspaceBindingsResponse, error)) *MockWorkspaceBindingsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetBindings provides a mock function with given fields: ctx, request
func (_m *MockWorkspaceBindingsInterface) GetBindings(ctx context.Context, request catalog.GetBindingsRequest) listing.Iterator[catalog.WorkspaceBinding] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetBindings")
	}

	var r0 listing.Iterator[catalog.WorkspaceBinding]
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetBindingsRequest) listing.Iterator[catalog.WorkspaceBinding]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[catalog.WorkspaceBinding])
		}
	}

	return r0
}

// MockWorkspaceBindingsInterface_GetBindings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBindings'
type MockWorkspaceBindingsInterface_GetBindings_Call struct {
	*mock.Call
}

// GetBindings is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.GetBindingsRequest
func (_e *MockWorkspaceBindingsInterface_Expecter) GetBindings(ctx interface{}, request interface{}) *MockWorkspaceBindingsInterface_GetBindings_Call {
	return &MockWorkspaceBindingsInterface_GetBindings_Call{Call: _e.mock.On("GetBindings", ctx, request)}
}

func (_c *MockWorkspaceBindingsInterface_GetBindings_Call) Run(run func(ctx context.Context, request catalog.GetBindingsRequest)) *MockWorkspaceBindingsInterface_GetBindings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.GetBindingsRequest))
	})
	return _c
}

func (_c *MockWorkspaceBindingsInterface_GetBindings_Call) Return(_a0 listing.Iterator[catalog.WorkspaceBinding]) *MockWorkspaceBindingsInterface_GetBindings_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorkspaceBindingsInterface_GetBindings_Call) RunAndReturn(run func(context.Context, catalog.GetBindingsRequest) listing.Iterator[catalog.WorkspaceBinding]) *MockWorkspaceBindingsInterface_GetBindings_Call {
	_c.Call.Return(run)
	return _c
}

// GetBindingsAll provides a mock function with given fields: ctx, request
func (_m *MockWorkspaceBindingsInterface) GetBindingsAll(ctx context.Context, request catalog.GetBindingsRequest) ([]catalog.WorkspaceBinding, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetBindingsAll")
	}

	var r0 []catalog.WorkspaceBinding
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetBindingsRequest) ([]catalog.WorkspaceBinding, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetBindingsRequest) []catalog.WorkspaceBinding); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]catalog.WorkspaceBinding)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.GetBindingsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkspaceBindingsInterface_GetBindingsAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBindingsAll'
type MockWorkspaceBindingsInterface_GetBindingsAll_Call struct {
	*mock.Call
}

// GetBindingsAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.GetBindingsRequest
func (_e *MockWorkspaceBindingsInterface_Expecter) GetBindingsAll(ctx interface{}, request interface{}) *MockWorkspaceBindingsInterface_GetBindingsAll_Call {
	return &MockWorkspaceBindingsInterface_GetBindingsAll_Call{Call: _e.mock.On("GetBindingsAll", ctx, request)}
}

func (_c *MockWorkspaceBindingsInterface_GetBindingsAll_Call) Run(run func(ctx context.Context, request catalog.GetBindingsRequest)) *MockWorkspaceBindingsInterface_GetBindingsAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.GetBindingsRequest))
	})
	return _c
}

func (_c *MockWorkspaceBindingsInterface_GetBindingsAll_Call) Return(_a0 []catalog.WorkspaceBinding, _a1 error) *MockWorkspaceBindingsInterface_GetBindingsAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkspaceBindingsInterface_GetBindingsAll_Call) RunAndReturn(run func(context.Context, catalog.GetBindingsRequest) ([]catalog.WorkspaceBinding, error)) *MockWorkspaceBindingsInterface_GetBindingsAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetBindingsBySecurableTypeAndSecurableName provides a mock function with given fields: ctx, securableType, securableName
func (_m *MockWorkspaceBindingsInterface) GetBindingsBySecurableTypeAndSecurableName(ctx context.Context, securableType string, securableName string) (*catalog.GetWorkspaceBindingsResponse, error) {
	ret := _m.Called(ctx, securableType, securableName)

	if len(ret) == 0 {
		panic("no return value specified for GetBindingsBySecurableTypeAndSecurableName")
	}

	var r0 *catalog.GetWorkspaceBindingsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*catalog.GetWorkspaceBindingsResponse, error)); ok {
		return rf(ctx, securableType, securableName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *catalog.GetWorkspaceBindingsResponse); ok {
		r0 = rf(ctx, securableType, securableName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.GetWorkspaceBindingsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, securableType, securableName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkspaceBindingsInterface_GetBindingsBySecurableTypeAndSecurableName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBindingsBySecurableTypeAndSecurableName'
type MockWorkspaceBindingsInterface_GetBindingsBySecurableTypeAndSecurableName_Call struct {
	*mock.Call
}

// GetBindingsBySecurableTypeAndSecurableName is a helper method to define mock.On call
//   - ctx context.Context
//   - securableType string
//   - securableName string
func (_e *MockWorkspaceBindingsInterface_Expecter) GetBindingsBySecurableTypeAndSecurableName(ctx interface{}, securableType interface{}, securableName interface{}) *MockWorkspaceBindingsInterface_GetBindingsBySecurableTypeAndSecurableName_Call {
	return &MockWorkspaceBindingsInterface_GetBindingsBySecurableTypeAndSecurableName_Call{Call: _e.mock.On("GetBindingsBySecurableTypeAndSecurableName", ctx, securableType, securableName)}
}

func (_c *MockWorkspaceBindingsInterface_GetBindingsBySecurableTypeAndSecurableName_Call) Run(run func(ctx context.Context, securableType string, securableName string)) *MockWorkspaceBindingsInterface_GetBindingsBySecurableTypeAndSecurableName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockWorkspaceBindingsInterface_GetBindingsBySecurableTypeAndSecurableName_Call) Return(_a0 *catalog.GetWorkspaceBindingsResponse, _a1 error) *MockWorkspaceBindingsInterface_GetBindingsBySecurableTypeAndSecurableName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkspaceBindingsInterface_GetBindingsBySecurableTypeAndSecurableName_Call) RunAndReturn(run func(context.Context, string, string) (*catalog.GetWorkspaceBindingsResponse, error)) *MockWorkspaceBindingsInterface_GetBindingsBySecurableTypeAndSecurableName_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *MockWorkspaceBindingsInterface) GetByName(ctx context.Context, name string) (*catalog.GetCatalogWorkspaceBindingsResponse, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *catalog.GetCatalogWorkspaceBindingsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*catalog.GetCatalogWorkspaceBindingsResponse, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *catalog.GetCatalogWorkspaceBindingsResponse); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.GetCatalogWorkspaceBindingsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkspaceBindingsInterface_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type MockWorkspaceBindingsInterface_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockWorkspaceBindingsInterface_Expecter) GetByName(ctx interface{}, name interface{}) *MockWorkspaceBindingsInterface_GetByName_Call {
	return &MockWorkspaceBindingsInterface_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *MockWorkspaceBindingsInterface_GetByName_Call) Run(run func(ctx context.Context, name string)) *MockWorkspaceBindingsInterface_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockWorkspaceBindingsInterface_GetByName_Call) Return(_a0 *catalog.GetCatalogWorkspaceBindingsResponse, _a1 error) *MockWorkspaceBindingsInterface_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkspaceBindingsInterface_GetByName_Call) RunAndReturn(run func(context.Context, string) (*catalog.GetCatalogWorkspaceBindingsResponse, error)) *MockWorkspaceBindingsInterface_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockWorkspaceBindingsInterface) Update(ctx context.Context, request catalog.UpdateWorkspaceBindings) (*catalog.UpdateCatalogWorkspaceBindingsResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *catalog.UpdateCatalogWorkspaceBindingsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdateWorkspaceBindings) (*catalog.UpdateCatalogWorkspaceBindingsResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdateWorkspaceBindings) *catalog.UpdateCatalogWorkspaceBindingsResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.UpdateCatalogWorkspaceBindingsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.UpdateWorkspaceBindings) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkspaceBindingsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockWorkspaceBindingsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.UpdateWorkspaceBindings
func (_e *MockWorkspaceBindingsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockWorkspaceBindingsInterface_Update_Call {
	return &MockWorkspaceBindingsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockWorkspaceBindingsInterface_Update_Call) Run(run func(ctx context.Context, request catalog.UpdateWorkspaceBindings)) *MockWorkspaceBindingsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.UpdateWorkspaceBindings))
	})
	return _c
}

func (_c *MockWorkspaceBindingsInterface_Update_Call) Return(_a0 *catalog.UpdateCatalogWorkspaceBindingsResponse, _a1 error) *MockWorkspaceBindingsInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkspaceBindingsInterface_Update_Call) RunAndReturn(run func(context.Context, catalog.UpdateWorkspaceBindings) (*catalog.UpdateCatalogWorkspaceBindingsResponse, error)) *MockWorkspaceBindingsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateBindings provides a mock function with given fields: ctx, request
func (_m *MockWorkspaceBindingsInterface) UpdateBindings(ctx context.Context, request catalog.UpdateWorkspaceBindingsParameters) (*catalog.UpdateWorkspaceBindingsResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for UpdateBindings")
	}

	var r0 *catalog.UpdateWorkspaceBindingsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdateWorkspaceBindingsParameters) (*catalog.UpdateWorkspaceBindingsResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdateWorkspaceBindingsParameters) *catalog.UpdateWorkspaceBindingsResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.UpdateWorkspaceBindingsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.UpdateWorkspaceBindingsParameters) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkspaceBindingsInterface_UpdateBindings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateBindings'
type MockWorkspaceBindingsInterface_UpdateBindings_Call struct {
	*mock.Call
}

// UpdateBindings is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.UpdateWorkspaceBindingsParameters
func (_e *MockWorkspaceBindingsInterface_Expecter) UpdateBindings(ctx interface{}, request interface{}) *MockWorkspaceBindingsInterface_UpdateBindings_Call {
	return &MockWorkspaceBindingsInterface_UpdateBindings_Call{Call: _e.mock.On("UpdateBindings", ctx, request)}
}

func (_c *MockWorkspaceBindingsInterface_UpdateBindings_Call) Run(run func(ctx context.Context, request catalog.UpdateWorkspaceBindingsParameters)) *MockWorkspaceBindingsInterface_UpdateBindings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.UpdateWorkspaceBindingsParameters))
	})
	return _c
}

func (_c *MockWorkspaceBindingsInterface_UpdateBindings_Call) Return(_a0 *catalog.UpdateWorkspaceBindingsResponse, _a1 error) *MockWorkspaceBindingsInterface_UpdateBindings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkspaceBindingsInterface_UpdateBindings_Call) RunAndReturn(run func(context.Context, catalog.UpdateWorkspaceBindingsParameters) (*catalog.UpdateWorkspaceBindingsResponse, error)) *MockWorkspaceBindingsInterface_UpdateBindings_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWorkspaceBindingsInterface creates a new instance of MockWorkspaceBindingsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWorkspaceBindingsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWorkspaceBindingsInterface {
	mock := &MockWorkspaceBindingsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
