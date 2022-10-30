package cmd

import (
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"fmt"
	"strings"

	"github.com/Ink-33/authn/api"
	"github.com/Ink-33/authn/api/define"
	"github.com/Ink-33/authn/api/share"
)

func ShowCertInfo(c *api.WebAuthNClient) (func(), error) {
	id := make([]byte, 32)
	_, _ = rand.Read(id)
	u := &testUser{id}

	printCallAPI()
	opts := api.NewMakeCredOpts()
	opts.AttestationConveyancePreference = define.WebAuthNAttestationConveyancePreferenceDirect
	opts.AuthenticatorAttachment = define.WebAuthNAuthenticatorAttachmentAny
	a, err := c.MakeCredential(u, "local://demo.app", opts)
	if err != nil {
		return nil, err
	}
	return func() { printCertInfo(a.AttestationDecode.X5C) }, nil
}

func printCertInfo(x5c []share.X5C) {
	for i := range x5c {
		fmt.Printf("\nCertificate %v:\n", i+1)
		cert, err := x509.ParseCertificate(x5c[i])
		if err != nil {
			fmt.Printf("cert: %v", err)
		}
		fmt.Printf("Version: %v (%#x)\n", cert.Version, cert.Version)
		fmt.Printf("Serial Number: %v (%#x)\n", cert.SerialNumber, cert.SerialNumber)
		fmt.Printf("Signature Algorithm: %v\n", cert.SignatureAlgorithm)
		fmt.Printf("Issuer: %v\n", cert.Issuer)
		fmt.Printf("Validity:\n\tNot Before: %v\n\tNot After: %v\n",
			cert.NotBefore, cert.NotAfter)
		fmt.Printf("Subject: %v\n", cert.Subject)

		ppk := &pkixPublicKey{}
		_, err = asn1.Unmarshal(cert.RawSubjectPublicKeyInfo, ppk)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		pksize := 0
		name := ""
		switch pk := cert.PublicKey.(type) {
		case *rsa.PublicKey:
			pksize = pk.Size()
			name = "rsa"
		case *dsa.PublicKey:
			pksize = int(pk.Parameters.Q.Int64()) // maybe?
			name = "dsa"
		case *ecdsa.PublicKey:
			pksize = pk.Params().BitSize
			name = pk.Curve.Params().Name
		case ed25519.PublicKey:
			pksize = len(pk) // maybe?
			name = "ed25519"
		default:

		}

		fmt.Printf("Subject Public Key Info:\n\tPublic-Key: (%v bit)\n\tpub:\n%v\n\tAlgorithm Identifier: %v (%v)\n",
			pksize, buildBytesStr(ppk.BitString.Bytes, "\t\t"), ppk.Algo.Algorithm.String(), name)

		fmt.Printf("Signature Algorithm: %v\n\n%v", cert.SignatureAlgorithm, buildBytesStr(cert.Signature, "\t"))
	}
}

func buildBytesStr(b []byte, sep string) (str string) {
	l := len(b) / 18
	rest := func() int {
		if len(b)%18 != 0 {
			return 1
		}
		return 0
	}()

	strs := make([]string, l+rest+2)
	for i := 0; i < len(strs)-2; i++ {
		bs := b[i*18 : func() int {
			if (i+1)*18 > len(b) {
				return len(b)
			}
			return (i + 1) * 18
		}()]
		ss := make([]string, len(bs))
		for k := range bs {
			ss[k] = fmt.Sprintf("%02x", bs[k])
		}
		strs[i+1] = strings.Join(ss, ":") + "\n"
	}
	return strings.Join(strs, sep)
}

type pkixPublicKey struct {
	Algo      pkix.AlgorithmIdentifier
	BitString asn1.BitString
}
