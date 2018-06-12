package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPublicPrivateKey(t *testing.T) {
	private, public := GeneratePrivatePublicKeys()
	assert.NotNil(t, private, "The private key should not be nil")
	assert.NotEmpty(t, private, "The private key should not be empty")
	assert.NotNil(t, public, "The public key should not be nil")
	assert.NotEmpty(t, public, "The public key should not be empty")

}
