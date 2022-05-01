package config

import (
	_ "embed"
)

//go:embed data/version-short.tmpl
var VersionShortTmpl string

//go:embed data/version-long.tmpl
var VersionLongTmpl string

// Software Version
var Version string = "v0.0.0-development"
