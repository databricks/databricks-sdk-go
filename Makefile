default: build

build: vendor
	@echo "✓ Building source code with go build ..."
	@go build -mod vendor -v

fmt:
	@echo "✓ Formatting source code with goimports ..."
	@go run golang.org/x/tools/cmd/goimports@v0.34.0 -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")
	@echo "✓ Formatting source code with gofmt ..."
	@gofmt -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")

lint: vendor
	@echo "✓ Linting source code with https://staticcheck.io/ ..."
	@go run honnef.co/go/tools/cmd/staticcheck@v0.5.1 ./...

test: vendor
	@echo "✓ Running tests ..."
	@go run gotest.tools/gotestsum@v1.12.2 --format pkgname-and-test-fails \
		--no-summary=skipped --raw-command go test -v \
		-json -short -coverprofile=coverage.txt ./...

coverage: test
	@echo "✓ Opening coverage for unit tests ..."
	@go tool cover -html=coverage.txt

vendor:
	@echo "✓ Filling vendor folder with library code ..."
	@go mod vendor

download:
	@echo "✓ Downloading dependencies ..."
	@go mod download

doc:
	@echo "Open http://localhost:6060"
	@go run golang.org/x/tools/cmd/godoc@v0.34.0 -http=localhost:6060

proto-verify:
	@echo "✓ Verifying protos with protoc ..."
	@protoc -I=protos -I=third_party/googleapis --include_imports --descriptor_set_out=/dev/null $(shell find protos -type f -name '*.proto')

proto-tools:
	@echo "✓ Installing protoc plugins ..."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

proto-gen:
	@echo "✓ Generating Go code from protos ..."
	@PATH="$(shell go env GOPATH)/bin:$$PATH" protoc -I=protos -I=third_party/googleapis --go_out=paths=source_relative:. $(shell find protos -type f -name '*.proto')

.PHONY: fmt vendor fmt coverage test lint doc download proto-verify proto-gen proto-tools
