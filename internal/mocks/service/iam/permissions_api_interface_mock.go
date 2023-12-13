// Code generated by mockery v2.38.0. DO NOT EDIT.

package iam

import (
	context "context"

	iam "github.com/databricks/databricks-sdk-go/service/iam"
	mock "github.com/stretchr/testify/mock"
)

// MockPermissionsAPIInterface is an autogenerated mock type for the PermissionsAPIInterface type
type MockPermissionsAPIInterface struct {
	mock.Mock
}

type MockPermissionsAPIInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPermissionsAPIInterface) EXPECT() *MockPermissionsAPIInterface_Expecter {
	return &MockPermissionsAPIInterface_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockPermissionsAPIInterface) Get(ctx context.Context, request iam.GetPermissionRequest) (*iam.ObjectPermissions, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *iam.ObjectPermissions
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.GetPermissionRequest) (*iam.ObjectPermissions, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.GetPermissionRequest) *iam.ObjectPermissions); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.ObjectPermissions)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.GetPermissionRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPermissionsAPIInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockPermissionsAPIInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.GetPermissionRequest
func (_e *MockPermissionsAPIInterface_Expecter) Get(ctx interface{}, request interface{}) *MockPermissionsAPIInterface_Get_Call {
	return &MockPermissionsAPIInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockPermissionsAPIInterface_Get_Call) Run(run func(ctx context.Context, request iam.GetPermissionRequest)) *MockPermissionsAPIInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.GetPermissionRequest))
	})
	return _c
}

func (_c *MockPermissionsAPIInterface_Get_Call) Return(_a0 *iam.ObjectPermissions, _a1 error) *MockPermissionsAPIInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPermissionsAPIInterface_Get_Call) RunAndReturn(run func(context.Context, iam.GetPermissionRequest) (*iam.ObjectPermissions, error)) *MockPermissionsAPIInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByRequestObjectTypeAndRequestObjectId provides a mock function with given fields: ctx, requestObjectType, requestObjectId
func (_m *MockPermissionsAPIInterface) GetByRequestObjectTypeAndRequestObjectId(ctx context.Context, requestObjectType string, requestObjectId string) (*iam.ObjectPermissions, error) {
	ret := _m.Called(ctx, requestObjectType, requestObjectId)

	if len(ret) == 0 {
		panic("no return value specified for GetByRequestObjectTypeAndRequestObjectId")
	}

	var r0 *iam.ObjectPermissions
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*iam.ObjectPermissions, error)); ok {
		return rf(ctx, requestObjectType, requestObjectId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *iam.ObjectPermissions); ok {
		r0 = rf(ctx, requestObjectType, requestObjectId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.ObjectPermissions)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, requestObjectType, requestObjectId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPermissionsAPIInterface_GetByRequestObjectTypeAndRequestObjectId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByRequestObjectTypeAndRequestObjectId'
type MockPermissionsAPIInterface_GetByRequestObjectTypeAndRequestObjectId_Call struct {
	*mock.Call
}

// GetByRequestObjectTypeAndRequestObjectId is a helper method to define mock.On call
//   - ctx context.Context
//   - requestObjectType string
//   - requestObjectId string
func (_e *MockPermissionsAPIInterface_Expecter) GetByRequestObjectTypeAndRequestObjectId(ctx interface{}, requestObjectType interface{}, requestObjectId interface{}) *MockPermissionsAPIInterface_GetByRequestObjectTypeAndRequestObjectId_Call {
	return &MockPermissionsAPIInterface_GetByRequestObjectTypeAndRequestObjectId_Call{Call: _e.mock.On("GetByRequestObjectTypeAndRequestObjectId", ctx, requestObjectType, requestObjectId)}
}

func (_c *MockPermissionsAPIInterface_GetByRequestObjectTypeAndRequestObjectId_Call) Run(run func(ctx context.Context, requestObjectType string, requestObjectId string)) *MockPermissionsAPIInterface_GetByRequestObjectTypeAndRequestObjectId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockPermissionsAPIInterface_GetByRequestObjectTypeAndRequestObjectId_Call) Return(_a0 *iam.ObjectPermissions, _a1 error) *MockPermissionsAPIInterface_GetByRequestObjectTypeAndRequestObjectId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPermissionsAPIInterface_GetByRequestObjectTypeAndRequestObjectId_Call) RunAndReturn(run func(context.Context, string, string) (*iam.ObjectPermissions, error)) *MockPermissionsAPIInterface_GetByRequestObjectTypeAndRequestObjectId_Call {
	_c.Call.Return(run)
	return _c
}

