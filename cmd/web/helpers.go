package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// The ServerError helper writes an error message and stack trace to the errorLog,
// then sends a generic internal Server Error Response to the user.
func (app *application) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Print(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The cleintError helper sends a specific status code and correspnindg description
// to the user. We'll use this in the book to send responses like 400 "Bad Request" when there's a problem with
// the request that the user sents.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// For consistency, we'll also implement a notFound helper. This is simply
// conveniece wrapper around clientError which sends a 404 Not Found response to the User
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
