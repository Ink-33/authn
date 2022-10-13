package cmd

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/Ink-33/authn/api"
)

func MakeCred(c *api.WebAuthNClient) (func(), error) {
	id := make([]byte, 32)
	_, _ = rand.Read(id)
	u := &testUser{id}

	printCallAPI()
	a, err := c.MakeCredential(u, "local://demo.app", nil)
	if err != nil {
		return nil, err
	}

	p := func() {
		// fmt.Printf("Version: %v\n", a.Version)
		// fmt.Printf("FormatType: %v\n", a.FormatType)
		// fmt.Printf("AttestationDecodeType: %v\n", a.AttestationDecode)
		// fmt.Printf("UsedTransport: %v\n", a.UsedTransport)
		// fmt.Printf("EpAtt: %v\n", a.EpAtt)
		// fmt.Printf("LargeBlobSupported: %v\n", a.LargeBlobSupported)
		// fmt.Printf("ResidentKey: %v\n", a.ResidentKey)

		// fmt.Printf("Extensions: %v\n", a.Extensions)

		fmt.Printf("RPID Hash: %v\n",
			base64.RawURLEncoding.EncodeToString(a.AuthenticatorData.RPIDHash))

		fmt.Printf("User Present: %v\n", a.AuthenticatorData.Flags.UserPresent)
		fmt.Printf("RFU1: %v\n", a.AuthenticatorData.Flags.RFU1)
		fmt.Printf("User Verified: %v\n", a.AuthenticatorData.Flags.UserVerified)
		fmt.Printf("Backup Eligibility: %v\n", a.AuthenticatorData.Flags.BackupEligibility)
		fmt.Printf("Backup State: %v\n", a.AuthenticatorData.Flags.BackupState)
		fmt.Printf("RFU2: %v\n", a.AuthenticatorData.Flags.RFU2)
		fmt.Printf("Attested credential data included: %v\n", a.AuthenticatorData.Flags.AttestedCredentialData)
		fmt.Printf("Extension data included: %v\n", a.AuthenticatorData.Flags.ExtensionData)

		fmt.Printf("Sign Counter: %v\n", a.AuthenticatorData.SignCounter)

		fmt.Printf("AAGUID: %v\n",
			base64.RawURLEncoding.EncodeToString(a.AuthenticatorData.AttestedCredentialData.AAGUID))

		fmt.Printf("CredentialID:%v\n",
			base64.RawURLEncoding.EncodeToString(a.AuthenticatorData.AttestedCredentialData.CredentialID))

		fmt.Printf("COSE: %v\n", a.AuthenticatorData.AttestedCredentialData.CredentialPublicKey)
		// atM := map[string]any{}
		// err = cbor.Unmarshal(a.Attestation, &atM)
		// if err != nil {
		// 	fmt.Printf("Err: %v\n", err)
		// 	return
		// }
		// fmt.Printf("Attestation: %v\n", atM)

		// atoM := map[string]any{}
		// err = cbor.Unmarshal(a.AttestationObject, &atoM)
		// if err != nil {
		// 	fmt.Printf("Err: %v\n", err)
		// 	return
		// }
		// fmt.Printf("AttestationObject: %v\n", atoM)

	}
	return p, nil
}
