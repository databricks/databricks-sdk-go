// Code generated by mockery v2.43.0. DO NOT EDIT.

package marketplace

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	marketplace "github.com/databricks/databricks-sdk-go/service/marketplace"

	mock "github.com/stretchr/testify/mock"
)

// MockConsumerListingsInterface is an autogenerated mock type for the ConsumerListingsInterface type
type MockConsumerListingsInterface struct {
	mock.Mock
}

type MockConsumerListingsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConsumerListingsInterface) EXPECT() *MockConsumerListingsInterface_Expecter {
	return &MockConsumerListingsInterface_Expecter{mock: &_m.Mock}
}

// BatchGet provides a mock function with given fields: ctx, request
func (_m *MockConsumerListingsInterface) BatchGet(ctx context.Context, request marketplace.BatchGetListingsRequest) (*marketplace.BatchGetListingsResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for BatchGet")
	}

	var r0 *marketplace.BatchGetListingsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.BatchGetListingsRequest) (*marketplace.BatchGetListingsResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.BatchGetListingsRequest) *marketplace.BatchGetListingsResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplace.BatchGetListingsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, marketplace.BatchGetListingsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConsumerListingsInterface_BatchGet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BatchGet'
type MockConsumerListingsInterface_BatchGet_Call struct {
	*mock.Call
}

// BatchGet is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.BatchGetListingsRequest
func (_e *MockConsumerListingsInterface_Expecter) BatchGet(ctx interface{}, request interface{}) *MockConsumerListingsInterface_BatchGet_Call {
	return &MockConsumerListingsInterface_BatchGet_Call{Call: _e.mock.On("BatchGet", ctx, request)}
}

func (_c *MockConsumerListingsInterface_BatchGet_Call) Run(run func(ctx context.Context, request marketplace.BatchGetListingsRequest)) *MockConsumerListingsInterface_BatchGet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.BatchGetListingsRequest))
	})
	return _c
}

func (_c *MockConsumerListingsInterface_BatchGet_Call) Return(_a0 *marketplace.BatchGetListingsResponse, _a1 error) *MockConsumerListingsInterface_BatchGet_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConsumerListingsInterface_BatchGet_Call) RunAndReturn(run func(context.Context, marketplace.BatchGetListingsRequest) (*marketplace.BatchGetListingsResponse, error)) *MockConsumerListingsInterface_BatchGet_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockConsumerListingsInterface) Get(ctx context.Context, request marketplace.GetListingRequest) (*marketplace.GetListingResponse, error) {
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

// MockConsumerListingsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockConsumerListingsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.GetListingRequest
func (_e *MockConsumerListingsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockConsumerListingsInterface_Get_Call {
	return &MockConsumerListingsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockConsumerListingsInterface_Get_Call) Run(run func(ctx context.Context, request marketplace.GetListingRequest)) *MockConsumerListingsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.GetListingRequest))
	})
	return _c
}

func (_c *MockConsumerListingsInterface_Get_Call) Return(_a0 *marketplace.GetListingResponse, _a1 error) *MockConsumerListingsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConsumerListingsInterface_Get_Call) RunAndReturn(run func(context.Context, marketplace.GetListingRequest) (*marketplace.GetListingResponse, error)) *MockConsumerListingsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: ctx, id
func (_m *MockConsumerListingsInterface) GetById(ctx context.Context, id string) (*marketplace.GetListingResponse, error) {
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

// MockConsumerListingsInterface_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type MockConsumerListingsInterface_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockConsumerListingsInterface_Expecter) GetById(ctx interface{}, id interface{}) *MockConsumerListingsInterface_GetById_Call {
	return &MockConsumerListingsInterface_GetById_Call{Call: _e.mock.On("GetById", ctx, id)}
}

func (_c *MockConsumerListingsInterface_GetById_Call) Run(run func(ctx context.Context, id string)) *MockConsumerListingsInterface_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockConsumerListingsInterface_GetById_Call) Return(_a0 *marketplace.GetListingResponse, _a1 error) *MockConsumerListingsInterface_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConsumerListingsInterface_GetById_Call) RunAndReturn(run func(context.Context, string) (*marketplace.GetListingResponse, error)) *MockConsumerListingsInterface_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// GetBySummaryName provides a mock function with given fields: ctx, name
func (_m *MockConsumerListingsInterface) GetBySummaryName(ctx context.Context, name string) (*marketplace.Listing, error) {
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

// MockConsumerListingsInterface_GetBySummaryName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBySummaryName'
type MockConsumerListingsInterface_GetBySummaryName_Call struct {
	*mock.Call
}

// GetBySummaryName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockConsumerListingsInterface_Expecter) GetBySummaryName(ctx interface{}, name interface{}) *MockConsumerListingsInterface_GetBySummaryName_Call {
	return &MockConsumerListingsInterface_GetBySummaryName_Call{Call: _e.mock.On("GetBySummaryName", ctx, name)}
}

