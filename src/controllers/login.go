package controllers

import (
	"encoding/json"
	"io"
	"modules/src/authentication"
	"modules/src/database"
	"modules/src/models"
	"modules/src/repositories"
	"modules/src/responses"
	"modules/src/security"
	"net/http"
	"strconv"
)

// Login function that is responsible for authenticating a user in the API
func Login(w http.ResponseWriter, r *http.Request) {
	request, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(request, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.UserRepo(db)
	userDatabase, err := repository.GetByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.Verify(userDatabase.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(userDatabase.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userID := strconv.FormatUint(userDatabase.ID, 10)

	responses.JSON(w, http.StatusOK, models.AuthenticationData{ID: userID, Token: token})
}
