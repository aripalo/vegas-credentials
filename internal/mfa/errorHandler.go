package mfa

import "context"

func getTokenErrorHandler(ctx context.Context, err error, errors chan *error) {
	if err != nil {
		errors <- &err
		return
	}
}
