package main

import "github.com/0x71rex/hun7rex/internal/runner"

func main() {
	opts := runner.ParseOptions()
	runner.RunEnumeration(opts)
}
