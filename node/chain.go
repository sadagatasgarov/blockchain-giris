package node

import (
	"encoding/hex"
	"fmt"

	"gitlab.com/sadagatasgarov/bchain/proto"
	"gitlab.com/sadagatasgarov/bchain/types"
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

func (list *HeaderList) Get(index int) *proto.Header {
	if index > list.Height() {
		panic("index too high!")
	}
	return list.headers[index]
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
		headers:    NewHeaderList(),
	}
}

func (c *Chain) Height() int {
	return c.headers.Height()
}

func (c *Chain) AddBlock(b *proto.Block) error {
	// Add the header to the list of headers
	c.headers.Add(b.Header)
	// validation
	return c.blockstore.Put(b)
}

func (c *Chain) GetBlockByHash(hash []byte) (*proto.Block, error) {
	hashHex := hex.EncodeToString(hash)
	return c.blockstore.Get(hashHex)
}

func (c *Chain) GetBlockByHeight(height int) (*proto.Block, error) {
	if c.Height() > height {
		return nil, fmt.Errorf("given height (%d) too high - height (%d)", height, c.Height())
	}

	header := c.headers.Get(height)
	hash := types.HashHeader(header)
	return c.GetBlockByHash(hash)
}
