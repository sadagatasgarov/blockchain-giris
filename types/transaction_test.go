package types

import (
	"fmt"
	"testing"

	"gitlab.com/sadagatasgarov/bchain/crypto"
	"gitlab.com/sadagatasgarov/bchain/proto"
	"gitlab.com/sadagatasgarov/bchain/util"
)

// my balance 100 coins

// want to send 5 coins to "AAA"

//input

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

	sig:=SignTransaction(fromPrivKey, tx)

	input.Signature = sig.Bytes()

	fmt.Printf("%+v\n", tx)

}
