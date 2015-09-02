package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/quorumsco/contacts/models"
	"github.com/quorumsco/logs"
)

type Contact struct {
	DB *gorm.DB
}

func (t *Contact) RetrieveCollection(args models.ContactArgs, reply *models.ContactReply) error {
	var (
		contactStore = models.ContactStore(t.DB)
		err          error
	)

	if reply.Contacts, err = contactStore.Find(args); err != nil {
		logs.Error(err)
		return err
	}

	return nil
}

func (t *Contact) RetrieveCollectionByMission(args models.ContactArgs, reply *models.ContactReply) error {
	var (
		err error

		contactStore = models.ContactStore(t.DB)
		m            = models.Mission{ID: args.MissionID, GroupID: args.Contact.GroupID}
	)

	if reply.Contacts, err = contactStore.FindByMission(&m, args); err != nil {
		logs.Error(err)
		return err
	}

	return nil
}

func (t *Contact) Retrieve(args models.ContactArgs, reply *models.ContactReply) error {
	var (
		contactStore = models.ContactStore(t.DB)
		err          error
	)

	if reply.Contact, err = contactStore.First(args); err != nil {
		logs.Error(err)
		return err
	}

	return nil
}

func (t *Contact) Update(args models.ContactArgs, reply *models.ContactReply) error {
	var (
		contactStore = models.ContactStore(t.DB)
		err          error
	)

	if reply.Contact, err = contactStore.First(args); err != nil {
		return err
	}

	if err = contactStore.Save(args.Contact, args); err != nil {
		logs.Error(err)
		return err
	}

	return nil
}

func (t *Contact) Create(args models.ContactArgs, reply *models.ContactReply) error {
	var (
		contactStore = models.ContactStore(t.DB)
		err          error
	)

	if err = contactStore.Save(args.Contact, args); err != nil {
		logs.Error(err)
		return err
	}

	reply.Contact = args.Contact

	return nil
}

func (t *Contact) Delete(args models.ContactArgs, reply *models.ContactReply) error {
	var (
		contactStore = models.ContactStore(t.DB)
		err          error
	)

	if err = contactStore.Delete(args.Contact, args); err != nil {
		logs.Debug(err)
		return err
	}

	return nil
}
