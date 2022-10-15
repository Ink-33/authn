package share

import (
	"unsafe"

	"github.com/Ink-33/authn/api/utils"
)

// RawCommonAttestation contains the common data for an attestation.
type RawCommonAttestation struct {
	// Version of this structure, to allow for modifications in the future.
	Version uint32 // DWORD dwVersion;

	// Hash and Padding Algorithm
	//
	// The following won't be set for "fido-u2f" which assumes "ES256".
	AlgString *uint16 // PCWSTR pwszAlg;
	Alg       int32   // LONG lAlg;      // COSE algorithm

	// Signature that was generated for this attestation.
	SignatureLen uint32 // DWORD cbSignature;
	// _Field_size_bytes_(cbSignature)
	SignaturePtr *byte // PBYTE pbSignature;

	// Following is set for Full Basic Attestation. If not, set then, this is Self Attestation.
	// Array of X.509 DER encoded certificates. The first certificate is the signer, leaf certificate.
	X5CLen uint32 // DWORD cX5c;
	// _Field_size_(cX5c)
	X5CPtr *RawX5C // PWEBAUTHN_X5C pX5c;

	// Following are also set for tpm
	Ver         *uint16 //  PCWSTR pwszVer; // L"2.0"
	CertInfoLen uint32  // DWORD cbCertInfo;
	// _Field_size_bytes_(cbCertInfo)
	CertInfoPtr *byte  // PBYTE pbCertInfo;
	PubAreaLen  uint32 // DWORD cbPubArea;
	// _Field_size_bytes_(cbPubArea)
	PubAreaPtr *byte // PBYTE pbPubArea;
}

// CommonAttestation contains the common data for an attestation.
type CommonAttestation struct {
	// Version of this structure, to allow for modifications in the future.
	Version uint32 // DWORD dwVersion;

	// Hash and Padding Algorithm
	//
	// The following won't be set for "fido-u2f" which assumes "ES256".
	AlgString string // PCWSTR pwszAlg;
	Alg       int32  // LONG lAlg;      // COSE algorithm

	// Signature that was generated for this attestation.
	Signature []byte // PBYTE pbSignature;

	// Following is set for Full Basic Attestation. If not, set then, this is Self Attestation.
	// Array of X.509 DER encoded certificates. The first certificate is the signer, leaf certificate.
	X5C []X5C // PWEBAUTHN_X5C pX5c;

	// Following are also set for tpm
	Ver      string //  PCWSTR pwszVer; // L"2.0"
	CertInfo []byte // PBYTE pbCertInfo;
	PubArea  []byte // PBYTE pbPubArea;
}

func (c *RawCommonAttestation) DeRaw() *CommonAttestation {
	if c == nil {
		return nil
	}
	rx5c := unsafe.Slice(c.X5CPtr, c.X5CLen)
	x5c := make([]X5C, c.X5CLen)
	for i := 0; i < int(c.X5CLen); i++ {
		x5c[i] = rx5c[i].DeRaw()
	}

	return &CommonAttestation{
		Version:   c.Version,
		AlgString: utils.UTF16PtrtoString(c.AlgString),
		Alg:       c.Alg,
		Signature: utils.BytesBuilder(c.SignaturePtr, c.SignatureLen),
		X5C:       x5c,
		Ver:       utils.UTF16PtrtoString(c.Ver),
		CertInfo:  utils.BytesBuilder(c.SignaturePtr, c.SignatureLen),
		PubArea:   utils.BytesBuilder(c.PubAreaPtr, c.PubAreaLen),
	}
}
