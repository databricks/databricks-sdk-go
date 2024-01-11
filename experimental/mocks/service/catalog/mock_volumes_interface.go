// Code generated by mockery v2.39.1. DO NOT EDIT.

package catalog

import (
	context "context"

	catalog "github.com/databricks/databricks-sdk-go/service/catalog"

	listing "github.com/databricks/databricks-sdk-go/listing"

	mock "github.com/stretchr/testify/mock"
)

// MockVolumesInterface is an autogenerated mock type for the VolumesInterface type
type MockVolumesInterface struct {
	mock.Mock
}

type MockVolumesInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockVolumesInterface) EXPECT() *MockVolumesInterface_Expecter {
	return &MockVolumesInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockVolumesInterface) Create(ctx context.Context, request catalog.CreateVolumeRequestContent) (*catalog.VolumeInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *catalog.VolumeInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateVolumeRequestContent) (*catalog.VolumeInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateVolumeRequestContent) *catalog.VolumeInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.VolumeInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.CreateVolumeRequestContent) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVolumesInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockVolumesInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.CreateVolumeRequestContent
func (_e *MockVolumesInterface_Expecter) Create(ctx interface{}, request interface{}) *MockVolumesInterface_Create_Call {
	return &MockVolumesInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockVolumesInterface_Create_Call) Run(run func(ctx context.Context, request catalog.CreateVolumeRequestContent)) *MockVolumesInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.CreateVolumeRequestContent))
	})
	return _c
}

func (_c *MockVolumesInterface_Create_Call) Return(_a0 *catalog.VolumeInfo, _a1 error) *MockVolumesInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVolumesInterface_Create_Call) RunAndReturn(run func(context.Context, catalog.CreateVolumeRequestContent) (*catalog.VolumeInfo, error)) *MockVolumesInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockVolumesInterface) Delete(ctx context.Context, request catalog.DeleteVolumeRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.DeleteVolumeRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockVolumesInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockVolumesInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.DeleteVolumeRequest
func (_e *MockVolumesInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockVolumesInterface_Delete_Call {
	return &MockVolumesInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockVolumesInterface_Delete_Call) Run(run func(ctx context.Context, request catalog.DeleteVolumeRequest)) *MockVolumesInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.DeleteVolumeRequest))
	})
	return _c
}

func (_c *MockVolumesInterface_Delete_Call) Return(_a0 error) *MockVolumesInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockVolumesInterface_Delete_Call) RunAndReturn(run func(context.Context, catalog.DeleteVolumeRequest) error) *MockVolumesInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByFullNameArg provides a mock function with given fields: ctx, fullNameArg
func (_m *MockVolumesInterface) DeleteByFullNameArg(ctx context.Context, fullNameArg string) error {
	ret := _m.Called(ctx, fullNameArg)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByFullNameArg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, fullNameArg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockVolumesInterface_DeleteByFullNameArg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByFullNameArg'
type MockVolumesInterface_DeleteByFullNameArg_Call struct {
	*mock.Call
}

// DeleteByFullNameArg is a helper method to define mock.On call
//   - ctx context.Context
//   - fullNameArg string
func (_e *MockVolumesInterface_Expecter) DeleteByFullNameArg(ctx interface{}, fullNameArg interface{}) *MockVolumesInterface_DeleteByFullNameArg_Call {
	return &MockVolumesInterface_DeleteByFullNameArg_Call{Call: _e.mock.On("DeleteByFullNameArg", ctx, fullNameArg)}
}

func (_c *MockVolumesInterface_DeleteByFullNameArg_Call) Run(run func(ctx context.Context, fullNameArg string)) *MockVolumesInterface_DeleteByFullNameArg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockVolumesInterface_DeleteByFullNameArg_Call) Return(_a0 error) *MockVolumesInterface_DeleteByFullNameArg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockVolumesInterface_DeleteByFullNameArg_Call) RunAndReturn(run func(context.Context, string) error) *MockVolumesInterface_DeleteByFullNameArg_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *MockVolumesInterface) GetByName(ctx context.Context, name string) (*catalog.VolumeInfo, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *catalog.VolumeInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*catalog.VolumeInfo, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *catalog.VolumeInfo); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.VolumeInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVolumesInterface_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type MockVolumesInterface_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockVolumesInterface_Expecter) GetByName(ctx interface{}, name interface{}) *MockVolumesInterface_GetByName_Call {
	return &MockVolumesInterface_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *MockVolumesInterface_GetByName_Call) Run(run func(ctx context.Context, name string)) *MockVolumesInterface_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockVolumesInterface_GetByName_Call) Return(_a0 *catalog.VolumeInfo, _a1 error) *MockVolumesInterface_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVolumesInterface_GetByName_Call) RunAndReturn(run func(context.Context, string) (*catalog.VolumeInfo, error)) *MockVolumesInterface_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockVolumesInterface) Impl() catalog.VolumesService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 catalog.VolumesService
	if rf, ok := ret.Get(0).(func() catalog.VolumesService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.VolumesService)
		}
	}

	return r0
}

// MockVolumesInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockVolumesInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockVolumesInterface_Expecter) Impl() *MockVolumesInterface_Impl_Call {
	return &MockVolumesInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockVolumesInterface_Impl_Call) Run(run func()) *MockVolumesInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockVolumesInterface_Impl_Call) Return(_a0 catalog.VolumesService) *MockVolumesInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockVolumesInterface_Impl_Call) RunAndReturn(run func() catalog.VolumesService) *MockVolumesInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockVolumesInterface) List(ctx context.Context, request catalog.ListVolumesRequest) *listing.PaginatingIterator[catalog.ListVolumesRequest, *catalog.ListVolumesResponseContent, catalog.VolumeInfo] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *listing.PaginatingIterator[catalog.ListVolumesRequest, *catalog.ListVolumesResponseContent, catalog.VolumeInfo]
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListVolumesRequest) *listing.PaginatingIterator[catalog.ListVolumesRequest, *catalog.ListVolumesResponseContent, catalog.VolumeInfo]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*listing.PaginatingIterator[catalog.ListVolumesRequest, *catalog.ListVolumesResponseContent, catalog.VolumeInfo])
		}
	}

	return r0
}

// MockVolumesInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockVolumesInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListVolumesRequest
func (_e *MockVolumesInterface_Expecter) List(ctx interface{}, request interface{}) *MockVolumesInterface_List_Call {
	return &MockVolumesInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockVolumesInterface_List_Call) Run(run func(ctx context.Context, request catalog.ListVolumesRequest)) *MockVolumesInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListVolumesRequest))
	})
	return _c
}

func (_c *MockVolumesInterface_List_Call) Return(_a0 *listing.PaginatingIterator[catalog.ListVolumesRequest, *catalog.ListVolumesResponseContent, catalog.VolumeInfo]) *MockVolumesInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockVolumesInterface_List_Call) RunAndReturn(run func(context.Context, catalog.ListVolumesRequest) *listing.PaginatingIterator[catalog.ListVolumesRequest, *catalog.ListVolumesResponseContent, catalog.VolumeInfo]) *MockVolumesInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockVolumesInterface) ListAll(ctx context.Context, request catalog.ListVolumesRequest) ([]catalog.VolumeInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []catalog.VolumeInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListVolumesRequest) ([]catalog.VolumeInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListVolumesRequest) []catalog.VolumeInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]catalog.VolumeInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.ListVolumesRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVolumesInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockVolumesInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListVolumesRequest
func (_e *MockVolumesInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockVolumesInterface_ListAll_Call {
	return &MockVolumesInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockVolumesInterface_ListAll_Call) Run(run func(ctx context.Context, request catalog.ListVolumesRequest)) *MockVolumesInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListVolumesRequest))
	})
	return _c
}

func (_c *MockVolumesInterface_ListAll_Call) Return(_a0 []catalog.VolumeInfo, _a1 error) *MockVolumesInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVolumesInterface_ListAll_Call) RunAndReturn(run func(context.Context, catalog.ListVolumesRequest) ([]catalog.VolumeInfo, error)) *MockVolumesInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: ctx, request
func (_m *MockVolumesInterface) Read(ctx context.Context, request catalog.ReadVolumeRequest) (*catalog.VolumeInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 *catalog.VolumeInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ReadVolumeRequest) (*catalog.VolumeInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ReadVolumeRequest) *catalog.VolumeInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.VolumeInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.ReadVolumeRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVolumesInterface_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type MockVolumesInterface_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ReadVolumeRequest
func (_e *MockVolumesInterface_Expecter) Read(ctx interface{}, request interface{}) *MockVolumesInterface_Read_Call {
	return &MockVolumesInterface_Read_Call{Call: _e.mock.On("Read", ctx, request)}
}

func (_c *MockVolumesInterface_Read_Call) Run(run func(ctx context.Context, request catalog.ReadVolumeRequest)) *MockVolumesInterface_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ReadVolumeRequest))
	})
	return _c
}

func (_c *MockVolumesInterface_Read_Call) Return(_a0 *catalog.VolumeInfo, _a1 error) *MockVolumesInterface_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVolumesInterface_Read_Call) RunAndReturn(run func(context.Context, catalog.ReadVolumeRequest) (*catalog.VolumeInfo, error)) *MockVolumesInterface_Read_Call {
	_c.Call.Return(run)
	return _c
}

