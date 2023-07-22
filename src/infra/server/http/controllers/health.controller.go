package controllers

import "net/http"

func HealthController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{\"msg\": \"Hello World!\"}"))
}
