// Code generated by mockery v2.38.0. DO NOT EDIT.

package oauth2

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	mock "github.com/stretchr/testify/mock"

	oauth2 "github.com/databricks/databricks-sdk-go/service/oauth2"
)

// MockPublishedAppIntegrationAPIInterface is an autogenerated mock type for the PublishedAppIntegrationAPIInterface type
type MockPublishedAppIntegrationAPIInterface struct {
	mock.Mock
}

type MockPublishedAppIntegrationAPIInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPublishedAppIntegrationAPIInterface) EXPECT() *MockPublishedAppIntegrationAPIInterface_Expecter {
	return &MockPublishedAppIntegrationAPIInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockPublishedAppIntegrationAPIInterface) Create(ctx context.Context, request oauth2.CreatePublishedAppIntegration) (*oauth2.CreatePublishedAppIntegrationOutput, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *oauth2.CreatePublishedAppIntegrationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.CreatePublishedAppIntegration) (*oauth2.CreatePublishedAppIntegrationOutput, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.CreatePublishedAppIntegration) *oauth2.CreatePublishedAppIntegrationOutput); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.CreatePublishedAppIntegrationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, oauth2.CreatePublishedAppIntegration) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPublishedAppIntegrationAPIInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockPublishedAppIntegrationAPIInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request oauth2.CreatePublishedAppIntegration
func (_e *MockPublishedAppIntegrationAPIInterface_Expecter) Create(ctx interface{}, request interface{}) *MockPublishedAppIntegrationAPIInterface_Create_Call {
	return &MockPublishedAppIntegrationAPIInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockPublishedAppIntegrationAPIInterface_Create_Call) Run(run func(ctx context.Context, request oauth2.CreatePublishedAppIntegration)) *MockPublishedAppIntegrationAPIInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oauth2.CreatePublishedAppIntegration))
	})
	return _c
}

func (_c *MockPublishedAppIntegrationAPIInterface_Create_Call) Return(_a0 *oauth2.CreatePublishedAppIntegrationOutput, _a1 error) *MockPublishedAppIntegrationAPIInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPublishedAppIntegrationAPIInterface_Create_Call) RunAndReturn(run func(context.Context, oauth2.CreatePublishedAppIntegration) (*oauth2.CreatePublishedAppIntegrationOutput, error)) *MockPublishedAppIntegrationAPIInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockPublishedAppIntegrationAPIInterface) Delete(ctx context.Context, request oauth2.DeletePublishedAppIntegrationRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.DeletePublishedAppIntegrationRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPublishedAppIntegrationAPIInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockPublishedAppIntegrationAPIInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request oauth2.DeletePublishedAppIntegrationRequest
func (_e *MockPublishedAppIntegrationAPIInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockPublishedAppIntegrationAPIInterface_Delete_Call {
	return &MockPublishedAppIntegrationAPIInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockPublishedAppIntegrationAPIInterface_Delete_Call) Run(run func(ctx context.Context, request oauth2.DeletePublishedAppIntegrationRequest)) *MockPublishedAppIntegrationAPIInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oauth2.DeletePublishedAppIntegrationRequest))
	})
	return _c
}

