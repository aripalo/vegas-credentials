package utils

import (
	"fmt"
	"log"
	"os"
)

// SafeLogger logs to stderr, since aws credential_process reads stdout
var SafeLogger *log.Logger = log.New(os.Stderr, "", 0)

// OutputToAwsCredentialProcess prints to stdout so aws credential_process can read it
func OutputToAwsCredentialProcess(output string) {
	fmt.Fprintf(os.Stdout, output)
}
