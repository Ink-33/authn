package share

import (
	"github.com/Ink-33/authn/api/utils"
)

// RawUserInfo is the information about an User Entity
//
// _WEBAUTHN_USER_ENTITY_INFORMATION
type RawUserInfo struct {
	// Version of this structure, to allow for modifications in the future.
	// This field is required and should be set to CURRENT_VERSION above.
	Version uint32 // DWORD dwVersion;

	// Identifier for the User. This field is required.
	IDLen uint32 // DWORD              cbId
	// _Field_size_bytes_(cbId)
	IDPtr *byte // PBYTE              pbId

	// Contains a detailed name for this account, such as "john.p.smith@example.com".
	Name *uint16 // PCWSTR pwszName

	// Optional URL that can be used to retrieve an image containing the user's current avatar,
	// or a data URI that contains the image data.
	Icon *uint16 // PCWSTR pwszIcon

	// For User: Contains the friendly name associated with the user account by the Relying Party, such as "John P. Smith".
	DisplayName *uint16 // PCWSTR pwszDisplayName
}

// UserInfo is the information about an User Entity
type UserInfo struct {
	// Version of this structure, to allow for modifications in the future.
	// This field is required and should be set to CURRENT_VERSION above.
	Version uint32 // DWORD dwVersion;

	ID []byte // PBYTE              pbId

	// Contains a detailed name for this account, such as "john.p.smith@example.com".
	Name string

	// Optional URL that can be used to retrieve an image containing the user's current avatar,
	// or a data URI that contains the image data.
	Icon string

	// For User: Contains the friendly name associated with the user account by the Relying Party, such as "John P. Smith".
	DisplayName string
}

// DeRaw ...
func (c *RawUserInfo) DeRaw() *UserInfo {
	if c == nil {
		return nil
	}
	return &UserInfo{
		Version:     c.Version,
		ID:          utils.BytesBuilder(c.IDPtr, c.IDLen),
		Name:        utils.UTF16PtrtoString(c.Name),
		Icon:        utils.UTF16PtrtoString(c.Icon),
		DisplayName: utils.UTF16PtrtoString(c.DisplayName),
	}
}
