package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// Chaange the signature of the home handler so it defined as a method against *application
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	// Initialize a slice containing the paths to the two files. It's important
	// to note that the file containing our base template must be the *first*
	// file in the slice.

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	// Use the template.ParseFiles() function to read the files and store the
	// templates in a template set. Notice that we can pass the slice of file
	// paths as a variadic parameter?
	ts, err := template.ParseFiles(files...)
	if err != nil {
		//because the home Handler function is now a method against application
		//it can access its fields, including the error logger. We'll write the log
		//message to this instead of the standard logger
		app.ServerError(w, err)
		// log.Print(err.Error())
	}

	// Use the ExecuteTemplate() method to write the content of the "base"
	// template as the response body.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.ServerError(w, err)
		// log.Print(err.Error())
	}

	// w.Write([]byte("Hello from snippetbox"))
}

// add a snippetView Handler - view  specific response
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Displaying a specific snippet with ID %d...", id)
}

// create a new snippet
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("creating a new snippet..."))
}
