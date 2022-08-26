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
	WebAuthNCoseCredentialParameterCurrentVersion = 1
	WebAuthNCredentialCurrentVersion              = 1
)

// credential parameters.
const (
	WebAuthNCredentialTypePublicKey = "public-key"

	WebAuthNCoseAlgorithmECDSAP256WithSHA256 = -7
	WebAuthNCoseAlgorithmECDSAP384WithSHA384 = -35
	WebAuthNCoseAlgorithmECDSAP521WithSHA512 = -36

	WebAuthNCoseAlgorithmRSASSAPKCS1V15WithSHA256 = -257
	WebAuthNCoseAlgorithmRSASSAPKCS1V15WithSHA384 = -258
	WebAuthNCoseAlgorithmRSASSAPKCS1V15WithSHA512 = -259

	WebAuthNCoseAlgorithmRSAPSSWithSHA256 = -37
	WebAuthNCoseAlgorithmRSAPSSWithSHA384 = -38
	WebAuthNCoseAlgorithmRSAPSSWithSHA512 = -39
)

// client data
const (
	WebAuthNHashAlgorithmSHA256 = "SHA-256"
	WebAuthNHashAlgorithmSHA384 = "SHA-384"
	WebAuthNHashAlgorithmSHA512 = "SHA-512"
)


