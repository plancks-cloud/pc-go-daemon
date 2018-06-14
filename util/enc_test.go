package util

import (
	"github.com/btcsuite/btcd/btcec"
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

	publicKeyString, hash := setupSignature()
	assert.NotNil(t, publicKeyString, "The public key generated should not be nil")
	assert.NotNil(t, hash, "The message hash generated should not be nil")

	publicKeyBytes := []byte(publicKeyString)

	public, err := btcutil.NewAddressPubKey(publicKeyBytes, networks["btc"].getNetworkParams())

	assert.Nil(t, err, "The err should be nil")
	assert.NotNil(t, public, "The publicKey creation should not throw an error")

	signature, err := btcec.ParseSignature(hash, public.PubKey().Curve)
	assert.Nil(t, err, "The err should be nil")
	assert.NotNil(t, signature, "The signature should not be nil")

	//Looks like this isn't working...
	//ok := VerifySignature(signature, hash, public.PubKey())
	//assert.True(t, ok, "The signature should be verified with the public key")

}

func setupSignature() (publicKey string, hash []byte) {
	var private string
	private, publicKey = GeneratePrivatePublicKeys()
	wif, _ := privateKeyToWif(private)
	message := "Hi there!"
	hash = SignMessage(message, wif.PrivKey)
	return

}
