package share

import "golang.org/x/sys/windows"

// RawAuthenticatorGetAssertionOptions ...
type RawAuthenticatorGetAssertionOptions struct {
	// Version of this structure, to allow for modifications in the future.
	Version uint32 // DWORD

	// Time that the operation is expected to complete within.
	// This is used as guidance, and can be overridden by the platform.
	TimeoutMilliseconds uint32 // DWORD

	// Allowed Credentials List.
	CredentialList RawCredentials // WEBAUTHN_CREDENTIALS CredentialList

	// Optional extensions to parse when performing the operation.
	Extensions RawExtensions // WEBAUTHN_EXTENSIONS Extensions

	// Optional. Platform vs Cross-Platform Authenticators.
	AuthenticatorAttachment uint32 // DWORD dwAuthenticatorAttachment

	// User Verification Requirement.
	UserVerificationRequirement uint32 // DWORD dwUserVerificationRequirement

	// Flags
	Flags uint32 // DWORD dwFlags

	//
	// The following fields have been added in WEBAUTHN_AUTHENTICATOR_GET_ASSERTION_OPTIONS_VERSION_2
	//

	// Optional identifier for the U2F AppId. Converted to UTF8 before being hashed. Not lower cased.
	U2fAppID *uint16 // PCWSTR pwszU2fAppId

	// If the following is non-NULL, then, set to TRUE if the above pwszU2fAppid was used instead of
	// PCWSTR pwszRpId;
	IsU2fAppIDUsed *bool // BOOL *pbU2fAppId

	//
	// The following fields have been added in WEBAUTHN_AUTHENTICATOR_GET_ASSERTION_OPTIONS_VERSION_3
	//

	// Cancellation Id - Optional - See WebAuthNGetCancellationId
	CancellationID *windows.GUID // GUID *pCancellationId

	//
	// The following fields have been added in WEBAUTHN_AUTHENTICATOR_GET_ASSERTION_OPTIONS_VERSION_4
	//

	// Allow Credential List. If present, "CredentialList" will be ignored.
	AllowCredentialList *RawCredentialList // PWEBAUTHN_CREDENTIAL_LIST pAllowCredentialList

	//
	// The following fields have been added in WEBAUTHN_AUTHENTICATOR_GET_ASSERTION_OPTIONS_VERSION_5
	//

	CredLargeBlobOperation uint32 // DWORD dwCredLargeBlobOperation

	// Size of pbCredLargeBlob
	CredLargeBlobLen uint32 // DWORD              cbCredLargeBlob
	// _Field_size_bytes_ (cbCredLargeBlob)
	CredLargeBlobPtr byte // PBYTE              pbCredLargeBlob

	//
	// The following fields have been added in WEBAUTHN_AUTHENTICATOR_GET_ASSERTION_OPTIONS_VERSION_6
	//

	// PRF values which will be converted into HMAC-SECRET values according to WebAuthn Spec.
	HMACSecretSaltValues *RawHMACSecretSaltValues //  PWEBAUTHN_HMAC_SECRET_SALT_VALUES pHmacSecretSaltValues

	// Optional. BrowserInPrivate Mode. Defaulting to FALSE.
	BrowserInPrivateMode bool // BOOL bBrowserInPrivateMode
}
