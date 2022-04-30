package mutex

import (
	"path"

	"github.com/alexflint/go-filemutex"
	"github.com/aripalo/vegas-credentials/internal/config/locations"
	"github.com/aripalo/vegas-credentials/internal/logger"
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

	logger.Trace("Mutex lock initializing")

	filePath := path.Join(dir, fileName)

	// create the filemutex
	fm, err := filemutex.New(filePath)
	if err != nil {
		logger.Error("Mutex lock init failed")
		return nil, err
	}

	logger.Trace("Mutex lock init success")

	err = fm.Lock()
	if err != nil {
		logger.Error("Mutex lock acquiring failed")
		return nil, err
	}

	logger.Trace("Mutex lock acquiring success")

	return func() error {
		err := fm.Unlock()
		if err != nil {
			logger.Error("Mutex unlock failed")
			return err
		}
		logger.Trace("Mutex unlock success")
		return nil
	}, nil

}
