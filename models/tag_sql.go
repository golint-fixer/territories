package models

import "github.com/jinzhu/gorm"

type TagSQL struct {
	DB *gorm.DB
}

func (s *TagSQL) Save(t *Tag, args TagArgs) error {
	var c = &Contact{ID: args.ContactID}

	if t.ID == 0 {
		s.DB.Debug().Model(c).Association("Tags").Append(t)

		return s.DB.Error
	}

	s.DB.Model(c).Association("Tags").Replace(t)

	return s.DB.Error
}

func (s *TagSQL) Delete(t *Tag, args TagArgs) error {
	s.DB.Model(&Contact{ID: args.ContactID}).Association("Tags").Delete(t)

	return s.DB.Error
}

func (s *TagSQL) Find(args TagArgs) ([]Tag, error) {
	var (
		tags []Tag
		c    = &Contact{ID: args.ContactID}
	)

	s.DB.Model(c).Association("Tags").Find(&tags)
	if s.DB.Error != nil {
		return nil, s.DB.Error
	}

	return tags, s.DB.Error
}

// func (s *TagSQL) FindTagById(t *Tag, c Contact) error {
// 	s.DB.Model(&c).Association("Tags").Find(t)
//
// 	return s.DB.Error
// }
