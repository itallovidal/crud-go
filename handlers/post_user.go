package handlers

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"rocketseat-desafio-1/dtos"
	"rocketseat-desafio-1/models"

	"github.com/google/uuid"
)


func CreateUserController(db map[string]models.User) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request ) {
		body, bodyError := io.ReadAll(r.Body)

		if bodyError != nil {
			slog.Error("Body parse error.")
			w.Write([]byte("Erro interno de servidor."))
		}

		var userList []models.User

		jsonError := json.Unmarshal(body, &userList)

		if jsonError != nil {
			slog.Error("Json parse error.")
			w.Write([]byte("Erro interno de servidor."))
		}

		for _, newUser := range userList {
			id := uuid.NewString()
			db[id] = newUser
		}

		response := dtos.GenericResponse{
			Message: "Usuários cadastrados com sucesso!",
			Status: http.StatusCreated,
		}

		jsonResponse, _ := json.Marshal(response)
		
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResponse)
	}
}