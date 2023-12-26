package pubsub

import (
	"context"
	"encoding/gob"
	"github.com/quic-go/quic-go"
)

func (server *Server) HandlePublisher(connection quic.Connection) {
	server.log.Println("New publisher connected", connection.RemoteAddr().String())

	stream, err := connection.AcceptStream(context.Background())
	if err != nil {
		server.log.Println("Stream error", err)
		return
	}

	var initialSubscribers int
	initialSubscribers = len(server.subscribers)
	if initialSubscribers == 0 {
		initialSubscribers = -1
	}

	for {
		decoder := gob.NewDecoder(stream)
		var receivedMessage Message
		err := decoder.Decode(&receivedMessage)
		if err != nil {
			if err.Error() == "EOF" {
				return
			} else {
				server.log.Fatal("Decode error?", err)

				return
			}
		}
		server.sendToSubscribers(receivedMessage)
		initialSubscribers = server.NotifyPublisher(stream, initialSubscribers)
	}
}

func (server *Server) NotifyPublisher(stream quic.Stream, subscribersBefore int) int {
	subscribersNow := len(server.subscribers)

	var msg string
	if subscribersBefore != subscribersNow {
		if subscribersNow == 0 {
			msg = "No subscribers connected"
		} else {
			msg = "New subscriber connected!"
		}

		m := NewMessage(msg, server.name)
		err := m.SendToStream(stream)
		if err != nil {
			server.log.Println("Error sending notification", err)
		}

	}

	return subscribersNow
}

func (server *Server) sendToSubscribers(message Message) {
	server.mu.RLock()

	server.log.Println("Sending msg to subscribers from", message.Sender, "total: ", len(server.subscribers))

	for subscriberStream := range server.subscribers {
		err := message.SendToStream(subscriberStream)
		if err != nil {
			if err.Error() == "timeout: no recent network activity" {
				defer server.RemoveSubscriber(subscriberStream)
			} else {
				server.log.Println("Sending error: ", err)
			}
		}
	}

	defer server.mu.RUnlock()
}
