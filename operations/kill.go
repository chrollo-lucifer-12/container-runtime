package operations

import "fmt"

type KillOpt struct {
	ID     string
	Signal string
}

func Kill(opts *KillOpt) error {
	fmt.Println(opts)
	return nil
}
