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

// See https://pkg.go.dev/crypto/cipher#NewCTR

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

func getSecretKey() ([]byte, error) {

	passphrase, err := getPassphrase()
	if err != nil {
		return []byte(passphrase), err
	}

	encKey := strings.ReplaceAll(fmt.Sprintf("%-64x\n", passphrase), " ", "0")[:AES_256_KEYSIZE]

	return []byte(encKey), nil
}

func createCipher() (cipher.Block, error) {
	aesKey, err := getSecretKey()
	if err != nil {
		return nil, err
	}
	c, err := aes.NewCipher(aesKey)
	return c, err
}

const AES_256_KEYSIZE int = 32

// Generate the string value used in AES-256-CTR encryption secret.
// The getPassphrase is generated from the environment (boot time, hostname and user UID).
// If those values change (i.e. reboot or running in Docker container), it's okay,
// since it only means that the existing cache is ignored and new temporary credentials
// will be fetched from AWS STS.
func getPassphrase() (string, error) {

	// Get system boot time, ignore error and use the default value (0) in that case
	bootedAt, _ := host.BootTime()
	bootedAtS := fmt.Sprint(bootedAt)

	// Resolve system hostname
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}

	// Resolve current user's UID
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	userUid := user.Uid

	// Join the resolved values, create a SHA1 out of them and return it
	joined := strings.Join([]string{hostname, userUid, bootedAtS}, "")
	pwd := utils.GenerateSHA1(joined)
	return pwd, nil
}
