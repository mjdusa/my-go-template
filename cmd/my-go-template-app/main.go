package main

import (
	"os"

	"github.com/my-go-template-org/my-go-template-app/internal/runner"
)

func main() {
	os.Exit(runner.Run())
}
