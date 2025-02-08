package helpers

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, response any, status int){
	jsonResponse, _ := json.Marshal(response)
	w.WriteHeader(status)
	w.Write(jsonResponse)
}