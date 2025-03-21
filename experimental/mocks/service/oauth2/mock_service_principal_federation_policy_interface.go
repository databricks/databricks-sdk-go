// Code generated by mockery v2.53.2. DO NOT EDIT.

package oauth2

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	mock "github.com/stretchr/testify/mock"

	oauth2 "github.com/databricks/databricks-sdk-go/service/oauth2"
)

// MockServicePrincipalFederationPolicyInterface is an autogenerated mock type for the ServicePrincipalFederationPolicyInterface type
type MockServicePrincipalFederationPolicyInterface struct {
	mock.Mock
}

type MockServicePrincipalFederationPolicyInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockServicePrincipalFederationPolicyInterface) EXPECT() *MockServicePrincipalFederationPolicyInterface_Expecter {
	return &MockServicePrincipalFederationPolicyInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockServicePrincipalFederationPolicyInterface) Create(ctx context.Context, request oauth2.CreateServicePrincipalFederationPolicyRequest) (*oauth2.FederationPolicy, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *oauth2.FederationPolicy
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.CreateServicePrincipalFederationPolicyRequest) (*oauth2.FederationPolicy, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.CreateServicePrincipalFederationPolicyRequest) *oauth2.FederationPolicy); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.FederationPolicy)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, oauth2.CreateServicePrincipalFederationPolicyRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockServicePrincipalFederationPolicyInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockServicePrincipalFederationPolicyInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request oauth2.CreateServicePrincipalFederationPolicyRequest
func (_e *MockServicePrincipalFederationPolicyInterface_Expecter) Create(ctx interface{}, request interface{}) *MockServicePrincipalFederationPolicyInterface_Create_Call {
	return &MockServicePrincipalFederationPolicyInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockServicePrincipalFederationPolicyInterface_Create_Call) Run(run func(ctx context.Context, request oauth2.CreateServicePrincipalFederationPolicyRequest)) *MockServicePrincipalFederationPolicyInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oauth2.CreateServicePrincipalFederationPolicyRequest))
	})
	return _c
}

func (_c *MockServicePrincipalFederationPolicyInterface_Create_Call) Return(_a0 *oauth2.FederationPolicy, _a1 error) *MockServicePrincipalFederationPolicyInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServicePrincipalFederationPolicyInterface_Create_Call) RunAndReturn(run func(context.Context, oauth2.CreateServicePrincipalFederationPolicyRequest) (*oauth2.FederationPolicy, error)) *MockServicePrincipalFederationPolicyInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockServicePrincipalFederationPolicyInterface) Delete(ctx context.Context, request oauth2.DeleteServicePrincipalFederationPolicyRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.DeleteServicePrincipalFederationPolicyRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockServicePrincipalFederationPolicyInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockServicePrincipalFederationPolicyInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request oauth2.DeleteServicePrincipalFederationPolicyRequest
func (_e *MockServicePrincipalFederationPolicyInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockServicePrincipalFederationPolicyInterface_Delete_Call {
	return &MockServicePrincipalFederationPolicyInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockServicePrincipalFederationPolicyInterface_Delete_Call) Run(run func(ctx context.Context, request oauth2.DeleteServicePrincipalFederationPolicyRequest)) *MockServicePrincipalFederationPolicyInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oauth2.DeleteServicePrincipalFederationPolicyRequest))
	})
	return _c
}

func (_c *MockServicePrincipalFederationPolicyInterface_Delete_Call) Return(_a0 error) *MockServicePrincipalFederationPolicyInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServicePrincipalFederationPolicyInterface_Delete_Call) RunAndReturn(run func(context.Context, oauth2.DeleteServicePrincipalFederationPolicyRequest) error) *MockServicePrincipalFederationPolicyInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByServicePrincipalIdAndPolicyId provides a mock function with given fields: ctx, servicePrincipalId, policyId
func (_m *MockServicePrincipalFederationPolicyInterface) DeleteByServicePrincipalIdAndPolicyId(ctx context.Context, servicePrincipalId int64, policyId string) error {
	ret := _m.Called(ctx, servicePrincipalId, policyId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByServicePrincipalIdAndPolicyId")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) error); ok {
		r0 = rf(ctx, servicePrincipalId, policyId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockServicePrincipalFederationPolicyInterface_DeleteByServicePrincipalIdAndPolicyId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByServicePrincipalIdAndPolicyId'
type MockServicePrincipalFederationPolicyInterface_DeleteByServicePrincipalIdAndPolicyId_Call struct {
	*mock.Call
}

// DeleteByServicePrincipalIdAndPolicyId is a helper method to define mock.On call
//   - ctx context.Context
//   - servicePrincipalId int64
//   - policyId string
func (_e *MockServicePrincipalFederationPolicyInterface_Expecter) DeleteByServicePrincipalIdAndPolicyId(ctx interface{}, servicePrincipalId interface{}, policyId interface{}) *MockServicePrincipalFederationPolicyInterface_DeleteByServicePrincipalIdAndPolicyId_Call {
	return &MockServicePrincipalFederationPolicyInterface_DeleteByServicePrincipalIdAndPolicyId_Call{Call: _e.mock.On("DeleteByServicePrincipalIdAndPolicyId", ctx, servicePrincipalId, policyId)}
}

func (_c *MockServicePrincipalFederationPolicyInterface_DeleteByServicePrincipalIdAndPolicyId_Call) Run(run func(ctx context.Context, servicePrincipalId int64, policyId string)) *MockServicePrincipalFederationPolicyInterface_DeleteByServicePrincipalIdAndPolicyId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(string))
	})
	return _c
}

