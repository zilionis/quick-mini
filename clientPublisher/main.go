package main

import (
	"encoding/gob"
	"golangjob/pubsub"
	"strconv"
	"time"
)

func main() {
	config := pubsub.HandleConfig()

	appName := "Publisher_1"
	logger := pubsub.Logger(appName)
	var tickerSeconds time.Duration = 10

	stream := pubsub.CreateConnectionStreamSync(config.PortPub)
	defer stream.Close()

	go func() {
		ticker := time.NewTicker(tickerSeconds * time.Second)
		for range ticker.C {
			m := pubsub.NewMessage("Hello from app", appName)
			logger.Println("<-- Sending to server message:" + m.Content + " at " + strconv.FormatInt(m.Timestamp, 10))
			err := m.SendToStream(stream)
			if err != nil {
				logger.Fatal("Error sending to stream", err)
			}
		}
	}()

	go func() {
		for {
			decoder := gob.NewDecoder(stream)
			var receivedMessage pubsub.Message
			err := decoder.Decode(&receivedMessage)
			if err != nil {
				if err.Error() == "EOF" {
					logger.Fatal("Stream session ended?")
				} else {
					logger.Println("decode error?", err)

					continue
				}
			}
			logger.Println("--> ["+receivedMessage.Sender+"]: "+receivedMessage.Content, "[", receivedMessage.Timestamp, "]")
		}
	}()

	select {}
}
