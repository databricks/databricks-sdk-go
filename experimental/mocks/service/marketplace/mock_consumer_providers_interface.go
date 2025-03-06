// Code generated by mockery v2.53.0. DO NOT EDIT.

package marketplace

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	marketplace "github.com/databricks/databricks-sdk-go/service/marketplace"

	mock "github.com/stretchr/testify/mock"
)

// MockConsumerProvidersInterface is an autogenerated mock type for the ConsumerProvidersInterface type
type MockConsumerProvidersInterface struct {
	mock.Mock
}

type MockConsumerProvidersInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConsumerProvidersInterface) EXPECT() *MockConsumerProvidersInterface_Expecter {
	return &MockConsumerProvidersInterface_Expecter{mock: &_m.Mock}
}

// BatchGet provides a mock function with given fields: ctx, request
func (_m *MockConsumerProvidersInterface) BatchGet(ctx context.Context, request marketplace.BatchGetProvidersRequest) (*marketplace.BatchGetProvidersResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for BatchGet")
	}

	var r0 *marketplace.BatchGetProvidersResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.BatchGetProvidersRequest) (*marketplace.BatchGetProvidersResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.BatchGetProvidersRequest) *marketplace.BatchGetProvidersResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplace.BatchGetProvidersResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, marketplace.BatchGetProvidersRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConsumerProvidersInterface_BatchGet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BatchGet'
type MockConsumerProvidersInterface_BatchGet_Call struct {
	*mock.Call
}

// BatchGet is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.BatchGetProvidersRequest
func (_e *MockConsumerProvidersInterface_Expecter) BatchGet(ctx interface{}, request interface{}) *MockConsumerProvidersInterface_BatchGet_Call {
	return &MockConsumerProvidersInterface_BatchGet_Call{Call: _e.mock.On("BatchGet", ctx, request)}
}

func (_c *MockConsumerProvidersInterface_BatchGet_Call) Run(run func(ctx context.Context, request marketplace.BatchGetProvidersRequest)) *MockConsumerProvidersInterface_BatchGet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.BatchGetProvidersRequest))
	})
	return _c
}

func (_c *MockConsumerProvidersInterface_BatchGet_Call) Return(_a0 *marketplace.BatchGetProvidersResponse, _a1 error) *MockConsumerProvidersInterface_BatchGet_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConsumerProvidersInterface_BatchGet_Call) RunAndReturn(run func(context.Context, marketplace.BatchGetProvidersRequest) (*marketplace.BatchGetProvidersResponse, error)) *MockConsumerProvidersInterface_BatchGet_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockConsumerProvidersInterface) Get(ctx context.Context, request marketplace.GetProviderRequest) (*marketplace.GetProviderResponse, error) {
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

// MockConsumerProvidersInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockConsumerProvidersInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.GetProviderRequest
func (_e *MockConsumerProvidersInterface_Expecter) Get(ctx interface{}, request interface{}) *MockConsumerProvidersInterface_Get_Call {
	return &MockConsumerProvidersInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockConsumerProvidersInterface_Get_Call) Run(run func(ctx context.Context, request marketplace.GetProviderRequest)) *MockConsumerProvidersInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.GetProviderRequest))
	})
	return _c
}

func (_c *MockConsumerProvidersInterface_Get_Call) Return(_a0 *marketplace.GetProviderResponse, _a1 error) *MockConsumerProvidersInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConsumerProvidersInterface_Get_Call) RunAndReturn(run func(context.Context, marketplace.GetProviderRequest) (*marketplace.GetProviderResponse, error)) *MockConsumerProvidersInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: ctx, id
func (_m *MockConsumerProvidersInterface) GetById(ctx context.Context, id string) (*marketplace.GetProviderResponse, error) {
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

// MockConsumerProvidersInterface_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type MockConsumerProvidersInterface_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockConsumerProvidersInterface_Expecter) GetById(ctx interface{}, id interface{}) *MockConsumerProvidersInterface_GetById_Call {
	return &MockConsumerProvidersInterface_GetById_Call{Call: _e.mock.On("GetById", ctx, id)}
}

