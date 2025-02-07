package controllers

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"rocketseat-desafio-1/interfaces"

	"github.com/google/uuid"
)


func CreateUserController(db map[string]interfaces.User) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request ) {
		body, bodyError := io.ReadAll(r.Body)

		if bodyError != nil {
			slog.Error("Body parse error.")
			w.Write([]byte("Erro interno de servidor."))
		}

		var userList []interfaces.User

		jsonError := json.Unmarshal(body, &userList)

		if jsonError != nil {
			slog.Error("Json parse error.")
			w.Write([]byte("Erro interno de servidor."))
		}

		for _, newUser := range userList {
			id := uuid.NewString()
			db[id] = newUser
		}

		response := interfaces.GenericResponse{
			Message: "Usuários cadastrados com sucesso!",
			Status: http.StatusCreated,
		}

		jsonResponse, _ := json.Marshal(response)
		
		w.Write(jsonResponse)
	}
}