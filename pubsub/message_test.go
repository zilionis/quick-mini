package pubsub

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMessage(t *testing.T) {
	m := NewMessage("a-content", "b-sender")
	assert.Equal(t, "a-content", m.Content)
	assert.Equal(t, "b-sender", m.Sender)
}
