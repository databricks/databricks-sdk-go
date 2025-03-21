// Code generated by mockery v2.53.2. DO NOT EDIT.

package settings

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	mock "github.com/stretchr/testify/mock"

	settings "github.com/databricks/databricks-sdk-go/service/settings"
)

// MockIpAccessListsInterface is an autogenerated mock type for the IpAccessListsInterface type
type MockIpAccessListsInterface struct {
	mock.Mock
}

type MockIpAccessListsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIpAccessListsInterface) EXPECT() *MockIpAccessListsInterface_Expecter {
	return &MockIpAccessListsInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockIpAccessListsInterface) Create(ctx context.Context, request settings.CreateIpAccessList) (*settings.CreateIpAccessListResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *settings.CreateIpAccessListResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.CreateIpAccessList) (*settings.CreateIpAccessListResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.CreateIpAccessList) *settings.CreateIpAccessListResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.CreateIpAccessListResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.CreateIpAccessList) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIpAccessListsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockIpAccessListsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.CreateIpAccessList
func (_e *MockIpAccessListsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockIpAccessListsInterface_Create_Call {
	return &MockIpAccessListsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockIpAccessListsInterface_Create_Call) Run(run func(ctx context.Context, request settings.CreateIpAccessList)) *MockIpAccessListsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.CreateIpAccessList))
	})
	return _c
}

func (_c *MockIpAccessListsInterface_Create_Call) Return(_a0 *settings.CreateIpAccessListResponse, _a1 error) *MockIpAccessListsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIpAccessListsInterface_Create_Call) RunAndReturn(run func(context.Context, settings.CreateIpAccessList) (*settings.CreateIpAccessListResponse, error)) *MockIpAccessListsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockIpAccessListsInterface) Delete(ctx context.Context, request settings.DeleteIpAccessListRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.DeleteIpAccessListRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIpAccessListsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockIpAccessListsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.DeleteIpAccessListRequest
func (_e *MockIpAccessListsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockIpAccessListsInterface_Delete_Call {
	return &MockIpAccessListsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockIpAccessListsInterface_Delete_Call) Run(run func(ctx context.Context, request settings.DeleteIpAccessListRequest)) *MockIpAccessListsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.DeleteIpAccessListRequest))
	})
	return _c
}

func (_c *MockIpAccessListsInterface_Delete_Call) Return(_a0 error) *MockIpAccessListsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIpAccessListsInterface_Delete_Call) RunAndReturn(run func(context.Context, settings.DeleteIpAccessListRequest) error) *MockIpAccessListsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByIpAccessListId provides a mock function with given fields: ctx, ipAccessListId
func (_m *MockIpAccessListsInterface) DeleteByIpAccessListId(ctx context.Context, ipAccessListId string) error {
	ret := _m.Called(ctx, ipAccessListId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByIpAccessListId")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, ipAccessListId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIpAccessListsInterface_DeleteByIpAccessListId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByIpAccessListId'
type MockIpAccessListsInterface_DeleteByIpAccessListId_Call struct {
	*mock.Call
}

// DeleteByIpAccessListId is a helper method to define mock.On call
//   - ctx context.Context
//   - ipAccessListId string
func (_e *MockIpAccessListsInterface_Expecter) DeleteByIpAccessListId(ctx interface{}, ipAccessListId interface{}) *MockIpAccessListsInterface_DeleteByIpAccessListId_Call {
	return &MockIpAccessListsInterface_DeleteByIpAccessListId_Call{Call: _e.mock.On("DeleteByIpAccessListId", ctx, ipAccessListId)}
}

func (_c *MockIpAccessListsInterface_DeleteByIpAccessListId_Call) Run(run func(ctx context.Context, ipAccessListId string)) *MockIpAccessListsInterface_DeleteByIpAccessListId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockIpAccessListsInterface_DeleteByIpAccessListId_Call) Return(_a0 error) *MockIpAccessListsInterface_DeleteByIpAccessListId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIpAccessListsInterface_DeleteByIpAccessListId_Call) RunAndReturn(run func(context.Context, string) error) *MockIpAccessListsInterface_DeleteByIpAccessListId_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockIpAccessListsInterface) Get(ctx context.Context, request settings.GetIpAccessListRequest) (*settings.FetchIpAccessListResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *settings.FetchIpAccessListResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetIpAccessListRequest) (*settings.FetchIpAccessListResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetIpAccessListRequest) *settings.FetchIpAccessListResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.FetchIpAccessListResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.GetIpAccessListRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIpAccessListsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockIpAccessListsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.GetIpAccessListRequest
func (_e *MockIpAccessListsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockIpAccessListsInterface_Get_Call {
	return &MockIpAccessListsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockIpAccessListsInterface_Get_Call) Run(run func(ctx context.Context, request settings.GetIpAccessListRequest)) *MockIpAccessListsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.GetIpAccessListRequest))
	})
	return _c
}

