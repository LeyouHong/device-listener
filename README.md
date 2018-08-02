# About device-listener 
device-listener is RPC server can receive register and deregister infomations from devices client.
Data will be saved in leveldb.

# Download.
````
go get github.com/LeyouHong/device-listener
````

# build
## build on local
````
cd device-listener && make
````
target is build/device-listener

## build docker image
````
make build_linux_amd64 && make docker
docker run -it lhong/device-listener
````

## How to use API
``````
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
``````