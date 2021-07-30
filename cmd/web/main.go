package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ashwani00002/go-course/pkg/config"
	"github.com/Ashwani00002/go-course/pkg/handlers"
	"github.com/Ashwani00002/go-course/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.Appconfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

// Use go run  cmd/web/*.go for compiling,
//go will fetch other *.go files using dependencies defined in main.go
