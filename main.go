package main

import (
	"myapp/config"
	"myapp/database"
	"myapp/routes"
)

func main() {
	config.LoadEnv()        // .env
	db := database.InitDB() // Conecta ao banco

	r := routes.SetupRoutes(db)
	r.Run(":8080") // inicia servidor
}
