package share

import "golang.org/x/sys/windows"

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

// CollectedClientData is the information about client data.
type CollectedClientData struct {
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

// CollectedClient is the information about client data json.
type CollectedClient struct {
	Type      string `json:"type"`
	Challenge string `json:"challenge"`
	Origin    string `json:"origin"`
}

// COSECredentialParameters is the information about credential parameters.
type COSECredentialParameters struct {
	CredentialParametersLen uint32 // DWORD cCredentialParameters;
	// _Field_size_(cCredentialParameters)
	CredentialParameters *COSECredentialParameter // PWEBAUTHN_COSE_CREDENTIAL_PARAMETER pCredentialParameters;
}

// COSECredentialParameter is the information about credential parameter.
type COSECredentialParameter struct {
	// Version of this structure, to allow for modifications in the future.
	Version uint32 // DWORD dwVersion;

	// Well-known credential type specifying a credential to create.
	CredentialType *uint16 // LPCWSTR pwszCredentialType

	// Well-known COSE algorithm specifying the algorithm to use for the credential.
	Alg int64 // LONG lAlg
}

//Credential is the information about credential.
type Credential struct {
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

//Credentials is the information about credentials.
type Credentials struct {
	CredentialsLen uint32 // DWORD cCredentials;
	// _Field_size_(cCredentials)
	CredentialsPtr *Credential // PWEBAUTHN_CREDENTIAL pCredentials;
}

// AuthenticatorMakeCredentialOptions ...
type AuthenticatorMakeCredentialOptions struct {
	// Version of this structure, to allow for modifications in the future.
	Version uint32 // DWORD dwVersion;

	// Time that the operation is expected to complete within.
	// This is used as guidance, and can be overridden by the platform.
	TimeoutMilliseconds uint32 // DWORD dwTimeoutMilliseconds

	// Credentials used for exclusion.
	CredentialList Credentials // WEBAUTHN_CREDENTIALS CredentialList

	// Optional extensions to parse when performing the operation.
	Extensions Extensions // WEBAUTHN_EXTENSIONS Extensions

	// Optional. Platform vs Cross-Platform Authenticators.
	AuthenticatorAttachment uint32 // DWORD dwAuthenticatorAttachment

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
	CancellationID *windows.GUID // GUID *pCancellationId

	//
	// The following fields have been added in WEBAUTHN_AUTHENTICATOR_MAKE_CREDENTIAL_OPTIONS_VERSION_3
	//

	// Exclude Credential List. If present, "CredentialList" will be ignored.
	ExcludeCredentialList *CredentialList // PWEBAUTHN_CREDENTIAL_LIST pExcludeCredentialList

	//
	// The following fields have been added in WEBAUTHN_AUTHENTICATOR_MAKE_CREDENTIAL_OPTIONS_VERSION_4
	//

	// Enterprise Attestation
	EnterpriseAttestation uint32 // DWORD dwEnterpriseAttestation

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
	BrowserInPrivateMode bool // BOOL bBrowserInPrivateMode
}

// GetCredentialsOptions for WebAuthNGetPlatformCredentialList API
type GetCredentialsOptions struct {
	// Version of this structure, to allow for modifications in the future.
	Version uint32 // DWORD dwVersion

	// Optional.
	RPID *uint16 // LPCWSTR pwszRpId

	// Optional. BrowserInPrivate Mode. Defaulting to FALSE.
	BrowserInPrivateMode bool // BOOL bBrowserInPrivateMode
}

// CredentialAttestation info.
type CredentialAttestation struct {
	// Version of this structure, to allow for modifications in the future.
	Version uint32 // DWORD dwVersion;

	// Attestation format type
	FormatType *uint16 // PCWSTR pwszFormatType

	// Size of cbAuthenticatorData.
	AuthenticatorDataLen uint32 // DWORD cbAuthenticatorData
	// Authenticator data that was created for this credential.
	// _Field_size_bytes_ (cbAuthenticatorData)
	AuthenticatorDataPtr *byte // PBYTE              pbAuthenticatorData

	// Size of CBOR encoded attestation information
	//0 => encoded as CBOR null value.
	AttestationLen uint32 // DWORD cbAttestation
	//Encoded CBOR attestation information
	// _Field_size_bytes_ (cbAttestation)
	AttestationPtr *byte // PBYTE              pbAttestation

	AttestationDecodeType uint32 // DWORD dwAttestationDecodeType
	// Following depends on the dwAttestationDecodeType
	//  WEBAUTHN_ATTESTATION_DECODE_NONE
	//      NULL - not able to decode the CBOR attestation information
	//  WEBAUTHN_ATTESTATION_DECODE_COMMON
	//      PWEBAUTHN_COMMON_ATTESTATION;
	AttestationDecode uintptr // PVOID pvAttestationDecode

	// The CBOR encoded Attestation Object to be returned to the RP.
	AttestationObjectLen uint32 // DWORD              cbAttestationObject
	// _Field_size_bytes_ (cbAttestationObject)
	AttestationObjectPtr *byte // PBYTE              pbAttestationObject

	// The CredentialId bytes extracted from the Authenticator Data.
	// Used by Edge to return to the RP.
	CredentialIDLen uint32 // DWORD              cbCredentialId
	// _Field_size_bytes_ (cbCredentialId)
	CredentialIDPtr *byte // PBYTE              pbCredentialId

	//
	// Following fields have been added in WEBAUTHN_CREDENTIAL_ATTESTATION_VERSION_2
	//

	Extensions *Extensions // WEBAUTHN_EXTENSIONS Extensions

	//
	// Following fields have been added in WEBAUTHN_CREDENTIAL_ATTESTATION_VERSION_3
	//

	// One of the WEBAUTHN_CTAP_TRANSPORT_* bits will be set corresponding to
	// the transport that was used.
	UsedTransport uint32 // DWORD dwUsedTransport

	//
	// Following fields have been added in WEBAUTHN_CREDENTIAL_ATTESTATION_VERSION_4
	//

	EpAtt              bool // BOOL bEpAtt
	LargeBlobSupported bool // BOOL  bLargeBlobSupported
	ResidentKey        bool // BOOL  bResidentKey
}

// Extensions info.
type Extensions struct {
	ExtensionsLen uint32 // DWORD cExtensions;
	// _Field_size_(cExtensions)
	ExtensionsPrt *Extension // PWEBAUTHN_EXTENSION pExtensions;
}

// Extension info。
type Extension struct {
	ExtensionIdentifier *uint16 // LPCWSTR pwszExtensionIdentifier
	ExtensionID         uint32  // DWORD               cbExtension
	ExtensionPtr        uintptr // PVOID               pvExtension
}

// CredentialEX is the information about credential with extra information, such as, dwTransports
type CredentialEX struct {
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

// CredentialList is the information about credential list with extra information
type CredentialList struct {
	Credentials uint32 // DWORD cCredentials;
	// _Field_size_            (cCredentials)
	CredentialsPtr *CredentialEX // PWEBAUTHN_CREDENTIAL_EX *ppCredentials
}
