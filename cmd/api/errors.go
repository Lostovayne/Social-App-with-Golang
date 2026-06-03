package main

import "net/http"

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	writeJsonError(w, http.StatusInternalServerError, "The server encountered a problem and could not process your request")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	writeJsonError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	writeJsonError(w, http.StatusNotFound, "The requested resource was not found")
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	writeJsonError(w, http.StatusConflict, "The resource already exists")
}
