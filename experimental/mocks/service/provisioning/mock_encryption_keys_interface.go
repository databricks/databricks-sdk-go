// Code generated by mockery v2.53.2. DO NOT EDIT.

package provisioning

import (
	context "context"

	provisioning "github.com/databricks/databricks-sdk-go/service/provisioning"
	mock "github.com/stretchr/testify/mock"
)

// MockEncryptionKeysInterface is an autogenerated mock type for the EncryptionKeysInterface type
type MockEncryptionKeysInterface struct {
	mock.Mock
}

type MockEncryptionKeysInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEncryptionKeysInterface) EXPECT() *MockEncryptionKeysInterface_Expecter {
	return &MockEncryptionKeysInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockEncryptionKeysInterface) Create(ctx context.Context, request provisioning.CreateCustomerManagedKeyRequest) (*provisioning.CustomerManagedKey, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *provisioning.CustomerManagedKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, provisioning.CreateCustomerManagedKeyRequest) (*provisioning.CustomerManagedKey, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, provisioning.CreateCustomerManagedKeyRequest) *provisioning.CustomerManagedKey); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provisioning.CustomerManagedKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, provisioning.CreateCustomerManagedKeyRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEncryptionKeysInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockEncryptionKeysInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request provisioning.CreateCustomerManagedKeyRequest
func (_e *MockEncryptionKeysInterface_Expecter) Create(ctx interface{}, request interface{}) *MockEncryptionKeysInterface_Create_Call {
	return &MockEncryptionKeysInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockEncryptionKeysInterface_Create_Call) Run(run func(ctx context.Context, request provisioning.CreateCustomerManagedKeyRequest)) *MockEncryptionKeysInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(provisioning.CreateCustomerManagedKeyRequest))
	})
	return _c
}

func (_c *MockEncryptionKeysInterface_Create_Call) Return(_a0 *provisioning.CustomerManagedKey, _a1 error) *MockEncryptionKeysInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEncryptionKeysInterface_Create_Call) RunAndReturn(run func(context.Context, provisioning.CreateCustomerManagedKeyRequest) (*provisioning.CustomerManagedKey, error)) *MockEncryptionKeysInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockEncryptionKeysInterface) Delete(ctx context.Context, request provisioning.DeleteEncryptionKeyRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, provisioning.DeleteEncryptionKeyRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEncryptionKeysInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockEncryptionKeysInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request provisioning.DeleteEncryptionKeyRequest
func (_e *MockEncryptionKeysInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockEncryptionKeysInterface_Delete_Call {
	return &MockEncryptionKeysInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockEncryptionKeysInterface_Delete_Call) Run(run func(ctx context.Context, request provisioning.DeleteEncryptionKeyRequest)) *MockEncryptionKeysInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(provisioning.DeleteEncryptionKeyRequest))
	})
	return _c
}

func (_c *MockEncryptionKeysInterface_Delete_Call) Return(_a0 error) *MockEncryptionKeysInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEncryptionKeysInterface_Delete_Call) RunAndReturn(run func(context.Context, provisioning.DeleteEncryptionKeyRequest) error) *MockEncryptionKeysInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByCustomerManagedKeyId provides a mock function with given fields: ctx, customerManagedKeyId
func (_m *MockEncryptionKeysInterface) DeleteByCustomerManagedKeyId(ctx context.Context, customerManagedKeyId string) error {
	ret := _m.Called(ctx, customerManagedKeyId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByCustomerManagedKeyId")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, customerManagedKeyId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEncryptionKeysInterface_DeleteByCustomerManagedKeyId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByCustomerManagedKeyId'
type MockEncryptionKeysInterface_DeleteByCustomerManagedKeyId_Call struct {
	*mock.Call
}

// DeleteByCustomerManagedKeyId is a helper method to define mock.On call
//   - ctx context.Context
//   - customerManagedKeyId string
func (_e *MockEncryptionKeysInterface_Expecter) DeleteByCustomerManagedKeyId(ctx interface{}, customerManagedKeyId interface{}) *MockEncryptionKeysInterface_DeleteByCustomerManagedKeyId_Call {
	return &MockEncryptionKeysInterface_DeleteByCustomerManagedKeyId_Call{Call: _e.mock.On("DeleteByCustomerManagedKeyId", ctx, customerManagedKeyId)}
}

func (_c *MockEncryptionKeysInterface_DeleteByCustomerManagedKeyId_Call) Run(run func(ctx context.Context, customerManagedKeyId string)) *MockEncryptionKeysInterface_DeleteByCustomerManagedKeyId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockEncryptionKeysInterface_DeleteByCustomerManagedKeyId_Call) Return(_a0 error) *MockEncryptionKeysInterface_DeleteByCustomerManagedKeyId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEncryptionKeysInterface_DeleteByCustomerManagedKeyId_Call) RunAndReturn(run func(context.Context, string) error) *MockEncryptionKeysInterface_DeleteByCustomerManagedKeyId_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockEncryptionKeysInterface) Get(ctx context.Context, request provisioning.GetEncryptionKeyRequest) (*provisioning.CustomerManagedKey, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *provisioning.CustomerManagedKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, provisioning.GetEncryptionKeyRequest) (*provisioning.CustomerManagedKey, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, provisioning.GetEncryptionKeyRequest) *provisioning.CustomerManagedKey); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provisioning.CustomerManagedKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, provisioning.GetEncryptionKeyRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEncryptionKeysInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockEncryptionKeysInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request provisioning.GetEncryptionKeyRequest
func (_e *MockEncryptionKeysInterface_Expecter) Get(ctx interface{}, request interface{}) *MockEncryptionKeysInterface_Get_Call {
	return &MockEncryptionKeysInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockEncryptionKeysInterface_Get_Call) Run(run func(ctx context.Context, request provisioning.GetEncryptionKeyRequest)) *MockEncryptionKeysInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(provisioning.GetEncryptionKeyRequest))
	})
	return _c
}

