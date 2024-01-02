package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivteKey(t *testing.T) {
	privKey := GeneratePrivteKey()
	assert.Equal(t, len(privKey.Bytes()), privKeyLen)

	pubKey := privKey.Public()
	assert.Equal(t, len(pubKey.Bytes()), pubKeyLen)
}

func TestNewPrivateKeyFromString(t *testing.T) {
	var (
		seed       = "31360b9368cab545e1d12786564445309b6365ce84601257c1a509d7eb1fcdde"
		privKey    = NewPrivateKeyFromString(seed)
		addressStr = "3bde855791bebf949b3d727b3e16651de76ff690"
	)
	// seed := make([]byte, 32)
	// io.ReadFull(rand.Reader, seed)
	// fmt.Println(hex.EncodeToString(seed))
	assert.Equal(t, privKeyLen, len(privKey.Bytes()))
	address := privKey.Public().Address()
	assert.Equal(t, addressStr, address.String())
	//fmt.Println(address)
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
	privKey := GeneratePrivteKey()
	pubKey := privKey.Public()
	address := pubKey.Address()
	assert.Equal(t, addressLen, len(address.Bytes()))
	fmt.Println(address)
}