func (_c *MockConsumerProvidersInterface_GetById_Call) Run(run func(ctx context.Context, id string)) *MockConsumerProvidersInterface_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockConsumerProvidersInterface_GetById_Call) Return(_a0 *marketplace.GetProviderResponse, _a1 error) *MockConsumerProvidersInterface_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConsumerProvidersInterface_GetById_Call) RunAndReturn(run func(context.Context, string) (*marketplace.GetProviderResponse, error)) *MockConsumerProvidersInterface_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *MockConsumerProvidersInterface) GetByName(ctx context.Context, name string) (*marketplace.ProviderInfo, error) {
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

// MockConsumerProvidersInterface_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type MockConsumerProvidersInterface_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockConsumerProvidersInterface_Expecter) GetByName(ctx interface{}, name interface{}) *MockConsumerProvidersInterface_GetByName_Call {
	return &MockConsumerProvidersInterface_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *MockConsumerProvidersInterface_GetByName_Call) Run(run func(ctx context.Context, name string)) *MockConsumerProvidersInterface_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockConsumerProvidersInterface_GetByName_Call) Return(_a0 *marketplace.ProviderInfo, _a1 error) *MockConsumerProvidersInterface_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConsumerProvidersInterface_GetByName_Call) RunAndReturn(run func(context.Context, string) (*marketplace.ProviderInfo, error)) *MockConsumerProvidersInterface_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockConsumerProvidersInterface) List(ctx context.Context, request marketplace.ListProvidersRequest) listing.Iterator[marketplace.ProviderInfo] {
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

// MockConsumerProvidersInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockConsumerProvidersInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.ListProvidersRequest
func (_e *MockConsumerProvidersInterface_Expecter) List(ctx interface{}, request interface{}) *MockConsumerProvidersInterface_List_Call {
	return &MockConsumerProvidersInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockConsumerProvidersInterface_List_Call) Run(run func(ctx context.Context, request marketplace.ListProvidersRequest)) *MockConsumerProvidersInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.ListProvidersRequest))
	})
	return _c
}

func (_c *MockConsumerProvidersInterface_List_Call) Return(_a0 listing.Iterator[marketplace.ProviderInfo]) *MockConsumerProvidersInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConsumerProvidersInterface_List_Call) RunAndReturn(run func(context.Context, marketplace.ListProvidersRequest) listing.Iterator[marketplace.ProviderInfo]) *MockConsumerProvidersInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockConsumerProvidersInterface) ListAll(ctx context.Context, request marketplace.ListProvidersRequest) ([]marketplace.ProviderInfo, error) {
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

// MockConsumerProvidersInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockConsumerProvidersInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.ListProvidersRequest
func (_e *MockConsumerProvidersInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockConsumerProvidersInterface_ListAll_Call {
	return &MockConsumerProvidersInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockConsumerProvidersInterface_ListAll_Call) Run(run func(ctx context.Context, request marketplace.ListProvidersRequest)) *MockConsumerProvidersInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.ListProvidersRequest))
	})
	return _c
}

func (_c *MockConsumerProvidersInterface_ListAll_Call) Return(_a0 []marketplace.ProviderInfo, _a1 error) *MockConsumerProvidersInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConsumerProvidersInterface_ListAll_Call) RunAndReturn(run func(context.Context, marketplace.ListProvidersRequest) ([]marketplace.ProviderInfo, error)) *MockConsumerProvidersInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// ProviderInfoNameToIdMap provides a mock function with given fields: ctx, request
func (_m *MockConsumerProvidersInterface) ProviderInfoNameToIdMap(ctx context.Context, request marketplace.ListProvidersRequest) (map[string]string, error) {
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

// MockConsumerProvidersInterface_ProviderInfoNameToIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProviderInfoNameToIdMap'
type MockConsumerProvidersInterface_ProviderInfoNameToIdMap_Call struct {
	*mock.Call
}

// ProviderInfoNameToIdMap is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.ListProvidersRequest
func (_e *MockConsumerProvidersInterface_Expecter) ProviderInfoNameToIdMap(ctx interface{}, request interface{}) *MockConsumerProvidersInterface_ProviderInfoNameToIdMap_Call {
	return &MockConsumerProvidersInterface_ProviderInfoNameToIdMap_Call{Call: _e.mock.On("ProviderInfoNameToIdMap", ctx, request)}
}

func (_c *MockConsumerProvidersInterface_ProviderInfoNameToIdMap_Call) Run(run func(ctx context.Context, request marketplace.ListProvidersRequest)) *MockConsumerProvidersInterface_ProviderInfoNameToIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.ListProvidersRequest))
	})
	return _c
}

func (_c *MockConsumerProvidersInterface_ProviderInfoNameToIdMap_Call) Return(_a0 map[string]string, _a1 error) *MockConsumerProvidersInterface_ProviderInfoNameToIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConsumerProvidersInterface_ProviderInfoNameToIdMap_Call) RunAndReturn(run func(context.Context, marketplace.ListProvidersRequest) (map[string]string, error)) *MockConsumerProvidersInterface_ProviderInfoNameToIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockConsumerProvidersInterface creates a new instance of MockConsumerProvidersInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConsumerProvidersInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConsumerProvidersInterface {
	mock := &MockConsumerProvidersInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
