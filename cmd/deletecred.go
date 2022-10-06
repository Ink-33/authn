package cmd

import (
	"encoding/base64"
	"fmt"
	"strings"
	"sync"

	"github.com/Ink-33/authn/api"
	"github.com/Ink-33/authn/interact"
)

type cred struct {
	cbCred  uint32
	pbCred  *byte
	b64Cred string
}

func DeletePlatformCred(c *api.WebAuthNClient) (func(), error) {
	choices := interact.Choose{
		Title: "Select sub operation:",
		Choices: []interact.Choice{
			interact.NewChoice(
				"Purge all credentials related with this cli tool",
				func() (func(), error) { msg, err := purge(c); return msg, interact.NewToPreviouswithErr(err) },
			),
			interact.NewChoice(
				"Choose credentials to delete",
				func() (func(), error) { msg, err := choose(c); return msg, interact.NewToPreviouswithErr(err) },
			),
		},
		Loop:                 true,
		ToPreviousChooseDesc: "Cancel",
	}
	return choices.Do()
}

func purge(c *api.WebAuthNClient) (func(), error) {
	printCallAPI()
	list, err := c.GetPlatformCredentialList("")
	if err != nil {
		if err.Error() == "NteNotFound" {
			return func() { fmt.Println("Nothing to do ...") }, nil
		}
		return nil, err
	}

	if len(list) == 0 {
		return func() { fmt.Println("Nothing to do ...") }, nil
	}

	printCredList(list)
	fmt.Printf("These credentials will be removed. Y/n?")
	q := ""
	fmt.Scanln(&q)

	if strings.ToLower(q) != "y" {
		return nil, fmt.Errorf("Cancelled")
	}

	deletelist := make([]cred, len(list))
	for i := 0; i < len(list); i++ {
		deletelist[i] = cred{
			cbCred:  uint32(len(list[i].CredentialID)),
			pbCred:  &list[i].CredentialID[0],
			b64Cred: base64.URLEncoding.EncodeToString(list[i].CredentialID),
		}
	}
	delete(c, deletelist)
	return func() { fmt.Println("All done") }, nil
}

func choose(c *api.WebAuthNClient) (func(), error) {
	return nil, fmt.Errorf("TODO")
}

func delete(c *api.WebAuthNClient, list []cred) {
	wg := sync.WaitGroup{}
	wg.Add(len(list))

	bucket := make(chan struct{}, 3)
	for i := 0; i < 3; i++ {
		bucket <- struct{}{}
	}

	for i := 0; i < len(list); i++ {
		<-bucket
		go func(i int) {
			fmt.Printf("Deleting %v ...\n", list[i].b64Cred)
			err := c.DeletePlatformCred(list[i].cbCred, list[i].pbCred)
			if err != nil {
				fmt.Println("Err:", err)
			}
			wg.Done()
			bucket <- struct{}{}
		}(i)
	}
	wg.Wait()
}
