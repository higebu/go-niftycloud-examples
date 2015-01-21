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
	resp, err := client.DescribeAvailabilityZones(nil)
	if err != nil {
		log.Fatal(err)
	}

	zones := resp.Zones
	j, err := json.Marshal(zones)
	os.Stdout.Write(j)
}
