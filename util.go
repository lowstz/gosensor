package main

import (
	"runtime"
)

func getFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name() + "  "
}
