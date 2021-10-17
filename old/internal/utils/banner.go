package utils

import "github.com/aripalo/aws-mfa-credential-process/internal/config"

func PrintBanner(commandName string) {
	SafeLogLn(TextGrayDark(CreateRuler("=")))
	SafeLogLn()
	SafeLogLn(FormatMessage(COLOR_TITLE, "", config.PRODUCT_NAME, commandName))
	SafeLogLn()
	SafeLogLn(FormatMessage(COLOR_DEBUG, "📝 ", "Repository & Docs", "https://github.com/aripalo/aws-mfa-credential-process"))
	SafeLogLn()
	SafeLogLn(TextGrayDark(CreateRuler("-")))
}
