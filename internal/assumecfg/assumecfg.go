// Package assumecfg implements loading & providing necessary configuration data
// required to perform AWS STS AssumeRole operations with MFA.
package assumecfg

// Configuration data required to perform AWS STS AssumeRole operation with MFA.
type AssumeCfg struct {
	// ProfileName of the profile in AWS Configuration file.
	//  "frank@concerts"
	ProfileName string

	// MfaSerial can be either a serial number for a hardware device
	// (such as GAHT12345678) or an Amazon Resource Name (ARN) for
	// a virtual MFA device (such as arn:aws:iam::123456789012:mfa/user).
	MfaSerial string

	// Yubikey Device Serial ID Number (8+ digits).
	//  "12345678"
	YubikeySerial string

	// YubikeyLabel is the Account Label in Yubikey OATH application.
	//  "<issuer>:<name>"
	//  "Amazon Web Services:frank@concerts"
	// It can also match the MfaSerial ARN format:
	//  "arn:aws:iam::123456789012:mfa/user"
	YubikeyLabel string

	// AWS Region to be used when interacting with AWS Security Token Service (STS).
	//  "eu-north-1"
	Region string

	// The source profile fomr AWS config file with long-term credentials.
	//  "default"
	SourceProfile string

	// Role ARN to assume.
	//  "arn:aws:iam::222222222222:role/SingerRole"
	RoleArn string

	// Duration Seconds
	//  4383
	DurationSeconds int

	// Role Session Name
	//  "SinatraAtTheSands"
	RoleSessionName string

	// External ID
	//  "0093624694724"
	ExternalID string

	// SHA1 hash checksum calculated from the current configuration.
	// Useful for detecting configuration changes e.g. for caching purposes.
	Checksum string
}
