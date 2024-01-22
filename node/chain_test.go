package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/sadagatasgarov/bchain/types"
	"gitlab.com/sadagatasgarov/bchain/util"
)

func TestChainHeight(t *testing.T) {
	chain:=NewChain(NewMemoryBlockStore())
	for i := 0; i < 1000; i++ {
		b := util.RandomBlock()
		assert.Nil(t, chain.AddBlock(b))
		assert.Equal(t, chain.Height(), i)
	}
}

func TestAddBlock(t *testing.T) {
	bs := NewMemoryBlockStore()
	chain := NewChain(bs)
	block := util.RandomBlock()
	blockHash := types.HashBlock(block)

	assert.Nil(t, chain.AddBlock(block))
	fetchedBlock, err := chain.GetBlockByHash(blockHash)
	assert.Nil(t, err)
	assert.Equal(t, block, fetchedBlock)

	fetchedBlockByHeight, err := chain.GetBlockByHeight(0)
	
	assert.Nil(t, err)
	assert.Equal(t, block, fetchedBlockByHeight)

}
