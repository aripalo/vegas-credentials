package main

import (
	"fmt"

	"github.com/aripalo/vegas-credentials/cmd"
	"github.com/aripalo/vegas-credentials/internal/msg"
	"github.com/aripalo/vegas-credentials/internal/mutex"
)

// TODO msg not initialized...?

func main() {

	unlock, err := mutex.Lock()
	if err != nil {
		msg.Bail(fmt.Sprintf("Lock Error: %s", err))
	}

	defer func() {
		_ = unlock()
	}()

	cmd.Execute()
}
