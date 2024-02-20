package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	assert.Equal(t, len(privKey.Bytes()), privKeyLen)

	pubKey := privKey.Public()
	assert.Equal(t, len(pubKey.Bytes()), pubKeyLen)
}

func TestNewPrivateKeyFromString(t *testing.T) {
	//seed := make([]byte, 32)
	// io.ReadFull(rand.Reader, seed)
	// fmt.Println(hex.EncodeToString(seed))
	var (
		seed       = "778962010c794a95d8ce02f4999b82685520525af374bd8c272307ba232a4b8a"
		addressStr = "78892dd9149a4f3ebbd9f31ce057dad4af878207"
		privKey    = NewPrivateKeyFromString(seed)
	)
	assert.Equal(t, privKeyLen, len(privKey.Bytes()))
	address := privKey.Public().Address()
	assert.Equal(t, addressStr, address.String())
	//fmt.Println(address)

}

func TestPrivateKeySign(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()

	msg := []byte("marvin key")
	sig := privKey.Sign(msg)

	assert.True(t, sig.Verify(pubKey, msg))
	//test will invalid msg
	assert.False(t, sig.Verify(pubKey, []byte("key")))
	//test with invalid pubkey

	invalidPrivKey := GeneratePrivateKey()
	invalidPubKey := invalidPrivKey.Public()
	assert.False(t, sig.Verify(invalidPubKey, msg))
}

func TestPublicKeyToAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	address := pubKey.Address()
	assert.Equal(t, addressLen, len(address.Bytes()))
}
