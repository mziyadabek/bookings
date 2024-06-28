package main

import (
	"github.com/mziyadabek/bookings/pkg/config"
	"github.com/mziyadabek/bookings/pkg/handlers"
	"github.com/mziyadabek/bookings/pkg/render"
	"log"
	"net/http"
)

const PortNum = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	_ = http.ListenAndServe(PortNum, nil)
}
