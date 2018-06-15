package util

import (
	"github.com/btcsuite/btcutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratePrivatePublicKeys(t *testing.T) {
	private, public := GeneratePrivatePublicKeys()
	assert.NotNil(t, private, "The private key should not be nil")
	assert.NotEmpty(t, private, "The private key should not be empty")
	assert.NotNil(t, public, "The public key should not be nil")
	assert.NotEmpty(t, public, "The public key should not be empty")

}

func TestPrivateKeyToWif(t *testing.T) {
	privateKey, publicKey := GeneratePrivatePublicKeys()
	assert.NotNil(t, privateKey, "The private key generated should not be nil")
	assert.NotNil(t, publicKey, "The public key generated should not be nil")
	wif, err := privateKeyToWif(privateKey)
	assert.NotNil(t, wif, "The wif generated should not be nil")
	assert.Nil(t, err, "The err generated should be nil")

}

func TestSignMessage(t *testing.T) {

	//Go sign a message.
	/*
		This is being done in another method to ensure we don't mix  things.
		When we verify a signature for a message we will only have access to a few things
		1. The message
		2. The signature
		3. The public key

		By doing the call in another method we ensure that only those particulars
		can be used in testing the VerifySignature method. This ensures a
		test that is valid in the real world.
	*/
	publicKeyString, hashString, messageStr := setupSignature()

	assert.NotNil(t, publicKeyString, "The public key generated should not be nil")
	assert.NotNil(t, hashString, "The message hash generated should not be nil")

	publicKeyBytes := []byte(publicKeyString)

	public, err := btcutil.NewAddressPubKey(publicKeyBytes, networks["btc"].getNetworkParams())

	assert.Nil(t, err, "The err should be nil")
	assert.NotNil(t, public, "The publicKey creation should not throw an error")

	ok := VerifySignature(public.PubKey(), hashString, messageStr)
	assert.True(t, ok, "The message signature should verify to true")

}

func setupSignature() (publicKey string, signature string, message string) {
	var private string
	private, publicKey = GeneratePrivatePublicKeys()
	wif, _ := privateKeyToWif(private)
	message = "Hi there!"
	signature = SignMessage(message, wif.PrivKey)
	return

}
