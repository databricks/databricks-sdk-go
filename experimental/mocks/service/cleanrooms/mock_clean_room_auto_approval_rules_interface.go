// Code generated by mockery v2.53.2. DO NOT EDIT.

package cleanrooms

import (
	context "context"

	cleanrooms "github.com/databricks/databricks-sdk-go/service/cleanrooms"

	listing "github.com/databricks/databricks-sdk-go/listing"

	mock "github.com/stretchr/testify/mock"
)

// MockCleanRoomAutoApprovalRulesInterface is an autogenerated mock type for the CleanRoomAutoApprovalRulesInterface type
type MockCleanRoomAutoApprovalRulesInterface struct {
	mock.Mock
}

type MockCleanRoomAutoApprovalRulesInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCleanRoomAutoApprovalRulesInterface) EXPECT() *MockCleanRoomAutoApprovalRulesInterface_Expecter {
	return &MockCleanRoomAutoApprovalRulesInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockCleanRoomAutoApprovalRulesInterface) Create(ctx context.Context, request cleanrooms.CreateCleanRoomAutoApprovalRuleRequest) (*cleanrooms.CleanRoomAutoApprovalRule, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *cleanrooms.CleanRoomAutoApprovalRule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cleanrooms.CreateCleanRoomAutoApprovalRuleRequest) (*cleanrooms.CleanRoomAutoApprovalRule, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cleanrooms.CreateCleanRoomAutoApprovalRuleRequest) *cleanrooms.CleanRoomAutoApprovalRule); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cleanrooms.CleanRoomAutoApprovalRule)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, cleanrooms.CreateCleanRoomAutoApprovalRuleRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCleanRoomAutoApprovalRulesInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockCleanRoomAutoApprovalRulesInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request cleanrooms.CreateCleanRoomAutoApprovalRuleRequest
func (_e *MockCleanRoomAutoApprovalRulesInterface_Expecter) Create(ctx interface{}, request interface{}) *MockCleanRoomAutoApprovalRulesInterface_Create_Call {
	return &MockCleanRoomAutoApprovalRulesInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockCleanRoomAutoApprovalRulesInterface_Create_Call) Run(run func(ctx context.Context, request cleanrooms.CreateCleanRoomAutoApprovalRuleRequest)) *MockCleanRoomAutoApprovalRulesInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(cleanrooms.CreateCleanRoomAutoApprovalRuleRequest))
	})
	return _c
}

func (_c *MockCleanRoomAutoApprovalRulesInterface_Create_Call) Return(_a0 *cleanrooms.CleanRoomAutoApprovalRule, _a1 error) *MockCleanRoomAutoApprovalRulesInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCleanRoomAutoApprovalRulesInterface_Create_Call) RunAndReturn(run func(context.Context, cleanrooms.CreateCleanRoomAutoApprovalRuleRequest) (*cleanrooms.CleanRoomAutoApprovalRule, error)) *MockCleanRoomAutoApprovalRulesInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockCleanRoomAutoApprovalRulesInterface) Delete(ctx context.Context, request cleanrooms.DeleteCleanRoomAutoApprovalRuleRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, cleanrooms.DeleteCleanRoomAutoApprovalRuleRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCleanRoomAutoApprovalRulesInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockCleanRoomAutoApprovalRulesInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request cleanrooms.DeleteCleanRoomAutoApprovalRuleRequest
func (_e *MockCleanRoomAutoApprovalRulesInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockCleanRoomAutoApprovalRulesInterface_Delete_Call {
	return &MockCleanRoomAutoApprovalRulesInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockCleanRoomAutoApprovalRulesInterface_Delete_Call) Run(run func(ctx context.Context, request cleanrooms.DeleteCleanRoomAutoApprovalRuleRequest)) *MockCleanRoomAutoApprovalRulesInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(cleanrooms.DeleteCleanRoomAutoApprovalRuleRequest))
	})
	return _c
}

