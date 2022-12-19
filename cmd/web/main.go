package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/buscaroli/secondGoWebApp/pkg/config"
	"github.com/buscaroli/secondGoWebApp/pkg/handlers"
	"github.com/buscaroli/secondGoWebApp/pkg/render"
)

const port = ":8080"

// making app and session available to the entire main package
var app config.AppConfig
var session *scs.SessionManager

func main() {
	// set to true once deploying for production
	app.IsProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour // default is 24h
	session.Cookie.Persist = true     // persist after closing browser
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.IsProduction // set true in production

	// adding session to the Repo in order to access it app wide
	app.Session = session

	// create the templates for every page and save them into a cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	// the cached templates are saven in the global config file
	app.TemplateCache = tc

	// In development mode set to false so any change to a template will be reflected in the browser without having to restart the server
	app.UseCache = app.IsProduction

	// the next two lines allows us to access the config from within the handlers package
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// the next line allows us to access the config from within the render package
	render.NewTemplates(&app)

	fmt.Println("Server up and running on port", port)
	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
