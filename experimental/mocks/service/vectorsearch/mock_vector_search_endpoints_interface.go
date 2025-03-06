// Code generated by mockery v2.53.0. DO NOT EDIT.

package vectorsearch

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	mock "github.com/stretchr/testify/mock"

	retries "github.com/databricks/databricks-sdk-go/retries"

	time "time"

	vectorsearch "github.com/databricks/databricks-sdk-go/service/vectorsearch"
)

// MockVectorSearchEndpointsInterface is an autogenerated mock type for the VectorSearchEndpointsInterface type
type MockVectorSearchEndpointsInterface struct {
	mock.Mock
}

type MockVectorSearchEndpointsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockVectorSearchEndpointsInterface) EXPECT() *MockVectorSearchEndpointsInterface_Expecter {
	return &MockVectorSearchEndpointsInterface_Expecter{mock: &_m.Mock}
}

// CreateEndpoint provides a mock function with given fields: ctx, createEndpoint
func (_m *MockVectorSearchEndpointsInterface) CreateEndpoint(ctx context.Context, createEndpoint vectorsearch.CreateEndpoint) (*vectorsearch.WaitGetEndpointVectorSearchEndpointOnline[vectorsearch.EndpointInfo], error) {
	ret := _m.Called(ctx, createEndpoint)

	if len(ret) == 0 {
		panic("no return value specified for CreateEndpoint")
	}

	var r0 *vectorsearch.WaitGetEndpointVectorSearchEndpointOnline[vectorsearch.EndpointInfo]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, vectorsearch.CreateEndpoint) (*vectorsearch.WaitGetEndpointVectorSearchEndpointOnline[vectorsearch.EndpointInfo], error)); ok {
		return rf(ctx, createEndpoint)
	}
	if rf, ok := ret.Get(0).(func(context.Context, vectorsearch.CreateEndpoint) *vectorsearch.WaitGetEndpointVectorSearchEndpointOnline[vectorsearch.EndpointInfo]); ok {
		r0 = rf(ctx, createEndpoint)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vectorsearch.WaitGetEndpointVectorSearchEndpointOnline[vectorsearch.EndpointInfo])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, vectorsearch.CreateEndpoint) error); ok {
		r1 = rf(ctx, createEndpoint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVectorSearchEndpointsInterface_CreateEndpoint_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateEndpoint'
type MockVectorSearchEndpointsInterface_CreateEndpoint_Call struct {
	*mock.Call
}

// CreateEndpoint is a helper method to define mock.On call
//   - ctx context.Context
//   - createEndpoint vectorsearch.CreateEndpoint
func (_e *MockVectorSearchEndpointsInterface_Expecter) CreateEndpoint(ctx interface{}, createEndpoint interface{}) *MockVectorSearchEndpointsInterface_CreateEndpoint_Call {
	return &MockVectorSearchEndpointsInterface_CreateEndpoint_Call{Call: _e.mock.On("CreateEndpoint", ctx, createEndpoint)}
}

func (_c *MockVectorSearchEndpointsInterface_CreateEndpoint_Call) Run(run func(ctx context.Context, createEndpoint vectorsearch.CreateEndpoint)) *MockVectorSearchEndpointsInterface_CreateEndpoint_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(vectorsearch.CreateEndpoint))
	})
	return _c
}

func (_c *MockVectorSearchEndpointsInterface_CreateEndpoint_Call) Return(_a0 *vectorsearch.WaitGetEndpointVectorSearchEndpointOnline[vectorsearch.EndpointInfo], _a1 error) *MockVectorSearchEndpointsInterface_CreateEndpoint_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVectorSearchEndpointsInterface_CreateEndpoint_Call) RunAndReturn(run func(context.Context, vectorsearch.CreateEndpoint) (*vectorsearch.WaitGetEndpointVectorSearchEndpointOnline[vectorsearch.EndpointInfo], error)) *MockVectorSearchEndpointsInterface_CreateEndpoint_Call {
	_c.Call.Return(run)
	return _c
}

