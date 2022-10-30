package main

import (
	"fmt"

	"github.com/Ink-33/authn/api"
	"github.com/Ink-33/authn/api/raw"
	"github.com/Ink-33/authn/cmd"
	"github.com/Ink-33/authn/interact"
)

func main() {
	fmt.Printf("Cli tool version: %v\n", cmd.Version)
	fmt.Printf("WebAuthN API Version: %v\n", raw.GetAPIVersionNumber())
	fmt.Printf("Is User Verifying Platform Authenticator Available: %v\n\n", raw.IsUserVerifyingPlatformAuthenticatorAvailable())
	c := api.NewClient("go.webauthn.demo.app", "WebAuthN From Golang", "")

	choices := interact.Choose{
		Title: "Select operation:",
		Choices: []interact.Choice{
			interact.NewChoice(
				"Make Credential",
				func() (func(), error) { return cmd.MakeCred(c) },
			),
			interact.NewChoice(
				"Show Authenticator Attestation Certificate",
				func() (func(), error) { return cmd.ShowCertInfo(c) },
			),
			interact.NewChoice(
				"Get Assertion",
				func() (func(), error) { return cmd.GetAssertion(c) },
			),
			interact.NewChoice(
				"Get Platform Credential List",
				func() (func(), error) { return cmd.GetPlatformCredList(c) },
			),
			interact.NewChoice(
				"Delete Platform Credential",
				func() (func(), error) { return cmd.DeletePlatformCred(c) },
			),
		},
		Loop: true,
	}
	_, err := choices.Do()
	if err != nil {
		panic(err)
	}
}
