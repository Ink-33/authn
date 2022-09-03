package api

import (
	"unsafe"

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
	RPInfo                   share.RPInfo
	COSECredentialParameters []*share.COSECredentialParameter
	makeCredOpts             share.AuthenticatorMakeCredentialOptions
	challengeLength          int
	Timeout                  uint32
}

// NewClient inits a WebAuthN client with the basic information.
//
// An rpID MUST be a valid domain string. See https://w3c.github.io/webauthn/#rp-id
func NewClient(rpID, rpName, rpIcon string) *WebAuthNClient {
	c := &WebAuthNClient{
		RPInfo:                   share.RPInfo{Version: define.WebAuthNRPEntityInformationCurrentVersion, ID: windows.StringToUTF16Ptr(rpID), Name: windows.StringToUTF16Ptr(rpName), Icon: windows.StringToUTF16Ptr(rpIcon)},
		COSECredentialParameters: []*share.COSECredentialParameter{},
		challengeLength:          32,
	}
	c.SetDefaultCOSE()
	return c
}

// SetDefaultCOSE sets COSECredentialParameters to ES256 and RS256.
func (c *WebAuthNClient) SetDefaultCOSE() {
	c.COSECredentialParameters = []*share.COSECredentialParameter{
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

func (c *WebAuthNClient) MakeCredential(user User, origin string, opts *share.AuthenticatorMakeCredentialOptions) (*share.CredentialAttestation, error) {
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
		opts = NewMakeCerdOpts()
	}
	if c.Timeout != 0 {
		opts.TimeoutMilliseconds = c.Timeout
	}
	cancelID, err := utils.CreateCancelID()
	if err != nil {
		return nil, err
	}
	opts.CancellationID = &cancelID
	return raw.AuthenticatorMakeCredential(utils.GetConsoleWindows(),
		&c.RPInfo,
		&share.UserInfo{
			Version:     define.WebAuthNUserEntityInformationCurrentVersion,
			IDLen:       uint32(len(user.GetID())),
			IDPtr:       &user.GetID()[0],
			Name:        windows.StringToUTF16Ptr(user.GetName()),
			Icon:        windows.StringToUTF16Ptr(user.GetIcon()),
			DisplayName: windows.StringToUTF16Ptr(user.GetDisplayName()),
		},
		&share.COSECredentialParameters{
			CredentialParametersLen: uint32(len(c.COSECredentialParameters)),
			CredentialParameters:    c.COSECredentialParameters[0],
		},
		cd,
		opts)
}

func (c *WebAuthNClient) GetAssertion(origin string, opts *share.AuthenticatorGetAssertionOptions) (*share.Assertion, error) {
	chanlleng, err := utils.CreateChallenge(c.challengeLength)
	if err != nil {
		return nil, err
	}
	cd, err := CreateClientData(define.CollectedClientCeremonyGet, origin, chanlleng, define.WebAuthNHashAlgorithmSHA256)
	if err != nil {
		return nil, err
	}
	return raw.AuthenticatorGetAssertion(
		utils.GetConsoleWindows(),
		c.RPInfo.ID,
		cd,
		opts,
	)
}

func (c *WebAuthNClient) GetPlatformCredentialList() ([]*share.CredentialDetails, error) {
	res, err := raw.GetPlatformCredentialList(
		&share.GetCredentialsOptions{
			Version:              define.WebAuthNGetCredentialsOptionsCurrentVersion,
			RPID:                 c.RPInfo.ID,
			BrowserInPrivateMode: false,
		},
	)
	if err != nil {
		return nil, err
	}
	return unsafe.Slice(res.CredentialDetailsPtr, int(res.CredentialDetailsLen)), nil
}
