// Code generated by mockery v2.53.2. DO NOT EDIT.

package marketplace

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	marketplace "github.com/databricks/databricks-sdk-go/service/marketplace"

	mock "github.com/stretchr/testify/mock"
)

// MockProviderListingsInterface is an autogenerated mock type for the ProviderListingsInterface type
type MockProviderListingsInterface struct {
	mock.Mock
}

type MockProviderListingsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockProviderListingsInterface) EXPECT() *MockProviderListingsInterface_Expecter {
	return &MockProviderListingsInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockProviderListingsInterface) Create(ctx context.Context, request marketplace.CreateListingRequest) (*marketplace.CreateListingResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *marketplace.CreateListingResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.CreateListingRequest) (*marketplace.CreateListingResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.CreateListingRequest) *marketplace.CreateListingResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplace.CreateListingResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, marketplace.CreateListingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderListingsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockProviderListingsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.CreateListingRequest
func (_e *MockProviderListingsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockProviderListingsInterface_Create_Call {
	return &MockProviderListingsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockProviderListingsInterface_Create_Call) Run(run func(ctx context.Context, request marketplace.CreateListingRequest)) *MockProviderListingsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.CreateListingRequest))
	})
	return _c
}

func (_c *MockProviderListingsInterface_Create_Call) Return(_a0 *marketplace.CreateListingResponse, _a1 error) *MockProviderListingsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderListingsInterface_Create_Call) RunAndReturn(run func(context.Context, marketplace.CreateListingRequest) (*marketplace.CreateListingResponse, error)) *MockProviderListingsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockProviderListingsInterface) Delete(ctx context.Context, request marketplace.DeleteListingRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.DeleteListingRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockProviderListingsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockProviderListingsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.DeleteListingRequest
func (_e *MockProviderListingsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockProviderListingsInterface_Delete_Call {
	return &MockProviderListingsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockProviderListingsInterface_Delete_Call) Run(run func(ctx context.Context, request marketplace.DeleteListingRequest)) *MockProviderListingsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.DeleteListingRequest))
	})
	return _c
}

func (_c *MockProviderListingsInterface_Delete_Call) Return(_a0 error) *MockProviderListingsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProviderListingsInterface_Delete_Call) RunAndReturn(run func(context.Context, marketplace.DeleteListingRequest) error) *MockProviderListingsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *MockProviderListingsInterface) DeleteById(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteById")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockProviderListingsInterface_DeleteById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteById'
type MockProviderListingsInterface_DeleteById_Call struct {
	*mock.Call
}

// DeleteById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockProviderListingsInterface_Expecter) DeleteById(ctx interface{}, id interface{}) *MockProviderListingsInterface_DeleteById_Call {
	return &MockProviderListingsInterface_DeleteById_Call{Call: _e.mock.On("DeleteById", ctx, id)}
}

func (_c *MockProviderListingsInterface_DeleteById_Call) Run(run func(ctx context.Context, id string)) *MockProviderListingsInterface_DeleteById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockProviderListingsInterface_DeleteById_Call) Return(_a0 error) *MockProviderListingsInterface_DeleteById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProviderListingsInterface_DeleteById_Call) RunAndReturn(run func(context.Context, string) error) *MockProviderListingsInterface_DeleteById_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockProviderListingsInterface) Get(ctx context.Context, request marketplace.GetListingRequest) (*marketplace.GetListingResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *marketplace.GetListingResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.GetListingRequest) (*marketplace.GetListingResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.GetListingRequest) *marketplace.GetListingResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplace.GetListingResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, marketplace.GetListingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderListingsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockProviderListingsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.GetListingRequest
func (_e *MockProviderListingsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockProviderListingsInterface_Get_Call {
	return &MockProviderListingsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockProviderListingsInterface_Get_Call) Run(run func(ctx context.Context, request marketplace.GetListingRequest)) *MockProviderListingsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.GetListingRequest))
	})
	return _c
}