func (_c *MockPublishedAppIntegrationAPIInterface_Delete_Call) Return(_a0 error) *MockPublishedAppIntegrationAPIInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPublishedAppIntegrationAPIInterface_Delete_Call) RunAndReturn(run func(context.Context, oauth2.DeletePublishedAppIntegrationRequest) error) *MockPublishedAppIntegrationAPIInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByIntegrationId provides a mock function with given fields: ctx, integrationId
func (_m *MockPublishedAppIntegrationAPIInterface) DeleteByIntegrationId(ctx context.Context, integrationId string) error {
	ret := _m.Called(ctx, integrationId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByIntegrationId")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, integrationId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPublishedAppIntegrationAPIInterface_DeleteByIntegrationId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByIntegrationId'
type MockPublishedAppIntegrationAPIInterface_DeleteByIntegrationId_Call struct {
	*mock.Call
}

// DeleteByIntegrationId is a helper method to define mock.On call
//   - ctx context.Context
//   - integrationId string
func (_e *MockPublishedAppIntegrationAPIInterface_Expecter) DeleteByIntegrationId(ctx interface{}, integrationId interface{}) *MockPublishedAppIntegrationAPIInterface_DeleteByIntegrationId_Call {
	return &MockPublishedAppIntegrationAPIInterface_DeleteByIntegrationId_Call{Call: _e.mock.On("DeleteByIntegrationId", ctx, integrationId)}
}

func (_c *MockPublishedAppIntegrationAPIInterface_DeleteByIntegrationId_Call) Run(run func(ctx context.Context, integrationId string)) *MockPublishedAppIntegrationAPIInterface_DeleteByIntegrationId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockPublishedAppIntegrationAPIInterface_DeleteByIntegrationId_Call) Return(_a0 error) *MockPublishedAppIntegrationAPIInterface_DeleteByIntegrationId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPublishedAppIntegrationAPIInterface_DeleteByIntegrationId_Call) RunAndReturn(run func(context.Context, string) error) *MockPublishedAppIntegrationAPIInterface_DeleteByIntegrationId_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockPublishedAppIntegrationAPIInterface) Get(ctx context.Context, request oauth2.GetPublishedAppIntegrationRequest) (*oauth2.GetPublishedAppIntegrationOutput, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *oauth2.GetPublishedAppIntegrationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.GetPublishedAppIntegrationRequest) (*oauth2.GetPublishedAppIntegrationOutput, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.GetPublishedAppIntegrationRequest) *oauth2.GetPublishedAppIntegrationOutput); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.GetPublishedAppIntegrationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, oauth2.GetPublishedAppIntegrationRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPublishedAppIntegrationAPIInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockPublishedAppIntegrationAPIInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request oauth2.GetPublishedAppIntegrationRequest
func (_e *MockPublishedAppIntegrationAPIInterface_Expecter) Get(ctx interface{}, request interface{}) *MockPublishedAppIntegrationAPIInterface_Get_Call {
	return &MockPublishedAppIntegrationAPIInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockPublishedAppIntegrationAPIInterface_Get_Call) Run(run func(ctx context.Context, request oauth2.GetPublishedAppIntegrationRequest)) *MockPublishedAppIntegrationAPIInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oauth2.GetPublishedAppIntegrationRequest))
	})
	return _c
}

func (_c *MockPublishedAppIntegrationAPIInterface_Get_Call) Return(_a0 *oauth2.GetPublishedAppIntegrationOutput, _a1 error) *MockPublishedAppIntegrationAPIInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPublishedAppIntegrationAPIInterface_Get_Call) RunAndReturn(run func(context.Context, oauth2.GetPublishedAppIntegrationRequest) (*oauth2.GetPublishedAppIntegrationOutput, error)) *MockPublishedAppIntegrationAPIInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByIntegrationId provides a mock function with given fields: ctx, integrationId
func (_m *MockPublishedAppIntegrationAPIInterface) GetByIntegrationId(ctx context.Context, integrationId string) (*oauth2.GetPublishedAppIntegrationOutput, error) {
	ret := _m.Called(ctx, integrationId)

	if len(ret) == 0 {
		panic("no return value specified for GetByIntegrationId")
	}

	var r0 *oauth2.GetPublishedAppIntegrationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*oauth2.GetPublishedAppIntegrationOutput, error)); ok {
		return rf(ctx, integrationId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *oauth2.GetPublishedAppIntegrationOutput); ok {
		r0 = rf(ctx, integrationId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.GetPublishedAppIntegrationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, integrationId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPublishedAppIntegrationAPIInterface_GetByIntegrationId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByIntegrationId'
type MockPublishedAppIntegrationAPIInterface_GetByIntegrationId_Call struct {
	*mock.Call
}

// GetByIntegrationId is a helper method to define mock.On call
//   - ctx context.Context
//   - integrationId string
func (_e *MockPublishedAppIntegrationAPIInterface_Expecter) GetByIntegrationId(ctx interface{}, integrationId interface{}) *MockPublishedAppIntegrationAPIInterface_GetByIntegrationId_Call {
	return &MockPublishedAppIntegrationAPIInterface_GetByIntegrationId_Call{Call: _e.mock.On("GetByIntegrationId", ctx, integrationId)}
}

func (_c *MockPublishedAppIntegrationAPIInterface_GetByIntegrationId_Call) Run(run func(ctx context.Context, integrationId string)) *MockPublishedAppIntegrationAPIInterface_GetByIntegrationId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockPublishedAppIntegrationAPIInterface_GetByIntegrationId_Call) Return(_a0 *oauth2.GetPublishedAppIntegrationOutput, _a1 error) *MockPublishedAppIntegrationAPIInterface_GetByIntegrationId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPublishedAppIntegrationAPIInterface_GetByIntegrationId_Call) RunAndReturn(run func(context.Context, string) (*oauth2.GetPublishedAppIntegrationOutput, error)) *MockPublishedAppIntegrationAPIInterface_GetByIntegrationId_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockPublishedAppIntegrationAPIInterface) Impl() oauth2.PublishedAppIntegrationService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 oauth2.PublishedAppIntegrationService
	if rf, ok := ret.Get(0).(func() oauth2.PublishedAppIntegrationService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(oauth2.PublishedAppIntegrationService)
		}
	}

	return r0
}

// MockPublishedAppIntegrationAPIInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockPublishedAppIntegrationAPIInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockPublishedAppIntegrationAPIInterface_Expecter) Impl() *MockPublishedAppIntegrationAPIInterface_Impl_Call {
	return &MockPublishedAppIntegrationAPIInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockPublishedAppIntegrationAPIInterface_Impl_Call) Run(run func()) *MockPublishedAppIntegrationAPIInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPublishedAppIntegrationAPIInterface_Impl_Call) Return(_a0 oauth2.PublishedAppIntegrationService) *MockPublishedAppIntegrationAPIInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPublishedAppIntegrationAPIInterface_Impl_Call) RunAndReturn(run func() oauth2.PublishedAppIntegrationService) *MockPublishedAppIntegrationAPIInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx
func (_m *MockPublishedAppIntegrationAPIInterface) List(ctx context.Context) *listing.PaginatingIterator[struct{}, *oauth2.GetPublishedAppIntegrationsOutput, oauth2.GetPublishedAppIntegrationOutput] {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *listing.PaginatingIterator[struct{}, *oauth2.GetPublishedAppIntegrationsOutput, oauth2.GetPublishedAppIntegrationOutput]
	if rf, ok := ret.Get(0).(func(context.Context) *listing.PaginatingIterator[struct{}, *oauth2.GetPublishedAppIntegrationsOutput, oauth2.GetPublishedAppIntegrationOutput]); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*listing.PaginatingIterator[struct{}, *oauth2.GetPublishedAppIntegrationsOutput, oauth2.GetPublishedAppIntegrationOutput])
		}
	}

	return r0
}

// MockPublishedAppIntegrationAPIInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockPublishedAppIntegrationAPIInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockPublishedAppIntegrationAPIInterface_Expecter) List(ctx interface{}) *MockPublishedAppIntegrationAPIInterface_List_Call {
	return &MockPublishedAppIntegrationAPIInterface_List_Call{Call: _e.mock.On("List", ctx)}
}

func (_c *MockPublishedAppIntegrationAPIInterface_List_Call) Run(run func(ctx context.Context)) *MockPublishedAppIntegrationAPIInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockPublishedAppIntegrationAPIInterface_List_Call) Return(_a0 *listing.PaginatingIterator[struct{}, *oauth2.GetPublishedAppIntegrationsOutput, oauth2.GetPublishedAppIntegrationOutput]) *MockPublishedAppIntegrationAPIInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPublishedAppIntegrationAPIInterface_List_Call) RunAndReturn(run func(context.Context) *listing.PaginatingIterator[struct{}, *oauth2.GetPublishedAppIntegrationsOutput, oauth2.GetPublishedAppIntegrationOutput]) *MockPublishedAppIntegrationAPIInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx
func (_m *MockPublishedAppIntegrationAPIInterface) ListAll(ctx context.Context) ([]oauth2.GetPublishedAppIntegrationOutput, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []oauth2.GetPublishedAppIntegrationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]oauth2.GetPublishedAppIntegrationOutput, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []oauth2.GetPublishedAppIntegrationOutput); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]oauth2.GetPublishedAppIntegrationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPublishedAppIntegrationAPIInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockPublishedAppIntegrationAPIInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockPublishedAppIntegrationAPIInterface_Expecter) ListAll(ctx interface{}) *MockPublishedAppIntegrationAPIInterface_ListAll_Call {
	return &MockPublishedAppIntegrationAPIInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx)}
}

