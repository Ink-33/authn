package api

import (
	"syscall"
	"unsafe"

	"github.com/Ink-33/authn/api/hresult"
	"github.com/Ink-33/authn/api/share"
)

var webauthn = syscall.MustLoadDLL("webauthn.dll")

// GetAPIVersionNumber returns supported webauthn version of the current system.
func GetAPIVersionNumber() uintptr {
	ver, _, _ := webauthn.MustFindProc("WebAuthNGetApiVersionNumber").Call()
	return ver
}

// AuthenticatorMakeCredential ...
func AuthenticatorMakeCredential(hWnd uintptr,
	rpInfo *share.RPInfo,
	userInfo *share.UserInfo,
	pkCredParams *share.CoseCredentialParameters,
	clientData *share.CollectedClientData,
	options *share.AuthenticatorMakeCredentialOptions,
	cerdAttestation uintptr) error {
	res, _, _ := webauthn.MustFindProc("WebAuthNAuthenticatorMakeCredential").Call(hWnd,
		uintptr(unsafe.Pointer(rpInfo)),
		uintptr(unsafe.Pointer(userInfo)),
		uintptr(unsafe.Pointer(pkCredParams)),
		uintptr(unsafe.Pointer(clientData)),
		uintptr(unsafe.Pointer(options)),
		cerdAttestation)
	if res == 0 {
		return nil
	}
	return hresult.HResult(res)
}

// IsUserVerifyingPlatformAuthenticatorAvailable checks if the device owns platform authenticators.
func IsUserVerifyingPlatformAuthenticatorAvailable() uintptr {
	var is uintptr
	_, _, _ = webauthn.MustFindProc("WebAuthNIsUserVerifyingPlatformAuthenticatorAvailable").Call(uintptr(unsafe.Pointer(&is)))
	return is
}
