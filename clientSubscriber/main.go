package main

import (
	"golangjob/pubsub"
)

func main() {
	config := pubsub.HandleConfig()

	pubsub.SubscriberApp("Publisher_1", *config.PortSub)
}
