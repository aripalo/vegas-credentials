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

	"github.com/aripalo/vegas-credentials/internal/cache/encryption/boottime"
	"github.com/aripalo/vegas-credentials/internal/checksum"
	"github.com/aripalo/vegas-credentials/internal/msg"
)

// Encrypt data with AES-256-CTR
func Encrypt(plaintext []byte) ([]byte, error) {

	var err error

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	iv := ciphertext[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
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

	bootTime := getTimeStamp()

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

	// Control value invalidates the cache if:
	// - machine hostname changed
	// - user UID changed
	// - system has been restarted
	control := buildControlValue(hostname, user.Uid, bootTime)

	// Print out control value for debugging purposes, but don't write it to log.
	msg.DebugNoLog("ℹ️", fmt.Sprintf("Control Value: %s", control))

	// Create a SHA1 hash out of the joined strings
	passphrase, err := checksum.Generate([]byte(control))
	if err != nil {
		return nil, err
	}

	// enforce length to 32 bytes for AES-256
	passphrase32 := passphrase[:AES_256_KEYSIZE]

	// Finally return the passphrase as byte array
	return []byte(passphrase32), nil
}

// Time formatting layout for cache control value.
// Essentially: year+month+day+hour+minute.
const bootedAtTimestampFormat string = "200601021504"

// Get system boot time minute (or fallback to previous day 4AM).
// See package "boottime" for more.
func getTimeStamp() string {
	bootedAt := boottime.Get()
	return bootedAt.Format(bootedAtTimestampFormat)
}

// Build cache control value from hostname, user UID and boot time.
func buildControlValue(hostname string, userUid string, bootTime string) string {
	var joined strings.Builder
	joined.WriteString(hostname)
	joined.WriteString("__")
	joined.WriteString(userUid)
	joined.WriteString("__")
	joined.WriteString(bootTime)
	return joined.String()
}
