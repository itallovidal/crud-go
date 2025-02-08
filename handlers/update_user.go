package handlers

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"rocketseat-desafio-1/dtos"
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
	
			jsonResponse, _ := json.Marshal(response)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonResponse)
			return
		} 

		body, bodyParserError := io.ReadAll(r.Body)

		if bodyParserError != nil{
			slog.Error("Body parser error.")

			response := dtos.GenericResponse{
				Message: "Erro interno do servidor.",
				Status: http.StatusInternalServerError,
			}

			jsonResponse, _ := json.Marshal(response)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonResponse)
		}

		var updatedUser models.User
		updateParserError := json.Unmarshal(body, &updatedUser)

		if updateParserError != nil{
			slog.Error("Update parser error.")

			response := dtos.GenericResponse{
				Message: "Erro interno do servidor.",
				Status: http.StatusInternalServerError,
			}

			jsonResponse, _ := json.Marshal(response)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonResponse)
		}

		db[id] = updatedUser

		response := dtos.GenericResponse{
			Message: "Usuário atualizado com sucesso!",
			Status: http.StatusNoContent,
		}

		jsonResponse, _ := json.Marshal(response)
		w.WriteHeader(http.StatusNoContent)
		w.Write(jsonResponse)
	}
}