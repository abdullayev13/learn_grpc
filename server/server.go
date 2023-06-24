package main

import (
	"github.com/abdullayev13/learn_grpc/chat"
	"log"
	"net"

	"google.golang.org/grpc"
)

var port = ":9000"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to liten port [%s] : %v", port, err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc over port [%s] : %v", port, err)
	}

}
