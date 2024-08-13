// Code generated by mockery v2.43.0. DO NOT EDIT.

package settings

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	mock "github.com/stretchr/testify/mock"

	settings "github.com/databricks/databricks-sdk-go/service/settings"
)

// MockAccountIpAccessListsInterface is an autogenerated mock type for the AccountIpAccessListsInterface type
type MockAccountIpAccessListsInterface struct {
	mock.Mock
}

type MockAccountIpAccessListsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAccountIpAccessListsInterface) EXPECT() *MockAccountIpAccessListsInterface_Expecter {
	return &MockAccountIpAccessListsInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockAccountIpAccessListsInterface) Create(ctx context.Context, request settings.CreateIpAccessList) (*settings.CreateIpAccessListResponse, error) {
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

// MockAccountIpAccessListsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockAccountIpAccessListsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.CreateIpAccessList
func (_e *MockAccountIpAccessListsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockAccountIpAccessListsInterface_Create_Call {
	return &MockAccountIpAccessListsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockAccountIpAccessListsInterface_Create_Call) Run(run func(ctx context.Context, request settings.CreateIpAccessList)) *MockAccountIpAccessListsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.CreateIpAccessList))
	})
	return _c
}

func (_c *MockAccountIpAccessListsInterface_Create_Call) Return(_a0 *settings.CreateIpAccessListResponse, _a1 error) *MockAccountIpAccessListsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountIpAccessListsInterface_Create_Call) RunAndReturn(run func(context.Context, settings.CreateIpAccessList) (*settings.CreateIpAccessListResponse, error)) *MockAccountIpAccessListsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockAccountIpAccessListsInterface) Delete(ctx context.Context, request settings.DeleteAccountIpAccessListRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.DeleteAccountIpAccessListRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAccountIpAccessListsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockAccountIpAccessListsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.DeleteAccountIpAccessListRequest
func (_e *MockAccountIpAccessListsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockAccountIpAccessListsInterface_Delete_Call {
	return &MockAccountIpAccessListsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockAccountIpAccessListsInterface_Delete_Call) Run(run func(ctx context.Context, request settings.DeleteAccountIpAccessListRequest)) *MockAccountIpAccessListsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.DeleteAccountIpAccessListRequest))
	})
	return _c
}

func (_c *MockAccountIpAccessListsInterface_Delete_Call) Return(_a0 error) *MockAccountIpAccessListsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountIpAccessListsInterface_Delete_Call) RunAndReturn(run func(context.Context, settings.DeleteAccountIpAccessListRequest) error) *MockAccountIpAccessListsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByIpAccessListId provides a mock function with given fields: ctx, ipAccessListId
func (_m *MockAccountIpAccessListsInterface) DeleteByIpAccessListId(ctx context.Context, ipAccessListId string) error {
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

// MockAccountIpAccessListsInterface_DeleteByIpAccessListId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByIpAccessListId'
type MockAccountIpAccessListsInterface_DeleteByIpAccessListId_Call struct {
	*mock.Call
}

// DeleteByIpAccessListId is a helper method to define mock.On call
//   - ctx context.Context
//   - ipAccessListId string
func (_e *MockAccountIpAccessListsInterface_Expecter) DeleteByIpAccessListId(ctx interface{}, ipAccessListId interface{}) *MockAccountIpAccessListsInterface_DeleteByIpAccessListId_Call {
	return &MockAccountIpAccessListsInterface_DeleteByIpAccessListId_Call{Call: _e.mock.On("DeleteByIpAccessListId", ctx, ipAccessListId)}
}

func (_c *MockAccountIpAccessListsInterface_DeleteByIpAccessListId_Call) Run(run func(ctx context.Context, ipAccessListId string)) *MockAccountIpAccessListsInterface_DeleteByIpAccessListId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAccountIpAccessListsInterface_DeleteByIpAccessListId_Call) Return(_a0 error) *MockAccountIpAccessListsInterface_DeleteByIpAccessListId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountIpAccessListsInterface_DeleteByIpAccessListId_Call) RunAndReturn(run func(context.Context, string) error) *MockAccountIpAccessListsInterface_DeleteByIpAccessListId_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockAccountIpAccessListsInterface) Get(ctx context.Context, request settings.GetAccountIpAccessListRequest) (*settings.GetIpAccessListResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *settings.GetIpAccessListResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetAccountIpAccessListRequest) (*settings.GetIpAccessListResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetAccountIpAccessListRequest) *settings.GetIpAccessListResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.GetIpAccessListResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.GetAccountIpAccessListRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountIpAccessListsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockAccountIpAccessListsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.GetAccountIpAccessListRequest
func (_e *MockAccountIpAccessListsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockAccountIpAccessListsInterface_Get_Call {
	return &MockAccountIpAccessListsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockAccountIpAccessListsInterface_Get_Call) Run(run func(ctx context.Context, request settings.GetAccountIpAccessListRequest)) *MockAccountIpAccessListsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.GetAccountIpAccessListRequest))
	})
	return _c
}

