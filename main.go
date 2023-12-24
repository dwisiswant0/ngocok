package main

import (
	"context"
	"os"

	"net/http"
)

func main() {
	runner := new(runner)
	runner.opt = opt

	if err := runner.start(context.Background()); err != nil && err != http.ErrServerClosed {
		log.Error(errRuntime, "err", err.Error())
		os.Exit(2)
	}
}