// ReadByFullNameArg provides a mock function with given fields: ctx, fullNameArg
func (_m *MockVolumesInterface) ReadByFullNameArg(ctx context.Context, fullNameArg string) (*catalog.VolumeInfo, error) {
	ret := _m.Called(ctx, fullNameArg)

	if len(ret) == 0 {
		panic("no return value specified for ReadByFullNameArg")
	}

	var r0 *catalog.VolumeInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*catalog.VolumeInfo, error)); ok {
		return rf(ctx, fullNameArg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *catalog.VolumeInfo); ok {
		r0 = rf(ctx, fullNameArg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.VolumeInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, fullNameArg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVolumesInterface_ReadByFullNameArg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadByFullNameArg'
type MockVolumesInterface_ReadByFullNameArg_Call struct {
	*mock.Call
}

// ReadByFullNameArg is a helper method to define mock.On call
//   - ctx context.Context
//   - fullNameArg string
func (_e *MockVolumesInterface_Expecter) ReadByFullNameArg(ctx interface{}, fullNameArg interface{}) *MockVolumesInterface_ReadByFullNameArg_Call {
	return &MockVolumesInterface_ReadByFullNameArg_Call{Call: _e.mock.On("ReadByFullNameArg", ctx, fullNameArg)}
}

func (_c *MockVolumesInterface_ReadByFullNameArg_Call) Run(run func(ctx context.Context, fullNameArg string)) *MockVolumesInterface_ReadByFullNameArg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockVolumesInterface_ReadByFullNameArg_Call) Return(_a0 *catalog.VolumeInfo, _a1 error) *MockVolumesInterface_ReadByFullNameArg_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVolumesInterface_ReadByFullNameArg_Call) RunAndReturn(run func(context.Context, string) (*catalog.VolumeInfo, error)) *MockVolumesInterface_ReadByFullNameArg_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockVolumesInterface) Update(ctx context.Context, request catalog.UpdateVolumeRequestContent) (*catalog.VolumeInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *catalog.VolumeInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdateVolumeRequestContent) (*catalog.VolumeInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdateVolumeRequestContent) *catalog.VolumeInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.VolumeInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.UpdateVolumeRequestContent) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVolumesInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockVolumesInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.UpdateVolumeRequestContent
func (_e *MockVolumesInterface_Expecter) Update(ctx interface{}, request interface{}) *MockVolumesInterface_Update_Call {
	return &MockVolumesInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockVolumesInterface_Update_Call) Run(run func(ctx context.Context, request catalog.UpdateVolumeRequestContent)) *MockVolumesInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.UpdateVolumeRequestContent))
	})
	return _c
}

func (_c *MockVolumesInterface_Update_Call) Return(_a0 *catalog.VolumeInfo, _a1 error) *MockVolumesInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVolumesInterface_Update_Call) RunAndReturn(run func(context.Context, catalog.UpdateVolumeRequestContent) (*catalog.VolumeInfo, error)) *MockVolumesInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// VolumeInfoNameToVolumeIdMap provides a mock function with given fields: ctx, request
func (_m *MockVolumesInterface) VolumeInfoNameToVolumeIdMap(ctx context.Context, request catalog.ListVolumesRequest) (map[string]string, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for VolumeInfoNameToVolumeIdMap")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListVolumesRequest) (map[string]string, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListVolumesRequest) map[string]string); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.ListVolumesRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVolumesInterface_VolumeInfoNameToVolumeIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VolumeInfoNameToVolumeIdMap'
type MockVolumesInterface_VolumeInfoNameToVolumeIdMap_Call struct {
	*mock.Call
}

// VolumeInfoNameToVolumeIdMap is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListVolumesRequest
func (_e *MockVolumesInterface_Expecter) VolumeInfoNameToVolumeIdMap(ctx interface{}, request interface{}) *MockVolumesInterface_VolumeInfoNameToVolumeIdMap_Call {
	return &MockVolumesInterface_VolumeInfoNameToVolumeIdMap_Call{Call: _e.mock.On("VolumeInfoNameToVolumeIdMap", ctx, request)}
}

func (_c *MockVolumesInterface_VolumeInfoNameToVolumeIdMap_Call) Run(run func(ctx context.Context, request catalog.ListVolumesRequest)) *MockVolumesInterface_VolumeInfoNameToVolumeIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListVolumesRequest))
	})
	return _c
}

func (_c *MockVolumesInterface_VolumeInfoNameToVolumeIdMap_Call) Return(_a0 map[string]string, _a1 error) *MockVolumesInterface_VolumeInfoNameToVolumeIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVolumesInterface_VolumeInfoNameToVolumeIdMap_Call) RunAndReturn(run func(context.Context, catalog.ListVolumesRequest) (map[string]string, error)) *MockVolumesInterface_VolumeInfoNameToVolumeIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockVolumesInterface) WithImpl(impl catalog.VolumesService) catalog.VolumesInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 catalog.VolumesInterface
	if rf, ok := ret.Get(0).(func(catalog.VolumesService) catalog.VolumesInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.VolumesInterface)
		}
	}

	return r0
}

// MockVolumesInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockVolumesInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl catalog.VolumesService
func (_e *MockVolumesInterface_Expecter) WithImpl(impl interface{}) *MockVolumesInterface_WithImpl_Call {
	return &MockVolumesInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockVolumesInterface_WithImpl_Call) Run(run func(impl catalog.VolumesService)) *MockVolumesInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(catalog.VolumesService))
	})
	return _c
}

func (_c *MockVolumesInterface_WithImpl_Call) Return(_a0 catalog.VolumesInterface) *MockVolumesInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockVolumesInterface_WithImpl_Call) RunAndReturn(run func(catalog.VolumesService) catalog.VolumesInterface) *MockVolumesInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockVolumesInterface creates a new instance of MockVolumesInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockVolumesInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockVolumesInterface {
	mock := &MockVolumesInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
