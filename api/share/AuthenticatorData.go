package share

import (
	"encoding/binary"

	"github.com/Ink-33/authn/api/cose"
	"github.com/fxamacker/cbor/v2"
)

type AuthenticatorData struct {
	RPIDHash               []byte
	Flags                  Flags
	SignCounter            uint32
	AttestedCredentialData *CredentialData
	Extensions             any
}

type CredentialData struct {
	AAGUID              []byte
	CredentialID        []byte
	CredentialPublicKey cose.COSEPublicKey
}

type Flags struct {
	UserPresent            bool
	RFU1                   bool
	UserVerified           bool
	BackupEligibility      bool
	BackupState            bool
	RFU2                   bool
	AttestedCredentialData bool
	ExtensionData          bool
}

func ParseAuthenticatorData(data []byte) (*AuthenticatorData, error) {
	d := &AuthenticatorData{
		Flags: Flags{
			UserPresent:            data[32]&(1<<0) == 1<<0,
			RFU1:                   data[32]&(1<<1) == 1<<1,
			UserVerified:           data[32]&(1<<2) == 1<<2,
			BackupEligibility:      data[32]&(1<<3) == 1<<3,
			BackupState:            data[32]&(1<<4) == 1<<4,
			RFU2:                   data[32]&(1<<5) == 1<<5,
			AttestedCredentialData: data[32]&(1<<6) == 1<<6,
			ExtensionData:          data[32]&(1<<7) == 1<<7,
		},
		SignCounter: binary.BigEndian.Uint32(data[33:37]),
		Extensions:  nil,
	}
	copy(d.RPIDHash[:], data[:32])
	if d.Flags.AttestedCredentialData {
		cidlen := binary.BigEndian.Uint16(data[53:55])
		cd := &CredentialData{
			CredentialID:        data[55 : 55+cidlen],
			CredentialPublicKey: nil,
		}
		copy(cd.AAGUID[:], data[37:53])
		i := 55 + int(cidlen) + 1
		for ; i < len(data)+1; i++ {
			pk, err := cose.ParseCOSEKey(data[55+cidlen : i])
			if err == nil {
				cd.CredentialPublicKey = pk
				break
			}
		}
		d.AttestedCredentialData = cd
		if i != len(data) {
			err := cbor.Unmarshal(data[i:], d.Extensions)
			if err != nil {
				return nil, err
			}
		}
	}

	return d, nil
}
