# Using the testkit

To test your application it is often necessary to mock external dependency.
The SDK contains a testkit you can use for such a purpose, so you don't have to 
use the [`gomock`](https://github.com/golang/mock) framework for code generating 
testing mocks. 

The entrypoint `mocks.NewMockWorkspaceClient` exposes an API that allows you
to retrieve a mock object, on which you can set expectations like so:

```go
client := mocks.NewMockWorkspaceClient(t)
client.GetMockJobsAPI().On("ListAll", mock.Anything, jobs.ListJobsRequest{Limit: 10}).Return(
    []jobs.BaseJob{
        {
            JobId: 64,
        },
        {
            JobId: 65,
        },
    }, nil,
)
```

Please read through [`listjobs_test.go`](listjobs_test.go) test example. 
