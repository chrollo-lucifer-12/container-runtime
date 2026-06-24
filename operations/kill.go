package operations

import (
	"fmt"

	"github.com/container-runtime/container"
)

type KillOpt struct {
	ID     string
	Signal string
}

func Kill(opts *KillOpt) error {

	cntr, err := container.Load(opts.ID)
	if err != nil {
		return fmt.Errorf("load container: %w", err)
	}

	// sig, err := strconv.Atoi(opts.Signal)
	// if err != nil {
	// 	return fmt.Errorf("convert signal to int: %w", err)
	// }

	if err := cntr.Kill(); err != nil {
		return fmt.Errorf("kill container: %w", err)
	}

	if err := cntr.Save(); err != nil {
		return fmt.Errorf("save container: %w", err)
	}

	return nil
}
