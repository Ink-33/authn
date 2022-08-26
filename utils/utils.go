package utils

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var kernel32 = windows.MustLoadDLL("Kernel32.dll")

// GetConsoleWindows retrieves the window handle used by the console associated with the calling process.
func GetConsoleWindows() (hWnd uintptr) {
	hWnd, _, _ = kernel32.MustFindProc("GetConsoleWindow").Call()
	return
}

// UTF16toString converts a pointer to a UTF16 string into a Go string.
func UTF16toString(p *uint16) string {
	return windows.UTF16ToString((*[4096]uint16)(unsafe.Pointer(p))[:])
}
