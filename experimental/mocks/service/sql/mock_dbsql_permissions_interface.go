// Code generated by mockery v2.53.2. DO NOT EDIT.

package sql

import (
	context "context"

	sql "github.com/databricks/databricks-sdk-go/service/sql"
	mock "github.com/stretchr/testify/mock"
)

// MockDbsqlPermissionsInterface is an autogenerated mock type for the DbsqlPermissionsInterface type
type MockDbsqlPermissionsInterface struct {
	mock.Mock
}

type MockDbsqlPermissionsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDbsqlPermissionsInterface) EXPECT() *MockDbsqlPermissionsInterface_Expecter {
	return &MockDbsqlPermissionsInterface_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockDbsqlPermissionsInterface) Get(ctx context.Context, request sql.GetDbsqlPermissionRequest) (*sql.GetResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *sql.GetResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.GetDbsqlPermissionRequest) (*sql.GetResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sql.GetDbsqlPermissionRequest) *sql.GetResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.GetResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sql.GetDbsqlPermissionRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDbsqlPermissionsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockDbsqlPermissionsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.GetDbsqlPermissionRequest
func (_e *MockDbsqlPermissionsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockDbsqlPermissionsInterface_Get_Call {
	return &MockDbsqlPermissionsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockDbsqlPermissionsInterface_Get_Call) Run(run func(ctx context.Context, request sql.GetDbsqlPermissionRequest)) *MockDbsqlPermissionsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.GetDbsqlPermissionRequest))
	})
	return _c
}

func (_c *MockDbsqlPermissionsInterface_Get_Call) Return(_a0 *sql.GetResponse, _a1 error) *MockDbsqlPermissionsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDbsqlPermissionsInterface_Get_Call) RunAndReturn(run func(context.Context, sql.GetDbsqlPermissionRequest) (*sql.GetResponse, error)) *MockDbsqlPermissionsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByObjectTypeAndObjectId provides a mock function with given fields: ctx, objectType, objectId
func (_m *MockDbsqlPermissionsInterface) GetByObjectTypeAndObjectId(ctx context.Context, objectType sql.ObjectTypePlural, objectId string) (*sql.GetResponse, error) {
	ret := _m.Called(ctx, objectType, objectId)

	if len(ret) == 0 {
		panic("no return value specified for GetByObjectTypeAndObjectId")
	}

	var r0 *sql.GetResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.ObjectTypePlural, string) (*sql.GetResponse, error)); ok {
		return rf(ctx, objectType, objectId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sql.ObjectTypePlural, string) *sql.GetResponse); ok {
		r0 = rf(ctx, objectType, objectId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.GetResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sql.ObjectTypePlural, string) error); ok {
		r1 = rf(ctx, objectType, objectId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDbsqlPermissionsInterface_GetByObjectTypeAndObjectId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByObjectTypeAndObjectId'
type MockDbsqlPermissionsInterface_GetByObjectTypeAndObjectId_Call struct {
	*mock.Call
}

// GetByObjectTypeAndObjectId is a helper method to define mock.On call
//   - ctx context.Context
//   - objectType sql.ObjectTypePlural
//   - objectId string
func (_e *MockDbsqlPermissionsInterface_Expecter) GetByObjectTypeAndObjectId(ctx interface{}, objectType interface{}, objectId interface{}) *MockDbsqlPermissionsInterface_GetByObjectTypeAndObjectId_Call {
	return &MockDbsqlPermissionsInterface_GetByObjectTypeAndObjectId_Call{Call: _e.mock.On("GetByObjectTypeAndObjectId", ctx, objectType, objectId)}
}

func (_c *MockDbsqlPermissionsInterface_GetByObjectTypeAndObjectId_Call) Run(run func(ctx context.Context, objectType sql.ObjectTypePlural, objectId string)) *MockDbsqlPermissionsInterface_GetByObjectTypeAndObjectId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.ObjectTypePlural), args[2].(string))
	})
	return _c
}

