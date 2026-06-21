package main

import (
	"fmt"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	writeJsonError(w, http.StatusInternalServerError, "The server encountered a problem and could not process your request")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	writeJsonError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	app.logError(r, fmt.Errorf("resource not found"))
	writeJsonError(w, http.StatusNotFound, "The requested resource was not found")
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request) {
	app.logError(r, fmt.Errorf("resource conflict"))
	writeJsonError(w, http.StatusConflict, "The resource already exists")
}

func (app *application) logError(r *http.Request, err error) {
	fmt.Printf("[%s] %s %s: %v\n", r.Method, r.URL.Path, r.RemoteAddr, err)
}
