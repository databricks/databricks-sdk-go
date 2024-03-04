// Code generated by mockery v2.39.1. DO NOT EDIT.

package iam

import (
	context "context"

	iam "github.com/databricks/databricks-sdk-go/service/iam"
	mock "github.com/stretchr/testify/mock"
)

// MockPermissionMigrationInterface is an autogenerated mock type for the PermissionMigrationInterface type
type MockPermissionMigrationInterface struct {
	mock.Mock
}

type MockPermissionMigrationInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPermissionMigrationInterface) EXPECT() *MockPermissionMigrationInterface_Expecter {
	return &MockPermissionMigrationInterface_Expecter{mock: &_m.Mock}
}

// Impl provides a mock function with given fields:
func (_m *MockPermissionMigrationInterface) Impl() iam.PermissionMigrationService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 iam.PermissionMigrationService
	if rf, ok := ret.Get(0).(func() iam.PermissionMigrationService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(iam.PermissionMigrationService)
		}
	}

	return r0
}

// MockPermissionMigrationInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockPermissionMigrationInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockPermissionMigrationInterface_Expecter) Impl() *MockPermissionMigrationInterface_Impl_Call {
	return &MockPermissionMigrationInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockPermissionMigrationInterface_Impl_Call) Run(run func()) *MockPermissionMigrationInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPermissionMigrationInterface_Impl_Call) Return(_a0 iam.PermissionMigrationService) *MockPermissionMigrationInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPermissionMigrationInterface_Impl_Call) RunAndReturn(run func() iam.PermissionMigrationService) *MockPermissionMigrationInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// MigratePermissions provides a mock function with given fields: ctx, request
func (_m *MockPermissionMigrationInterface) MigratePermissions(ctx context.Context, request iam.PermissionMigrationRequest) (*iam.PermissionMigrationResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for MigratePermissions")
	}

	var r0 *iam.PermissionMigrationResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.PermissionMigrationRequest) (*iam.PermissionMigrationResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.PermissionMigrationRequest) *iam.PermissionMigrationResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.PermissionMigrationResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.PermissionMigrationRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPermissionMigrationInterface_MigratePermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MigratePermissions'
type MockPermissionMigrationInterface_MigratePermissions_Call struct {
	*mock.Call
}

// MigratePermissions is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.PermissionMigrationRequest
func (_e *MockPermissionMigrationInterface_Expecter) MigratePermissions(ctx interface{}, request interface{}) *MockPermissionMigrationInterface_MigratePermissions_Call {
	return &MockPermissionMigrationInterface_MigratePermissions_Call{Call: _e.mock.On("MigratePermissions", ctx, request)}
}

func (_c *MockPermissionMigrationInterface_MigratePermissions_Call) Run(run func(ctx context.Context, request iam.PermissionMigrationRequest)) *MockPermissionMigrationInterface_MigratePermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.PermissionMigrationRequest))
	})
	return _c
}

func (_c *MockPermissionMigrationInterface_MigratePermissions_Call) Return(_a0 *iam.PermissionMigrationResponse, _a1 error) *MockPermissionMigrationInterface_MigratePermissions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPermissionMigrationInterface_MigratePermissions_Call) RunAndReturn(run func(context.Context, iam.PermissionMigrationRequest) (*iam.PermissionMigrationResponse, error)) *MockPermissionMigrationInterface_MigratePermissions_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockPermissionMigrationInterface) WithImpl(impl iam.PermissionMigrationService) iam.PermissionMigrationInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 iam.PermissionMigrationInterface
	if rf, ok := ret.Get(0).(func(iam.PermissionMigrationService) iam.PermissionMigrationInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(iam.PermissionMigrationInterface)
		}
	}

	return r0
}

// MockPermissionMigrationInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockPermissionMigrationInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl iam.PermissionMigrationService
func (_e *MockPermissionMigrationInterface_Expecter) WithImpl(impl interface{}) *MockPermissionMigrationInterface_WithImpl_Call {
	return &MockPermissionMigrationInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockPermissionMigrationInterface_WithImpl_Call) Run(run func(impl iam.PermissionMigrationService)) *MockPermissionMigrationInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(iam.PermissionMigrationService))
	})
	return _c
}

func (_c *MockPermissionMigrationInterface_WithImpl_Call) Return(_a0 iam.PermissionMigrationInterface) *MockPermissionMigrationInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPermissionMigrationInterface_WithImpl_Call) RunAndReturn(run func(iam.PermissionMigrationService) iam.PermissionMigrationInterface) *MockPermissionMigrationInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPermissionMigrationInterface creates a new instance of MockPermissionMigrationInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPermissionMigrationInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPermissionMigrationInterface {
	mock := &MockPermissionMigrationInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