func (_c *MockIpAccessListsInterface_Get_Call) Return(_a0 *settings.FetchIpAccessListResponse, _a1 error) *MockIpAccessListsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIpAccessListsInterface_Get_Call) RunAndReturn(run func(context.Context, settings.GetIpAccessListRequest) (*settings.FetchIpAccessListResponse, error)) *MockIpAccessListsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByIpAccessListId provides a mock function with given fields: ctx, ipAccessListId
func (_m *MockIpAccessListsInterface) GetByIpAccessListId(ctx context.Context, ipAccessListId string) (*settings.FetchIpAccessListResponse, error) {
	ret := _m.Called(ctx, ipAccessListId)

	if len(ret) == 0 {
		panic("no return value specified for GetByIpAccessListId")
	}

	var r0 *settings.FetchIpAccessListResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*settings.FetchIpAccessListResponse, error)); ok {
		return rf(ctx, ipAccessListId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *settings.FetchIpAccessListResponse); ok {
		r0 = rf(ctx, ipAccessListId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.FetchIpAccessListResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ipAccessListId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIpAccessListsInterface_GetByIpAccessListId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByIpAccessListId'
type MockIpAccessListsInterface_GetByIpAccessListId_Call struct {
	*mock.Call
}

// GetByIpAccessListId is a helper method to define mock.On call
//   - ctx context.Context
//   - ipAccessListId string
func (_e *MockIpAccessListsInterface_Expecter) GetByIpAccessListId(ctx interface{}, ipAccessListId interface{}) *MockIpAccessListsInterface_GetByIpAccessListId_Call {
	return &MockIpAccessListsInterface_GetByIpAccessListId_Call{Call: _e.mock.On("GetByIpAccessListId", ctx, ipAccessListId)}
}

func (_c *MockIpAccessListsInterface_GetByIpAccessListId_Call) Run(run func(ctx context.Context, ipAccessListId string)) *MockIpAccessListsInterface_GetByIpAccessListId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockIpAccessListsInterface_GetByIpAccessListId_Call) Return(_a0 *settings.FetchIpAccessListResponse, _a1 error) *MockIpAccessListsInterface_GetByIpAccessListId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIpAccessListsInterface_GetByIpAccessListId_Call) RunAndReturn(run func(context.Context, string) (*settings.FetchIpAccessListResponse, error)) *MockIpAccessListsInterface_GetByIpAccessListId_Call {
	_c.Call.Return(run)
	return _c
}

// GetByLabel provides a mock function with given fields: ctx, name
func (_m *MockIpAccessListsInterface) GetByLabel(ctx context.Context, name string) (*settings.IpAccessListInfo, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByLabel")
	}

	var r0 *settings.IpAccessListInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*settings.IpAccessListInfo, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *settings.IpAccessListInfo); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.IpAccessListInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIpAccessListsInterface_GetByLabel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByLabel'
type MockIpAccessListsInterface_GetByLabel_Call struct {
	*mock.Call
}

// GetByLabel is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockIpAccessListsInterface_Expecter) GetByLabel(ctx interface{}, name interface{}) *MockIpAccessListsInterface_GetByLabel_Call {
	return &MockIpAccessListsInterface_GetByLabel_Call{Call: _e.mock.On("GetByLabel", ctx, name)}
}

func (_c *MockIpAccessListsInterface_GetByLabel_Call) Run(run func(ctx context.Context, name string)) *MockIpAccessListsInterface_GetByLabel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockIpAccessListsInterface_GetByLabel_Call) Return(_a0 *settings.IpAccessListInfo, _a1 error) *MockIpAccessListsInterface_GetByLabel_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIpAccessListsInterface_GetByLabel_Call) RunAndReturn(run func(context.Context, string) (*settings.IpAccessListInfo, error)) *MockIpAccessListsInterface_GetByLabel_Call {
	_c.Call.Return(run)
	return _c
}

// IpAccessListInfoLabelToListIdMap provides a mock function with given fields: ctx
func (_m *MockIpAccessListsInterface) IpAccessListInfoLabelToListIdMap(ctx context.Context) (map[string]string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for IpAccessListInfoLabelToListIdMap")
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

// MockIpAccessListsInterface_IpAccessListInfoLabelToListIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IpAccessListInfoLabelToListIdMap'
type MockIpAccessListsInterface_IpAccessListInfoLabelToListIdMap_Call struct {
	*mock.Call
}

// IpAccessListInfoLabelToListIdMap is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockIpAccessListsInterface_Expecter) IpAccessListInfoLabelToListIdMap(ctx interface{}) *MockIpAccessListsInterface_IpAccessListInfoLabelToListIdMap_Call {
	return &MockIpAccessListsInterface_IpAccessListInfoLabelToListIdMap_Call{Call: _e.mock.On("IpAccessListInfoLabelToListIdMap", ctx)}
}

