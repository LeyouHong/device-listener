package proto

import (
	"github.com/golang/protobuf/proto"
	"log"
	"testing"
)

func TestDeviceProto(t *testing.T) {
	dev := SignUpRequest{
		Id:     		123,
		Name:   		"Device1",
		Addr:			"172.17.0.2:8080",
		Capabilities : 	"for testing",
		Action: 		"register",
		Ts:     		"1526668370",
		Services:   	"Demo_service",
		Uuid:			"12345",
	}

	mData, err := proto.Marshal(&dev)
	if err != nil {
		log.Fatal(err)
	}

	var umData SignUpRequest
	err = proto.Unmarshal(mData, &umData)
	if err != nil {
		log.Fatal(err)
	}
}
