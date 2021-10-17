package config

import _ "embed"

const APP_NAME = "aws-mfa-credential-process"

//go:embed data/app-description-short.txt
var APP_DESCRIPTION_SHORT string

//go:embed data/app-description-long.txt
var APP_DESCRIPTION_LONG string

//go:embed data/assume-description-short.txt
var ASSUME_DESCRIPTION_SHORT string

//go:embed data/assume-description-long.txt
var ASSUME_DESCRIPTION_LONG string
