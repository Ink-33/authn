package share

// RPInfo is the information about an RP Entity
type RPInfo struct {
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

// UserInfo is the information about an User Entity
type UserInfo struct {
	// Version of this structure, to allow for modifications in the future.
	// This field is required and should be set to CURRENT_VERSION above.
	Version uint32 //DWORD dwVersion;

	// Identifier for the User. This field is required.
	IDSize uint32 //DWORD              cbId
	// _Field_size_bytes_(cbId)
	IDbuf uintptr //PBYTE              pbId

	// Contains a detailed name for this account, such as "john.p.smith@example.com".
	Name *uint16 //PCWSTR pwszName

	// Optional URL that can be used to retrieve an image containing the user's current avatar,
	// or a data URI that contains the image data.
	Icon *uint16 //PCWSTR pwszIcon

	// For User: Contains the friendly name associated with the user account by the Relying Party, such as "John P. Smith".
	DisplayName *uint16 //PCWSTR pwszDisplayName
}

// CollectedClientData is the information about client data.
type CollectedClientData struct {
	// Version of this structure, to allow for modifications in the future.
	// This field is required and should be set to CURRENT_VERSION above.
	Version uint32 // DWORD dwVersion

	// Size of the pbClientDataJSON field.
	ClientDataJSONSize uint32 // DWORD cbClientDataJSON
	// UTF-8 encoded JSON serialization of the client data.
	// _Field_size_bytes_ (cbClientDataJSON)
	ClientDataJSON uintptr // PBYTE              pbClientDataJSON

	// Hash algorithm ID used to hash the pbClientDataJSON field.
	HashAlgID *uint16 // LPCWSTR pwszHashAlgId
}

// CoseCredentialParameters is the information about credential parameters.
type CoseCredentialParameters struct {
	CredentialParametersSize uint32 // DWORD cCredentialParameters;
	// _Field_size_(cCredentialParameters)
	CredentialParameters uintptr // PWEBAUTHN_COSE_CREDENTIAL_PARAMETER pCredentialParameters;
}

// CoseCredentialParameter is the information about credential parameter.
type CoseCredentialParameter struct {
	// Version of this structure, to allow for modifications in the future.
	Version uint32 //DWORD dwVersion;

	// Well-known credential type specifying a credential to create.
	CredentialType *uint16 // LPCWSTR pwszCredentialType

	// Well-known COSE algorithm specifying the algorithm to use for the credential.
	Alg int64 //LONG lAlg
}

//Credential is the information about credential.
type Credential struct {
	// Version of this structure, to allow for modifications in the future.
	Version uint32 // DWORD dwVersion;

	// Size of pbID.
	IDSize uint32 // DWORD cbId;
	// Unique ID for this particular credential.
	// _Field_size_bytes_(cbId)
	ID uintptr // PBYTE pbId;

	// Well-known credential type specifying what this particular credential is.
	CredentialType *uint16 // LPCWSTR pwszCredentialType;
}

//Credentials is the information about credentials.
type Credentials struct {
	CredentialsSize uint32 // DWORD cCredentials;
	// _Field_size_(cCredentials)
	CredentialsList uintptr // PWEBAUTHN_CREDENTIAL pCredentials;
}

// AuthenticatorMakeCredentialOptions ...
type AuthenticatorMakeCredentialOptions struct {
	// Version of this structure, to allow for modifications in the future.
	Version uint32 // DWORD dwVersion;

	// Time that the operation is expected to complete within.
	// This is used as guidance, and can be overridden by the platform.
	TimeoutMilliseconds uint32 // DWORD dwTimeoutMilliseconds

	// Credentials used for exclusion.
	CredentialList uintptr // WEBAUTHN_CREDENTIALS CredentialList

	// Optional extensions to parse when performing the operation.
	Extensions uintptr // WEBAUTHN_EXTENSIONS Extensions

	// Optional. Platform vs Cross-Platform Authenticators.
	AuthenticatorAttachment uint32 //DWORD dwAuthenticatorAttachment

	// Optional. Require key to be resident or not. Defaulting to FALSE.
	RequireResidentKey bool // BOOL bRequireResidentKey

	// User Verification Requirement.
	UserVerificationRequirement uint32 // DWORD dwUserVerificationRequirement

	// Attestation Conveyance Preference.
	AttestationConveyancePreference uint32 // DWORD dwAttestationConveyancePreference

	// Reserved for future Use
	Flags uint32 // DWORD dwFlags

	//
	// The following fields have been added in WEBAUTHN_AUTHENTICATOR_MAKE_CREDENTIAL_OPTIONS_VERSION_2
	//

	// Cancellation Id - Optional - See WebAuthNGetCancellationId
	CancellationID uintptr // GUID *pCancellationId

	//
	// The following fields have been added in WEBAUTHN_AUTHENTICATOR_MAKE_CREDENTIAL_OPTIONS_VERSION_3
	//

	// Exclude Credential List. If present, "CredentialList" will be ignored.
	ExcludeCredentialList uintptr // PWEBAUTHN_CREDENTIAL_LIST pExcludeCredentialList

	//
	// The following fields have been added in WEBAUTHN_AUTHENTICATOR_MAKE_CREDENTIAL_OPTIONS_VERSION_4
	//

	// Enterprise Attestation
	EnterpriseAttestation uint32 //DWORD dwEnterpriseAttestation

	// Large Blob Support: none, required or preferred
	//
	// NTE_INVALID_PARAMETER when large blob required or preferred and
	//   bRequireResidentKey isn't set to TRUE
	LargeBlobSupport uint32 // DWORD dwLargeBlobSupport

	// Optional. Prefer key to be resident. Defaulting to FALSE. When TRUE,
	// overrides the above bRequireResidentKey.
	PreferResidentKey bool // BOOL bPreferResidentKey

	//
	// The following fields have been added in WEBAUTHN_AUTHENTICATOR_MAKE_CREDENTIAL_OPTIONS_VERSION_5
	//

	// Optional. BrowserInPrivate Mode. Defaulting to FALSE.
	BrowserInPrivateMode bool //BOOL bBrowserInPrivateMode
}
