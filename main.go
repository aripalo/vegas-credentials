package main

import (
	"log"
	"os"
	"runtime/pprof"

	"github.com/aripalo/aws-mfa-credential-process/cmd"
)

func main() {

	cpuProfile := os.Getenv("CPU_PROFILE")

	if cpuProfile != "" {
		f, err := os.Create(cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		err = pprof.StartCPUProfile(f)
		if err != nil {
			panic(err)
		}
		defer pprof.StopCPUProfile()
	}
	cmd.Execute()
}
