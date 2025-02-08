package handlers

import (
	"net/http"
	"rocketseat-desafio-1/helpers"
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
			response := GetUserByIdResponse{
				Status: 404,
				Message: "Usuário não encontrado.",
				User: nil,
			}

			helpers.SendResponse(w, response, http.StatusNotFound)
			return
		}

		response := GetUserByIdResponse{
			Status: 200,
			User: &user,
		}

		helpers.SendResponse(w, response, http.StatusOK)
	}
}