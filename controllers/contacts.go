package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Quorumsco/contact/models"
	"github.com/Quorumsco/contact/views"
	. "github.com/iogo-framework/jsonapi"
	"github.com/iogo-framework/router"

	"github.com/iogo-framework/logs"
)

func RetrieveContactCollection(w http.ResponseWriter, req *http.Request) {
	db := getDB(req)
	store := models.ContactStore(db)

	contacts, err := store.Find()
	if err != nil {
		logs.Error(err)
		Error(w, req, err.Error(), http.StatusInternalServerError)
		return
	}

	Success(w, req, views.Contacts{Contacts: contacts}, http.StatusOK)
}

func RetrieveContactByID(w http.ResponseWriter, req *http.Request) {
	db := getDB(req)
	store := models.ContactStore(db)

	id, err := strconv.Atoi(router.GetContext(req).GetParam("id"))
	if err != nil {
		logs.Debug(err)
		Fail(w, req, map[string]interface{}{"id": "not integer"}, http.StatusBadRequest)
		return
	}

	var c = models.Contact{
		ID: int64(id),
	}
	err = store.First(&c)
	fmt.Println(&c)
	if err != nil {
		if err == sql.ErrNoRows {
			Success(w, req, nil, http.StatusNotFound)
			return
		}
		logs.Error(err)
		Error(w, req, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	Success(w, req, views.Contact{Contact: &c}, http.StatusOK)
}

func UpdateContactByID(w http.ResponseWriter, req *http.Request) {
	db := getDB(req)
	store := models.ContactStore(db)

	id, err := strconv.Atoi(router.GetContext(req).GetParam("id"))
	if err != nil {
		logs.Debug(err)
		Fail(w, req, map[string]interface{}{"id": "not integer"}, http.StatusBadRequest)
		return
	}

	var c = new(models.Contact)
	c.ID = int64(id)
	err = store.First(c)
	if err != nil {
		Fail(w, req, map[string]interface{}{"contact": err.Error()}, http.StatusBadRequest)
		return
	}

	err = Request(&views.Contact{Contact: c}, req)
	if err != nil {
		logs.Debug(err)
		Fail(w, req, map[string]interface{}{"contact": err.Error()}, http.StatusBadRequest)
		return
	}

	c.ID = int64(id)

	errs := c.Validate()
	if len(errs) > 0 {
		logs.Debug(errs)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		Fail(w, req, map[string]interface{}{"contact": errs}, http.StatusBadRequest)
		return
	}

	err = store.Save(c)
	if err != nil {
		logs.Error(err)
		Error(w, req, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	Success(w, req, views.Contact{Contact: c}, http.StatusOK)
}

func CreateContact(w http.ResponseWriter, req *http.Request) {
	db := getDB(req)
	store := models.ContactStore(db)

	var c = new(models.Contact)
	err := Request(&views.Contact{Contact: c}, req)
	if err != nil {
		logs.Debug(err)
		Fail(w, req, map[string]interface{}{"contact": err.Error()}, http.StatusBadRequest)
		return
	}

	errs := c.Validate()
	if len(errs) > 0 {
		logs.Debug(errs)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		Fail(w, req, map[string]interface{}{"contact": errs}, http.StatusBadRequest)
		return
	}

	err = store.Save(c)
	if err != nil {
		logs.Error(err)
		Error(w, req, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/v1.0/%s/%d", "contacts", c.ID))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	Success(w, req, views.Contact{Contact: c}, http.StatusCreated)
}

func CreateContactOptions(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin,content-type")
}

func DeleteContactByID(w http.ResponseWriter, req *http.Request) {
	db := getDB(req)
	store := models.ContactStore(db)

	id, err := strconv.Atoi(router.GetContext(req).GetParam("id"))
	if err != nil {
		logs.Debug(err)
		Fail(w, req, map[string]interface{}{"id": "not integer"}, http.StatusBadRequest)
		return
	}

	var c = new(models.Contact)
	c.ID = int64(id)
	err = store.Delete(c)
	if err != nil {
		logs.Debug(err)
		Fail(w, req, map[string]interface{}{"id": "not integer"}, http.StatusBadRequest)
		return
	}

	Success(w, req, nil, http.StatusNoContent)
}
