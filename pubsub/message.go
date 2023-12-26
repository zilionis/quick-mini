package pubsub

import (
	"encoding/gob"
	"github.com/quic-go/quic-go"
	"time"
)

type Message struct {
	Sender    string
	Content   string
	Timestamp int64
}

func NewMessage(content string, Sender string) *Message {
	return &Message{Content: content, Sender: Sender, Timestamp: time.Now().Unix()}
}

func (m Message) getContent() string {
	return m.Content
}

func (m Message) SendToStream(stream quic.Stream) error {
	encoder := gob.NewEncoder(stream)
	if err := encoder.Encode(m); err != nil {
		return err
	}

	return nil
}
