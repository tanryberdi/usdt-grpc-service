build:
	go build -o usdt-grpc-service ./cmd/app

docker-build:
	docker build -t usdt-grpc-service .

run:
	./usdt-grpc-service

test:
	go test -v ./...

lint:
	golangci-lint run