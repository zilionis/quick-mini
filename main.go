package main

import (
	"golangjob/pubsub"
)

func main() {
	config := pubsub.HandleConfig()
	server := pubsub.CreateServer()

	go server.ListenAndServe(config.PortSub, "Subscriber", server.HandleSubscriber)
	go server.ListenAndServe(config.PortPub, "Publisher", server.HandlePublisher)

	select {}
}
