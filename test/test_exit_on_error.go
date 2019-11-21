package main

import (
	"errors"
	"github.com/hekonsek/osexit"
)

func main() {
	osexit.ExitOnError(errors.New("ops!"))
}
