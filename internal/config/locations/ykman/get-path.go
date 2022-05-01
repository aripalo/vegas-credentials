package ykman

import "os/exec"

var executable string = "ykman"

func GetPath() string {
	if ykmanPath, err := exec.LookPath(executable); err != nil {
		return ""
	} else {
		return ykmanPath
	}
}
