package main

import (
	"fmt"

	"github.com/Ink-33/authn/api"
	"github.com/Ink-33/authn/api/raw"
	"github.com/Ink-33/authn/api/utils"
	"github.com/Ink-33/authn/cmd"
)

func main() {
	fmt.Printf("Cli tool version: %v\n", cmd.Version)
	fmt.Printf("WebAuthN API Version: %v\n", raw.GetAPIVersionNumber())
	fmt.Printf("Is User Verifying Platform Authenticator Available: %v\n", raw.IsUserVerifyingPlatformAuthenticatorAvailable())
	c := api.NewClient("go.webauthn.demo.app", "WebAuthN From Golang", "")
loop:
	fmt.Println("Select operation:")
	for i := range cmd.Actions {
		fmt.Println(cmd.Actions[i].Desp)
	}
	fmt.Println("0:", "Exit")

	op := utils.ScanInputAndCheck()

	if op == 0 {
		return
	}

	if op > len(cmd.Actions) || op < 0 {
		fmt.Println("Invaild input")
		println()
		goto loop
	}

	err := cmd.Actions[op-1].Function(c)
	if err != nil {
		fmt.Println("Err:", err)
	}

	println()
	goto loop
}
