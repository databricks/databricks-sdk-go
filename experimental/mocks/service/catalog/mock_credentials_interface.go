// Code generated by mockery v2.53.2. DO NOT EDIT.

package catalog

import (
	context "context"

	catalog "github.com/databricks/databricks-sdk-go/service/catalog"

	listing "github.com/databricks/databricks-sdk-go/listing"

	mock "github.com/stretchr/testify/mock"
)

// MockCredentialsInterface is an autogenerated mock type for the CredentialsInterface type
type MockCredentialsInterface struct {
	mock.Mock
}

type MockCredentialsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCredentialsInterface) EXPECT() *MockCredentialsInterface_Expecter {
	return &MockCredentialsInterface_Expecter{mock: &_m.Mock}
}

// CreateCredential provides a mock function with given fields: ctx, request
func (_m *MockCredentialsInterface) CreateCredential(ctx context.Context, request catalog.CreateCredentialRequest) (*catalog.CredentialInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for CreateCredential")
	}

	var r0 *catalog.CredentialInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateCredentialRequest) (*catalog.CredentialInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateCredentialRequest) *catalog.CredentialInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.CredentialInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.CreateCredentialRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCredentialsInterface_CreateCredential_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCredential'
type MockCredentialsInterface_CreateCredential_Call struct {
	*mock.Call
}

// CreateCredential is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.CreateCredentialRequest
func (_e *MockCredentialsInterface_Expecter) CreateCredential(ctx interface{}, request interface{}) *MockCredentialsInterface_CreateCredential_Call {
	return &MockCredentialsInterface_CreateCredential_Call{Call: _e.mock.On("CreateCredential", ctx, request)}
}

func (_c *MockCredentialsInterface_CreateCredential_Call) Run(run func(ctx context.Context, request catalog.CreateCredentialRequest)) *MockCredentialsInterface_CreateCredential_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.CreateCredentialRequest))
	})
	return _c
}

func (_c *MockCredentialsInterface_CreateCredential_Call) Return(_a0 *catalog.CredentialInfo, _a1 error) *MockCredentialsInterface_CreateCredential_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCredentialsInterface_CreateCredential_Call) RunAndReturn(run func(context.Context, catalog.CreateCredentialRequest) (*catalog.CredentialInfo, error)) *MockCredentialsInterface_CreateCredential_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCredential provides a mock function with given fields: ctx, request
func (_m *MockCredentialsInterface) DeleteCredential(ctx context.Context, request catalog.DeleteCredentialRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCredential")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.DeleteCredentialRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCredentialsInterface_DeleteCredential_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCredential'
type MockCredentialsInterface_DeleteCredential_Call struct {
	*mock.Call
}

// DeleteCredential is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.DeleteCredentialRequest
func (_e *MockCredentialsInterface_Expecter) DeleteCredential(ctx interface{}, request interface{}) *MockCredentialsInterface_DeleteCredential_Call {
	return &MockCredentialsInterface_DeleteCredential_Call{Call: _e.mock.On("DeleteCredential", ctx, request)}
}

func (_c *MockCredentialsInterface_DeleteCredential_Call) Run(run func(ctx context.Context, request catalog.DeleteCredentialRequest)) *MockCredentialsInterface_DeleteCredential_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.DeleteCredentialRequest))
	})
	return _c
}

func (_c *MockCredentialsInterface_DeleteCredential_Call) Return(_a0 error) *MockCredentialsInterface_DeleteCredential_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCredentialsInterface_DeleteCredential_Call) RunAndReturn(run func(context.Context, catalog.DeleteCredentialRequest) error) *MockCredentialsInterface_DeleteCredential_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCredentialByNameArg provides a mock function with given fields: ctx, nameArg
func (_m *MockCredentialsInterface) DeleteCredentialByNameArg(ctx context.Context, nameArg string) error {
	ret := _m.Called(ctx, nameArg)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCredentialByNameArg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, nameArg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCredentialsInterface_DeleteCredentialByNameArg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCredentialByNameArg'
type MockCredentialsInterface_DeleteCredentialByNameArg_Call struct {
	*mock.Call
}

// DeleteCredentialByNameArg is a helper method to define mock.On call
//   - ctx context.Context
//   - nameArg string
func (_e *MockCredentialsInterface_Expecter) DeleteCredentialByNameArg(ctx interface{}, nameArg interface{}) *MockCredentialsInterface_DeleteCredentialByNameArg_Call {
	return &MockCredentialsInterface_DeleteCredentialByNameArg_Call{Call: _e.mock.On("DeleteCredentialByNameArg", ctx, nameArg)}
}

func (_c *MockCredentialsInterface_DeleteCredentialByNameArg_Call) Run(run func(ctx context.Context, nameArg string)) *MockCredentialsInterface_DeleteCredentialByNameArg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockCredentialsInterface_DeleteCredentialByNameArg_Call) Return(_a0 error) *MockCredentialsInterface_DeleteCredentialByNameArg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCredentialsInterface_DeleteCredentialByNameArg_Call) RunAndReturn(run func(context.Context, string) error) *MockCredentialsInterface_DeleteCredentialByNameArg_Call {
	_c.Call.Return(run)
	return _c
}

