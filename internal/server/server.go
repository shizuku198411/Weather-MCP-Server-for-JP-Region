package server

import (
	"context"
	"log"
	"mcp_server/internal/config"
	"mcp_server/internal/tool"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func NewMCPServer() *MCPServer {
	return &MCPServer{
		Config: config.LoadConfig(),
		Tools:  *tool.NewToolHandler(),
	}
}

type MCPServer struct {
	Config config.Config
	Tools  tool.ToolHandler
}

func (s *MCPServer) Run() {
	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    s.Config.ServerName,
			Version: s.Config.ServerVersion,
		},
		nil,
	)

	// add tools
	mcp.AddTool(server,
		&mcp.Tool{
			Name:        "get_forecast",
			Description: "Get weather forecast for a location",
		},
		s.Tools.GetForecast,
	)

	// run server
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
