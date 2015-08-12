package models

import "github.com/jinzhu/gorm"

type MissionSQL struct {
	DB *gorm.DB
}

func (s *MissionSQL) SaveMission(m *Mission) error {
	if m.ID == 0 {
		s.DB.Model(m).Association("Contacts").Append(m)

		return s.DB.Error
	}

	s.DB.Model(m).Association("Contacts").Replace(m)

	return s.DB.Error
}

func (s *MissionSQL) DeleteMission(m *Mission) error {
	s.DB.Model(m).Association("Contacts").Delete(m)

	return s.DB.Error
}

func (s *MissionSQL) FindMissions(m Mission) ([]Mission, error) {
	var missions []Mission
	s.DB.Model(&m).Association("Contacts").Find(&missions)
	if s.DB.Error != nil {
		return make([]Mission, 0), nil
	}

	return missions, s.DB.Error
}

func (s *MissionSQL) FindMissionById(m *Mission) error {
	s.DB.Model(m).Association("Contacts").Find(m)

	return s.DB.Error
}
