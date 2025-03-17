// Code generated by mockery v2.53.2. DO NOT EDIT.

package sql

import (
	context "context"

	sql "github.com/databricks/databricks-sdk-go/service/sql"
	mock "github.com/stretchr/testify/mock"
)

// MockStatementExecutionInterface is an autogenerated mock type for the StatementExecutionInterface type
type MockStatementExecutionInterface struct {
	mock.Mock
}

type MockStatementExecutionInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStatementExecutionInterface) EXPECT() *MockStatementExecutionInterface_Expecter {
	return &MockStatementExecutionInterface_Expecter{mock: &_m.Mock}
}

// CancelExecution provides a mock function with given fields: ctx, request
func (_m *MockStatementExecutionInterface) CancelExecution(ctx context.Context, request sql.CancelExecutionRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for CancelExecution")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.CancelExecutionRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStatementExecutionInterface_CancelExecution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CancelExecution'
type MockStatementExecutionInterface_CancelExecution_Call struct {
	*mock.Call
}

// CancelExecution is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.CancelExecutionRequest
func (_e *MockStatementExecutionInterface_Expecter) CancelExecution(ctx interface{}, request interface{}) *MockStatementExecutionInterface_CancelExecution_Call {
	return &MockStatementExecutionInterface_CancelExecution_Call{Call: _e.mock.On("CancelExecution", ctx, request)}
}

func (_c *MockStatementExecutionInterface_CancelExecution_Call) Run(run func(ctx context.Context, request sql.CancelExecutionRequest)) *MockStatementExecutionInterface_CancelExecution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.CancelExecutionRequest))
	})
	return _c
}

func (_c *MockStatementExecutionInterface_CancelExecution_Call) Return(_a0 error) *MockStatementExecutionInterface_CancelExecution_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStatementExecutionInterface_CancelExecution_Call) RunAndReturn(run func(context.Context, sql.CancelExecutionRequest) error) *MockStatementExecutionInterface_CancelExecution_Call {
	_c.Call.Return(run)
	return _c
}

// ExecuteAndWait provides a mock function with given fields: ctx, request
func (_m *MockStatementExecutionInterface) ExecuteAndWait(ctx context.Context, request sql.ExecuteStatementRequest) (*sql.StatementResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ExecuteAndWait")
	}

	var r0 *sql.StatementResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.ExecuteStatementRequest) (*sql.StatementResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sql.ExecuteStatementRequest) *sql.StatementResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.StatementResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sql.ExecuteStatementRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStatementExecutionInterface_ExecuteAndWait_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteAndWait'
type MockStatementExecutionInterface_ExecuteAndWait_Call struct {
	*mock.Call
}

// ExecuteAndWait is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.ExecuteStatementRequest
func (_e *MockStatementExecutionInterface_Expecter) ExecuteAndWait(ctx interface{}, request interface{}) *MockStatementExecutionInterface_ExecuteAndWait_Call {
	return &MockStatementExecutionInterface_ExecuteAndWait_Call{Call: _e.mock.On("ExecuteAndWait", ctx, request)}
}

func (_c *MockStatementExecutionInterface_ExecuteAndWait_Call) Run(run func(ctx context.Context, request sql.ExecuteStatementRequest)) *MockStatementExecutionInterface_ExecuteAndWait_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.ExecuteStatementRequest))
	})
	return _c
}

func (_c *MockStatementExecutionInterface_ExecuteAndWait_Call) Return(_a0 *sql.StatementResponse, _a1 error) *MockStatementExecutionInterface_ExecuteAndWait_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStatementExecutionInterface_ExecuteAndWait_Call) RunAndReturn(run func(context.Context, sql.ExecuteStatementRequest) (*sql.StatementResponse, error)) *MockStatementExecutionInterface_ExecuteAndWait_Call {
	_c.Call.Return(run)
	return _c
}

// ExecuteStatement provides a mock function with given fields: ctx, request
func (_m *MockStatementExecutionInterface) ExecuteStatement(ctx context.Context, request sql.ExecuteStatementRequest) (*sql.StatementResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ExecuteStatement")
	}

	var r0 *sql.StatementResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.ExecuteStatementRequest) (*sql.StatementResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sql.ExecuteStatementRequest) *sql.StatementResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.StatementResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sql.ExecuteStatementRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStatementExecutionInterface_ExecuteStatement_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteStatement'
type MockStatementExecutionInterface_ExecuteStatement_Call struct {
	*mock.Call
}

// ExecuteStatement is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.ExecuteStatementRequest
func (_e *MockStatementExecutionInterface_Expecter) ExecuteStatement(ctx interface{}, request interface{}) *MockStatementExecutionInterface_ExecuteStatement_Call {
	return &MockStatementExecutionInterface_ExecuteStatement_Call{Call: _e.mock.On("ExecuteStatement", ctx, request)}
}

