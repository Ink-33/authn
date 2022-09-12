package cmd

import (
	"encoding/base64"
	"fmt"
	"strings"
	"sync"

	"github.com/Ink-33/authn/api"
	"github.com/Ink-33/authn/api/utils"
)

type cred struct {
	cbCred  uint32
	pbCred  *byte
	b64Cred string
}

func DeletePlatformCred(c *api.WebAuthNClient) error {
	fmt.Println("Select sub operation:")
	fmt.Println("1:", "Purge all credentials related with this cli tool")
	fmt.Println("2:", "Choose credentials to delete")
	fmt.Println("0:", "Cancel")

	op := utils.ScanInputAndCheck()

	switch op {
	case 0:
		return nil
	case 1:
		return purge(c)
	case 2:
		return choose()
	}
	return nil
}

func purge(c *api.WebAuthNClient) error {
	printCallAPI()
	list, err := c.GetPlatformCredentialList("")
	if err != nil {
		if err.Error() == "NteNotFound" {
			fmt.Println("Nothing to do ...")
			return nil
		}
		return err
	}

	if len(list) == 0 {
		fmt.Println("Nothing to do...")
		return nil
	}

	printCredList(list)
	fmt.Printf("These credentials will be removed. Y/n?")
	q := ""
	fmt.Scanln(&q)

	if strings.ToLower(q) != "y" {
		return fmt.Errorf("Cancelled")
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
	return nil
}

func choose() error {
	return fmt.Errorf("TODO")
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
	fmt.Println("All done")
}
