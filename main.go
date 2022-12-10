package main

import (
	// "q/queue"
	// "fmt"
	"os"
	"os/signal"
	mq "q/messagequeue"
	c "q/messagequeue/consumer"
	msg "q/messagequeue/message"
	s "q/messagequeue/server"
	"syscall"
	// "time"
	// "time"
)



func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	// ! simple queue
	// q := queue.Queue{}
	// q.Enqueue(1)
	// q.Enqueue(2)
	// q.Enqueue(3)
	// q.Print()
	// fmt.Println(q.Dequeue())
	// q.Print()
	// fmt.Println(q.IsEmpty())
	// fmt.Println(q.Peek())
	// fmt.Println(q.Size())
// message queue

	server := s.NewServer("main")
	q1 := mq.NewMessageQueue(1)
	q2 := mq.NewMessageQueue(2)
	server.AddQueue(q1)
	server.AddQueue(q2)
	
	consumer1 := c.NewConsumer(q1)
	consumer2 := c.NewConsumer(q2)

	go consumer1.Consume()		
	go consumer2.Consume()
	go func() {
		for {
			// time.Sleep(1 * time.Second)
			q1.AddMessage(msg.NewMessage( "1", 1))
		}
		}()
		server.Run()
		
	<-sigs
		
		
		
		
		
		







}

