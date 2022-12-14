package utils

import (
	"crypto/rand"
	"encoding/base64"
	"reflect"
	"unsafe"

	"github.com/google/uuid"
	"golang.org/x/sys/windows"
)

var kernel32 = windows.NewLazySystemDLL("Kernel32.dll")
var user32 = windows.NewLazySystemDLL("User32.dll")

// GetConsoleWindow retrieves the window handle used by the console associated with the calling process.
func GetConsoleWindow() (hWnd uintptr) {
	hWnd, _, _ = kernel32.NewProc("GetConsoleWindow").Call()
	return
}

// GetHostWindow retrieves the window handle used by the foreground window who starts the caliing process.
func GetHostWindow() (hWnd uintptr) {
	return GetForegroundWindow()
}

// UTF16PtrtoString converts a pointer to a UTF16 string into a Go string.
func UTF16PtrtoString(p *uint16) string {
	if p == nil {
		return ""
	}
	return windows.UTF16ToString((*[4096]uint16)(unsafe.Pointer(p))[:])
}

// B2S convert byte slice to string.
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// S2B convert string to []byte slice.
func S2B(s string) (b []byte) {
	(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data = (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
	(*reflect.SliceHeader)(unsafe.Pointer(&b)).Cap = len(s)
	(*reflect.SliceHeader)(unsafe.Pointer(&b)).Len = len(s)
	return
}

// CreateChallenge generates a new chanllenge that will be sent to the authenticator.
//
// In order to prevent replay attacks, the challenges MUST contain enough entropy to
// make guessing them infeasible. Challenges SHOULD therefore be at least 16 bytes long.
// See https://w3c.github.io/webauthn/#sctn-cryptographic-challenges
//
// # Default 32 bytes length will be used if the given length is less than 16 or more than 256
//
// Challenge is encoded in base64.
func CreateChallenge(len int) (string, error) {
	if len < 16 || len > 256 {
		len = 16
	}
	challenge := make([]byte, len)
	_, err := rand.Read(challenge)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(challenge), nil
}

// BytesBuilder converts PBYTE to go []byte.
func BytesBuilder(ptr *byte, len uint32) (buf []byte) {
	if len == 0 {
		return
	}
	buf = make([]byte, len)
	raw := unsafe.Slice(ptr, len)
	copy(buf, raw) // make sure it is safe after call free api.
	return buf
}

// CreateCancelID returns a new windows guid
func CreateCancelID() (windows.GUID, error) {
	return windows.GUIDFromString("{" + uuid.New().String() + "}")
}

func GetForegroundWindow() (hWnd uintptr) {
	hWnd, _, _ = user32.NewProc("GetForegroundWindow").Call()
	return
}
