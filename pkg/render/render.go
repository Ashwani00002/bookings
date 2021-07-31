package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Ashwani00002/go-course/pkg/config"
	"github.com/Ashwani00002/go-course/pkg/models"
)

var functions = template.FuncMap{}
var app *config.Appconfig

// NewTemplate sets the config for the template package
func NewTemplate(a *config.Appconfig) {
	app = a
}

// Render templates for handlers
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	// get template cache from app config
	var tc map[string]*template.Template
	if app.UseCache {
		// Read information from template cache & don't rebuild, for production env
		tc = app.TemplateCache
	} else {
		// rebuild template cache if app.UseCache = false, for development env
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Page is currently at:", name)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, nil
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
