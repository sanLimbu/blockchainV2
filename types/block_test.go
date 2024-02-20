package types

import (
	"testing"

	"github.com/sanLimbu/blockchain/crypto"

	"github.com/sanLimbu/blockchain/util"
	"github.com/stretchr/testify/assert"
)

func TestHasBlock(t *testing.T) {
	block := util.RandomBlock()
	hash := HashBlock(block)

	assert.Equal(t, 32, len(hash))
}

func TestSignBlock(t *testing.T) {
	var (
		block   = util.RandomBlock()
		privKey = crypto.GeneratePrivateKey()
		pubKey  = privKey.Public()
	)

	sig := SignBlock(privKey, block)
	assert.Equal(t, 64, len(sig.Bytes()))
	assert.True(t, sig.Verify(pubKey, HashBlock(block)))

}
