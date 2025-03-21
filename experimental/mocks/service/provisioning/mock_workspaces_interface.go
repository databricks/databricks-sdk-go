// Code generated by mockery v2.53.2. DO NOT EDIT.

package provisioning

import (
	context "context"

	provisioning "github.com/databricks/databricks-sdk-go/service/provisioning"
	mock "github.com/stretchr/testify/mock"

	retries "github.com/databricks/databricks-sdk-go/retries"

	time "time"
)

// MockWorkspacesInterface is an autogenerated mock type for the WorkspacesInterface type
type MockWorkspacesInterface struct {
	mock.Mock
}

type MockWorkspacesInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWorkspacesInterface) EXPECT() *MockWorkspacesInterface_Expecter {
	return &MockWorkspacesInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, createWorkspaceRequest
func (_m *MockWorkspacesInterface) Create(ctx context.Context, createWorkspaceRequest provisioning.CreateWorkspaceRequest) (*provisioning.WaitGetWorkspaceRunning[provisioning.Workspace], error) {
	ret := _m.Called(ctx, createWorkspaceRequest)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *provisioning.WaitGetWorkspaceRunning[provisioning.Workspace]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, provisioning.CreateWorkspaceRequest) (*provisioning.WaitGetWorkspaceRunning[provisioning.Workspace], error)); ok {
		return rf(ctx, createWorkspaceRequest)
	}
	if rf, ok := ret.Get(0).(func(context.Context, provisioning.CreateWorkspaceRequest) *provisioning.WaitGetWorkspaceRunning[provisioning.Workspace]); ok {
		r0 = rf(ctx, createWorkspaceRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provisioning.WaitGetWorkspaceRunning[provisioning.Workspace])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, provisioning.CreateWorkspaceRequest) error); ok {
		r1 = rf(ctx, createWorkspaceRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkspacesInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockWorkspacesInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - createWorkspaceRequest provisioning.CreateWorkspaceRequest
func (_e *MockWorkspacesInterface_Expecter) Create(ctx interface{}, createWorkspaceRequest interface{}) *MockWorkspacesInterface_Create_Call {
	return &MockWorkspacesInterface_Create_Call{Call: _e.mock.On("Create", ctx, createWorkspaceRequest)}
}

func (_c *MockWorkspacesInterface_Create_Call) Run(run func(ctx context.Context, createWorkspaceRequest provisioning.CreateWorkspaceRequest)) *MockWorkspacesInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(provisioning.CreateWorkspaceRequest))
	})
	return _c
}

