package hresult

// HResult defines error message
type HResult uintptr

func (e HResult) Error() string {
	switch e {
	case 0:
		return "Ok"
	case 0x80070032:
		return "NotSupported" // Win32: The request is not supported
	case 0x800704C7:
		return "Canceled" // Win32: The operation was canceled by the user
	case 0x800705B4:
		return "Timeout" // Win32: This operation returned because the timeout period expired
	case 0x8009000F:
		return "NteExists" // Object already exists
	case 0x80090011:
		return "NteNotFound" // Object was not found
	case 0x80090016:
		return "NteBadKeyset" // Keyset does not exist (Windows Hello not active)
	case 0x80090023:
		return "NteTokenKeysetStorageFull" // The security token does not have storage space available for an additional container
	case 0x80090027:
		return "NteInvalidParameter" // The parameter is incorrect
	case 0x80090029:
		return "NteNotSupported" // The requested operation is not supported
	case 0x80090035:
		return "NteDeviceNotFound" // The device that is required by this cryptographic provider is not found on this platform
	case 0x80090036:
		return "NteUserCanceled" // The action was cancelled by the user
	default:
		return "UnknownError"
	}
}
