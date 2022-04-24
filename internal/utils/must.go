package utils

func Must[T any](operation func() (T, error)) T {
	result, err := operation()
	if err != nil {
		panic(err)
	}
	return result
}
