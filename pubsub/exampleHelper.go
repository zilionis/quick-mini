package pubsub

import (
	"context"
	"encoding/gob"
	"github.com/TwiN/go-color"
	"github.com/quic-go/quic-go"
	"log"
	"strconv"
	"time"
)

func PublisherApp(appName string, tickerDuration time.Duration, publisherPort int) {
	logger := Logger(appName, color.BlackBackground)

	stream := CreateConnectionStreamSync(&publisherPort)
	defer stream.Close()

	go func() {
		ticker := time.NewTicker(tickerDuration * time.Second)
		for range ticker.C {
			m := NewMessage("Hello from app", appName)
			logger.Println("<-- " + m.Content + " at " + strconv.FormatInt(m.Timestamp, 10))
			err := m.SendToStream(stream)
			if err != nil {
				logger.Fatal("Error sending to stream", err)
			}
		}
	}()

	for {
		messageReceiver(&stream, logger)
	}
}

func SubscriberApp(appName string, subscriberPort int) {
	logger := Logger(appName, color.GreenBackground)
	connection := CreateConnection(&subscriberPort)

	stream, err := connection.AcceptStream(context.Background())
	if err != nil {

		log.Println("Subscriber accept stream error", err)
	}

	for {
		messageReceiver(&stream, logger)
	}

	select {}
}

func messageReceiver(stream *quic.Stream, logger *log.Logger) {
	decoder := gob.NewDecoder(*stream)
	var receivedMessage Message
	err := decoder.Decode(&receivedMessage)
	if err != nil {
		if err.Error() == "EOF" {
			logger.Fatal("Stream session ended?")
		} else {
			logger.Println("decode error?", err)

			return
		}
	}
	logger.Println("--> ["+receivedMessage.Sender+"]: "+receivedMessage.Content, "[", receivedMessage.Timestamp, "]")
}
