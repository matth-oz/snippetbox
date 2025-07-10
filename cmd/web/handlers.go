package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Display the home page
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Add custom header
	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// Display a specific snippet
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet ID = %d", id)
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
