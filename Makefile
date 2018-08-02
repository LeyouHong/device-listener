PACKAGES = $(shell go list ./... | grep -v '/vendor/')  

build:
	go build -o build/device-listener ./main.go
	go build -o build/concurrent_demo github.com/LeyouHong/device-listener/test/concurrent_demo
	go build -o build/signup_demo github.com/LeyouHong/device-listener/test/signup_demo
	go build -o build/database_demo github.com/LeyouHong/device-listener/test/database_demo

build_linux_amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -o build/device-listener

docker: 
	docker build -t lhong/device-listener .

test:
	@echo "====> Running go test"
	@go test $(PACKAGES)

clean:
	rm -rf build
	rm -rf database/db

.PHONY: test