func (_c *MockWorkspacesInterface_Create_Call) Return(_a0 *provisioning.WaitGetWorkspaceRunning[provisioning.Workspace], _a1 error) *MockWorkspacesInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkspacesInterface_Create_Call) RunAndReturn(run func(context.Context, provisioning.CreateWorkspaceRequest) (*provisioning.WaitGetWorkspaceRunning[provisioning.Workspace], error)) *MockWorkspacesInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// CreateAndWait provides a mock function with given fields: ctx, createWorkspaceRequest, options
func (_m *MockWorkspacesInterface) CreateAndWait(ctx context.Context, createWorkspaceRequest provisioning.CreateWorkspaceRequest, options ...retries.Option[provisioning.Workspace]) (*provisioning.Workspace, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, createWorkspaceRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateAndWait")
	}

	var r0 *provisioning.Workspace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, provisioning.CreateWorkspaceRequest, ...retries.Option[provisioning.Workspace]) (*provisioning.Workspace, error)); ok {
		return rf(ctx, createWorkspaceRequest, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, provisioning.CreateWorkspaceRequest, ...retries.Option[provisioning.Workspace]) *provisioning.Workspace); ok {
		r0 = rf(ctx, createWorkspaceRequest, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provisioning.Workspace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, provisioning.CreateWorkspaceRequest, ...retries.Option[provisioning.Workspace]) error); ok {
		r1 = rf(ctx, createWorkspaceRequest, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkspacesInterface_CreateAndWait_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAndWait'
type MockWorkspacesInterface_CreateAndWait_Call struct {
	*mock.Call
}

// CreateAndWait is a helper method to define mock.On call
//   - ctx context.Context
//   - createWorkspaceRequest provisioning.CreateWorkspaceRequest
//   - options ...retries.Option[provisioning.Workspace]
func (_e *MockWorkspacesInterface_Expecter) CreateAndWait(ctx interface{}, createWorkspaceRequest interface{}, options ...interface{}) *MockWorkspacesInterface_CreateAndWait_Call {
	return &MockWorkspacesInterface_CreateAndWait_Call{Call: _e.mock.On("CreateAndWait",
		append([]interface{}{ctx, createWorkspaceRequest}, options...)...)}
}

func (_c *MockWorkspacesInterface_CreateAndWait_Call) Run(run func(ctx context.Context, createWorkspaceRequest provisioning.CreateWorkspaceRequest, options ...retries.Option[provisioning.Workspace])) *MockWorkspacesInterface_CreateAndWait_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]retries.Option[provisioning.Workspace], len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(retries.Option[provisioning.Workspace])
			}
		}
		run(args[0].(context.Context), args[1].(provisioning.CreateWorkspaceRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockWorkspacesInterface_CreateAndWait_Call) Return(_a0 *provisioning.Workspace, _a1 error) *MockWorkspacesInterface_CreateAndWait_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkspacesInterface_CreateAndWait_Call) RunAndReturn(run func(context.Context, provisioning.CreateWorkspaceRequest, ...retries.Option[provisioning.Workspace]) (*provisioning.Workspace, error)) *MockWorkspacesInterface_CreateAndWait_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockWorkspacesInterface) Delete(ctx context.Context, request provisioning.DeleteWorkspaceRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, provisioning.DeleteWorkspaceRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWorkspacesInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockWorkspacesInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request provisioning.DeleteWorkspaceRequest
func (_e *MockWorkspacesInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockWorkspacesInterface_Delete_Call {
	return &MockWorkspacesInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockWorkspacesInterface_Delete_Call) Run(run func(ctx context.Context, request provisioning.DeleteWorkspaceRequest)) *MockWorkspacesInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(provisioning.DeleteWorkspaceRequest))
	})
	return _c
}

func (_c *MockWorkspacesInterface_Delete_Call) Return(_a0 error) *MockWorkspacesInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorkspacesInterface_Delete_Call) RunAndReturn(run func(context.Context, provisioning.DeleteWorkspaceRequest) error) *MockWorkspacesInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByWorkspaceId provides a mock function with given fields: ctx, workspaceId
func (_m *MockWorkspacesInterface) DeleteByWorkspaceId(ctx context.Context, workspaceId int64) error {
	ret := _m.Called(ctx, workspaceId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByWorkspaceId")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, workspaceId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWorkspacesInterface_DeleteByWorkspaceId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByWorkspaceId'
type MockWorkspacesInterface_DeleteByWorkspaceId_Call struct {
	*mock.Call
}

// DeleteByWorkspaceId is a helper method to define mock.On call
//   - ctx context.Context
//   - workspaceId int64
func (_e *MockWorkspacesInterface_Expecter) DeleteByWorkspaceId(ctx interface{}, workspaceId interface{}) *MockWorkspacesInterface_DeleteByWorkspaceId_Call {
	return &MockWorkspacesInterface_DeleteByWorkspaceId_Call{Call: _e.mock.On("DeleteByWorkspaceId", ctx, workspaceId)}
}

func (_c *MockWorkspacesInterface_DeleteByWorkspaceId_Call) Run(run func(ctx context.Context, workspaceId int64)) *MockWorkspacesInterface_DeleteByWorkspaceId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockWorkspacesInterface_DeleteByWorkspaceId_Call) Return(_a0 error) *MockWorkspacesInterface_DeleteByWorkspaceId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorkspacesInterface_DeleteByWorkspaceId_Call) RunAndReturn(run func(context.Context, int64) error) *MockWorkspacesInterface_DeleteByWorkspaceId_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockWorkspacesInterface) Get(ctx context.Context, request provisioning.GetWorkspaceRequest) (*provisioning.Workspace, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *provisioning.Workspace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, provisioning.GetWorkspaceRequest) (*provisioning.Workspace, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, provisioning.GetWorkspaceRequest) *provisioning.Workspace); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provisioning.Workspace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, provisioning.GetWorkspaceRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkspacesInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockWorkspacesInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request provisioning.GetWorkspaceRequest
func (_e *MockWorkspacesInterface_Expecter) Get(ctx interface{}, request interface{}) *MockWorkspacesInterface_Get_Call {
	return &MockWorkspacesInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockWorkspacesInterface_Get_Call) Run(run func(ctx context.Context, request provisioning.GetWorkspaceRequest)) *MockWorkspacesInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(provisioning.GetWorkspaceRequest))
	})
	return _c
}

