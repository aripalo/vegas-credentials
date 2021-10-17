package assume

import (
	"encoding/json"
	"fmt"
	"os"
)

// OutputToAwsCredentialProcess prints to stdout so aws credential_process can read it
// https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html
func outputToAwsCredentialProcess(output json.RawMessage) {
	fmt.Fprintf(os.Stdout, string(output))
}
