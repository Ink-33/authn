package hresult

// GetHResultErr ...
func GetHResultErr(res uintptr) error {
	switch res {
	case 0:
		return nil
	case 0x80070032:
		var msg ErrNotSupported = "NotSupported"
		return msg
	case 0x800704C7:
		var msg ErrCanceled = "Canceled"
		return msg
	case 0x800705B4:
		var msg ErrTimeout = "Timeout"
		return msg
	case 0x8009000F:
		var msg ErrNteExists = "NteExists"
		return msg
	case 0x80090011:
		var msg ErrNteNotFound = "NteNotFound"
		return msg
	case 0x80090016:
		var msg ErrNteBadKeyset = "NteBadKeyset"
		return msg
	case 0x80090023:
		var msg ErrNteTokenKeysetStorageFull = "NteTokenKeysetStorageFull"
		return msg
	case 0x80090027:
		var msg ErrNteInvalidParameter = "NteInvalidParameter"
		return msg
	case 0x80090029:
		var msg ErrNteNotSupported = "NteNotSupported"
		return msg
	case 0x80090035:
		var msg ErrNteDeviceNotFound = "NteDeviceNotFound"
		return msg
	case 0x80090036:
		var msg ErrNteUserCanceled = "NteUserCanceled"
		return msg
	default:
		var msg ErrUnknownError = "UnknownError"
		return msg
	}
}

// ErrNotSupported Win32: The request is not supported
type ErrNotSupported string

func (e ErrNotSupported) Error() string {
	return string(e)
}

// ErrCanceled Win32: The operation was canceled by the user
type ErrCanceled string

func (e ErrCanceled) Error() string {
	return string(e)
}

// ErrTimeout This operation returned because the timeout period expired
type ErrTimeout string

func (e ErrTimeout) Error() string {
	return string(e)
}

// ErrNteExists Object already exists
type ErrNteExists string

func (e ErrNteExists) Error() string {
	return string(e)
}

// ErrNteNotFound Object was not found
type ErrNteNotFound string

func (e ErrNteNotFound) Error() string {
	return string(e)
}

// ErrNteBadKeyset Keyset does not exist (Windows Hello not active)
type ErrNteBadKeyset string

func (e ErrNteBadKeyset) Error() string {
	return string(e)
}

// ErrNteTokenKeysetStorageFull The security token does not have storage space available for an additional container
type ErrNteTokenKeysetStorageFull string

func (e ErrNteTokenKeysetStorageFull) Error() string {
	return string(e)
}

// ErrNteInvalidParameter The parameter is incorrect
type ErrNteInvalidParameter string

func (e ErrNteInvalidParameter) Error() string {
	return string(e)
}

// ErrNteNotSupported The requested operation is not supported
type ErrNteNotSupported string

func (e ErrNteNotSupported) Error() string {
	return string(e)
}

// ErrNteDeviceNotFound The device that is required by this cryptographic provider is not found on this platform
type ErrNteDeviceNotFound string

func (e ErrNteDeviceNotFound) Error() string {
	return string(e)
}

// ErrNteUserCanceled The operation was canceled by the user
type ErrNteUserCanceled string

func (e ErrNteUserCanceled) Error() string {
	return string(e)
}

// ErrUnknownError ...
type ErrUnknownError string

func (e ErrUnknownError) Error() string {
	return string(e)
}
