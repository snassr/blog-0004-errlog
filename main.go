package main

import (
	"fmt"
	"io"
	"os"

	"github.com/snassr/blog-0004-errlog/errors"
	"github.com/snassr/blog-0004-errlog/logg"
)

var log *logg.Logger

func main() {
	// initialize logging to standard out (can add file, or database io.Writer as well)
	log = logg.Init(
		[]io.Writer{os.Stdout},
		[]io.Writer{os.Stdout},
		[]io.Writer{os.Stdout},
		[]io.Writer{os.Stdout},
		[]io.Writer{os.Stdout},
	)

	// mock hasPermission error
	if hasPermission(false) {
		log.Warn(errors.Error{
			Op:   "main.hasPermission",
			Kind: errors.Permission,
			Err:  "permission denied",
		})
	}

	// mock data error
	_, err1 := getData("no data")
	if err1 != nil {
		log.Info(errors.Error{
			Op:   "main.getData",
			Kind: errors.Other,
			Err:  err1.Error(),
		})
	}

	// mock operation error
	_, err2 := mathoperation("^+", 1, 2)
	if err2 != nil {
		log.Err(errors.Error{
			Op:   "main.mathoperation",
			Kind: errors.Invalid,
			Err:  err2.Error(),
		})
	}

	// mock fatal error (server failed...)
	log.Fatal(errors.Error{
		Op:   "main",
		Kind: errors.Other,
		Err:  "fatal service error",
	})
}

func hasPermission(b bool) bool {
	return !b
}

func getData(data string) (string, error) {
	return "", fmt.Errorf(data)
}

func mathoperation(operator string, a, b int) (int, error) {
	return -1, fmt.Errorf("invalid operator: %v", operator)
}
