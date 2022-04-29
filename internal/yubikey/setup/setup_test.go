package setup

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStateMachine(t *testing.T) {
	tests := []struct {
		name     string
		input    State
		op       Operation
		expected State
	}{
		{
			name:     "default",
			input:    State{Name: 99},
			op:       Operation{},
			expected: State{Name: ERROR, Error: errors.New("unknown error")},
		},
		{
			name:     "init",
			input:    State{Name: INIT},
			op:       Operation{},
			expected: State{Name: CHECK_DEVICE_AVAILABLE},
		},
		{
			name:     "device not available",
			input:    State{Name: CHECK_DEVICE_AVAILABLE},
			op:       Operation{IsAvailable: func() bool { return false }},
			expected: State{Name: ERROR, Error: errors.New("yubikey: device not available")},
		},
		{
			name:     "device is available",
			input:    State{Name: CHECK_DEVICE_AVAILABLE},
			op:       Operation{IsAvailable: func() bool { return true }},
			expected: State{Name: CHECK_DEVICE_PASSWORD_PROTECTED},
		},
		{
			name:     "device not password protected",
			input:    State{Name: CHECK_DEVICE_PASSWORD_PROTECTED},
			op:       Operation{IsPasswordProtected: func() bool { return false }},
			expected: State{Name: CHECK_DEVICE_HAS_ACCOUNT},
		},
		{
			name:     "device is password protected",
			input:    State{Name: CHECK_DEVICE_PASSWORD_PROTECTED},
			op:       Operation{IsPasswordProtected: func() bool { return true }},
			expected: State{Name: GET_PASSWORD_FROM_CACHE},
		},
		{
			name:     "password not found from cache",
			input:    State{Name: GET_PASSWORD_FROM_CACHE},
			op:       Operation{GetPassword: func() (string, error) { return "", errors.New("key not found") }},
			expected: State{Name: PASSWORD_NOT_FOUND_FROM_CACHE},
		},
		{
			name:     "password is found from cache",
			input:    State{Name: GET_PASSWORD_FROM_CACHE},
			op:       Operation{GetPassword: func() (string, error) { return "p4ssword", nil }},
			expected: State{Name: AUTHENTICATE_WITH_CACHED_PASSWORD, Password: "p4ssword", Count: 1},
		},
		{
			name:     "handle error in password not found from cache",
			input:    State{Name: PASSWORD_NOT_FOUND_FROM_CACHE},
			op:       Operation{RemovePassword: func() error { return errors.New("fail") }},
			expected: State{Name: ERROR, Error: errors.New("fail")},
		},
		{
			name:     "handle password not found from cache",
			input:    State{Name: PASSWORD_NOT_FOUND_FROM_CACHE},
			op:       Operation{RemovePassword: func() error { return nil }},
			expected: State{Name: GET_PASSWORD_FROM_USER},
		},
		{
			name:     "authenticate error with cached password",
			input:    State{Name: AUTHENTICATE_WITH_CACHED_PASSWORD, Password: "p4ssword"},
			op:       Operation{Authenticate: func(string) (bool, error) { return false, errors.New("fail") }},
			expected: State{Name: ERROR, Error: errors.New("fail")},
		},
		{
			name:     "authenticate failed with cached password",
			input:    State{Name: AUTHENTICATE_WITH_CACHED_PASSWORD, Count: 1, Password: "p4ssword"},
			op:       Operation{Authenticate: func(string) (bool, error) { return false, nil }},
			expected: State{Name: GET_PASSWORD_FROM_USER, Count: 1},
		},
		{
			name:     "authenticate success with cached password",
			input:    State{Name: AUTHENTICATE_WITH_CACHED_PASSWORD, Count: 1, Password: "p4ssword"},
			op:       Operation{Authenticate: func(string) (bool, error) { return true, nil }},
			expected: State{Name: DONE, Password: "p4ssword"},
		},
		{
			name:     "get password from user too many retried",
			input:    State{Name: GET_PASSWORD_FROM_USER, Count: maximumPasswordTries + 1},
			op:       Operation{AskPass: func() (string, error) { return "p4ssword", nil }},
			expected: State{Name: ERROR, Count: 4, Error: errors.New("password is incorrect: failed with too many attempts (4)")},
		},
		{
			name:     "get password from user error",
			input:    State{Name: GET_PASSWORD_FROM_USER, Count: 1},
			op:       Operation{AskPass: func() (string, error) { return "p4ssword", errors.New("fail") }},
			expected: State{Name: ERROR, Error: errors.New("fail")},
		},
		{
			name:     "get password from user success",
			input:    State{Name: GET_PASSWORD_FROM_USER, Count: 1},
			op:       Operation{AskPass: func() (string, error) { return "p4ssword", nil }},
			expected: State{Name: AUTHENTICATE_WITH_USER_PASSWORD, Count: 2, Password: "p4ssword"},
		},
		{
			name:     "authenticate error with user password",
			input:    State{Name: AUTHENTICATE_WITH_USER_PASSWORD, Password: "p4ssword"},
			op:       Operation{Authenticate: func(string) (bool, error) { return false, errors.New("fail") }},
			expected: State{Name: ERROR, Error: errors.New("fail")},
		},
		{
			name:     "authenticate failed with user password",
			input:    State{Name: AUTHENTICATE_WITH_USER_PASSWORD, Count: 1, Password: "p4ssword"},
			op:       Operation{Authenticate: func(string) (bool, error) { return false, nil }},
			expected: State{Name: GET_PASSWORD_FROM_USER, Count: 1},
		},
		{
			name:     "authenticate success with user password",
			input:    State{Name: AUTHENTICATE_WITH_USER_PASSWORD, Count: 1, Password: "p4ssword"},
			op:       Operation{Authenticate: func(string) (bool, error) { return true, nil }},
			expected: State{Name: SAVE_PASSWORD, Password: "p4ssword"},
		},
		{
			name:     "save password error",
			input:    State{Name: SAVE_PASSWORD, Count: 1, Password: "p4ssword"},
			op:       Operation{SetPassword: func(string) error { return errors.New("fail") }},
			expected: State{Name: ERROR, Error: errors.New("fail")},
		},
		{
			name:     "save password success",
			input:    State{Name: SAVE_PASSWORD, Count: 1, Password: "p4ssword"},
			op:       Operation{SetPassword: func(string) error { return nil }},
			expected: State{Name: DONE, Password: "p4ssword"},
		},
		{
			name:     "check device has account error",
			input:    State{Name: CHECK_DEVICE_HAS_ACCOUNT},
			op:       Operation{HasAccount: func() (bool, error) { return false, errors.New("fail") }},
			expected: State{Name: ERROR, Error: errors.New("yubikey: could not read accounts")},
		},
		{
			name:     "check device has account fail",
			input:    State{Name: CHECK_DEVICE_HAS_ACCOUNT},
			op:       Operation{HasAccount: func() (bool, error) { return false, nil }},
			expected: State{Name: ERROR, Error: errors.New("yubikey: account not found")},
		},
		{
			name:     "check device has account success",
			input:    State{Name: CHECK_DEVICE_HAS_ACCOUNT},
			op:       Operation{HasAccount: func() (bool, error) { return true, nil }},
			expected: State{Name: DONE},
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			actual := stateMachine(test.input, test.op)
			assert.Equal(t, test.expected, actual)
		})
	}
}
