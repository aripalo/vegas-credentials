package locations

import "os/exec"

func getYkmanPath() string {
	if ykmanPath, err := exec.LookPath("ykman"); err != nil {
		return ""
	} else {
		return ykmanPath
	}
}
