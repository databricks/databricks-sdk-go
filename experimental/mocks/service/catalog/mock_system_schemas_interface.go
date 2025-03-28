// Code generated by mockery v2.53.2. DO NOT EDIT.

package catalog

import (
	context "context"

	catalog "github.com/databricks/databricks-sdk-go/service/catalog"

	listing "github.com/databricks/databricks-sdk-go/listing"

	mock "github.com/stretchr/testify/mock"
)

// MockSystemSchemasInterface is an autogenerated mock type for the SystemSchemasInterface type
type MockSystemSchemasInterface struct {
	mock.Mock
}

type MockSystemSchemasInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSystemSchemasInterface) EXPECT() *MockSystemSchemasInterface_Expecter {
	return &MockSystemSchemasInterface_Expecter{mock: &_m.Mock}
}

// Disable provides a mock function with given fields: ctx, request
func (_m *MockSystemSchemasInterface) Disable(ctx context.Context, request catalog.DisableRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Disable")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.DisableRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSystemSchemasInterface_Disable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Disable'
type MockSystemSchemasInterface_Disable_Call struct {
	*mock.Call
}

// Disable is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.DisableRequest
func (_e *MockSystemSchemasInterface_Expecter) Disable(ctx interface{}, request interface{}) *MockSystemSchemasInterface_Disable_Call {
	return &MockSystemSchemasInterface_Disable_Call{Call: _e.mock.On("Disable", ctx, request)}
}

func (_c *MockSystemSchemasInterface_Disable_Call) Run(run func(ctx context.Context, request catalog.DisableRequest)) *MockSystemSchemasInterface_Disable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.DisableRequest))
	})
	return _c
}

func (_c *MockSystemSchemasInterface_Disable_Call) Return(_a0 error) *MockSystemSchemasInterface_Disable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSystemSchemasInterface_Disable_Call) RunAndReturn(run func(context.Context, catalog.DisableRequest) error) *MockSystemSchemasInterface_Disable_Call {
	_c.Call.Return(run)
	return _c
}

// DisableByMetastoreIdAndSchemaName provides a mock function with given fields: ctx, metastoreId, schemaName
func (_m *MockSystemSchemasInterface) DisableByMetastoreIdAndSchemaName(ctx context.Context, metastoreId string, schemaName string) error {
	ret := _m.Called(ctx, metastoreId, schemaName)

	if len(ret) == 0 {
		panic("no return value specified for DisableByMetastoreIdAndSchemaName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, metastoreId, schemaName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSystemSchemasInterface_DisableByMetastoreIdAndSchemaName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DisableByMetastoreIdAndSchemaName'
type MockSystemSchemasInterface_DisableByMetastoreIdAndSchemaName_Call struct {
	*mock.Call
}

// DisableByMetastoreIdAndSchemaName is a helper method to define mock.On call
//   - ctx context.Context
//   - metastoreId string
//   - schemaName string
func (_e *MockSystemSchemasInterface_Expecter) DisableByMetastoreIdAndSchemaName(ctx interface{}, metastoreId interface{}, schemaName interface{}) *MockSystemSchemasInterface_DisableByMetastoreIdAndSchemaName_Call {
	return &MockSystemSchemasInterface_DisableByMetastoreIdAndSchemaName_Call{Call: _e.mock.On("DisableByMetastoreIdAndSchemaName", ctx, metastoreId, schemaName)}
}

func (_c *MockSystemSchemasInterface_DisableByMetastoreIdAndSchemaName_Call) Run(run func(ctx context.Context, metastoreId string, schemaName string)) *MockSystemSchemasInterface_DisableByMetastoreIdAndSchemaName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockSystemSchemasInterface_DisableByMetastoreIdAndSchemaName_Call) Return(_a0 error) *MockSystemSchemasInterface_DisableByMetastoreIdAndSchemaName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSystemSchemasInterface_DisableByMetastoreIdAndSchemaName_Call) RunAndReturn(run func(context.Context, string, string) error) *MockSystemSchemasInterface_DisableByMetastoreIdAndSchemaName_Call {
	_c.Call.Return(run)
	return _c
}

// Enable provides a mock function with given fields: ctx, request
func (_m *MockSystemSchemasInterface) Enable(ctx context.Context, request catalog.EnableRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Enable")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.EnableRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSystemSchemasInterface_Enable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Enable'
type MockSystemSchemasInterface_Enable_Call struct {
	*mock.Call
}

// Enable is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.EnableRequest
func (_e *MockSystemSchemasInterface_Expecter) Enable(ctx interface{}, request interface{}) *MockSystemSchemasInterface_Enable_Call {
	return &MockSystemSchemasInterface_Enable_Call{Call: _e.mock.On("Enable", ctx, request)}
}

func (_c *MockSystemSchemasInterface_Enable_Call) Run(run func(ctx context.Context, request catalog.EnableRequest)) *MockSystemSchemasInterface_Enable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.EnableRequest))
	})
	return _c
}

