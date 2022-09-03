package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"unsafe"

	"github.com/Ink-33/authn/api"
	"github.com/Ink-33/authn/utils"
	"github.com/fxamacker/cbor/v2"
)

func main() {
	c := api.NewClient("go.webauthn.demo.app", "WebAuthN From Golang", "")
	id := make([]byte, 32)
	_, _ = rand.Read(id)
	u := &testUser{id}
	a, err := c.MakeCredential(u, "local://demo.app", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("a.Version: %v\n", a.Version)
	fmt.Printf("a.FormatType: %v\n", utils.UTF16toString(a.FormatType))
	fmt.Printf("a.AttestationDecodeType: %v\n", a.AttestationDecodeType)
	fmt.Printf("a.UsedTransport: %v\n", a.UsedTransport)
	fmt.Printf("a.EpAtt: %v\n", a.EpAtt)
	fmt.Printf("a.LargeBlobSupported: %v\n", a.LargeBlobSupported)
	fmt.Printf("a.ResidentKey: %v\n", a.ResidentKey)

	fmt.Printf("a.Extensions: %v\n", a.Extensions)

	fmt.Printf("AuthenticatorData: %v\n",
		base64.RawStdEncoding.EncodeToString(unsafe.Slice(a.AuthenticatorDataPtr, a.AuthenticatorDataLen)))

	atM := map[string]any{}
	err = cbor.Unmarshal(unsafe.Slice(a.AttestationPtr, a.AttestationLen), &atM)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Attestation: %v\n", atM)

	atoM := map[string]any{}
	err = cbor.Unmarshal(unsafe.Slice(a.AttestationObjectPtr, a.AttestationObjectLen), &atoM)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AttestationObject: %v\n", atoM)

	fmt.Printf("CredentialID: %v\n",
		base64.RawURLEncoding.EncodeToString(unsafe.Slice(a.CredentialIDPtr, a.CredentialIDLen)))

}

type testUser struct {
	id []byte
}

func (user *testUser) GetID() []byte {
	return user.id
}

func (user *testUser) GetName() string {
	return "test@example.com"
}

func (user *testUser) GetDisplayName() string {
	return "Test User"
}

func (user *testUser) GetIcon() string {
	return ""
}
