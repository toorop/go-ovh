package main

import (
	"fmt"
	"os"

	ovh "github.com/toorop/go-ovh"
	"github.com/toorop/go-ovh/auth"
)

func main() {
	// Get AK & CK from env
	if len(os.Args) != 2 {
		fmt.Println("Usage: getck Application_Key")
		os.Exit(1)
	}
	ak := os.Args[1]
	authClient := auth.New(ovh.NewClient("ovh-eu").WithKeyring(ak, "", ""))

	// Define access rules
	accessRules := auth.PostAuthCredentialParamsBodyAccessRules{}

	// Access rule 1
	// Ressource /me
	// Methods: all
	accessRules = append(accessRules, &auth.AuthAccessRule{
		Method: "GET",
		Path:   "/me/*",
	})

	body := auth.PostAuthCredentialParamsBody{
		AccessRules: accessRules,
	}

	credentials, err := authClient.RequestCredential(body)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("Consumer Key (CK): %s\n", credentials.ConsumerKey)
	fmt.Printf("Validation URL: %s\n", credentials.ValidationURL)
}
