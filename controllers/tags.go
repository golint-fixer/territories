package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/quorumsco/contacts/models"
	"github.com/quorumsco/logs"
)

type Tag struct {
	DB *gorm.DB
}

func (t *Tag) RetrieveCollection(args models.TagArgs, reply *models.TagReply) error {
	var (
		err error

		tagStore = models.TagStore(t.DB)
	)

	reply.Tags, err = tagStore.Find(args)
	if err != nil {
		logs.Error(err)
		return err
	}

	return nil
}

func (t *Tag) Create(args models.TagArgs, reply *models.TagReply) error {
	var (
		err error

		tagStore = models.TagStore(t.DB)
	)

	if err = tagStore.Save(args.Tag, args); err != nil {
		logs.Error(err)
		return err
	}

	return nil
}

func (t *Tag) Delete(args models.TagArgs, reply *models.TagReply) error {
	var (
		err error

		tagStore = models.TagStore(t.DB)
	)

	if err = tagStore.Delete(args.Tag, args); err != nil {
		logs.Debug(err)
		return err
	}

	return nil
}
