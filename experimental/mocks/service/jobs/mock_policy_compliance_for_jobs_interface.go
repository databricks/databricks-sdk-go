// Code generated by mockery v2.53.0. DO NOT EDIT.

package jobs

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	jobs "github.com/databricks/databricks-sdk-go/service/jobs"

	mock "github.com/stretchr/testify/mock"
)

// MockPolicyComplianceForJobsInterface is an autogenerated mock type for the PolicyComplianceForJobsInterface type
type MockPolicyComplianceForJobsInterface struct {
	mock.Mock
}

type MockPolicyComplianceForJobsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPolicyComplianceForJobsInterface) EXPECT() *MockPolicyComplianceForJobsInterface_Expecter {
	return &MockPolicyComplianceForJobsInterface_Expecter{mock: &_m.Mock}
}

// EnforceCompliance provides a mock function with given fields: ctx, request
func (_m *MockPolicyComplianceForJobsInterface) EnforceCompliance(ctx context.Context, request jobs.EnforcePolicyComplianceRequest) (*jobs.EnforcePolicyComplianceResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for EnforceCompliance")
	}

	var r0 *jobs.EnforcePolicyComplianceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, jobs.EnforcePolicyComplianceRequest) (*jobs.EnforcePolicyComplianceResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, jobs.EnforcePolicyComplianceRequest) *jobs.EnforcePolicyComplianceResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jobs.EnforcePolicyComplianceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, jobs.EnforcePolicyComplianceRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPolicyComplianceForJobsInterface_EnforceCompliance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnforceCompliance'
type MockPolicyComplianceForJobsInterface_EnforceCompliance_Call struct {
	*mock.Call
}

// EnforceCompliance is a helper method to define mock.On call
//   - ctx context.Context
//   - request jobs.EnforcePolicyComplianceRequest
func (_e *MockPolicyComplianceForJobsInterface_Expecter) EnforceCompliance(ctx interface{}, request interface{}) *MockPolicyComplianceForJobsInterface_EnforceCompliance_Call {
	return &MockPolicyComplianceForJobsInterface_EnforceCompliance_Call{Call: _e.mock.On("EnforceCompliance", ctx, request)}
}

func (_c *MockPolicyComplianceForJobsInterface_EnforceCompliance_Call) Run(run func(ctx context.Context, request jobs.EnforcePolicyComplianceRequest)) *MockPolicyComplianceForJobsInterface_EnforceCompliance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(jobs.EnforcePolicyComplianceRequest))
	})
	return _c
}

func (_c *MockPolicyComplianceForJobsInterface_EnforceCompliance_Call) Return(_a0 *jobs.EnforcePolicyComplianceResponse, _a1 error) *MockPolicyComplianceForJobsInterface_EnforceCompliance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPolicyComplianceForJobsInterface_EnforceCompliance_Call) RunAndReturn(run func(context.Context, jobs.EnforcePolicyComplianceRequest) (*jobs.EnforcePolicyComplianceResponse, error)) *MockPolicyComplianceForJobsInterface_EnforceCompliance_Call {
	_c.Call.Return(run)
	return _c
}

// GetCompliance provides a mock function with given fields: ctx, request
func (_m *MockPolicyComplianceForJobsInterface) GetCompliance(ctx context.Context, request jobs.GetPolicyComplianceRequest) (*jobs.GetPolicyComplianceResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetCompliance")
	}

	var r0 *jobs.GetPolicyComplianceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, jobs.GetPolicyComplianceRequest) (*jobs.GetPolicyComplianceResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, jobs.GetPolicyComplianceRequest) *jobs.GetPolicyComplianceResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jobs.GetPolicyComplianceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, jobs.GetPolicyComplianceRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPolicyComplianceForJobsInterface_GetCompliance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCompliance'
type MockPolicyComplianceForJobsInterface_GetCompliance_Call struct {
	*mock.Call
}

