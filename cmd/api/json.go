package main

import (
	"encoding/json"
	"net/http"
)

// writeJson se encarga de escribir la respuesta en formato JSON. Recibe el status code y los datos a enviar, establece el header Content-Type a application/json, escribe el status code y codifica los datos en JSON para enviarlos al cliente.
func writeJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

// readJson se encarga de leer el JSON del body de la petición y decodificarlo en la estructura que le pasemos como argumento. Además, limita el tamaño del body para evitar ataques de denegación de servicio (DoS) y deshabilita la decodificación de campos desconocidos para evitar errores silenciosos.
func readJson(w http.ResponseWriter, r *http.Request, data any) error {
	// Block Atacks
	maxBytes := 1_048_576 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(data)

}

// writeJsonError es una función auxiliar que se encarga de enviar un mensaje de error en formato JSON. Recibe el status code y el mensaje de error, crea una estructura con el campo "error" y llama a writeJson para enviar la respuesta al cliente.
func writeJsonError(w http.ResponseWriter, status int, message string) error {

	type envelope struct {
		Error string `json:"error"`
	}

	return writeJson(w, status, &envelope{Error: message})
}
