#!/usr/bin/env sh

command -v golangci-lint >/dev/null 2>&1 || {
  curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.17.1
}
golangci-lint run ./...
