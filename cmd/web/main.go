package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/amir-moshfegh/web-game/pkg/config"
	"github.com/amir-moshfegh/web-game/pkg/handler"
	"github.com/amir-moshfegh/web-game/pkg/render"
	"log"
	"net/http"
	"time"
)

const addr = ":8080"

var session *scs.SessionManager
var app config.AppConfig

func main() {
	//TODO:: change true after protecion
	app.Protection = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = false
	session.Cookie.Secure = app.Protection
	session.Cookie.SameSite = http.SameSiteLaxMode
	app.Session = session

	ts, err := render.CreateRenderCache()
	if err != nil {
		log.Fatalln("Coud not create template cache.")
	}
	app.TemplateCache = ts
	app.UseCache = false
	repo := handler.NewRepo(&app)
	handler.NewHandler(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    addr,
		Handler: routes(),
	}

	_ = srv.ListenAndServe()
}
