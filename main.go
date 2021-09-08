package main

import (
	"Balance/internal/server"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	app := server.NewServer()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to export env vars: %v", err)
	}

	port, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		log.Fatalf("Failed to export port from env")
	}

	if err := app.Run(port); err != nil {
		log.Fatal(err)
	}
}