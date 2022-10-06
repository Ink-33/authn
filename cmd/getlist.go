package cmd

import (
	"github.com/Ink-33/authn/api"
)

func GetPlatformCredList(c *api.WebAuthNClient) (func(), error) {
	printCallAPI()
	res, err := c.GetPlatformCredentialList("-")
	if err != nil {
		return nil, err
	}

	return func() { printCredList(res) }, nil
}
