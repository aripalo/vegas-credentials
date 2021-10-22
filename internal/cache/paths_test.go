package cache

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPathXDG(t *testing.T) {
	base := "/xdg-cache-home"
	os.Setenv("XDG_CACHE_HOME", base)

	cacheName := "foo"
	fileName := "bar"
	p := cachePathForGOOS(cacheName, fileName, "darwin")

	os.Setenv("XDG_CACHE_HOME", "")

	want := filepath.Join(base, cacheName, fileName)

	if p != want {
		t.Fatalf("Got %q, want %q", p, want)
	}
}

func TestPathWindows(t *testing.T) {
	base := "%LOCALAPPDATA%"
	cacheName := "foo"
	fileName := "bar"
	p := cachePathForGOOS(cacheName, fileName, "windows")

	want := filepath.Join(base, cacheName, fileName)

	if p != want {
		t.Fatalf("Got %q, want %q", p, want)
	}
}

func TestPathMacOS(t *testing.T) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	base := filepath.Join(homedir, "Library", "Caches")
	cacheName := "foo"
	fileName := "bar"
	p := cachePathForGOOS(cacheName, fileName, "darwin")

	want := filepath.Join(base, cacheName, fileName)

	if p != want {
		t.Fatalf("Got %q, want %q", p, want)
	}
}

func TestPathLinux(t *testing.T) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	base := filepath.Join(homedir, ".cache")
	cacheName := "foo"
	fileName := "bar"
	p := cachePathForGOOS(cacheName, fileName, "linux")

	want := filepath.Join(base, cacheName, fileName)

	if p != want {
		t.Fatalf("Got %q, want %q", p, want)
	}
}
