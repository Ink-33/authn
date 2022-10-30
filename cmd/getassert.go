package cmd

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/Ink-33/authn/api"
)

func GetAssertion(c *api.WebAuthNClient) (func(), error) {
	id := make([]byte, 32)
	_, _ = rand.Read(id)

	printCallAPI()
	b, err := c.GetAssertion("local://demo.app", nil)
	if err != nil {
		return nil, err
	}

	return func() {
		fmt.Printf("RPID Hash: %v\n",
			base64.RawURLEncoding.EncodeToString(b.AuthenticatorData.RPIDHash[:]))

		fmt.Printf("User Present: %v\n", b.AuthenticatorData.Flags.UserPresent)
		fmt.Printf("RFU1: %v\n", b.AuthenticatorData.Flags.RFU1)
		fmt.Printf("User Verified: %v\n", b.AuthenticatorData.Flags.UserVerified)
		fmt.Printf("Backup Eligibility: %v\n", b.AuthenticatorData.Flags.BackupEligibility)
		fmt.Printf("Backup State: %v\n", b.AuthenticatorData.Flags.BackupState)
		fmt.Printf("RFU2: %v\n", b.AuthenticatorData.Flags.RFU2)
		fmt.Printf("Attested credential data included: %v\n", b.AuthenticatorData.Flags.AttestedCredentialData)
		fmt.Printf("Extension data included: %v\n", b.AuthenticatorData.Flags.ExtensionData)

		fmt.Printf("UserID: %v\n",
			base64.RawURLEncoding.EncodeToString(b.UserID))
		fmt.Printf("CredentialID: %v\n",
			base64.RawURLEncoding.EncodeToString(b.Credential.ID))
		fmt.Printf("Signature: %v\n",
			base64.RawURLEncoding.EncodeToString(b.Signature))

		fmt.Printf("Sign Counter: %v\n", b.AuthenticatorData.SignCounter)

	}, nil
}
