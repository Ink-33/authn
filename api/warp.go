package api

import (
	"encoding/json"

	"github.com/Ink-33/authn/api/define"
	"github.com/Ink-33/authn/api/share"
	"golang.org/x/sys/windows"
)

// NewMakeCerdOpts returns an AuthenticatorMakeCredentialOptions struct pointer whih default value.
func NewMakeCerdOpts() *share.AuthenticatorMakeCredentialOptions {
	return &share.AuthenticatorMakeCredentialOptions{
		Version:             define.WebAuthNAuthenticatorMakeCredentialOptionsCurrentVersion,
		TimeoutMilliseconds: 60000,
		CredentialList: share.Credentials{
			CredentialsLen: 0,
			CredentialsPtr: nil,
		},
		Extensions: share.Extensions{
			ExtensionsLen: 0,
			ExtensionsPrt: nil,
		},
		AuthenticatorAttachment:         define.WebAuthNAuthenticatorAttachmentAny,
		RequireResidentKey:              false,
		UserVerificationRequirement:     define.WebAuthNUserVerificationRequirementDiscouraged,
		AttestationConveyancePreference: define.WebAuthNAttestationConveyancePreferenceNone,
		Flags:                           0,
		CancellationID:                  nil,
		ExcludeCredentialList: &share.CredentialList{
			Credentials:    0,
			CredentialsPtr: nil,
		},
		EnterpriseAttestation: define.WebAuthNEnterpriseAttestationNone,
		LargeBlobSupport:      define.WebAuthNCerdLargeBlobOperationNone,
		PreferResidentKey:     false,
		BrowserInPrivateMode:  false,
	}
}

// CreateClientData ...
func CreateClientData(action, origin, challenge, HashAlgID string) (*share.CollectedClientData, error) {
	cd := share.CollectedClient{
		Type:      action,
		Challenge: challenge,
		Origin:    origin,
	}
	cdjson, err := json.Marshal(cd)
	if err != nil {
		return nil, err
	}
	return &share.CollectedClientData{
		Version:           define.WebAuthNClientDataCurrentVersion,
		ClientDataJSONLen: uint32(len(cdjson)),
		ClientDataJSONPtr: &cdjson[0],
		HashAlgID:         windows.StringToUTF16Ptr(HashAlgID),
	}, nil
}
