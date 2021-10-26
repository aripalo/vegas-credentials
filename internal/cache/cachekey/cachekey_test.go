package cachekey

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/profile"
	"github.com/aripalo/vegas-credentials/internal/profile/source"
	"github.com/aripalo/vegas-credentials/internal/profile/target"
	"github.com/aripalo/vegas-credentials/internal/vegastestapp"
)

func TestCombineStringsWithSimpleInput(t *testing.T) {
	input1 := "foo"
	input2 := "__"
	input3 := "bar"
	want := "foo__bar"
	output := strings.Join([]string{input1, input2, input3}, "")
	if output != want {
		t.Fatalf(`combineStrings("%s", "%s, "%s) = %q, want match for %#q`, input1, input2, input3, output, want)
	}
}

func TestCombineStringsWithRealInput(t *testing.T) {
	input1 := "my-profile"
	input2 := "__"
	input3 := "8843d7f92416211de9ebb963ff4ce28125932878"
	want := "my-profile__8843d7f92416211de9ebb963ff4ce28125932878"
	output := strings.Join([]string{input1, input2, input3}, "")
	fmt.Println("COMBINATION=====", output)
	if output != want {
		t.Fatalf(`combineStrings("%s", "%s, "%s) = %q, want match for %#q`, input1, input2, input3, output, want)
	}
}

func TestConfigToString(t *testing.T) {
	/*
		input := profile.NewProfile{
			RoleArn:       "arn:aws:iam::123456789012:role/ExampleRole",
			YubikeySerial: "123456",
			YubikeyLabel:  "foobar",
			MfaSerial:     "arn:aws:iam::123456789012:mfa/example",
			SourceProfile: "default",
		}
	*/
	input := profile.Profile{
		Source: &source.SourceProfile{
			YubikeySerial: "123456",
			YubikeyLabel:  "foobar",
			MfaSerial:     "arn:aws:iam::123456789012:mfa/example",
		},
		Target: &target.TargetProfile{
			SourceProfile: "default",
			RoleArn:       "arn:aws:iam::123456789012:role/ExampleRole",
		},
	}

	want := `{"Source":{"Name":"","YubikeySerial":"123456","YubikeyLabel":"foobar","MfaSerial":"arn:aws:iam::123456789012:mfa/example","Region":""},"Target":{"SourceProfile":"default","RoleArn":"arn:aws:iam::123456789012:role/ExampleRole","DurationSeconds":0,"Region":"","RoleSessionName":"","ExternalID":""}}`

	output, err := configToString(input)

	if output != want || err != nil {
		t.Fatalf(`configToString(input) = %q, want match for %#q`, output, want)
	}
}

func TestGet(t *testing.T) {
	f := config.Flags{
		Profile: "my-profile",
	}
	p := profile.Profile{
		Source: &source.SourceProfile{
			YubikeySerial: "123456",
			YubikeyLabel:  "foobar",
			MfaSerial:     "arn:aws:iam::123456789012:mfa/example",
		},
		Target: &target.TargetProfile{
			SourceProfile: "default",
			RoleArn:       "arn:aws:iam::123456789012:role/ExampleRole",
		},
	}

	// want generated with https://passwordsgenerator.net/sha1-hash-generator/
	// with data: {"YubikeySerial":"123456","YubikeyLabel":"foobar","SourceProfile":"default","RoleArn":"arn:aws:iam::123456789012:role/ExampleRole","MfaSerial":"arn:aws:iam::123456789012:mfa/example","DurationSeconds":0,"Region":"","RoleSessionName":"","ExternalID":""}
	want := "my-profile__18b65be949ff29f60fb833e1326f095019115293"

	foo := vegastestapp.New(f, p)

	output, err := Get(foo)

	if output != want || err != nil {
		t.Fatalf(`configToString(input) = %q, want match for %#q`, output, want)
	}
}
