package cachekey

import (
	"fmt"
	"io"
	"testing"

	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
)

func TestGenerateSha1Hash(t *testing.T) {
	input := "foobar"

	// want generated with https://passwordsgenerator.net/sha1-hash-generator/
	want := "8843d7f92416211de9ebb963ff4ce28125932878"

	output := generateSha1Hash(input)
	if output != want {
		t.Fatalf(`generateSha1Hash("%s") = %q, want match for %#q`, input, output, want)
	}
}

func TestCombineStringsWithSimpleInput(t *testing.T) {
	input1 := "foo"
	input2 := "__"
	input3 := "bar"
	want := "foo__bar"
	output := combineStrings(input1, input2, input3)
	if output != want {
		t.Fatalf(`combineStrings("%s", "%s, "%s) = %q, want match for %#q`, input1, input2, input3, output, want)
	}
}

func TestCombineStringsWithRealInput(t *testing.T) {
	input1 := "my-profile"
	input2 := "__"
	input3 := "8843d7f92416211de9ebb963ff4ce28125932878"
	want := "my-profile__8843d7f92416211de9ebb963ff4ce28125932878"
	output := combineStrings(input1, input2, input3)
	fmt.Println("COMBINATION=====", output)
	if output != want {
		t.Fatalf(`combineStrings("%s", "%s, "%s) = %q, want match for %#q`, input1, input2, input3, output, want)
	}
}

func TestConfigToString(t *testing.T) {
	input := profile.Profile{
		RoleArn:       "arn:aws:iam::123456789012:role/ExampleRole",
		YubikeySerial: "123456",
		YubikeyLabel:  "foobar",
		MfaSerial:     "arn:aws:iam::123456789012:mfa/example",
		SourceProfile: "default",
	}

	want := `{"YubikeySerial":"123456","YubikeyLabel":"foobar","SourceProfile":"default","RoleArn":"arn:aws:iam::123456789012:role/ExampleRole","MfaSerial":"arn:aws:iam::123456789012:mfa/example","DurationSeconds":0,"Region":"","RoleSessionName":"","ExternalID":""}`

	output, err := configToString(input)

	if output != want || err != nil {
		t.Fatalf(`configToString(input) = %q, want match for %#q`, output, want)
	}
}

func TestGet(t *testing.T) {
	c := config.Config{
		Profile: "my-profile",
	}
	p := profile.Profile{
		RoleArn:       "arn:aws:iam::123456789012:role/ExampleRole",
		YubikeySerial: "123456",
		YubikeyLabel:  "foobar",
		MfaSerial:     "arn:aws:iam::123456789012:mfa/example",
		SourceProfile: "default",
	}

	// want generated with https://passwordsgenerator.net/sha1-hash-generator/
	// with data: {"YubikeySerial":"123456","YubikeyLabel":"foobar","SourceProfile":"default","RoleArn":"arn:aws:iam::123456789012:role/ExampleRole","MfaSerial":"arn:aws:iam::123456789012:mfa/example","DurationSeconds":0,"Region":"","RoleSessionName":"","ExternalID":""}
	want := "my-profile__3eb841cc0a378c607534bd21202ef4f9a721572a"

	foo := NewDpForTest(c, p)

	output, err := Get(foo)

	if output != want || err != nil {
		t.Fatalf(`configToString(input) = %q, want match for %#q`, output, want)
	}
}

type DpForTest struct {
	c config.Config
	p profile.Profile
	w io.Writer
}

func (d *DpForTest) GetWriteStream() io.Writer {
	return d.w
}

func (d *DpForTest) GetProfile() *profile.Profile {
	return &d.p
}

func (d *DpForTest) GetConfig() *config.Config {
	return &d.c
}

func NewDpForTest(c config.Config, p profile.Profile) *DpForTest {
	return &DpForTest{
		c: c,
		p: p,
		w: io.Discard,
	}
}
