package server

import (
	"context"
	"log"
	"mcp_server/internal/config"
	"mcp_server/internal/tool"
	"net/http"
	"strings"

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

func (s *MCPServer) Run() error {
	server := s.newServer()

	switch strings.ToLower(s.Config.Transport) {
	case "http", "streamable-http", "streamable_http":
		return s.runHTTP(server)
	case "stdio", "":
		return server.Run(context.Background(), &mcp.StdioTransport{})
	default:
		log.Printf("unknown MCP_TRANSPORT %q, fallback to stdio", s.Config.Transport)
		return server.Run(context.Background(), &mcp.StdioTransport{})
	}
}

func (s *MCPServer) newServer() *mcp.Server {
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
			Name:        "get_coordinates",
			Description: "Resolve Japanese location names into latitude and longitude using Open-Meteo geocoding",
		},
		s.Tools.GetCoordinates,
	)

	mcp.AddTool(server,
		&mcp.Tool{
			Name:        "get_forecast",
			Description: "Get weather forecast for a location",
		},
		s.Tools.GetForecast,
	)

	return server
}

func (s *MCPServer) runHTTP(server *mcp.Server) error {
	mux := http.NewServeMux()

	mux.Handle(s.Config.HTTPPath, mcp.NewStreamableHTTPHandler(
		func(_ *http.Request) *mcp.Server {
			return server
		},
		&mcp.StreamableHTTPOptions{
			Stateless:    true,
			JSONResponse: true,
		},
	))

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	addr := s.Config.HTTPHost + ":" + s.Config.HTTPPort
	log.Printf("starting MCP server over HTTP on %s%s", addr, s.Config.HTTPPath)

	return http.ListenAndServe(addr, mux)
}
