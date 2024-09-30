build:
	go build -o app ./cmd/app


docker-build:
	docker build -t usdt-grpc-service .

run:
	./app

test:
	go test -v ./...

lint:
	golangci-lint run