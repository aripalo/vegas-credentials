package response

import (
	"time"

	"github.com/aripalo/vegas-credentials/internal/cache/cachekey"
	"github.com/aripalo/vegas-credentials/internal/interfaces"
	"github.com/aripalo/vegas-credentials/internal/logger"
)

// SaveToCache saves response to cache in cache database
func (r *Response) SaveToCache(a interfaces.AssumeCredentialProcess) error {
	data, err := r.Serialize()
	if err != nil {
		return err
	}

	key, err := cachekey.Get(a)
	if err != nil {
		return err
	}

	now := time.Now()
	ttl := r.Expiration.Sub(now)

	err = r.cache.Set(key, data, ttl)
	if err != nil {
		return err
	}

	logger.Debugln(a, "ℹ️ ", "Credentials", "Saved to cache")

	return nil
}

// ReadFromCache gets the cached response from cache database
func (r *Response) ReadFromCache(a interfaces.AssumeCredentialProcess) error {
	key, err := cachekey.Get(a)
	if err != nil {
		return err
	}

	data, err := r.cache.Get(key)
	if err != nil {
		return err
	}

	logger.DebugJSON(a, "🔧 ", "Cached Response", data)

	err = r.Deserialize(data)
	if err != nil {
		return err
	}

	return nil
}

// DeleteFromCache deletes the cached response cache database
func (r *Response) DeleteFromCache(a interfaces.AssumeCredentialProcess) error {
	key, err := cachekey.Get(a)
	if err != nil {
		return err
	}

	return r.cache.Remove(key)
}
