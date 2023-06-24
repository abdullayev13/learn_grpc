package chat

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

type Server struct {
	count int64
}

func (s *Server) SayHello(c context.Context, msg *Message) (*Message, error) {
	count := s.incrementCount()
	fmt.Println("Message on server : ")
	fmt.Println(msg.Body)
	fmt.Println("=============================")
	fmt.Printf("Total count :				 || %d ||\n", count)
	fmt.Println("=============================")

	time.Sleep(time.Millisecond)

	return &Message{Body: fmt.Sprintf("Hello From Server. It's %d-msg", count)}, nil
}

func (s *Server) incrementCount() int64 {
	return atomic.AddInt64(&s.count, 1)
}
