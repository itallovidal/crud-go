package controllers

import "net/http"

func UserController(w http.ResponseWriter, r *http.Request )   {
	w.Write([]byte("Oi!"))
}