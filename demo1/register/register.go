package main

import (
	"encoding/json"
	"log"

	consulApi "github.com/hashicorp/consul/api"
)

const (
	ConsulAddress = "192.168.88.11:8500"
)

func main() {
	config := consulApi.DefaultConfig()
	config.Address = ConsulAddress

	client, err := consulApi.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}
	service, meta, err := client.Health().Service("Node1", "", true, nil)
	if err != nil {
		log.Fatalln(err)
	}

	marshal, err := json.Marshal(service)
	if err == nil {
		log.Println(string(marshal))
	}
	log.Println(meta)
}
