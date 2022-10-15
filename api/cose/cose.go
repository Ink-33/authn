package cose

import (
	"fmt"

	"github.com/fxamacker/cbor/v2"
)

// WebAuthN use COSE format to encode public key.
// See https://w3c.github.io/webauthn/#sctn-encoded-credPubKey-examples for example key format.

type COSEPublicKey interface{}

type PKMetadata struct {
	KeyType   int64 `cbor:"1,keyasint,omitempty"`
	Algorithm int64 `cbor:"3,keyasint,omitempty"`
}

type EC2 struct {
	PKMetadata
	Curve int64  `cbor:"-1,keyasint,omitempty"`
	X     []byte `cbor:"-2,keyasint,omitempty"`
	Y     []byte `cbor:"-3,keyasint,omitempty"`
}

type OKP struct {
	PKMetadata
	Curve int64  `cbor:"-1,keyasint,omitempty"`
	X     []byte `cbor:"-2,keyasint,omitempty"`
}

type RSA struct {
	PKMetadata
	Modulus  []byte `cbor:"-1,keyasint,omitempty"`
	Exponent []byte `cbor:"-3,keyasint,omitempty"`
}

func ParseCOSEKey(cose []byte) (pkdata COSEPublicKey, err error) {
	pkd := &PKMetadata{}
	err = cbor.Unmarshal(cose, pkd)
	if err != nil {
		return nil, err
	}
	switch pkd.KeyType {
	case 1: // Octet Key
		pk := &OKP{}
		err = cbor.Unmarshal(cose, pk)
		if err != nil {
			return nil, err
		}
		return pk, nil
	case 2: // EC2 key
		pk := &EC2{}
		err = cbor.Unmarshal(cose, pk)
		if err != nil {
			return nil, err
		}
		return pk, nil
	case 3: // RSA key
		pk := &RSA{}
		err = cbor.Unmarshal(cose, pk)
		if err != nil {
			return nil, err
		}
		return pk, nil
	default:
		return nil, fmt.Errorf("unknown key type")
	}
}
