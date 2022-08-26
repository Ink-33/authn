package define

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

	WebAuthNAuthenticatorMakeCredentialOptionsVersion1       = 1
	WebAuthNAuthenticatorMakeCredentialOptionsVersion2       = 2
	WebAuthNAuthenticatorMakeCredentialOptionsVersion3       = 3
	WebAuthNAuthenticatorMakeCredentialOptionsVersion4       = 4
	WebAuthNAuthenticatorMakeCredentialOptionsVersion5       = 5
	WebAuthNAuthenticatorMakeCredentialOptionsCurrentVersion = WebAuthNAuthenticatorMakeCredentialOptionsVersion5
)

// version
const (
	WebAuthNRPEntityInformationCurrentVersion     = 1
	WebAuthNUserEntityInformationCurrentVersion   = 1
	WebAuthNClientDataCurrentVersion              = 1
	WebAuthNCOSECredentialParameterCurrentVersion = 1
	WebAuthNCredentialCurrentVersion              = 1
	WebAuthNCredentialDetailsCurrentVersion       = 1
	WebAuthNGetCredentialsOptionsCurrentVersion   = 1
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
