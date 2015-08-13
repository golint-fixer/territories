package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type MissionSQL struct {
	DB *gorm.DB
}

func (s *MissionSQL) SaveMission(m *Mission) error {
	c := new([]Contact)
	if m.ID == 0 {
		fmt.Println("")
		fmt.Println(m)
		fmt.Println("")
		s.DB.Model(m).Related(c, "Contacts").Save(m)

		return s.DB.Error
	}

	s.DB.Model(m).Related(c, "Contacts").Update(m)

	return s.DB.Error
}

func (s *MissionSQL) DeleteMission(m *Mission) error {
	c := new([]Contact)
	s.DB.Model(m).Related(c, "Contacts").Delete(m)

	return s.DB.Error
}

func (s *MissionSQL) FindMissions(m Mission) ([]Mission, error) {
	var missions []Mission
	c := new([]Contact)
	s.DB.Model(m).Related(c, "Contacts").Find(&missions)

	if s.DB.Error != nil {
		return make([]Mission, 0), nil
	}

	return missions, s.DB.Error
}

func (s *MissionSQL) FindMissionById(m *Mission) error {
	c := new([]Contact)
	s.DB.Model(m).Related(c, "Contacts").Find(m)

	return s.DB.Error
}
