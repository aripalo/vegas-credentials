package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"strings"
)

func Encrypt(plaintext []byte) ([]byte, error) {

	var err error
	//var encrypted []byte

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	fmt.Println("IV length: ", len(iv))

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

	passphrase := "password" // TODO

	encKey := strings.ReplaceAll(fmt.Sprintf("%-64x\n", passphrase), " ", "0")[:AES_256_KEYSIZE]

	fmt.Println("SECRET KEY ", encKey)

	return []byte(encKey), nil
}

func createCipher() (cipher.Block, error) {
	aesKey, err := getSecretKey()
	c, err := aes.NewCipher(aesKey)
	return c, err
}

const AES_256_KEYSIZE int = 32
