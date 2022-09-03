package share

import "golang.org/x/sys/windows"

// RawAuthenticatorMakeCredentialOptions ...
type RawAuthenticatorMakeCredentialOptions struct {
	// Version of this structure, to allow for modifications in the future.
	Version uint32 // DWORD dwVersion;

	// Time that the operation is expected to complete within.
	// This is used as guidance, and can be overridden by the platform.
	TimeoutMilliseconds uint32 // DWORD dwTimeoutMilliseconds

	// Credentials used for exclusion.
	CredentialList RawCredentials // WEBAUTHN_CREDENTIALS CredentialList

	// Optional extensions to parse when performing the operation.
	Extensions RawExtensions // WEBAUTHN_EXTENSIONS Extensions

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
	ExcludeCredentialList *RawCredentialList // PWEBAUTHN_CREDENTIAL_LIST pExcludeCredentialList

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
