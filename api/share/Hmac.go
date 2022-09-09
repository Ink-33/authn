package share

import "github.com/Ink-33/authn/api/utils"

// RawHMACSecretSalt ...
//
// _WEBAUTHN_HMAC_SECRET_SALT
type RawHMACSecretSalt struct {
	// Size of pbFirst.
	FirstLen uint32 // DWORD              cbFirst
	// _Field_size_bytes_ (cbFirst)
	FirstPtr *byte // PBYTE              pbFirst // Required

	// Size of pbSecond.
	SecondLen uint32 // DWORD              cbSecond
	// _Field_size_bytes_ (cbSecond)
	SecondPtr *byte // PBYTE              pbSecond
}

// RawCredWithHMACSecretSalt ...
//
// _WEBAUTHN_CRED_WITH_HMAC_SECRET_SALT
type RawCredWithHMACSecretSalt struct {
	// Size of pbCredID.
	CredIDLen uint32 // DWORD              cbCredID
	// _Field_size_bytes_ (cbCredID)
	CredIDPtr *byte // PBYTE              pbCredID // Required

	// PRF Values for above credential
	HMACSecretSalt *RawHMACSecretSalt // PWEBAUTHN_HMAC_SECRET_SALT pHmacSecretSalt // Required
}

// RawHMACSecretSaltValues ...
//
// _WEBAUTHN_HMAC_SECRET_SALT_VALUES
type RawHMACSecretSaltValues struct {
	GlobalHmacSalt *RawHMACSecretSalt // PWEBAUTHN_HMAC_SECRET_SALT pGlobalHmacSalt

	CredWithHMACSecretSaltListLen uint32 // DWORD  cCredWithHmacSecretSaltList
	// _Field_size_                         (cCredWithHmacSecretSaltList)
	CredWithHMACSecretSaltListPtr *RawCredWithHMACSecretSalt // PWEBAUTHN_CRED_WITH_HMAC_SECRET_SALT pCredWithHmacSecretSaltList
}

// HMACSecretSalt ...
type HMACSecretSalt struct {
	First  []byte
	Second []byte
}

// DeRaw ...
func (c *RawHMACSecretSalt) DeRaw() *HMACSecretSalt {
	if c == nil {
		return nil
	}
	return &HMACSecretSalt{
		First:  utils.BytesBuilder(c.FirstPtr, c.FirstLen),
		Second: utils.BytesBuilder(c.SecondPtr, c.SecondLen),
	}
}
