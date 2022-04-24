package password

import (
	"context"
	"errors"
	"fmt"
	"vegas3/internal/msg"
	"vegas3/internal/yubikey/askpass"
)

type machineStateName string

type machineState struct {
	Name     machineStateName
	Password string
	count    int
	Error    error
}

const (
	INITIALIZE                        machineStateName = "INITIALIZE"
	GET_PASSWORD_FROM_CACHE                            = "GET_PASSWORD_FROM_CACHE"
	GET_PASSWORD_FROM_USER                             = "GET_PASSWORD_FROM_USER"
	PASSWORD_NOT_FOUND_FROM_CACHE                      = "PASSWORD_NOT_FOUND_FROM_CACHE"
	AUTHENTICATE_WITH_USER_PASSWORD                    = "AUTHENTICATE_WITH_USER_PASSWORD"
	AUTHENTICATE_WITH_CACHED_PASSWORD                  = "AUTHENTICATE_WITH_CACHED_PASSWORD"
	SAVE_TO_CACHE                                      = "SAVE_TO_CACHE"
	ERROR                                              = "ERROR"
	DONE                                               = "DONE"
)

type PasswordCache interface {
	Read() (string, error)
	Write(string) error
	Delete() error
}

const maximumPasswordTries int = 3

func Resolve(ctx context.Context, device string, cache PasswordCache, enableGui bool) error {

	// Loop through the state machine
	state := machineState{Name: INITIALIZE}
	for state.Name != DONE && state.Name != ERROR {
		state = stateMachine(ctx, state, device, cache, enableGui)
	}

	// Once done...
	if state.Error != nil {
		msg.Message.Warningln("⚠️", fmt.Sprintf("Yubikey Password: %s", state.Error))
		return state.Error
	}
	return nil
}

func stateMachine(ctx context.Context, state machineState, device string, cache PasswordCache, enableGui bool) machineState {

	switch state.Name {

	case GET_PASSWORD_FROM_CACHE:
		value, err := cache.Read()
		if err != nil {
			return machineState{
				Name: PASSWORD_NOT_FOUND_FROM_CACHE,
			}
		}

		msg.Message.Debugln("ℹ️", "Yubikey OATH Password: Received from Cache")
		return machineState{
			Name:     AUTHENTICATE_WITH_CACHED_PASSWORD,
			Password: value,
			count:    state.count + 1,
		}

	case PASSWORD_NOT_FOUND_FROM_CACHE:
		msg.Message.Debugln("ℹ️", "Yubikey OATH Password: Not found from Cache")
		err := cache.Delete()
		if err != nil {
			return machineState{
				Name:  ERROR,
				Error: err,
			}
		}
		return machineState{
			Name:  GET_PASSWORD_FROM_USER,
			count: state.count + 1,
		}

	case GET_PASSWORD_FROM_USER:
		value, err := askpass.AskPassword(ctx, enableGui)
		if err != nil {
			return machineState{
				Name:  ERROR,
				Error: err,
			}
		}
		if state.count >= maximumPasswordTries {
			return machineState{
				Name:  ERROR,
				Error: fmt.Errorf("Incorrect: Failed with too many attempts (%d)", state.count),
				count: state.count,
			}
		}
		return machineState{
			Name:     AUTHENTICATE_WITH_USER_PASSWORD,
			Password: value,
			count:    state.count + 1,
		}

	case AUTHENTICATE_WITH_CACHED_PASSWORD:
		ok, err := authenticate(ctx, device, state.Password)
		if err != nil {
			return machineState{
				Name:  ERROR,
				Error: err,
			}
		}
		if !ok {
			msg.Message.Warningln("⚠️", "Yubikey OATH Password: Incorrect")
			return machineState{
				Name:  GET_PASSWORD_FROM_USER,
				count: state.count,
			}
		}
		return machineState{
			Name:     DONE,
			Password: state.Password,
		}

	case AUTHENTICATE_WITH_USER_PASSWORD:
		ok, err := authenticate(ctx, device, state.Password)
		if err != nil {
			return machineState{
				Name:  ERROR,
				Error: err,
			}
		}
		if !ok {
			msg.Message.Warningln("⚠️", "Yubikey OATH Password: Incorrect")
			return machineState{
				Name:  GET_PASSWORD_FROM_USER,
				count: state.count,
			}
		}
		return machineState{
			Name:     SAVE_TO_CACHE,
			Password: state.Password,
		}

	case SAVE_TO_CACHE:

		err := cache.Write(state.Password)
		if err != nil {
			return machineState{
				Name:  ERROR,
				Error: err,
			}
		}

		msg.Message.Debugln("ℹ️", "Yubikey OATH Password: Saved to Cache")
		return machineState{
			Name:     DONE,
			Password: state.Password,
		}

	case INITIALIZE:
		return machineState{
			Name: GET_PASSWORD_FROM_CACHE,
		}

	default:
		return machineState{
			Name:  ERROR,
			Error: errors.New("uknown error"),
		}
	}
}
