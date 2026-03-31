package tool

import (
	"context"
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
