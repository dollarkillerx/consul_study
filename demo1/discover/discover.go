package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	consulApi "github.com/hashicorp/consul/api"
)

const (
	ConsulAddress = "192.168.88.11:8500"
	LocalIP       = "192.168.88.39"
	LocalPort     = 8181
)

func main() {
	config := consulApi.DefaultConfig()
	config.Address = ConsulAddress
	client, err := consulApi.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}

	node := consulApi.AgentServiceRegistration{
		ID:      uuid.New().String(), // 节点名称
		Name:    "Node1",             // 服务名称
		Port:    LocalPort,
		Tags:    []string{"sp1", "sp2"},
		Address: LocalIP,
	}

	check := consulApi.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d/heartbeat", LocalIP, LocalPort),
		Timeout:                        "5s",  // 超时时间
		Interval:                       "5s",  // 尝试间隔
		DeregisterCriticalServiceAfter: "10s", // 实效时间
	}

	node.Check = &check
	if err := client.Agent().ServiceRegister(&node); err != nil {
		log.Fatalln(err)
	}

	log.Println("Register Success")

	checkServer()
}

func checkServer() {
	http.HandleFunc("/heartbeat", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("success"))
	})

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", LocalIP, LocalPort), nil); err != nil {
		log.Fatalln(err)
	}
}
