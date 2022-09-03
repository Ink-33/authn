package share

// RawExtensions info.
type RawExtensions struct {
	ExtensionsLen uint32 // DWORD cExtensions;
	// _Field_size_(cExtensions)
	ExtensionsPrt *RawExtension // PWEBAUTHN_EXTENSION pExtensions;
}

// RawExtension info。
type RawExtension struct {
	ExtensionIdentifier *uint16 // LPCWSTR pwszExtensionIdentifier
	ExtensionID         uint32  // DWORD               cbExtension
	ExtensionPtr        uintptr // PVOID               pvExtension
}