func (_c *MockCleanRoomAutoApprovalRulesInterface_Delete_Call) Return(_a0 error) *MockCleanRoomAutoApprovalRulesInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCleanRoomAutoApprovalRulesInterface_Delete_Call) RunAndReturn(run func(context.Context, cleanrooms.DeleteCleanRoomAutoApprovalRuleRequest) error) *MockCleanRoomAutoApprovalRulesInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockCleanRoomAutoApprovalRulesInterface) Get(ctx context.Context, request cleanrooms.GetCleanRoomAutoApprovalRuleRequest) (*cleanrooms.CleanRoomAutoApprovalRule, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *cleanrooms.CleanRoomAutoApprovalRule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cleanrooms.GetCleanRoomAutoApprovalRuleRequest) (*cleanrooms.CleanRoomAutoApprovalRule, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cleanrooms.GetCleanRoomAutoApprovalRuleRequest) *cleanrooms.CleanRoomAutoApprovalRule); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cleanrooms.CleanRoomAutoApprovalRule)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, cleanrooms.GetCleanRoomAutoApprovalRuleRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCleanRoomAutoApprovalRulesInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockCleanRoomAutoApprovalRulesInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request cleanrooms.GetCleanRoomAutoApprovalRuleRequest
func (_e *MockCleanRoomAutoApprovalRulesInterface_Expecter) Get(ctx interface{}, request interface{}) *MockCleanRoomAutoApprovalRulesInterface_Get_Call {
	return &MockCleanRoomAutoApprovalRulesInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockCleanRoomAutoApprovalRulesInterface_Get_Call) Run(run func(ctx context.Context, request cleanrooms.GetCleanRoomAutoApprovalRuleRequest)) *MockCleanRoomAutoApprovalRulesInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(cleanrooms.GetCleanRoomAutoApprovalRuleRequest))
	})
	return _c
}

func (_c *MockCleanRoomAutoApprovalRulesInterface_Get_Call) Return(_a0 *cleanrooms.CleanRoomAutoApprovalRule, _a1 error) *MockCleanRoomAutoApprovalRulesInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCleanRoomAutoApprovalRulesInterface_Get_Call) RunAndReturn(run func(context.Context, cleanrooms.GetCleanRoomAutoApprovalRuleRequest) (*cleanrooms.CleanRoomAutoApprovalRule, error)) *MockCleanRoomAutoApprovalRulesInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockCleanRoomAutoApprovalRulesInterface) List(ctx context.Context, request cleanrooms.ListCleanRoomAutoApprovalRulesRequest) listing.Iterator[cleanrooms.CleanRoomAutoApprovalRule] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[cleanrooms.CleanRoomAutoApprovalRule]
	if rf, ok := ret.Get(0).(func(context.Context, cleanrooms.ListCleanRoomAutoApprovalRulesRequest) listing.Iterator[cleanrooms.CleanRoomAutoApprovalRule]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[cleanrooms.CleanRoomAutoApprovalRule])
		}
	}

	return r0
}

// MockCleanRoomAutoApprovalRulesInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockCleanRoomAutoApprovalRulesInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request cleanrooms.ListCleanRoomAutoApprovalRulesRequest
func (_e *MockCleanRoomAutoApprovalRulesInterface_Expecter) List(ctx interface{}, request interface{}) *MockCleanRoomAutoApprovalRulesInterface_List_Call {
	return &MockCleanRoomAutoApprovalRulesInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockCleanRoomAutoApprovalRulesInterface_List_Call) Run(run func(ctx context.Context, request cleanrooms.ListCleanRoomAutoApprovalRulesRequest)) *MockCleanRoomAutoApprovalRulesInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(cleanrooms.ListCleanRoomAutoApprovalRulesRequest))
	})
	return _c
}

