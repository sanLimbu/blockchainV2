package types

import (
	"fmt"
	"testing"

	"github.com/sanLimbu/blockchain/crypto"
	"github.com/sanLimbu/blockchain/proto"
	"github.com/sanLimbu/blockchain/util"
	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {

	fromPrivKey := crypto.GeneratePrivateKey()
	fromAddress := fromPrivKey.Public().Address().Bytes()

	toPrivKey := crypto.GeneratePrivateKey()
	toAddress := toPrivKey.Public().Address().Bytes()

	input := &proto.TxInput{
		PrevTxHash:   util.RandomHash(),
		PrevOutIndex: 0,
		PublicKey:    fromPrivKey.Public().Bytes(),
	}

	output1 := &proto.TxOutput{
		Amount:  50,
		Address: toAddress,
	}

	output2 := &proto.TxOutput{
		Amount:  1000,
		Address: fromAddress,
	}

	tx := &proto.Transaction{
		Version: 1,
		Inputs:  []*proto.TxInput{input},
		Outputs: []*proto.TxOutput{output1, output2},
	}
	sig := SignTransaction(fromPrivKey, tx)
	input.Signature = sig.Bytes()
	fmt.Printf("%+v\n", tx)
	assert.True(t, VerifyTransaction(tx))
}
