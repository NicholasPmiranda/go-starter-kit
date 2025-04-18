package main

import (
	logger "boilerPlate/config/looger"
	"fmt"

	"github.com/joho/godotenv"

	"boilerPlate/database/seeds"
	"boilerPlate/routes"
)

func main() {

	logger.SetupLogger()

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	seeds.Run()

	router := routes.SetupRoutes()

	router.Run(":3030")

}
