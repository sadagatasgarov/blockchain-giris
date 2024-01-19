package node

import (
	"encoding/hex"

	"gitlab.com/sadagatasgarov/bchain/proto"
)

type Chain struct {
	blockstore BlockStorer
}

func NewChain(bs BlockStorer) *Chain {
	return &Chain{
		blockstore: bs,
	}
}

func (c *Chain) AddBlock(b *proto.Block) error {
	// validation
	return c.blockstore.Put(b)
}

func (c *Chain) GetBlockByHash(hash []byte) (*proto.Block, error) {
	hashHex := hex.EncodeToString(hash)
	return c.blockstore.Get(hashHex)
}

func (c *Chain) GetBlockByHeight(height int) (*proto.Block, error) {
	return nil, nil
}
