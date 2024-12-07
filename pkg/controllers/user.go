package controllers

import (
	"devbook-api/pkg/database"
	"devbook-api/pkg/models"
	"devbook-api/pkg/repositories"
	"devbook-api/pkg/response"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = user.Prepare("create"); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	newUser, err := repo.Insert(user)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	response.JSON(w, http.StatusCreated, newUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	userID, err := strconv.ParseUint(param["userId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	repo.Delete(userID)
	response.JSON(w, http.StatusOK, nil)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	userID, err := strconv.ParseUint(param["userId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = user.Prepare(""); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	repo.Update(userID, user)

	response.JSON(w, http.StatusOK, nil)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	userID, err := strconv.ParseUint(param["userId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	user, err := repo.Get(userID)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	response.JSON(w, http.StatusOK, user)
}

func GetUserList(w http.ResponseWriter, r *http.Request) {
	nameOrNick := r.URL.Query().Get("user")

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	users, err := repo.List(nameOrNick)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	response.JSON(w, http.StatusOK, users)
}