// GenerateTemporaryServiceCredential provides a mock function with given fields: ctx, request
func (_m *MockCredentialsInterface) GenerateTemporaryServiceCredential(ctx context.Context, request catalog.GenerateTemporaryServiceCredentialRequest) (*catalog.TemporaryCredentials, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GenerateTemporaryServiceCredential")
	}

	var r0 *catalog.TemporaryCredentials
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GenerateTemporaryServiceCredentialRequest) (*catalog.TemporaryCredentials, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GenerateTemporaryServiceCredentialRequest) *catalog.TemporaryCredentials); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.TemporaryCredentials)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.GenerateTemporaryServiceCredentialRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCredentialsInterface_GenerateTemporaryServiceCredential_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateTemporaryServiceCredential'
type MockCredentialsInterface_GenerateTemporaryServiceCredential_Call struct {
	*mock.Call
}

// GenerateTemporaryServiceCredential is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.GenerateTemporaryServiceCredentialRequest
func (_e *MockCredentialsInterface_Expecter) GenerateTemporaryServiceCredential(ctx interface{}, request interface{}) *MockCredentialsInterface_GenerateTemporaryServiceCredential_Call {
	return &MockCredentialsInterface_GenerateTemporaryServiceCredential_Call{Call: _e.mock.On("GenerateTemporaryServiceCredential", ctx, request)}
}

func (_c *MockCredentialsInterface_GenerateTemporaryServiceCredential_Call) Run(run func(ctx context.Context, request catalog.GenerateTemporaryServiceCredentialRequest)) *MockCredentialsInterface_GenerateTemporaryServiceCredential_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.GenerateTemporaryServiceCredentialRequest))
	})
	return _c
}

func (_c *MockCredentialsInterface_GenerateTemporaryServiceCredential_Call) Return(_a0 *catalog.TemporaryCredentials, _a1 error) *MockCredentialsInterface_GenerateTemporaryServiceCredential_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCredentialsInterface_GenerateTemporaryServiceCredential_Call) RunAndReturn(run func(context.Context, catalog.GenerateTemporaryServiceCredentialRequest) (*catalog.TemporaryCredentials, error)) *MockCredentialsInterface_GenerateTemporaryServiceCredential_Call {
	_c.Call.Return(run)
	return _c
}

// GetCredential provides a mock function with given fields: ctx, request
func (_m *MockCredentialsInterface) GetCredential(ctx context.Context, request catalog.GetCredentialRequest) (*catalog.CredentialInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetCredential")
	}

	var r0 *catalog.CredentialInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetCredentialRequest) (*catalog.CredentialInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetCredentialRequest) *catalog.CredentialInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.CredentialInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.GetCredentialRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCredentialsInterface_GetCredential_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCredential'
type MockCredentialsInterface_GetCredential_Call struct {
	*mock.Call
}

// GetCredential is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.GetCredentialRequest
func (_e *MockCredentialsInterface_Expecter) GetCredential(ctx interface{}, request interface{}) *MockCredentialsInterface_GetCredential_Call {
	return &MockCredentialsInterface_GetCredential_Call{Call: _e.mock.On("GetCredential", ctx, request)}
}

func (_c *MockCredentialsInterface_GetCredential_Call) Run(run func(ctx context.Context, request catalog.GetCredentialRequest)) *MockCredentialsInterface_GetCredential_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.GetCredentialRequest))
	})
	return _c
}

func (_c *MockCredentialsInterface_GetCredential_Call) Return(_a0 *catalog.CredentialInfo, _a1 error) *MockCredentialsInterface_GetCredential_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCredentialsInterface_GetCredential_Call) RunAndReturn(run func(context.Context, catalog.GetCredentialRequest) (*catalog.CredentialInfo, error)) *MockCredentialsInterface_GetCredential_Call {
	_c.Call.Return(run)
	return _c
}

