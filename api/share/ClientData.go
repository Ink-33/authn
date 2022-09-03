package share

// RawCollectedClientData is the information about client data.
type RawCollectedClientData struct {
	// Version of this structure, to allow for modifications in the future.
	// This field is required and should be set to CURRENT_VERSION above.
	Version uint32 // DWORD dwVersion

	// Size of the pbClientDataJSON field.
	ClientDataJSONLen uint32 // DWORD cbClientDataJSON
	// UTF-8 encoded JSON serialization of the client data.
	// _Field_size_bytes_ (cbClientDataJSON)
	ClientDataJSONPtr *byte // PBYTE              pbClientDataJSON

	// Hash algorithm ID used to hash the pbClientDataJSON field.
	HashAlgID *uint16 // LPCWSTR pwszHashAlgId
}

// RawCollectedClient is the information about client data json.
type RawCollectedClient struct {
	Type      string `json:"type"`
	Challenge string `json:"challenge"`
	Origin    string `json:"origin"`
}
