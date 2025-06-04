package main

import(
	"log"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2"
	"github.com/francinemaaciel/estudos-go/handlers"
)

func main()  {

	//Carrega a chave da api do arquivo .env
	erro := godotenv.Load()
	if erro != nil{
		log.Fatal("erro ao carregar o arquivo .env")
	}

	app := fiber.New()

	app.Static("/", "./views")

	//Registra a rota da API
	app.Get("/clima/:cidade", handlers.GetWeather) //O :cidade é um parâmetro dinâmico
	//handlers.GetWeather é a função que vai ser chamada quando essa rota for acessada

	app.Listen(":3000") //Inicia o servidor pela porta 3000 do pc
}