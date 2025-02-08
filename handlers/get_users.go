package handlers

import (
	"encoding/json"
	"net/http"
	"rocketseat-desafio-1/models"
)

type GetAllUsersResponse struct {
	Users []models.User
	Status int32
}

func GetAllUsers(db map[string]models.User) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request ) {
		
		var userList []models.User
		for _, user := range db {
			userList = append(userList, user)
		}

		response := GetAllUsersResponse{
			Users: userList,
			Status: http.StatusOK,
		}

		jsonResponse, _ := json.Marshal(response)
		
		w.Write(jsonResponse)
	}
}