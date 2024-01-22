package node

import (
	"encoding/hex"

	"gitlab.com/sadagatasgarov/bchain/proto"
)

type HeaderList struct {
	headers []*proto.Header
}

func NewHeaderList() *HeaderList {
	return &HeaderList{
		headers: []*proto.Header{},
	}
}

func (list *HeaderList) Add(h *proto.Header) {
	list.headers = append(list.headers, h)
}

func (list *HeaderList) Height() int {
	return list.Len() - 1
}

// [h1, h2, h3, ] len=5 height=4
func (list *HeaderList) Len() int {
	return len(list.headers)
}

type Chain struct {
	blockstore BlockStorer
	headers    *HeaderList
}

func NewChain(bs BlockStorer) *Chain {
	return &Chain{
		blockstore: bs,
		headers: NewHeaderList(),
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
