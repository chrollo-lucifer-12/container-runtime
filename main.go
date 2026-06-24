package main

import (
	"github.com/container-runtime/cli"
)

func main() {
	if err := cli.RootCmd().Execute(); err != nil {
		panic(err)
	}
}
