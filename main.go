package main

import (
	"fmt"
	"path/filepath"

	"github.com/aripalo/vegas-credentials/cmd"
	"github.com/aripalo/vegas-credentials/internal/config/locations"
	"github.com/aripalo/vegas-credentials/internal/msg"
	"github.com/aripalo/vegas-credentials/internal/mutex"
)

const lockFile string = "mutex-lock"
const lockErrPrefix string = "Lock Error: %s"

// TODO msg not initialized...?

func main() {
	lockPath := filepath.Join(locations.StateDir, lockFile)
	mc, err := mutex.New(lockPath)
	if err != nil {
		msg.Bail(fmt.Sprintf(lockErrPrefix, err))
	}
	err = mc.Lock()
	if err != nil {
		msg.Bail(fmt.Sprintf(lockErrPrefix, err))
	}

	defer func() {
		_ = mc.Unlock()
	}()

	cmd.Execute()
}
