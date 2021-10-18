package response

import (
	"github.com/aripalo/aws-mfa-credential-process/internal/cache"
	"github.com/aripalo/aws-mfa-credential-process/internal/data"
	"github.com/aripalo/aws-mfa-credential-process/internal/logger"
)

// SaveToCache saves response to cache in secure keyring
func (r *Response) SaveToCache(d data.Provider) error {
	data, err := r.Serialize()
	if err != nil {
		return err
	}
	err = cache.Save(d, data)
	if err != nil {
		return err
	}

	logger.Debugln(d, "ℹ️ ", "Credentials", "Saved to cache")

	return nil
}

// ReadFromCache gets the cached response from secure keyring
func (r *Response) ReadFromCache(d data.Provider) error {
	data, err := cache.Get(d)
	if err != nil {
		return err
	}

	logger.DebugJSON(d, "🔧 ", "Cached Response", data)

	err = r.Deserialize(data)
	if err != nil {
		return err
	}

	return nil
}
