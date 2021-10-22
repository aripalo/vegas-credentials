package cache

import (
	"os"
	"path/filepath"
	"runtime"
)

// CachePath provides the location for cache file
func CachePath(cacheName string, fileName string) string {
	return cachePathForGOOS(cacheName, fileName, runtime.GOOS)
}

func cachePathForGOOS(cacheName string, fileName string, goos string) string {
	xdgCacheHome := os.Getenv("XDG_CACHE_HOME")
	if xdgCacheHome != "" {
		return filepath.Join(xdgCacheHome, cacheName, fileName)
	}
	if goos == "windows" {
		return filepath.Join("%LOCALAPPDATA%", cacheName, fileName)
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	if goos == "darwin" {
		return filepath.Join(homedir, "Library", "Caches", cacheName, fileName)
	}

	// TODO maybe check if this exists (without fileName and if not, use ~)
	return filepath.Join(homedir, ".cache", cacheName, fileName)
}
