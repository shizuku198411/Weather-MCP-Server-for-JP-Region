package main

import (
	"log"
	"mcp_server/internal/server"
)

func main() {
	server := server.NewMCPServer()
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
