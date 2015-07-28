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

func RetrieveContactCollection(w http.ResponseWriter, req *http.Request) {
	var (
		db           = getDB(req)
		contactStore = models.ContactStore(db)
	)
	contacts, err := contactStore.Find()
	if err != nil {
		logs.Error(err)
		Error(w, req, err.Error(), http.StatusInternalServerError)
		return
	}

	Success(w, req, views.Contacts{Contacts: contacts}, http.StatusOK)
}

func RetrieveContact(w http.ResponseWriter, req *http.Request) {

	id, err := strconv.Atoi(router.Context(req).Param("id"))
	if err != nil {
		logs.Debug(err)
		Fail(w, req, map[string]interface{}{"id": "not integer"}, http.StatusBadRequest)
		return
	}

	var (
		c            = models.Contact{ID: uint(id)}
		db           = getDB(req)
		contactStore = models.ContactStore(db)
	)
	if err = contactStore.First(&c); err != nil {
		if err == sql.ErrNoRows {
			Success(w, req, nil, http.StatusNotFound)
			return
		}
		logs.Error(err)
		Error(w, req, err.Error(), http.StatusInternalServerError)
		return
	}

	Success(w, req, views.Contact{Contact: &c}, http.StatusOK)
}

func UpdateContact(w http.ResponseWriter, req *http.Request) {
	var (
		contactID int
		err       error
	)
	if contactID, err = strconv.Atoi(router.Context(req).Param("id")); err != nil {
		logs.Debug(err)
		Fail(w, req, map[string]interface{}{"id": "not integer"}, http.StatusBadRequest)
		return
	}

	var (
		db           = getDB(req)
		contactStore = models.ContactStore(db)
		c            = &models.Contact{ID: uint(contactID)}
	)
	if err = contactStore.First(c); err != nil {
		Fail(w, req, map[string]interface{}{"contact": err.Error()}, http.StatusBadRequest)
		return
	}

	if err = Request(&views.Contact{Contact: c}, req); err != nil {
		logs.Debug(err)
		Fail(w, req, map[string]interface{}{"contact": err.Error()}, http.StatusBadRequest)
		return
	}
	c.ID = uint(contactID)

	var errs = c.Validate()
	if len(errs) > 0 {
		logs.Debug(errs)
		Fail(w, req, map[string]interface{}{"contact": errs}, http.StatusBadRequest)
		return
	}

	if err = contactStore.Save(c); err != nil {
		logs.Error(err)
		Error(w, req, err.Error(), http.StatusInternalServerError)
		return
	}

	Success(w, req, views.Contact{Contact: c}, http.StatusOK)
}

func CreateContact(w http.ResponseWriter, req *http.Request) {
	var (
		c = new(models.Contact)

		err error
	)
	if err := Request(&views.Contact{Contact: c}, req); err != nil {
		logs.Debug(err)
		Fail(w, req, map[string]interface{}{"contact": err.Error()}, http.StatusBadRequest)
		return
	}

	errs := c.Validate()
	if len(errs) > 0 {
		logs.Debug(errs)
		Fail(w, req, map[string]interface{}{"contact": errs}, http.StatusBadRequest)
		return
	}

	var (
		db           = getDB(req)
		contactStore = models.ContactStore(db)
	)
	if err = contactStore.Save(c); err != nil {
		logs.Error(err)
		Error(w, req, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/%s/%d", "contacts", c.ID))
	Success(w, req, views.Contact{Contact: c}, http.StatusCreated)
}

func ContactCollectionOptions(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin,access-control-allow-methods,content-type")
}

func ContactOptions(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin,access-control-allow-methods,content-type")
}

func DeleteContact(w http.ResponseWriter, req *http.Request) {
	var (
		contactID int
		err       error
	)
	if contactID, err = strconv.Atoi(router.Context(req).Param("id")); err != nil {
		logs.Debug(err)
		Fail(w, req, map[string]interface{}{"id": "not integer"}, http.StatusBadRequest)
		return
	}

	var (
		db           = getDB(req)
		contactStore = models.ContactStore(db)
		c            = &models.Contact{ID: uint(contactID)}
	)
	if err = contactStore.Delete(c); err != nil {
		logs.Debug(err)
		Fail(w, req, map[string]interface{}{"id": "not integer"}, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	Success(w, req, nil, http.StatusNoContent)
}