// GetCompliance is a helper method to define mock.On call
//   - ctx context.Context
//   - request jobs.GetPolicyComplianceRequest
func (_e *MockPolicyComplianceForJobsInterface_Expecter) GetCompliance(ctx interface{}, request interface{}) *MockPolicyComplianceForJobsInterface_GetCompliance_Call {
	return &MockPolicyComplianceForJobsInterface_GetCompliance_Call{Call: _e.mock.On("GetCompliance", ctx, request)}
}

func (_c *MockPolicyComplianceForJobsInterface_GetCompliance_Call) Run(run func(ctx context.Context, request jobs.GetPolicyComplianceRequest)) *MockPolicyComplianceForJobsInterface_GetCompliance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(jobs.GetPolicyComplianceRequest))
	})
	return _c
}

func (_c *MockPolicyComplianceForJobsInterface_GetCompliance_Call) Return(_a0 *jobs.GetPolicyComplianceResponse, _a1 error) *MockPolicyComplianceForJobsInterface_GetCompliance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPolicyComplianceForJobsInterface_GetCompliance_Call) RunAndReturn(run func(context.Context, jobs.GetPolicyComplianceRequest) (*jobs.GetPolicyComplianceResponse, error)) *MockPolicyComplianceForJobsInterface_GetCompliance_Call {
	_c.Call.Return(run)
	return _c
}

// GetComplianceByJobId provides a mock function with given fields: ctx, jobId
func (_m *MockPolicyComplianceForJobsInterface) GetComplianceByJobId(ctx context.Context, jobId int64) (*jobs.GetPolicyComplianceResponse, error) {
	ret := _m.Called(ctx, jobId)

	if len(ret) == 0 {
		panic("no return value specified for GetComplianceByJobId")
	}

	var r0 *jobs.GetPolicyComplianceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*jobs.GetPolicyComplianceResponse, error)); ok {
		return rf(ctx, jobId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *jobs.GetPolicyComplianceResponse); ok {
		r0 = rf(ctx, jobId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jobs.GetPolicyComplianceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, jobId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPolicyComplianceForJobsInterface_GetComplianceByJobId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetComplianceByJobId'
type MockPolicyComplianceForJobsInterface_GetComplianceByJobId_Call struct {
	*mock.Call
}

// GetComplianceByJobId is a helper method to define mock.On call
//   - ctx context.Context
//   - jobId int64
func (_e *MockPolicyComplianceForJobsInterface_Expecter) GetComplianceByJobId(ctx interface{}, jobId interface{}) *MockPolicyComplianceForJobsInterface_GetComplianceByJobId_Call {
	return &MockPolicyComplianceForJobsInterface_GetComplianceByJobId_Call{Call: _e.mock.On("GetComplianceByJobId", ctx, jobId)}
}

func (_c *MockPolicyComplianceForJobsInterface_GetComplianceByJobId_Call) Run(run func(ctx context.Context, jobId int64)) *MockPolicyComplianceForJobsInterface_GetComplianceByJobId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockPolicyComplianceForJobsInterface_GetComplianceByJobId_Call) Return(_a0 *jobs.GetPolicyComplianceResponse, _a1 error) *MockPolicyComplianceForJobsInterface_GetComplianceByJobId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPolicyComplianceForJobsInterface_GetComplianceByJobId_Call) RunAndReturn(run func(context.Context, int64) (*jobs.GetPolicyComplianceResponse, error)) *MockPolicyComplianceForJobsInterface_GetComplianceByJobId_Call {
	_c.Call.Return(run)
	return _c
}

// ListCompliance provides a mock function with given fields: ctx, request
func (_m *MockPolicyComplianceForJobsInterface) ListCompliance(ctx context.Context, request jobs.ListJobComplianceRequest) listing.Iterator[jobs.JobCompliance] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListCompliance")
	}

	var r0 listing.Iterator[jobs.JobCompliance]
	if rf, ok := ret.Get(0).(func(context.Context, jobs.ListJobComplianceRequest) listing.Iterator[jobs.JobCompliance]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[jobs.JobCompliance])
		}
	}

	return r0
}

