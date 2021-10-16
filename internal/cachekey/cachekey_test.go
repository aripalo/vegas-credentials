package cachekey

import (
	"fmt"
	"testing"

	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
)

func TestGenerateSha1Hash(t *testing.T) {
	input := "foobar"

	// want generated with https://passwordsgenerator.net/sha1-hash-generator/
	want := "8843d7f92416211de9ebb963ff4ce28125932878"

	output := generateSha1Hash(input)
	if output != want {
		t.Fatalf(`generateSha1Hash("%s") = %q, want match for %#q, nil`, input, output, want)
	}
}

func TestCombineStringsWithSimpleInput(t *testing.T) {
	input1 := "foo"
	input2 := "bar"
	want := "foobar"
	output := combineStrings(input1, input2)
	if output != want {
		t.Fatalf(`combineStrings("%s", "%s) = %q, want match for %#q, nil`, input1, input2, output, want)
	}
}

func TestCombineStringsWithRealInput(t *testing.T) {
	input1 := "my-profile"
	input2 := "{\"YubikeySerial\":\"123456\",\"YubikeyLabel\":\"foobar\",\"SourceProfile\":\"default\",\"AssumeRoleArn\":\"arn:aws:iam::123456789012:role/ExampleRole\",\"MfaSerial\":\"arn:aws:iam::123456789012:mfa/example\",\"DurationSeconds\":0,\"Region\":\"\",\"RoleSessionName\":\"\",\"ExternalID\":\"\"}"
	want := "my-profile{\"YubikeySerial\":\"123456\",\"YubikeyLabel\":\"foobar\",\"SourceProfile\":\"default\",\"AssumeRoleArn\":\"arn:aws:iam::123456789012:role/ExampleRole\",\"MfaSerial\":\"arn:aws:iam::123456789012:mfa/example\",\"DurationSeconds\":0,\"Region\":\"\",\"RoleSessionName\":\"\",\"ExternalID\":\"\"}"
	output := combineStrings(input1, input2)
	fmt.Println("COMBINATION=====", output)
	if output != want {
		t.Fatalf(`combineStrings("%s", "%s) = %q, want match for %#q, nil`, input1, input2, output, want)
	}
}

func TestConfigToString(t *testing.T) {
	input := profile.Profile{
		AssumeRoleArn: "arn:aws:iam::123456789012:role/ExampleRole",
		YubikeySerial: "123456",
		YubikeyLabel:  "foobar",
		MfaSerial:     "arn:aws:iam::123456789012:mfa/example",
		SourceProfile: "default",
	}

	want := "{\"YubikeySerial\":\"123456\",\"YubikeyLabel\":\"foobar\",\"SourceProfile\":\"default\",\"AssumeRoleArn\":\"arn:aws:iam::123456789012:role/ExampleRole\",\"MfaSerial\":\"arn:aws:iam::123456789012:mfa/example\",\"DurationSeconds\":0,\"Region\":\"\",\"RoleSessionName\":\"\",\"ExternalID\":\"\"}"

	output, err := configToString(input)

	if output != want || err != nil {
		t.Fatalf(`configToString(input) = %q, want match for %#q, nil`, output, want)
	}
}

func TestGet(t *testing.T) {
	input1 := "my-profile"
	input2 := profile.Profile{
		AssumeRoleArn: "arn:aws:iam::123456789012:role/ExampleRole",
		YubikeySerial: "123456",
		YubikeyLabel:  "foobar",
		MfaSerial:     "arn:aws:iam::123456789012:mfa/example",
		SourceProfile: "default",
	}

	// want generated with https://passwordsgenerator.net/sha1-hash-generator/
	// with data: my-profile{"YubikeySerial":"123456","YubikeyLabel":"foobar","SourceProfile":"default","AssumeRoleArn":"arn:aws:iam::123456789012:role/ExampleRole","MfaSerial":"arn:aws:iam::123456789012:mfa/example","DurationSeconds":0,"Region":"","RoleSessionName":"","ExternalID":""}
	want := "cf1635f6d6e5e5a54dc3439924868c6a440f4a21"

	output, err := Get(input1, input2)

	if output != want || err != nil {
		t.Fatalf(`configToString(input) = %q, want match for %#q, nil`, output, want)
	}
}
