package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/matth-oz/snippetbox/internal/models"
)

// Display the home page
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Add custom header
	w.Header().Add("Server", "Go")

	snippets, err := app.Snippets.Latest()

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, r, http.StatusOK, "home.tmpl", data)
}

// Display a specific snippet
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.Snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, r, http.StatusOK, "view.tmpl", data)
	// fmt.Fprintf(w, "%+v", data)
}

// Display a form for creating a new snippet
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form creating a new snippet…"))
}

// Create a new snippet
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.Snippets.Insert(title, content, expires)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
	//w.WriteHeader(http.StatusCreated)
	//w.Write([]byte("Save a new snippet…"))
}
