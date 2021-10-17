package securestorage

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/99designs/keyring"
	"github.com/aripalo/aws-mfa-credential-process/internal/cachekey"
	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/aripalo/aws-mfa-credential-process/internal/prompt"
	"github.com/aripalo/aws-mfa-credential-process/internal/utils"
)

const KEYRING_LABEL string = config.PRODUCT_NAME
const KEYPREFIX string = KEYRING_LABEL + "__"

var ring keyring.Keyring

// Init initializes a keyring
func Init(disableDialog bool) {

	var err error

	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	keyringPath := filepath.Join(homedir, config.PRODUCT_CONFIG_LOCATION, "keyring")

	os.MkdirAll(keyringPath, os.ModePerm)

	utils.SafeLogLn(utils.FormatMessage(utils.COLOR_DEBUG, "ℹ️  ", "Cache Keyring", KEYRING_LABEL))

	ring, err = keyring.Open(keyring.Config{

		// Common config
		ServiceName: KEYRING_LABEL,
		AllowedBackends: []keyring.BackendType{
			keyring.KeychainBackend,
			keyring.WinCredBackend,
			keyring.KWalletBackend,
			keyring.SecretServiceBackend,
			keyring.PassBackend,
			keyring.FileBackend,
		},

		// MacOS keychain
		KeychainName:                   KEYRING_LABEL,
		KeychainTrustApplication:       true,
		KeychainSynchronizable:         false,
		KeychainAccessibleWhenUnlocked: false,

		// Windows
		WinCredPrefix: KEYRING_LABEL,

		// KDE Wallet
		KWalletAppID:  KEYRING_LABEL,
		KWalletFolder: KEYRING_LABEL,

		// freedesktop.org's Secret Service
		LibSecretCollectionName: KEYRING_LABEL,

		// Pass (https://www.passwordstore.org/)
		PassPrefix: KEYRING_LABEL,

		// Fallback encrypted file
		FileDir: keyringPath,
		FilePasswordFunc: func(message string) (string, error) {
			ctx, cancel := context.WithTimeout(nil, 5*time.Minute)
			defer cancel()

			if disableDialog {
				return prompt.Cli(ctx, message)
			} else {
				return prompt.Dialog(ctx, "Keyring Unlock", message)
			}
		},
	})

	if err != nil {
		panic(err)
	}
}

func Set(key string, data []byte) error {
	ensureRing()
	err := ring.Set(keyring.Item{
		Key:  key,
		Data: data,
	})
	return err
}

func Get(key string) ([]byte, error) {
	ensureRing()
	i, err := ring.Get(key)
	return i.Data, err
}

func Remove(key string) error {
	ensureRing()
	err := ring.Remove(key)
	return err
}

func RemoveAll(profileName string) error {
	ensureRing()
	var err error

	keys, err := ring.Keys()

	for _, v := range keys {
		if profileName == "" {
			err = Remove(v)
			return err
		} else {
			if strings.HasPrefix(v, fmt.Sprintf("%s%s", profileName, cachekey.Separator)) {
				err = Remove(v)
				return err
			}
		}
	}

	return err
}

func ensureRing() {
	if ring == nil {
		panic(errors.New("keyring not initilized"))
	}
}
