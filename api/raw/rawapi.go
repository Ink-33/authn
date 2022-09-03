package raw

import (
	"unsafe"

	"github.com/Ink-33/authn/api/hresult"
	"github.com/Ink-33/authn/api/share"
	"github.com/Ink-33/authn/api/utils"
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
		Call(
			hWnd,
			uintptr(unsafe.Pointer(rpInfo)),
			uintptr(unsafe.Pointer(userInfo)),
			uintptr(unsafe.Pointer(pkCredParams)),
			uintptr(unsafe.Pointer(clientData)),
			uintptr(unsafe.Pointer(options)),
			uintptr(unsafe.Pointer(&cerdAttestation)),
		)
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
//
// Returns the following Error Names:
//  L"Success"              - S_OK
//  L"InvalidStateError"    - NTE_EXISTS
//  L"ConstraintError"      - HRESULT_FROM_WIN32(ERROR_NOT_SUPPORTED),
//                            NTE_NOT_SUPPORTED,
//                            NTE_TOKEN_KEYSET_STORAGE_FULL
//  L"NotSupportedError"    - NTE_INVALID_PARAMETER
//  L"NotAllowedError"      - NTE_DEVICE_NOT_FOUND,
//                            NTE_NOT_FOUND,
//                            HRESULT_FROM_WIN32(ERROR_CANCELLED),
//                            NTE_USER_CANCELLED,
//                            HRESULT_FROM_WIN32(ERROR_TIMEOUT)
//  L"UnknownError"         - All other hr values
//
func GetErrorName(hr hresult.HResult) string {
	msg, _, _ := webauthn.MustFindProc("WebAuthNGetErrorName").Call(uintptr(hr))
	return utils.UTF16PtrtoString((*uint16)(unsafe.Pointer(msg)))
}

// GetPlatformCredentialList ...
func GetPlatformCredentialList(options *share.GetCredentialsOptions) (*share.CredentialDetailsList, error) {
	var result uintptr
	res, _, _ := webauthn.MustFindProc("WebAuthNGetPlatformCredentialList").
		Call(uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&result)))
	if res == 0 {
		return (*share.CredentialDetailsList)(unsafe.Pointer(result)), nil
	}
	return nil, hresult.HResult(res)
}

// AuthenticatorGetAssertion ...
func AuthenticatorGetAssertion(hWnd uintptr,
	rpID *uint16,
	clientData *share.CollectedClientData,
	opts *share.AuthenticatorGetAssertionOptions) (assertion *share.Assertion, err error) {
	var result uintptr
	res, _, _ := webauthn.MustFindProc("WebAuthNAuthenticatorGetAssertion").
		Call(
			hWnd,
			uintptr(unsafe.Pointer(rpID)),
			uintptr(unsafe.Pointer(clientData)),
			uintptr(unsafe.Pointer(opts)),
			uintptr(unsafe.Pointer(&result)),
		)
	if res == 0 {
		return (*share.Assertion)(unsafe.Pointer(result)), nil
	}
	return nil, hresult.HResult(res)
}