func (_c *MockPublishedAppIntegrationAPIInterface_ListAll_Call) Run(run func(ctx context.Context)) *MockPublishedAppIntegrationAPIInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockPublishedAppIntegrationAPIInterface_ListAll_Call) Return(_a0 []oauth2.GetPublishedAppIntegrationOutput, _a1 error) *MockPublishedAppIntegrationAPIInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPublishedAppIntegrationAPIInterface_ListAll_Call) RunAndReturn(run func(context.Context) ([]oauth2.GetPublishedAppIntegrationOutput, error)) *MockPublishedAppIntegrationAPIInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockPublishedAppIntegrationAPIInterface) Update(ctx context.Context, request oauth2.UpdatePublishedAppIntegration) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.UpdatePublishedAppIntegration) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPublishedAppIntegrationAPIInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockPublishedAppIntegrationAPIInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request oauth2.UpdatePublishedAppIntegration
func (_e *MockPublishedAppIntegrationAPIInterface_Expecter) Update(ctx interface{}, request interface{}) *MockPublishedAppIntegrationAPIInterface_Update_Call {
	return &MockPublishedAppIntegrationAPIInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockPublishedAppIntegrationAPIInterface_Update_Call) Run(run func(ctx context.Context, request oauth2.UpdatePublishedAppIntegration)) *MockPublishedAppIntegrationAPIInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oauth2.UpdatePublishedAppIntegration))
	})
	return _c
}

func (_c *MockPublishedAppIntegrationAPIInterface_Update_Call) Return(_a0 error) *MockPublishedAppIntegrationAPIInterface_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPublishedAppIntegrationAPIInterface_Update_Call) RunAndReturn(run func(context.Context, oauth2.UpdatePublishedAppIntegration) error) *MockPublishedAppIntegrationAPIInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockPublishedAppIntegrationAPIInterface) WithImpl(impl oauth2.PublishedAppIntegrationService) oauth2.PublishedAppIntegrationAPIInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 oauth2.PublishedAppIntegrationAPIInterface
	if rf, ok := ret.Get(0).(func(oauth2.PublishedAppIntegrationService) oauth2.PublishedAppIntegrationAPIInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(oauth2.PublishedAppIntegrationAPIInterface)
		}
	}

	return r0
}

// MockPublishedAppIntegrationAPIInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockPublishedAppIntegrationAPIInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl oauth2.PublishedAppIntegrationService
func (_e *MockPublishedAppIntegrationAPIInterface_Expecter) WithImpl(impl interface{}) *MockPublishedAppIntegrationAPIInterface_WithImpl_Call {
	return &MockPublishedAppIntegrationAPIInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockPublishedAppIntegrationAPIInterface_WithImpl_Call) Run(run func(impl oauth2.PublishedAppIntegrationService)) *MockPublishedAppIntegrationAPIInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(oauth2.PublishedAppIntegrationService))
	})
	return _c
}

func (_c *MockPublishedAppIntegrationAPIInterface_WithImpl_Call) Return(_a0 oauth2.PublishedAppIntegrationAPIInterface) *MockPublishedAppIntegrationAPIInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPublishedAppIntegrationAPIInterface_WithImpl_Call) RunAndReturn(run func(oauth2.PublishedAppIntegrationService) oauth2.PublishedAppIntegrationAPIInterface) *MockPublishedAppIntegrationAPIInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPublishedAppIntegrationAPIInterface creates a new instance of MockPublishedAppIntegrationAPIInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPublishedAppIntegrationAPIInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPublishedAppIntegrationAPIInterface {
	mock := &MockPublishedAppIntegrationAPIInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
