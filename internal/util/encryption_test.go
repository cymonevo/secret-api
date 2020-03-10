package util

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestEncryptDecrypt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	data := struct {
		username string
		password string
	}{
		username: "username",
		password: "password",
	}
	enc, _ := json.Marshal(data)
	tests := []struct {
		name     string
		request  []byte
		expected []byte
		isErr    bool
	}{
		{
			name:    "no error",
			request: enc,
			isErr:   false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			key := NewEncryptionKey()
			enc, err := Encrypt(test.request, key)
			assert.NoError(t, err)
			dec, err := Decrypt(enc, key)
			assert.NoError(t, err)
			assert.Equal(t, test.request, dec)
		})
	}
}
