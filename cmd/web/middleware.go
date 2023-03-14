package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

func NoSurf(next http.Handler) http.Handler {
	nosurfHandler := nosurf.New(next)
	nosurfHandler.SetBaseCookie(http.Cookie{
		Path:     "/",
		Secure:   app.Protection,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return nosurfHandler
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
