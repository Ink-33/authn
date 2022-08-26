package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/Ink-33/authn/api"
	"github.com/Ink-33/authn/api/define"
	"github.com/Ink-33/authn/api/share"
	"github.com/Ink-33/authn/utils"
)

func main() {
	hWnd := utils.GetConsoleWindows()
	fmt.Printf("hWnd: %v\n", hWnd)
	is := api.IsUserVerifyingPlatformAuthenticatorAvailable()
	fmt.Printf("is: %v\n", is)
	var a uintptr
	err := api.AuthenticatorMakeCredential(utils.GetConsoleWindows(),
		&share.RPInfo{
			Version: define.WebAuthNRPEntityInformationCurrentVersion,
			ID:      syscall.StringToUTF16Ptr("tttssss"),
			Name:    syscall.StringToUTF16Ptr("Test1"),
			Icon:    syscall.StringToUTF16Ptr(""),
		},
		&share.UserInfo{
			Version:     define.WebAuthNUserEntityInformationCurrentVersion,
			IDSize:      0,
			IDbuf:       0,
			Name:        syscall.StringToUTF16Ptr("Test1"),
			Icon:        syscall.StringToUTF16Ptr(""),
			DisplayName: syscall.StringToUTF16Ptr("Test1"),
		},
		&share.CoseCredentialParameters{
			CredentialParametersSize: 0,
			CredentialParameters:     0,
		},
		&share.CollectedClientData{
			Version:            define.WebAuthNClientDataCurrentVersion,
			ClientDataJSONSize: 0,
			ClientDataJSON:     0,
			HashAlgID:          syscall.StringToUTF16Ptr(define.WebAuthNHashAlgorithmSHA256),
		},
		nil,
		uintptr(unsafe.Pointer(&a)))
	fmt.Println(a, err)

}
