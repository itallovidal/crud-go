package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"rocketseat-desafio-1/dtos"
	"rocketseat-desafio-1/helpers"
	"rocketseat-desafio-1/models"

	"github.com/google/uuid"
)


func CreateUserController(db map[string]models.User) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request ) {
		body, bodyError := io.ReadAll(r.Body)

		if bodyError != nil {
			response := dtos.GenericResponse{
				Message: "Erro interno de servidor.",
				Status: http.StatusInternalServerError,
			}

			helpers.SendResponse(w, response, http.StatusInternalServerError)
			return
		}

		var userList []models.User

		jsonError := json.Unmarshal(body, &userList)

		if jsonError != nil {
			response := dtos.GenericResponse{
				Message: "Erro interno de servidor.",
				Status: http.StatusInternalServerError,
			}
			
			helpers.SendResponse(w, response, http.StatusInternalServerError)
			return
		}

		for _, newUser := range userList {
			id := uuid.NewString()
			db[id] = newUser
		}

		response := dtos.GenericResponse{
			Message: "Usuários cadastrados com sucesso!",
			Status: http.StatusCreated,
		}
		
		helpers.SendResponse(w, response, http.StatusCreated)
	}
}