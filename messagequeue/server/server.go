package server

import (
	"fmt"
	"os"
	"os/signal"
	mq "q/messagequeue"
	"sync/atomic"
	"syscall"
	"time"
)

type Server struct {
	serverName string
	Queues     []*mq.MessageQueue
}

func NewServer(serverName string) *Server {
	return &Server{serverName: serverName}
}

func (server *Server) AddQueue(queue *mq.MessageQueue) {
	server.Queues = append(server.Queues, queue)
}

func (server *Server) Run() {
	ticker := time.NewTicker(1 * time.Second)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		for {
			select {
			case <-ticker.C:
				for _, queue := range server.Queues {
					queue.Lock()

					for _, message := range queue.Items {
						

						if message.LivedSeconds >= message.ExpirySeconds {
							mq.RemoveMessageFromQueue(queue, message)
							atomic.AddInt32(&message.LivedSeconds, 1)

							continue
						}
						if message.LivedSeconds >= queue.RetentionSeconds {
							mq.RemoveMessageFromQueue(queue, message)

						}
						atomic.AddInt32(&message.LivedSeconds, 1)

					}
					queue.Unlock()

				}

			case <-sigs:
				ticker.Stop()
				fmt.Println("stoping the server ...")
				return

			}
		}
	}()

}

// fmt.Println("queue", i, "has", len(queue.Items), "messages")
// 					for j, message := range queue.Items {
// 						if message.LivedSeconds >= message.ExpirySeconds {
// 							fmt.Println("removing message from queue")
// 							mq.RemoveMessageFromQueue(server.Queues[i], &queue.Items[j])
// 							fmt.Println("message removed")
// 						}
// 						if message.LivedSeconds >= queue.RetentionSeconds {
// 							fmt.Println("removing message from queue")
// 							mq.RemoveMessageFromQueue(server.Queues[i], &queue.Items[j])
// 						}
// 						message.LivedSeconds += 1
// 					}
