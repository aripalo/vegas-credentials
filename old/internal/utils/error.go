package utils

func PrintError(err error) {
	SafeLogLn()
	SafeLogLn(TextGrayDark(CreateRuler("=")))
	SafeLogLn(FormatMessage(COLOR_ERROR, "🚨 ", "credential_process", "Failed"))
	SafeLogLn(FormatMessage(COLOR_DEBUG, "ℹ️  ", "credential_process", "\n\n"+err.Error()))
	SafeLogLn(TextGrayDark(CreateRuler("=")))
}