func (_c *MockServicePrincipalFederationPolicyInterface_DeleteByServicePrincipalIdAndPolicyId_Call) Return(_a0 error) *MockServicePrincipalFederationPolicyInterface_DeleteByServicePrincipalIdAndPolicyId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServicePrincipalFederationPolicyInterface_DeleteByServicePrincipalIdAndPolicyId_Call) RunAndReturn(run func(context.Context, int64, string) error) *MockServicePrincipalFederationPolicyInterface_DeleteByServicePrincipalIdAndPolicyId_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockServicePrincipalFederationPolicyInterface) Get(ctx context.Context, request oauth2.GetServicePrincipalFederationPolicyRequest) (*oauth2.FederationPolicy, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *oauth2.FederationPolicy
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.GetServicePrincipalFederationPolicyRequest) (*oauth2.FederationPolicy, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.GetServicePrincipalFederationPolicyRequest) *oauth2.FederationPolicy); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.FederationPolicy)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, oauth2.GetServicePrincipalFederationPolicyRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockServicePrincipalFederationPolicyInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockServicePrincipalFederationPolicyInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request oauth2.GetServicePrincipalFederationPolicyRequest
func (_e *MockServicePrincipalFederationPolicyInterface_Expecter) Get(ctx interface{}, request interface{}) *MockServicePrincipalFederationPolicyInterface_Get_Call {
	return &MockServicePrincipalFederationPolicyInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockServicePrincipalFederationPolicyInterface_Get_Call) Run(run func(ctx context.Context, request oauth2.GetServicePrincipalFederationPolicyRequest)) *MockServicePrincipalFederationPolicyInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oauth2.GetServicePrincipalFederationPolicyRequest))
	})
	return _c
}

func (_c *MockServicePrincipalFederationPolicyInterface_Get_Call) Return(_a0 *oauth2.FederationPolicy, _a1 error) *MockServicePrincipalFederationPolicyInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServicePrincipalFederationPolicyInterface_Get_Call) RunAndReturn(run func(context.Context, oauth2.GetServicePrincipalFederationPolicyRequest) (*oauth2.FederationPolicy, error)) *MockServicePrincipalFederationPolicyInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByServicePrincipalIdAndPolicyId provides a mock function with given fields: ctx, servicePrincipalId, policyId
func (_m *MockServicePrincipalFederationPolicyInterface) GetByServicePrincipalIdAndPolicyId(ctx context.Context, servicePrincipalId int64, policyId string) (*oauth2.FederationPolicy, error) {
	ret := _m.Called(ctx, servicePrincipalId, policyId)

	if len(ret) == 0 {
		panic("no return value specified for GetByServicePrincipalIdAndPolicyId")
	}

	var r0 *oauth2.FederationPolicy
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) (*oauth2.FederationPolicy, error)); ok {
		return rf(ctx, servicePrincipalId, policyId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) *oauth2.FederationPolicy); ok {
		r0 = rf(ctx, servicePrincipalId, policyId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.FederationPolicy)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, servicePrincipalId, policyId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockServicePrincipalFederationPolicyInterface_GetByServicePrincipalIdAndPolicyId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByServicePrincipalIdAndPolicyId'
type MockServicePrincipalFederationPolicyInterface_GetByServicePrincipalIdAndPolicyId_Call struct {
	*mock.Call
}

// GetByServicePrincipalIdAndPolicyId is a helper method to define mock.On call
//   - ctx context.Context
//   - servicePrincipalId int64
//   - policyId string
func (_e *MockServicePrincipalFederationPolicyInterface_Expecter) GetByServicePrincipalIdAndPolicyId(ctx interface{}, servicePrincipalId interface{}, policyId interface{}) *MockServicePrincipalFederationPolicyInterface_GetByServicePrincipalIdAndPolicyId_Call {
	return &MockServicePrincipalFederationPolicyInterface_GetByServicePrincipalIdAndPolicyId_Call{Call: _e.mock.On("GetByServicePrincipalIdAndPolicyId", ctx, servicePrincipalId, policyId)}
}

func (_c *MockServicePrincipalFederationPolicyInterface_GetByServicePrincipalIdAndPolicyId_Call) Run(run func(ctx context.Context, servicePrincipalId int64, policyId string)) *MockServicePrincipalFederationPolicyInterface_GetByServicePrincipalIdAndPolicyId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(string))
	})
	return _c
}

