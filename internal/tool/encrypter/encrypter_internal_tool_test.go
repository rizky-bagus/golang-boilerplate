package encrypter_test

import (
	"testing"

	"api-gorm-setting/internal/tool/encrypter"

	"github.com/stretchr/testify/assert"
)

func TestNewBcrypt(t *testing.T) {
	t.Run("successfully create an instance of Bcrypt", func(t *testing.T) {
		crypter := encrypter.NewBcrypt()
		assert.NotNil(t, crypter)
	})
}

func TestBcrypt_Encrypt(t *testing.T) {
	t.Run("encypted string is different from the original one", func(t *testing.T) {
		crypter := encrypter.NewBcrypt()

		passwords := []string{"ki hajar dewantara", "tut wuri handayani", "pendidikan indonesia"}
		for _, password := range passwords {
			result, err := crypter.Encrypt(password)

			assert.Nil(t, err)
			assert.NotEqual(t, password, result)
		}
	})
}
