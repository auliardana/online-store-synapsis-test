package main

import (
	"online-store/app/db"
	"online-store/app/router"
	"os"

	"github.com/joho/godotenv"
)

// @title           API Documentation for online-store rest api
// @version         1.0
// @description     synapsis.id test
// @termsOfService  http://swagger.io/terms/

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	// fmt.Println("Hello World")

	if err := godotenv.Load(".env"); err != nil {
		panic("ERROR: Could not load .env")
	}

	db.ConnectDatabase()
	r := router.SetupRouter()
	r.Run(":" + os.Getenv("DEV_PORT"))

}
