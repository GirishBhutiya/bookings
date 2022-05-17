package main

import (
	"log"
	"net/http"
	"time"

	"github.com/GirishBhutiya/bookings/pkg/config"
	"github.com/GirishBhutiya/bookings/pkg/handlers"
	"github.com/GirishBhutiya/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"
const HomeLink = "/"
const AboutPageLink = "/about"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//set to true in production
	app.InProduction = false

	app.UseCache = false

	session = scs.New()
	session.Cookie.Persist = true
	session.Lifetime = 24 * time.Hour
	session.Cookie.Secure = app.InProduction
	session.Cookie.SameSite = http.SameSiteLaxMode

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can not create template cache")
	}

	app.TemplateCache = tc

	render.NewTemplate(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// http.handlefunc(homelink, handlers.repo.home)
	// http.HandleFunc(AboutPageLink, handlers.Repo.About)
	//_ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr: portNumber,
		//Handler: routes(&app),
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
