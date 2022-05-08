package resolver

import (
	"context"
	"strings"

	"github.com/aripalo/vegas-credentials/internal/multinput"
	"github.com/aripalo/vegas-credentials/internal/prompt"
)

var guiPrompt = prompt.Dialog

func GUI(ctx context.Context) (*multinput.Result, error) {
	result, err := guiPrompt(ctx, "Multifactor Authentication", "Enter TOTP MFA Token Code:")
	code := strings.TrimSpace(result)
	return &multinput.Result{Value: code, ResolverID: ResolverIdGuiDialog}, err
}
