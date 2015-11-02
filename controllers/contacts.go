package controllers

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/quorumsco/contacts/models"
	"github.com/quorumsco/logs"

	"gopkg.in/olivere/elastic.v2"
)

type Contact struct {
	DB *gorm.DB
}

func (t *Contact) Search(args models.ContactArgs, reply *models.ContactReply) error {
	client, err := elastic.NewClient()
	if err != nil {
		logs.Critical(err)
		return err
	}
	termQuery := elastic.NewMultiMatchQuery(args.Search.Query, args.Search.Field, "firstname")
	termQuery = termQuery.Type("cross_fields")
	termQuery = termQuery.Operator("and")
	searchResult, err := client.Search().
		Index("contacts").
		Query(&termQuery).
		Sort("surname", true).
		Pretty(true).
		Do()
	if err != nil {
		logs.Critical(err)
		return err
	}

	if searchResult.Hits != nil {
		for _, hit := range searchResult.Hits.Hits {
			var c models.Contact
			err := json.Unmarshal(*hit.Source, &c)
			if err != nil {
				logs.Error(err)
				return err
			}
			reply.Contacts = append(reply.Contacts, c)
		}
	} else {
		reply.Contacts = nil
	}

	return nil
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

	if err = contactStore.Save(args.Contact, args); err != nil {
		logs.Error(err)
		return err
	}

	if reply.Contact, err = contactStore.First(args); err != nil {
		return err
	}

	id := strconv.Itoa(int(args.Contact.ID))
	if id == "" {
		logs.Error("id is nil")
		return errors.New("id is nil")
	}

	client, err := elastic.NewClient()
	if err != nil {
		logs.Critical(err)
	} else {
		_, err = client.Index().
			Index("contacts").
			Type("contact").
			Id(id).
			BodyJson(reply.Contact).
			Do()
		if err != nil {
			logs.Critical(err)
			return err
		}
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

	id := strconv.Itoa(int(reply.Contact.ID))
	if id == "" {
		logs.Error("id is nil")
		return errors.New("id is nil")
	}

	client, err := elastic.NewClient()
	if err != nil {
		logs.Critical(err)
	} else {
		_, err = client.Index().
			Index("contacts").
			Type("contact").
			Id(id).
			BodyJson(reply.Contact).
			Do()
		if err != nil {
			logs.Critical(err)
			return err
		}
		_, err = client.Flush().Index("contacts").Do()
		if err != nil {
			logs.Critical(err)
			return err
		}
	}

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

	id := strconv.Itoa(int(args.Contact.ID))
	if id == "" {
		logs.Error("id is nil")
		return errors.New("id is nil")
	}

	client, err := elastic.NewClient()
	if err != nil {
		logs.Critical(err)
	} else {
		_, err = client.Delete().
			Index("contacts").
			Type("contact").
			Id(id).
			Do()
		if err != nil {
			logs.Critical(err)
			return err
		}
	}

	return nil
}
