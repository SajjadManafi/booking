package main

import (
	"booking/internal/config"
	"booking/internal/handler"
	"booking/internal/models"
	"booking/internal/render"
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const Port = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// what am I going to put in the session
	gob.Register(models.Reservation{})

	// change this to true in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handler.NewRepo(&app)
	handler.NewHandler(repo)

	render.NewTemplates(&app)

	//http.HandleFunc("/", handler.Repo.Home)
	//http.HandleFunc("/About", handler.Repo.About)

	fmt.Println("Starting App on Port ", Port)
	//_ = http.ListenAndServe(Port, nil)

	srv := &http.Server{
		Addr:    Port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
