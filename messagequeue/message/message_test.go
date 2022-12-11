package massage

import (
	"testing"
	"github.com/stretchr/testify/assert"

)

func TestMessage(t *testing.T) {
	m := NewMessage("hello", 10)
	assert.Equal(t, "hello", m.Content)
	assert.Equal(t, int32(10), m.ExpirySeconds)
	assert.Equal(t, int32(0), m.LivedSeconds)
	assert.NotEqual(t, "", m.Id)
}
	