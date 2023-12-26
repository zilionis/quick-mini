package pubsub

import (
	"flag"
	"fmt"
	"os"
)

type AppConfig struct {
	PortSub *int
	PortPub *int
	Verbose *bool
	Name    *string
}

func HandleConfig() AppConfig {

	config := AppConfig{
		PortSub: flag.Int("port", 6001, "PortSub number for clientSubscriber"),
		PortPub: flag.Int("portPub", 6002, "PortSub number for publisher"),
		Verbose: flag.Bool("v", false, "verbose"),
		Name:    flag.String("name", "Server", "Name"),
	}

	flag.Usage = Usage
	flag.Parse()

	args := flag.Args()

	if len(args) > 0 && args[0] == "help" {
		flag.Usage()
	}

	return config
}

var Usage = func() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}
