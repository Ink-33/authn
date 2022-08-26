package main

import (
	"encoding/json"
	"fmt"

	"unsafe"

	"github.com/Ink-33/authn/api"
	"github.com/Ink-33/authn/api/define"
	"github.com/Ink-33/authn/api/share"
	"github.com/Ink-33/authn/utils"
	"github.com/duo-labs/webauthn/protocol"
	"golang.org/x/sys/windows"
)

func main() {
	hWnd := utils.GetConsoleWindows()
	fmt.Printf("hWnd: %v\n", hWnd)
	is := api.IsUserVerifyingPlatformAuthenticatorAvailable()
	fmt.Printf("is: %v\n", is)
	var a uintptr
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
	chanlleng, err := protocol.CreateChallenge()
	if err != nil {
		panic(err)
	}
	cd := share.CollectedClient{
		Type:      "webauthn.create",
		Challenge: chanlleng.String(),
		Origin:    "local://demo.app",
	}
	cdjson, err := json.Marshal(cd)
	if err != nil {
		panic(err)
	}
	cose := share.COSECredentialParameter{
		Version:        define.WebAuthNCOSECredentialParameterCurrentVersion,
		CredentialType: windows.StringToUTF16Ptr(define.WebAuthNCredentialTypePublicKey),
		Alg:            define.WebAuthNCOSEAlgorithmECDSAP256WithSHA256,
	}
	err = api.AuthenticatorMakeCredential(utils.GetConsoleWindows(),
		&share.RPInfo{
			Version: define.WebAuthNRPEntityInformationCurrentVersion,
			ID:      windows.StringToUTF16Ptr("go.webauthn.demo.app"),
			Name:    windows.StringToUTF16Ptr("WebAuthN From Golang"),
			Icon:    windows.StringToUTF16Ptr(""),
		},
		&share.UserInfo{
			Version:     define.WebAuthNUserEntityInformationCurrentVersion,
			IDSize:      uint32(unsafe.Sizeof(u.ID)),
			IDbuf:       uintptr(unsafe.Pointer(&u.ID[0])),
			Name:        windows.StringToUTF16Ptr(u.Name),
			Icon:        windows.StringToUTF16Ptr(""),
			DisplayName: windows.StringToUTF16Ptr(u.DisplayName),
		},
		&share.COSECredentialParameters{
			CredentialParametersLen: uint32(1),
			CredentialParameters:    uintptr(unsafe.Pointer(&cose)),
		},
		&share.CollectedClientData{
			Version:           define.WebAuthNClientDataCurrentVersion,
			ClientDataJSONLen: uint32(len(cdjson)),
			ClientDataJSON:    uintptr(unsafe.Pointer(&cdjson[0])),
			HashAlgID:         windows.StringToUTF16Ptr(define.WebAuthNHashAlgorithmSHA256),
		},
		nil,
		uintptr(unsafe.Pointer(&a)))
	fmt.Println(a, err)
}