func (_c *MockAccountIpAccessListsInterface_Get_Call) Return(_a0 *settings.GetIpAccessListResponse, _a1 error) *MockAccountIpAccessListsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountIpAccessListsInterface_Get_Call) RunAndReturn(run func(context.Context, settings.GetAccountIpAccessListRequest) (*settings.GetIpAccessListResponse, error)) *MockAccountIpAccessListsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByIpAccessListId provides a mock function with given fields: ctx, ipAccessListId
func (_m *MockAccountIpAccessListsInterface) GetByIpAccessListId(ctx context.Context, ipAccessListId string) (*settings.GetIpAccessListResponse, error) {
	ret := _m.Called(ctx, ipAccessListId)

	if len(ret) == 0 {
		panic("no return value specified for GetByIpAccessListId")
	}

	var r0 *settings.GetIpAccessListResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*settings.GetIpAccessListResponse, error)); ok {
		return rf(ctx, ipAccessListId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *settings.GetIpAccessListResponse); ok {
		r0 = rf(ctx, ipAccessListId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.GetIpAccessListResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ipAccessListId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountIpAccessListsInterface_GetByIpAccessListId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByIpAccessListId'
type MockAccountIpAccessListsInterface_GetByIpAccessListId_Call struct {
	*mock.Call
}

// GetByIpAccessListId is a helper method to define mock.On call
//   - ctx context.Context
//   - ipAccessListId string
func (_e *MockAccountIpAccessListsInterface_Expecter) GetByIpAccessListId(ctx interface{}, ipAccessListId interface{}) *MockAccountIpAccessListsInterface_GetByIpAccessListId_Call {
	return &MockAccountIpAccessListsInterface_GetByIpAccessListId_Call{Call: _e.mock.On("GetByIpAccessListId", ctx, ipAccessListId)}
}

func (_c *MockAccountIpAccessListsInterface_GetByIpAccessListId_Call) Run(run func(ctx context.Context, ipAccessListId string)) *MockAccountIpAccessListsInterface_GetByIpAccessListId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAccountIpAccessListsInterface_GetByIpAccessListId_Call) Return(_a0 *settings.GetIpAccessListResponse, _a1 error) *MockAccountIpAccessListsInterface_GetByIpAccessListId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountIpAccessListsInterface_GetByIpAccessListId_Call) RunAndReturn(run func(context.Context, string) (*settings.GetIpAccessListResponse, error)) *MockAccountIpAccessListsInterface_GetByIpAccessListId_Call {
	_c.Call.Return(run)
	return _c
}

// GetByLabel provides a mock function with given fields: ctx, name
func (_m *MockAccountIpAccessListsInterface) GetByLabel(ctx context.Context, name string) (*settings.IpAccessListInfo, error) {
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

// MockAccountIpAccessListsInterface_GetByLabel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByLabel'
type MockAccountIpAccessListsInterface_GetByLabel_Call struct {
	*mock.Call
}

// GetByLabel is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockAccountIpAccessListsInterface_Expecter) GetByLabel(ctx interface{}, name interface{}) *MockAccountIpAccessListsInterface_GetByLabel_Call {
	return &MockAccountIpAccessListsInterface_GetByLabel_Call{Call: _e.mock.On("GetByLabel", ctx, name)}
}

func (_c *MockAccountIpAccessListsInterface_GetByLabel_Call) Run(run func(ctx context.Context, name string)) *MockAccountIpAccessListsInterface_GetByLabel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAccountIpAccessListsInterface_GetByLabel_Call) Return(_a0 *settings.IpAccessListInfo, _a1 error) *MockAccountIpAccessListsInterface_GetByLabel_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountIpAccessListsInterface_GetByLabel_Call) RunAndReturn(run func(context.Context, string) (*settings.IpAccessListInfo, error)) *MockAccountIpAccessListsInterface_GetByLabel_Call {
	_c.Call.Return(run)
	return _c
}

// IpAccessListInfoLabelToListIdMap provides a mock function with given fields: ctx
func (_m *MockAccountIpAccessListsInterface) IpAccessListInfoLabelToListIdMap(ctx context.Context) (map[string]string, error) {
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

// MockAccountIpAccessListsInterface_IpAccessListInfoLabelToListIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IpAccessListInfoLabelToListIdMap'
type MockAccountIpAccessListsInterface_IpAccessListInfoLabelToListIdMap_Call struct {
	*mock.Call
}

// IpAccessListInfoLabelToListIdMap is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockAccountIpAccessListsInterface_Expecter) IpAccessListInfoLabelToListIdMap(ctx interface{}) *MockAccountIpAccessListsInterface_IpAccessListInfoLabelToListIdMap_Call {
	return &MockAccountIpAccessListsInterface_IpAccessListInfoLabelToListIdMap_Call{Call: _e.mock.On("IpAccessListInfoLabelToListIdMap", ctx)}
}

