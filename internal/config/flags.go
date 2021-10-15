package config

type Flags struct {
	ProfileName    string
	Verbose        bool
	HideArns       bool
	DisableDialog  bool
	DisableRefresh bool
}

const (
	FLAG_PROFILE_NAME    string = "profile"
	FLAG_VERBOSE         string = "verbose"
	FLAG_HIDE_ARNS       string = "hide-arns"
	FLAG_DISABLE_DIALOG  string = "disable-dialog"
	FLAG_DISABLE_REFRESH string = "disable-refresh"
)
