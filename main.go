package main

import (
	"fmt"

	"github.com/container-runtime/cli"
)

func main() {
	if err := cli.RootCmd().Execute(); err != nil {
		fmt.Println(err)
	}
}
