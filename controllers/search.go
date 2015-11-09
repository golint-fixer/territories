// Bundle of functions managing the CRUD and the elasticsearch engine
package controllers

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/quorumsco/contacts/models"
	"github.com/quorumsco/elastic"
	"github.com/quorumsco/logs"
)

// Search contains the search related methods and a gorm client
type Search struct {
	Client *elastic.Client
}

// Index indexes a contact into elasticsearch
func (s *Search) Index(args models.ContactArgs, reply *models.ContactReply) error {
	id := strconv.Itoa(int(args.Contact.ID))
	if id == "" {
		logs.Error("id is nil")
		return errors.New("id is nil")
	}

	_, err := s.Client.Index().
		Index("contacts").
		Type("contact").
		Id(id).
		BodyJson(args.Contact).
		Do()
	if err != nil {
		logs.Critical(err)
		return err
	}

	return nil
}

// UnIndex unindexes a contact from elasticsearch
func (s *Search) UnIndex(args models.ContactArgs, reply *models.ContactReply) error {
	id := strconv.Itoa(int(args.Contact.ID))
	if id == "" {
		logs.Error("id is nil")
		return errors.New("id is nil")
	}

	_, err := s.Client.Delete().
		Index("contacts").
		Type("contact").
		Id(id).
		Do()
	if err != nil {
		logs.Critical(err)
		return err
	}

	return nil
}

// SearchContacts performs a cross_field search request to elasticsearch and returns the results via RPC
func (s *Search) SearchContacts(args models.SearchArgs, reply *models.SearchReply) error {
	termQuery := elastic.NewMultiMatchQuery(args.Search.Query, args.Search.Field, "firstname")
	termQuery = termQuery.Type("cross_fields")
	termQuery = termQuery.Operator("and")
	searchResult, err := s.Client.Search().
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
