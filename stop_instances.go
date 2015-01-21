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
	opts := compute.StopInstancesOptions{
		//Force: true,
		InstanceIds: []string{
			"9145d31f",
			"3fa4cac5",
		},
	}
	resp, err := client.StopInstances(&opts)
	if err != nil {
		log.Fatal(err)
	}

	instances := resp.StateChanges
	j, err := json.Marshal(instances)
	os.Stdout.Write(j)
}
