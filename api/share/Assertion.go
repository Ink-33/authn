package share

import "github.com/Ink-33/authn/api/utils"

// RawAssertion is authenticatorGetAssertion output.
type RawAssertion struct {
	// Version of this structure, to allow for modifications in the future.
	Version uint32 // DWORD dwVersion;

	// Size of cbAuthenticatorData.
	AuthenticatorDataLen uint32 // DWORD cbAuthenticatorData;
	// Authenticator data that was created for this assertion.
	// _Field_size_bytes_(cbAuthenticatorData)
	AuthenticatorDataPtr *byte // PBYTE pbAuthenticatorData;

	// Size of pbSignature.
	SignatureLen uint32 // DWORD cbSignature;
	// Signature that was generated for this assertion.
	// _Field_size_bytes_(cbSignature)
	SignaturePtr *byte // PBYTE pbSignature;

	// Credential that was used for this assertion.
	Credential RawCredential // WEBAUTHN_CREDENTIAL Credential;

	// Size of User Id
	UserIDLen uint32 // DWORD cbUserId;
	// UserId
	// _Field_size_bytes_(cbUserId)
	UserIDPtr *byte // PBYTE pbUserId;

	//
	// Following fields have been added in WEBAUTHN_ASSERTION_VERSION_2
	//

	Extensions RawExtensions // WEBAUTHN_EXTENSIONS Extensions

	// Size of pbCredLargeBlob
	CredLargeBlobLen uint32 // DWORD              cbCredLargeBlob
	// _Field_size_bytes_ (cbCredLargeBlob)
	CredLargeBlobPtr *byte // PBYTE              pbCredLargeBlob

	CredLargeBlobStatus uint32 // DWORD dwCredLargeBlobStatus

	//
	// Following fields have been added in WEBAUTHN_ASSERTION_VERSION_3
	//

	HMACSecret *RawHMACSecretSalt //PWEBAUTHN_HMAC_SECRET_SALT pHmacSecret
}

// Assertion ...
type Assertion struct {
	// Version of this structure, to allow for modifications in the future.
	Version uint32

	// Authenticator data that was created for this assertion.
	AuthenticatorData *AuthenticatorData

	// Signature that was generated for this assertion.
	Signature []byte

	// Credential that was used for this assertion.
	Credential Credential

	// UserID
	UserID []byte

	//
	// Following fields have been added in WEBAUTHN_ASSERTION_VERSION_2
	//
	// TODO: warp this field
	Extensions RawExtensions

	CredLargeBlob []byte

	CredLargeBlobStatus uint32 // DWORD dwCredLargeBlobStatus

	//
	// Following fields have been added in WEBAUTHN_ASSERTION_VERSION_3
	//

	HMACSecret *HMACSecretSalt //PWEBAUTHN_HMAC_SECRET_SALT pHmacSecret
}

// DeRaw ...
func (c *RawAssertion) DeRaw() *Assertion {
	if c == nil {
		return nil
	}
	return &Assertion{
		Version: c.Version,
		AuthenticatorData: func() *AuthenticatorData {
			d, _ := ParseAuthenticatorData(
				utils.BytesBuilder(c.AuthenticatorDataPtr, c.AuthenticatorDataLen),
			)
			return d
		}(),
		Signature:           utils.BytesBuilder(c.SignaturePtr, c.SignatureLen),
		Credential:          *c.Credential.DeRaw(),
		UserID:              utils.BytesBuilder(c.UserIDPtr, c.UserIDLen),
		Extensions:          c.Extensions,
		CredLargeBlob:       utils.BytesBuilder(c.CredLargeBlobPtr, c.CredLargeBlobLen),
		CredLargeBlobStatus: c.CredLargeBlobStatus,
		HMACSecret:          c.HMACSecret.DeRaw(),
	}
}
