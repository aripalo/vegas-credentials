package mfa

import (
	"bufio"
	"context"
	"os"
	"strings"
)

func getCliToken(ctx context.Context, out chan *Result, errors chan *error) {

	var err error
	var result Result
	result.Provider = TOKEN_PROVIDER_CLI

	defer getTokenErrorHandler(ctx, err, errors)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	text = strings.TrimSpace(text)

	result.Value = strings.TrimSpace(text)
	out <- &result
}
