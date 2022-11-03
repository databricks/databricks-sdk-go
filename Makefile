default: build

fmt:
	@echo "✓ Formatting source code with goimports ..."
	@goimports -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")
	@echo "✓ Formatting source code with gofmt ..."
	@gofmt -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")

lint: vendor
	@echo "✓ Linting source code with https://staticcheck.io/ ..."
	@staticcheck ./...

test: lint
	@echo "✓ Running tests ..."
	@gotestsum --format pkgname-and-test-fails \
		--no-summary=skipped --raw-command go test -v \
		-json -short -coverprofile=coverage.txt ./...

coverage: test
	@echo "✓ Opening coverage for unit tests ..."
	@go tool cover -html=coverage.txt

vendor:
	@echo "✓ Filling vendor folder with library code ..."
	@go mod vendor

doc:
	@echo "Open http://localhost:6060"
	@godoc -http=localhost:6060

.PHONY: fmt vendor fmt coverage test lint doc
