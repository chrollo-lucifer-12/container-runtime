package operations

import (
	"encoding/json"
	"fmt"

	"github.com/container-runtime/container"
)

type StateOpts struct {
	ID string
}

func State(opts *StateOpts) (string, error) {
	cntr, err := container.Load(opts.ID)
	if err != nil {
		return "", fmt.Errorf("Load container: %w", err)
	}

	state, err := json.Marshal(cntr.State)
	if err != nil {
		return "", fmt.Errorf("marshal state: %w", err)
	}

	return string(state), nil
}
