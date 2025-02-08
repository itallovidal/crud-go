package handlers

import (
	"encoding/json"
	"net/http"
	"rocketseat-desafio-1/models"

	"github.com/go-chi/chi/v5"
)


type GetUserByIdResponse struct{
	Status 	int32
	User 	*models.User 	`json:",omitempty"`
	Message string 				`json:",omitempty"`
}



func GetUserById(db map[string]models.User) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		id := chi.URLParam(r, "id")

		user, exists := db[id]

		if !exists {
			w.WriteHeader(http.StatusNotFound)

			response := GetUserByIdResponse{
				Status: 404,
				Message: "Usuário não encontrado.",
				User: nil,
			}

			jsonReponse, _ := json.Marshal(response)

			w.WriteHeader(http.StatusNotFound)
			w.Write(jsonReponse)
			return
		}

		response := GetUserByIdResponse{
			Status: 200,
			User: &user,
		}

		jsonReponse, _ := json.Marshal(response)

		w.WriteHeader(http.StatusOK)
		w.Write(jsonReponse)
	}
}