func (_c *MockConsumerListingsInterface_GetBySummaryName_Call) Run(run func(ctx context.Context, name string)) *MockConsumerListingsInterface_GetBySummaryName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockConsumerListingsInterface_GetBySummaryName_Call) Return(_a0 *marketplace.Listing, _a1 error) *MockConsumerListingsInterface_GetBySummaryName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConsumerListingsInterface_GetBySummaryName_Call) RunAndReturn(run func(context.Context, string) (*marketplace.Listing, error)) *MockConsumerListingsInterface_GetBySummaryName_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockConsumerListingsInterface) Impl() marketplace.ConsumerListingsService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 marketplace.ConsumerListingsService
	if rf, ok := ret.Get(0).(func() marketplace.ConsumerListingsService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(marketplace.ConsumerListingsService)
		}
	}

	return r0
}

// MockConsumerListingsInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockConsumerListingsInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockConsumerListingsInterface_Expecter) Impl() *MockConsumerListingsInterface_Impl_Call {
	return &MockConsumerListingsInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockConsumerListingsInterface_Impl_Call) Run(run func()) *MockConsumerListingsInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConsumerListingsInterface_Impl_Call) Return(_a0 marketplace.ConsumerListingsService) *MockConsumerListingsInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConsumerListingsInterface_Impl_Call) RunAndReturn(run func() marketplace.ConsumerListingsService) *MockConsumerListingsInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockConsumerListingsInterface) List(ctx context.Context, request marketplace.ListListingsRequest) listing.Iterator[marketplace.Listing] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[marketplace.Listing]
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.ListListingsRequest) listing.Iterator[marketplace.Listing]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[marketplace.Listing])
		}
	}

	return r0
}

// MockConsumerListingsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockConsumerListingsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.ListListingsRequest
func (_e *MockConsumerListingsInterface_Expecter) List(ctx interface{}, request interface{}) *MockConsumerListingsInterface_List_Call {
	return &MockConsumerListingsInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockConsumerListingsInterface_List_Call) Run(run func(ctx context.Context, request marketplace.ListListingsRequest)) *MockConsumerListingsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.ListListingsRequest))
	})
	return _c
}

func (_c *MockConsumerListingsInterface_List_Call) Return(_a0 listing.Iterator[marketplace.Listing]) *MockConsumerListingsInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConsumerListingsInterface_List_Call) RunAndReturn(run func(context.Context, marketplace.ListListingsRequest) listing.Iterator[marketplace.Listing]) *MockConsumerListingsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockConsumerListingsInterface) ListAll(ctx context.Context, request marketplace.ListListingsRequest) ([]marketplace.Listing, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []marketplace.Listing
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.ListListingsRequest) ([]marketplace.Listing, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.ListListingsRequest) []marketplace.Listing); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]marketplace.Listing)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, marketplace.ListListingsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConsumerListingsInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockConsumerListingsInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.ListListingsRequest
func (_e *MockConsumerListingsInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockConsumerListingsInterface_ListAll_Call {
	return &MockConsumerListingsInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockConsumerListingsInterface_ListAll_Call) Run(run func(ctx context.Context, request marketplace.ListListingsRequest)) *MockConsumerListingsInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.ListListingsRequest))
	})
	return _c
}

func (_c *MockConsumerListingsInterface_ListAll_Call) Return(_a0 []marketplace.Listing, _a1 error) *MockConsumerListingsInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConsumerListingsInterface_ListAll_Call) RunAndReturn(run func(context.Context, marketplace.ListListingsRequest) ([]marketplace.Listing, error)) *MockConsumerListingsInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// ListingSummaryNameToIdMap provides a mock function with given fields: ctx, request
func (_m *MockConsumerListingsInterface) ListingSummaryNameToIdMap(ctx context.Context, request marketplace.ListListingsRequest) (map[string]string, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListingSummaryNameToIdMap")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.ListListingsRequest) (map[string]string, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.ListListingsRequest) map[string]string); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, marketplace.ListListingsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConsumerListingsInterface_ListingSummaryNameToIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListingSummaryNameToIdMap'
type MockConsumerListingsInterface_ListingSummaryNameToIdMap_Call struct {
	*mock.Call
}

// ListingSummaryNameToIdMap is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.ListListingsRequest
func (_e *MockConsumerListingsInterface_Expecter) ListingSummaryNameToIdMap(ctx interface{}, request interface{}) *MockConsumerListingsInterface_ListingSummaryNameToIdMap_Call {
	return &MockConsumerListingsInterface_ListingSummaryNameToIdMap_Call{Call: _e.mock.On("ListingSummaryNameToIdMap", ctx, request)}
}

func (_c *MockConsumerListingsInterface_ListingSummaryNameToIdMap_Call) Run(run func(ctx context.Context, request marketplace.ListListingsRequest)) *MockConsumerListingsInterface_ListingSummaryNameToIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.ListListingsRequest))
	})
	return _c
}

