package main

import (
	"fmt"
	"log"

	ovh "github.com/toorop/go-ovh"
	"github.com/toorop/go-ovh/domain"
)

func main() {
	ovh := domain.New(ovh.NewClient("ovh-eu").WithKeyringFromEnv())

	/*zones, err := ovh.GetZones()
	if err != nil {
		log.Fatalf("unable to get zones: %v", err)
	}
	fmt.Println(zones)*/

	zone, err := ovh.GetZone("respublica.fr")
	if err != nil {
		log.Fatalf("unable to get zone: %v", err)
	}
	fmt.Println(zone)

}
