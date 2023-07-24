package main

import (
	"net/http"
	"os"

	"github.com/fatih/color"
	_ "github.com/lib/pq"
)

func main() {

	projectName := os.Getenv("PROJECT_NAME")
	println("project name", projectName)

	app := initApplication()
	//app.ListenAndServe()

	err := http.ListenAndServe(":3000", app.Routes)
	if err != nil {
		color.Red("error %s", err)
		return
	}
}
