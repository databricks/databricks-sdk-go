default: build

build: vendor
	@echo "✓ Building source code with go build ..."
	@go build -mod vendor -v

fmt:
	@echo "✓ Installing goimports ..."
	@go install golang.org/x/tools/cmd/goimports@latest
	@echo "✓ Formatting source code with goimports ..."
	@find . -type f -name '*.go' -not -path "./vendor/*" -exec goimports -w {} +
	@echo "✓ Formatting source code with gofmt ..."
	@find . -type f -name '*.go' -not -path "./vendor/*" -exec gofmt -w {} +

lint: vendor
	@echo "✓ Linting source code with https://staticcheck.io/ ..."
	@for dir in */; do \
	    if [ -f "$$dir/go.mod" ]; then \
	        echo "Linting $$dir..."; \
	        cd $$dir && \
			go run honnef.co/go/tools/cmd/staticcheck@v0.5.1 ./... || exit 1 && \
	        cd ..; \
	    fi \
    done

test: vendor
	@echo "✓ Running tests in all modules ..."
	@for dir in */; do \
	    if [ -f "$$dir/go.mod" ]; then \
	        echo "Testing $$dir..."; \
	        cd $$dir && \
			go run gotest.tools/gotestsum@latest --format pkgname-and-test-fails \
			--no-summary=skipped --raw-command go test -v \
			-json -short -coverprofile=coverage.txt  -covermode=count ./... || exit 1 && \
	        cd ..; \
	    fi \
    done

vendor:
	@echo "✓ Filling vendor folders in all modules ..."
	@for dir in */; do \
	    if [ -f "$$dir/go.mod" ]; then \
	        echo "Vendoring $$dir..."; \
	        cd $$dir && \
	        go mod vendor && \
	        cd ..; \
	    fi \
	done

download:
	@echo "✓ Downloading dependencies in all modules ..."
	@for dir in */; do \
		if [ -f "$$dir/go.mod" ]; then \
			echo "Downloading $$dir..."; \
			cd $$dir && \
			go mod download && \
			cd ..; \
		fi \
	done

doc:
	@echo "Open http://localhost:6060"
	@go run golang.org/x/tools/cmd/godoc@latest -http=localhost:6060

.PHONY: fmt vendor fmt coverage test lint doc
