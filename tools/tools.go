//go:build tools
// +build tools

// This file uses the recommended method for tracking developer tools in a Go
// module.
//
// REF: https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
package tools

import (
	_ "github.com/golang/mock/mockgen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/osmosis-labs/go-mutesting/cmd/go-mutesting"
	_ "golang.org/x/perf/cmd/benchstat"
)
