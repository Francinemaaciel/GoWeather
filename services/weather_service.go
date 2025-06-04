package services

import (

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/francinemaaciel/estudos-go/models"
)

	//Faz a requisição e retorna os dados estruturados
func GetWeatherAPI (city string, apiKey string)(*models.WeatherData, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric&lang=pt_br", city, apiKey)

	resp, erro := http.Get(url) //Fazendo a requisição
	if erro != nil {
		return nil, fmt.Errorf("erro na conexão com a API: %v", erro)
	} 
	defer resp.Body.Close() //Fechando a conexão da resposta da API

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("api tornou erro: %s", resp.Status)
	}

	//Cria uma váriavel para armazenar a resposta convertida
	var data models.WeatherData

	//Converte o json em struct
	if erro := json.NewDecoder(resp.Body).Decode(&data); erro != nil {  //Lê o jason e coloca a resposta na váriavel data
		return nil, fmt.Errorf("erro ao codificar o JSON: %v", erro)
	}

	//Retorna dados para o handler
	return &data, nil
}