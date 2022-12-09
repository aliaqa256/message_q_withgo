package massage

import (
	"github.com/satori/go.uuid"
	"sync"
)

type Message struct {
	sync.RWMutex
	Id string
	Content string
	ExpirySeconds int32
	LivedSeconds int32
}


func NewMessage(content string, expirySeconds int32) *Message {
	id := uuid.NewV4().String()
	return &Message{Id: id, Content: content, ExpirySeconds: expirySeconds, LivedSeconds: 0}
}
