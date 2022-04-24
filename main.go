package main

import (
	"fmt"
	"path/filepath"

	"github.com/aripalo/vegas-credentials/cmd"
	"github.com/aripalo/vegas-credentials/internal/locations"
	"github.com/aripalo/vegas-credentials/internal/mutex"
	"github.com/aripalo/vegas-credentials/internal/utils"
)

const lockFile string = "mutex-lock"

func main() {
	lockPath := filepath.Join(locations.StateDir, lockFile)
	mc, err := mutex.New(lockPath)
	if err != nil {
		utils.Bail(fmt.Sprintf("Configuration Error: %s", err))
	}
	err = mc.Lock()
	if err != nil {
		utils.Bail(fmt.Sprintf("Configuration Error: %s", err))
	}

	defer func() {
		err := mc.Unlock()
		if err != nil {
			utils.Bail(fmt.Sprintf("Configuration Error: %s", err))
		}
	}()

	cmd.Execute()
}
