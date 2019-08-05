GO ?= go

all: install vet lint test

install: go.mod
	$(GO) mod download
lint: install
	./scripts/lint.sh
vet: install
	$(GO) vet ./...
test: install
	$(GO) test -v -coverprofile=coverage.txt -covermode=atomic ./...
test-race:
	$(GO) test --race ./...
clean:
	rm -f coverage.out
.PHONY: all install lint vet test test-race clean