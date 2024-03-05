package main

import (
	"fmt"
	"log"
	"mygithub/pkg/config"
	"mygithub/pkg/handlers"
	"mygithub/pkg/render"
	"net/http"
)

const port_number = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UserCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println(fmt.Sprintf("Starting application on post %s", port_number))
	_ = http.ListenAndServe(port_number, nil)

}