func (_c *MockStatementExecutionInterface_ExecuteStatement_Call) Run(run func(ctx context.Context, request sql.ExecuteStatementRequest)) *MockStatementExecutionInterface_ExecuteStatement_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.ExecuteStatementRequest))
	})
	return _c
}

func (_c *MockStatementExecutionInterface_ExecuteStatement_Call) Return(_a0 *sql.StatementResponse, _a1 error) *MockStatementExecutionInterface_ExecuteStatement_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStatementExecutionInterface_ExecuteStatement_Call) RunAndReturn(run func(context.Context, sql.ExecuteStatementRequest) (*sql.StatementResponse, error)) *MockStatementExecutionInterface_ExecuteStatement_Call {
	_c.Call.Return(run)
	return _c
}

// GetStatement provides a mock function with given fields: ctx, request
func (_m *MockStatementExecutionInterface) GetStatement(ctx context.Context, request sql.GetStatementRequest) (*sql.StatementResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetStatement")
	}

	var r0 *sql.StatementResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.GetStatementRequest) (*sql.StatementResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sql.GetStatementRequest) *sql.StatementResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.StatementResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sql.GetStatementRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStatementExecutionInterface_GetStatement_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatement'
type MockStatementExecutionInterface_GetStatement_Call struct {
	*mock.Call
}

// GetStatement is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.GetStatementRequest
func (_e *MockStatementExecutionInterface_Expecter) GetStatement(ctx interface{}, request interface{}) *MockStatementExecutionInterface_GetStatement_Call {
	return &MockStatementExecutionInterface_GetStatement_Call{Call: _e.mock.On("GetStatement", ctx, request)}
}

func (_c *MockStatementExecutionInterface_GetStatement_Call) Run(run func(ctx context.Context, request sql.GetStatementRequest)) *MockStatementExecutionInterface_GetStatement_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.GetStatementRequest))
	})
	return _c
}

func (_c *MockStatementExecutionInterface_GetStatement_Call) Return(_a0 *sql.StatementResponse, _a1 error) *MockStatementExecutionInterface_GetStatement_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStatementExecutionInterface_GetStatement_Call) RunAndReturn(run func(context.Context, sql.GetStatementRequest) (*sql.StatementResponse, error)) *MockStatementExecutionInterface_GetStatement_Call {
	_c.Call.Return(run)
	return _c
}

// GetStatementByStatementId provides a mock function with given fields: ctx, statementId
func (_m *MockStatementExecutionInterface) GetStatementByStatementId(ctx context.Context, statementId string) (*sql.StatementResponse, error) {
	ret := _m.Called(ctx, statementId)

	if len(ret) == 0 {
		panic("no return value specified for GetStatementByStatementId")
	}

	var r0 *sql.StatementResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*sql.StatementResponse, error)); ok {
		return rf(ctx, statementId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *sql.StatementResponse); ok {
		r0 = rf(ctx, statementId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.StatementResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, statementId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStatementExecutionInterface_GetStatementByStatementId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatementByStatementId'
type MockStatementExecutionInterface_GetStatementByStatementId_Call struct {
	*mock.Call
}

// GetStatementByStatementId is a helper method to define mock.On call
//   - ctx context.Context
//   - statementId string
func (_e *MockStatementExecutionInterface_Expecter) GetStatementByStatementId(ctx interface{}, statementId interface{}) *MockStatementExecutionInterface_GetStatementByStatementId_Call {
	return &MockStatementExecutionInterface_GetStatementByStatementId_Call{Call: _e.mock.On("GetStatementByStatementId", ctx, statementId)}
}

func (_c *MockStatementExecutionInterface_GetStatementByStatementId_Call) Run(run func(ctx context.Context, statementId string)) *MockStatementExecutionInterface_GetStatementByStatementId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockStatementExecutionInterface_GetStatementByStatementId_Call) Return(_a0 *sql.StatementResponse, _a1 error) *MockStatementExecutionInterface_GetStatementByStatementId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStatementExecutionInterface_GetStatementByStatementId_Call) RunAndReturn(run func(context.Context, string) (*sql.StatementResponse, error)) *MockStatementExecutionInterface_GetStatementByStatementId_Call {
	_c.Call.Return(run)
	return _c
}

// GetStatementResultChunkN provides a mock function with given fields: ctx, request
func (_m *MockStatementExecutionInterface) GetStatementResultChunkN(ctx context.Context, request sql.GetStatementResultChunkNRequest) (*sql.ResultData, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetStatementResultChunkN")
	}

	var r0 *sql.ResultData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sql.GetStatementResultChunkNRequest) (*sql.ResultData, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sql.GetStatementResultChunkNRequest) *sql.ResultData); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.ResultData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sql.GetStatementResultChunkNRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStatementExecutionInterface_GetStatementResultChunkN_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatementResultChunkN'
