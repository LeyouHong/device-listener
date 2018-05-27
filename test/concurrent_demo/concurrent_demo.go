package main

import (
	"fmt"
	pb "github.com/device-listener/rpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func test(ts string, i int) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("172.17.0.2:32770", opts...)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewDeviceClient(conn)
	response, err := client.SignUp(context.Background(), &pb.SignUpRequest{
		Id:           123,
		Name:         "Device",
		Addr:         "172.17.0.2:8080",
		Capabilities: "for testing",
		Action:       "register",
		Ts:           ts,
		Services:     "Demo_service",
		Uuid:         "12345",
	})
	fmt.Print(response)
	fmt.Println(i)
}

func main() {

	ts := "152666"
	tail := 8378
	for i := 0; i < 1000; i++ {
		tail++
		go test(ts+string(tail), i)
	}

	time.Sleep(10 * time.Minute)
}
