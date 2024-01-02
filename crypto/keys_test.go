package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivteKey(t *testing.T) {
	privKey := GeneratePrivteKey()
	assert.Equal(t, len(privKey.Bytes()), privKeyLen)

	pubKey := privKey.Public()
	assert.Equal(t, len(pubKey.Bytes()), pubKeyLen)
}

func TestPrivateKeySign(t *testing.T) {
	privKey := GeneratePrivteKey()
	publicKey := privKey.Public()
	msg := []byte("foo bar baza")

	sig := privKey.Sign(msg)
	assert.True(t, sig.Verify(publicKey, msg))

	// Test with inwalid msg
	assert.False(t, sig.Verify(publicKey, []byte("foo")))

	// Test with invalid pubKey
	invalidPrivKey := GeneratePrivteKey()
	invalidPubKey := invalidPrivKey.Public()
	assert.False(t, sig.Verify(invalidPubKey, msg))

}

func TestPublictestToAddress(t *testing.T) {
	privKey :=GeneratePrivteKey()
	pubKey := privKey.Public()
	address:=pubKey.Address()
	assert.Equal(t, addressLen, len(address.Bytes()))
}
