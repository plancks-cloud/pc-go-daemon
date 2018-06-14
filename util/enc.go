package util

import (
	"errors"

	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcutil"
)

func GeneratePrivatePublicKeys() (privateKey string, publicKey string) {
	wif, _ := networks["btc"].createPrivateKey()
	return wif.String(), string(wif.PrivKey.PubKey().SerializeCompressed())

}

func privateKeyToWif(privateKey string) (wif *btcutil.WIF, err error) {
	wif, err = networks["btc"].importWIF(privateKey)
	return
}

func SignMessage(message string, privateKey *btcec.PrivateKey) (result []byte) {
	messageHash := chainhash.DoubleHashB([]byte(message))
	signature, err := privateKey.Sign(messageHash)
	if err != nil {
		fmt.Println(err)
		return
	}
	result = signature.Serialize()
	return
}

func VerifySignature(signature *btcec.Signature, hash []byte, publicKey *btcec.PublicKey) bool {
	return signature.Verify(hash, publicKey)
}

type network struct {
	name        string
	symbol      string
	xpubkey     byte
	xprivatekey byte
}

var networks = map[string]network{
	"rdd": {name: "reddcoin", symbol: "rdd", xpubkey: 0x3d, xprivatekey: 0xbd},
	"dgb": {name: "digibyte", symbol: "dgb", xpubkey: 0x1e, xprivatekey: 0x80},
	"btc": {name: "bitcoin", symbol: "btc", xpubkey: 0x00, xprivatekey: 0x80},
	"ltc": {name: "litecoin", symbol: "ltc", xpubkey: 0x30, xprivatekey: 0xb0},
}

func (network network) getNetworkParams() *chaincfg.Params {
	networkParams := &chaincfg.MainNetParams
	networkParams.PubKeyHashAddrID = network.xpubkey
	networkParams.PrivateKeyID = network.xprivatekey
	return networkParams
}

func (network network) createPrivateKey() (*btcutil.WIF, error) {
	secret, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, err
	}
	return btcutil.NewWIF(secret, network.getNetworkParams(), true)
}

func (network network) importWIF(wifStr string) (*btcutil.WIF, error) {
	wif, err := btcutil.DecodeWIF(wifStr)
	if err != nil {
		return nil, err
	}
	if !wif.IsForNet(network.getNetworkParams()) {
		return nil, errors.New("The WIF string is not valid for the `" + network.name + "` network")
	}
	return wif, nil
}

func (network network) getAddress(wif *btcutil.WIF) (*btcutil.AddressPubKey, error) {
	return btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeCompressed(), network.getNetworkParams())
}