// CreateEndpointAndWait provides a mock function with given fields: ctx, createEndpoint, options
func (_m *MockVectorSearchEndpointsInterface) CreateEndpointAndWait(ctx context.Context, createEndpoint vectorsearch.CreateEndpoint, options ...retries.Option[vectorsearch.EndpointInfo]) (*vectorsearch.EndpointInfo, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, createEndpoint)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateEndpointAndWait")
	}

	var r0 *vectorsearch.EndpointInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, vectorsearch.CreateEndpoint, ...retries.Option[vectorsearch.EndpointInfo]) (*vectorsearch.EndpointInfo, error)); ok {
		return rf(ctx, createEndpoint, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, vectorsearch.CreateEndpoint, ...retries.Option[vectorsearch.EndpointInfo]) *vectorsearch.EndpointInfo); ok {
		r0 = rf(ctx, createEndpoint, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vectorsearch.EndpointInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, vectorsearch.CreateEndpoint, ...retries.Option[vectorsearch.EndpointInfo]) error); ok {
		r1 = rf(ctx, createEndpoint, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVectorSearchEndpointsInterface_CreateEndpointAndWait_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateEndpointAndWait'
type MockVectorSearchEndpointsInterface_CreateEndpointAndWait_Call struct {
	*mock.Call
}

// CreateEndpointAndWait is a helper method to define mock.On call
//   - ctx context.Context
//   - createEndpoint vectorsearch.CreateEndpoint
//   - options ...retries.Option[vectorsearch.EndpointInfo]
func (_e *MockVectorSearchEndpointsInterface_Expecter) CreateEndpointAndWait(ctx interface{}, createEndpoint interface{}, options ...interface{}) *MockVectorSearchEndpointsInterface_CreateEndpointAndWait_Call {
	return &MockVectorSearchEndpointsInterface_CreateEndpointAndWait_Call{Call: _e.mock.On("CreateEndpointAndWait",
		append([]interface{}{ctx, createEndpoint}, options...)...)}
}

func (_c *MockVectorSearchEndpointsInterface_CreateEndpointAndWait_Call) Run(run func(ctx context.Context, createEndpoint vectorsearch.CreateEndpoint, options ...retries.Option[vectorsearch.EndpointInfo])) *MockVectorSearchEndpointsInterface_CreateEndpointAndWait_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]retries.Option[vectorsearch.EndpointInfo], len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(retries.Option[vectorsearch.EndpointInfo])
			}
		}
		run(args[0].(context.Context), args[1].(vectorsearch.CreateEndpoint), variadicArgs...)
	})
	return _c
}

func (_c *MockVectorSearchEndpointsInterface_CreateEndpointAndWait_Call) Return(_a0 *vectorsearch.EndpointInfo, _a1 error) *MockVectorSearchEndpointsInterface_CreateEndpointAndWait_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVectorSearchEndpointsInterface_CreateEndpointAndWait_Call) RunAndReturn(run func(context.Context, vectorsearch.CreateEndpoint, ...retries.Option[vectorsearch.EndpointInfo]) (*vectorsearch.EndpointInfo, error)) *MockVectorSearchEndpointsInterface_CreateEndpointAndWait_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteEndpoint provides a mock function with given fields: ctx, request
func (_m *MockVectorSearchEndpointsInterface) DeleteEndpoint(ctx context.Context, request vectorsearch.DeleteEndpointRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for DeleteEndpoint")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, vectorsearch.DeleteEndpointRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockVectorSearchEndpointsInterface_DeleteEndpoint_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteEndpoint'
type MockVectorSearchEndpointsInterface_DeleteEndpoint_Call struct {
	*mock.Call
}

// DeleteEndpoint is a helper method to define mock.On call
//   - ctx context.Context
//   - request vectorsearch.DeleteEndpointRequest
func (_e *MockVectorSearchEndpointsInterface_Expecter) DeleteEndpoint(ctx interface{}, request interface{}) *MockVectorSearchEndpointsInterface_DeleteEndpoint_Call {
	return &MockVectorSearchEndpointsInterface_DeleteEndpoint_Call{Call: _e.mock.On("DeleteEndpoint", ctx, request)}
}

func (_c *MockVectorSearchEndpointsInterface_DeleteEndpoint_Call) Run(run func(ctx context.Context, request vectorsearch.DeleteEndpointRequest)) *MockVectorSearchEndpointsInterface_DeleteEndpoint_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(vectorsearch.DeleteEndpointRequest))
	})
	return _c
}

