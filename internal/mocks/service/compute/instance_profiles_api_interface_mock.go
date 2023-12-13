// Code generated by mockery v2.38.0. DO NOT EDIT.

package compute

import (
	context "context"

	compute "github.com/databricks/databricks-sdk-go/service/compute"

	listing "github.com/databricks/databricks-sdk-go/listing"

	mock "github.com/stretchr/testify/mock"
)

// MockInstanceProfilesAPIInterface is an autogenerated mock type for the InstanceProfilesAPIInterface type
type MockInstanceProfilesAPIInterface struct {
	mock.Mock
}

type MockInstanceProfilesAPIInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInstanceProfilesAPIInterface) EXPECT() *MockInstanceProfilesAPIInterface_Expecter {
	return &MockInstanceProfilesAPIInterface_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: ctx, request
func (_m *MockInstanceProfilesAPIInterface) Add(ctx context.Context, request compute.AddInstanceProfile) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, compute.AddInstanceProfile) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockInstanceProfilesAPIInterface_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type MockInstanceProfilesAPIInterface_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - ctx context.Context
//   - request compute.AddInstanceProfile
func (_e *MockInstanceProfilesAPIInterface_Expecter) Add(ctx interface{}, request interface{}) *MockInstanceProfilesAPIInterface_Add_Call {
	return &MockInstanceProfilesAPIInterface_Add_Call{Call: _e.mock.On("Add", ctx, request)}
}

func (_c *MockInstanceProfilesAPIInterface_Add_Call) Run(run func(ctx context.Context, request compute.AddInstanceProfile)) *MockInstanceProfilesAPIInterface_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(compute.AddInstanceProfile))
	})
	return _c
}

func (_c *MockInstanceProfilesAPIInterface_Add_Call) Return(_a0 error) *MockInstanceProfilesAPIInterface_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInstanceProfilesAPIInterface_Add_Call) RunAndReturn(run func(context.Context, compute.AddInstanceProfile) error) *MockInstanceProfilesAPIInterface_Add_Call {
	_c.Call.Return(run)
	return _c
}

// Edit provides a mock function with given fields: ctx, request
func (_m *MockInstanceProfilesAPIInterface) Edit(ctx context.Context, request compute.InstanceProfile) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Edit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, compute.InstanceProfile) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockInstanceProfilesAPIInterface_Edit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Edit'
type MockInstanceProfilesAPIInterface_Edit_Call struct {
	*mock.Call
}

// Edit is a helper method to define mock.On call
//   - ctx context.Context
//   - request compute.InstanceProfile
func (_e *MockInstanceProfilesAPIInterface_Expecter) Edit(ctx interface{}, request interface{}) *MockInstanceProfilesAPIInterface_Edit_Call {
	return &MockInstanceProfilesAPIInterface_Edit_Call{Call: _e.mock.On("Edit", ctx, request)}
}

func (_c *MockInstanceProfilesAPIInterface_Edit_Call) Run(run func(ctx context.Context, request compute.InstanceProfile)) *MockInstanceProfilesAPIInterface_Edit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(compute.InstanceProfile))
	})
	return _c
}

func (_c *MockInstanceProfilesAPIInterface_Edit_Call) Return(_a0 error) *MockInstanceProfilesAPIInterface_Edit_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInstanceProfilesAPIInterface_Edit_Call) RunAndReturn(run func(context.Context, compute.InstanceProfile) error) *MockInstanceProfilesAPIInterface_Edit_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockInstanceProfilesAPIInterface) Impl() compute.InstanceProfilesService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 compute.InstanceProfilesService
	if rf, ok := ret.Get(0).(func() compute.InstanceProfilesService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(compute.InstanceProfilesService)
		}
	}

	return r0
}

// MockInstanceProfilesAPIInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockInstanceProfilesAPIInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockInstanceProfilesAPIInterface_Expecter) Impl() *MockInstanceProfilesAPIInterface_Impl_Call {
	return &MockInstanceProfilesAPIInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockInstanceProfilesAPIInterface_Impl_Call) Run(run func()) *MockInstanceProfilesAPIInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInstanceProfilesAPIInterface_Impl_Call) Return(_a0 compute.InstanceProfilesService) *MockInstanceProfilesAPIInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInstanceProfilesAPIInterface_Impl_Call) RunAndReturn(run func() compute.InstanceProfilesService) *MockInstanceProfilesAPIInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx
func (_m *MockInstanceProfilesAPIInterface) List(ctx context.Context) *listing.PaginatingIterator[struct{}, *compute.ListInstanceProfilesResponse, compute.InstanceProfile] {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *listing.PaginatingIterator[struct{}, *compute.ListInstanceProfilesResponse, compute.InstanceProfile]
	if rf, ok := ret.Get(0).(func(context.Context) *listing.PaginatingIterator[struct{}, *compute.ListInstanceProfilesResponse, compute.InstanceProfile]); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*listing.PaginatingIterator[struct{}, *compute.ListInstanceProfilesResponse, compute.InstanceProfile])
		}
	}

	return r0
}

// MockInstanceProfilesAPIInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockInstanceProfilesAPIInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockInstanceProfilesAPIInterface_Expecter) List(ctx interface{}) *MockInstanceProfilesAPIInterface_List_Call {
	return &MockInstanceProfilesAPIInterface_List_Call{Call: _e.mock.On("List", ctx)}
}

func (_c *MockInstanceProfilesAPIInterface_List_Call) Run(run func(ctx context.Context)) *MockInstanceProfilesAPIInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockInstanceProfilesAPIInterface_List_Call) Return(_a0 *listing.PaginatingIterator[struct{}, *compute.ListInstanceProfilesResponse, compute.InstanceProfile]) *MockInstanceProfilesAPIInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInstanceProfilesAPIInterface_List_Call) RunAndReturn(run func(context.Context) *listing.PaginatingIterator[struct{}, *compute.ListInstanceProfilesResponse, compute.InstanceProfile]) *MockInstanceProfilesAPIInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx
func (_m *MockInstanceProfilesAPIInterface) ListAll(ctx context.Context) ([]compute.InstanceProfile, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []compute.InstanceProfile
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]compute.InstanceProfile, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []compute.InstanceProfile); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]compute.InstanceProfile)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInstanceProfilesAPIInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockInstanceProfilesAPIInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockInstanceProfilesAPIInterface_Expecter) ListAll(ctx interface{}) *MockInstanceProfilesAPIInterface_ListAll_Call {
	return &MockInstanceProfilesAPIInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx)}
}

func (_c *MockInstanceProfilesAPIInterface_ListAll_Call) Run(run func(ctx context.Context)) *MockInstanceProfilesAPIInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockInstanceProfilesAPIInterface_ListAll_Call) Return(_a0 []compute.InstanceProfile, _a1 error) *MockInstanceProfilesAPIInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInstanceProfilesAPIInterface_ListAll_Call) RunAndReturn(run func(context.Context) ([]compute.InstanceProfile, error)) *MockInstanceProfilesAPIInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// Remove provides a mock function with given fields: ctx, request
func (_m *MockInstanceProfilesAPIInterface) Remove(ctx context.Context, request compute.RemoveInstanceProfile) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Remove")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, compute.RemoveInstanceProfile) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockInstanceProfilesAPIInterface_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type MockInstanceProfilesAPIInterface_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//   - ctx context.Context
//   - request compute.RemoveInstanceProfile
func (_e *MockInstanceProfilesAPIInterface_Expecter) Remove(ctx interface{}, request interface{}) *MockInstanceProfilesAPIInterface_Remove_Call {
	return &MockInstanceProfilesAPIInterface_Remove_Call{Call: _e.mock.On("Remove", ctx, request)}
}