func (_c *MockConsumerListingsInterface_ListingSummaryNameToIdMap_Call) Return(_a0 map[string]string, _a1 error) *MockConsumerListingsInterface_ListingSummaryNameToIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConsumerListingsInterface_ListingSummaryNameToIdMap_Call) RunAndReturn(run func(context.Context, marketplace.ListListingsRequest) (map[string]string, error)) *MockConsumerListingsInterface_ListingSummaryNameToIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// Search provides a mock function with given fields: ctx, request
func (_m *MockConsumerListingsInterface) Search(ctx context.Context, request marketplace.SearchListingsRequest) listing.Iterator[marketplace.Listing] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Search")
	}

	var r0 listing.Iterator[marketplace.Listing]
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.SearchListingsRequest) listing.Iterator[marketplace.Listing]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[marketplace.Listing])
		}
	}

	return r0
}

// MockConsumerListingsInterface_Search_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Search'
type MockConsumerListingsInterface_Search_Call struct {
	*mock.Call
}

// Search is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.SearchListingsRequest
func (_e *MockConsumerListingsInterface_Expecter) Search(ctx interface{}, request interface{}) *MockConsumerListingsInterface_Search_Call {
	return &MockConsumerListingsInterface_Search_Call{Call: _e.mock.On("Search", ctx, request)}
}

func (_c *MockConsumerListingsInterface_Search_Call) Run(run func(ctx context.Context, request marketplace.SearchListingsRequest)) *MockConsumerListingsInterface_Search_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.SearchListingsRequest))
	})
	return _c
}

func (_c *MockConsumerListingsInterface_Search_Call) Return(_a0 listing.Iterator[marketplace.Listing]) *MockConsumerListingsInterface_Search_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConsumerListingsInterface_Search_Call) RunAndReturn(run func(context.Context, marketplace.SearchListingsRequest) listing.Iterator[marketplace.Listing]) *MockConsumerListingsInterface_Search_Call {
	_c.Call.Return(run)
	return _c
}

// SearchAll provides a mock function with given fields: ctx, request
func (_m *MockConsumerListingsInterface) SearchAll(ctx context.Context, request marketplace.SearchListingsRequest) ([]marketplace.Listing, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for SearchAll")
	}

	var r0 []marketplace.Listing
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.SearchListingsRequest) ([]marketplace.Listing, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, marketplace.SearchListingsRequest) []marketplace.Listing); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]marketplace.Listing)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, marketplace.SearchListingsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConsumerListingsInterface_SearchAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchAll'
type MockConsumerListingsInterface_SearchAll_Call struct {
	*mock.Call
}

// SearchAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request marketplace.SearchListingsRequest
func (_e *MockConsumerListingsInterface_Expecter) SearchAll(ctx interface{}, request interface{}) *MockConsumerListingsInterface_SearchAll_Call {
	return &MockConsumerListingsInterface_SearchAll_Call{Call: _e.mock.On("SearchAll", ctx, request)}
}

func (_c *MockConsumerListingsInterface_SearchAll_Call) Run(run func(ctx context.Context, request marketplace.SearchListingsRequest)) *MockConsumerListingsInterface_SearchAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(marketplace.SearchListingsRequest))
	})
	return _c
}

func (_c *MockConsumerListingsInterface_SearchAll_Call) Return(_a0 []marketplace.Listing, _a1 error) *MockConsumerListingsInterface_SearchAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConsumerListingsInterface_SearchAll_Call) RunAndReturn(run func(context.Context, marketplace.SearchListingsRequest) ([]marketplace.Listing, error)) *MockConsumerListingsInterface_SearchAll_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockConsumerListingsInterface) WithImpl(impl marketplace.ConsumerListingsService) marketplace.ConsumerListingsInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 marketplace.ConsumerListingsInterface
	if rf, ok := ret.Get(0).(func(marketplace.ConsumerListingsService) marketplace.ConsumerListingsInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(marketplace.ConsumerListingsInterface)
		}
	}

	return r0
}

// MockConsumerListingsInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockConsumerListingsInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl marketplace.ConsumerListingsService
func (_e *MockConsumerListingsInterface_Expecter) WithImpl(impl interface{}) *MockConsumerListingsInterface_WithImpl_Call {
	return &MockConsumerListingsInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockConsumerListingsInterface_WithImpl_Call) Run(run func(impl marketplace.ConsumerListingsService)) *MockConsumerListingsInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(marketplace.ConsumerListingsService))
	})
	return _c
}

func (_c *MockConsumerListingsInterface_WithImpl_Call) Return(_a0 marketplace.ConsumerListingsInterface) *MockConsumerListingsInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConsumerListingsInterface_WithImpl_Call) RunAndReturn(run func(marketplace.ConsumerListingsService) marketplace.ConsumerListingsInterface) *MockConsumerListingsInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockConsumerListingsInterface creates a new instance of MockConsumerListingsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConsumerListingsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConsumerListingsInterface {
	mock := &MockConsumerListingsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
