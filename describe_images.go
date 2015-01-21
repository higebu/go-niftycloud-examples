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
	resp, err := client.Images(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	images := resp.Images
	j, err := json.Marshal(images)
	os.Stdout.Write(j)
}
