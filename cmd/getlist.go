package cmd

import (
	"github.com/Ink-33/authn/api"
)

func GetPlatformCredList(c *api.WebAuthNClient) error {
	printCallAPI()
	res, err := c.GetPlatformCredentialList("-")
	if err != nil {
		return err
	}
	printCredList(res)
	return nil
}
