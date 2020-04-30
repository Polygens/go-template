package handler

import "net/http"

// Health will respond with "pong" and a status 200 code
func Health(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte("pong"))
}
