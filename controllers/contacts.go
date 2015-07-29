package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/quorumsco/contacts/models"
	"github.com/quorumsco/contacts/views"
	. "github.com/quorumsco/jsonapi"
	"github.com/quorumsco/logs"
	"github.com/quorumsco/router"
)

const (
	ErrMissingUserID = errors.New("missing")
)

func RetrieveContactCollection(w http.ResponseWriter, r *http.Request) {
	var (
		query = r.URL.Query()

		err error
	)

	var userID int
	userID, err = strconv.Atoi(query.Get("user_id"))
	if err != nil {
		logs.Debug(err)
		Error(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	var (
		contacts     []models.Contact
		db           = getDB(r)
		contactStore = models.ContactStore(db)
	)
	if contacts, err = contactStore.Find(uint(userID)); err != nil {
		logs.Error(err)
		Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	Success(w, r, views.Contacts{Contacts: contacts}, http.StatusOK)
}

func RetrieveContact(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(router.Context(r).Param("id"))
	if err != nil {
		logs.Debug(err)
		Fail(w, r, map[string]interface{}{"id": "not integer"}, http.StatusBadRequest)
		return
	}

	var (
		c            = models.Contact{ID: uint(id)}
		db           = getDB(r)
		contactStore = models.ContactStore(db)
	)
	if err = contactStore.First(&c); err != nil {
		if err == sql.ErrNoRows {
			Success(w, r, nil, http.StatusNotFound)
			return
		}
		logs.Error(err)
		Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	Success(w, r, views.Contact{Contact: &c}, http.StatusOK)
}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	var (
		contactID int
		err       error
	)
	if contactID, err = strconv.Atoi(router.Context(r).Param("id")); err != nil {
		logs.Debug(err)
		Fail(w, r, map[string]interface{}{"id": "not integer"}, http.StatusBadRequest)
		return
	}

	var (
		db           = getDB(r)
		contactStore = models.ContactStore(db)
		c            = &models.Contact{ID: uint(contactID)}
	)
	if err = contactStore.First(c); err != nil {
		Fail(w, r, map[string]interface{}{"contact": err.Error()}, http.StatusBadRequest)
		return
	}

	if err = Request(&views.Contact{Contact: c}, r); err != nil {
		logs.Debug(err)
		Fail(w, r, map[string]interface{}{"contact": err.Error()}, http.StatusBadRequest)
		return
	}
	c.ID = uint(contactID)

	var errs = c.Validate()
	if len(errs) > 0 {
		logs.Debug(errs)
		Fail(w, r, map[string]interface{}{"contact": errs}, http.StatusBadRequest)
		return
	}

	if err = contactStore.Save(c); err != nil {
		logs.Error(err)
		Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	Success(w, r, views.Contact{Contact: c}, http.StatusOK)
}

func CreateContact(w http.ResponseWriter, r *http.Request) {
	var (
		c = new(models.Contact)

		err error
	)
	if err := Request(&views.Contact{Contact: c}, r); err != nil {
		logs.Debug(err)
		Fail(w, r, map[string]interface{}{"contact": err.Error()}, http.StatusBadRequest)
		return
	}

	errs := c.Validate()
	if len(errs) > 0 {
		logs.Debug(errs)
		Fail(w, r, map[string]interface{}{"contact": errs}, http.StatusBadRequest)
		return
	}

	var (
		db           = getDB(r)
		contactStore = models.ContactStore(db)
	)
	if err = contactStore.Save(c); err != nil {
		logs.Error(err)
		Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/%s/%d", "contacts", c.ID))
	Success(w, r, views.Contact{Contact: c}, http.StatusCreated)
}

func ContactCollectionOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin,access-control-allow-methods,content-type")
}

func ContactOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin,access-control-allow-methods,content-type")
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	var (
		contactID int
		err       error
	)
	if contactID, err = strconv.Atoi(router.Context(r).Param("id")); err != nil {
		logs.Debug(err)
		Fail(w, r, map[string]interface{}{"id": "not integer"}, http.StatusBadRequest)
		return
	}

	var (
		db           = getDB(r)
		contactStore = models.ContactStore(db)
		c            = &models.Contact{ID: uint(contactID)}
	)
	if err = contactStore.Delete(c); err != nil {
		logs.Debug(err)
		Fail(w, r, map[string]interface{}{"id": "not integer"}, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	Success(w, r, nil, http.StatusNoContent)
}
