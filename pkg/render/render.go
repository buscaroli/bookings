package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/buscaroli/secondGoWebApp/pkg/config"
	"github.com/buscaroli/secondGoWebApp/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		// get the templates cache from app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from cache")
	}

	// creating a buffer for easier error debugging
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	//render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all the files named *.page.html
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.html (collected with the previous loop)
	for _, page := range pages {
		name := filepath.Base(page)
		// Create the template associated with the given filename
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// get all the layout files
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		// if there is at least one layout file use it together with the page layout to create a full page, if not skip this step as ts is already a full page (with no dynamic content)
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, err
}
