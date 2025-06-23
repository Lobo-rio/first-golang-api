package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"modules/src/authentication"
	"modules/src/database"
	"modules/src/models"
	"modules/src/repositories"
	"modules/src/responses"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateNote function that adds a new annotation to the database
func CreateNote(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	request, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return 
	}

	var note models.Note
	if err = json.Unmarshal(request, &note); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return 
	}

	note.AuthorID = userID

	if err = note.Prepare("registration"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return 
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return 
	}
	defer db.Close()

	repository := repositories.NoteRepo(db)
	note.ID, err = repository.Create(note)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return 
	}

	responses.JSON(w, http.StatusCreated, note)
}

// GetAllNotes function that retrieves all notes saved in the database
func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return 
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return 
	}
	defer db.Close()

	repository := repositories.NoteRepo(db)
	notes, err := repository.GetAll(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return 
	}

	responses.JSON(w, http.StatusOK, notes)
}

// GetByIDNote function that brings a single note
func GetByIDNote(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	noteID, err := strconv.ParseUint(parameters["noteId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NoteRepo(db)
	noteDatabase, err := repository.GetByID(noteID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if noteDatabase.AuthorID != userID {
		responses.Error(w, http.StatusForbidden, errors.New("não é possível visualizar uma anotação que não seja sua"))
		return
	}

	responses.JSON(w, http.StatusOK, noteDatabase)
}

// UpdateNote function that changes the data of a note
func UpdateNote(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	parameters := mux.Vars(r)
	noteID, erro := strconv.ParseUint(parameters["noteId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NoteRepo(db)
	noteDatabase, err := repository.GetByID(noteID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	if noteDatabase.AuthorID != userID {
		responses.Error(w, http.StatusForbidden, errors.New("não é possível visualizar uma anotação que não seja sua"))
		return
	}
	
	request, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	
	var note models.Note
	if err = json.Unmarshal(request, &note); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	
	
	if err = note.Prepare("edit"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	
	if err = repository.Update(noteID, note); err != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteNote function that removes a note
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	noteID, err := strconv.ParseUint(parameters["noteId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NoteRepo(db)
	noteDatabase, err := repository.GetByID(noteID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if noteDatabase.AuthorID != userID {
		responses.Error(w, http.StatusForbidden, errors.New("não é possível visualizar uma anotação que não seja sua"))
		return
	}

	if err = repository.Delete(noteID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}