package main

import (
	"github.com/device-listener/api"
	"log"
	"strconv"
	"time"
)

func main() {
	request := &api.DBRequest{
		Address: "172.17.0.2:32770",
	}

	client, err := api.NewDBClient(request)
	if err != nil {
		log.Fatal(err)
	}

	res, err := api.GetDeviceInfo(client, "Device1", "1526941382", strconv.FormatInt(time.Now().Unix(), 10))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)

	//m := api.GetDevices([]byte("Device1"))
	//log.Println(m)

	api.CloseDBClient(client)
}
