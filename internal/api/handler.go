package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mcp_server/internal/config"
	"net/http"
	"net/url"
	"strings"
)

func NewNWSAPI() *NWSAPI {
	return &NWSAPI{
		Config: config.LoadConfig(),
	}
}

type NWSAPI struct {
	Config config.Config
}

func (api *NWSAPI) doRequest(ctx context.Context, requestURL string, out any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", api.Config.UserAgent)
	req.Header.Set("Accept", "application/geo+json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request to %s: %w", requestURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(body))
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}

func (api *NWSAPI) GetForecast(ctx context.Context, latitude, longitude float64) (*ForecastResponse, error) {
	requestURL := fmt.Sprintf(
		"%s/forecast?latitude=%f&longitude=%f&daily=weather_code,temperature_2m_max,temperature_2m_min,precipitation_probability_max,precipitation_sum,wind_speed_10m_max&forecast_days=5&timezone=%s",
		strings.TrimRight(api.Config.APIURL, "/"),
		latitude,
		longitude,
		url.QueryEscape("Asia/Tokyo"),
	)

	var result ForecastResponse
	if err := api.doRequest(ctx, requestURL, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (api *NWSAPI) GetCoordinates(ctx context.Context, query string) (*GeocodingResponse, error) {
	requestURL := fmt.Sprintf(
		"%s/search?name=%s&count=10&language=ja&format=json",
		strings.TrimRight(api.Config.GeocodingURL, "/"),
		url.QueryEscape(strings.TrimSpace(query)),
	)

	var result GeocodingResponse
	if err := api.doRequest(ctx, requestURL, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (api *NWSAPI) GetAlerts(ctx context.Context, state string) (*AlertsResponse, error) {
	requestURL := fmt.Sprintf("%s/alerts/active/area/%s", strings.TrimRight(api.Config.APIURL, "/"), url.PathEscape(state))

	var result AlertsResponse
	if err := api.doRequest(ctx, requestURL, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
