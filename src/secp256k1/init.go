package secp256k1

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	 "github.com/filecoin-project/go-crypto"
	 crypto2 "github.com/filecoin-project/go-state-types/crypto"
	"github.com/minio/blake2b-simd"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/lib/sigs"

	"golang.org/x/xerrors"

)

type KeyType string
type secpSigner struct{}

func (secpSigner) GenPrivate() ([]byte, error) {
	priv, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return priv, nil
}

func GenerateKey() (*Key, error) {
	ctyp := crypto2.SigTypeSecp256k1
	pk, err := sigs.Generate(ctyp)
	if err != nil{
		return nil, err
	}

	ki := types.KeyInfo{
		Type:       "secp256k1",
		PrivateKey: pk,
	}
	return NewKey(ki)
}

func NewKey(keyinfo types.KeyInfo) (*Key, error) {
	k := &Key{
		KeyInfo: keyinfo,
	}

	var err error
	k.PublicKey, err = sigs.ToPublic(ActSigType(k.Type), k.PrivateKey)
	if err != nil {
		return nil, err
	}

	k.Address, err = address.NewSecp256k1Address(k.PublicKey)
	if err != nil{
		return nil, xerrors.Errorf("converting Secp256k1 to address: %w",err)
	}
	return k,nil
}

type Key struct {
	types.KeyInfo

	PublicKey []byte
	Address   address.Address
}

func (secpSigner) ToPublic(pk []byte) ([]byte, error) {
	return crypto.PublicKey(pk), nil
}

func (secpSigner) Sign(pk []byte, msg []byte) ([]byte, error) {
	b2sum := blake2b.Sum256(msg)
	sig, err := crypto.Sign(pk, b2sum[:])
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func (secpSigner) Verify(sig []byte, a address.Address, msg []byte) error {
	b2sum := blake2b.Sum256(msg)
	pubk, err := crypto.EcRecover(b2sum[:], sig)
	if err != nil {
		return err
	}

	maybeaddr, err := address.NewSecp256k1Address(pubk)
	if err != nil {
		return err
	}

	if a != maybeaddr {
		return fmt.Errorf("signature did not match")
	}

	return nil
}


func ActSigType(typ types.KeyType) crypto2.SigType {
	switch typ {
	case types.KTBLS:
		return crypto2.SigTypeBLS
	case types.KTSecp256k1:
		return crypto2.SigTypeSecp256k1
	default:
		return crypto2.SigTypeUnknown
	}
}