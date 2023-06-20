package main

import (
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	// read .env
	err := godotenv.Load()
	if err != nil {
		color.Red("could not load .env file, rename \".env.example\" to \".env\" and configure it")
		log.Fatal(err)
	}

	projectName := os.Getenv("PROJECT_NAME")
	println("project name", projectName)

}
