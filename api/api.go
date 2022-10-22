package api

import (
	"github.com/Ink-33/authn/api/define"
	"github.com/Ink-33/authn/api/raw"
	"github.com/Ink-33/authn/api/share"
	"github.com/Ink-33/authn/api/utils"
	"golang.org/x/sys/windows"
)

type User interface {
	GetID() []byte
	GetName() string
	GetDisplayName() string
	GetIcon() string
}

// WebAuthNClient ...
type WebAuthNClient struct {
	RPInfo                   share.RawRPInfo
	COSECredentialParameters []*share.RawCOSECredentialParameter
	makeCredOpts             share.RawAuthenticatorMakeCredentialOptions
	challengeLength          int
	Timeout                  uint32
}

// NewClient inits a WebAuthN client with the basic information.
//
// An rpID MUST be a valid domain string. See https://w3c.github.io/webauthn/#rp-id
func NewClient(rpID, rpName, rpIcon string) *WebAuthNClient {
	c := &WebAuthNClient{
		RPInfo:                   share.RawRPInfo{Version: define.WebAuthNRPEntityInformationCurrentVersion, ID: windows.StringToUTF16Ptr(rpID), Name: windows.StringToUTF16Ptr(rpName), Icon: windows.StringToUTF16Ptr(rpIcon)},
		COSECredentialParameters: []*share.RawCOSECredentialParameter{},
		challengeLength:          32,
	}
	c.SetDefaultCOSE()
	return c
}

// SetDefaultCOSE sets COSECredentialParameters to ES256 and RS256.
func (c *WebAuthNClient) SetDefaultCOSE() {
	c.COSECredentialParameters = []*share.RawCOSECredentialParameter{
		{
			Version:        define.WebAuthNCOSECredentialParameterCurrentVersion,
			CredentialType: windows.StringToUTF16Ptr(define.WebAuthNCredentialTypePublicKey),
			Alg:            define.WebAuthNCOSEAlgorithmECDSAP256WithSHA256,
		},
		{
			Version:        define.WebAuthNCOSECredentialParameterCurrentVersion,
			CredentialType: windows.StringToUTF16Ptr(define.WebAuthNCredentialTypePublicKey),
			Alg:            define.WebAuthNCOSEAlgorithmRSASSAPKCS1V15WithSHA256,
		},
	}
}

func (c *WebAuthNClient) MakeCredential(user User, origin string, opts *share.RawAuthenticatorMakeCredentialOptions) (*share.CredentialAttestation, error) {
	if len(c.COSECredentialParameters) == 0 {
		c.SetDefaultCOSE()
	}
	chanlleng, err := utils.CreateChallenge(c.challengeLength)
	if err != nil {
		return nil, err
	}
	cd, err := CreateClientData(define.CollectedClientCeremonyCreate, origin, chanlleng, define.WebAuthNHashAlgorithmSHA256)
	if err != nil {
		return nil, err
	}
	if opts == nil {
		opts = NewMakeCredOpts()
	}
	if c.Timeout != 0 {
		opts.TimeoutMilliseconds = c.Timeout
	}
	cancelID, err := utils.CreateCancelID()
	if err != nil {
		return nil, err
	}
	opts.CancellationID = &cancelID
	return raw.AuthenticatorMakeCredential(utils.GetHostWindow(),
		&c.RPInfo,
		&share.RawUserInfo{
			Version:     define.WebAuthNUserEntityInformationCurrentVersion,
			IDLen:       uint32(len(user.GetID())),
			IDPtr:       &user.GetID()[0],
			Name:        windows.StringToUTF16Ptr(user.GetName()),
			Icon:        windows.StringToUTF16Ptr(user.GetIcon()),
			DisplayName: windows.StringToUTF16Ptr(user.GetDisplayName()),
		},
		&share.RawCOSECredentialParameters{
			CredentialParametersLen: uint32(len(c.COSECredentialParameters)),
			CredentialParameters:    c.COSECredentialParameters[0],
		},
		cd,
		opts)
}

func (c *WebAuthNClient) GetAssertion(origin string, opts *share.RawAuthenticatorGetAssertionOptions) (*share.Assertion, error) {
	chanlleng, err := utils.CreateChallenge(c.challengeLength)
	if err != nil {
		return nil, err
	}
	cd, err := CreateClientData(define.CollectedClientCeremonyGet, origin, chanlleng, define.WebAuthNHashAlgorithmSHA256)
	if err != nil {
		return nil, err
	}
	return raw.AuthenticatorGetAssertion(
		utils.GetHostWindow(),
		c.RPInfo.ID,
		cd,
		opts,
	)
}

// GetPlatformCredentialList gets the list of WebAuthN/FIDO2 credentials currently stored for the user.
//
// If rpid is not given, will use client rpid.
//
// If rpid is "-", will let this argurement empty when calling win32 api.
func (c *WebAuthNClient) GetPlatformCredentialList(rpid string) ([]*share.CredentialDetails, error) {
	rpidu16 := (*uint16)(nil)
	switch rpid {
	case "-":
		rpidu16 = nil
	case "":
		rpidu16 = c.RPInfo.ID
	default:
		rpidu16 = windows.StringToUTF16Ptr(rpid)
	}
	return raw.GetPlatformCredentialList(
		&share.RawGetCredentialsOptions{
			Version:              define.WebAuthNGetCredentialsOptionsCurrentVersion,
			RPID:                 rpidu16,
			BrowserInPrivateMode: false,
		},
	)
}

// DeletePlatformCred removes a Public Key Credential Source stored on a Virtual Authenticator.
func (c *WebAuthNClient) DeletePlatformCred(cbCred uint32, pbCred *byte) error {
	return raw.DeletePlatformCred(cbCred, pbCred)
}
