// Code generated by mockery v2.43.0. DO NOT EDIT.

package billing

import (
	context "context"

	billing "github.com/databricks/databricks-sdk-go/service/billing"

	listing "github.com/databricks/databricks-sdk-go/listing"

	mock "github.com/stretchr/testify/mock"
)

// MockBudgetsInterface is an autogenerated mock type for the BudgetsInterface type
type MockBudgetsInterface struct {
	mock.Mock
}

type MockBudgetsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBudgetsInterface) EXPECT() *MockBudgetsInterface_Expecter {
	return &MockBudgetsInterface_Expecter{mock: &_m.Mock}
}

// BudgetWithStatusNameToBudgetIdMap provides a mock function with given fields: ctx
func (_m *MockBudgetsInterface) BudgetWithStatusNameToBudgetIdMap(ctx context.Context) (map[string]string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for BudgetWithStatusNameToBudgetIdMap")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (map[string]string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) map[string]string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBudgetsInterface_BudgetWithStatusNameToBudgetIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BudgetWithStatusNameToBudgetIdMap'
type MockBudgetsInterface_BudgetWithStatusNameToBudgetIdMap_Call struct {
	*mock.Call
}

// BudgetWithStatusNameToBudgetIdMap is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockBudgetsInterface_Expecter) BudgetWithStatusNameToBudgetIdMap(ctx interface{}) *MockBudgetsInterface_BudgetWithStatusNameToBudgetIdMap_Call {
	return &MockBudgetsInterface_BudgetWithStatusNameToBudgetIdMap_Call{Call: _e.mock.On("BudgetWithStatusNameToBudgetIdMap", ctx)}
}

func (_c *MockBudgetsInterface_BudgetWithStatusNameToBudgetIdMap_Call) Run(run func(ctx context.Context)) *MockBudgetsInterface_BudgetWithStatusNameToBudgetIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockBudgetsInterface_BudgetWithStatusNameToBudgetIdMap_Call) Return(_a0 map[string]string, _a1 error) *MockBudgetsInterface_BudgetWithStatusNameToBudgetIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBudgetsInterface_BudgetWithStatusNameToBudgetIdMap_Call) RunAndReturn(run func(context.Context) (map[string]string, error)) *MockBudgetsInterface_BudgetWithStatusNameToBudgetIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockBudgetsInterface) Create(ctx context.Context, request billing.WrappedBudget) (*billing.WrappedBudgetWithStatus, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *billing.WrappedBudgetWithStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, billing.WrappedBudget) (*billing.WrappedBudgetWithStatus, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, billing.WrappedBudget) *billing.WrappedBudgetWithStatus); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billing.WrappedBudgetWithStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, billing.WrappedBudget) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBudgetsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockBudgetsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request billing.WrappedBudget
func (_e *MockBudgetsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockBudgetsInterface_Create_Call {
	return &MockBudgetsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockBudgetsInterface_Create_Call) Run(run func(ctx context.Context, request billing.WrappedBudget)) *MockBudgetsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(billing.WrappedBudget))
	})
	return _c
}

func (_c *MockBudgetsInterface_Create_Call) Return(_a0 *billing.WrappedBudgetWithStatus, _a1 error) *MockBudgetsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBudgetsInterface_Create_Call) RunAndReturn(run func(context.Context, billing.WrappedBudget) (*billing.WrappedBudgetWithStatus, error)) *MockBudgetsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockBudgetsInterface) Delete(ctx context.Context, request billing.DeleteBudgetRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, billing.DeleteBudgetRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBudgetsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockBudgetsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request billing.DeleteBudgetRequest
func (_e *MockBudgetsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockBudgetsInterface_Delete_Call {
	return &MockBudgetsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockBudgetsInterface_Delete_Call) Run(run func(ctx context.Context, request billing.DeleteBudgetRequest)) *MockBudgetsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(billing.DeleteBudgetRequest))
	})
	return _c
}