func (_c *MockServicePrincipalFederationPolicyInterface_GetByServicePrincipalIdAndPolicyId_Call) Return(_a0 *oauth2.FederationPolicy, _a1 error) *MockServicePrincipalFederationPolicyInterface_GetByServicePrincipalIdAndPolicyId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServicePrincipalFederationPolicyInterface_GetByServicePrincipalIdAndPolicyId_Call) RunAndReturn(run func(context.Context, int64, string) (*oauth2.FederationPolicy, error)) *MockServicePrincipalFederationPolicyInterface_GetByServicePrincipalIdAndPolicyId_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockServicePrincipalFederationPolicyInterface) List(ctx context.Context, request oauth2.ListServicePrincipalFederationPoliciesRequest) listing.Iterator[oauth2.FederationPolicy] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[oauth2.FederationPolicy]
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.ListServicePrincipalFederationPoliciesRequest) listing.Iterator[oauth2.FederationPolicy]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[oauth2.FederationPolicy])
		}
	}

	return r0
}

// MockServicePrincipalFederationPolicyInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockServicePrincipalFederationPolicyInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request oauth2.ListServicePrincipalFederationPoliciesRequest
func (_e *MockServicePrincipalFederationPolicyInterface_Expecter) List(ctx interface{}, request interface{}) *MockServicePrincipalFederationPolicyInterface_List_Call {
	return &MockServicePrincipalFederationPolicyInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockServicePrincipalFederationPolicyInterface_List_Call) Run(run func(ctx context.Context, request oauth2.ListServicePrincipalFederationPoliciesRequest)) *MockServicePrincipalFederationPolicyInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oauth2.ListServicePrincipalFederationPoliciesRequest))
	})
	return _c
}

func (_c *MockServicePrincipalFederationPolicyInterface_List_Call) Return(_a0 listing.Iterator[oauth2.FederationPolicy]) *MockServicePrincipalFederationPolicyInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServicePrincipalFederationPolicyInterface_List_Call) RunAndReturn(run func(context.Context, oauth2.ListServicePrincipalFederationPoliciesRequest) listing.Iterator[oauth2.FederationPolicy]) *MockServicePrincipalFederationPolicyInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockServicePrincipalFederationPolicyInterface) ListAll(ctx context.Context, request oauth2.ListServicePrincipalFederationPoliciesRequest) ([]oauth2.FederationPolicy, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []oauth2.FederationPolicy
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.ListServicePrincipalFederationPoliciesRequest) ([]oauth2.FederationPolicy, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.ListServicePrincipalFederationPoliciesRequest) []oauth2.FederationPolicy); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]oauth2.FederationPolicy)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, oauth2.ListServicePrincipalFederationPoliciesRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockServicePrincipalFederationPolicyInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockServicePrincipalFederationPolicyInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request oauth2.ListServicePrincipalFederationPoliciesRequest
func (_e *MockServicePrincipalFederationPolicyInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockServicePrincipalFederationPolicyInterface_ListAll_Call {
	return &MockServicePrincipalFederationPolicyInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockServicePrincipalFederationPolicyInterface_ListAll_Call) Run(run func(ctx context.Context, request oauth2.ListServicePrincipalFederationPoliciesRequest)) *MockServicePrincipalFederationPolicyInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oauth2.ListServicePrincipalFederationPoliciesRequest))
	})
	return _c
}

