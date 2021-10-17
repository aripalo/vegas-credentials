package prompt

import (
	"bufio"
	"context"
	"os"
	"strings"

	"github.com/aripalo/aws-mfa-credential-process/internal/utils"
	"github.com/ncruces/zenity"
)

func Dialog(ctx context.Context, title string, text string) (string, error) {
	value, err := zenity.Entry(
		text,
		zenity.Title(title),
		zenity.Context(ctx),
	)

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(value), nil
}

func Cli(ctx context.Context, text string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	if text != "" {
		utils.SafeLogLn(text)
	}

	value, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(value), nil
}
