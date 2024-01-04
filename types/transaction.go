package types

import (
	"crypto/sha256"

	"gitlab.com/sadagatasgarov/bchain/crypto"
	"gitlab.com/sadagatasgarov/bchain/proto"
	pb "google.golang.org/protobuf/proto"
)

func SignTransaction(pk *crypto.PrivateKey, tx *proto.Transaction) *crypto.Signature {
	return pk.Sign(HashTansaction(tx))
}

func HashTansaction(tx *proto.Transaction) []byte {
	b, err := pb.Marshal(tx)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(b)
	return hash[:]
}
