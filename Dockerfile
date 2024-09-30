FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o usdt-grpc-service ./cmd/app

CMD ["./usdt-grpc-service"]