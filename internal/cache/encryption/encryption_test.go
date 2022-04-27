package encryption

import (
	"math/rand"
	"testing"
)

func TestEncryption(t *testing.T) {

	inputs := []string{
		"foobar",
		randStringBytes(10),
		randStringBytes(1000),
		randStringBytes(10000),
		randStringBytes(100000),
		`{"Version":1,"AccessKeyId":"ASIA123456789EXAMPLE","SecretAccessKey":"noyieBae8oob4azeiquoo4al1mua1ooch7vieEXAMPLE","SessionToken":"EXAMPLEth3ohChoh0wo0ohfeiheichae/ahR0iejieBiesein1seno1aeRuZae4jeoreecheequooy0ahyohdaeraizu9Aeth3eeshoo6phaes1gahrofeeW/zahli0ooS9yeequ5aith6aikie2nahp3aoch2jo4vaebiab4aiChooshooNiefei6oow8ZohCheeyahXo8ee1ahm9GahgheeYae1Xo8/AhThae1ahpo3si7vaegee92khY2V0Vg+/f+RDAhThae1ahpo3si7vaegee9zsKLBjItJr0P+oaY/IGA4KFEBg9GmyAxGAhThae1ahpo3si7vaegee97xuJWrqcCnyFv","Expiration":"2019-05-20T23:49:23Z"}`,
	}

	for _, v := range inputs {
		encryptDecrypt(v, t)
	}
}

func encryptDecrypt(input string, t *testing.T) {
	encrypted, err := Encrypt([]byte(input))
	if err != nil {
		t.Fatalf("Encryption error not expected, got %q", err)
	}

	decrypted, err := Decrypt(encrypted)
	if err != nil {
		t.Fatalf("Decryption error not expected, got %q", err)
	}

	output := string(decrypted)

	if output != input {
		t.Fatalf("Output %q does not match input %q", output, input)
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
