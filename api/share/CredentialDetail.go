package share

import (
	"unsafe"

	"github.com/Ink-33/authn/api/utils"
)

// RawCredentialDetails is the Credential Information for WebAuthNGetPlatformCredentialList API
//
// _WEBAUTHN_CREDENTIAL_DETAILS
type RawCredentialDetails struct {
	// Version of this structure, to allow for modifications in the future.
	Version uint32 // DWORD dwVersion

	// Size of pbCredentialID.
	CredentialIDLen uint32 // DWORD              cbCredentialID
	// _Field_size_bytes_ (cbCredentialID)
	CredentialIDPtr *byte // PBYTE              pbCredentialID

	// RP Info
	RPInformation *RawRPInfo // PWEBAUTHN_RP_ENTITY_INFORMATION pRpInformation

	// User Info
	UserInformation *RawUserInfo // PWEBAUTHN_USER_ENTITY_INFORMATION pUserInformation

	// Removable or not.
	Removable bool // BOOL bRemovable
}

// RawCredentialDetailsList ...
//
// _WEBAUTHN_CREDENTIAL_DETAILS_LIST
type RawCredentialDetailsList struct {
	CredentialDetailsLen uint32 // DWORD                        cCredentialDetails
	// _Field_size_                 (cCredentialDetails)
	CredentialDetailsPtr **RawCredentialDetails // PWEBAUTHN_CREDENTIAL_DETAILS *ppCredentialDetails
}

// CredentialDetails ...
type CredentialDetails struct {
	// Version of this structure, to allow for modifications in the future.
	Version uint32

	CredentialID []byte

	// RP Info
	RPInformation *RPInfo

	// User Info
	UserInformation *UserInfo

	// Removable or not.
	Removable bool // BOOL bRemovable
}

// DeRaw ...
func (c *RawCredentialDetails) DeRaw() *CredentialDetails {
	if c == nil {
		return nil
	}
	return &CredentialDetails{
		Version:         c.Version,
		CredentialID:    utils.BytesBuilder(c.CredentialIDPtr, c.CredentialIDLen),
		RPInformation:   c.RPInformation.DeRaw(),
		UserInformation: c.UserInformation.DeRaw(),
		Removable:       c.Removable,
	}
}

// DeRaw ...
func (c *RawCredentialDetailsList) DeRaw() []*CredentialDetails {
	if c == nil {
		return nil
	}
	rl := unsafe.Slice(c.CredentialDetailsPtr, c.CredentialDetailsLen)
	l := make([]*CredentialDetails, c.CredentialDetailsLen)
	for i := uint32(0); i < c.CredentialDetailsLen; i++ {
		l[i] = rl[i].DeRaw()
	}
	return l
}
