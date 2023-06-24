package main

import (
	"context"
	"fmt"
	"github.com/abdullayev13/learn_grpc/chat"
	"google.golang.org/grpc"
	"log"
	"time"
)

var (
	serverPort = ":9000"
	ctx        = context.Background()
	c          chat.ChatServiceClient
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(serverPort, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connet port [%s] : %v", serverPort, err)
	}

	defer conn.Close()

	c = chat.NewChatServiceClient(conn)

	for i := 0; i < 1299; i++ {
		go fun()
	}
	fun()
	time.Sleep(23 * time.Second)
}

func fun() {
	since := time.Now()
	for i := 0; i < 100; i++ {
		msg := &chat.Message{Body: fmt.Sprintf("I'm Client%d", i+1)}

		res, err := c.SayHello(ctx, msg)

		if err != nil {
			log.Fatalf("server faled to SayHello : %v", err)
		}

		fmt.Println(res.Body)
	}

	done := time.Now()

	fmt.Print("time gone : ")
	fmt.Println(done.Sub(since))

}