func (_c *MockWorkspacesInterface_Get_Call) Return(_a0 *provisioning.Workspace, _a1 error) *MockWorkspacesInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkspacesInterface_Get_Call) RunAndReturn(run func(context.Context, provisioning.GetWorkspaceRequest) (*provisioning.Workspace, error)) *MockWorkspacesInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByWorkspaceId provides a mock function with given fields: ctx, workspaceId
func (_m *MockWorkspacesInterface) GetByWorkspaceId(ctx context.Context, workspaceId int64) (*provisioning.Workspace, error) {
	ret := _m.Called(ctx, workspaceId)

	if len(ret) == 0 {
		panic("no return value specified for GetByWorkspaceId")
	}

	var r0 *provisioning.Workspace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*provisioning.Workspace, error)); ok {
		return rf(ctx, workspaceId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *provisioning.Workspace); ok {
		r0 = rf(ctx, workspaceId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provisioning.Workspace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, workspaceId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkspacesInterface_GetByWorkspaceId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByWorkspaceId'
type MockWorkspacesInterface_GetByWorkspaceId_Call struct {
	*mock.Call
}

// GetByWorkspaceId is a helper method to define mock.On call
//   - ctx context.Context
//   - workspaceId int64
func (_e *MockWorkspacesInterface_Expecter) GetByWorkspaceId(ctx interface{}, workspaceId interface{}) *MockWorkspacesInterface_GetByWorkspaceId_Call {
	return &MockWorkspacesInterface_GetByWorkspaceId_Call{Call: _e.mock.On("GetByWorkspaceId", ctx, workspaceId)}
}

func (_c *MockWorkspacesInterface_GetByWorkspaceId_Call) Run(run func(ctx context.Context, workspaceId int64)) *MockWorkspacesInterface_GetByWorkspaceId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockWorkspacesInterface_GetByWorkspaceId_Call) Return(_a0 *provisioning.Workspace, _a1 error) *MockWorkspacesInterface_GetByWorkspaceId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkspacesInterface_GetByWorkspaceId_Call) RunAndReturn(run func(context.Context, int64) (*provisioning.Workspace, error)) *MockWorkspacesInterface_GetByWorkspaceId_Call {
	_c.Call.Return(run)
	return _c
}

// GetByWorkspaceName provides a mock function with given fields: ctx, name
func (_m *MockWorkspacesInterface) GetByWorkspaceName(ctx context.Context, name string) (*provisioning.Workspace, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByWorkspaceName")
	}

	var r0 *provisioning.Workspace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*provisioning.Workspace, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *provisioning.Workspace); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provisioning.Workspace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkspacesInterface_GetByWorkspaceName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByWorkspaceName'
type MockWorkspacesInterface_GetByWorkspaceName_Call struct {
	*mock.Call
}

