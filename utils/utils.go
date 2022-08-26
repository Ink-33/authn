package utils

import (
	"syscall"
)

var kernel32 = syscall.MustLoadDLL("Kernel32.dll")

// GetConsoleWindows retrieves the window handle used by the console associated with the calling process.
func GetConsoleWindows() (hWnd uintptr) {
	hWnd, _, _ = kernel32.MustFindProc("GetConsoleWindow").Call()
	return
}
