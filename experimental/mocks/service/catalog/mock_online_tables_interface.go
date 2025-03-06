// Code generated by mockery v2.53.0. DO NOT EDIT.

package catalog

import (
	context "context"

	catalog "github.com/databricks/databricks-sdk-go/service/catalog"

	mock "github.com/stretchr/testify/mock"

	retries "github.com/databricks/databricks-sdk-go/retries"

	time "time"
)

// MockOnlineTablesInterface is an autogenerated mock type for the OnlineTablesInterface type
type MockOnlineTablesInterface struct {
	mock.Mock
}

type MockOnlineTablesInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockOnlineTablesInterface) EXPECT() *MockOnlineTablesInterface_Expecter {
	return &MockOnlineTablesInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, createOnlineTableRequest
func (_m *MockOnlineTablesInterface) Create(ctx context.Context, createOnlineTableRequest catalog.CreateOnlineTableRequest) (*catalog.WaitGetOnlineTableActive[catalog.OnlineTable], error) {
	ret := _m.Called(ctx, createOnlineTableRequest)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *catalog.WaitGetOnlineTableActive[catalog.OnlineTable]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateOnlineTableRequest) (*catalog.WaitGetOnlineTableActive[catalog.OnlineTable], error)); ok {
		return rf(ctx, createOnlineTableRequest)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateOnlineTableRequest) *catalog.WaitGetOnlineTableActive[catalog.OnlineTable]); ok {
		r0 = rf(ctx, createOnlineTableRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.WaitGetOnlineTableActive[catalog.OnlineTable])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.CreateOnlineTableRequest) error); ok {
		r1 = rf(ctx, createOnlineTableRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOnlineTablesInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockOnlineTablesInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - createOnlineTableRequest catalog.CreateOnlineTableRequest
func (_e *MockOnlineTablesInterface_Expecter) Create(ctx interface{}, createOnlineTableRequest interface{}) *MockOnlineTablesInterface_Create_Call {
	return &MockOnlineTablesInterface_Create_Call{Call: _e.mock.On("Create", ctx, createOnlineTableRequest)}
}

func (_c *MockOnlineTablesInterface_Create_Call) Run(run func(ctx context.Context, createOnlineTableRequest catalog.CreateOnlineTableRequest)) *MockOnlineTablesInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.CreateOnlineTableRequest))
	})
	return _c
}

