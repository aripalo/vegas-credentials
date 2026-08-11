package boottime

import (
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/host"
)

// Return the system boot time or fallback to previous day 4AM.
func Get() time.Time {

	// Since Windows and WSL2 boot time can not be trusted (due to drift)
	// just use the previous day 4AM for now. It's okay for our cache
	// control purposes.
	if runtime.GOOS == "windows" || isWSL2() {
		return getPreviousDay4AM()
	}

	bootTime, err := getRealBootTime()
	if err != nil {
		return getPreviousDay4AM()
	}
	return bootTime
}

// Query system boot time with shirou/gopsutil.
func getRealBootTime() (time.Time, error) {
	var bootTime time.Time

	bootEpochSeconds, err := host.BootTime()
	if err != nil {
		return bootTime, err
	}

	return time.Unix(int64(bootEpochSeconds), 0), nil
}

// isWSL2 checks if the OS is running in Windows Subsystem for Linux 2.
func isWSL2() bool {
	data, err := os.ReadFile("/proc/version")
	if err != nil {
		return false
	}
	lower := strings.ToLower(string(data))
	return strings.Contains(lower, "microsoft") || strings.Contains(lower, "wsl")
}

// Return time for previous day 4AM.
func getPreviousDay4AM() time.Time {
	t := time.Now().AddDate(0, 0, -1)
	year, month, day := t.Date()
	return time.Date(year, month, day, 4, 0, 0, 0, t.Location())
}