// GetCredentialByNameArg provides a mock function with given fields: ctx, nameArg
func (_m *MockCredentialsInterface) GetCredentialByNameArg(ctx context.Context, nameArg string) (*catalog.CredentialInfo, error) {
	ret := _m.Called(ctx, nameArg)

	if len(ret) == 0 {
		panic("no return value specified for GetCredentialByNameArg")
	}

	var r0 *catalog.CredentialInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*catalog.CredentialInfo, error)); ok {
		return rf(ctx, nameArg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *catalog.CredentialInfo); ok {
		r0 = rf(ctx, nameArg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.CredentialInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, nameArg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCredentialsInterface_GetCredentialByNameArg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCredentialByNameArg'
type MockCredentialsInterface_GetCredentialByNameArg_Call struct {
	*mock.Call
}

// GetCredentialByNameArg is a helper method to define mock.On call
//   - ctx context.Context
//   - nameArg string
func (_e *MockCredentialsInterface_Expecter) GetCredentialByNameArg(ctx interface{}, nameArg interface{}) *MockCredentialsInterface_GetCredentialByNameArg_Call {
	return &MockCredentialsInterface_GetCredentialByNameArg_Call{Call: _e.mock.On("GetCredentialByNameArg", ctx, nameArg)}
}

func (_c *MockCredentialsInterface_GetCredentialByNameArg_Call) Run(run func(ctx context.Context, nameArg string)) *MockCredentialsInterface_GetCredentialByNameArg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockCredentialsInterface_GetCredentialByNameArg_Call) Return(_a0 *catalog.CredentialInfo, _a1 error) *MockCredentialsInterface_GetCredentialByNameArg_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCredentialsInterface_GetCredentialByNameArg_Call) RunAndReturn(run func(context.Context, string) (*catalog.CredentialInfo, error)) *MockCredentialsInterface_GetCredentialByNameArg_Call {
	_c.Call.Return(run)
	return _c
}

// ListCredentials provides a mock function with given fields: ctx, request
func (_m *MockCredentialsInterface) ListCredentials(ctx context.Context, request catalog.ListCredentialsRequest) listing.Iterator[catalog.CredentialInfo] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListCredentials")
	}

	var r0 listing.Iterator[catalog.CredentialInfo]
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListCredentialsRequest) listing.Iterator[catalog.CredentialInfo]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[catalog.CredentialInfo])
		}
	}

	return r0
}

// MockCredentialsInterface_ListCredentials_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListCredentials'
type MockCredentialsInterface_ListCredentials_Call struct {
	*mock.Call
}

// ListCredentials is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListCredentialsRequest
func (_e *MockCredentialsInterface_Expecter) ListCredentials(ctx interface{}, request interface{}) *MockCredentialsInterface_ListCredentials_Call {
	return &MockCredentialsInterface_ListCredentials_Call{Call: _e.mock.On("ListCredentials", ctx, request)}
}

func (_c *MockCredentialsInterface_ListCredentials_Call) Run(run func(ctx context.Context, request catalog.ListCredentialsRequest)) *MockCredentialsInterface_ListCredentials_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListCredentialsRequest))
	})
	return _c
}

func (_c *MockCredentialsInterface_ListCredentials_Call) Return(_a0 listing.Iterator[catalog.CredentialInfo]) *MockCredentialsInterface_ListCredentials_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCredentialsInterface_ListCredentials_Call) RunAndReturn(run func(context.Context, catalog.ListCredentialsRequest) listing.Iterator[catalog.CredentialInfo]) *MockCredentialsInterface_ListCredentials_Call {
	_c.Call.Return(run)
	return _c
}

// ListCredentialsAll provides a mock function with given fields: ctx, request
func (_m *MockCredentialsInterface) ListCredentialsAll(ctx context.Context, request catalog.ListCredentialsRequest) ([]catalog.CredentialInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListCredentialsAll")
	}

	var r0 []catalog.CredentialInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListCredentialsRequest) ([]catalog.CredentialInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListCredentialsRequest) []catalog.CredentialInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]catalog.CredentialInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.ListCredentialsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCredentialsInterface_ListCredentialsAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListCredentialsAll'
type MockCredentialsInterface_ListCredentialsAll_Call struct {
	*mock.Call
}

// ListCredentialsAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListCredentialsRequest
func (_e *MockCredentialsInterface_Expecter) ListCredentialsAll(ctx interface{}, request interface{}) *MockCredentialsInterface_ListCredentialsAll_Call {
	return &MockCredentialsInterface_ListCredentialsAll_Call{Call: _e.mock.On("ListCredentialsAll", ctx, request)}
}

func (_c *MockCredentialsInterface_ListCredentialsAll_Call) Run(run func(ctx context.Context, request catalog.ListCredentialsRequest)) *MockCredentialsInterface_ListCredentialsAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListCredentialsRequest))
	})
	return _c
}

func (_c *MockCredentialsInterface_ListCredentialsAll_Call) Return(_a0 []catalog.CredentialInfo, _a1 error) *MockCredentialsInterface_ListCredentialsAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCredentialsInterface_ListCredentialsAll_Call) RunAndReturn(run func(context.Context, catalog.ListCredentialsRequest) ([]catalog.CredentialInfo, error)) *MockCredentialsInterface_ListCredentialsAll_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateCredential provides a mock function with given fields: ctx, request
func (_m *MockCredentialsInterface) UpdateCredential(ctx context.Context, request catalog.UpdateCredentialRequest) (*catalog.CredentialInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCredential")
	}

	var r0 *catalog.CredentialInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdateCredentialRequest) (*catalog.CredentialInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdateCredentialRequest) *catalog.CredentialInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.CredentialInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.UpdateCredentialRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCredentialsInterface_UpdateCredential_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateCredential'
type MockCredentialsInterface_UpdateCredential_Call struct {
	*mock.Call
}

// UpdateCredential is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.UpdateCredentialRequest
func (_e *MockCredentialsInterface_Expecter) UpdateCredential(ctx interface{}, request interface{}) *MockCredentialsInterface_UpdateCredential_Call {
	return &MockCredentialsInterface_UpdateCredential_Call{Call: _e.mock.On("UpdateCredential", ctx, request)}
}

func (_c *MockCredentialsInterface_UpdateCredential_Call) Run(run func(ctx context.Context, request catalog.UpdateCredentialRequest)) *MockCredentialsInterface_UpdateCredential_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.UpdateCredentialRequest))
	})
	return _c
}

func (_c *MockCredentialsInterface_UpdateCredential_Call) Return(_a0 *catalog.CredentialInfo, _a1 error) *MockCredentialsInterface_UpdateCredential_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCredentialsInterface_UpdateCredential_Call) RunAndReturn(run func(context.Context, catalog.UpdateCredentialRequest) (*catalog.CredentialInfo, error)) *MockCredentialsInterface_UpdateCredential_Call {
	_c.Call.Return(run)
	return _c
}

// ValidateCredential provides a mock function with given fields: ctx, request
func (_m *MockCredentialsInterface) ValidateCredential(ctx context.Context, request catalog.ValidateCredentialRequest) (*catalog.ValidateCredentialResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ValidateCredential")
	}

	var r0 *catalog.ValidateCredentialResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ValidateCredentialRequest) (*catalog.ValidateCredentialResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ValidateCredentialRequest) *catalog.ValidateCredentialResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.ValidateCredentialResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.ValidateCredentialRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCredentialsInterface_ValidateCredential_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateCredential'
type MockCredentialsInterface_ValidateCredential_Call struct {
	*mock.Call
}

// ValidateCredential is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ValidateCredentialRequest
func (_e *MockCredentialsInterface_Expecter) ValidateCredential(ctx interface{}, request interface{}) *MockCredentialsInterface_ValidateCredential_Call {
	return &MockCredentialsInterface_ValidateCredential_Call{Call: _e.mock.On("ValidateCredential", ctx, request)}
}

func (_c *MockCredentialsInterface_ValidateCredential_Call) Run(run func(ctx context.Context, request catalog.ValidateCredentialRequest)) *MockCredentialsInterface_ValidateCredential_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ValidateCredentialRequest))
	})
	return _c
}

func (_c *MockCredentialsInterface_ValidateCredential_Call) Return(_a0 *catalog.ValidateCredentialResponse, _a1 error) *MockCredentialsInterface_ValidateCredential_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCredentialsInterface_ValidateCredential_Call) RunAndReturn(run func(context.Context, catalog.ValidateCredentialRequest) (*catalog.ValidateCredentialResponse, error)) *MockCredentialsInterface_ValidateCredential_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCredentialsInterface creates a new instance of MockCredentialsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCredentialsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCredentialsInterface {
	mock := &MockCredentialsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
