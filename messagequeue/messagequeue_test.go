package messagequeue

import (
	"testing"

	"github.com/stretchr/testify/assert"
	msg "q/messagequeue/message"
)

func TestMessageQueue(t *testing.T) {
	mq := NewMessageQueue(10)
	assert.Equal(t, 0, len(mq.Items))
	assert.Equal(t, 10, mq.RetentionSeconds)
	assert.NotEqual(t, "", mq.Id)

}

func TestAddMessage(t *testing.T) {
	mq := NewMessageQueue(10)
	assert.Equal(t, 0, len(mq.Items))
	mq.AddMessage(msg.NewMessage("hello", 10))
	assert.Equal(t, 1, len(mq.Items))
	assert.Equal(t, "hello", mq.Items[0].Content)
	assert.Equal(t, 10, mq.Items[0].ExpirySeconds)
	assert.Equal(t, 0, mq.Items[0].LivedSeconds)
	assert.NotEqual(t, "", mq.Items[0].Id)
}


 
func TestRemoveMessageFromQueue(t *testing.T) {
	mq := NewMessageQueue(10)
	assert.Equal(t, 0, len(mq.Items))
	mq.AddMessage(msg.NewMessage("hello", 10))
	assert.Equal(t, 1, len(mq.Items))
	assert.Equal(t, "hello", mq.Items[0].Content)
	assert.Equal(t, 10, mq.Items[0].ExpirySeconds)
	assert.Equal(t, 0, mq.Items[0].LivedSeconds)
	assert.NotEqual(t, "", mq.Items[0].Id)
	RemoveMessageFromQueue(mq, mq.Items[0])
	assert.Equal(t, 0, len(mq.Items))
}


