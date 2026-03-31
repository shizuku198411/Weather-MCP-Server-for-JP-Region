package api

type ForecastResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Daily     struct {
		Time                        []string  `json:"time"`
		WeatherCode                 []int     `json:"weather_code"`
		Temperature2MMax            []float64 `json:"temperature_2m_max"`
		Temperature2MMin            []float64 `json:"temperature_2m_min"`
		PrecipitationProbabilityMax []int     `json:"precipitation_probability_max"`
		PrecipitationSum            []float64 `json:"precipitation_sum"`
		WindSpeed10MMax             []float64 `json:"wind_speed_10m_max"`
	} `json:"daily"`
}

type ForecastPeriod struct {
	Date                        string
	WeatherDescription          string
	TemperatureMax              float64
	TemperatureMin              float64
	PrecipitationProbabilityMax int
	PrecipitationSum            float64
	WindSpeedMax                float64
}

type AlertsResponse struct {
	Features []AlertFeature `json:"features"`
}

type AlertFeature struct {
	Properties AlertProperties `json:"properties"`
}

type AlertProperties struct {
	Event       string `json:"event"`
	AreaDesc    string `json:"areaDesc"`
	Severity    string `json:"severity"`
	Desctiption string `json:"description"`
	Instruction string `json:"instruction"`
}

type ForecastInput struct {
	Latitude  float64 `json:"latitude" jsonschema:"Latitude of the location"`
	Longitude float64 `json:"longitude" jsonschema:"Longitude of the location"`
}

type AlertsInput struct {
	State string `json:"state" jsonschema:"state code"`
}
