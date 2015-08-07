package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/quorumsco/contacts/models"
	"github.com/quorumsco/contacts/views"
	. "github.com/quorumsco/jsonapi"
	"github.com/quorumsco/logs"
	"github.com/quorumsco/router"
)

func RetrieveTagById(w http.ResponseWriter, r *http.Request) {
	var (
		contactID int
		err       error
	)
	contactID, err = strconv.Atoi(router.Context(r).Param("id"))
	if err != nil {
		logs.Debug(err)
		Fail(w, r, map[string]interface{}{"contact_id": "not integer"}, http.StatusBadRequest)
		return
	}

	var (
		userID       = getUID(r)
		db           = getDB(r)
		contactStore = models.ContactStore(db)
		c            = models.Contact{ID: uint(contactID)}
	)
	if err = contactStore.First(&c, userID); err != nil {
		if err == sql.ErrNoRows {
			Fail(w, r, nil, http.StatusNotFound)
			return
		}
		logs.Error(err)
		Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	var tagID int
	tagID, err = strconv.Atoi(router.Context(r).Param("tag_id"))
	if err != nil {
		logs.Debug(err)
		Fail(w, r, map[string]interface{}{"tag_id": "not integer"}, http.StatusBadRequest)
		return
	}

	var (
		tagStore = models.TagStore(db)
		t        = &models.Tag{ID: uint(tagID)}
	)
	err = tagStore.FindTagById(t, c)
	if err != nil {
		logs.Error(err)
		Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	Success(w, r, views.Tag{Tag: t}, http.StatusOK)
}

func RetrieveTagCollection(w http.ResponseWriter, r *http.Request) {
	var (
		contactID int
		err       error
	)
	contactID, err = strconv.Atoi(router.Context(r).Param("id"))
	if err != nil {
		logs.Debug(err)
		Fail(w, r, map[string]interface{}{"contact_id": "not integer"}, http.StatusBadRequest)
		return
	}

	var (
		userID       = getUID(r)
		db           = getDB(r)
		contactStore = models.ContactStore(db)
		c            = models.Contact{ID: uint(contactID)}
	)
	if err = contactStore.First(&c, userID); err != nil {
		if err == sql.ErrNoRows {
			Fail(w, r, nil, http.StatusNotFound)
			return
		}
		logs.Error(err)
		Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	err = contactStore.FindTags(&c)
	if err != nil {
		logs.Error(err)
		Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	Success(w, r, views.Tags{Tags: c.Tags}, http.StatusOK)
}

func CreateTag(w http.ResponseWriter, r *http.Request) {
	var (
		contactID int
		err       error

		t = new(models.Tag)
	)
	contactID, err = strconv.Atoi(router.Context(r).Param("id"))
	if err != nil {
		logs.Debug(err)
		Fail(w, r, map[string]interface{}{"contact_id": "not integer"}, http.StatusBadRequest)
		return
	}

	if err = Request(&views.Tag{Tag: t}, r); err != nil {
		logs.Debug(err)
		Fail(w, r, map[string]interface{}{"note": err.Error()}, http.StatusBadRequest)
		return
	}

	var (
		userID       = getUID(r)
		db           = getDB(r)
		tagStore     = models.TagStore(db)
		contactStore = models.ContactStore(db)
		c            = models.Contact{ID: uint(contactID)}
	)

	if err = contactStore.First(&c, userID); err != nil {
		if err == sql.ErrNoRows {
			Fail(w, r, nil, http.StatusNotFound)
			return
		}
		logs.Error(err)
		Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = tagStore.SaveTag(t, c); err != nil {
		logs.Error(err)
		Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/%s/%d", "tag", t.ID))
	Success(w, r, views.Tag{Tag: t}, http.StatusCreated)
}

func DeleteTag(w http.ResponseWriter, r *http.Request) {
	var (
		tagID int
		err   error
	)
	if tagID, err = strconv.Atoi(router.Context(r).Param("tag_id")); err != nil {
		logs.Debug(err)
		Fail(w, r, map[string]interface{}{"tag_id": "not integer"}, http.StatusBadRequest)
		return
	}

	var contactID int
	contactID, err = strconv.Atoi(router.Context(r).Param("id"))
	if err != nil {
		logs.Debug(err)
		Fail(w, r, map[string]interface{}{"contact_id": "not integer"}, http.StatusBadRequest)
		return
	}

	var (
		db           = getDB(r)
		userID       = getUID(r)
		contactStore = models.ContactStore(db)
		c            = models.Contact{ID: uint(contactID)}
	)
	if err = contactStore.First(&c, userID); err != nil {
		if err == sql.ErrNoRows {
			Fail(w, r, nil, http.StatusNotFound)
			return
		}
		logs.Error(err)
		Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	var (
		t        = &models.Tag{ID: uint(tagID)}
		tagStore = models.TagStore(db)
	)
	if err = tagStore.DeleteTag(t, c); err != nil {
		logs.Debug(err)
		Fail(w, r, nil, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	Success(w, r, nil, http.StatusNoContent)
}