func (_c *MockBudgetsInterface_Delete_Call) Return(_a0 error) *MockBudgetsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBudgetsInterface_Delete_Call) RunAndReturn(run func(context.Context, billing.DeleteBudgetRequest) error) *MockBudgetsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByBudgetId provides a mock function with given fields: ctx, budgetId
func (_m *MockBudgetsInterface) DeleteByBudgetId(ctx context.Context, budgetId string) error {
	ret := _m.Called(ctx, budgetId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByBudgetId")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, budgetId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBudgetsInterface_DeleteByBudgetId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByBudgetId'
type MockBudgetsInterface_DeleteByBudgetId_Call struct {
	*mock.Call
}

// DeleteByBudgetId is a helper method to define mock.On call
//   - ctx context.Context
//   - budgetId string
func (_e *MockBudgetsInterface_Expecter) DeleteByBudgetId(ctx interface{}, budgetId interface{}) *MockBudgetsInterface_DeleteByBudgetId_Call {
	return &MockBudgetsInterface_DeleteByBudgetId_Call{Call: _e.mock.On("DeleteByBudgetId", ctx, budgetId)}
}

func (_c *MockBudgetsInterface_DeleteByBudgetId_Call) Run(run func(ctx context.Context, budgetId string)) *MockBudgetsInterface_DeleteByBudgetId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockBudgetsInterface_DeleteByBudgetId_Call) Return(_a0 error) *MockBudgetsInterface_DeleteByBudgetId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBudgetsInterface_DeleteByBudgetId_Call) RunAndReturn(run func(context.Context, string) error) *MockBudgetsInterface_DeleteByBudgetId_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockBudgetsInterface) Get(ctx context.Context, request billing.GetBudgetRequest) (*billing.WrappedBudgetWithStatus, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *billing.WrappedBudgetWithStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, billing.GetBudgetRequest) (*billing.WrappedBudgetWithStatus, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, billing.GetBudgetRequest) *billing.WrappedBudgetWithStatus); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billing.WrappedBudgetWithStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, billing.GetBudgetRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBudgetsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockBudgetsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request billing.GetBudgetRequest
func (_e *MockBudgetsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockBudgetsInterface_Get_Call {
	return &MockBudgetsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockBudgetsInterface_Get_Call) Run(run func(ctx context.Context, request billing.GetBudgetRequest)) *MockBudgetsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(billing.GetBudgetRequest))
	})
	return _c
}

func (_c *MockBudgetsInterface_Get_Call) Return(_a0 *billing.WrappedBudgetWithStatus, _a1 error) *MockBudgetsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBudgetsInterface_Get_Call) RunAndReturn(run func(context.Context, billing.GetBudgetRequest) (*billing.WrappedBudgetWithStatus, error)) *MockBudgetsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByBudgetId provides a mock function with given fields: ctx, budgetId
func (_m *MockBudgetsInterface) GetByBudgetId(ctx context.Context, budgetId string) (*billing.WrappedBudgetWithStatus, error) {
	ret := _m.Called(ctx, budgetId)

	if len(ret) == 0 {
		panic("no return value specified for GetByBudgetId")
	}

	var r0 *billing.WrappedBudgetWithStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*billing.WrappedBudgetWithStatus, error)); ok {
		return rf(ctx, budgetId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *billing.WrappedBudgetWithStatus); ok {
		r0 = rf(ctx, budgetId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billing.WrappedBudgetWithStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, budgetId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBudgetsInterface_GetByBudgetId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByBudgetId'
type MockBudgetsInterface_GetByBudgetId_Call struct {
	*mock.Call
}

// GetByBudgetId is a helper method to define mock.On call
//   - ctx context.Context
//   - budgetId string
func (_e *MockBudgetsInterface_Expecter) GetByBudgetId(ctx interface{}, budgetId interface{}) *MockBudgetsInterface_GetByBudgetId_Call {
	return &MockBudgetsInterface_GetByBudgetId_Call{Call: _e.mock.On("GetByBudgetId", ctx, budgetId)}
}

func (_c *MockBudgetsInterface_GetByBudgetId_Call) Run(run func(ctx context.Context, budgetId string)) *MockBudgetsInterface_GetByBudgetId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockBudgetsInterface_GetByBudgetId_Call) Return(_a0 *billing.WrappedBudgetWithStatus, _a1 error) *MockBudgetsInterface_GetByBudgetId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBudgetsInterface_GetByBudgetId_Call) RunAndReturn(run func(context.Context, string) (*billing.WrappedBudgetWithStatus, error)) *MockBudgetsInterface_GetByBudgetId_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *MockBudgetsInterface) GetByName(ctx context.Context, name string) (*billing.BudgetWithStatus, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *billing.BudgetWithStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*billing.BudgetWithStatus, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *billing.BudgetWithStatus); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billing.BudgetWithStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBudgetsInterface_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type MockBudgetsInterface_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockBudgetsInterface_Expecter) GetByName(ctx interface{}, name interface{}) *MockBudgetsInterface_GetByName_Call {
	return &MockBudgetsInterface_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *MockBudgetsInterface_GetByName_Call) Run(run func(ctx context.Context, name string)) *MockBudgetsInterface_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockBudgetsInterface_GetByName_Call) Return(_a0 *billing.BudgetWithStatus, _a1 error) *MockBudgetsInterface_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBudgetsInterface_GetByName_Call) RunAndReturn(run func(context.Context, string) (*billing.BudgetWithStatus, error)) *MockBudgetsInterface_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockBudgetsInterface) Impl() billing.BudgetsService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 billing.BudgetsService
	if rf, ok := ret.Get(0).(func() billing.BudgetsService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(billing.BudgetsService)
		}
	}

	return r0
}

// MockBudgetsInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockBudgetsInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockBudgetsInterface_Expecter) Impl() *MockBudgetsInterface_Impl_Call {
	return &MockBudgetsInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockBudgetsInterface_Impl_Call) Run(run func()) *MockBudgetsInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBudgetsInterface_Impl_Call) Return(_a0 billing.BudgetsService) *MockBudgetsInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBudgetsInterface_Impl_Call) RunAndReturn(run func() billing.BudgetsService) *MockBudgetsInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx
func (_m *MockBudgetsInterface) List(ctx context.Context) listing.Iterator[billing.BudgetWithStatus] {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[billing.BudgetWithStatus]
	if rf, ok := ret.Get(0).(func(context.Context) listing.Iterator[billing.BudgetWithStatus]); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[billing.BudgetWithStatus])
		}
	}

	return r0
}

// MockBudgetsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockBudgetsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockBudgetsInterface_Expecter) List(ctx interface{}) *MockBudgetsInterface_List_Call {
	return &MockBudgetsInterface_List_Call{Call: _e.mock.On("List", ctx)}
}

func (_c *MockBudgetsInterface_List_Call) Run(run func(ctx context.Context)) *MockBudgetsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockBudgetsInterface_List_Call) Return(_a0 listing.Iterator[billing.BudgetWithStatus]) *MockBudgetsInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBudgetsInterface_List_Call) RunAndReturn(run func(context.Context) listing.Iterator[billing.BudgetWithStatus]) *MockBudgetsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx
func (_m *MockBudgetsInterface) ListAll(ctx context.Context) ([]billing.BudgetWithStatus, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []billing.BudgetWithStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]billing.BudgetWithStatus, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []billing.BudgetWithStatus); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]billing.BudgetWithStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBudgetsInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockBudgetsInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockBudgetsInterface_Expecter) ListAll(ctx interface{}) *MockBudgetsInterface_ListAll_Call {
	return &MockBudgetsInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx)}
}

func (_c *MockBudgetsInterface_ListAll_Call) Run(run func(ctx context.Context)) *MockBudgetsInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockBudgetsInterface_ListAll_Call) Return(_a0 []billing.BudgetWithStatus, _a1 error) *MockBudgetsInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBudgetsInterface_ListAll_Call) RunAndReturn(run func(context.Context) ([]billing.BudgetWithStatus, error)) *MockBudgetsInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockBudgetsInterface) Update(ctx context.Context, request billing.WrappedBudget) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, billing.WrappedBudget) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBudgetsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockBudgetsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request billing.WrappedBudget
func (_e *MockBudgetsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockBudgetsInterface_Update_Call {
	return &MockBudgetsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockBudgetsInterface_Update_Call) Run(run func(ctx context.Context, request billing.WrappedBudget)) *MockBudgetsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(billing.WrappedBudget))
	})
	return _c
}

func (_c *MockBudgetsInterface_Update_Call) Return(_a0 error) *MockBudgetsInterface_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBudgetsInterface_Update_Call) RunAndReturn(run func(context.Context, billing.WrappedBudget) error) *MockBudgetsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockBudgetsInterface) WithImpl(impl billing.BudgetsService) billing.BudgetsInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 billing.BudgetsInterface
	if rf, ok := ret.Get(0).(func(billing.BudgetsService) billing.BudgetsInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(billing.BudgetsInterface)
		}
	}

	return r0
}

// MockBudgetsInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockBudgetsInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl billing.BudgetsService
func (_e *MockBudgetsInterface_Expecter) WithImpl(impl interface{}) *MockBudgetsInterface_WithImpl_Call {
	return &MockBudgetsInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockBudgetsInterface_WithImpl_Call) Run(run func(impl billing.BudgetsService)) *MockBudgetsInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(billing.BudgetsService))
	})
	return _c
}

func (_c *MockBudgetsInterface_WithImpl_Call) Return(_a0 billing.BudgetsInterface) *MockBudgetsInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBudgetsInterface_WithImpl_Call) RunAndReturn(run func(billing.BudgetsService) billing.BudgetsInterface) *MockBudgetsInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBudgetsInterface creates a new instance of MockBudgetsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBudgetsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBudgetsInterface {
	mock := &MockBudgetsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
