# Interoperability with `gomock`

When developing large applications, you find yourself in need of mocking APIs. For Go, there's [`gomock`](https://github.com/golang/mock) framework for code generating testing mocks. In this small example, we'll show how to use `gomock` with Databricks SDK for Go.

Please read through [`dbfs_test.go`](dbfs_test.go) test example. 

## Declaring which mocks to generate

```go
//go:generate go run github.com/golang/mock/mockgen@latest -package=mocks -destination=mocks/dbfs.go github.com/databricks/databricks-sdk-go/service/dbfs DbfsService
```

* `go run github.com/golang/mock/mockgen@latest` downloads and executes the latest version of `mockgen` command
* `-package=mocks` instructs to generate mocks in the `mocks` package
* `-destination=mocks/dbfs.go` instructs to create `dbfs.go` file with mock stubs.
* `github.com/databricks/databricks-sdk-go/service/dbfs` tells which Databricks package to look services in.
* `DbfsService` tells which services to generate mocks for.

## Initializing `gomock`

Every test needs the following preamble:

```go
ctrl := gomock.NewController(t)
defer ctrl.Finish()
```

## Mocking individual methods with `gomock`

Every actual method call must be mocked for the test to pass:

```go
mockDbfs := mocks.NewMockDbfsService(ctrl)
mockDbfs.EXPECT().Create(gomock.Any(), gomock.Eq(dbfs.Create{
    Path:      "/a/b/c",
    Overwrite: true,
})).Return(&dbfs.CreateResponse{
    Handle: 123,
}, nil)
```

## Testing idioms with Databricks SDK for Go

You can stub out the auth with the `databricks.NewMockConfig(nil)` helper function. Every service has a public `WithImpl()` method, that you can use to set the stubs for every service that is called in the unit tests.

```go
w := workspaces.New(databricks.NewMockConfig(nil))
w.Dbfs.WithImpl(mockDbfs)
```

## Running this example

1. Run `go mod tidy` in this folder to create `go.sum` file to pick dependency versions.
2. Run `go mod vendor` to download dependencies into `vendor/` directory.
3. Run `go generate ./...` to create `mocks/` directory.
4. Run `go test ./...` to invoke tests with mocks.