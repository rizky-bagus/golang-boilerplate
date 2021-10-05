package encrypter

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	cost = 10
)

// Bcrypt holds the job to encrypt a string.
// It uses https://godoc.org/golang.org/x/crypto/bcrypt.
type Bcrypt struct {
}

// NewBcrypt creates an instance of Bcrypt.
func NewBcrypt() *Bcrypt {
	return &Bcrypt{}
}

// Encrypt encrypts a string.
func (b *Bcrypt) Encrypt(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hashed), err
}

// CompareEncryptedAndPlain compares encrypted and plain string.
func (b *Bcrypt) CompareEncryptedAndPlain(encrypted, plain string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(plain)); err != nil {
		return false
	}
	return true
}
