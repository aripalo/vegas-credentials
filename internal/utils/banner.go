package utils

func PrintBanner() {
	SafeLogLn(TextGrayDark(CreateRuler("=")))
	SafeLogLn()
	SafeLogLn(FormatMessage(COLOR_TITLE, "", "credential_process", "AWS MFA Assume Credential Process"))
	SafeLogLn()
	SafeLogLn(FormatMessage(COLOR_DEBUG, "📝 ", "Repository & Docs", "https://github.com/aripalo/aws-mfa-assume-credential-process"))
	SafeLogLn()
	SafeLogLn(TextGrayDark(CreateRuler("-")))
}
