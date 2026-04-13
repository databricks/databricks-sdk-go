// Hand-written mock stubs for filesAPIUploadUtilities methods.
// These will be replaced by mockery-generated code on the next `make codegen`.

package files

import (
	"context"
	"io"

	files "github.com/databricks/databricks-sdk-go/service/files"
)

// UploadWithChunking provides a mock function for the FilesInterface.
func (_m *MockFilesInterface) UploadWithChunking(ctx context.Context, filePath string, content io.ReadSeeker, contentLength int64, opts ...files.UploadOption) error {
	_va := make([]any, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []any
	_ca = append(_ca, ctx, filePath, content, contentLength)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)
	if len(ret) == 0 {
		panic("no return value specified for UploadWithChunking")
	}
	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, io.ReadSeeker, int64, ...files.UploadOption) error); ok {
		r0 = rf(ctx, filePath, content, contentLength, opts...)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// UploadFromFile provides a mock function for the FilesInterface.
func (_m *MockFilesInterface) UploadFromFile(ctx context.Context, filePath string, sourcePath string, opts ...files.UploadOption) error {
	_va := make([]any, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []any
	_ca = append(_ca, ctx, filePath, sourcePath)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)
	if len(ret) == 0 {
		panic("no return value specified for UploadFromFile")
	}
	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...files.UploadOption) error); ok {
		r0 = rf(ctx, filePath, sourcePath, opts...)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}
