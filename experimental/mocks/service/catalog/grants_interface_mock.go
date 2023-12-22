// Code generated by mockery v2.38.0. DO NOT EDIT.

package catalog

import (
	context "context"

	catalog "github.com/databricks/databricks-sdk-go/service/catalog"

	mock "github.com/stretchr/testify/mock"
)

// MockGrantsInterface is an autogenerated mock type for the GrantsInterface type
type MockGrantsInterface struct {
	mock.Mock
}

type MockGrantsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGrantsInterface) EXPECT() *MockGrantsInterface_Expecter {
	return &MockGrantsInterface_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockGrantsInterface) Get(ctx context.Context, request catalog.GetGrantRequest) (*catalog.PermissionsList, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *catalog.PermissionsList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetGrantRequest) (*catalog.PermissionsList, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetGrantRequest) *catalog.PermissionsList); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.PermissionsList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.GetGrantRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGrantsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockGrantsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.GetGrantRequest
func (_e *MockGrantsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockGrantsInterface_Get_Call {
	return &MockGrantsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockGrantsInterface_Get_Call) Run(run func(ctx context.Context, request catalog.GetGrantRequest)) *MockGrantsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.GetGrantRequest))
	})
	return _c
}

func (_c *MockGrantsInterface_Get_Call) Return(_a0 *catalog.PermissionsList, _a1 error) *MockGrantsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGrantsInterface_Get_Call) RunAndReturn(run func(context.Context, catalog.GetGrantRequest) (*catalog.PermissionsList, error)) *MockGrantsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetBySecurableTypeAndFullName provides a mock function with given fields: ctx, securableType, fullName
func (_m *MockGrantsInterface) GetBySecurableTypeAndFullName(ctx context.Context, securableType catalog.SecurableType, fullName string) (*catalog.PermissionsList, error) {
	ret := _m.Called(ctx, securableType, fullName)

	if len(ret) == 0 {
		panic("no return value specified for GetBySecurableTypeAndFullName")
	}

	var r0 *catalog.PermissionsList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.SecurableType, string) (*catalog.PermissionsList, error)); ok {
		return rf(ctx, securableType, fullName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.SecurableType, string) *catalog.PermissionsList); ok {
		r0 = rf(ctx, securableType, fullName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.PermissionsList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.SecurableType, string) error); ok {
		r1 = rf(ctx, securableType, fullName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGrantsInterface_GetBySecurableTypeAndFullName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBySecurableTypeAndFullName'
type MockGrantsInterface_GetBySecurableTypeAndFullName_Call struct {
	*mock.Call
}

// GetBySecurableTypeAndFullName is a helper method to define mock.On call
//   - ctx context.Context
//   - securableType catalog.SecurableType
//   - fullName string
func (_e *MockGrantsInterface_Expecter) GetBySecurableTypeAndFullName(ctx interface{}, securableType interface{}, fullName interface{}) *MockGrantsInterface_GetBySecurableTypeAndFullName_Call {
	return &MockGrantsInterface_GetBySecurableTypeAndFullName_Call{Call: _e.mock.On("GetBySecurableTypeAndFullName", ctx, securableType, fullName)}
}

func (_c *MockGrantsInterface_GetBySecurableTypeAndFullName_Call) Run(run func(ctx context.Context, securableType catalog.SecurableType, fullName string)) *MockGrantsInterface_GetBySecurableTypeAndFullName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.SecurableType), args[2].(string))
	})
	return _c
}

func (_c *MockGrantsInterface_GetBySecurableTypeAndFullName_Call) Return(_a0 *catalog.PermissionsList, _a1 error) *MockGrantsInterface_GetBySecurableTypeAndFullName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGrantsInterface_GetBySecurableTypeAndFullName_Call) RunAndReturn(run func(context.Context, catalog.SecurableType, string) (*catalog.PermissionsList, error)) *MockGrantsInterface_GetBySecurableTypeAndFullName_Call {
	_c.Call.Return(run)
	return _c
}

// GetEffective provides a mock function with given fields: ctx, request
func (_m *MockGrantsInterface) GetEffective(ctx context.Context, request catalog.GetEffectiveRequest) (*catalog.EffectivePermissionsList, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetEffective")
	}

	var r0 *catalog.EffectivePermissionsList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetEffectiveRequest) (*catalog.EffectivePermissionsList, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetEffectiveRequest) *catalog.EffectivePermissionsList); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.EffectivePermissionsList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.GetEffectiveRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGrantsInterface_GetEffective_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEffective'
type MockGrantsInterface_GetEffective_Call struct {
	*mock.Call
}

// GetEffective is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.GetEffectiveRequest
func (_e *MockGrantsInterface_Expecter) GetEffective(ctx interface{}, request interface{}) *MockGrantsInterface_GetEffective_Call {
	return &MockGrantsInterface_GetEffective_Call{Call: _e.mock.On("GetEffective", ctx, request)}
}

func (_c *MockGrantsInterface_GetEffective_Call) Run(run func(ctx context.Context, request catalog.GetEffectiveRequest)) *MockGrantsInterface_GetEffective_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.GetEffectiveRequest))
	})
	return _c
}

func (_c *MockGrantsInterface_GetEffective_Call) Return(_a0 *catalog.EffectivePermissionsList, _a1 error) *MockGrantsInterface_GetEffective_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGrantsInterface_GetEffective_Call) RunAndReturn(run func(context.Context, catalog.GetEffectiveRequest) (*catalog.EffectivePermissionsList, error)) *MockGrantsInterface_GetEffective_Call {
	_c.Call.Return(run)
	return _c
}

