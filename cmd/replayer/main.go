package main

import (
	"github.com/v2pro/koala/envarg"
	_ "github.com/v2pro/koala/gateway/gw4libc"
)

func init() {
	envarg.SetupLogging()
}

func main() {
}
