package pubsub

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/quic-go/quic-go"
	"log"
	"os"
)

func Logger(appName string) *log.Logger {
	return log.New(
		os.Stderr,
		color.Ize(color.Bold+color.BlueBackground+color.White, "[ "+appName+" ]")+" ",
		log.Ltime,
	)
}
func GetAddressFromPort(port *int) string {
	return fmt.Sprintf("0.0.0.0:%d", *port)
}

func CreateConnection(portName *int) quic.Connection {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"jobExample"},
	}

	con, err := quic.DialAddr(
		context.Background(),
		GetAddressFromPort(portName),
		tlsConf,
		nil,
	)
	if err != nil {
		log.Fatal("Cant connect to app server. Error: ", err)
	}

	return con
}

func CreateConnectionStreamSync(portName *int) quic.Stream {
	con := CreateConnection(portName)

	stream, err := con.OpenStreamSync(context.Background())
	if err != nil {
		log.Fatal("Error creating stream", err)
	}

	return stream
}