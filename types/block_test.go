package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/sadagatasgarov/bchain/util"
)

func TestHashBlock(t *testing.T) {
	block := util.RandomBlock()
	hash := HashBlock(block)
	assert.Equal(t, 32, len(hash))
}
