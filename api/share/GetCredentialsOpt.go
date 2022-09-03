package share

// RawGetCredentialsOptions for WebAuthNGetPlatformCredentialList API
type RawGetCredentialsOptions struct {
	// Version of this structure, to allow for modifications in the future.
	Version uint32 // DWORD dwVersion

	// Optional.
	RPID *uint16 // LPCWSTR pwszRpId

	// Optional. BrowserInPrivate Mode. Defaulting to FALSE.
	BrowserInPrivateMode bool // BOOL bBrowserInPrivateMode
}
