package main

import (
	"fileserver/internal/server"
	"flag"
	"fmt"
	"log"
)

const DEFAULT_PORT = 5050

func main() {
	port := flag.Int("port", DEFAULT_PORT, "Port from which listen for incoming connections")
	flag.Parse()

	server := server.NewTcpChatServer()
	defer server.Close()
	log.Fatal(server.ListenAndServe(fmt.Sprintf("localhost:%d", *port)))
}