func (_c *MockDbsqlPermissionsInterface_GetByObjectTypeAndObjectId_Call) Return(_a0 *sql.GetResponse, _a1 error) *MockDbsqlPermissionsInterface_GetByObjectTypeAndObjectId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDbsqlPermissionsInterface_GetByObjectTypeAndObjectId_Call) RunAndReturn(run func(context.Context, sql.ObjectTypePlural, string) (*sql.GetResponse, error)) *MockDbsqlPermissionsInterface_GetByObjectTypeAndObjectId_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: ctx, request
func (_m *MockDbsqlPermissionsInterface) Set(ctx context.Context, request sql.SetRequest) (*sql.SetResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 *sql.SetResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.SetRequest) (*sql.SetResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sql.SetRequest) *sql.SetResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.SetResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sql.SetRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDbsqlPermissionsInterface_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type MockDbsqlPermissionsInterface_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.SetRequest
func (_e *MockDbsqlPermissionsInterface_Expecter) Set(ctx interface{}, request interface{}) *MockDbsqlPermissionsInterface_Set_Call {
	return &MockDbsqlPermissionsInterface_Set_Call{Call: _e.mock.On("Set", ctx, request)}
}

func (_c *MockDbsqlPermissionsInterface_Set_Call) Run(run func(ctx context.Context, request sql.SetRequest)) *MockDbsqlPermissionsInterface_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.SetRequest))
	})
	return _c
}

func (_c *MockDbsqlPermissionsInterface_Set_Call) Return(_a0 *sql.SetResponse, _a1 error) *MockDbsqlPermissionsInterface_Set_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDbsqlPermissionsInterface_Set_Call) RunAndReturn(run func(context.Context, sql.SetRequest) (*sql.SetResponse, error)) *MockDbsqlPermissionsInterface_Set_Call {
	_c.Call.Return(run)
	return _c
}

// TransferOwnership provides a mock function with given fields: ctx, request
func (_m *MockDbsqlPermissionsInterface) TransferOwnership(ctx context.Context, request sql.TransferOwnershipRequest) (*sql.Success, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for TransferOwnership")
	}

	var r0 *sql.Success
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.TransferOwnershipRequest) (*sql.Success, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sql.TransferOwnershipRequest) *sql.Success); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Success)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sql.TransferOwnershipRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDbsqlPermissionsInterface_TransferOwnership_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TransferOwnership'
type MockDbsqlPermissionsInterface_TransferOwnership_Call struct {
	*mock.Call
}

// TransferOwnership is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.TransferOwnershipRequest
func (_e *MockDbsqlPermissionsInterface_Expecter) TransferOwnership(ctx interface{}, request interface{}) *MockDbsqlPermissionsInterface_TransferOwnership_Call {
	return &MockDbsqlPermissionsInterface_TransferOwnership_Call{Call: _e.mock.On("TransferOwnership", ctx, request)}
}

func (_c *MockDbsqlPermissionsInterface_TransferOwnership_Call) Run(run func(ctx context.Context, request sql.TransferOwnershipRequest)) *MockDbsqlPermissionsInterface_TransferOwnership_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.TransferOwnershipRequest))
	})
	return _c
}

func (_c *MockDbsqlPermissionsInterface_TransferOwnership_Call) Return(_a0 *sql.Success, _a1 error) *MockDbsqlPermissionsInterface_TransferOwnership_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDbsqlPermissionsInterface_TransferOwnership_Call) RunAndReturn(run func(context.Context, sql.TransferOwnershipRequest) (*sql.Success, error)) *MockDbsqlPermissionsInterface_TransferOwnership_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDbsqlPermissionsInterface creates a new instance of MockDbsqlPermissionsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDbsqlPermissionsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDbsqlPermissionsInterface {
	mock := &MockDbsqlPermissionsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
