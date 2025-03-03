package keys

import (
	"crypto/ed25519"
	"crypto/x509"
)

type PublicKey ed25519.PublicKey

func (pub PublicKey) ToBytes() (bs []byte, err error) {
	return x509.MarshalPKIXPublicKey(ed25519.PublicKey(pub))
}

func (pub PublicKey) ToPEMBlock() (pem PEMBlock, err error) {
	var bs []byte

	bs, err = pub.ToBytes()
	if err != nil {
		return pem, err
	}

	pem = PEMBlock{
		Type:  publicKeyPEMType,
		Bytes: bs,
	}

	return pem, nil
}
