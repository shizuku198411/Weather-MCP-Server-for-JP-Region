package config

import "os"

type Config struct {
	ServerName    string
	ServerVersion string
	APIURL        string
	UserAgent     string
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

	userAgent := os.Getenv("MCP_USER_AGENT")
	if userAgent == "" {
		userAgent = "weather-app/0.1.0"
	}

	return Config{
		ServerName:    serverName,
		ServerVersion: serverVersion,
		APIURL:        baseURL,
		UserAgent:     userAgent,
	}
}
