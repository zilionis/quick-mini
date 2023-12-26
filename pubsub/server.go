package pubsub

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"github.com/quic-go/quic-go"
	"log"
	"math/big"
	"sync"
)

type Server struct {
	subscribers map[quic.Stream]struct{}
	mu          sync.RWMutex
	log         *log.Logger
	name        string
}

func CreateServer() *Server {
	return &Server{
		subscribers: make(map[quic.Stream]struct{}),
		name:        "Server",
		log:         Logger("Server"),
	}
}

func (server *Server) ListenAndServe(port *int, appName string, serveWithConnection func(conn quic.Connection)) {
	addr := GetAddressFromPort(port)
	listener, err := quic.ListenAddr(addr, GenerateTLSConfig(), nil)

	if err != nil {
		server.log.Fatal(err)
		return
	}
	defer func() {
		server.log.Printf("[ %v ] Listening on port\n", appName)
		_ = listener.Close()
	}()

	server.log.Println("[ "+appName+" ] Listening on port", addr)

	for {
		connection, err := listener.Accept(context.Background())
		if err != nil {
			server.log.Fatal(err)
		}

		go serveWithConnection(connection)
	}
}

func GenerateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{tlsCert},
		NextProtos:         []string{"jobExample"},
	}
}
