package awsini

// User models the ~/.aws/confg profile configuration for source profile.
// Essentially it matches IAM User.
type User struct {
	MfaSerial     string `ini:"mfa_serial"`
	YubikeySerial string `ini:"vegas_yubikey_serial"`
	YubikeyLabel  string `ini:"vegas_yubikey_label"`
	Region        string `ini:"region"`
}

// Role models the ~/.aws/confg profile configuration for target profile
// which is used to assume an IAM Role.
type Role struct {
	SourceProfile   string `ini:"vegas_source_profile"`
	RoleArn         string `ini:"vegas_role_arn"`
	DurationSeconds int    `ini:"duration_seconds"`
	RoleSessionName string `ini:"role_session_name"`
	ExternalID      string `ini:"external_id"`
	Region          string `ini:"region"`
}
