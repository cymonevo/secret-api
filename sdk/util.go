package sdk

import (
	"net/http"

	"github.com/cymonevo/secret-api/internal/util"
)

func IsSuccess(code int) bool {
	return code >= http.StatusOK && code < 300
}

func NewEncryptionKey() *[32]byte {
	return util.NewEncryptionKey()
}

func Encrypt(plaintext []byte, key *[32]byte) (ciphertext []byte, err error) {
	return util.Encrypt(plaintext, key)
}

func Decrypt(ciphertext []byte, key *[32]byte) (plaintext []byte, err error) {
	return util.Decrypt(ciphertext, key)
}
