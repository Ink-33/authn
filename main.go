package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"unsafe"

	"github.com/Ink-33/authn/api/define"
	"github.com/Ink-33/authn/api/raw"
	"github.com/Ink-33/authn/api/share"
	"github.com/Ink-33/authn/utils"
	"github.com/fxamacker/cbor/v2"
	"github.com/google/uuid"
	"golang.org/x/sys/windows"
)

func main() {
	hWnd := utils.GetConsoleWindows()
	fmt.Printf("hWnd: %v\n", hWnd)
	is := raw.IsUserVerifyingPlatformAuthenticatorAvailable()
	fmt.Printf("is: %v\n", is)
	type user struct {
		Name        string
		DisplayName string
		ID          []byte
	}
	u := user{
		Name:        "test",
		DisplayName: "Test",
		ID:          []byte("oI24"),
	}
	chanlleng, err := utils.CreateChallenge()
	if err != nil {
		panic(err)
	}
	cd := share.CollectedClient{
		Type:      "webauthn.create",
		Challenge: chanlleng,
		Origin:    "local://demo.app",
	}
	cdjson, err := json.Marshal(cd)
	if err != nil {
		panic(err)
	}
	winguid, err := windows.GUIDFromString("{" + uuid.New().String() + "}")
	if err != nil {
		panic(err)
	}
	opt := &share.AuthenticatorMakeCredentialOptions{
		Version:             define.WebAuthNAuthenticatorMakeCredentialOptionsCurrentVersion,
		TimeoutMilliseconds: 30000,
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
		CancellationID:                  &winguid,
		ExcludeCredentialList: &share.CredentialList{
			Credentials:    0,
			CredentialsPtr: &share.CredentialEX{},
		},
		EnterpriseAttestation: define.WebAuthNEnterpriseAttestationNone,
		LargeBlobSupport:      define.WebAuthNLargeBlobSupportNone,
		PreferResidentKey:     false,
		BrowserInPrivateMode:  false,
	}
	coses := []share.COSECredentialParameter{
		{
			Version:        define.WebAuthNCOSECredentialParameterCurrentVersion,
			CredentialType: windows.StringToUTF16Ptr(define.WebAuthNCredentialTypePublicKey),
			Alg:            define.WebAuthNCOSEAlgorithmECDSAP256WithSHA256,
		},
		{Version: define.WebAuthNCOSECredentialParameterCurrentVersion,
			CredentialType: windows.StringToUTF16Ptr(define.WebAuthNCredentialTypePublicKey),
			Alg:            define.WebAuthNCOSEAlgorithmRSASSAPKCS1V15WithSHA256,
		}}

	a, err := raw.AuthenticatorMakeCredential(utils.GetConsoleWindows(),
		&share.RPInfo{
			Version: define.WebAuthNRPEntityInformationCurrentVersion,
			ID:      windows.StringToUTF16Ptr("go.webauthn.demo.app"),
			Name:    windows.StringToUTF16Ptr("WebAuthN From Golang"),
			Icon:    windows.StringToUTF16Ptr(""),
		},
		&share.UserInfo{
			Version:     define.WebAuthNUserEntityInformationCurrentVersion,
			IDLen:       uint32(len(u.ID)),
			IDPtr:       &u.ID[0],
			Name:        windows.StringToUTF16Ptr(u.Name),
			Icon:        windows.StringToUTF16Ptr(""),
			DisplayName: windows.StringToUTF16Ptr(u.DisplayName),
		},
		&share.COSECredentialParameters{
			CredentialParametersLen: uint32(len(coses)),
			CredentialParameters:    &coses[0],
		},
		&share.CollectedClientData{
			Version:           define.WebAuthNClientDataCurrentVersion,
			ClientDataJSONLen: uint32(len(cdjson)),
			ClientDataJSONPtr: &cdjson[0],
			HashAlgID:         windows.StringToUTF16Ptr(define.WebAuthNHashAlgorithmSHA256),
		},
		opt,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("a.Version: %v\n", a.Version)
	fmt.Printf("a.FormatType: %v\n", utils.UTF16toString(a.FormatType))
	fmt.Printf("a.AttestationDecodeType: %v\n", a.AttestationDecodeType)
	fmt.Printf("a.UsedTransport: %v\n", a.UsedTransport)
	fmt.Printf("a.EpAtt: %v\n", a.EpAtt)
	fmt.Printf("a.LargeBlobSupported: %v\n", a.LargeBlobSupported)
	fmt.Printf("a.ResidentKey: %v\n", a.ResidentKey)

	fmt.Printf("a.Extensions: %v\n", a.Extensions)

	fmt.Printf("AuthenticatorData: %v\n",
		base64.RawStdEncoding.EncodeToString(unsafe.Slice(a.AuthenticatorDataPtr, a.AuthenticatorDataLen)))

	atM := map[string]any{}
	err = cbor.Unmarshal(unsafe.Slice(a.AttestationPtr, a.AttestationLen), &atM)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Attestation: %v\n", atM)

	atoM := map[string]any{}
	err = cbor.Unmarshal(unsafe.Slice(a.AttestationObjectPtr, a.AttestationObjectLen), &atoM)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AttestationObject: %v\n", atoM)

	fmt.Printf("CredentialID: %v\n",
		base64.RawURLEncoding.EncodeToString(unsafe.Slice(a.CredentialIDPtr, a.CredentialIDLen)))

}
