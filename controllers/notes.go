package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/quorumsco/contacts/models"
	"github.com/quorumsco/logs"
)

type Note struct {
	DB *gorm.DB
}

func (t *Note) RetrieveCollection(args models.NoteArgs, reply *models.NoteReply) error {
	var (
		err error

		NoteStore = models.NoteStore(t.DB)
	)

	reply.Notes, err = NoteStore.Find(args)
	if err != nil {
		logs.Error(err)
		return err
	}

	return nil
}

func (t *Note) Retrieve(args models.NoteArgs, reply *models.NoteReply) error {
	var (
		NoteStore = models.NoteStore(t.DB)
		err       error
	)

	if reply.Note, err = NoteStore.First(args); err != nil {
		logs.Error(err)
		return err
	}

	return nil
}

func (t *Note) Create(args models.NoteArgs, reply *models.NoteReply) error {
	var (
		err error

		NoteStore = models.NoteStore(t.DB)
	)

	if err = NoteStore.Save(args.Note, args); err != nil {
		logs.Error(err)
		return err
	}

	reply.Note = args.Note

	return nil
}

func (t *Note) Delete(args models.NoteArgs, reply *models.NoteReply) error {
	var (
		err error

		NoteStore = models.NoteStore(t.DB)
	)

	if err = NoteStore.Delete(args.Note, args); err != nil {
		logs.Debug(err)
		return err
	}

	return nil
}
