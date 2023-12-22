// Code generated by mockery v2.38.0. DO NOT EDIT.

package catalog

import (
	context "context"

	catalog "github.com/databricks/databricks-sdk-go/service/catalog"

	mock "github.com/stretchr/testify/mock"
)

// MockTableConstraintsInterface is an autogenerated mock type for the TableConstraintsInterface type
type MockTableConstraintsInterface struct {
	mock.Mock
}

type MockTableConstraintsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTableConstraintsInterface) EXPECT() *MockTableConstraintsInterface_Expecter {
	return &MockTableConstraintsInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockTableConstraintsInterface) Create(ctx context.Context, request catalog.CreateTableConstraint) (*catalog.TableConstraint, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *catalog.TableConstraint
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateTableConstraint) (*catalog.TableConstraint, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateTableConstraint) *catalog.TableConstraint); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.TableConstraint)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.CreateTableConstraint) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTableConstraintsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockTableConstraintsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.CreateTableConstraint
func (_e *MockTableConstraintsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockTableConstraintsInterface_Create_Call {
	return &MockTableConstraintsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockTableConstraintsInterface_Create_Call) Run(run func(ctx context.Context, request catalog.CreateTableConstraint)) *MockTableConstraintsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.CreateTableConstraint))
	})
	return _c
}

func (_c *MockTableConstraintsInterface_Create_Call) Return(_a0 *catalog.TableConstraint, _a1 error) *MockTableConstraintsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTableConstraintsInterface_Create_Call) RunAndReturn(run func(context.Context, catalog.CreateTableConstraint) (*catalog.TableConstraint, error)) *MockTableConstraintsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockTableConstraintsInterface) Delete(ctx context.Context, request catalog.DeleteTableConstraintRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.DeleteTableConstraintRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTableConstraintsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockTableConstraintsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.DeleteTableConstraintRequest
func (_e *MockTableConstraintsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockTableConstraintsInterface_Delete_Call {
	return &MockTableConstraintsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockTableConstraintsInterface_Delete_Call) Run(run func(ctx context.Context, request catalog.DeleteTableConstraintRequest)) *MockTableConstraintsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.DeleteTableConstraintRequest))
	})
	return _c
}

func (_c *MockTableConstraintsInterface_Delete_Call) Return(_a0 error) *MockTableConstraintsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTableConstraintsInterface_Delete_Call) RunAndReturn(run func(context.Context, catalog.DeleteTableConstraintRequest) error) *MockTableConstraintsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByFullName provides a mock function with given fields: ctx, fullName
func (_m *MockTableConstraintsInterface) DeleteByFullName(ctx context.Context, fullName string) error {
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

// MockTableConstraintsInterface_DeleteByFullName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByFullName'
type MockTableConstraintsInterface_DeleteByFullName_Call struct {
	*mock.Call
}

// DeleteByFullName is a helper method to define mock.On call
//   - ctx context.Context
//   - fullName string
func (_e *MockTableConstraintsInterface_Expecter) DeleteByFullName(ctx interface{}, fullName interface{}) *MockTableConstraintsInterface_DeleteByFullName_Call {
	return &MockTableConstraintsInterface_DeleteByFullName_Call{Call: _e.mock.On("DeleteByFullName", ctx, fullName)}
}

func (_c *MockTableConstraintsInterface_DeleteByFullName_Call) Run(run func(ctx context.Context, fullName string)) *MockTableConstraintsInterface_DeleteByFullName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockTableConstraintsInterface_DeleteByFullName_Call) Return(_a0 error) *MockTableConstraintsInterface_DeleteByFullName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTableConstraintsInterface_DeleteByFullName_Call) RunAndReturn(run func(context.Context, string) error) *MockTableConstraintsInterface_DeleteByFullName_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockTableConstraintsInterface) Impl() catalog.TableConstraintsService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 catalog.TableConstraintsService
	if rf, ok := ret.Get(0).(func() catalog.TableConstraintsService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.TableConstraintsService)
		}
	}

	return r0
}

// MockTableConstraintsInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockTableConstraintsInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockTableConstraintsInterface_Expecter) Impl() *MockTableConstraintsInterface_Impl_Call {
	return &MockTableConstraintsInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockTableConstraintsInterface_Impl_Call) Run(run func()) *MockTableConstraintsInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTableConstraintsInterface_Impl_Call) Return(_a0 catalog.TableConstraintsService) *MockTableConstraintsInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTableConstraintsInterface_Impl_Call) RunAndReturn(run func() catalog.TableConstraintsService) *MockTableConstraintsInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockTableConstraintsInterface) WithImpl(impl catalog.TableConstraintsService) catalog.TableConstraintsInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 catalog.TableConstraintsInterface
	if rf, ok := ret.Get(0).(func(catalog.TableConstraintsService) catalog.TableConstraintsInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.TableConstraintsInterface)
		}
	}

	return r0
}

// MockTableConstraintsInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockTableConstraintsInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl catalog.TableConstraintsService
func (_e *MockTableConstraintsInterface_Expecter) WithImpl(impl interface{}) *MockTableConstraintsInterface_WithImpl_Call {
	return &MockTableConstraintsInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockTableConstraintsInterface_WithImpl_Call) Run(run func(impl catalog.TableConstraintsService)) *MockTableConstraintsInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(catalog.TableConstraintsService))
	})
	return _c
}

func (_c *MockTableConstraintsInterface_WithImpl_Call) Return(_a0 catalog.TableConstraintsInterface) *MockTableConstraintsInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTableConstraintsInterface_WithImpl_Call) RunAndReturn(run func(catalog.TableConstraintsService) catalog.TableConstraintsInterface) *MockTableConstraintsInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTableConstraintsInterface creates a new instance of MockTableConstraintsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTableConstraintsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTableConstraintsInterface {
	mock := &MockTableConstraintsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