// GetByWorkspaceName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockWorkspacesInterface_Expecter) GetByWorkspaceName(ctx interface{}, name interface{}) *MockWorkspacesInterface_GetByWorkspaceName_Call {
	return &MockWorkspacesInterface_GetByWorkspaceName_Call{Call: _e.mock.On("GetByWorkspaceName", ctx, name)}
}

func (_c *MockWorkspacesInterface_GetByWorkspaceName_Call) Run(run func(ctx context.Context, name string)) *MockWorkspacesInterface_GetByWorkspaceName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockWorkspacesInterface_GetByWorkspaceName_Call) Return(_a0 *provisioning.Workspace, _a1 error) *MockWorkspacesInterface_GetByWorkspaceName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkspacesInterface_GetByWorkspaceName_Call) RunAndReturn(run func(context.Context, string) (*provisioning.Workspace, error)) *MockWorkspacesInterface_GetByWorkspaceName_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx
func (_m *MockWorkspacesInterface) List(ctx context.Context) ([]provisioning.Workspace, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []provisioning.Workspace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]provisioning.Workspace, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []provisioning.Workspace); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]provisioning.Workspace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkspacesInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockWorkspacesInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockWorkspacesInterface_Expecter) List(ctx interface{}) *MockWorkspacesInterface_List_Call {
	return &MockWorkspacesInterface_List_Call{Call: _e.mock.On("List", ctx)}
}

func (_c *MockWorkspacesInterface_List_Call) Run(run func(ctx context.Context)) *MockWorkspacesInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockWorkspacesInterface_List_Call) Return(_a0 []provisioning.Workspace, _a1 error) *MockWorkspacesInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkspacesInterface_List_Call) RunAndReturn(run func(context.Context) ([]provisioning.Workspace, error)) *MockWorkspacesInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, updateWorkspaceRequest
func (_m *MockWorkspacesInterface) Update(ctx context.Context, updateWorkspaceRequest provisioning.UpdateWorkspaceRequest) (*provisioning.WaitGetWorkspaceRunning[struct{}], error) {
	ret := _m.Called(ctx, updateWorkspaceRequest)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *provisioning.WaitGetWorkspaceRunning[struct{}]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, provisioning.UpdateWorkspaceRequest) (*provisioning.WaitGetWorkspaceRunning[struct{}], error)); ok {
		return rf(ctx, updateWorkspaceRequest)
	}
	if rf, ok := ret.Get(0).(func(context.Context, provisioning.UpdateWorkspaceRequest) *provisioning.WaitGetWorkspaceRunning[struct{}]); ok {
		r0 = rf(ctx, updateWorkspaceRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provisioning.WaitGetWorkspaceRunning[struct{}])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, provisioning.UpdateWorkspaceRequest) error); ok {
		r1 = rf(ctx, updateWorkspaceRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkspacesInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockWorkspacesInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - updateWorkspaceRequest provisioning.UpdateWorkspaceRequest
func (_e *MockWorkspacesInterface_Expecter) Update(ctx interface{}, updateWorkspaceRequest interface{}) *MockWorkspacesInterface_Update_Call {
	return &MockWorkspacesInterface_Update_Call{Call: _e.mock.On("Update", ctx, updateWorkspaceRequest)}
}

func (_c *MockWorkspacesInterface_Update_Call) Run(run func(ctx context.Context, updateWorkspaceRequest provisioning.UpdateWorkspaceRequest)) *MockWorkspacesInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(provisioning.UpdateWorkspaceRequest))
	})
	return _c
}

