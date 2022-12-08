package massage

import (
	"github.com/satori/go.uuid"
	"sync"
)

type Message struct {
	sync.Mutex
	Id string
	Content string
	ExpirySeconds int
	LivedSeconds int
}


func NewMessage(content string, expirySeconds int) *Message {
	id := uuid.NewV4().String()
	return &Message{Id: id, Content: content, ExpirySeconds: expirySeconds, LivedSeconds: 0}
}
