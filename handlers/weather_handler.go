package handlers

import(
	"github.com/gofiber/fiber/v2"
	"github.com/francinemaaciel/estudos-go/services"
	"os" //Pacote padrão que permite interagir com o sistema operacional
)
//Função que trata a rota do Get
//c da biblioteca do fiber é o contetxo de requisição -- vai acessar os parâmetros da URL
func GetWeather(c *fiber.Ctx) error{
	cidade := c.Params("cidade") //Acessa e pega a cidade da URL

	if cidade == ""{
		//Se nenhuma cidade for informada ele retorna o erro 400
		return c.Status(400).JSON(fiber.Map{  // .JSON(fiber.Map) vai criar um objeto json para retornar uam repsosta sobre o erro
			"erro" : "você não forneceu nenhuma cidade",
		})

	}
	//Consulta a chave da API
	apiKey := os.Getenv("API_WEATHER_KEY")
	if apiKey == "" {
		//500 Erro Interno do Servidor
		return c.Status(500).JSON(fiber.Map{ 
			"erro" : "chave da api não encontrada",
		})
	}

	//Chama o services que consulta a APi
	weather, erro := services.GetWeatherAPI(cidade, apiKey)
	if erro != nil{
		return c.Status(500).JSON(fiber.Map{
			"erro" : "consulta na api falhou",
		})
	}
	//Se tudo estiver certo, retorna o clima
	return c.JSON(weather)
}