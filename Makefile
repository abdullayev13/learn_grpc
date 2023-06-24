
protoc:
	protoc --go_out=plugins=grpc:chat chat.proto

server:
	go run server/server.go

client:
	go run client/client.go



