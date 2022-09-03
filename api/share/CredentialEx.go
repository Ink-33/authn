package share

// RawCredentialEX is the information about credential with extra information, such as, dwTransports
//
// _WEBAUTHN_CREDENTIAL_EX
type RawCredentialEX struct {
	// Version of this structure, to allow for modifications in the future.
	Version uint32 // DWORD dwVersion;

	// Size of pbID.
	IDLen uint32 // DWORD cbId;
	// Unique ID for this particular credential.
	// _Field_size_bytes_(cbId)
	InPtr *byte // PBYTE pbId;

	// Well-known credential type specifying what this particular credential is.
	CredentialType *uint16 // LPCWSTR pwszCredentialType;

	// Transports. 0 implies no transport restrictions.
	Transports uint32 // DWORD dwTransports;
}

// RawCredentialList is the information about credential list with extra information
//
// _WEBAUTHN_CREDENTIAL_LIST
type RawCredentialList struct {
	Credentials uint32 // DWORD cCredentials;
	// _Field_size_            (cCredentials)
	CredentialsPtr *RawCredentialEX // PWEBAUTHN_CREDENTIAL_EX *ppCredentials
}
