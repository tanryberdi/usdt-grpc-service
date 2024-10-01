# usdt-grpc-service

## Обзор
`usdt-grpc-service` — это основанная на gRPC служба для извлечения и хранения курсов обмена USDT. Она подключается к базе данных PostgreSQL для сохранения курсов и предоставляет конечные точки для извлечения курсов и проверки работоспособности службы.

## Функции
- Получайте (Fetch) курсы обмена USDT из внешнего API.
- Сохраняйте курсы валют в базе данных PostgreSQL.
- Извлечение сохраненных курсов обмена через gRPC.
- Конечная точка проверки работоспособности для мониторинга состояния службы.

## Предпосылки
- Go 1.10 or или более поздние версии
- PostgreSQL
- `protoc` (Protocol Buffers compiler)
- `protoc-gen-go` и `protoc-gen-go-grpc` плагины

## Установка

1. **Клонировать репозиторий:**
    ```sh
    git clone https://github.com/tanryberdi/usdt-grpc-service.git
    cd usdt-grpc-service
    ```

2. **Установить зависимости:**
    ```sh
    go mod tidy
    ```

3. **Установить компилятор Protocol Buffers и плагины:**
    ```sh
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    ```
   
## Использование

1. **Генерация кода буферов протокола:**
    ```sh
    protoc --go_out=. --go-grpc_out=. proto/*.proto
    ```

2. **Настройте базу данных PostgreSQL:**
    ```sh
    export DB_HOST=localhost
    export DB_USER=yourusername
    export DB_PASSWORD=yourpassword
    export DB_NAME=testdb
    ```
3. **Создайте (build) службу**
    ```shell
    make build
    ```

4. **Запустить службу:**
    ```sh
    make run
    ```
5. **Получайте и сохраняйте курсы обмена USDT:**
    ```sh
    grpcurl -plaintext localhost:50051 rates.RateService/GetRates
    ```


## Тестирование (Testing)

1. **Run unit tests:**
    ```sh
    make test 
    ```

## Линтер (Linter)

1. **Check for linter:**
    ```sh
    make lint 
    ```

## Структура проекта

- `cmd/app/main.go`: Entry point of the application.
- `internal/db`: Database connection and operations.
- `internal/handler`: gRPC service handlers.
- `proto`: Protocol Buffers definitions.
- `internal/logger.go`: Logger initialization (zap).