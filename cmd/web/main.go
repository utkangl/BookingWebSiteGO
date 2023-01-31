package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/utkangl/GoWEB/pkg/config"
	"github.com/utkangl/GoWEB/pkg/handlers"
	"github.com/utkangl/GoWEB/pkg/render"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	tempCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tempCache

	repository := handlers.CreateRepo(&app)
	handlers.SetRepo(repository)

	render.SetConfig(&app)

	http.HandleFunc("/", handlers.Repo.HomePage)
	http.HandleFunc("/about", handlers.Repo.AboutPage)

	fmt.Println("Starting the Application on Port: ", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