// GetPermissionLevels provides a mock function with given fields: ctx, request
func (_m *MockPermissionsAPIInterface) GetPermissionLevels(ctx context.Context, request iam.GetPermissionLevelsRequest) (*iam.GetPermissionLevelsResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetPermissionLevels")
	}

	var r0 *iam.GetPermissionLevelsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.GetPermissionLevelsRequest) (*iam.GetPermissionLevelsResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.GetPermissionLevelsRequest) *iam.GetPermissionLevelsResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.GetPermissionLevelsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.GetPermissionLevelsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPermissionsAPIInterface_GetPermissionLevels_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPermissionLevels'
type MockPermissionsAPIInterface_GetPermissionLevels_Call struct {
	*mock.Call
}

// GetPermissionLevels is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.GetPermissionLevelsRequest
func (_e *MockPermissionsAPIInterface_Expecter) GetPermissionLevels(ctx interface{}, request interface{}) *MockPermissionsAPIInterface_GetPermissionLevels_Call {
	return &MockPermissionsAPIInterface_GetPermissionLevels_Call{Call: _e.mock.On("GetPermissionLevels", ctx, request)}
}

func (_c *MockPermissionsAPIInterface_GetPermissionLevels_Call) Run(run func(ctx context.Context, request iam.GetPermissionLevelsRequest)) *MockPermissionsAPIInterface_GetPermissionLevels_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.GetPermissionLevelsRequest))
	})
	return _c
}

func (_c *MockPermissionsAPIInterface_GetPermissionLevels_Call) Return(_a0 *iam.GetPermissionLevelsResponse, _a1 error) *MockPermissionsAPIInterface_GetPermissionLevels_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPermissionsAPIInterface_GetPermissionLevels_Call) RunAndReturn(run func(context.Context, iam.GetPermissionLevelsRequest) (*iam.GetPermissionLevelsResponse, error)) *MockPermissionsAPIInterface_GetPermissionLevels_Call {
	_c.Call.Return(run)
	return _c
}

// GetPermissionLevelsByRequestObjectTypeAndRequestObjectId provides a mock function with given fields: ctx, requestObjectType, requestObjectId
func (_m *MockPermissionsAPIInterface) GetPermissionLevelsByRequestObjectTypeAndRequestObjectId(ctx context.Context, requestObjectType string, requestObjectId string) (*iam.GetPermissionLevelsResponse, error) {
	ret := _m.Called(ctx, requestObjectType, requestObjectId)

	if len(ret) == 0 {
		panic("no return value specified for GetPermissionLevelsByRequestObjectTypeAndRequestObjectId")
	}

	var r0 *iam.GetPermissionLevelsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*iam.GetPermissionLevelsResponse, error)); ok {
		return rf(ctx, requestObjectType, requestObjectId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *iam.GetPermissionLevelsResponse); ok {
		r0 = rf(ctx, requestObjectType, requestObjectId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.GetPermissionLevelsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, requestObjectType, requestObjectId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPermissionsAPIInterface_GetPermissionLevelsByRequestObjectTypeAndRequestObjectId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPermissionLevelsByRequestObjectTypeAndRequestObjectId'
type MockPermissionsAPIInterface_GetPermissionLevelsByRequestObjectTypeAndRequestObjectId_Call struct {
	*mock.Call
}

// GetPermissionLevelsByRequestObjectTypeAndRequestObjectId is a helper method to define mock.On call
//   - ctx context.Context
//   - requestObjectType string
//   - requestObjectId string
func (_e *MockPermissionsAPIInterface_Expecter) GetPermissionLevelsByRequestObjectTypeAndRequestObjectId(ctx interface{}, requestObjectType interface{}, requestObjectId interface{}) *MockPermissionsAPIInterface_GetPermissionLevelsByRequestObjectTypeAndRequestObjectId_Call {
	return &MockPermissionsAPIInterface_GetPermissionLevelsByRequestObjectTypeAndRequestObjectId_Call{Call: _e.mock.On("GetPermissionLevelsByRequestObjectTypeAndRequestObjectId", ctx, requestObjectType, requestObjectId)}
}

func (_c *MockPermissionsAPIInterface_GetPermissionLevelsByRequestObjectTypeAndRequestObjectId_Call) Run(run func(ctx context.Context, requestObjectType string, requestObjectId string)) *MockPermissionsAPIInterface_GetPermissionLevelsByRequestObjectTypeAndRequestObjectId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockPermissionsAPIInterface_GetPermissionLevelsByRequestObjectTypeAndRequestObjectId_Call) Return(_a0 *iam.GetPermissionLevelsResponse, _a1 error) *MockPermissionsAPIInterface_GetPermissionLevelsByRequestObjectTypeAndRequestObjectId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPermissionsAPIInterface_GetPermissionLevelsByRequestObjectTypeAndRequestObjectId_Call) RunAndReturn(run func(context.Context, string, string) (*iam.GetPermissionLevelsResponse, error)) *MockPermissionsAPIInterface_GetPermissionLevelsByRequestObjectTypeAndRequestObjectId_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockPermissionsAPIInterface) Impl() iam.PermissionsService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 iam.PermissionsService
	if rf, ok := ret.Get(0).(func() iam.PermissionsService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(iam.PermissionsService)
		}
	}

	return r0
}

// MockPermissionsAPIInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockPermissionsAPIInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockPermissionsAPIInterface_Expecter) Impl() *MockPermissionsAPIInterface_Impl_Call {
	return &MockPermissionsAPIInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockPermissionsAPIInterface_Impl_Call) Run(run func()) *MockPermissionsAPIInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPermissionsAPIInterface_Impl_Call) Return(_a0 iam.PermissionsService) *MockPermissionsAPIInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPermissionsAPIInterface_Impl_Call) RunAndReturn(run func() iam.PermissionsService) *MockPermissionsAPIInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: ctx, request
func (_m *MockPermissionsAPIInterface) Set(ctx context.Context, request iam.PermissionsRequest) (*iam.ObjectPermissions, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 *iam.ObjectPermissions
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.PermissionsRequest) (*iam.ObjectPermissions, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.PermissionsRequest) *iam.ObjectPermissions); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.ObjectPermissions)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.PermissionsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPermissionsAPIInterface_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type MockPermissionsAPIInterface_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.PermissionsRequest
func (_e *MockPermissionsAPIInterface_Expecter) Set(ctx interface{}, request interface{}) *MockPermissionsAPIInterface_Set_Call {
	return &MockPermissionsAPIInterface_Set_Call{Call: _e.mock.On("Set", ctx, request)}
}

func (_c *MockPermissionsAPIInterface_Set_Call) Run(run func(ctx context.Context, request iam.PermissionsRequest)) *MockPermissionsAPIInterface_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.PermissionsRequest))
	})
	return _c
}

func (_c *MockPermissionsAPIInterface_Set_Call) Return(_a0 *iam.ObjectPermissions, _a1 error) *MockPermissionsAPIInterface_Set_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPermissionsAPIInterface_Set_Call) RunAndReturn(run func(context.Context, iam.PermissionsRequest) (*iam.ObjectPermissions, error)) *MockPermissionsAPIInterface_Set_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockPermissionsAPIInterface) Update(ctx context.Context, request iam.PermissionsRequest) (*iam.ObjectPermissions, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *iam.ObjectPermissions
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, iam.PermissionsRequest) (*iam.ObjectPermissions, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, iam.PermissionsRequest) *iam.ObjectPermissions); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.ObjectPermissions)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, iam.PermissionsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPermissionsAPIInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockPermissionsAPIInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request iam.PermissionsRequest
func (_e *MockPermissionsAPIInterface_Expecter) Update(ctx interface{}, request interface{}) *MockPermissionsAPIInterface_Update_Call {
	return &MockPermissionsAPIInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockPermissionsAPIInterface_Update_Call) Run(run func(ctx context.Context, request iam.PermissionsRequest)) *MockPermissionsAPIInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(iam.PermissionsRequest))
	})
	return _c
}

func (_c *MockPermissionsAPIInterface_Update_Call) Return(_a0 *iam.ObjectPermissions, _a1 error) *MockPermissionsAPIInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPermissionsAPIInterface_Update_Call) RunAndReturn(run func(context.Context, iam.PermissionsRequest) (*iam.ObjectPermissions, error)) *MockPermissionsAPIInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockPermissionsAPIInterface) WithImpl(impl iam.PermissionsService) iam.PermissionsAPIInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 iam.PermissionsAPIInterface
	if rf, ok := ret.Get(0).(func(iam.PermissionsService) iam.PermissionsAPIInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(iam.PermissionsAPIInterface)
		}
	}

	return r0
}

// MockPermissionsAPIInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockPermissionsAPIInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl iam.PermissionsService
func (_e *MockPermissionsAPIInterface_Expecter) WithImpl(impl interface{}) *MockPermissionsAPIInterface_WithImpl_Call {
	return &MockPermissionsAPIInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockPermissionsAPIInterface_WithImpl_Call) Run(run func(impl iam.PermissionsService)) *MockPermissionsAPIInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(iam.PermissionsService))
	})
	return _c
}

func (_c *MockPermissionsAPIInterface_WithImpl_Call) Return(_a0 iam.PermissionsAPIInterface) *MockPermissionsAPIInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPermissionsAPIInterface_WithImpl_Call) RunAndReturn(run func(iam.PermissionsService) iam.PermissionsAPIInterface) *MockPermissionsAPIInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPermissionsAPIInterface creates a new instance of MockPermissionsAPIInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPermissionsAPIInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPermissionsAPIInterface {
	mock := &MockPermissionsAPIInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