// MockPolicyComplianceForJobsInterface_ListCompliance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListCompliance'
type MockPolicyComplianceForJobsInterface_ListCompliance_Call struct {
	*mock.Call
}

// ListCompliance is a helper method to define mock.On call
//   - ctx context.Context
//   - request jobs.ListJobComplianceRequest
func (_e *MockPolicyComplianceForJobsInterface_Expecter) ListCompliance(ctx interface{}, request interface{}) *MockPolicyComplianceForJobsInterface_ListCompliance_Call {
	return &MockPolicyComplianceForJobsInterface_ListCompliance_Call{Call: _e.mock.On("ListCompliance", ctx, request)}
}

func (_c *MockPolicyComplianceForJobsInterface_ListCompliance_Call) Run(run func(ctx context.Context, request jobs.ListJobComplianceRequest)) *MockPolicyComplianceForJobsInterface_ListCompliance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(jobs.ListJobComplianceRequest))
	})
	return _c
}

func (_c *MockPolicyComplianceForJobsInterface_ListCompliance_Call) Return(_a0 listing.Iterator[jobs.JobCompliance]) *MockPolicyComplianceForJobsInterface_ListCompliance_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPolicyComplianceForJobsInterface_ListCompliance_Call) RunAndReturn(run func(context.Context, jobs.ListJobComplianceRequest) listing.Iterator[jobs.JobCompliance]) *MockPolicyComplianceForJobsInterface_ListCompliance_Call {
	_c.Call.Return(run)
	return _c
}

// ListComplianceAll provides a mock function with given fields: ctx, request
func (_m *MockPolicyComplianceForJobsInterface) ListComplianceAll(ctx context.Context, request jobs.ListJobComplianceRequest) ([]jobs.JobCompliance, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListComplianceAll")
	}

	var r0 []jobs.JobCompliance
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, jobs.ListJobComplianceRequest) ([]jobs.JobCompliance, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, jobs.ListJobComplianceRequest) []jobs.JobCompliance); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]jobs.JobCompliance)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, jobs.ListJobComplianceRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPolicyComplianceForJobsInterface_ListComplianceAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListComplianceAll'
type MockPolicyComplianceForJobsInterface_ListComplianceAll_Call struct {
	*mock.Call
}

// ListComplianceAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request jobs.ListJobComplianceRequest
func (_e *MockPolicyComplianceForJobsInterface_Expecter) ListComplianceAll(ctx interface{}, request interface{}) *MockPolicyComplianceForJobsInterface_ListComplianceAll_Call {
	return &MockPolicyComplianceForJobsInterface_ListComplianceAll_Call{Call: _e.mock.On("ListComplianceAll", ctx, request)}
}

func (_c *MockPolicyComplianceForJobsInterface_ListComplianceAll_Call) Run(run func(ctx context.Context, request jobs.ListJobComplianceRequest)) *MockPolicyComplianceForJobsInterface_ListComplianceAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(jobs.ListJobComplianceRequest))
	})
	return _c
}

func (_c *MockPolicyComplianceForJobsInterface_ListComplianceAll_Call) Return(_a0 []jobs.JobCompliance, _a1 error) *MockPolicyComplianceForJobsInterface_ListComplianceAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPolicyComplianceForJobsInterface_ListComplianceAll_Call) RunAndReturn(run func(context.Context, jobs.ListJobComplianceRequest) ([]jobs.JobCompliance, error)) *MockPolicyComplianceForJobsInterface_ListComplianceAll_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPolicyComplianceForJobsInterface creates a new instance of MockPolicyComplianceForJobsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPolicyComplianceForJobsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPolicyComplianceForJobsInterface {
	mock := &MockPolicyComplianceForJobsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
