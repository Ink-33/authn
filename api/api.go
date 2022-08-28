package api

import (
	"unsafe"

	"github.com/Ink-33/authn/api/hresult"
	"github.com/Ink-33/authn/api/share"
	"github.com/Ink-33/authn/utils"
	"golang.org/x/sys/windows"
)

var webauthn = windows.MustLoadDLL("webauthn.dll")

// GetAPIVersionNumber returns supported webauthn version of the current system.
func GetAPIVersionNumber() uintptr {
	ver, _, _ := webauthn.MustFindProc("WebAuthNGetApiVersionNumber").Call()
	return ver
}

// AuthenticatorMakeCredential ...
func AuthenticatorMakeCredential(hWnd uintptr,
	rpInfo *share.RPInfo,
	userInfo *share.UserInfo,
	pkCredParams *share.COSECredentialParameters,
	clientData *share.CollectedClientData,
	options *share.AuthenticatorMakeCredentialOptions) (attestation *share.CredentialAttestation, err error) {
	cerdAttestation := uintptr(0)
	res, _, _ := webauthn.MustFindProc("WebAuthNAuthenticatorMakeCredential").
		Call(hWnd,
			uintptr(unsafe.Pointer(rpInfo)),
			uintptr(unsafe.Pointer(userInfo)),
			uintptr(unsafe.Pointer(pkCredParams)),
			uintptr(unsafe.Pointer(clientData)),
			uintptr(unsafe.Pointer(options)),
			uintptr(unsafe.Pointer(&cerdAttestation)))
	if res == 0 {
		return (*share.CredentialAttestation)(unsafe.Pointer(cerdAttestation)), nil
	}
	return nil, hresult.HResult(res)
}

// IsUserVerifyingPlatformAuthenticatorAvailable checks if the device owns platform authenticators.
func IsUserVerifyingPlatformAuthenticatorAvailable() uintptr {
	var is uintptr
	_, _, _ = webauthn.MustFindProc("WebAuthNIsUserVerifyingPlatformAuthenticatorAvailable").
		Call(uintptr(unsafe.Pointer(&is)))
	return is
}

// GetErrorName ...
func GetErrorName(hr uintptr) string {
	hrs, _, _ := webauthn.MustFindProc("WebAuthNGetErrorName").Call(hr)
	return utils.UTF16toString((*uint16)(unsafe.Pointer(hrs)))
}

func GetPlatformCredentialList(options share.GetCredentialsOptions) (uintptr, error) {
	var result uintptr
	res, _, _ := webauthn.MustFindProc("WebAuthNGetPlatformCredentialList").
		Call(uintptr(unsafe.Pointer(&options)), uintptr(unsafe.Pointer(&result)))
	if res == 0 {
		return result, nil
	}
	return result, hresult.HResult(res)
}
