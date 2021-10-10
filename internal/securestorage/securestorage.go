package securestorage

import (
	"github.com/99designs/keyring"
	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/aripalo/aws-mfa-credential-process/internal/utils"
)

const KEYRING_LABEL string = config.PRODUCT_NAME
const KEYPREFIX string = KEYRING_LABEL + "__"

func open() keyring.Keyring {

	utils.SafeLogLn(utils.FormatMessage(utils.COLOR_DEBUG, "ℹ️  ", "Cache Keyring", KEYRING_LABEL))

	ring, err := keyring.Open(keyring.Config{
		ServiceName:                    KEYRING_LABEL,
		KeychainSynchronizable:         false,
		KeychainAccessibleWhenUnlocked: false,
		KeychainName:                   KEYRING_LABEL,
	})

	if err != nil {
		panic(err)
	}

	return ring
}

func Set(key string, data []byte) error {
	ring := open()
	err := ring.Set(keyring.Item{
		Key:  KEYPREFIX + key,
		Data: data,
	})
	return err
}

func Get(key string) ([]byte, error) {
	ring := open()
	i, err := ring.Get(KEYPREFIX + key)
	return i.Data, err
}

func Remove(key string) error {
	ring := open()
	err := ring.Remove(KEYPREFIX + key)
	return err
}
