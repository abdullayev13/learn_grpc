package main

import (
	"github.com/abdullayev13/learn_grpc/chat"
	"github.com/abdullayev13/learn_grpc/config"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", config.ServerPort)
	if err != nil {
		log.Fatalf("failed to liten port [%s] : %v", config.ServerPort, err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc over port [%s] : %v", config.ServerPort, err)
	}

}
