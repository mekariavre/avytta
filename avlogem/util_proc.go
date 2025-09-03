package avlogem

import (
	"runtime"
)

// getCallerInfo returns the file and line of the caller of LogStack
func getCallerInfo() (pc uintptr, file string, line int, ok bool) {
	// 2 skips: getCallerInfo -> LogStack -> user
	return getCaller(2)
}

// getCaller is a helper for testability
func getCaller(skip int) (uintptr, string, int, bool) {
	return runtime.Caller(skip)
}
