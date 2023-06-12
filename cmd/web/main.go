package main

import (
	"fmt"
	"log"
	"time"

	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/waterpigeon420/bookings/pkg/config"
	"github.com/waterpigeon420/bookings/pkg/handlers"
	"github.com/waterpigeon420/bookings/pkg/render"
)

var port = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	//lmao
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app) //creating a new repository calls NewRepo
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting server %s", port))

	srv := &http.Server{
		Addr:    port,
		Handler: routes((&app)),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
	//bruh
}
