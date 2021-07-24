package main

import (
	"fmt"
	"net/http"

	"github.com/Ashwani00002/go-course/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.Home)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

// Use go run  cmd/web/*.go for compiling,
//go will fetch other *.go files using dependencies defined in main.go