func (_c *MockServicePrincipalFederationPolicyInterface_ListAll_Call) Return(_a0 []oauth2.FederationPolicy, _a1 error) *MockServicePrincipalFederationPolicyInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServicePrincipalFederationPolicyInterface_ListAll_Call) RunAndReturn(run func(context.Context, oauth2.ListServicePrincipalFederationPoliciesRequest) ([]oauth2.FederationPolicy, error)) *MockServicePrincipalFederationPolicyInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// ListByServicePrincipalId provides a mock function with given fields: ctx, servicePrincipalId
func (_m *MockServicePrincipalFederationPolicyInterface) ListByServicePrincipalId(ctx context.Context, servicePrincipalId int64) (*oauth2.ListFederationPoliciesResponse, error) {
	ret := _m.Called(ctx, servicePrincipalId)

	if len(ret) == 0 {
		panic("no return value specified for ListByServicePrincipalId")
	}

	var r0 *oauth2.ListFederationPoliciesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*oauth2.ListFederationPoliciesResponse, error)); ok {
		return rf(ctx, servicePrincipalId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *oauth2.ListFederationPoliciesResponse); ok {
		r0 = rf(ctx, servicePrincipalId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.ListFederationPoliciesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, servicePrincipalId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockServicePrincipalFederationPolicyInterface_ListByServicePrincipalId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListByServicePrincipalId'
type MockServicePrincipalFederationPolicyInterface_ListByServicePrincipalId_Call struct {
	*mock.Call
}

// ListByServicePrincipalId is a helper method to define mock.On call
//   - ctx context.Context
//   - servicePrincipalId int64
func (_e *MockServicePrincipalFederationPolicyInterface_Expecter) ListByServicePrincipalId(ctx interface{}, servicePrincipalId interface{}) *MockServicePrincipalFederationPolicyInterface_ListByServicePrincipalId_Call {
	return &MockServicePrincipalFederationPolicyInterface_ListByServicePrincipalId_Call{Call: _e.mock.On("ListByServicePrincipalId", ctx, servicePrincipalId)}
}

func (_c *MockServicePrincipalFederationPolicyInterface_ListByServicePrincipalId_Call) Run(run func(ctx context.Context, servicePrincipalId int64)) *MockServicePrincipalFederationPolicyInterface_ListByServicePrincipalId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockServicePrincipalFederationPolicyInterface_ListByServicePrincipalId_Call) Return(_a0 *oauth2.ListFederationPoliciesResponse, _a1 error) *MockServicePrincipalFederationPolicyInterface_ListByServicePrincipalId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServicePrincipalFederationPolicyInterface_ListByServicePrincipalId_Call) RunAndReturn(run func(context.Context, int64) (*oauth2.ListFederationPoliciesResponse, error)) *MockServicePrincipalFederationPolicyInterface_ListByServicePrincipalId_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockServicePrincipalFederationPolicyInterface) Update(ctx context.Context, request oauth2.UpdateServicePrincipalFederationPolicyRequest) (*oauth2.FederationPolicy, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *oauth2.FederationPolicy
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.UpdateServicePrincipalFederationPolicyRequest) (*oauth2.FederationPolicy, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.UpdateServicePrincipalFederationPolicyRequest) *oauth2.FederationPolicy); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.FederationPolicy)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, oauth2.UpdateServicePrincipalFederationPolicyRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockServicePrincipalFederationPolicyInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockServicePrincipalFederationPolicyInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request oauth2.UpdateServicePrincipalFederationPolicyRequest
func (_e *MockServicePrincipalFederationPolicyInterface_Expecter) Update(ctx interface{}, request interface{}) *MockServicePrincipalFederationPolicyInterface_Update_Call {
	return &MockServicePrincipalFederationPolicyInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockServicePrincipalFederationPolicyInterface_Update_Call) Run(run func(ctx context.Context, request oauth2.UpdateServicePrincipalFederationPolicyRequest)) *MockServicePrincipalFederationPolicyInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oauth2.UpdateServicePrincipalFederationPolicyRequest))
	})
	return _c
}

func (_c *MockServicePrincipalFederationPolicyInterface_Update_Call) Return(_a0 *oauth2.FederationPolicy, _a1 error) *MockServicePrincipalFederationPolicyInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServicePrincipalFederationPolicyInterface_Update_Call) RunAndReturn(run func(context.Context, oauth2.UpdateServicePrincipalFederationPolicyRequest) (*oauth2.FederationPolicy, error)) *MockServicePrincipalFederationPolicyInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockServicePrincipalFederationPolicyInterface creates a new instance of MockServicePrincipalFederationPolicyInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockServicePrincipalFederationPolicyInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockServicePrincipalFederationPolicyInterface {
	mock := &MockServicePrincipalFederationPolicyInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