func (_c *MockWorkspacesInterface_Update_Call) Return(_a0 *provisioning.WaitGetWorkspaceRunning[struct{}], _a1 error) *MockWorkspacesInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkspacesInterface_Update_Call) RunAndReturn(run func(context.Context, provisioning.UpdateWorkspaceRequest) (*provisioning.WaitGetWorkspaceRunning[struct{}], error)) *MockWorkspacesInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateAndWait provides a mock function with given fields: ctx, updateWorkspaceRequest, options
func (_m *MockWorkspacesInterface) UpdateAndWait(ctx context.Context, updateWorkspaceRequest provisioning.UpdateWorkspaceRequest, options ...retries.Option[provisioning.Workspace]) (*provisioning.Workspace, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, updateWorkspaceRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAndWait")
	}

	var r0 *provisioning.Workspace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, provisioning.UpdateWorkspaceRequest, ...retries.Option[provisioning.Workspace]) (*provisioning.Workspace, error)); ok {
		return rf(ctx, updateWorkspaceRequest, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, provisioning.UpdateWorkspaceRequest, ...retries.Option[provisioning.Workspace]) *provisioning.Workspace); ok {
		r0 = rf(ctx, updateWorkspaceRequest, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provisioning.Workspace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, provisioning.UpdateWorkspaceRequest, ...retries.Option[provisioning.Workspace]) error); ok {
		r1 = rf(ctx, updateWorkspaceRequest, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkspacesInterface_UpdateAndWait_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAndWait'
type MockWorkspacesInterface_UpdateAndWait_Call struct {
	*mock.Call
}

// UpdateAndWait is a helper method to define mock.On call
//   - ctx context.Context
//   - updateWorkspaceRequest provisioning.UpdateWorkspaceRequest
//   - options ...retries.Option[provisioning.Workspace]
func (_e *MockWorkspacesInterface_Expecter) UpdateAndWait(ctx interface{}, updateWorkspaceRequest interface{}, options ...interface{}) *MockWorkspacesInterface_UpdateAndWait_Call {
	return &MockWorkspacesInterface_UpdateAndWait_Call{Call: _e.mock.On("UpdateAndWait",
		append([]interface{}{ctx, updateWorkspaceRequest}, options...)...)}
}

func (_c *MockWorkspacesInterface_UpdateAndWait_Call) Run(run func(ctx context.Context, updateWorkspaceRequest provisioning.UpdateWorkspaceRequest, options ...retries.Option[provisioning.Workspace])) *MockWorkspacesInterface_UpdateAndWait_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]retries.Option[provisioning.Workspace], len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(retries.Option[provisioning.Workspace])
			}
		}
		run(args[0].(context.Context), args[1].(provisioning.UpdateWorkspaceRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockWorkspacesInterface_UpdateAndWait_Call) Return(_a0 *provisioning.Workspace, _a1 error) *MockWorkspacesInterface_UpdateAndWait_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkspacesInterface_UpdateAndWait_Call) RunAndReturn(run func(context.Context, provisioning.UpdateWorkspaceRequest, ...retries.Option[provisioning.Workspace]) (*provisioning.Workspace, error)) *MockWorkspacesInterface_UpdateAndWait_Call {
	_c.Call.Return(run)
	return _c
}

