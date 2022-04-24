package awsini

type SourceProfile struct {
	Name          string
	MfaSerial     string `ini:"mfa_serial"`
	YubikeySerial string `ini:"vegas_yubikey_serial"`
	YubikeyLabel  string `ini:"vegas_yubikey_label"`
	Region        string `ini:"region"`
}

type TargetProfile struct {
	Name            string
	SourceProfile   string `ini:"vegas_source_profile"`
	RoleArn         string `ini:"vegas_role_arn"`
	DurationSeconds int    `ini:"duration_seconds"`
	RoleSessionName string `ini:"role_session_name"`
	ExternalID      string `ini:"external_id"`
	Region          string `ini:"region"`
}