func (_c *MockVectorSearchEndpointsInterface_DeleteEndpoint_Call) Return(_a0 error) *MockVectorSearchEndpointsInterface_DeleteEndpoint_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockVectorSearchEndpointsInterface_DeleteEndpoint_Call) RunAndReturn(run func(context.Context, vectorsearch.DeleteEndpointRequest) error) *MockVectorSearchEndpointsInterface_DeleteEndpoint_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteEndpointByEndpointName provides a mock function with given fields: ctx, endpointName
func (_m *MockVectorSearchEndpointsInterface) DeleteEndpointByEndpointName(ctx context.Context, endpointName string) error {
	ret := _m.Called(ctx, endpointName)

	if len(ret) == 0 {
		panic("no return value specified for DeleteEndpointByEndpointName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, endpointName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockVectorSearchEndpointsInterface_DeleteEndpointByEndpointName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteEndpointByEndpointName'
type MockVectorSearchEndpointsInterface_DeleteEndpointByEndpointName_Call struct {
	*mock.Call
}

// DeleteEndpointByEndpointName is a helper method to define mock.On call
//   - ctx context.Context
//   - endpointName string
func (_e *MockVectorSearchEndpointsInterface_Expecter) DeleteEndpointByEndpointName(ctx interface{}, endpointName interface{}) *MockVectorSearchEndpointsInterface_DeleteEndpointByEndpointName_Call {
	return &MockVectorSearchEndpointsInterface_DeleteEndpointByEndpointName_Call{Call: _e.mock.On("DeleteEndpointByEndpointName", ctx, endpointName)}
}

func (_c *MockVectorSearchEndpointsInterface_DeleteEndpointByEndpointName_Call) Run(run func(ctx context.Context, endpointName string)) *MockVectorSearchEndpointsInterface_DeleteEndpointByEndpointName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockVectorSearchEndpointsInterface_DeleteEndpointByEndpointName_Call) Return(_a0 error) *MockVectorSearchEndpointsInterface_DeleteEndpointByEndpointName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockVectorSearchEndpointsInterface_DeleteEndpointByEndpointName_Call) RunAndReturn(run func(context.Context, string) error) *MockVectorSearchEndpointsInterface_DeleteEndpointByEndpointName_Call {
	_c.Call.Return(run)
	return _c
}

// GetEndpoint provides a mock function with given fields: ctx, request
func (_m *MockVectorSearchEndpointsInterface) GetEndpoint(ctx context.Context, request vectorsearch.GetEndpointRequest) (*vectorsearch.EndpointInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetEndpoint")
	}

	var r0 *vectorsearch.EndpointInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, vectorsearch.GetEndpointRequest) (*vectorsearch.EndpointInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, vectorsearch.GetEndpointRequest) *vectorsearch.EndpointInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vectorsearch.EndpointInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, vectorsearch.GetEndpointRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVectorSearchEndpointsInterface_GetEndpoint_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEndpoint'
type MockVectorSearchEndpointsInterface_GetEndpoint_Call struct {
	*mock.Call
}

// GetEndpoint is a helper method to define mock.On call
//   - ctx context.Context
//   - request vectorsearch.GetEndpointRequest
func (_e *MockVectorSearchEndpointsInterface_Expecter) GetEndpoint(ctx interface{}, request interface{}) *MockVectorSearchEndpointsInterface_GetEndpoint_Call {
	return &MockVectorSearchEndpointsInterface_GetEndpoint_Call{Call: _e.mock.On("GetEndpoint", ctx, request)}
}

func (_c *MockVectorSearchEndpointsInterface_GetEndpoint_Call) Run(run func(ctx context.Context, request vectorsearch.GetEndpointRequest)) *MockVectorSearchEndpointsInterface_GetEndpoint_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(vectorsearch.GetEndpointRequest))
	})
	return _c
}

func (_c *MockVectorSearchEndpointsInterface_GetEndpoint_Call) Return(_a0 *vectorsearch.EndpointInfo, _a1 error) *MockVectorSearchEndpointsInterface_GetEndpoint_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVectorSearchEndpointsInterface_GetEndpoint_Call) RunAndReturn(run func(context.Context, vectorsearch.GetEndpointRequest) (*vectorsearch.EndpointInfo, error)) *MockVectorSearchEndpointsInterface_GetEndpoint_Call {
	_c.Call.Return(run)
	return _c
}