type MockStatementExecutionInterface_GetStatementResultChunkN_Call struct {
	*mock.Call
}

// GetStatementResultChunkN is a helper method to define mock.On call
//   - ctx context.Context
//   - request sql.GetStatementResultChunkNRequest
func (_e *MockStatementExecutionInterface_Expecter) GetStatementResultChunkN(ctx interface{}, request interface{}) *MockStatementExecutionInterface_GetStatementResultChunkN_Call {
	return &MockStatementExecutionInterface_GetStatementResultChunkN_Call{Call: _e.mock.On("GetStatementResultChunkN", ctx, request)}
}

func (_c *MockStatementExecutionInterface_GetStatementResultChunkN_Call) Run(run func(ctx context.Context, request sql.GetStatementResultChunkNRequest)) *MockStatementExecutionInterface_GetStatementResultChunkN_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sql.GetStatementResultChunkNRequest))
	})
	return _c
}

func (_c *MockStatementExecutionInterface_GetStatementResultChunkN_Call) Return(_a0 *sql.ResultData, _a1 error) *MockStatementExecutionInterface_GetStatementResultChunkN_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStatementExecutionInterface_GetStatementResultChunkN_Call) RunAndReturn(run func(context.Context, sql.GetStatementResultChunkNRequest) (*sql.ResultData, error)) *MockStatementExecutionInterface_GetStatementResultChunkN_Call {
	_c.Call.Return(run)
	return _c
}

// GetStatementResultChunkNByStatementIdAndChunkIndex provides a mock function with given fields: ctx, statementId, chunkIndex
func (_m *MockStatementExecutionInterface) GetStatementResultChunkNByStatementIdAndChunkIndex(ctx context.Context, statementId string, chunkIndex int) (*sql.ResultData, error) {
	ret := _m.Called(ctx, statementId, chunkIndex)

	if len(ret) == 0 {
		panic("no return value specified for GetStatementResultChunkNByStatementIdAndChunkIndex")
	}

	var r0 *sql.ResultData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) (*sql.ResultData, error)); ok {
		return rf(ctx, statementId, chunkIndex)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int) *sql.ResultData); ok {
		r0 = rf(ctx, statementId, chunkIndex)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.ResultData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int) error); ok {
		r1 = rf(ctx, statementId, chunkIndex)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStatementExecutionInterface_GetStatementResultChunkNByStatementIdAndChunkIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatementResultChunkNByStatementIdAndChunkIndex'
type MockStatementExecutionInterface_GetStatementResultChunkNByStatementIdAndChunkIndex_Call struct {
	*mock.Call
}

// GetStatementResultChunkNByStatementIdAndChunkIndex is a helper method to define mock.On call
//   - ctx context.Context
//   - statementId string
//   - chunkIndex int
func (_e *MockStatementExecutionInterface_Expecter) GetStatementResultChunkNByStatementIdAndChunkIndex(ctx interface{}, statementId interface{}, chunkIndex interface{}) *MockStatementExecutionInterface_GetStatementResultChunkNByStatementIdAndChunkIndex_Call {
	return &MockStatementExecutionInterface_GetStatementResultChunkNByStatementIdAndChunkIndex_Call{Call: _e.mock.On("GetStatementResultChunkNByStatementIdAndChunkIndex", ctx, statementId, chunkIndex)}
}

func (_c *MockStatementExecutionInterface_GetStatementResultChunkNByStatementIdAndChunkIndex_Call) Run(run func(ctx context.Context, statementId string, chunkIndex int)) *MockStatementExecutionInterface_GetStatementResultChunkNByStatementIdAndChunkIndex_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int))
	})
	return _c
}

func (_c *MockStatementExecutionInterface_GetStatementResultChunkNByStatementIdAndChunkIndex_Call) Return(_a0 *sql.ResultData, _a1 error) *MockStatementExecutionInterface_GetStatementResultChunkNByStatementIdAndChunkIndex_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStatementExecutionInterface_GetStatementResultChunkNByStatementIdAndChunkIndex_Call) RunAndReturn(run func(context.Context, string, int) (*sql.ResultData, error)) *MockStatementExecutionInterface_GetStatementResultChunkNByStatementIdAndChunkIndex_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStatementExecutionInterface creates a new instance of MockStatementExecutionInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStatementExecutionInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStatementExecutionInterface {
	mock := &MockStatementExecutionInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