func (_c *MockOnlineTablesInterface_Create_Call) Return(_a0 *catalog.WaitGetOnlineTableActive[catalog.OnlineTable], _a1 error) *MockOnlineTablesInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockOnlineTablesInterface_Create_Call) RunAndReturn(run func(context.Context, catalog.CreateOnlineTableRequest) (*catalog.WaitGetOnlineTableActive[catalog.OnlineTable], error)) *MockOnlineTablesInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// CreateAndWait provides a mock function with given fields: ctx, createOnlineTableRequest, options
func (_m *MockOnlineTablesInterface) CreateAndWait(ctx context.Context, createOnlineTableRequest catalog.CreateOnlineTableRequest, options ...retries.Option[catalog.OnlineTable]) (*catalog.OnlineTable, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, createOnlineTableRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateAndWait")
	}

	var r0 *catalog.OnlineTable
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateOnlineTableRequest, ...retries.Option[catalog.OnlineTable]) (*catalog.OnlineTable, error)); ok {
		return rf(ctx, createOnlineTableRequest, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateOnlineTableRequest, ...retries.Option[catalog.OnlineTable]) *catalog.OnlineTable); ok {
		r0 = rf(ctx, createOnlineTableRequest, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.OnlineTable)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.CreateOnlineTableRequest, ...retries.Option[catalog.OnlineTable]) error); ok {
		r1 = rf(ctx, createOnlineTableRequest, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOnlineTablesInterface_CreateAndWait_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAndWait'
type MockOnlineTablesInterface_CreateAndWait_Call struct {
	*mock.Call
}

// CreateAndWait is a helper method to define mock.On call
//   - ctx context.Context
//   - createOnlineTableRequest catalog.CreateOnlineTableRequest
//   - options ...retries.Option[catalog.OnlineTable]
func (_e *MockOnlineTablesInterface_Expecter) CreateAndWait(ctx interface{}, createOnlineTableRequest interface{}, options ...interface{}) *MockOnlineTablesInterface_CreateAndWait_Call {
	return &MockOnlineTablesInterface_CreateAndWait_Call{Call: _e.mock.On("CreateAndWait",
		append([]interface{}{ctx, createOnlineTableRequest}, options...)...)}
}

func (_c *MockOnlineTablesInterface_CreateAndWait_Call) Run(run func(ctx context.Context, createOnlineTableRequest catalog.CreateOnlineTableRequest, options ...retries.Option[catalog.OnlineTable])) *MockOnlineTablesInterface_CreateAndWait_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]retries.Option[catalog.OnlineTable], len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(retries.Option[catalog.OnlineTable])
			}
		}
		run(args[0].(context.Context), args[1].(catalog.CreateOnlineTableRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockOnlineTablesInterface_CreateAndWait_Call) Return(_a0 *catalog.OnlineTable, _a1 error) *MockOnlineTablesInterface_CreateAndWait_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockOnlineTablesInterface_CreateAndWait_Call) RunAndReturn(run func(context.Context, catalog.CreateOnlineTableRequest, ...retries.Option[catalog.OnlineTable]) (*catalog.OnlineTable, error)) *MockOnlineTablesInterface_CreateAndWait_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockOnlineTablesInterface) Delete(ctx context.Context, request catalog.DeleteOnlineTableRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.DeleteOnlineTableRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockOnlineTablesInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockOnlineTablesInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.DeleteOnlineTableRequest
func (_e *MockOnlineTablesInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockOnlineTablesInterface_Delete_Call {
	return &MockOnlineTablesInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockOnlineTablesInterface_Delete_Call) Run(run func(ctx context.Context, request catalog.DeleteOnlineTableRequest)) *MockOnlineTablesInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.DeleteOnlineTableRequest))
	})
	return _c
}

func (_c *MockOnlineTablesInterface_Delete_Call) Return(_a0 error) *MockOnlineTablesInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOnlineTablesInterface_Delete_Call) RunAndReturn(run func(context.Context, catalog.DeleteOnlineTableRequest) error) *MockOnlineTablesInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByName provides a mock function with given fields: ctx, name
func (_m *MockOnlineTablesInterface) DeleteByName(ctx context.Context, name string) error {
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

// MockOnlineTablesInterface_DeleteByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByName'
type MockOnlineTablesInterface_DeleteByName_Call struct {
	*mock.Call
}

// DeleteByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockOnlineTablesInterface_Expecter) DeleteByName(ctx interface{}, name interface{}) *MockOnlineTablesInterface_DeleteByName_Call {
	return &MockOnlineTablesInterface_DeleteByName_Call{Call: _e.mock.On("DeleteByName", ctx, name)}
}

func (_c *MockOnlineTablesInterface_DeleteByName_Call) Run(run func(ctx context.Context, name string)) *MockOnlineTablesInterface_DeleteByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockOnlineTablesInterface_DeleteByName_Call) Return(_a0 error) *MockOnlineTablesInterface_DeleteByName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOnlineTablesInterface_DeleteByName_Call) RunAndReturn(run func(context.Context, string) error) *MockOnlineTablesInterface_DeleteByName_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockOnlineTablesInterface) Get(ctx context.Context, request catalog.GetOnlineTableRequest) (*catalog.OnlineTable, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *catalog.OnlineTable
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetOnlineTableRequest) (*catalog.OnlineTable, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetOnlineTableRequest) *catalog.OnlineTable); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.OnlineTable)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.GetOnlineTableRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOnlineTablesInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockOnlineTablesInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.GetOnlineTableRequest
func (_e *MockOnlineTablesInterface_Expecter) Get(ctx interface{}, request interface{}) *MockOnlineTablesInterface_Get_Call {
	return &MockOnlineTablesInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockOnlineTablesInterface_Get_Call) Run(run func(ctx context.Context, request catalog.GetOnlineTableRequest)) *MockOnlineTablesInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.GetOnlineTableRequest))
	})
	return _c
}

