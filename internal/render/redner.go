package render

import (
	"booking/internal/config"
	"booking/internal/models"
	"bytes"
	"errors"
	"fmt"
	"github.com/justinas/nosurf"
	"html/template"
	"net/http"
	"path/filepath"
)

// functions we can use them in templates
var functions = template.FuncMap{}

var app *config.AppConfig
var pathToTemplates = "../templates"

// NewRender sets the config for the template package
func NewRenderer(a *config.AppConfig) {
	app = a
}

func AddDefaultData(dt *models.TemplateData, r *http.Request) *models.TemplateData {
	dt.Flash = app.Session.PopString(r.Context(), "flash")
	dt.Error = app.Session.PopString(r.Context(), "error")
	dt.Warning = app.Session.PopString(r.Context(), "warning")
	dt.CSRFToken = nosurf.Token(r)
	return dt
}

// Template renders templates using html/template
func Template(w http.ResponseWriter, r *http.Request, tmpl string, dt *models.TemplateData) error {
	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		return errors.New("cant get template from cache")
	}

	dt = AddDefaultData(dt, r)
	buf := new(bytes.Buffer)

	//store temp in buf var
	_ = t.Execute(buf, dt)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
		return err
	}

	return nil

}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.gohtml", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.gohtml", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.gohtml", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
