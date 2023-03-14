package main

import (
	"github.com/amir-moshfegh/web-game/pkg/handler"
	"github.com/go-chi/chi"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	
	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)

	return mux
}
