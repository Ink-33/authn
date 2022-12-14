package api

import (
	"encoding/json"

	"github.com/Ink-33/authn/api/define"
	"github.com/Ink-33/authn/api/share"
	"golang.org/x/sys/windows"
)

// NewMakeCredOpts returns an AuthenticatorMakeCredentialOptions struct pointer whih default value.
func NewMakeCredOpts() *share.RawAuthenticatorMakeCredentialOptions {
	return &share.RawAuthenticatorMakeCredentialOptions{
		Version:             define.WebAuthNAuthenticatorMakeCredentialOptionsCurrentVersion,
		TimeoutMilliseconds: 60000,
		CredentialList: share.RawCredentials{
			CredentialsLen: 0,
			CredentialsPtr: nil,
		},
		Extensions: share.RawExtensions{
			ExtensionsLen: 0,
			ExtensionsPrt: nil,
		},
		AuthenticatorAttachment:         define.WebAuthNAuthenticatorAttachmentAny,
		RequireResidentKey:              false,
		UserVerificationRequirement:     define.WebAuthNUserVerificationRequirementDiscouraged,
		AttestationConveyancePreference: define.WebAuthNAttestationConveyancePreferenceNone,
		Flags:                           0,
		CancellationID:                  nil,
		ExcludeCredentialList: &share.RawCredentialList{
			Credentials:    0,
			CredentialsPtr: nil,
		},
		EnterpriseAttestation: define.WebAuthNEnterpriseAttestationNone,
		LargeBlobSupport:      define.WebAuthNCredLargeBlobOperationNone,
		PreferResidentKey:     false,
		BrowserInPrivateMode:  false,
	}
}

// NewGetAssertionOptions returns an AuthenticatorGetAssertionOptions struct pointer whih default value.
func NewGetAssertionOptions() *share.RawAuthenticatorGetAssertionOptions {
	return &share.RawAuthenticatorGetAssertionOptions{
		Version:             define.WebAuthNAuthenticatorGetAssertionOptionsCurrentVersion,
		TimeoutMilliseconds: 60000,
		CredentialList: share.RawCredentials{
			CredentialsLen: 0,
			CredentialsPtr: nil,
		},
		Extensions: share.RawExtensions{
			ExtensionsLen: 0,
			ExtensionsPrt: nil,
		},
		AuthenticatorAttachment:     define.WebAuthNAuthenticatorAttachmentAny,
		UserVerificationRequirement: define.WebAuthNUserVerificationRequirementDiscouraged,
		Flags:                       0,
		U2fAppID:                    nil,
		IsU2fAppIDUsed:              nil,
		CancellationID:              nil,
		AllowCredentialList: &share.RawCredentialList{
			Credentials:    0,
			CredentialsPtr: nil,
		},
		CredLargeBlobOperation: define.WebAuthNCredLargeBlobOperationNone,
		CredLargeBlobLen:       0,
		CredLargeBlobPtr:       0,
		HMACSecretSaltValues:   nil,
	}
}

// CreateClientData ...
func CreateClientData(action, origin, challenge, HashAlgID string) (*share.RawCollectedClientData, error) {
	cd := share.RawCollectedClient{
		Type:      action,
		Challenge: challenge,
		Origin:    origin,
	}
	cdjson, err := json.Marshal(cd)
	if err != nil {
		return nil, err
	}
	return &share.RawCollectedClientData{
		Version:           define.WebAuthNClientDataCurrentVersion,
		ClientDataJSONLen: uint32(len(cdjson)),
		ClientDataJSONPtr: &cdjson[0],
		HashAlgID:         windows.StringToUTF16Ptr(HashAlgID),
	}, nil
}