// GetEndpointByEndpointName provides a mock function with given fields: ctx, endpointName
func (_m *MockVectorSearchEndpointsInterface) GetEndpointByEndpointName(ctx context.Context, endpointName string) (*vectorsearch.EndpointInfo, error) {
	ret := _m.Called(ctx, endpointName)

	if len(ret) == 0 {
		panic("no return value specified for GetEndpointByEndpointName")
	}

	var r0 *vectorsearch.EndpointInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*vectorsearch.EndpointInfo, error)); ok {
		return rf(ctx, endpointName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *vectorsearch.EndpointInfo); ok {
		r0 = rf(ctx, endpointName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vectorsearch.EndpointInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, endpointName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVectorSearchEndpointsInterface_GetEndpointByEndpointName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEndpointByEndpointName'
type MockVectorSearchEndpointsInterface_GetEndpointByEndpointName_Call struct {
	*mock.Call
}

// GetEndpointByEndpointName is a helper method to define mock.On call
//   - ctx context.Context
//   - endpointName string
func (_e *MockVectorSearchEndpointsInterface_Expecter) GetEndpointByEndpointName(ctx interface{}, endpointName interface{}) *MockVectorSearchEndpointsInterface_GetEndpointByEndpointName_Call {
	return &MockVectorSearchEndpointsInterface_GetEndpointByEndpointName_Call{Call: _e.mock.On("GetEndpointByEndpointName", ctx, endpointName)}
}

func (_c *MockVectorSearchEndpointsInterface_GetEndpointByEndpointName_Call) Run(run func(ctx context.Context, endpointName string)) *MockVectorSearchEndpointsInterface_GetEndpointByEndpointName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockVectorSearchEndpointsInterface_GetEndpointByEndpointName_Call) Return(_a0 *vectorsearch.EndpointInfo, _a1 error) *MockVectorSearchEndpointsInterface_GetEndpointByEndpointName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVectorSearchEndpointsInterface_GetEndpointByEndpointName_Call) RunAndReturn(run func(context.Context, string) (*vectorsearch.EndpointInfo, error)) *MockVectorSearchEndpointsInterface_GetEndpointByEndpointName_Call {
	_c.Call.Return(run)
	return _c
}

// ListEndpoints provides a mock function with given fields: ctx, request
func (_m *MockVectorSearchEndpointsInterface) ListEndpoints(ctx context.Context, request vectorsearch.ListEndpointsRequest) listing.Iterator[vectorsearch.EndpointInfo] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListEndpoints")
	}

	var r0 listing.Iterator[vectorsearch.EndpointInfo]
	if rf, ok := ret.Get(0).(func(context.Context, vectorsearch.ListEndpointsRequest) listing.Iterator[vectorsearch.EndpointInfo]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[vectorsearch.EndpointInfo])
		}
	}

	return r0
}

// MockVectorSearchEndpointsInterface_ListEndpoints_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListEndpoints'
type MockVectorSearchEndpointsInterface_ListEndpoints_Call struct {
	*mock.Call
}

// ListEndpoints is a helper method to define mock.On call
//   - ctx context.Context
//   - request vectorsearch.ListEndpointsRequest
func (_e *MockVectorSearchEndpointsInterface_Expecter) ListEndpoints(ctx interface{}, request interface{}) *MockVectorSearchEndpointsInterface_ListEndpoints_Call {
	return &MockVectorSearchEndpointsInterface_ListEndpoints_Call{Call: _e.mock.On("ListEndpoints", ctx, request)}
}

func (_c *MockVectorSearchEndpointsInterface_ListEndpoints_Call) Run(run func(ctx context.Context, request vectorsearch.ListEndpointsRequest)) *MockVectorSearchEndpointsInterface_ListEndpoints_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(vectorsearch.ListEndpointsRequest))
	})
	return _c
}

func (_c *MockVectorSearchEndpointsInterface_ListEndpoints_Call) Return(_a0 listing.Iterator[vectorsearch.EndpointInfo]) *MockVectorSearchEndpointsInterface_ListEndpoints_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockVectorSearchEndpointsInterface_ListEndpoints_Call) RunAndReturn(run func(context.Context, vectorsearch.ListEndpointsRequest) listing.Iterator[vectorsearch.EndpointInfo]) *MockVectorSearchEndpointsInterface_ListEndpoints_Call {
	_c.Call.Return(run)
	return _c
}

// ListEndpointsAll provides a mock function with given fields: ctx, request
func (_m *MockVectorSearchEndpointsInterface) ListEndpointsAll(ctx context.Context, request vectorsearch.ListEndpointsRequest) ([]vectorsearch.EndpointInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListEndpointsAll")
	}

	var r0 []vectorsearch.EndpointInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, vectorsearch.ListEndpointsRequest) ([]vectorsearch.EndpointInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, vectorsearch.ListEndpointsRequest) []vectorsearch.EndpointInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]vectorsearch.EndpointInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, vectorsearch.ListEndpointsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVectorSearchEndpointsInterface_ListEndpointsAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListEndpointsAll'
type MockVectorSearchEndpointsInterface_ListEndpointsAll_Call struct {
	*mock.Call
}

// ListEndpointsAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request vectorsearch.ListEndpointsRequest
func (_e *MockVectorSearchEndpointsInterface_Expecter) ListEndpointsAll(ctx interface{}, request interface{}) *MockVectorSearchEndpointsInterface_ListEndpointsAll_Call {
	return &MockVectorSearchEndpointsInterface_ListEndpointsAll_Call{Call: _e.mock.On("ListEndpointsAll", ctx, request)}
}

