// Bundle of functions managing the CRUD and the elasticsearch engine
package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/quorumsco/contacts/models"
	"github.com/quorumsco/logs"
)

// Mission contains the mission related methods and a gorm client
type Mission struct {
	DB *gorm.DB
}

// RetrieveCollection calls the MissionSQL Find method and returns the results via RPC
func (t *Mission) RetrieveCollection(args models.MissionArgs, reply *models.MissionReply) error {
	var (
		missionStore = models.MissionStore(t.DB)
		err          error
	)

	if reply.Missions, err = missionStore.Find(args); err != nil {
		logs.Error(err)
		return err
	}

	return nil
}

// Create calls the MissionSQL Save method and returns the results via RPC
func (t *Mission) Create(args models.MissionArgs, reply *models.MissionReply) error {
	var (
		missionStore = models.MissionStore(t.DB)
		err          error
	)

	if err = missionStore.Save(args.Mission, args); err != nil {
		logs.Error(err)
		return err
	}

	reply.Mission = args.Mission

	return nil
}

// Update calls the MissionSQL First method and returns the results via RPC
func (t *Mission) Update(args models.MissionArgs, reply *models.MissionReply) error {
	var (
		missionStore = models.MissionStore(t.DB)
		err          error
	)

	if reply.Mission, err = missionStore.First(args); err != nil {
		return err
	}

	if err = missionStore.Save(args.Mission, args); err != nil {
		logs.Error(err)
		return err
	}

	return nil
}

// Delete calls the MissionSQL Delete method and returns the results via RPC
func (t *Mission) Delete(args models.MissionArgs, reply *models.MissionReply) error {
	var (
		missionStore = models.MissionStore(t.DB)
		err          error
	)

	if err = missionStore.Delete(args.Mission, args); err != nil {
		logs.Debug(err)
		return err
	}

	return nil
}
