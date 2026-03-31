package main

import "mcp_server/internal/server"

func main() {
	// run server
	server := server.NewMCPServer()
	server.Run()
}
