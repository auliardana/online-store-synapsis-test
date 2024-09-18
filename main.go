package main

import (
	"fmt"

	"online-store/app/db"
	"online-store/app/router"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World")

	if err := godotenv.Load(".env"); err != nil {
		panic("ERROR: Could not load .env")
	}

	db.ConnectDatabase()
	r := router.SetupRouter()
	r.Run("localhost:8080")

}
