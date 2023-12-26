package pubsub

import (
	"context"
	"github.com/quic-go/quic-go"
)

func (server *Server) HandleSubscriber(connection quic.Connection) {

	server.log.Println("Subscriber connected: ", connection.RemoteAddr().String(), ". Total now:", len(server.subscribers)+1)

	stream, err := connection.OpenStreamSync(context.Background())
	if err != nil {
		server.log.Println("Subscriber AcceptStream err: ", err)
	}

	defer func() {
		server.RemoveSubscriber(stream)
		stream.Close()
	}()

	server.AddSubscriber(stream)

	select {}
}

func (server *Server) AddSubscriber(stream quic.Stream) {
	server.mu.Lock()
	defer server.mu.Unlock()

	server.subscribers[stream] = struct{}{}
}

func (server *Server) RemoveSubscriber(stream quic.Stream) {
	server.mu.Lock()
	defer server.mu.Unlock()

	delete(server.subscribers, stream)
}
