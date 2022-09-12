package cmd

import (
	"crypto/rand"
	"fmt"

	"github.com/Ink-33/authn/api"
)

func GetAssertion(c *api.WebAuthNClient) error {
	id := make([]byte, 32)
	_, _ = rand.Read(id)
	
	printCallAPI()
	b, err := c.GetAssertion("local://demo.app", nil)
	if err != nil {
		return err
	}
	fmt.Printf("b: %v\n", b)
	return nil
}
