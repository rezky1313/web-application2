package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

/*
Install a middleware
Middleware actualy allows you to process a request as it comes into your web
and perform an action on it

every time somebody hits the page just write something to the console
*/

func WritetoConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit The Page")
		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: false,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}
