package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/quorumsco/contacts/models"
	"github.com/quorumsco/logs"
)

// Tags contains the tag related methods and a gorm client
type Tag struct {
	DB *gorm.DB
}

// RetrieveCollection calls the TagSQL Find method and returns the results via RPC
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

// Create calls the TagSQL Save method and returns the results via RPC
func (t *Tag) Create(args models.TagArgs, reply *models.TagReply) error {
	var (
		err error

		tagStore = models.TagStore(t.DB)
	)

	if err = tagStore.Save(args.Tag, args); err != nil {
		logs.Error(err)
		return err
	}

	reply.Tag = args.Tag

	return nil
}

// Delete calls the TagSQL Delete method and returns the results via RPC
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
