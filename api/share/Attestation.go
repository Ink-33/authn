package share

// RawCredentialAttestation info.
type RawCredentialAttestation struct {
	// Version of this structure, to allow for modifications in the future.
	Version uint32 // DWORD dwVersion;

	// Attestation format type
	FormatType *uint16 // PCWSTR pwszFormatType

	// Size of cbAuthenticatorData.
	AuthenticatorDataLen uint32 // DWORD cbAuthenticatorData
	// Authenticator data that was created for this credential.
	// _Field_size_bytes_ (cbAuthenticatorData)
	AuthenticatorDataPtr *byte // PBYTE              pbAuthenticatorData

	// Size of CBOR encoded attestation information
	//0 => encoded as CBOR null value.
	AttestationLen uint32 // DWORD cbAttestation
	//Encoded CBOR attestation information
	// _Field_size_bytes_ (cbAttestation)
	AttestationPtr *byte // PBYTE              pbAttestation

	AttestationDecodeType uint32 // DWORD dwAttestationDecodeType
	// Following depends on the dwAttestationDecodeType
	//  WEBAUTHN_ATTESTATION_DECODE_NONE
	//      NULL - not able to decode the CBOR attestation information
	//  WEBAUTHN_ATTESTATION_DECODE_COMMON
	//      PWEBAUTHN_COMMON_ATTESTATION;
	AttestationDecode uintptr // PVOID pvAttestationDecode

	// The CBOR encoded Attestation Object to be returned to the RP.
	AttestationObjectLen uint32 // DWORD              cbAttestationObject
	// _Field_size_bytes_ (cbAttestationObject)
	AttestationObjectPtr *byte // PBYTE              pbAttestationObject

	// The CredentialId bytes extracted from the Authenticator Data.
	// Used by Edge to return to the RP.
	CredentialIDLen uint32 // DWORD              cbCredentialId
	// _Field_size_bytes_ (cbCredentialId)
	CredentialIDPtr *byte // PBYTE              pbCredentialId

	//
	// Following fields have been added in WEBAUTHN_CREDENTIAL_ATTESTATION_VERSION_2
	//

	Extensions *RawExtensions // WEBAUTHN_EXTENSIONS Extensions

	//
	// Following fields have been added in WEBAUTHN_CREDENTIAL_ATTESTATION_VERSION_3
	//

	// One of the WEBAUTHN_CTAP_TRANSPORT_* bits will be set corresponding to
	// the transport that was used.
	UsedTransport uint32 // DWORD dwUsedTransport

	//
	// Following fields have been added in WEBAUTHN_CREDENTIAL_ATTESTATION_VERSION_4
	//

	EpAtt              bool // BOOL bEpAtt
	LargeBlobSupported bool // BOOL  bLargeBlobSupported
	ResidentKey        bool // BOOL  bResidentKey
}