func (_c *MockProviderListingsInterface_Get_Call) Return(_a0 *marketplace.GetListingResponse, _a1 error) *MockProviderListingsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderListingsInterface_Get_Call) RunAndReturn(run func(context.Context, marketplace.GetListingRequest) (*marketplace.GetListingResponse, error)) *MockProviderListingsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: ctx, id
func (_m *MockProviderListingsInterface) GetById(ctx context.Context, id string) (*marketplace.GetListingResponse, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *marketplace.GetListingResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*marketplace.GetListingResponse, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *marketplace.GetListingResponse); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplace.GetListingResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderListingsInterface_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type MockProviderListingsInterface_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockProviderListingsInterface_Expecter) GetById(ctx interface{}, id interface{}) *MockProviderListingsInterface_GetById_Call {
	return &MockProviderListingsInterface_GetById_Call{Call: _e.mock.On("GetById", ctx, id)}
}

func (_c *MockProviderListingsInterface_GetById_Call) Run(run func(ctx context.Context, id string)) *MockProviderListingsInterface_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockProviderListingsInterface_GetById_Call) Return(_a0 *marketplace.GetListingResponse, _a1 error) *MockProviderListingsInterface_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderListingsInterface_GetById_Call) RunAndReturn(run func(context.Context, string) (*marketplace.GetListingResponse, error)) *MockProviderListingsInterface_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// GetBySummaryName provides a mock function with given fields: ctx, name
func (_m *MockProviderListingsInterface) GetBySummaryName(ctx context.Context, name string) (*marketplace.Listing, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetBySummaryName")
	}

	var r0 *marketplace.Listing
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*marketplace.Listing, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *marketplace.Listing); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplace.Listing)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderListingsInterface_GetBySummaryName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBySummaryName'
type MockProviderListingsInterface_GetBySummaryName_Call struct {
	*mock.Call
}

// GetBySummaryName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockProviderListingsInterface_Expecter) GetBySummaryName(ctx interface{}, name interface{}) *MockProviderListingsInterface_GetBySummaryName_Call {
	return &MockProviderListingsInterface_GetBySummaryName_Call{Call: _e.mock.On("GetBySummaryName", ctx, name)}
}

func (_c *MockProviderListingsInterface_GetBySummaryName_Call) Run(run func(ctx context.Context, name string)) *MockProviderListingsInterface_GetBySummaryName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockProviderListingsInterface_GetBySummaryName_Call) Return(_a0 *marketplace.Listing, _a1 error) *MockProviderListingsInterface_GetBySummaryName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderListingsInterface_GetBySummaryName_Call) RunAndReturn(run func(context.Context, string) (*marketplace.Listing, error)) *MockProviderListingsInterface_GetBySummaryName_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockProviderListingsInterface) List(ctx context.Context, request marketplace.GetListingsRequest) listing.Iterator[marketplace.Listing] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[marketplace.Listing]
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.GetListingsRequest) listing.Iterator[marketplace.Listing]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[marketplace.Listing])
		}
	}

	return r0
}

// MockProviderListingsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockProviderListingsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.GetListingsRequest
func (_e *MockProviderListingsInterface_Expecter) List(ctx interface{}, request interface{}) *MockProviderListingsInterface_List_Call {
	return &MockProviderListingsInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockProviderListingsInterface_List_Call) Run(run func(ctx context.Context, request marketplace.GetListingsRequest)) *MockProviderListingsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.GetListingsRequest))
	})
	return _c
}