func (_c *MockInstanceProfilesAPIInterface_Remove_Call) Run(run func(ctx context.Context, request compute.RemoveInstanceProfile)) *MockInstanceProfilesAPIInterface_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(compute.RemoveInstanceProfile))
	})
	return _c
}

func (_c *MockInstanceProfilesAPIInterface_Remove_Call) Return(_a0 error) *MockInstanceProfilesAPIInterface_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInstanceProfilesAPIInterface_Remove_Call) RunAndReturn(run func(context.Context, compute.RemoveInstanceProfile) error) *MockInstanceProfilesAPIInterface_Remove_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveByInstanceProfileArn provides a mock function with given fields: ctx, instanceProfileArn
func (_m *MockInstanceProfilesAPIInterface) RemoveByInstanceProfileArn(ctx context.Context, instanceProfileArn string) error {
	ret := _m.Called(ctx, instanceProfileArn)

	if len(ret) == 0 {
		panic("no return value specified for RemoveByInstanceProfileArn")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, instanceProfileArn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockInstanceProfilesAPIInterface_RemoveByInstanceProfileArn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveByInstanceProfileArn'
type MockInstanceProfilesAPIInterface_RemoveByInstanceProfileArn_Call struct {
	*mock.Call
}

// RemoveByInstanceProfileArn is a helper method to define mock.On call
//   - ctx context.Context
//   - instanceProfileArn string
func (_e *MockInstanceProfilesAPIInterface_Expecter) RemoveByInstanceProfileArn(ctx interface{}, instanceProfileArn interface{}) *MockInstanceProfilesAPIInterface_RemoveByInstanceProfileArn_Call {
	return &MockInstanceProfilesAPIInterface_RemoveByInstanceProfileArn_Call{Call: _e.mock.On("RemoveByInstanceProfileArn", ctx, instanceProfileArn)}
}

func (_c *MockInstanceProfilesAPIInterface_RemoveByInstanceProfileArn_Call) Run(run func(ctx context.Context, instanceProfileArn string)) *MockInstanceProfilesAPIInterface_RemoveByInstanceProfileArn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockInstanceProfilesAPIInterface_RemoveByInstanceProfileArn_Call) Return(_a0 error) *MockInstanceProfilesAPIInterface_RemoveByInstanceProfileArn_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInstanceProfilesAPIInterface_RemoveByInstanceProfileArn_Call) RunAndReturn(run func(context.Context, string) error) *MockInstanceProfilesAPIInterface_RemoveByInstanceProfileArn_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockInstanceProfilesAPIInterface) WithImpl(impl compute.InstanceProfilesService) compute.InstanceProfilesAPIInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 compute.InstanceProfilesAPIInterface
	if rf, ok := ret.Get(0).(func(compute.InstanceProfilesService) compute.InstanceProfilesAPIInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(compute.InstanceProfilesAPIInterface)
		}
	}

	return r0
}

// MockInstanceProfilesAPIInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockInstanceProfilesAPIInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl compute.InstanceProfilesService
func (_e *MockInstanceProfilesAPIInterface_Expecter) WithImpl(impl interface{}) *MockInstanceProfilesAPIInterface_WithImpl_Call {
	return &MockInstanceProfilesAPIInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockInstanceProfilesAPIInterface_WithImpl_Call) Run(run func(impl compute.InstanceProfilesService)) *MockInstanceProfilesAPIInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(compute.InstanceProfilesService))
	})
	return _c
}

func (_c *MockInstanceProfilesAPIInterface_WithImpl_Call) Return(_a0 compute.InstanceProfilesAPIInterface) *MockInstanceProfilesAPIInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInstanceProfilesAPIInterface_WithImpl_Call) RunAndReturn(run func(compute.InstanceProfilesService) compute.InstanceProfilesAPIInterface) *MockInstanceProfilesAPIInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockInstanceProfilesAPIInterface creates a new instance of MockInstanceProfilesAPIInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInstanceProfilesAPIInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInstanceProfilesAPIInterface {
	mock := &MockInstanceProfilesAPIInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