func (_c *MockSystemSchemasInterface_Enable_Call) Return(_a0 error) *MockSystemSchemasInterface_Enable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSystemSchemasInterface_Enable_Call) RunAndReturn(run func(context.Context, catalog.EnableRequest) error) *MockSystemSchemasInterface_Enable_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockSystemSchemasInterface) List(ctx context.Context, request catalog.ListSystemSchemasRequest) listing.Iterator[catalog.SystemSchemaInfo] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[catalog.SystemSchemaInfo]
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListSystemSchemasRequest) listing.Iterator[catalog.SystemSchemaInfo]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[catalog.SystemSchemaInfo])
		}
	}

	return r0
}

// MockSystemSchemasInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockSystemSchemasInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListSystemSchemasRequest
func (_e *MockSystemSchemasInterface_Expecter) List(ctx interface{}, request interface{}) *MockSystemSchemasInterface_List_Call {
	return &MockSystemSchemasInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockSystemSchemasInterface_List_Call) Run(run func(ctx context.Context, request catalog.ListSystemSchemasRequest)) *MockSystemSchemasInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListSystemSchemasRequest))
	})
	return _c
}

func (_c *MockSystemSchemasInterface_List_Call) Return(_a0 listing.Iterator[catalog.SystemSchemaInfo]) *MockSystemSchemasInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSystemSchemasInterface_List_Call) RunAndReturn(run func(context.Context, catalog.ListSystemSchemasRequest) listing.Iterator[catalog.SystemSchemaInfo]) *MockSystemSchemasInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockSystemSchemasInterface) ListAll(ctx context.Context, request catalog.ListSystemSchemasRequest) ([]catalog.SystemSchemaInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []catalog.SystemSchemaInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListSystemSchemasRequest) ([]catalog.SystemSchemaInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListSystemSchemasRequest) []catalog.SystemSchemaInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]catalog.SystemSchemaInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.ListSystemSchemasRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSystemSchemasInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockSystemSchemasInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListSystemSchemasRequest
func (_e *MockSystemSchemasInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockSystemSchemasInterface_ListAll_Call {
	return &MockSystemSchemasInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockSystemSchemasInterface_ListAll_Call) Run(run func(ctx context.Context, request catalog.ListSystemSchemasRequest)) *MockSystemSchemasInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListSystemSchemasRequest))
	})
	return _c
}

func (_c *MockSystemSchemasInterface_ListAll_Call) Return(_a0 []catalog.SystemSchemaInfo, _a1 error) *MockSystemSchemasInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSystemSchemasInterface_ListAll_Call) RunAndReturn(run func(context.Context, catalog.ListSystemSchemasRequest) ([]catalog.SystemSchemaInfo, error)) *MockSystemSchemasInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// ListByMetastoreId provides a mock function with given fields: ctx, metastoreId
func (_m *MockSystemSchemasInterface) ListByMetastoreId(ctx context.Context, metastoreId string) (*catalog.ListSystemSchemasResponse, error) {
	ret := _m.Called(ctx, metastoreId)

	if len(ret) == 0 {
		panic("no return value specified for ListByMetastoreId")
	}

	var r0 *catalog.ListSystemSchemasResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*catalog.ListSystemSchemasResponse, error)); ok {
		return rf(ctx, metastoreId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *catalog.ListSystemSchemasResponse); ok {
		r0 = rf(ctx, metastoreId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.ListSystemSchemasResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, metastoreId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSystemSchemasInterface_ListByMetastoreId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListByMetastoreId'
type MockSystemSchemasInterface_ListByMetastoreId_Call struct {
	*mock.Call
}

// ListByMetastoreId is a helper method to define mock.On call
//   - ctx context.Context
//   - metastoreId string
func (_e *MockSystemSchemasInterface_Expecter) ListByMetastoreId(ctx interface{}, metastoreId interface{}) *MockSystemSchemasInterface_ListByMetastoreId_Call {
	return &MockSystemSchemasInterface_ListByMetastoreId_Call{Call: _e.mock.On("ListByMetastoreId", ctx, metastoreId)}
}

func (_c *MockSystemSchemasInterface_ListByMetastoreId_Call) Run(run func(ctx context.Context, metastoreId string)) *MockSystemSchemasInterface_ListByMetastoreId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSystemSchemasInterface_ListByMetastoreId_Call) Return(_a0 *catalog.ListSystemSchemasResponse, _a1 error) *MockSystemSchemasInterface_ListByMetastoreId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSystemSchemasInterface_ListByMetastoreId_Call) RunAndReturn(run func(context.Context, string) (*catalog.ListSystemSchemasResponse, error)) *MockSystemSchemasInterface_ListByMetastoreId_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSystemSchemasInterface creates a new instance of MockSystemSchemasInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSystemSchemasInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSystemSchemasInterface {
	mock := &MockSystemSchemasInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
