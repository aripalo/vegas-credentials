package response

import (
	"fmt"
)

// Output to stdout so aws credential_process can read it
// https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html
func (r *Response) Output() error {
	output, err := r.Serialize()
	if err != nil {
		return err
	}
	fmt.Fprintf(r.destination, string(output))
	return nil
}
