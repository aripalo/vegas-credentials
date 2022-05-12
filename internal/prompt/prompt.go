package prompt

import (
	"bufio"
	"context"
	"os"
	"strings"
	"syscall"

	"github.com/ncruces/zenity"
	"golang.org/x/term"
)

func Password(ctx context.Context, title string, text string) (string, error) {
	_, value, err := zenity.Password(
		zenity.Title(title),
		zenity.Context(ctx),
	)

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(value), nil
}

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

func CliPassword(ctx context.Context, text string) (string, error) {

	value, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(value)), nil
}

func Cli(ctx context.Context, text string) (string, error) {

	reader := bufio.NewReader(os.Stdin)

	value, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(value), nil
}
