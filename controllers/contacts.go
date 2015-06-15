package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Quorumsco/contact/models"
	"github.com/Quorumsco/contact/views"
	"github.com/silverwyrda/iogo"
)

func ContactList(w http.ResponseWriter, req *http.Request) {
	db, _ := getEnv(req)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	contacts := models.FindAllContacts(db)

	b, err := json.Marshal(views.ContactsView{Contacts: contacts})
	if err != nil {
		io.WriteString(w, "Error")
	}
	w.Write(b)
}

func ContactByID(w http.ResponseWriter, req *http.Request) {
	db, _ := getEnv(req)
	id, err := strconv.Atoi(iogo.GetContext(req).GetParam("id"))
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	contact := models.FindContactByID(db, id)

	b, err := json.Marshal(views.ContactView{Contact: contact})
	if err != nil {
		io.WriteString(w, "Error")
	}
	w.Write(b)
}

func ContactNew(w http.ResponseWriter, req *http.Request) {
	db, _ := getEnv(req)

	decoder := json.NewDecoder(req.Body)
	var c models.Contact
	err := decoder.Decode(&c)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(c)

	err = c.NewRecord(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	b, err := json.Marshal(views.ContactView{Contact: &c})
	if err != nil {
		io.WriteString(w, "Error")
	}
	w.Write(b)
}

func ContactNewOptions(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin,content-type")
}
