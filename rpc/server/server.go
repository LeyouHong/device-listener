package server

import (
	"encoding/json"
	"fmt"
	c "github.com/device-listener/conf"
	"github.com/device-listener/database"
	pb "github.com/device-listener/rpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type RpcServer struct{}

func (s *RpcServer) SignUp(ctx context.Context, request *pb.SignUpRequest) (*pb.Response, error) {
	reqJson, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	database.Put(c.Conf.GetDB(), []byte(request.Name+string(request.Ts)), reqJson)
	log.Println(string(reqJson))

	return &pb.Response{
		Res: "register successful",
	}, nil
}

func (s *RpcServer) GetDeviceInfo(ctx context.Context, request *pb.DataRequest) (*pb.Response, error) {
	m := make(map[string]string)

	database.Search(c.Conf.GetDB(), request.Key, request.Start, request.End, m)
	log.Println(m)

	resJson, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	return &pb.Response{
		Res: string(resJson),
	}, nil
}

func Start() {
	addr := fmt.Sprintf("%v:%v", c.Conf.GetIP(), c.Conf.GetPORT())
	log.Println(addr)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDeviceServer(grpcServer, &RpcServer{})
	grpcServer.Serve(lis)
}
