PROTOC=protoc

proto-gen:
	$(PROTOC) --go_out=. --go_opt=paths=import \
    --go-grpc_out=. --go-grpc_opt=paths=import \
    proto/mailer.proto

dep: # Download required dependencies
	go mod tidy
	go mod download
	go mod vendor

build: dep
	CGO_ENABLED=1 go build -mod=vendor -o ./bin/${BIN_NAME} -a ./cmd/meeting-app

run: dep
	go run ./cmd/app