func (_c *MockVectorSearchEndpointsInterface_ListEndpointsAll_Call) Run(run func(ctx context.Context, request vectorsearch.ListEndpointsRequest)) *MockVectorSearchEndpointsInterface_ListEndpointsAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(vectorsearch.ListEndpointsRequest))
	})
	return _c
}

func (_c *MockVectorSearchEndpointsInterface_ListEndpointsAll_Call) Return(_a0 []vectorsearch.EndpointInfo, _a1 error) *MockVectorSearchEndpointsInterface_ListEndpointsAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVectorSearchEndpointsInterface_ListEndpointsAll_Call) RunAndReturn(run func(context.Context, vectorsearch.ListEndpointsRequest) ([]vectorsearch.EndpointInfo, error)) *MockVectorSearchEndpointsInterface_ListEndpointsAll_Call {
	_c.Call.Return(run)
	return _c
}

// WaitGetEndpointVectorSearchEndpointOnline provides a mock function with given fields: ctx, endpointName, timeout, callback
func (_m *MockVectorSearchEndpointsInterface) WaitGetEndpointVectorSearchEndpointOnline(ctx context.Context, endpointName string, timeout time.Duration, callback func(*vectorsearch.EndpointInfo)) (*vectorsearch.EndpointInfo, error) {
	ret := _m.Called(ctx, endpointName, timeout, callback)

	if len(ret) == 0 {
		panic("no return value specified for WaitGetEndpointVectorSearchEndpointOnline")
	}

	var r0 *vectorsearch.EndpointInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration, func(*vectorsearch.EndpointInfo)) (*vectorsearch.EndpointInfo, error)); ok {
		return rf(ctx, endpointName, timeout, callback)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration, func(*vectorsearch.EndpointInfo)) *vectorsearch.EndpointInfo); ok {
		r0 = rf(ctx, endpointName, timeout, callback)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vectorsearch.EndpointInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, time.Duration, func(*vectorsearch.EndpointInfo)) error); ok {
		r1 = rf(ctx, endpointName, timeout, callback)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVectorSearchEndpointsInterface_WaitGetEndpointVectorSearchEndpointOnline_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WaitGetEndpointVectorSearchEndpointOnline'
type MockVectorSearchEndpointsInterface_WaitGetEndpointVectorSearchEndpointOnline_Call struct {
	*mock.Call
}

// WaitGetEndpointVectorSearchEndpointOnline is a helper method to define mock.On call
//   - ctx context.Context
//   - endpointName string
//   - timeout time.Duration
//   - callback func(*vectorsearch.EndpointInfo)
func (_e *MockVectorSearchEndpointsInterface_Expecter) WaitGetEndpointVectorSearchEndpointOnline(ctx interface{}, endpointName interface{}, timeout interface{}, callback interface{}) *MockVectorSearchEndpointsInterface_WaitGetEndpointVectorSearchEndpointOnline_Call {
	return &MockVectorSearchEndpointsInterface_WaitGetEndpointVectorSearchEndpointOnline_Call{Call: _e.mock.On("WaitGetEndpointVectorSearchEndpointOnline", ctx, endpointName, timeout, callback)}
}

func (_c *MockVectorSearchEndpointsInterface_WaitGetEndpointVectorSearchEndpointOnline_Call) Run(run func(ctx context.Context, endpointName string, timeout time.Duration, callback func(*vectorsearch.EndpointInfo))) *MockVectorSearchEndpointsInterface_WaitGetEndpointVectorSearchEndpointOnline_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(time.Duration), args[3].(func(*vectorsearch.EndpointInfo)))
	})
	return _c
}

func (_c *MockVectorSearchEndpointsInterface_WaitGetEndpointVectorSearchEndpointOnline_Call) Return(_a0 *vectorsearch.EndpointInfo, _a1 error) *MockVectorSearchEndpointsInterface_WaitGetEndpointVectorSearchEndpointOnline_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVectorSearchEndpointsInterface_WaitGetEndpointVectorSearchEndpointOnline_Call) RunAndReturn(run func(context.Context, string, time.Duration, func(*vectorsearch.EndpointInfo)) (*vectorsearch.EndpointInfo, error)) *MockVectorSearchEndpointsInterface_WaitGetEndpointVectorSearchEndpointOnline_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockVectorSearchEndpointsInterface creates a new instance of MockVectorSearchEndpointsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockVectorSearchEndpointsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockVectorSearchEndpointsInterface {
	mock := &MockVectorSearchEndpointsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
