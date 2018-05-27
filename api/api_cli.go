package api

import (
	pb "github.com/device-listener/rpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
)

type SignUpClient struct {
	request    SignUpRequest
	connection *grpc.ClientConn
}

func NewSignUpClient(request *SignUpRequest) (*SignUpClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(request.Address, opts...)
	if err != nil {
		log.Fatal(err)
	}

	return &SignUpClient{request: *request, connection: conn}, nil
}

func CloseSignUpClient(client *SignUpClient) {
	defer client.connection.Close()
}

func Register(client *SignUpClient) (string, error) {
	return SignUp(client, "Register")
}

func DeRegister(client *SignUpClient) (string, error) {
	return SignUp(client, "DeRegister")
}

func SignUp(client *SignUpClient, action string) (string, error) {
	cli := pb.NewDeviceClient(client.connection)
	response, err := cli.SignUp(context.Background(), &pb.SignUpRequest{
		Id:           client.request.Id,
		Name:         client.request.Name,
		Addr:         client.request.Address,
		Capabilities: client.request.Capabilities,
		Action:       action,
		Ts:           strconv.FormatInt(time.Now().Unix(), 10),
		Services:     client.request.Services,
		Uuid:         client.request.Uuid,
	})

	return response.Res, err
}

type DBClient struct {
	request    DBRequest
	connection *grpc.ClientConn
}

func NewDBClient(request *DBRequest) (*DBClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(request.Address, opts...)
	if err != nil {
		log.Fatal(err)
	}

	return &DBClient{request: *request, connection: conn}, nil
}

func CloseDBClient(client *DBClient) {
	defer client.connection.Close()
}

func GetDeviceInfo(client *DBClient, key string, start string, end string) (string, error) {
	cli := pb.NewDeviceClient(client.connection)
	response, err := cli.GetDeviceInfo(context.Background(), &pb.DataRequest{
		Key:   key,
		Start: start,
		End:   end,
	})

	return response.Res, err
}
