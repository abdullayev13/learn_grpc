package main

import (
	"context"
	"fmt"
	"github.com/abdullayev13/learn_grpc/chat"
	"github.com/abdullayev13/learn_grpc/config"
	"google.golang.org/grpc"
	"log"
	"time"
)

var (
	ctx = context.Background()
	c   chat.ChatServiceClient
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(config.ServerPort, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connet port [%s] : %v", config.ServerPort, err)
	}

	defer conn.Close()

	c = chat.NewChatServiceClient(conn)

	fun()
}

func fun() {
	since := time.Now()
	for i := 0; i < 100; i++ {
		sincei := time.Now()
		msg := &chat.Message{Body: fmt.Sprintf("I'm Client%d", i+1)}

		res, err := c.SayHello(ctx, msg)

		if err != nil {
			log.Fatalf("server faled to SayHello : %v", err)
		}

		fmt.Printf("%s\n%s\n==================\n",
			res.Body, time.Now().Sub(sincei).String())
	}

	done := time.Now()

	fmt.Print("time gone : ")
	fmt.Println(done.Sub(since))

}
