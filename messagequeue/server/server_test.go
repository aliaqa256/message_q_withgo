package server

import (
	mq "q/messagequeue"
	msg "q/messagequeue/message"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	s := NewServer("test")
	assert.Equal(t, 0, len(s.Queues))
	assert.Equal(t, "test", s.serverName)
}


func AddQueue(t *testing.T) {
	s := NewServer("test")
	msgq := mq.NewMessageQueue(10)
	s.AddQueue(msgq)
	assert.Equal(t, 1, len(s.Queues))
	assert.Equal(t, 10, s.Queues[0].RetentionSeconds)
	assert.Equal(t, 0, len(s.Queues[0].Items))
	assert.NotEqual(t, "", s.Queues[0].Id)

}


func TestAddMassage(t *testing.T) {
	s := NewServer("test")
	msgq := mq.NewMessageQueue(10)
	s.AddQueue(msgq)

	m := msg.NewMessage("hello", 10)
	msgq.AddMessage(m)
	assert.Equal(t, 1, len(s.Queues[0].Items))
	assert.Equal(t, "hello", s.Queues[0].Items[0].Content)
	assert.Equal(t, 10, s.Queues[0].Items[0].ExpirySeconds)
	assert.Equal(t, 0, s.Queues[0].Items[0].LivedSeconds)
	assert.NotEqual(t, "", s.Queues[0].Items[0].Id)
}

