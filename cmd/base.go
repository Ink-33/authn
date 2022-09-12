package cmd

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/Ink-33/authn/api"
	"github.com/Ink-33/authn/api/share"
)

type Action struct {
	Desp     string
	Function func(*api.WebAuthNClient) error
}

var Actions = []Action{
	{
		Desp:     strings.Join([]string{"1:", "Make Credential"}, " "),
		Function: MakeCred,
	},
	{
		Desp:     strings.Join([]string{"2:", "Get Assertion"}, " "),
		Function: GetAssertion,
	},
	{
		Desp:     strings.Join([]string{"3:", "Get Platform Credential List"}, " "),
		Function: GetPlatformCredList,
	},
	{
		Desp:     strings.Join([]string{"4:", "Delete Platform Credential"}, " "),
		Function: DeletePlatformCred,
	},
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

func printCallAPI() {
	fmt.Println("Calling win32 api ...")
}

func printCredList(list []*share.CredentialDetails) {
	for i, v := range list {
		fmt.Printf("[%v]\tCredID:\t%v\n\tUser:\t%v[%v]\n\tRP:\t%v[%v]\n\tRemovable:\t%v\n\n",
			i,
			base64.RawURLEncoding.EncodeToString(v.CredentialID),
			v.UserInformation.Name,
			v.UserInformation.DisplayName,
			v.RPInformation.Name,
			v.RPInformation.ID,
			v.Removable,
		)
	}
}
