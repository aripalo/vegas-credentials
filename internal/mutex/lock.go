package mutex

import (
	"path"

	"github.com/alexflint/go-filemutex"
	"github.com/aripalo/vegas-credentials/internal/config/locations"
)

// Diretory to store the lock file into.
// By default uses locations.StateDir which is ensured to exists.
var dir string = locations.StateDir

// File name used to control the mutex lock.
var fileName string = "mutexlock"

// Return type fo the unlock function.
type MutexUnlock func() error

// Helps with parallel executions (e.g. with Terraform parallelism=n)
// ensuring only a single process at a time can interact with AWS and
// the internal cache – as the BadgerDB requires a filelock:
// https://github.com/dgraph-io/badger/blob/69926151f6532f2fe97a9b11ee9281519c8ec5e6/dir_unix.go#L45
func Lock() (MutexUnlock, error) {
	filePath := path.Join(dir, fileName)

	// create the filemutex
	fm, err := filemutex.New(filePath)
	if err != nil {
		return nil, err
	}

	err = fm.Lock()
	if err != nil {
		return nil, err
	}

	return fm.Unlock, nil

}
