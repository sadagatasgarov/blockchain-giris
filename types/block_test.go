package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/sadagatasgarov/bchain/crypto"
	"gitlab.com/sadagatasgarov/bchain/util"
)

func TestSignBlock(t *testing.T) {
	block:=util.RandomBlock()
	privKey := crypto.GeneratePrivteKey()
	pubKey := privKey.Public()

	sig := SignBlock(privKey, block)
	assert.Equal(t, 64, len(sig.Bytes()))
	assert.True(t, sig.Verify(pubKey, HashBlock(block)))
}

func TestHashBlock(t *testing.T) {
	block := util.RandomBlock()
	hash := HashBlock(block)
	assert.Equal(t, 32, len(hash))
}