// WaitGetWorkspaceRunning provides a mock function with given fields: ctx, workspaceId, timeout, callback
func (_m *MockWorkspacesInterface) WaitGetWorkspaceRunning(ctx context.Context, workspaceId int64, timeout time.Duration, callback func(*provisioning.Workspace)) (*provisioning.Workspace, error) {
	ret := _m.Called(ctx, workspaceId, timeout, callback)

	if len(ret) == 0 {
		panic("no return value specified for WaitGetWorkspaceRunning")
	}

	var r0 *provisioning.Workspace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, time.Duration, func(*provisioning.Workspace)) (*provisioning.Workspace, error)); ok {
		return rf(ctx, workspaceId, timeout, callback)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, time.Duration, func(*provisioning.Workspace)) *provisioning.Workspace); ok {
		r0 = rf(ctx, workspaceId, timeout, callback)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provisioning.Workspace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, time.Duration, func(*provisioning.Workspace)) error); ok {
		r1 = rf(ctx, workspaceId, timeout, callback)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkspacesInterface_WaitGetWorkspaceRunning_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WaitGetWorkspaceRunning'
type MockWorkspacesInterface_WaitGetWorkspaceRunning_Call struct {
	*mock.Call
}

// WaitGetWorkspaceRunning is a helper method to define mock.On call
//   - ctx context.Context
//   - workspaceId int64
//   - timeout time.Duration
//   - callback func(*provisioning.Workspace)
func (_e *MockWorkspacesInterface_Expecter) WaitGetWorkspaceRunning(ctx interface{}, workspaceId interface{}, timeout interface{}, callback interface{}) *MockWorkspacesInterface_WaitGetWorkspaceRunning_Call {
	return &MockWorkspacesInterface_WaitGetWorkspaceRunning_Call{Call: _e.mock.On("WaitGetWorkspaceRunning", ctx, workspaceId, timeout, callback)}
}

func (_c *MockWorkspacesInterface_WaitGetWorkspaceRunning_Call) Run(run func(ctx context.Context, workspaceId int64, timeout time.Duration, callback func(*provisioning.Workspace))) *MockWorkspacesInterface_WaitGetWorkspaceRunning_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(time.Duration), args[3].(func(*provisioning.Workspace)))
	})
	return _c
}

func (_c *MockWorkspacesInterface_WaitGetWorkspaceRunning_Call) Return(_a0 *provisioning.Workspace, _a1 error) *MockWorkspacesInterface_WaitGetWorkspaceRunning_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkspacesInterface_WaitGetWorkspaceRunning_Call) RunAndReturn(run func(context.Context, int64, time.Duration, func(*provisioning.Workspace)) (*provisioning.Workspace, error)) *MockWorkspacesInterface_WaitGetWorkspaceRunning_Call {
	_c.Call.Return(run)
	return _c
}

// WorkspaceWorkspaceNameToWorkspaceIdMap provides a mock function with given fields: ctx
func (_m *MockWorkspacesInterface) WorkspaceWorkspaceNameToWorkspaceIdMap(ctx context.Context) (map[string]int64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for WorkspaceWorkspaceNameToWorkspaceIdMap")
	}

	var r0 map[string]int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (map[string]int64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) map[string]int64); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]int64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWorkspacesInterface_WorkspaceWorkspaceNameToWorkspaceIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WorkspaceWorkspaceNameToWorkspaceIdMap'
type MockWorkspacesInterface_WorkspaceWorkspaceNameToWorkspaceIdMap_Call struct {
	*mock.Call
}

// WorkspaceWorkspaceNameToWorkspaceIdMap is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockWorkspacesInterface_Expecter) WorkspaceWorkspaceNameToWorkspaceIdMap(ctx interface{}) *MockWorkspacesInterface_WorkspaceWorkspaceNameToWorkspaceIdMap_Call {
	return &MockWorkspacesInterface_WorkspaceWorkspaceNameToWorkspaceIdMap_Call{Call: _e.mock.On("WorkspaceWorkspaceNameToWorkspaceIdMap", ctx)}
}

func (_c *MockWorkspacesInterface_WorkspaceWorkspaceNameToWorkspaceIdMap_Call) Run(run func(ctx context.Context)) *MockWorkspacesInterface_WorkspaceWorkspaceNameToWorkspaceIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockWorkspacesInterface_WorkspaceWorkspaceNameToWorkspaceIdMap_Call) Return(_a0 map[string]int64, _a1 error) *MockWorkspacesInterface_WorkspaceWorkspaceNameToWorkspaceIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWorkspacesInterface_WorkspaceWorkspaceNameToWorkspaceIdMap_Call) RunAndReturn(run func(context.Context) (map[string]int64, error)) *MockWorkspacesInterface_WorkspaceWorkspaceNameToWorkspaceIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWorkspacesInterface creates a new instance of MockWorkspacesInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWorkspacesInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWorkspacesInterface {
	mock := &MockWorkspacesInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
