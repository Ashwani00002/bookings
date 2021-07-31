package main

import (
	"net/http"

	"github.com/Ashwani00002/go-course/pkg/config"
	"github.com/Ashwani00002/go-course/pkg/handlers"
	"github.com/bmizerany/pat"
)

func routes(app *config.Appconfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}
