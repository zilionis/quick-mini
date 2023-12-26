package main

import (
	"golangjob/pubsub"
)

func main() {
	config := pubsub.HandleConfig()

	pubsub.PublisherApp("Publisher_1", 5, *config.PortPub)
}