func (_c *MockOnlineTablesInterface_Get_Call) Return(_a0 *catalog.OnlineTable, _a1 error) *MockOnlineTablesInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockOnlineTablesInterface_Get_Call) RunAndReturn(run func(context.Context, catalog.GetOnlineTableRequest) (*catalog.OnlineTable, error)) *MockOnlineTablesInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *MockOnlineTablesInterface) GetByName(ctx context.Context, name string) (*catalog.OnlineTable, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *catalog.OnlineTable
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*catalog.OnlineTable, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *catalog.OnlineTable); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.OnlineTable)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOnlineTablesInterface_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type MockOnlineTablesInterface_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockOnlineTablesInterface_Expecter) GetByName(ctx interface{}, name interface{}) *MockOnlineTablesInterface_GetByName_Call {
	return &MockOnlineTablesInterface_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *MockOnlineTablesInterface_GetByName_Call) Run(run func(ctx context.Context, name string)) *MockOnlineTablesInterface_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockOnlineTablesInterface_GetByName_Call) Return(_a0 *catalog.OnlineTable, _a1 error) *MockOnlineTablesInterface_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockOnlineTablesInterface_GetByName_Call) RunAndReturn(run func(context.Context, string) (*catalog.OnlineTable, error)) *MockOnlineTablesInterface_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// WaitGetOnlineTableActive provides a mock function with given fields: ctx, name, timeout, callback
func (_m *MockOnlineTablesInterface) WaitGetOnlineTableActive(ctx context.Context, name string, timeout time.Duration, callback func(*catalog.OnlineTable)) (*catalog.OnlineTable, error) {
	ret := _m.Called(ctx, name, timeout, callback)

	if len(ret) == 0 {
		panic("no return value specified for WaitGetOnlineTableActive")
	}

	var r0 *catalog.OnlineTable
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration, func(*catalog.OnlineTable)) (*catalog.OnlineTable, error)); ok {
		return rf(ctx, name, timeout, callback)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration, func(*catalog.OnlineTable)) *catalog.OnlineTable); ok {
		r0 = rf(ctx, name, timeout, callback)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.OnlineTable)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, time.Duration, func(*catalog.OnlineTable)) error); ok {
		r1 = rf(ctx, name, timeout, callback)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOnlineTablesInterface_WaitGetOnlineTableActive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WaitGetOnlineTableActive'
type MockOnlineTablesInterface_WaitGetOnlineTableActive_Call struct {
	*mock.Call
}

// WaitGetOnlineTableActive is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - timeout time.Duration
//   - callback func(*catalog.OnlineTable)
func (_e *MockOnlineTablesInterface_Expecter) WaitGetOnlineTableActive(ctx interface{}, name interface{}, timeout interface{}, callback interface{}) *MockOnlineTablesInterface_WaitGetOnlineTableActive_Call {
	return &MockOnlineTablesInterface_WaitGetOnlineTableActive_Call{Call: _e.mock.On("WaitGetOnlineTableActive", ctx, name, timeout, callback)}
}

func (_c *MockOnlineTablesInterface_WaitGetOnlineTableActive_Call) Run(run func(ctx context.Context, name string, timeout time.Duration, callback func(*catalog.OnlineTable))) *MockOnlineTablesInterface_WaitGetOnlineTableActive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(time.Duration), args[3].(func(*catalog.OnlineTable)))
	})
	return _c
}

func (_c *MockOnlineTablesInterface_WaitGetOnlineTableActive_Call) Return(_a0 *catalog.OnlineTable, _a1 error) *MockOnlineTablesInterface_WaitGetOnlineTableActive_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockOnlineTablesInterface_WaitGetOnlineTableActive_Call) RunAndReturn(run func(context.Context, string, time.Duration, func(*catalog.OnlineTable)) (*catalog.OnlineTable, error)) *MockOnlineTablesInterface_WaitGetOnlineTableActive_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockOnlineTablesInterface creates a new instance of MockOnlineTablesInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOnlineTablesInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOnlineTablesInterface {
	mock := &MockOnlineTablesInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
