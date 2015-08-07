package models

import "github.com/jinzhu/gorm"

type TagSQL struct {
	DB *gorm.DB
}

func (s *TagSQL) SaveTag(t *Tag, c Contact) error {
	if t.ID == 0 {
		s.DB.Model(&c).Association("Tags").Append(t)

		return s.DB.Error
	}

	s.DB.Model(&c).Association("Tags").Replace(t)

	return s.DB.Error
}

func (s *TagSQL) DeleteTag(t *Tag, c Contact) error {
	s.DB.Model(&c).Association("Tags").Delete(t)

	return s.DB.Error
}

func (s *TagSQL) FindTagsByContact(c Contact) ([]Tag, error) {
	var tags []Tag
	s.DB.Model(&c).Association("Tags").Find(&tags)
	if s.DB.Error != nil {
		return make([]Tag, 0), nil
	}

	return tags, s.DB.Error
}

func (s *TagSQL) FindTagById(t *Tag, c Contact) error {
	s.DB.Model(&c).Association("Tags").Find(t)

	return s.DB.Error
}
