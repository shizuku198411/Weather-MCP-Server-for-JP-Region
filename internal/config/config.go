package config

import "os"

type Config struct {
	ServerName    string
	ServerVersion string
	APIURL        string
	GeocodingURL  string
	UserAgent     string
	Transport     string
	HTTPHost      string
	HTTPPort      string
	HTTPPath      string
}

func LoadConfig() Config {
	serverName := os.Getenv("MCP_SERVER_NAME")
	if serverName == "" {
		serverName = "weather"
	}

	serverVersion := os.Getenv("MCP_SERVER_VERSION")
	if serverVersion == "" {
		serverVersion = "0.1.0"
	}

	baseURL := os.Getenv("MCP_API_BASE_URL")
	if baseURL == "" {
		baseURL = "https://api.open-meteo.com/v1"
	}

	geocodingURL := os.Getenv("MCP_GEOCODING_API_BASE_URL")
	if geocodingURL == "" {
		geocodingURL = "https://geocoding-api.open-meteo.com/v1"
	}

	userAgent := os.Getenv("MCP_USER_AGENT")
	if userAgent == "" {
		userAgent = "weather-app/0.1.0"
	}

	transport := os.Getenv("MCP_TRANSPORT")
	if transport == "" {
		transport = "stdio"
	}

	httpHost := os.Getenv("MCP_HTTP_HOST")
	if httpHost == "" {
		httpHost = "0.0.0.0"
	}

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = os.Getenv("MCP_HTTP_PORT")
	}
	if httpPort == "" {
		httpPort = "8080"
	}

	httpPath := os.Getenv("MCP_HTTP_PATH")
	if httpPath == "" {
		httpPath = "/mcp"
	}

	return Config{
		ServerName:    serverName,
		ServerVersion: serverVersion,
		APIURL:        baseURL,
		GeocodingURL:  geocodingURL,
		UserAgent:     userAgent,
		Transport:     transport,
		HTTPHost:      httpHost,
		HTTPPort:      httpPort,
		HTTPPath:      httpPath,
	}
}
