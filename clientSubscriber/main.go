package main

import (
	"context"
	"encoding/gob"
	"golangjob/pubsub"
	"log"
)

func main() {
	config := pubsub.HandleConfig()

	connection := pubsub.CreateConnection(config.PortSub)

	stream, err := connection.AcceptStream(context.Background())
	if err != nil {

		log.Println("Subscriber accept stream error", err)
	}

	for {
		decoder := gob.NewDecoder(stream)
		var receivedMessage pubsub.Message
		err := decoder.Decode(&receivedMessage)
		if err != nil {
			if err.Error() == "EOF" {
				return
			} else {
				log.Fatal("decode error?", err)

				return
			}
		}
		log.Println("["+receivedMessage.Sender+"]: "+receivedMessage.Content, "[", receivedMessage.Timestamp, "]")

	}
}
