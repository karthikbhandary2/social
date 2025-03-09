package main

import (
	"net/http"

	"github.com/karthikbhandary2/social/internal/store"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Working"))

	app.store.Posts.Create(r.Context(), &store.Post{})
}