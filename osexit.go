package osexit

import (
	"fmt"
	"os"
	"runtime/debug"
)

const UnixExitCodeGeneralError = 1

func PrintErr(err error) bool {
	if err != nil {
		fmt.Printf("Something went wrong: %s\n", err)
		if os.Getenv("OSEXIT_DEBUG") == "on" {
			debug.PrintStack()
		}
		return true
	}
	return false
}

func ExitOnError(err error) {
	if PrintErr(err) {
		os.Exit(UnixExitCodeGeneralError)
	}
}
