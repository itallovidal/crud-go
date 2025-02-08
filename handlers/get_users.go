package handlers

import (
	"net/http"
	"rocketseat-desafio-1/helpers"
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

		helpers.SendResponse(w, response, http.StatusOK)
	}
}