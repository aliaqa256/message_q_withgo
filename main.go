package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"q/queue"
	"syscall"
	"time"

)

func main() {
	rand.Seed(time.Now().UnixNano())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	// ! simple queue
	q := queue.Queue{}

	go func() {
		for i := 0; i < 1000000; i++ {
			q.Enqueue(i)
			time.Sleep(time.Duration(rand.Intn(500)+100) * time.Millisecond)
		}
	}()

	go func() {
		for {
			out := q.Dequeue()
			if out == -1 {
				continue
			}
			fmt.Println(out,"+++++++++")
			time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
		}
	}()
	go func() {
		for {
			out := q.Dequeue()
			if out == -1 {
				continue
			}
			fmt.Println(out,"---------")
			time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
		}
	}()
	// server := s.NewServer("main")
	// q1 := mq.NewMessageQueue(1)
	// q2 := mq.NewMessageQueue(2)
	// server.AddQueue(q1)
	// server.AddQueue(q2)

	// consumer1 := c.NewConsumer(q1)
	// consumer2 := c.NewConsumer(q2)

	// go consumer1.Consume()
	// go consumer2.Consume()
	// go func() {
	// 	for {
	// 		// time.Sleep(1 * time.Second)
	// 		q1.AddMessage(msg.NewMessage( "1", 1))
	// 	}
	// 	}()
	// 	server.Run()

	<-sigs

}
