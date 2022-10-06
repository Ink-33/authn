package cmd

import (
	"crypto/rand"
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

	return func() { fmt.Printf("b: %v\n", b) }, nil
}
