package response

import (
	"time"

	"github.com/aripalo/aws-mfa-credential-process/internal/cache/cachekey"
	"github.com/aripalo/aws-mfa-credential-process/internal/data"
	"github.com/aripalo/aws-mfa-credential-process/internal/logger"
)

// SaveToCache saves response to cache in cache database
func (r *Response) SaveToCache(d data.Provider) error {
	data, err := r.Serialize()
	if err != nil {
		return err
	}

	key, err := cachekey.Get(d)
	if err != nil {
		return err
	}

	now := time.Now()
	ttl := r.Expiration.Sub(now)

	err = r.cache.Set(key, data, ttl) //TODO change exp to dur
	if err != nil {
		return err
	}

	logger.Debugln(d, "ℹ️ ", "Credentials", "Saved to cache")

	return nil
}

// ReadFromCache gets the cached response from cache database
func (r *Response) ReadFromCache(d data.Provider) error {
	key, err := cachekey.Get(d)
	if err != nil {
		return err
	}

	data, err := r.cache.Get(key)
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

// DeleteFromCache deletes the cached response cache database
func (r *Response) DeleteFromCache(d data.Provider) error {
	key, err := cachekey.Get(d)
	if err != nil {
		return err
	}

	return r.cache.Remove(key)
}
