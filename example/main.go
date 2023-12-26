package main

import (
	"golangjob/pubsub"
	"time"
)

func main() {
	config := pubsub.HandleConfig()
	server := pubsub.CreateServer()

	go server.ListenAndServe(config.PortSub, "Subscriber", server.HandleSubscriber)
	go server.ListenAndServe(config.PortPub, "Publisher", server.HandlePublisher)

	go pubsub.PublisherApp("Pub__1", 3, *config.PortPub)
	go pubsub.PublisherApp("Pub__2", 7, *config.PortPub)

	time.Sleep(8 * time.Second)

	go pubsub.SubscriberApp("Sub__1", *config.PortSub)
	go pubsub.SubscriberApp("Sub__2", *config.PortSub)

	select {}
}
