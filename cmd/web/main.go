package main

import (
	"log"

	"github.com/krzysztofkaptur/basic-go-website/pkg/config"
	handlers "github.com/krzysztofkaptur/basic-go-website/pkg/handlers"
	"github.com/krzysztofkaptur/basic-go-website/pkg/render"
)

func main() {
	app := config.AppConfig{}

	tc, err := render.CreateCacheTemplate()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	RunServer(repo)
}
