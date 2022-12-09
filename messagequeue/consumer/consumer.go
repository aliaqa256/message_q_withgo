package consumer

import (
	"fmt"
	mq "q/messagequeue"

	uuid "github.com/satori/go.uuid"
)



type Consumer struct {
	ConsumerId   string
	MessageQueue *mq.MessageQueue
	ReadMessagesId []string
}

func NewConsumer(mq *mq.MessageQueue) *Consumer {
	return &Consumer{ConsumerId: uuid.NewV4().String(), MessageQueue: mq}
}

func (consumer *Consumer) Consume()  {
	fmt.Println("consumer is consuming ...")
	for {
			consumer.MessageQueue.Lock()
		for _, message := range consumer.MessageQueue.Items {
			if !contains(consumer.ReadMessagesId, message.Id) {
				fmt.Println(message)
				consumer.ReadMessagesId = append(consumer.ReadMessagesId, message.Id)				
			}
		}
		consumer.MessageQueue.Unlock()
		

	}	
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
