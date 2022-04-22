package config

// Describes the path for a file used to control parallelism via file locks.
var MutexLockFile string = func(filename string) string {
	return TempFilePath(filename)
}("mutex-lock")