func (_c *MockIpAccessListsInterface_IpAccessListInfoLabelToListIdMap_Call) Run(run func(ctx context.Context)) *MockIpAccessListsInterface_IpAccessListInfoLabelToListIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockIpAccessListsInterface_IpAccessListInfoLabelToListIdMap_Call) Return(_a0 map[string]string, _a1 error) *MockIpAccessListsInterface_IpAccessListInfoLabelToListIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIpAccessListsInterface_IpAccessListInfoLabelToListIdMap_Call) RunAndReturn(run func(context.Context) (map[string]string, error)) *MockIpAccessListsInterface_IpAccessListInfoLabelToListIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx
func (_m *MockIpAccessListsInterface) List(ctx context.Context) listing.Iterator[settings.IpAccessListInfo] {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[settings.IpAccessListInfo]
	if rf, ok := ret.Get(0).(func(context.Context) listing.Iterator[settings.IpAccessListInfo]); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[settings.IpAccessListInfo])
		}
	}

	return r0
}

// MockIpAccessListsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockIpAccessListsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockIpAccessListsInterface_Expecter) List(ctx interface{}) *MockIpAccessListsInterface_List_Call {
	return &MockIpAccessListsInterface_List_Call{Call: _e.mock.On("List", ctx)}
}

func (_c *MockIpAccessListsInterface_List_Call) Run(run func(ctx context.Context)) *MockIpAccessListsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockIpAccessListsInterface_List_Call) Return(_a0 listing.Iterator[settings.IpAccessListInfo]) *MockIpAccessListsInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIpAccessListsInterface_List_Call) RunAndReturn(run func(context.Context) listing.Iterator[settings.IpAccessListInfo]) *MockIpAccessListsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx
func (_m *MockIpAccessListsInterface) ListAll(ctx context.Context) ([]settings.IpAccessListInfo, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []settings.IpAccessListInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]settings.IpAccessListInfo, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []settings.IpAccessListInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]settings.IpAccessListInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIpAccessListsInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockIpAccessListsInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockIpAccessListsInterface_Expecter) ListAll(ctx interface{}) *MockIpAccessListsInterface_ListAll_Call {
	return &MockIpAccessListsInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx)}
}

func (_c *MockIpAccessListsInterface_ListAll_Call) Run(run func(ctx context.Context)) *MockIpAccessListsInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockIpAccessListsInterface_ListAll_Call) Return(_a0 []settings.IpAccessListInfo, _a1 error) *MockIpAccessListsInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIpAccessListsInterface_ListAll_Call) RunAndReturn(run func(context.Context) ([]settings.IpAccessListInfo, error)) *MockIpAccessListsInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// Replace provides a mock function with given fields: ctx, request
func (_m *MockIpAccessListsInterface) Replace(ctx context.Context, request settings.ReplaceIpAccessList) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Replace")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.ReplaceIpAccessList) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIpAccessListsInterface_Replace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Replace'
type MockIpAccessListsInterface_Replace_Call struct {
	*mock.Call
}

// Replace is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.ReplaceIpAccessList
func (_e *MockIpAccessListsInterface_Expecter) Replace(ctx interface{}, request interface{}) *MockIpAccessListsInterface_Replace_Call {
	return &MockIpAccessListsInterface_Replace_Call{Call: _e.mock.On("Replace", ctx, request)}
}

func (_c *MockIpAccessListsInterface_Replace_Call) Run(run func(ctx context.Context, request settings.ReplaceIpAccessList)) *MockIpAccessListsInterface_Replace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.ReplaceIpAccessList))
	})
	return _c
}

func (_c *MockIpAccessListsInterface_Replace_Call) Return(_a0 error) *MockIpAccessListsInterface_Replace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIpAccessListsInterface_Replace_Call) RunAndReturn(run func(context.Context, settings.ReplaceIpAccessList) error) *MockIpAccessListsInterface_Replace_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockIpAccessListsInterface) Update(ctx context.Context, request settings.UpdateIpAccessList) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateIpAccessList) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIpAccessListsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockIpAccessListsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.UpdateIpAccessList
func (_e *MockIpAccessListsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockIpAccessListsInterface_Update_Call {
	return &MockIpAccessListsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockIpAccessListsInterface_Update_Call) Run(run func(ctx context.Context, request settings.UpdateIpAccessList)) *MockIpAccessListsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.UpdateIpAccessList))
	})
	return _c
}

func (_c *MockIpAccessListsInterface_Update_Call) Return(_a0 error) *MockIpAccessListsInterface_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIpAccessListsInterface_Update_Call) RunAndReturn(run func(context.Context, settings.UpdateIpAccessList) error) *MockIpAccessListsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIpAccessListsInterface creates a new instance of MockIpAccessListsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIpAccessListsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIpAccessListsInterface {
	mock := &MockIpAccessListsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
