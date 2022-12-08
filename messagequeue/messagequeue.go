package messagequeue

// message queue with go
import (
	msg "q/messagequeue/message"
	"sync"

	uuid "github.com/satori/go.uuid"
)


type MessageQueue struct {
	sync.Mutex 
	Id string
	Items []*msg.Message
	RetentionSeconds int
}



func NewMessageQueue(retentionSeconds int) *MessageQueue {
	id := uuid.NewV4().String()
	return &MessageQueue{Id:id,RetentionSeconds: retentionSeconds}
}


func (mq *MessageQueue) AddMessage(Message *msg.Message) {
	mq.Items = append(mq.Items , Message)
}



func RemoveMessageFromQueue(queue *MessageQueue, message *msg.Message) {
	for i, item := range queue.Items {
		if item.Id == message.Id {
			queue.Items = append(queue.Items[:i], queue.Items[i+1:]...)
		}
	}
}