func (_c *MockCleanRoomAutoApprovalRulesInterface_List_Call) Return(_a0 listing.Iterator[cleanrooms.CleanRoomAutoApprovalRule]) *MockCleanRoomAutoApprovalRulesInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCleanRoomAutoApprovalRulesInterface_List_Call) RunAndReturn(run func(context.Context, cleanrooms.ListCleanRoomAutoApprovalRulesRequest) listing.Iterator[cleanrooms.CleanRoomAutoApprovalRule]) *MockCleanRoomAutoApprovalRulesInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockCleanRoomAutoApprovalRulesInterface) ListAll(ctx context.Context, request cleanrooms.ListCleanRoomAutoApprovalRulesRequest) ([]cleanrooms.CleanRoomAutoApprovalRule, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []cleanrooms.CleanRoomAutoApprovalRule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cleanrooms.ListCleanRoomAutoApprovalRulesRequest) ([]cleanrooms.CleanRoomAutoApprovalRule, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cleanrooms.ListCleanRoomAutoApprovalRulesRequest) []cleanrooms.CleanRoomAutoApprovalRule); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cleanrooms.CleanRoomAutoApprovalRule)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, cleanrooms.ListCleanRoomAutoApprovalRulesRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCleanRoomAutoApprovalRulesInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockCleanRoomAutoApprovalRulesInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request cleanrooms.ListCleanRoomAutoApprovalRulesRequest
func (_e *MockCleanRoomAutoApprovalRulesInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockCleanRoomAutoApprovalRulesInterface_ListAll_Call {
	return &MockCleanRoomAutoApprovalRulesInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockCleanRoomAutoApprovalRulesInterface_ListAll_Call) Run(run func(ctx context.Context, request cleanrooms.ListCleanRoomAutoApprovalRulesRequest)) *MockCleanRoomAutoApprovalRulesInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(cleanrooms.ListCleanRoomAutoApprovalRulesRequest))
	})
	return _c
}

func (_c *MockCleanRoomAutoApprovalRulesInterface_ListAll_Call) Return(_a0 []cleanrooms.CleanRoomAutoApprovalRule, _a1 error) *MockCleanRoomAutoApprovalRulesInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCleanRoomAutoApprovalRulesInterface_ListAll_Call) RunAndReturn(run func(context.Context, cleanrooms.ListCleanRoomAutoApprovalRulesRequest) ([]cleanrooms.CleanRoomAutoApprovalRule, error)) *MockCleanRoomAutoApprovalRulesInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockCleanRoomAutoApprovalRulesInterface) Update(ctx context.Context, request cleanrooms.UpdateCleanRoomAutoApprovalRuleRequest) (*cleanrooms.CleanRoomAutoApprovalRule, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *cleanrooms.CleanRoomAutoApprovalRule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cleanrooms.UpdateCleanRoomAutoApprovalRuleRequest) (*cleanrooms.CleanRoomAutoApprovalRule, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cleanrooms.UpdateCleanRoomAutoApprovalRuleRequest) *cleanrooms.CleanRoomAutoApprovalRule); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cleanrooms.CleanRoomAutoApprovalRule)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, cleanrooms.UpdateCleanRoomAutoApprovalRuleRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCleanRoomAutoApprovalRulesInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockCleanRoomAutoApprovalRulesInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request cleanrooms.UpdateCleanRoomAutoApprovalRuleRequest
func (_e *MockCleanRoomAutoApprovalRulesInterface_Expecter) Update(ctx interface{}, request interface{}) *MockCleanRoomAutoApprovalRulesInterface_Update_Call {
	return &MockCleanRoomAutoApprovalRulesInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockCleanRoomAutoApprovalRulesInterface_Update_Call) Run(run func(ctx context.Context, request cleanrooms.UpdateCleanRoomAutoApprovalRuleRequest)) *MockCleanRoomAutoApprovalRulesInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(cleanrooms.UpdateCleanRoomAutoApprovalRuleRequest))
	})
	return _c
}

func (_c *MockCleanRoomAutoApprovalRulesInterface_Update_Call) Return(_a0 *cleanrooms.CleanRoomAutoApprovalRule, _a1 error) *MockCleanRoomAutoApprovalRulesInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCleanRoomAutoApprovalRulesInterface_Update_Call) RunAndReturn(run func(context.Context, cleanrooms.UpdateCleanRoomAutoApprovalRuleRequest) (*cleanrooms.CleanRoomAutoApprovalRule, error)) *MockCleanRoomAutoApprovalRulesInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCleanRoomAutoApprovalRulesInterface creates a new instance of MockCleanRoomAutoApprovalRulesInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCleanRoomAutoApprovalRulesInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCleanRoomAutoApprovalRulesInterface {
	mock := &MockCleanRoomAutoApprovalRulesInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