func (_c *MockProviderListingsInterface_List_Call) Return(_a0 listing.Iterator[marketplace.Listing]) *MockProviderListingsInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProviderListingsInterface_List_Call) RunAndReturn(run func(context.Context, marketplace.GetListingsRequest) listing.Iterator[marketplace.Listing]) *MockProviderListingsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockProviderListingsInterface) ListAll(ctx context.Context, request marketplace.GetListingsRequest) ([]marketplace.Listing, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []marketplace.Listing
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.GetListingsRequest) ([]marketplace.Listing, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.GetListingsRequest) []marketplace.Listing); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]marketplace.Listing)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, marketplace.GetListingsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderListingsInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockProviderListingsInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.GetListingsRequest
func (_e *MockProviderListingsInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockProviderListingsInterface_ListAll_Call {
	return &MockProviderListingsInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockProviderListingsInterface_ListAll_Call) Run(run func(ctx context.Context, request marketplace.GetListingsRequest)) *MockProviderListingsInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.GetListingsRequest))
	})
	return _c
}

func (_c *MockProviderListingsInterface_ListAll_Call) Return(_a0 []marketplace.Listing, _a1 error) *MockProviderListingsInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderListingsInterface_ListAll_Call) RunAndReturn(run func(context.Context, marketplace.GetListingsRequest) ([]marketplace.Listing, error)) *MockProviderListingsInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// ListingSummaryNameToIdMap provides a mock function with given fields: ctx, request
func (_m *MockProviderListingsInterface) ListingSummaryNameToIdMap(ctx context.Context, request marketplace.GetListingsRequest) (map[string]string, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListingSummaryNameToIdMap")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.GetListingsRequest) (map[string]string, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.GetListingsRequest) map[string]string); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, marketplace.GetListingsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderListingsInterface_ListingSummaryNameToIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListingSummaryNameToIdMap'
type MockProviderListingsInterface_ListingSummaryNameToIdMap_Call struct {
	*mock.Call
}

// ListingSummaryNameToIdMap is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.GetListingsRequest
func (_e *MockProviderListingsInterface_Expecter) ListingSummaryNameToIdMap(ctx interface{}, request interface{}) *MockProviderListingsInterface_ListingSummaryNameToIdMap_Call {
	return &MockProviderListingsInterface_ListingSummaryNameToIdMap_Call{Call: _e.mock.On("ListingSummaryNameToIdMap", ctx, request)}
}

func (_c *MockProviderListingsInterface_ListingSummaryNameToIdMap_Call) Run(run func(ctx context.Context, request marketplace.GetListingsRequest)) *MockProviderListingsInterface_ListingSummaryNameToIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.GetListingsRequest))
	})
	return _c
}

func (_c *MockProviderListingsInterface_ListingSummaryNameToIdMap_Call) Return(_a0 map[string]string, _a1 error) *MockProviderListingsInterface_ListingSummaryNameToIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderListingsInterface_ListingSummaryNameToIdMap_Call) RunAndReturn(run func(context.Context, marketplace.GetListingsRequest) (map[string]string, error)) *MockProviderListingsInterface_ListingSummaryNameToIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockProviderListingsInterface) Update(ctx context.Context, request marketplace.UpdateListingRequest) (*marketplace.UpdateListingResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *marketplace.UpdateListingResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.UpdateListingRequest) (*marketplace.UpdateListingResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.UpdateListingRequest) *marketplace.UpdateListingResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplace.UpdateListingResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, marketplace.UpdateListingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderListingsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockProviderListingsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.UpdateListingRequest
func (_e *MockProviderListingsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockProviderListingsInterface_Update_Call {
	return &MockProviderListingsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockProviderListingsInterface_Update_Call) Run(run func(ctx context.Context, request marketplace.UpdateListingRequest)) *MockProviderListingsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.UpdateListingRequest))
	})
	return _c
}

func (_c *MockProviderListingsInterface_Update_Call) Return(_a0 *marketplace.UpdateListingResponse, _a1 error) *MockProviderListingsInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderListingsInterface_Update_Call) RunAndReturn(run func(context.Context, marketplace.UpdateListingRequest) (*marketplace.UpdateListingResponse, error)) *MockProviderListingsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockProviderListingsInterface creates a new instance of MockProviderListingsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockProviderListingsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockProviderListingsInterface {
	mock := &MockProviderListingsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
