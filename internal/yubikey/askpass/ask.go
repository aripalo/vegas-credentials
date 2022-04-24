package askpass

import (
	"context"
	"vegas3/internal/msg"
	"vegas3/internal/multinput"
	"vegas3/internal/prompt"
)

// Assign CLI Prompt to variable (useful for testing).
var cliPrompt = prompt.Cli

// Assign GUI Prompt to variable (useful for testing).
var guiPrompt = prompt.Password

// Query Yubikey OATH application password via CLI stdin.
func passwordQueryCLI(ctx context.Context) (*multinput.Result, error) {
	value, err := cliPrompt(ctx, "")
	return &multinput.Result{Value: value, ResolverID: "CLI"}, err
}

// Query Yubikey OATH application password via GUI Password diaglog.
func passwordQueryGUI(ctx context.Context) (*multinput.Result, error) {
	value, err := guiPrompt(ctx, "Yubikey OATH Password", "Password:")
	return &multinput.Result{Value: value, ResolverID: "GUI"}, err
}

// Ask the Yubikey OATH application password from user.
func AskPassword(ctx context.Context, enableGui bool) (string, error) {
	// resolvers used with multinput to query Yubikey password from user
	var resolvers []multinput.InputResolver
	resolvers = append(resolvers, passwordQueryCLI)
	if enableGui {
		resolvers = append(resolvers, passwordQueryGUI)
	}

	msg.Message.Promptln("🔑", "Input Yubikey OATH password: ")

	// Assign multinput query with given resolvers
	mi := multinput.New(resolvers)
	result, err := mi.Provide(ctx)

	if err != nil {
		return "", err
	}

	return result.Value, nil
}
