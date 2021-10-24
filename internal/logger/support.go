package logger

import "fmt"

const issueUrl string = "https://github.com/aripalo/vegas-credentials/issues/new"

func GetSupportString(padding string) string {
	return fmt.Sprintf("%sIf you believe this is an error with this tool, create a new issue:\n%s%s\n", padding, padding, issueUrl)
}
