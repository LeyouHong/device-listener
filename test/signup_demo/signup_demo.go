package main

import (
	"github.com/device-listener/api"
	"log"
)

func main() {
	request := &api.SignUpRequest{
		Id:           123,
		Name:         "Device1",
		Address:      "172.17.0.2:32770",
		Capabilities: "xxx",
		Services:     "DemoService",
		Uuid:         "54321",
	}

	client, err := api.NewSignUpClient(request)
	if err != nil {
		log.Fatal(err)
	}

	res, err := api.Register(client)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)

	res, err = api.DeRegister(client)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)

	api.CloseSignUpClient(client)
}
