package share

import "github.com/Ink-33/authn/api/utils"

// RawRPInfo is the information about an RP Entity
//
// _WEBAUTHN_RP_ENTITY_INFORMATION
type RawRPInfo struct {
	// Version of this structure, to allow for modifications in the future.
	// This field is required and should be set to CURRENT_VERSION above.
	Version uint32 //DWORD dwVersion;

	// Identifier for the RP. This field is required.
	ID *uint16 // PCWSTR pwszId;

	// Contains the friendly name of the Relying Party, such as "Acme Corporation", "Widgets Inc" or "Awesome Site".
	// This field is required.
	Name *uint16 // PCWSTR pwszName

	// Optional URL pointing to RP's logo.
	Icon *uint16 // PCWSTR pwszIcon
}

// RPInfo is the information about an RP Entity
type RPInfo struct {
	// Version of this structure, to allow for modifications in the future.
	// This field is required and should be set to CURRENT_VERSION above.
	Version uint32 //DWORD dwVersion;

	// Identifier for the RP. This field is required.
	ID string

	// Contains the friendly name of the Relying Party, such as "Acme Corporation", "Widgets Inc" or "Awesome Site".
	// This field is required.
	Name string

	// Optional URL pointing to RP's logo.
	Icon string
}

// DeRaw ...
func (c *RawRPInfo) DeRaw() *RPInfo {
	if c == nil {
		return nil
	}
	return &RPInfo{
		Version: c.Version,
		ID:      utils.UTF16PtrtoString(c.ID),
		Name:    utils.UTF16PtrtoString(c.Name),
		Icon:    utils.UTF16PtrtoString(c.Icon),
	}
}
