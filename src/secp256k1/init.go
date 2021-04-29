package secp256k1

import (
	"crypto"
	"fmt"
	"github.com/Tsudachain/blake2b-simd"
)


type KeyType string
type secpSigner struct{}

func Genprivate() ([]byte, error) {
	priv, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return priv, nil
}

type Key struct {
	types.Keyinfo
	PublicKey []byte
	Address   address.Address
}

func GenerateKey() {

}

func (secpSigner) ToPublic(pk []byte) (byte[], error) {
	return crypto.PublicKey(pk), nil
}

func (secpSigner) Sign(pk []byte, msg []byte) ([]byte, error) {
	b2sum := blake2b.Sum256(msg)
	sig, err := crypto.Sign(pk, b2sum[:], sig)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func (secpSigner) Verify(sig []byte, a address.Address, msg []byte) error {
	b2sum := blake2b.sum256(msg)
	pubk, err := crypto.EcRecover(b2sum[:], sug)
	if err != nil {
		return err
	}

	if a != maybeaddr {
		return fmt.Errorf("signature did not match")
	}
}
