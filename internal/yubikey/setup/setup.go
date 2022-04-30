package setup

import (
	"context"
	"errors"
	"fmt"

	"github.com/aripalo/vegas-credentials/internal/msg"
	"github.com/aripalo/vegas-credentials/internal/yubikey/askpass"
	"github.com/aripalo/ykmangoath"
)

type PasswordStore interface {
	SetPassword(password string) error
	GetPassword() (string, error)
	RemovePassword() error
}

// Options passed in by the caller
type Options struct {
	Device    string
	Account   string
	EnableGui bool
}

type Operation struct {
	IsAvailable         func() bool
	IsPasswordProtected func() bool
	HasAccount          func() (bool, error)
	Authenticate        func(string) (bool, error)
	AskPass             func() (string, error)
	SetPassword         func(string) error
	GetPassword         func() (string, error)
	RemovePassword      func() error
}

func Setup(options Options, store PasswordStore) error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	oathAccounts, err := ykmangoath.New(ctx, options.Device)
	if err != nil {
		return fmt.Errorf("ykmangoat init: %w", err)
	}

	op := Operation{
		IsAvailable:         oathAccounts.IsAvailable,
		IsPasswordProtected: oathAccounts.IsPasswordProtected,
		HasAccount: func() (bool, error) {
			return oathAccounts.HasAccount(options.Account)
		},
		Authenticate: func(password string) (bool, error) {
			msg.Debug("⚠️", "AUTHENTICAT RECEIVED: "+password)
			return authenticate(oathAccounts, password)
		},
		AskPass: func() (string, error) {
			return askpass.AskPassword(ctx, options.EnableGui)
		},
		SetPassword:    store.SetPassword,
		GetPassword:    store.GetPassword,
		RemovePassword: store.RemovePassword,
	}

	result := run(op)
	// Once done...
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func run(op Operation) State {

	// Loop through the state machine
	state := State{Name: INIT}
	for state.Name != DONE && state.Name != ERROR {
		state = stateMachine(state, op)
	}

	return state

}

type StateName uint8

type State struct {
	Name     StateName
	Password string
	Count    int
	Error    error
}

const (
	INIT StateName = iota
	CHECK_DEVICE_AVAILABLE
	CHECK_DEVICE_PASSWORD_PROTECTED
	GET_PASSWORD_FROM_CACHE
	PASSWORD_NOT_FOUND_FROM_CACHE
	AUTHENTICATE_WITH_CACHED_PASSWORD
	GET_PASSWORD_FROM_USER
	AUTHENTICATE_WITH_USER_PASSWORD
	SAVE_PASSWORD
	CHECK_DEVICE_HAS_ACCOUNT
	ERROR
	DONE
)

const maximumPasswordTries int = 3

func stateMachine(state State, op Operation) State {

	switch state.Name {

	case INIT:
		return State{Name: CHECK_DEVICE_AVAILABLE}

	case CHECK_DEVICE_AVAILABLE:
		if !op.IsAvailable() {
			msg.Debug("⚠️", "Yubikey: Device not available")
			return State{
				Name:  ERROR,
				Error: errors.New("yubikey: device not available"),
			}
		}
		return State{
			Name: CHECK_DEVICE_PASSWORD_PROTECTED,
		}

	case CHECK_DEVICE_PASSWORD_PROTECTED:
		if op.IsPasswordProtected() {
			msg.Debug("⚠️", "Yubikey: Device is password protected")
			return State{
				Name: GET_PASSWORD_FROM_CACHE,
			}
		}
		return State{
			Name: CHECK_DEVICE_HAS_ACCOUNT,
		}

	case GET_PASSWORD_FROM_CACHE:
		password, err := op.GetPassword()
		if err != nil {
			msg.Debug("⚠️", "Yubikey: Password not found from cache")
			return State{
				Name: PASSWORD_NOT_FOUND_FROM_CACHE,
			}
		}
		msg.Debug("⚠️", "Yubikey: Password found from cache")
		return State{
			Name:     AUTHENTICATE_WITH_CACHED_PASSWORD,
			Password: password,
			Count:    1,
		}

	case PASSWORD_NOT_FOUND_FROM_CACHE:
		err := op.RemovePassword()
		if err != nil {
			return State{
				Name:  ERROR,
				Error: err,
			}
		}
		return State{
			Name: GET_PASSWORD_FROM_USER,
		}

	case AUTHENTICATE_WITH_CACHED_PASSWORD:
		ok, err := op.Authenticate(state.Password)
		if err != nil {
			msg.Warn("⚠️", "Yubikey: Error authentication with cached password")
			return State{
				Name:  ERROR,
				Error: err,
			}
		}
		if !ok {
			msg.Warn("⚠️", "Yubikey: Incorrect password from cache")
			return State{
				Name:  GET_PASSWORD_FROM_USER,
				Count: state.Count,
			}
		}
		return State{
			Name:     DONE,
			Password: state.Password,
		}

	case GET_PASSWORD_FROM_USER:
		if state.Count >= maximumPasswordTries {
			return State{
				Name:  ERROR,
				Error: fmt.Errorf("password is incorrect: failed with too many attempts (%d)", state.Count),
				Count: state.Count,
			}
		}
		value, err := op.AskPass()

		msg.Debug("⚠️", "Yubikey OATH Password: "+value)

		if err != nil {
			return State{
				Name:  ERROR,
				Error: err,
			}
		}
		return State{
			Name:     AUTHENTICATE_WITH_USER_PASSWORD,
			Password: value,
			Count:    state.Count + 1,
		}

	case AUTHENTICATE_WITH_USER_PASSWORD:
		ok, err := op.Authenticate(state.Password)
		if err != nil {
			msg.Warn("⚠️", "Yubikey OATH Password: Authentication Error")
			return State{
				Name:  ERROR,
				Error: err,
			}
		}
		if !ok {
			msg.Warn("⚠️", "Yubikey OATH Password: Incorrect")
			return State{
				Name:  GET_PASSWORD_FROM_USER,
				Count: state.Count,
			}
		}
		return State{
			Name:     SAVE_PASSWORD,
			Password: state.Password,
		}

	case SAVE_PASSWORD:
		err := op.SetPassword(state.Password)
		if err != nil {
			return State{
				Name:  ERROR,
				Error: err,
			}
		}

		//msg.Debug("ℹ️", "Yubikey OATH Password: Saved to Cache")
		return State{
			Name:     DONE,
			Password: state.Password,
		}

	case CHECK_DEVICE_HAS_ACCOUNT:
		has, err := op.HasAccount()
		if err != nil {
			return State{
				Name:  ERROR,
				Error: errors.New("yubikey: could not read accounts"),
			}
		}
		if !has {
			return State{
				Name:  ERROR,
				Error: errors.New("yubikey: account not found"),
			}
		}
		return State{
			Name: DONE,
		}

	default:
		return State{
			Name:  ERROR,
			Error: errors.New("unknown error"),
		}
	}
}

func authenticate(oathAccounts ykmangoath.OathAccounts, password string) (bool, error) {
	if password != "" {
		err := oathAccounts.SetPassword(password)
		if err != nil {
			return false, err
		}
	}

	_, err := oathAccounts.List()

	msg.Debug("⚠️", fmt.Sprintf("List Error: %v", err))

	if err != nil {
		return false, err
	}

	return true, nil
}
