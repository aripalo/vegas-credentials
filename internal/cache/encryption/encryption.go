package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"strings"
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

	passphrase, err := password()
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

func password() (string, error) {
	tmp := "password" // TODO
	return strings.TrimSpace(string(tmp)), nil
}
