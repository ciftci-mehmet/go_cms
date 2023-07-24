package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/joho/godotenv"

	"github.com/ciftci-mehmet/go_cms/cmd/web/handlers"
	db "github.com/ciftci-mehmet/go_cms/db/sqlc"

	_ "github.com/lib/pq"
)

func initApplication() *Application {

	// read .env
	err := godotenv.Load()
	if err != nil {
		color.Red("could not load .env file, rename \".env.example\" to \".env\" and configure it")
		log.Fatal(err)
	}

	secure, err := strconv.ParseBool(os.Getenv("SERVER_SECURE"))
	if err != nil {
		log.Fatal(err)
	}

	app := &Application{
		Server: Server{
			ServerName: os.Getenv("SERVER_NAME"),
			Port:       os.Getenv("SERVER_PORT"),
			Secure:     secure,
			URL:        os.Getenv("SERVER_URL"),
		},
		Handlers: &handlers.Handlers{},
	}

	app.Routes = app.createRoutes()

	var dbQueries *db.Queries
	dbDriver := os.Getenv("DB_DRIVER")
	dbSource := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_URL"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"))
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	dbQueries = db.New(conn)

	app.Query = dbQueries
	app.Handlers.Query = dbQueries

	return app
}
