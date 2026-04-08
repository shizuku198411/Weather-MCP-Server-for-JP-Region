package tool

import (
	"context"
	"fmt"
	"mcp_server/internal/api"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func NewToolHandler() *ToolHandler {
	return &ToolHandler{
		APIHandler: *api.NewNWSAPI(),
	}
}

type ToolHandler struct {
	APIHandler api.NWSAPI
}

func (h *ToolHandler) GetCoordinates(ctx context.Context, req *mcp.CallToolRequest, input api.CoordinatesInput) (*mcp.CallToolResult, any, error) {
	query := strings.TrimSpace(input.Query)
	if query == "" {
		return h.MCPTextResult([]string{"Location query is required"}), nil, nil
	}

	coordinates, err := h.APIHandler.GetCoordinates(ctx, query)
	if err != nil {
		return h.MCPTextResult([]string{"Unable to resolve coordinates for this location"}), nil, nil
	}

	var results []string
	for _, result := range coordinates.Results {
		if result.CountryCode != "JP" {
			continue
		}

		results = append(results, formatCoordinateResult(result))
	}

	if len(results) == 0 {
		return h.MCPTextResult([]string{"No coordinates found for this location in Japan"}), nil, nil
	}

	return h.MCPTextResult(results), nil, nil
}

func (h *ToolHandler) GetForecast(ctx context.Context, req *mcp.CallToolRequest, input api.ForecastInput) (*mcp.CallToolResult, any, error) {
	forecastData, err := h.APIHandler.GetForecast(ctx, input.Latitude, input.Longitude)
	if err != nil {
		return h.MCPTextResult([]string{"Unable to fetch forecast data for this location"}), nil, nil
	}

	daily := forecastData.Daily
	periodCount := min(
		len(daily.Time),
		len(daily.WeatherCode),
		len(daily.Temperature2MMax),
		len(daily.Temperature2MMin),
		len(daily.PrecipitationProbabilityMax),
		len(daily.PrecipitationSum),
		len(daily.WindSpeed10MMax),
	)
	if periodCount == 0 {
		return h.MCPTextResult([]string{"No forecast periods available"}), nil, nil
	}

	var forecasts []string
	for i := range periodCount {
		forecasts = append(forecasts, api.FormatPeriod(api.ForecastPeriod{
			Date:                        daily.Time[i],
			WeatherDescription:          api.WeatherCodeDescription(daily.WeatherCode[i]),
			TemperatureMax:              daily.Temperature2MMax[i],
			TemperatureMin:              daily.Temperature2MMin[i],
			PrecipitationProbabilityMax: daily.PrecipitationProbabilityMax[i],
			PrecipitationSum:            daily.PrecipitationSum[i],
			WindSpeedMax:                daily.WindSpeed10MMax[i],
		}))
	}

	result := strings.Join(forecasts, "\n--\n")

	return h.MCPTextResult([]string{result}), nil, nil
}

func (h *ToolHandler) MCPTextResult(texts []string) *mcp.CallToolResult {
	var textContent []mcp.Content
	for _, text := range texts {
		textContent = append(textContent,
			&mcp.TextContent{
				Text: text,
			},
		)
	}

	return &mcp.CallToolResult{Content: textContent}
}

func formatCoordinateResult(result api.GeocodingResult) string {
	var lines []string

	lines = append(lines, fmt.Sprintf("地点: %s", result.Name))
	lines = append(lines, fmt.Sprintf("緯度: %.6f", result.Latitude))
	lines = append(lines, fmt.Sprintf("経度: %.6f", result.Longitude))

	if result.Country != "" {
		lines = append(lines, fmt.Sprintf("国: %s", result.Country))
	}
	if result.Admin1 != "" {
		lines = append(lines, fmt.Sprintf("都道府県: %s", result.Admin1))
	}
	if result.Admin2 != "" {
		lines = append(lines, fmt.Sprintf("地域2: %s", result.Admin2))
	}
	if result.Admin3 != "" {
		lines = append(lines, fmt.Sprintf("地域3: %s", result.Admin3))
	}

	return strings.Join(lines, "\n")
}
