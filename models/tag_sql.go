package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type TagSQL struct {
	DB *gorm.DB
}

func (s *TagSQL) Save(t *Tag, args TagArgs) error {
	if t == nil {
		return errors.New("save: tag is nil")
	}

	var c = &Contact{ID: args.ContactID}

	if t.ID == 0 {
		err := s.DB.Debug().Model(c).Association("Tags").Append(t).Error
		s.DB.Last(t)
		return err
	}

	return s.DB.Model(c).Association("Tags").Replace(t).Error
}

func (s *TagSQL) Delete(t *Tag, args TagArgs) error {
	return s.DB.Model(&Contact{ID: args.ContactID}).Association("Tags").Delete(t).Error
}

func (s *TagSQL) Find(args TagArgs) ([]Tag, error) {
	var (
		tags []Tag
		c    = &Contact{ID: args.ContactID}
	)

	err := s.DB.Model(c).Association("Tags").Find(&tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}
