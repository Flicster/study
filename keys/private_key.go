package keys

import (
	"crypto/ed25519"
	"crypto/x509"
)

type PrivateKey ed25519.PrivateKey

func (priv PrivateKey) ToBytes() (bs []byte, err error) {
	return x509.MarshalPKCS8PrivateKey(ed25519.PrivateKey(priv))
}

func (priv PrivateKey) ToPEMBlock() (pem PEMBlock, err error) {
	var bs []byte

	bs, err = priv.ToBytes()
	if err != nil {
		return pem, err
	}

	pem = PEMBlock{
		Type:  privateKeyPEMType,
		Bytes: bs,
	}

	return pem, nil
}
