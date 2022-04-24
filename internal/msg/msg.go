package msg

import (
	"github.com/aripalo/go-delightful"
	"github.com/aripalo/vegas-credentials/internal/config"
)

// Initialize
var Message = delightful.New(config.AppName)
