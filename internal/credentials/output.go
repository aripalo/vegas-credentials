package credentials

import "fmt"

// Output to stdout so aws credential_process can read it
// https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html
func (r *Credentials) Output() error {
	output, err := r.Serialize()
	if err != nil {
		return err
	}
	fmt.Fprint(r.output, string(output))
	return nil
}
