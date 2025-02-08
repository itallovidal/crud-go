package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"rocketseat-desafio-1/dtos"
	"rocketseat-desafio-1/helpers"
	"rocketseat-desafio-1/models"

	"github.com/go-chi/chi/v5"
)


func UpdateUser(db map[string]models.User) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request ) {
		id := chi.URLParam(r, "id")

		_, exists := db[id]

		if !exists {
			response := dtos.GenericResponse{
				Message: "Não é possível atualizar um usuário inexistente.",
				Status: http.StatusBadRequest,
			}

			helpers.SendResponse(w, response, http.StatusBadRequest)
			return
		} 

		body, bodyParserError := io.ReadAll(r.Body)

		if bodyParserError != nil{
			response := dtos.GenericResponse{
				Message: "Erro interno do servidor.",
				Status: http.StatusInternalServerError,
			}

			helpers.SendResponse(w, response, http.StatusInternalServerError)
			return
		}

		var updatedUser models.User
		updateParserError := json.Unmarshal(body, &updatedUser)

		if updateParserError != nil{
			response := dtos.GenericResponse{
				Message: "Erro interno do servidor.",
				Status: http.StatusInternalServerError,
			}

			helpers.SendResponse(w, response, http.StatusInternalServerError)
			return
		}

		db[id] = updatedUser

		response := dtos.GenericResponse{
			Message: "Usuário atualizado com sucesso!",
			Status: http.StatusNoContent,
		}

		helpers.SendResponse(w, response, http.StatusNoContent)
	}
}