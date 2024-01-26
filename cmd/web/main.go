package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/stellyes/go-bed-n-breakfast/pkg/config"
	"github.com/stellyes/go-bed-n-breakfast/pkg/handlers"
	"github.com/stellyes/go-bed-n-breakfast/pkg/render"
)

const portNumber = ":4000"

var app config.AppConfig
var session *scs.SessionManager

// Lowercase function names make them private
// to this specific module
func main() {
	// !!! CHANGE THIS TO TRUE WHEN IN PRODUCTION !!!
	app.InProduction = false

	// Create new user session
	session = scs.New()
	session.Lifetime = 24 * time.Hour // 24 hour expire time
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	// In production, this encrypts cookie (https)
	session.Cookie.Secure = app.InProduction

	// Set configuration session to current user session
	app.Session = session

	// Templates are HTML fragments, so to speak
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	_, _ = fmt.Println(fmt.Sprintf("Starting applicatiton on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
