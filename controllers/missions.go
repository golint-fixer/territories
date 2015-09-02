package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/quorumsco/contacts/models"
	"github.com/quorumsco/logs"
)

type Mission struct {
	DB *gorm.DB
}

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
