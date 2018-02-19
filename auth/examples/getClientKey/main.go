package main

import (
	"fmt"
	"os"

	ovh "github.com/toorop/go-ovh"
	"github.com/toorop/go-ovh/auth"
	"github.com/toorop/go-ovh/auth/authswag"
)

func main() {

	authClient := auth.New(ovh.NewClient("ovh-eu").WithKeyringFromEnv())

	// Define access rules
	accessRules := authswag.PostAuthCredentialParamsBodyAccessRules{}

	// Access rule 1
	// Ressource /me
	// Methods: all
	accessRules = append(accessRules, &authswag.AuthAccessRule{
		Method: "GET",
		Path:   "/domain/*",
	})

	accessRules = append(accessRules, &authswag.AuthAccessRule{
		Method: "POST",
		Path:   "/domain/*",
	})

	accessRules = append(accessRules, &authswag.AuthAccessRule{
		Method: "PUT",
		Path:   "/domain/*",
	})

	body := authswag.PostAuthCredentialParamsBody{
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
