//go:build windows && amd64

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

// IsUserVerifyingPlatformAuthenticatorAvailable checks if the device owns platform authenticators.
func IsUserVerifyingPlatformAuthenticatorAvailable() (is bool) {
	_, _, _ = webauthn.MustFindProc("WebAuthNIsUserVerifyingPlatformAuthenticatorAvailable").
		Call(uintptr(unsafe.Pointer(&is)))
	return is
}

// AuthenticatorMakeCredential ...
func AuthenticatorMakeCredential(hWnd uintptr,
	rpInfo *share.RawRPInfo,
	userInfo *share.RawUserInfo,
	pkCredParams *share.RawCOSECredentialParameters,
	clientData *share.RawCollectedClientData,
	options *share.RawAuthenticatorMakeCredentialOptions) (attestation *share.CredentialAttestation, err error) {
	raw := &share.RawCredentialAttestation{}
	res, _, _ := webauthn.MustFindProc("WebAuthNAuthenticatorMakeCredential").
		Call(
			hWnd,
			uintptr(unsafe.Pointer(rpInfo)),
			uintptr(unsafe.Pointer(userInfo)),
			uintptr(unsafe.Pointer(pkCredParams)),
			uintptr(unsafe.Pointer(clientData)),
			uintptr(unsafe.Pointer(options)),
			uintptr(unsafe.Pointer(&raw)),
		)
	if res == 0 {
		defer freeCredentialAttestation(raw)
		return raw.DeRaw(), nil
	}
	return nil, hresult.HResult(res)
}
func freeCredentialAttestation(attestation *share.RawCredentialAttestation) {
	_, _, _ = webauthn.MustFindProc("WebAuthNFreeCredentialAttestation").
		Call(uintptr(unsafe.Pointer(attestation)))
}

// GetErrorName returns the following Error Names:
//
//	L"Success"              - S_OK
//	L"InvalidStateError"    - NTE_EXISTS
//	L"ConstraintError"      - HRESULT_FROM_WIN32(ERROR_NOT_SUPPORTED),
//	                          NTE_NOT_SUPPORTED,
//	                          NTE_TOKEN_KEYSET_STORAGE_FULL
//	L"NotSupportedError"    - NTE_INVALID_PARAMETER
//	L"NotAllowedError"      - NTE_DEVICE_NOT_FOUND,
//	                          NTE_NOT_FOUND,
//	                          HRESULT_FROM_WIN32(ERROR_CANCELLED),
//	                          NTE_USER_CANCELLED,
//	                          HRESULT_FROM_WIN32(ERROR_TIMEOUT)
//	L"UnknownError"         - All other hr values
func GetErrorName(hr hresult.HResult) string {
	msg, _, _ := webauthn.MustFindProc("WebAuthNGetErrorName").Call(uintptr(hr))
	return utils.UTF16PtrtoString((*uint16)(unsafe.Pointer(msg)))
}

// GetPlatformCredentialList ...
func GetPlatformCredentialList(options *share.RawGetCredentialsOptions) (credList []*share.CredentialDetails, err error) {
	proc, err := webauthn.FindProc("WebAuthNGetPlatformCredentialList")
	if err != nil {
		return nil, err
	}
	raw := &share.RawCredentialDetailsList{}
	res, _, _ := proc.
		Call(uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(&raw)))
	if res == 0 {
		defer freePlatformCredentialList(raw)
		return raw.DeRaw(), nil
	}
	return nil, hresult.HResult(res)
}

// freePlatformCredentialList frees the allocation for the WEBAUTHN_CREDENTIAL_DETAILS_LIST.
func freePlatformCredentialList(list *share.RawCredentialDetailsList) {
	_, _, _ = webauthn.MustFindProc("WebAuthNFreePlatformCredentialList").
		Call(uintptr(unsafe.Pointer(list)))
}

// AuthenticatorGetAssertion produces an assertion signature representing
// an assertion by the authenticator that the user has consented to a specific transaction,
// such as logging in or completing a purchase.
func AuthenticatorGetAssertion(hWnd uintptr,
	rpID *uint16,
	clientData *share.RawCollectedClientData,
	opts *share.RawAuthenticatorGetAssertionOptions) (assertion *share.Assertion, err error) {
	raw := &share.RawAssertion{}
	res, _, _ := webauthn.MustFindProc("WebAuthNAuthenticatorGetAssertion").
		Call(
			hWnd,
			uintptr(unsafe.Pointer(rpID)),
			uintptr(unsafe.Pointer(clientData)),
			uintptr(unsafe.Pointer(opts)),
			uintptr(unsafe.Pointer(&raw)),
		)
	if res == 0 {
		defer freeAssertion(raw)

		return raw.DeRaw(), nil
	}
	return nil, hresult.HResult(res)
}

// FreeAssertion frees an assertion previously allocated by calling WebAuthNAuthenticatorGetAssertion.
func freeAssertion(assertion *share.RawAssertion) {
	_, _, _ = webauthn.MustFindProc("WebAuthNFreeAssertion").
		Call(uintptr(unsafe.Pointer(assertion)))
}
