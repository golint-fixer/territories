package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/quorumsco/contacts/models"
	"github.com/quorumsco/contacts/views"
	. "github.com/quorumsco/jsonapi"
	"github.com/quorumsco/logs"
	"github.com/quorumsco/router"
)

func RetrieveNoteCollection(w http.ResponseWriter, r *http.Request) {
	var (
		contactID int
		err       error
	)
	contactID, err = strconv.Atoi(router.Context(r).Param("id"))
	if err != nil {
		logs.Debug(err)
		Fail(w, r, map[string]interface{}{"id": "not integer"}, http.StatusBadRequest)
		return
	}

	var c = models.Contact{
		ID: uint(contactID),
	}

	var (
		db           = getDB(r)
		contactStore = models.ContactStore(db)
	)
	err = contactStore.FindNotes(&c)
	if err != nil {
		logs.Error(err)
		Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	Success(w, r, views.Notes{Notes: c.Notes}, http.StatusOK)
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var (
		n = new(models.Note)

		err error
	)
	if err = Request(&views.Note{Note: n}, r); err != nil {
		logs.Debug(err)
		Fail(w, r, map[string]interface{}{"note": err.Error()}, http.StatusBadRequest)
		return
	}

	var (
		db        = getDB(r)
		noteStore = models.NoteStore(db)
	)
	if err = noteStore.Save(n); err != nil {
		logs.Error(err)
		Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/%s/%d", "notes", n.ID))
	Success(w, r, views.Note{Note: n}, http.StatusCreated)
}

func DeleteNote(w http.ResponseWriter, req *http.Request) {
	var (
		contactID int
		err       error
	)
	if contactID, err = strconv.Atoi(router.Context(req).Param("id")); err != nil {
		logs.Debug(err)
		Fail(w, req, map[string]interface{}{"id": "not integer"}, http.StatusBadRequest)
		return
	}

	var noteID int
	if noteID, err = strconv.Atoi(router.Context(req).Param("note_id")); err != nil {
		logs.Debug(err)
		Fail(w, req, map[string]interface{}{"id": "not integer"}, http.StatusBadRequest)
		return
	}

	var (
		db        = getDB(req)
		noteStore = models.NoteStore(db)
		n         = &models.Note{ID: uint(noteID), ContactID: uint(contactID)}
	)
	if err = noteStore.Delete(n); err != nil {
		logs.Debug(err)
		Fail(w, req, map[string]interface{}{"id": "not integer"}, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	Success(w, req, nil, http.StatusNoContent)
}
