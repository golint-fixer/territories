package models

import "github.com/jinzhu/gorm"

type MissionSQL struct {
	DB *gorm.DB
}

func (s *MissionSQL) FindContactByMission(m *Mission) ([]Contact, error) {
	var contacts []Contact
	s.DB.Model(m).Related(new([]Contact), "Contacts").Find(&contacts)

	if s.DB.Error != nil {
		return make([]Contact, 0), nil
	}

	return contacts, s.DB.Error
}

func (s *MissionSQL) SaveMission(m *Mission) error {
	if m.ID == 0 {
		s.DB.Model(m).Related(new([]Contact), "Contacts").Save(m)

		return s.DB.Error
	}

	s.DB.Model(m).Related(new([]Contact), "Contacts").Update(m)

	return s.DB.Error
}

func (s *MissionSQL) DeleteMission(m *Mission) error {
	s.DB.Model(m).Related(new([]Contact), "Contacts").Delete(m)

	return s.DB.Error
}

func (s *MissionSQL) FindMissions(m Mission) ([]Mission, error) {
	var missions []Mission
	s.DB.Model(m).Related(new([]Contact), "Contacts").Find(&missions)

	if s.DB.Error != nil {
		return make([]Mission, 0), nil
	}

	return missions, s.DB.Error
}

func (s *MissionSQL) FindMissionById(m *Mission) error {
	s.DB.Model(m).Related(new([]Contact), "Contacts").Find(m)

	return s.DB.Error
}
