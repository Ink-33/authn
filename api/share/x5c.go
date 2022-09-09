package share

import "github.com/Ink-33/authn/api/utils"

// RawX5C is the X.509 certificate chain.
type RawX5C struct {
	// Length of X.509 encoded certificate
	DataLen uint32 // DWORD cbData;
	// X.509 encoded certificate bytes
	// _Field_size_bytes_(cbData)
	DataPtr *byte // PBYTE pbData;
}

// X5C is the X.509 certificate chain.
type X5C []byte

func (c *RawX5C) DeRaw() X5C {
	if c == nil {
		return nil
	}
	return utils.BytesBuilder(c.DataPtr, c.DataLen)
}