// GetEffectiveBySecurableTypeAndFullName provides a mock function with given fields: ctx, securableType, fullName
func (_m *MockGrantsInterface) GetEffectiveBySecurableTypeAndFullName(ctx context.Context, securableType catalog.SecurableType, fullName string) (*catalog.EffectivePermissionsList, error) {
	ret := _m.Called(ctx, securableType, fullName)

	if len(ret) == 0 {
		panic("no return value specified for GetEffectiveBySecurableTypeAndFullName")
	}

	var r0 *catalog.EffectivePermissionsList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.SecurableType, string) (*catalog.EffectivePermissionsList, error)); ok {
		return rf(ctx, securableType, fullName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.SecurableType, string) *catalog.EffectivePermissionsList); ok {
		r0 = rf(ctx, securableType, fullName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.EffectivePermissionsList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.SecurableType, string) error); ok {
		r1 = rf(ctx, securableType, fullName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGrantsInterface_GetEffectiveBySecurableTypeAndFullName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEffectiveBySecurableTypeAndFullName'
type MockGrantsInterface_GetEffectiveBySecurableTypeAndFullName_Call struct {
	*mock.Call
}

// GetEffectiveBySecurableTypeAndFullName is a helper method to define mock.On call
//   - ctx context.Context
//   - securableType catalog.SecurableType
//   - fullName string
func (_e *MockGrantsInterface_Expecter) GetEffectiveBySecurableTypeAndFullName(ctx interface{}, securableType interface{}, fullName interface{}) *MockGrantsInterface_GetEffectiveBySecurableTypeAndFullName_Call {
	return &MockGrantsInterface_GetEffectiveBySecurableTypeAndFullName_Call{Call: _e.mock.On("GetEffectiveBySecurableTypeAndFullName", ctx, securableType, fullName)}
}

func (_c *MockGrantsInterface_GetEffectiveBySecurableTypeAndFullName_Call) Run(run func(ctx context.Context, securableType catalog.SecurableType, fullName string)) *MockGrantsInterface_GetEffectiveBySecurableTypeAndFullName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.SecurableType), args[2].(string))
	})
	return _c
}

func (_c *MockGrantsInterface_GetEffectiveBySecurableTypeAndFullName_Call) Return(_a0 *catalog.EffectivePermissionsList, _a1 error) *MockGrantsInterface_GetEffectiveBySecurableTypeAndFullName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGrantsInterface_GetEffectiveBySecurableTypeAndFullName_Call) RunAndReturn(run func(context.Context, catalog.SecurableType, string) (*catalog.EffectivePermissionsList, error)) *MockGrantsInterface_GetEffectiveBySecurableTypeAndFullName_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockGrantsInterface) Impl() catalog.GrantsService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 catalog.GrantsService
	if rf, ok := ret.Get(0).(func() catalog.GrantsService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.GrantsService)
		}
	}

	return r0
}

// MockGrantsInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockGrantsInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockGrantsInterface_Expecter) Impl() *MockGrantsInterface_Impl_Call {
	return &MockGrantsInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockGrantsInterface_Impl_Call) Run(run func()) *MockGrantsInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockGrantsInterface_Impl_Call) Return(_a0 catalog.GrantsService) *MockGrantsInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGrantsInterface_Impl_Call) RunAndReturn(run func() catalog.GrantsService) *MockGrantsInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockGrantsInterface) Update(ctx context.Context, request catalog.UpdatePermissions) (*catalog.PermissionsList, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *catalog.PermissionsList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdatePermissions) (*catalog.PermissionsList, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdatePermissions) *catalog.PermissionsList); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.PermissionsList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.UpdatePermissions) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGrantsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockGrantsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.UpdatePermissions
func (_e *MockGrantsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockGrantsInterface_Update_Call {
	return &MockGrantsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockGrantsInterface_Update_Call) Run(run func(ctx context.Context, request catalog.UpdatePermissions)) *MockGrantsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.UpdatePermissions))
	})
	return _c
}

func (_c *MockGrantsInterface_Update_Call) Return(_a0 *catalog.PermissionsList, _a1 error) *MockGrantsInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGrantsInterface_Update_Call) RunAndReturn(run func(context.Context, catalog.UpdatePermissions) (*catalog.PermissionsList, error)) *MockGrantsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockGrantsInterface) WithImpl(impl catalog.GrantsService) catalog.GrantsInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 catalog.GrantsInterface
	if rf, ok := ret.Get(0).(func(catalog.GrantsService) catalog.GrantsInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.GrantsInterface)
		}
	}

	return r0
}

// MockGrantsInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockGrantsInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl catalog.GrantsService
func (_e *MockGrantsInterface_Expecter) WithImpl(impl interface{}) *MockGrantsInterface_WithImpl_Call {
	return &MockGrantsInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockGrantsInterface_WithImpl_Call) Run(run func(impl catalog.GrantsService)) *MockGrantsInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(catalog.GrantsService))
	})
	return _c
}

func (_c *MockGrantsInterface_WithImpl_Call) Return(_a0 catalog.GrantsInterface) *MockGrantsInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGrantsInterface_WithImpl_Call) RunAndReturn(run func(catalog.GrantsService) catalog.GrantsInterface) *MockGrantsInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockGrantsInterface creates a new instance of MockGrantsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGrantsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGrantsInterface {
	mock := &MockGrantsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
