package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"os/user"
	"strings"

	"github.com/aripalo/aws-mfa-credential-process/internal/utils"
	"github.com/shirou/gopsutil/host"
)

// Encrypt data with AES-256-CTR
func Encrypt(plaintext []byte) ([]byte, error) {

	var err error

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	blockCipher, err := createCipher()
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(blockCipher, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, err
}

// Decrypt AES-256-CTR encrypted data
func Decrypt(ciphertext []byte) ([]byte, error) {

	var err error
	plaintext := make([]byte, len(ciphertext[aes.BlockSize:]))

	// IV read from the beginning of the ciphertext
	iv := ciphertext[:aes.BlockSize]

	blockCipher, err := createCipher()
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(blockCipher, iv)
	stream.XORKeyStream(plaintext, ciphertext[aes.BlockSize:])

	return plaintext, err
}

// createCipher defines new cipherBlock used in AES encryption
func createCipher() (cipher.Block, error) {
	aesKey, err := getPassphrase()
	if err != nil {
		return nil, err
	}
	c, err := aes.NewCipher(aesKey)
	return c, err
}

// AES_256_KEYSIZE describes the length for key AES-256 encryption
const AES_256_KEYSIZE int = 32

// Generate the string value used in AES-256-CTR encryption secret.
// The getPassphrase is generated from the environment (boot time, hostname and user UID).
// If those values change (i.e. reboot or running in Docker container), it's okay,
// since it only means that the existing cache is ignored and new temporary credentials
// will be fetched from AWS STS.
func getPassphrase() ([]byte, error) {

	// Get system boot time, ignore error and use the default value (0) in that case
	bootedAt, _ := host.BootTime()
	bootedAtS := fmt.Sprint(bootedAt)

	// Resolve system hostname
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	// Resolve current user's UID
	user, err := user.Current()
	if err != nil {
		return nil, err
	}
	userUid := user.Uid

	// Join the resolved values
	var joined strings.Builder
	joined.WriteString(hostname)
	joined.WriteString(userUid)
	joined.WriteString(bootedAtS)

	// Create a SHA1 hash out of the joined strings
	passphrase, err := utils.GenerateSHA1(joined.String())
	if err != nil {
		return nil, err
	}

	// enforce length to 32 bytes for AES-256
	passphrase32 := passphrase[:AES_256_KEYSIZE]

	// Finally return the passphrase as byte array
	return []byte(passphrase32), nil
}
