// Code generated by mockery v2.38.0. DO NOT EDIT.

package compute

import (
	context "context"

	compute "github.com/databricks/databricks-sdk-go/service/compute"

	listing "github.com/databricks/databricks-sdk-go/listing"

	mock "github.com/stretchr/testify/mock"
)

// MockPolicyFamiliesInterface is an autogenerated mock type for the PolicyFamiliesInterface type
type MockPolicyFamiliesInterface struct {
	mock.Mock
}

type MockPolicyFamiliesInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPolicyFamiliesInterface) EXPECT() *MockPolicyFamiliesInterface_Expecter {
	return &MockPolicyFamiliesInterface_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockPolicyFamiliesInterface) Get(ctx context.Context, request compute.GetPolicyFamilyRequest) (*compute.PolicyFamily, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *compute.PolicyFamily
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, compute.GetPolicyFamilyRequest) (*compute.PolicyFamily, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, compute.GetPolicyFamilyRequest) *compute.PolicyFamily); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.PolicyFamily)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, compute.GetPolicyFamilyRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPolicyFamiliesInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockPolicyFamiliesInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request compute.GetPolicyFamilyRequest
func (_e *MockPolicyFamiliesInterface_Expecter) Get(ctx interface{}, request interface{}) *MockPolicyFamiliesInterface_Get_Call {
	return &MockPolicyFamiliesInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockPolicyFamiliesInterface_Get_Call) Run(run func(ctx context.Context, request compute.GetPolicyFamilyRequest)) *MockPolicyFamiliesInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(compute.GetPolicyFamilyRequest))
	})
	return _c
}

func (_c *MockPolicyFamiliesInterface_Get_Call) Return(_a0 *compute.PolicyFamily, _a1 error) *MockPolicyFamiliesInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPolicyFamiliesInterface_Get_Call) RunAndReturn(run func(context.Context, compute.GetPolicyFamilyRequest) (*compute.PolicyFamily, error)) *MockPolicyFamiliesInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByPolicyFamilyId provides a mock function with given fields: ctx, policyFamilyId
func (_m *MockPolicyFamiliesInterface) GetByPolicyFamilyId(ctx context.Context, policyFamilyId string) (*compute.PolicyFamily, error) {
	ret := _m.Called(ctx, policyFamilyId)

	if len(ret) == 0 {
		panic("no return value specified for GetByPolicyFamilyId")
	}

	var r0 *compute.PolicyFamily
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*compute.PolicyFamily, error)); ok {
		return rf(ctx, policyFamilyId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *compute.PolicyFamily); ok {
		r0 = rf(ctx, policyFamilyId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.PolicyFamily)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, policyFamilyId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPolicyFamiliesInterface_GetByPolicyFamilyId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByPolicyFamilyId'
type MockPolicyFamiliesInterface_GetByPolicyFamilyId_Call struct {
	*mock.Call
}

// GetByPolicyFamilyId is a helper method to define mock.On call
//   - ctx context.Context
//   - policyFamilyId string
func (_e *MockPolicyFamiliesInterface_Expecter) GetByPolicyFamilyId(ctx interface{}, policyFamilyId interface{}) *MockPolicyFamiliesInterface_GetByPolicyFamilyId_Call {
	return &MockPolicyFamiliesInterface_GetByPolicyFamilyId_Call{Call: _e.mock.On("GetByPolicyFamilyId", ctx, policyFamilyId)}
}

func (_c *MockPolicyFamiliesInterface_GetByPolicyFamilyId_Call) Run(run func(ctx context.Context, policyFamilyId string)) *MockPolicyFamiliesInterface_GetByPolicyFamilyId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockPolicyFamiliesInterface_GetByPolicyFamilyId_Call) Return(_a0 *compute.PolicyFamily, _a1 error) *MockPolicyFamiliesInterface_GetByPolicyFamilyId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPolicyFamiliesInterface_GetByPolicyFamilyId_Call) RunAndReturn(run func(context.Context, string) (*compute.PolicyFamily, error)) *MockPolicyFamiliesInterface_GetByPolicyFamilyId_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockPolicyFamiliesInterface) Impl() compute.PolicyFamiliesService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 compute.PolicyFamiliesService
	if rf, ok := ret.Get(0).(func() compute.PolicyFamiliesService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(compute.PolicyFamiliesService)
		}
	}

	return r0
}

// MockPolicyFamiliesInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockPolicyFamiliesInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockPolicyFamiliesInterface_Expecter) Impl() *MockPolicyFamiliesInterface_Impl_Call {
	return &MockPolicyFamiliesInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockPolicyFamiliesInterface_Impl_Call) Run(run func()) *MockPolicyFamiliesInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPolicyFamiliesInterface_Impl_Call) Return(_a0 compute.PolicyFamiliesService) *MockPolicyFamiliesInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPolicyFamiliesInterface_Impl_Call) RunAndReturn(run func() compute.PolicyFamiliesService) *MockPolicyFamiliesInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockPolicyFamiliesInterface) List(ctx context.Context, request compute.ListPolicyFamiliesRequest) *listing.PaginatingIterator[compute.ListPolicyFamiliesRequest, *compute.ListPolicyFamiliesResponse, compute.PolicyFamily] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *listing.PaginatingIterator[compute.ListPolicyFamiliesRequest, *compute.ListPolicyFamiliesResponse, compute.PolicyFamily]
	if rf, ok := ret.Get(0).(func(context.Context, compute.ListPolicyFamiliesRequest) *listing.PaginatingIterator[compute.ListPolicyFamiliesRequest, *compute.ListPolicyFamiliesResponse, compute.PolicyFamily]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*listing.PaginatingIterator[compute.ListPolicyFamiliesRequest, *compute.ListPolicyFamiliesResponse, compute.PolicyFamily])
		}
	}

	return r0
}

// MockPolicyFamiliesInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockPolicyFamiliesInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request compute.ListPolicyFamiliesRequest
func (_e *MockPolicyFamiliesInterface_Expecter) List(ctx interface{}, request interface{}) *MockPolicyFamiliesInterface_List_Call {
	return &MockPolicyFamiliesInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockPolicyFamiliesInterface_List_Call) Run(run func(ctx context.Context, request compute.ListPolicyFamiliesRequest)) *MockPolicyFamiliesInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(compute.ListPolicyFamiliesRequest))
	})
	return _c
}

func (_c *MockPolicyFamiliesInterface_List_Call) Return(_a0 *listing.PaginatingIterator[compute.ListPolicyFamiliesRequest, *compute.ListPolicyFamiliesResponse, compute.PolicyFamily]) *MockPolicyFamiliesInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPolicyFamiliesInterface_List_Call) RunAndReturn(run func(context.Context, compute.ListPolicyFamiliesRequest) *listing.PaginatingIterator[compute.ListPolicyFamiliesRequest, *compute.ListPolicyFamiliesResponse, compute.PolicyFamily]) *MockPolicyFamiliesInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockPolicyFamiliesInterface) ListAll(ctx context.Context, request compute.ListPolicyFamiliesRequest) ([]compute.PolicyFamily, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []compute.PolicyFamily
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, compute.ListPolicyFamiliesRequest) ([]compute.PolicyFamily, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, compute.ListPolicyFamiliesRequest) []compute.PolicyFamily); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]compute.PolicyFamily)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, compute.ListPolicyFamiliesRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPolicyFamiliesInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockPolicyFamiliesInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request compute.ListPolicyFamiliesRequest
func (_e *MockPolicyFamiliesInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockPolicyFamiliesInterface_ListAll_Call {
	return &MockPolicyFamiliesInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockPolicyFamiliesInterface_ListAll_Call) Run(run func(ctx context.Context, request compute.ListPolicyFamiliesRequest)) *MockPolicyFamiliesInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(compute.ListPolicyFamiliesRequest))
	})
	return _c
}

func (_c *MockPolicyFamiliesInterface_ListAll_Call) Return(_a0 []compute.PolicyFamily, _a1 error) *MockPolicyFamiliesInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPolicyFamiliesInterface_ListAll_Call) RunAndReturn(run func(context.Context, compute.ListPolicyFamiliesRequest) ([]compute.PolicyFamily, error)) *MockPolicyFamiliesInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockPolicyFamiliesInterface) WithImpl(impl compute.PolicyFamiliesService) compute.PolicyFamiliesInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 compute.PolicyFamiliesInterface
	if rf, ok := ret.Get(0).(func(compute.PolicyFamiliesService) compute.PolicyFamiliesInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(compute.PolicyFamiliesInterface)
		}
	}

	return r0
}

// MockPolicyFamiliesInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockPolicyFamiliesInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl compute.PolicyFamiliesService
func (_e *MockPolicyFamiliesInterface_Expecter) WithImpl(impl interface{}) *MockPolicyFamiliesInterface_WithImpl_Call {
	return &MockPolicyFamiliesInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockPolicyFamiliesInterface_WithImpl_Call) Run(run func(impl compute.PolicyFamiliesService)) *MockPolicyFamiliesInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(compute.PolicyFamiliesService))
	})
	return _c
}

func (_c *MockPolicyFamiliesInterface_WithImpl_Call) Return(_a0 compute.PolicyFamiliesInterface) *MockPolicyFamiliesInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPolicyFamiliesInterface_WithImpl_Call) RunAndReturn(run func(compute.PolicyFamiliesService) compute.PolicyFamiliesInterface) *MockPolicyFamiliesInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPolicyFamiliesInterface creates a new instance of MockPolicyFamiliesInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPolicyFamiliesInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPolicyFamiliesInterface {
	mock := &MockPolicyFamiliesInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
