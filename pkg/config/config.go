package config

import (
	"html/template"
	"log"
)

// Appconfig holds the application config
type Appconfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
