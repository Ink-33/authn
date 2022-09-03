package define

// Ceremony type.
const (
	CollectedClientCeremonyCreate = "webauthn.create"
	CollectedClientCeremonyGet    = "webauthn.get"
)

// Options.
const (
	WebAuthNAuthenticatorAttachmentAny                = 0
	WebAuthNAuthenticatorAttachmentPlatform           = 1
	WebAuthNAuthenticatorAttachmentCrossPlatform      = 2
	WebAuthNAuthenticatorAttachmentCrossPlatformU2Fv2 = 3

	WebAuthNUserVerificationRequirementAny         = 0
	WebAuthNUserVerificationRequirementRequired    = 1
	WebAuthNUserVerificationRequirementPreferred   = 2
	WebAuthNUserVerificationRequirementDiscouraged = 3

	WebAuthNAttestationConveyancePreferenceAny      = 0
	WebAuthNAttestationConveyancePreferenceNone     = 1
	WebAuthNAttestationConveyancePreferenceIndirect = 2
	WebAuthNAttestationConveyancePreferenceDirect   = 3

	WebAuthNEnterpriseAttestationNone              = 0
	WebAuthNEnterpriseAttestationVendorFacilitated = 1
	WebAuthNEnterpriseAttestationPlatformManaged   = 2

	WebAuthNLargeBlobSupportNone      = 0
	WebAuthNLargeBlobSupportRequired  = 1
	WebAuthNLargeBlobSupportPreferred = 2
)

// version
const (
	WebAuthNRPEntityInformationCurrentVersion     = 1
	WebAuthNUserEntityInformationCurrentVersion   = 1
	WebAuthNClientDataCurrentVersion              = 1
	WebAuthNCOSECredentialParameterCurrentVersion = 1
	WebAuthNCredentialCurrentVersion              = 1
	WebAuthNCredentialDetailsCurrentVersion       = 1

	WebAuthNCredentialAttestationVersion1       = 1
	WebAuthNCredentialAttestationVersion2       = 2
	WebAuthNCredentialAttestationVersion3       = 3
	WebAuthNCredentialAttestationVersion4       = 4
	WebAuthNCredentialAttestationCurrentVersion = WebAuthNCredentialAttestationVersion4

	WebAuthNCredentialEXCurrentVersion = 1

	WebAuthNGetCredentialsOptionsVersion1       = 1
	WebAuthNGetCredentialsOptionsCurrentVersion = WebAuthNGetCredentialsOptionsVersion1

	WebauthnCredentialDetailsVersion1       = 1
	WebauthnCredentialDetailsCurrentVersion = WebauthnCredentialDetailsVersion1

	WebAuthNAuthenticatorMakeCredentialOptionsVersion1       = 1
	WebAuthNAuthenticatorMakeCredentialOptionsVersion2       = 2
	WebAuthNAuthenticatorMakeCredentialOptionsVersion3       = 3
	WebAuthNAuthenticatorMakeCredentialOptionsVersion4       = 4
	WebAuthNAuthenticatorMakeCredentialOptionsVersion5       = 5
	WebAuthNAuthenticatorMakeCredentialOptionsCurrentVersion = WebAuthNAuthenticatorMakeCredentialOptionsVersion5

	WebAuthNAuthenticatorGetAssertionOptionsVersion1       = 1
	WebAuthNAuthenticatorGetAssertionOptionsVersion2       = 2
	WebAuthNAuthenticatorGetAssertionOptionsVersion3       = 3
	WebAuthNAuthenticatorGetAssertionOptionsVersion4       = 4
	WebAuthNAuthenticatorGetAssertionOptionsVersion5       = 5
	WebAuthNAuthenticatorGetAssertionOptionsVersion6       = 6
	WebAuthNAuthenticatorGetAssertionOptionsCurrentVersion = WebAuthNAuthenticatorGetAssertionOptionsVersion6
)

// credential parameters.
const (
	WebAuthNCredentialTypePublicKey = "public-key"

	WebAuthNCOSEAlgorithmECDSAP256WithSHA256 = -7
	WebAuthNCOSEAlgorithmECDSAP384WithSHA384 = -35
	WebAuthNCOSEAlgorithmECDSAP521WithSHA512 = -36

	WebAuthNCOSEAlgorithmRSASSAPKCS1V15WithSHA256 = -257
	WebAuthNCOSEAlgorithmRSASSAPKCS1V15WithSHA384 = -258
	WebAuthNCOSEAlgorithmRSASSAPKCS1V15WithSHA512 = -259

	WebAuthNCOSEAlgorithmRSAPSSWithSHA256 = -37
	WebAuthNCOSEAlgorithmRSAPSSWithSHA384 = -38
	WebAuthNCOSEAlgorithmRSAPSSWithSHA512 = -39
)

// client data
const (
	WebAuthNHashAlgorithmSHA256 = "SHA-256"
	WebAuthNHashAlgorithmSHA384 = "SHA-384"
	WebAuthNHashAlgorithmSHA512 = "SHA-512"
)

// credential attestation
const (
	WebAuthNAttestationTypePacked  = "packed"
	WebAuthNAttestationTypeU2F     = "fido-u2f"
	WebAuthNAttestationTypeTPM     = "tpm"
	WebAuthNAttestationTypeNone    = "none"
	WebAuthAttestationDecodeNone   = 0
	WebAuthAttestationDecodeCommon = 1
	WebAuthAttestationVerTPM20     = "2.0"
)

// credential extra information (Transports)
const (
	WebAuthNCTAPTransportUSB       = 0x00000001
	WebAuthNCTAPTransportNFC       = 0x00000002
	WebAuthNCTAPTransportBLE       = 0x00000004
	WebAuthNCTAPTransportTest      = 0x00000008
	WebAuthNCTAPTransportInternal  = 0x00000010
	WebAuthNCTAPTransportFlagsMask = 0x0000001F
)

// get attestation options
const (
	WebAuthNCerdLargeBlobOperationNone   = 0
	WebAuthNCerdLargeBlobOperationGet    = 1
	WebAuthNCerdLargeBlobOperationSet    = 2
	WebAuthNCerdLargeBlobOperationDelete = 3

	/*
	   Information about flags.
	*/
	WebAuthNAuthenticatorHMACSecretValuesFlag = 0x00100000
)
