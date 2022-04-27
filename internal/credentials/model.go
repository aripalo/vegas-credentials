package credentials

import (
	"io"
	"time"

	"github.com/aripalo/vegas-credentials/internal/assumable"
	"github.com/aripalo/vegas-credentials/internal/sts"
)

type StsCache interface {
	Set(key string, value []byte, ttl time.Duration) error
	Get(key string) ([]byte, error)
	Remove(key string) error
	RemoveByPrefix(keyPrefix string) error
	RemoveAll() error
	Disconnect() error
}

// Credentials defines the output format expected by AWS credential_process
// https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html
type Credentials struct {
	opts            assumable.Opts
	provider        sts.Provider
	cache           StsCache
	output          io.Writer
	Version         int       `json:"Version"`
	AccessKeyID     string    `json:"AccessKeyId"`
	SecretAccessKey string    `json:"SecretAccessKey"`
	SessionToken    string    `json:"SessionToken"`
	Expiration      time.Time `json:"Expiration"`
}
