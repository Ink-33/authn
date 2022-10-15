package share

// RawCOSECredentialParameters is the information about credential parameters.
type RawCOSECredentialParameters struct {
	CredentialParametersLen uint32 // DWORD cCredentialParameters;
	// _Field_size_(cCredentialParameters)
	CredentialParameters *RawCOSECredentialParameter // PWEBAUTHN_COSE_CREDENTIAL_PARAMETER pCredentialParameters;
}

// RawCOSECredentialParameter is the information about credential parameter.
type RawCOSECredentialParameter struct {
	// Version of this structure, to allow for modifications in the future.
	Version uint32 // DWORD dwVersion;

	// Well-known credential type specifying a credential to create.
	CredentialType *uint16 // LPCWSTR pwszCredentialType

	// Well-known COSE algorithm specifying the algorithm to use for the credential.
	Alg int32 // LONG lAlg
}
