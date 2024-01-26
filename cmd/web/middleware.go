package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	// Returns anonymous function, casts to handler function,
	// serving our middlware to client
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("We hit the page!")
		next.ServeHTTP(w, r)
	})
}

// Adds cross-site request forgery protection for all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",              // Referencing the "entire" site
		Secure:   app.InProduction, // For HTTPS, which I'm not using rn
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads and saves session on every request
func SessionLoad(next http.Handler) http.Handler {
	// CMD + Click on Load and Save to see the logic
	return session.LoadAndSave(next)
}
