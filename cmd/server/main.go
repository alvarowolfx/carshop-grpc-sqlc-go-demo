package main

import (
	"log"

	"com.aviebrantz.carshop/pkg/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	log.Fatal(server.Run())
}
