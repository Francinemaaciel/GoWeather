package models

type WeatherData struct{

	Nome string `json:"name"`
	Main struct{
		Temp float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
	} `json:"main"`
	Weather []struct {
		Descricao string `json:"description"`
	} `json:"weather"`
}