func (_c *MockEncryptionKeysInterface_Get_Call) Return(_a0 *provisioning.CustomerManagedKey, _a1 error) *MockEncryptionKeysInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEncryptionKeysInterface_Get_Call) RunAndReturn(run func(context.Context, provisioning.GetEncryptionKeyRequest) (*provisioning.CustomerManagedKey, error)) *MockEncryptionKeysInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByCustomerManagedKeyId provides a mock function with given fields: ctx, customerManagedKeyId
func (_m *MockEncryptionKeysInterface) GetByCustomerManagedKeyId(ctx context.Context, customerManagedKeyId string) (*provisioning.CustomerManagedKey, error) {
	ret := _m.Called(ctx, customerManagedKeyId)

	if len(ret) == 0 {
		panic("no return value specified for GetByCustomerManagedKeyId")
	}

	var r0 *provisioning.CustomerManagedKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*provisioning.CustomerManagedKey, error)); ok {
		return rf(ctx, customerManagedKeyId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *provisioning.CustomerManagedKey); ok {
		r0 = rf(ctx, customerManagedKeyId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provisioning.CustomerManagedKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, customerManagedKeyId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEncryptionKeysInterface_GetByCustomerManagedKeyId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByCustomerManagedKeyId'
type MockEncryptionKeysInterface_GetByCustomerManagedKeyId_Call struct {
	*mock.Call
}

// GetByCustomerManagedKeyId is a helper method to define mock.On call
//   - ctx context.Context
//   - customerManagedKeyId string
func (_e *MockEncryptionKeysInterface_Expecter) GetByCustomerManagedKeyId(ctx interface{}, customerManagedKeyId interface{}) *MockEncryptionKeysInterface_GetByCustomerManagedKeyId_Call {
	return &MockEncryptionKeysInterface_GetByCustomerManagedKeyId_Call{Call: _e.mock.On("GetByCustomerManagedKeyId", ctx, customerManagedKeyId)}
}

func (_c *MockEncryptionKeysInterface_GetByCustomerManagedKeyId_Call) Run(run func(ctx context.Context, customerManagedKeyId string)) *MockEncryptionKeysInterface_GetByCustomerManagedKeyId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockEncryptionKeysInterface_GetByCustomerManagedKeyId_Call) Return(_a0 *provisioning.CustomerManagedKey, _a1 error) *MockEncryptionKeysInterface_GetByCustomerManagedKeyId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEncryptionKeysInterface_GetByCustomerManagedKeyId_Call) RunAndReturn(run func(context.Context, string) (*provisioning.CustomerManagedKey, error)) *MockEncryptionKeysInterface_GetByCustomerManagedKeyId_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx
func (_m *MockEncryptionKeysInterface) List(ctx context.Context) ([]provisioning.CustomerManagedKey, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []provisioning.CustomerManagedKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]provisioning.CustomerManagedKey, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []provisioning.CustomerManagedKey); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]provisioning.CustomerManagedKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEncryptionKeysInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockEncryptionKeysInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockEncryptionKeysInterface_Expecter) List(ctx interface{}) *MockEncryptionKeysInterface_List_Call {
	return &MockEncryptionKeysInterface_List_Call{Call: _e.mock.On("List", ctx)}
}

func (_c *MockEncryptionKeysInterface_List_Call) Run(run func(ctx context.Context)) *MockEncryptionKeysInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockEncryptionKeysInterface_List_Call) Return(_a0 []provisioning.CustomerManagedKey, _a1 error) *MockEncryptionKeysInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEncryptionKeysInterface_List_Call) RunAndReturn(run func(context.Context) ([]provisioning.CustomerManagedKey, error)) *MockEncryptionKeysInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockEncryptionKeysInterface creates a new instance of MockEncryptionKeysInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEncryptionKeysInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEncryptionKeysInterface {
	mock := &MockEncryptionKeysInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
