package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/higebu/go-niftycloud/compute"
	"github.com/higebu/go-niftycloud/niftycloud"
)

func main() {
	auth, err := niftycloud.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}
	client := compute.New(auth, niftycloud.JPEast)
	resp, err := client.DescribeInstances(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	reservations := resp.Reservations
	j, err := json.Marshal(reservations)
	os.Stdout.Write(j)
}
