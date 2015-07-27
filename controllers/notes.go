package controllers

import (
	"net/http"
	"strconv"

	"github.com/quorumsco/contacts/models"
	"github.com/quorumsco/contacts/views"
	. "github.com/quorumsco/jsonapi"
	"github.com/quorumsco/logs"
	"github.com/quorumsco/router"
)

func RetrieveNotesByContact(w http.ResponseWriter, req *http.Request) {
	db := getDB(req)
	contactStore := models.ContactStore(db)

	contactID, err := strconv.Atoi(router.Context(req).Param("id"))
	if err != nil {
		logs.Debug(err)
		Fail(w, req, map[string]interface{}{"id": "not integer"}, http.StatusBadRequest)
		return
	}

	var c = models.Contact{
		ID: uint(contactID),
	}

	err = contactStore.FindNotes(&c)
	if err != nil {
		logs.Error(err)
		Error(w, req, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	Success(w, req, views.Notes{Notes: c.Notes}, http.StatusOK)
}

func DeleteNote(w http.ResponseWriter, req *http.Request) {
	db := getDB(req)
	noteStore := models.NoteStore(db)

	contactID, err := strconv.Atoi(router.Context(req).Param("id"))
	if err != nil {
		logs.Debug(err)
		Fail(w, req, map[string]interface{}{"id": "not integer"}, http.StatusBadRequest)
		return
	}

	noteID, err := strconv.Atoi(router.Context(req).Param("note_id"))
	if err != nil {
		logs.Debug(err)
		Fail(w, req, map[string]interface{}{"id": "not integer"}, http.StatusBadRequest)
		return
	}

	var n = new(models.Note)
	n.ID = uint(noteID)
	n.ContactID = uint(contactID)
	err = noteStore.Delete(n)
	if err != nil {
		logs.Debug(err)
		Fail(w, req, map[string]interface{}{"id": "not integer"}, http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	Success(w, req, nil, http.StatusNoContent)
}
