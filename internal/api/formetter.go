package api

import (
	"cmp"
	"fmt"
)

func FormatAlert(alert AlertFeature) string {
	props := alert.Properties
	event := cmp.Or(props.Event, "不明")
	areaDesc := cmp.Or(props.AreaDesc, "不明")
	severity := cmp.Or(props.Severity, "不明")
	description := cmp.Or(props.Desctiption, "説明はありません")
	instruction := cmp.Or(props.Instruction, "特別な指示はありません")

	return fmt.Sprintf(`
警報: %s
対象地域: %s
重要度: %s
説明: %s
指示: %s
`, event, areaDesc, severity, description, instruction)
}

func FormatPeriod(period ForecastPeriod) string {
	precipitationProbability := "不明"
	if period.PrecipitationProbabilityMax >= 0 {
		precipitationProbability = fmt.Sprintf("%d%%", period.PrecipitationProbabilityMax)
	}

	return fmt.Sprintf(`
%s:
天気: %s
気温: %.1f-%.1f°C
降水量: %.1f mm
降水確率: %s
最大風速: %.1f km/h
`, period.Date, period.WeatherDescription, period.TemperatureMin, period.TemperatureMax, period.PrecipitationSum, precipitationProbability, period.WindSpeedMax)
}

func WeatherCodeDescription(code int) string {
	switch code {
	case 0:
		return "快晴"
	case 1:
		return "おおむね晴れ"
	case 2:
		return "晴れ時々くもり"
	case 3:
		return "くもり"
	case 45, 48:
		return "霧"
	case 51, 53, 55:
		return "霧雨"
	case 56, 57:
		return "着氷性の霧雨"
	case 61, 63, 65:
		return "雨"
	case 66, 67:
		return "着氷性の雨"
	case 71, 73, 75, 77:
		return "雪"
	case 80, 81, 82:
		return "にわか雨"
	case 85, 86:
		return "にわか雪"
	case 95:
		return "雷雨"
	case 96, 99:
		return "ひょうを伴う雷雨"
	default:
		return "不明"
	}
}
