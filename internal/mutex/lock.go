package mutex

import (
	"errors"
	"os"
	"path"

	"github.com/alexflint/go-filemutex"
)

type MutexControl struct {
	dir  string
	file string
	path string

	// Will block until lock can be acquired
	Lock func() error

	// Releases the lock
	Unlock func() error
}

// Helps with parallel executions (e.g. with Terraform parallelism=n)
// ensuring only a single process at a time can interact with AWS and
// the internal cache – as the BadgerDB requires a filelock:
// https://github.com/dgraph-io/badger/blob/69926151f6532f2fe97a9b11ee9281519c8ec5e6/dir_unix.go#L45
func New(dir string) (MutexControl, error) {
	file := "mutexlock"
	m := MutexControl{
		dir:  dir,
		file: file,
		path: path.Join(dir, file),
	}

	// the base directory must exists before hand so use MkdirAll
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return m, errors.New("Could not create cache lock-control file")
	}

	// create the filemutex
	fm, err := filemutex.New(m.path)
	if err != nil {
		return m, errors.New("Directory did not exist or file could not created")
	}

	m.Lock = fm.Lock
	m.Unlock = fm.Unlock

	return m, nil
}