func (_c *MockAccountIpAccessListsInterface_IpAccessListInfoLabelToListIdMap_Call) Run(run func(ctx context.Context)) *MockAccountIpAccessListsInterface_IpAccessListInfoLabelToListIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockAccountIpAccessListsInterface_IpAccessListInfoLabelToListIdMap_Call) Return(_a0 map[string]string, _a1 error) *MockAccountIpAccessListsInterface_IpAccessListInfoLabelToListIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountIpAccessListsInterface_IpAccessListInfoLabelToListIdMap_Call) RunAndReturn(run func(context.Context) (map[string]string, error)) *MockAccountIpAccessListsInterface_IpAccessListInfoLabelToListIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx
func (_m *MockAccountIpAccessListsInterface) List(ctx context.Context) listing.Iterator[settings.IpAccessListInfo] {
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

// MockAccountIpAccessListsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockAccountIpAccessListsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockAccountIpAccessListsInterface_Expecter) List(ctx interface{}) *MockAccountIpAccessListsInterface_List_Call {
	return &MockAccountIpAccessListsInterface_List_Call{Call: _e.mock.On("List", ctx)}
}

func (_c *MockAccountIpAccessListsInterface_List_Call) Run(run func(ctx context.Context)) *MockAccountIpAccessListsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockAccountIpAccessListsInterface_List_Call) Return(_a0 listing.Iterator[settings.IpAccessListInfo]) *MockAccountIpAccessListsInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountIpAccessListsInterface_List_Call) RunAndReturn(run func(context.Context) listing.Iterator[settings.IpAccessListInfo]) *MockAccountIpAccessListsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx
func (_m *MockAccountIpAccessListsInterface) ListAll(ctx context.Context) ([]settings.IpAccessListInfo, error) {
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

// MockAccountIpAccessListsInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockAccountIpAccessListsInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockAccountIpAccessListsInterface_Expecter) ListAll(ctx interface{}) *MockAccountIpAccessListsInterface_ListAll_Call {
	return &MockAccountIpAccessListsInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx)}
}

func (_c *MockAccountIpAccessListsInterface_ListAll_Call) Run(run func(ctx context.Context)) *MockAccountIpAccessListsInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockAccountIpAccessListsInterface_ListAll_Call) Return(_a0 []settings.IpAccessListInfo, _a1 error) *MockAccountIpAccessListsInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountIpAccessListsInterface_ListAll_Call) RunAndReturn(run func(context.Context) ([]settings.IpAccessListInfo, error)) *MockAccountIpAccessListsInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// Replace provides a mock function with given fields: ctx, request
func (_m *MockAccountIpAccessListsInterface) Replace(ctx context.Context, request settings.ReplaceIpAccessList) error {
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

// MockAccountIpAccessListsInterface_Replace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Replace'
type MockAccountIpAccessListsInterface_Replace_Call struct {
	*mock.Call
}

// Replace is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.ReplaceIpAccessList
func (_e *MockAccountIpAccessListsInterface_Expecter) Replace(ctx interface{}, request interface{}) *MockAccountIpAccessListsInterface_Replace_Call {
	return &MockAccountIpAccessListsInterface_Replace_Call{Call: _e.mock.On("Replace", ctx, request)}
}

func (_c *MockAccountIpAccessListsInterface_Replace_Call) Run(run func(ctx context.Context, request settings.ReplaceIpAccessList)) *MockAccountIpAccessListsInterface_Replace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.ReplaceIpAccessList))
	})
	return _c
}

func (_c *MockAccountIpAccessListsInterface_Replace_Call) Return(_a0 error) *MockAccountIpAccessListsInterface_Replace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountIpAccessListsInterface_Replace_Call) RunAndReturn(run func(context.Context, settings.ReplaceIpAccessList) error) *MockAccountIpAccessListsInterface_Replace_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockAccountIpAccessListsInterface) Update(ctx context.Context, request settings.UpdateIpAccessList) error {
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

// MockAccountIpAccessListsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockAccountIpAccessListsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.UpdateIpAccessList
func (_e *MockAccountIpAccessListsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockAccountIpAccessListsInterface_Update_Call {
	return &MockAccountIpAccessListsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockAccountIpAccessListsInterface_Update_Call) Run(run func(ctx context.Context, request settings.UpdateIpAccessList)) *MockAccountIpAccessListsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.UpdateIpAccessList))
	})
	return _c
}

func (_c *MockAccountIpAccessListsInterface_Update_Call) Return(_a0 error) *MockAccountIpAccessListsInterface_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountIpAccessListsInterface_Update_Call) RunAndReturn(run func(context.Context, settings.UpdateIpAccessList) error) *MockAccountIpAccessListsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAccountIpAccessListsInterface creates a new instance of MockAccountIpAccessListsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAccountIpAccessListsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAccountIpAccessListsInterface {
	mock := &MockAccountIpAccessListsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
