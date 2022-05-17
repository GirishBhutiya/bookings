package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/GirishBhutiya/bookings/pkg/models"

	"github.com/GirishBhutiya/bookings/pkg/config"
)

const HtmlTemplatePath = "."

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplate sets the config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateDate) *models.TemplateDate {
	return td
}

//Render HTML templates
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateDate) {
	//path := "./templates/" + tmpl

	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser :", err)
	}

}

//create template cache in map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(HtmlTemplatePath + "/templates/*.page.tmpl")
	if err != nil {
		fmt.Println("Erros is :", err)
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Page is currently ", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			fmt.Println("Erros is :", err)
			return myCache, err
		}
		matches, err := filepath.Glob(HtmlTemplatePath + "/templates/*.layout.tmpl")
		if err != nil {
			fmt.Println("Erros is :", err)
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(HtmlTemplatePath + "/templates/*.layout.tmpl")
			if err != nil {
				fmt.Println("Erros is :", err)
				return myCache, err
			}
		}
		myCache[name] = ts

	}
	return myCache, nil

}
