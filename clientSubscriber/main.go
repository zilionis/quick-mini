package main

import (
	"golangjob/pubsub"
)

func main() {
	config := pubsub.HandleConfig()

	pubsub.SubscriberApp("Subscriber_1", *config.PortSub)
}
