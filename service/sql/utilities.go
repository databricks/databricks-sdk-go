package sql

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/retries"
)

// [EXPERIMENTAL] Execute a query and wait for results to be available
func (a *StatementExecutionAPI) ExecuteAndWait(ctx context.Context, request ExecuteStatementRequest) (*ExecuteStatementResponse, error) {
	immediateResponse, err := a.impl.ExecuteStatement(ctx, request)
	if err != nil {
		return nil, err
	}
	status := immediateResponse.Status
	switch status.State {
	case StatementStateSucceeded:
		return immediateResponse, nil
	case StatementStateFailed, StatementStateCanceled, StatementStateClosed:
		msg := status.State.String()
		if status.Error != nil {
			msg = fmt.Sprintf("%s: %s %s", msg, status.Error.ErrorCode, status.Error.Message)
		}
		return nil, fmt.Errorf(msg)
	default:
		// TODO: parse request.WaitTimeout and use it here
		return retries.Poll[ExecuteStatementResponse](ctx, 20*time.Minute,
			func() (*ExecuteStatementResponse, *retries.Err) {
				res, err := a.GetStatementByStatementId(ctx, immediateResponse.StatementId)
				if err != nil {
					return nil, retries.Halt(err)
				}
				status := res.Status
				switch status.State {
				case StatementStateSucceeded:
					return &ExecuteStatementResponse{
						Manifest:    res.Manifest,
						Result:      res.Result,
						StatementId: res.StatementId,
						Status:      res.Status,
					}, nil
				case StatementStateFailed, StatementStateCanceled, StatementStateClosed:
					msg := status.State.String()
					if status.Error != nil {
						msg = fmt.Sprintf("%s: %s %s", msg, status.Error.ErrorCode, status.Error.Message)
					}
					return nil, retries.Halt(fmt.Errorf(msg))
				default:
					return nil, retries.Continues(status.State.String())
				}
			})
	}
}
