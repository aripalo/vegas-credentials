package cache

import (
	"os"
	"path/filepath"
)

// CachePath provides the location for cache file
func CachePath(cacheName string, fileName string) string {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}

	return filepath.Join(cacheDir, cacheName, fileName)
}
