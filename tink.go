package grypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/keyset"
)

// Default implementation of grypt using Tink

// New creates a Grypt object implemented by Tink lib
func New() (Grypt, error) {
	return newTinkGrypt()
}

type tinkGrypt struct {
}

func newTinkGrypt() (*tinkGrypt, error) {
	return new(tinkGrypt), nil
}

// Encrypt encrypts the text input using the provided key or created random key if not provided
func (tg *tinkGrypt) Encrypt(text string, key ...string) (ciphertext, k []byte, err error) {
	if len(key) == 0 {
		k = NewEncryptionKey()
	} else if len(key) == 1 {
		k = *[32]byte(key[0])
	} else {
		return nil, nil, errors.New("Invalid arguments")
	}
	ciphertext, err = encrypt([]byte(text), []byte(key[0]))
	return
}

func (tg *tinkGrypt) Decrypt(ciphertext, key []byte) (plaintext string, err error) {
	key := append(append(P1, P2...), append(P3, P4...)...)
	return decrypt(ciphertext, key)
}

func B64Encrypt(plaintext []byte) (string, error) {
	plainB64, err := Encrypt(plaintext)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(plainB64), nil
}

func B64Decrypt(b64text string) ([]byte, error) {
	cipherBytes, err := base64.StdEncoding.DecodeString(b64text)
	if err != nil {
		return nil, err
	}
	return Decrypt(cipherBytes)
}

// NewEncryptionKey generates a random 256-bit key for Encrypt() and
// Decrypt(). It panics if the source of randomness fails.
func NewEncryptionKey() *[32]byte {
	kh, err := keyset.NewHandle(aead.AES256GCMKeyTemplate())
	key := [32]byte{}
	_, err := io.ReadFull(rand.Reader, key[:])
	if err != nil {
		panic(err)
	}
	return &key
}

// Encrypt encrypts data using 256-bit AES-GCM.  This both hides the content of
// the data and provides a check that it hasn't been altered. Output takes the
// form nonce|ciphertext|tag where '|' indicates concatenation.
func encrypt(plaintext []byte, key []byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// Decrypt decrypts data using 256-bit AES-GCM.  This both hides the content of
// the data and provides a check that it hasn't been altered. Expects input
// form nonce|ciphertext|tag where '|' indicates concatenation.
func decrypt(ciphertext []byte, key []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("malformed ciphertext")
	}

	return gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
}
