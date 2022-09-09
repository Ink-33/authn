package share

import (
	"github.com/Ink-33/authn/api/utils"
)

// RawCredential is the information about credential.
//
// _WEBAUTHN_CREDENTIAL
type RawCredential struct {
	// Version of this structure, to allow for modifications in the future.
	Version uint32 // DWORD dwVersion;

	// Size of pbID.
	IDLen uint32 // DWORD cbId;
	// Unique ID for this particular credential.
	// _Field_size_bytes_(cbId)
	IDPtr *byte // PBYTE pbId;

	// Well-known credential type specifying what this particular credential is.
	CredentialType *uint16 // LPCWSTR pwszCredentialType;
}

// RawCredentials is the information about credentials.
//
// _WEBAUTHN_CREDENTIALS
type RawCredentials struct {
	CredentialsLen uint32 // DWORD cCredentials;
	// _Field_size_(cCredentials)
	CredentialsPtr *RawCredential // PWEBAUTHN_CREDENTIAL pCredentials;
}

// Credential is the information about credential.
type Credential struct {
	// Version of this structure, to allow for modifications in the future.
	Version uint32 // DWORD dwVersion;

	// Unique ID for this particular credential.
	ID []byte

	// Well-known credential type specifying what this particular credential is.
	CredentialType string
}

// DeRaw ...
func (c *RawCredential) DeRaw() *Credential {
	if c == nil {
		return nil
	}
	return &Credential{
		Version:        c.Version,
		ID:             utils.BytesBuilder(c.IDPtr, c.IDLen),
		CredentialType: utils.UTF16PtrtoString(c.CredentialType),
	}
}
