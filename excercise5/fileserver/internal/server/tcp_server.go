package server

import (
	"fileserver/internal/protocol"
	"io"
	"log"
	"net"
)

type client struct {
	conn   net.Conn
	writer *protocol.MessageWriter
}

type TcpChatServer struct {
	listener net.Listener
}

func NewTcpChatServer() *TcpChatServer {
	return &TcpChatServer{}
}

func (s *TcpChatServer) ListenAndServe(address string) error {
	l, err := net.Listen("tcp", address)

	if err == nil {
		s.listener = l
	}

	s.start()
	return err
}

func (s *TcpChatServer) Close() {
	s.listener.Close()
}

func (s *TcpChatServer) start() {
	for {
		conn, err := s.listener.Accept()

		if err != nil {
			log.Println(err)
		} else {
			client := s.accept(conn)
			go s.serve(client)
		}
	}
}

func (s *TcpChatServer) accept(conn net.Conn) *client {
	log.Printf("Accepting connection from %v", conn.RemoteAddr())
	client := client{conn: conn,
		writer: protocol.NewMessageWriter(conn),
	}

	return &client
}

func (s *TcpChatServer) serve(client *client) {
	msgReader := protocol.NewMessageReader(client.conn)

	for {
		command, err := msgReader.Read(client.writer)

		if err != nil && err != io.EOF {
			log.Printf("Read error: %v\n", err)
		}

		if err == io.EOF {
			break
		}

		if command != nil {
			log.Println("Executing")
			command.ExecuteAndSend()
		}
	}
}
