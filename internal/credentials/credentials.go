package credentials

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/aripalo/vegas-credentials/internal/assumecfg"
	"github.com/aripalo/vegas-credentials/internal/cache"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/dgraph-io/badger/v3"
	"github.com/dustin/go-humanize"
)

// AWS_CREDENTIAL_PROCESS_VERSION defines the supported AWS credential_process version
const AWS_CREDENTIAL_PROCESS_VERSION int = 1

// Credentials defines the output format expected by AWS credential_process
// https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html
type Credentials struct {
	cfg             assumecfg.AssumeCfg
	repo            cache.Repository
	output          io.Writer
	Version         int       `json:"Version"`
	AccessKeyID     string    `json:"AccessKeyId"`
	SecretAccessKey string    `json:"SecretAccessKey"`
	SessionToken    string    `json:"SessionToken"`
	Expiration      time.Time `json:"Expiration"`
}

// New defines a response waiting to be fulfilled
func New(cfg assumecfg.AssumeCfg) *Credentials {
	r := &Credentials{
		cfg:     cfg,
		output:  os.Stdout,
		repo:    NewCredentialCache(),
		Version: AWS_CREDENTIAL_PROCESS_VERSION,
	}

	return r
}

// Load STS Temporary Session Credentials from cache.
func (c *Credentials) Load() error {
	err := c.readFromCache()
	if err != nil {
		if err == badger.ErrKeyNotFound { // TODO maybe don't expose badger internals?
			return errors.New("Not found")
		}
		return err
	}

	err = c.Validate()
	if err != nil {
		return err
	}

	now := time.Now()

	if c.isExpired(now) {
		return fmt.Errorf("Expired %s", humanize.RelTime(c.Expiration, now, "ago", "in future"))
	}

	if c.isRefreshNeeded() {
		return fmt.Errorf("Refresh required because expiration in %s", humanize.Time(c.Expiration))
	}

	return nil
}

// Retrieve new STS Temporary Session Credentials credentials from AWS.
func (c *Credentials) New(code string) error {

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithSharedConfigProfile(c.cfg.SourceProfile),
		config.WithRegion(c.cfg.Region),
	)

	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := sts.NewFromConfig(cfg)

	input := &sts.AssumeRoleInput{
		RoleArn:      aws.String(c.cfg.RoleArn),
		SerialNumber: aws.String(c.cfg.MfaSerial),
		TokenCode:    aws.String(code),
	}

	if c.cfg.DurationSeconds != 0 {
		input.DurationSeconds = aws.Int32(int32(c.cfg.DurationSeconds))
	}

	if c.cfg.RoleSessionName != "" {
		input.RoleSessionName = aws.String(c.cfg.RoleSessionName)
	}

	if c.cfg.ExternalID != "" {
		input.ExternalId = aws.String(c.cfg.ExternalID)
	}

	result, err := client.AssumeRole(ctx, input)
	if err != nil {
		panic(err)
	}

	c.AccessKeyID = aws.ToString(result.Credentials.AccessKeyId)
	c.SecretAccessKey = aws.ToString(result.Credentials.SecretAccessKey)
	c.SessionToken = aws.ToString(result.Credentials.SessionToken)
	c.Expiration = aws.ToTime(result.Credentials.Expiration)

	err = c.Validate()
	if err != nil {
		return err
	}

	err = c.saveToCache()
	if err != nil {
		return err
	}

	return nil

}

// Teardown operations for response, use with defer
func (r *Credentials) Teardown() error {
	return r.repo.Close()
}
