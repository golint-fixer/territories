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

func RetrieveContactsByMission(w http.ResponseWriter, r *http.Request) {
	var (
		missionID int
		err       error
	)
	missionID, err = strconv.Atoi(router.Context(r).Param("mission_id"))
	if err != nil {
		logs.Debug(err)
		Fail(w, r, map[string]interface{}{"mission_id": "not integer"}, http.StatusBadRequest)
		return
	}

	var (
		groupID      = getGID(r)
		db           = getDB(r)
		missionStore = models.MissionStore(db)
		m            = models.Mission{ID: uint(missionID), GroupID: groupID}
	)
	if err = missionStore.FindMissionById(&m); err != nil {
		if err == sql.ErrNoRows {
			logs.Error(err)
			Fail(w, r, nil, http.StatusNotFound)
			return
		}
		logs.Error(err)
		Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	var c []models.Contact
	if c, err = missionStore.FindContactByMission(&m); err != nil {
		logs.Error(err)
		Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	Success(w, r, views.Contacts{Contacts: c}, http.StatusOK)
}

func RetrieveMissionById(w http.ResponseWriter, r *http.Request) {
	var (
		missionID int
		err       error
	)
	missionID, err = strconv.Atoi(router.Context(r).Param("mission_id"))
	if err != nil {
		logs.Debug(err)
		Fail(w, r, map[string]interface{}{"mission_id": "not integer"}, http.StatusBadRequest)
		return
	}

	var (
		groupID      = getGID(r)
		db           = getDB(r)
		missionStore = models.MissionStore(db)
		m            = models.Mission{ID: uint(missionID), GroupID: groupID}
	)
	if err = missionStore.FindMissionById(&m); err != nil {
		if err == sql.ErrNoRows {
			logs.Error(err)
			Fail(w, r, nil, http.StatusNotFound)
			return
		}
		logs.Error(err)
		Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	Success(w, r, views.Mission{Mission: &m}, http.StatusOK)
}

func RetrieveMissionCollection(w http.ResponseWriter, r *http.Request) {
	var (
		groupID      = getGID(r)
		db           = getDB(r)
		missionStore = models.MissionStore(db)
		m            = models.Mission{GroupID: uint(groupID)}
	)
	missions, err := missionStore.FindMissions(m)
	if err != nil {
		logs.Error(err)
		Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	Success(w, r, views.Missions{Missions: missions}, http.StatusOK)
}

func CreateMission(w http.ResponseWriter, r *http.Request) {
	var (
		err error

		m = new(models.Mission)
	)
	if err = Request(&views.Mission{Mission: m}, r); err != nil {
		logs.Debug(err)
		Fail(w, r, map[string]interface{}{"mission": err.Error()}, http.StatusBadRequest)
		return
	}
	var (
		groupID      = getGID(r)
		db           = getDB(r)
		missionStore = models.MissionStore(db)
	)
	m.GroupID = groupID
	for i, _ := range m.Contacts {
		m.Contacts[i].GroupID = groupID
	}
	if err = missionStore.SaveMission(m); err != nil {
		logs.Error(err)
		Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/%s/%d", "mission", m.ID))
	Success(w, r, views.Mission{Mission: m}, http.StatusCreated)
}

func UpdateMission(w http.ResponseWriter, r *http.Request) {
	var (
		missionID int
		err       error
	)
	if missionID, err = strconv.Atoi(router.Context(r).Param("mission_id")); err != nil {
		logs.Debug(err)
		Fail(w, r, map[string]interface{}{"mission_id": "not integer"}, http.StatusBadRequest)
		return
	}

	var (
		groupID      = getGID(r)
		db           = getDB(r)
		missionStore = models.MissionStore(db)
		m            = &models.Mission{ID: uint(missionID), GroupID: groupID}
	)
	if err = missionStore.FindMissionById(m); err != nil {
		Fail(w, r, map[string]interface{}{"mission": err.Error()}, http.StatusBadRequest)
		return
	}

	if err = Request(&views.Mission{Mission: m}, r); err != nil {
		logs.Debug(err)
		Fail(w, r, map[string]interface{}{"mission": err.Error()}, http.StatusBadRequest)
		return
	}

	if err = missionStore.SaveMission(m); err != nil {
		logs.Error(err)
		Error(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	Success(w, r, views.Mission{Mission: m}, http.StatusOK)
}

func DeleteMission(w http.ResponseWriter, r *http.Request) {
	var (
		missionID int
		err       error
	)
	if missionID, err = strconv.Atoi(router.Context(r).Param("mission_id")); err != nil {
		logs.Debug(err)
		Fail(w, r, map[string]interface{}{"mission_id": "not integer"}, http.StatusBadRequest)
		return
	}

	var (
		db           = getDB(r)
		groupID      = getGID(r)
		missionStore = models.MissionStore(db)
		m            = models.Mission{ID: uint(missionID), GroupID: groupID}
	)
	if err = missionStore.DeleteMission(&m); err != nil {
		logs.Debug(err)
		Fail(w, r, nil, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	Success(w, r, nil, http.StatusNoContent)
}
