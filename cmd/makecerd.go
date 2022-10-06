package cmd

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/Ink-33/authn/api"
	"github.com/fxamacker/cbor/v2"
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
		fmt.Printf("a.Version: %v\n", a.Version)
		fmt.Printf("a.FormatType: %v\n", a.FormatType)
		fmt.Printf("a.AttestationDecodeType: %v\n", a.AttestationDecode)
		fmt.Printf("a.UsedTransport: %v\n", a.UsedTransport)
		fmt.Printf("a.EpAtt: %v\n", a.EpAtt)
		fmt.Printf("a.LargeBlobSupported: %v\n", a.LargeBlobSupported)
		fmt.Printf("a.ResidentKey: %v\n", a.ResidentKey)

		fmt.Printf("a.Extensions: %v\n", a.Extensions)

		fmt.Printf("AuthenticatorData: %v\n",
			base64.RawStdEncoding.EncodeToString(a.AuthenticatorData))

		atM := map[string]any{}
		err = cbor.Unmarshal(a.Attestation, &atM)
		if err != nil {
			fmt.Printf("Err: %v\n", err)
			return
		}
		fmt.Printf("Attestation: %v\n", atM)

		atoM := map[string]any{}
		err = cbor.Unmarshal(a.AttestationObject, &atoM)
		if err != nil {
			fmt.Printf("Err: %v\n", err)
			return
		}
		fmt.Printf("AttestationObject: %v\n", atoM)

		fmt.Printf("CredentialID: %v\n",
			base64.RawURLEncoding.EncodeToString(a.CredentialID))
	}
	return p, nil
}
