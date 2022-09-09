package share

import (
	"unsafe"

	"github.com/Ink-33/authn/api/utils"
)

// RawExtensions contains a list of extensions
type RawExtensions struct {
	ExtensionsLen uint32 // DWORD cExtensions;
	// _Field_size_(cExtensions)
	ExtensionsPrt *RawExtension // PWEBAUTHN_EXTENSION pExtensions;
}

// RawExtension contains information about an extension.
type RawExtension struct {
	Identifier *uint16 // LPCWSTR pwszExtensionIdentifier
	ID         uint32  // DWORD               cbExtension
	Ptr        uintptr // PVOID               pvExtension
}

// Extensions contains a list of extensions
type Extensions []*Extension

// Extension contains information about an extension.
type Extension struct {
	Identifier string
	ID         uint32
	// TODO
	Ptr unsafe.Pointer
}

func (c *RawExtensions) DeRaw() Extensions {
	if c == nil || c.ExtensionsLen == 0 {
		return Extensions{}
	}
	raw := unsafe.Slice(c.ExtensionsPrt, c.ExtensionsLen)
	exts := make(Extensions, c.ExtensionsLen)
	for i := 0; i < int(c.ExtensionsLen); i++ {
		exts[i] = raw[i].DeRaw()
	}
	return exts
}

func (c *RawExtension) DeRaw() *Extension {
	if c == nil {
		return nil
	}
	return &Extension{
		Identifier: utils.UTF16PtrtoString(c.Identifier),
		ID:         c.ID,
		Ptr:        nil,
	}
}
