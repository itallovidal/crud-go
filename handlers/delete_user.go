package handlers

import (
	"encoding/json"
	"net/http"
	"rocketseat-desafio-1/dtos"
	"rocketseat-desafio-1/models"

	"github.com/go-chi/chi/v5"
)


func DeleteUser(db map[string]models.User) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request ) {
		id := chi.URLParam(r, "id")

		_, exists := db[id]

		if !exists {
			response := dtos.GenericResponse{
				Message: "Não é possível deletar um usuário inexistente.",
				Status: http.StatusBadRequest,
			}
	
			jsonResponse, _ := json.Marshal(response)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonResponse)
			return
		} 

		delete(db, id)
		response := dtos.GenericResponse{
			Message: "Usuário deletado com sucesso!",
			Status: http.StatusOK,
		}

		jsonResponse, _ := json.Marshal(response)
		w.Write(jsonResponse)
	}
}