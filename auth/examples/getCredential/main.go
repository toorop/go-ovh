package main

import (
	"log"

	ovh "github.com/toorop/go-ovh"
	"github.com/toorop/go-ovh/auth"
)

func main() {
	credential, err := auth.New(ovh.NewClient("ovh-eu").WithKeyringFromEnv()).GetCurrentCredential()
	log.Println(credential, err)
}
