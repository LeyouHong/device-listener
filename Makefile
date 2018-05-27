PACKAGES = $(shell go list ./... | grep -v '/vendor/')  

build:
	go build -o build/device-listener ./main.go
	go build -o build/concurrent_demo github.com/device-listener/test/concurrent_demo
	go build -o build/signup_demo github.com/device-listener/test/signup_demo
	go build -o build/database_demo github.com/device-listener/test/database_demo

docker: 
	docker build -t device-listener .

test:
	@echo "====> Running go test"
	@go test $(PACKAGES)

clean:
	rm -rf build
	rm -rf database/db

.PHONY: test

