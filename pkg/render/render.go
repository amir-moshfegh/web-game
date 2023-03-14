package render

import (
	"bytes"
	"github.com/amir-moshfegh/web-game/pkg/config"
	"github.com/amir-moshfegh/web-game/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func defaultTemplateData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var cache map[string]*template.Template

	if app.UseCache {
		cache = app.TemplateCache
	} else {
		cache, _ = CreateRenderCache()
	}

	t, ok := cache[tmpl]
	if !ok {
		log.Fatalln("can't find this page")
	}

	buf := new(bytes.Buffer)
	td = defaultTemplateData(td)
	err := t.Execute(buf, td)

	if err != nil {
		log.Fatalln(err)
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatalln(err)
	}
}

func CreateRenderCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pageGlob := "./templates/*.page.tmpl"
	layoutGlob := "./templates/*.layout.tmpl"

	pages, err := filepath.Glob(pageGlob)
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(layoutGlob)
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(layoutGlob)
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
