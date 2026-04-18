package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	// Data Map
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}

	if err := writeJson(w, http.StatusOK, data); err != nil {
		writeJsonError(w, http.StatusInternalServerError, "error encoding response data: "+err.Error())

	}

}
