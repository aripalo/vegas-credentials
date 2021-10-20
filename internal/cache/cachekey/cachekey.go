package cachekey

import (
	"encoding/json"
	"strings"

	"github.com/aripalo/aws-mfa-credential-process/internal/data"
	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
	"github.com/aripalo/aws-mfa-credential-process/internal/utils"
)

const Separator = "__"

// Get is responsible for creating a unique cache key for given profile configuration, therefore ensuring mutated profile configuration will not use previous cached data
func Get(d data.Provider) (string, error) {
	c := d.GetConfig()
	p := d.GetProfile()
	configString, err := configToString(*p)
	hash := utils.GenerateSHA1(configString)
	key := strings.Join([]string{c.Profile, Separator, hash}, "")
	return key, err
}

// configToString convertts profile config into stringified JSON
func configToString(p profile.Profile) (string, error) {
	result, err := json.Marshal(p)
	return string(result), err
}
