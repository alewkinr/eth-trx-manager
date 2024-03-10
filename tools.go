//go:build tools
// +build tools

// This file is used to track dependencies for go modules.
// ref: https://go.dev/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

package eth_trx_manager

import (
	_ "github.com/abice/go-enum"
	_ "golang.org/x/tools/cmd/goimports"
)
