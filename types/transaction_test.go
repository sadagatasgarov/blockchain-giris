package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/sadagatasgarov/bchain/crypto"
	"gitlab.com/sadagatasgarov/bchain/proto"
	"gitlab.com/sadagatasgarov/bchain/util"
)

// my balance 100 coins

// want to send 5 coins to "AAA"

// 2 output
// 5 to the dude we wanne send
// 95 back to our address

func TestNewTransaction(t *testing.T) {
	fromPrivKey := crypto.GeneratePrivteKey()
	fromAddress := fromPrivKey.Public().Bytes()

	toPrivKey := crypto.GeneratePrivteKey()
	toAddress := toPrivKey.Public().Address().Bytes()

	input := &proto.TxInput{
		PrevTxHash:   util.RandomHash(),
		PrevOutIndex: 0,
		PublicKey:    fromAddress,
	}

	output1 := &proto.TxOutput{
		Amount:  5,
		Address: toAddress,
	}

	output2 := &proto.TxOutput{
		Amount:  95,
		Address: fromAddress,
	}

	tx := &proto.Transaction{
		Version: 1,
		Inputs:  []*proto.TxInput{input},
		Outputs: []*proto.TxOutput{output1, output2},
	}

	sig := SignTransaction(fromPrivKey, tx)

	input.Signature = sig.Bytes()

	//fmt.Printf("%+v\n", tx)

	assert.True(t, VerifyTransaction(tx))

}
