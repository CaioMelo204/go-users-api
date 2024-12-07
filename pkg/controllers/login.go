package controllers

import (
	"devbook-api/pkg/auth"
	"devbook-api/pkg/database"
	"devbook-api/pkg/models"
	"devbook-api/pkg/repositories"
	"devbook-api/pkg/response"
	"devbook-api/pkg/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	var got models.User

	repo := repositories.NewUserRepository(db)
	got, err = repo.GetByEmail(user.Email)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = security.CheckPasswordHash(user.Password, got.Password)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, _ := auth.CreateToken(got.Id)
	response.JSON(w, http.StatusOK, models.Auth{
		Id:    got.Id,
		Token: token,
	})
}
