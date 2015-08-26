package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type MissionSQL struct {
	DB *gorm.DB
}

func (s *MissionSQL) Save(m *Mission, args MissionArgs) error {
	if m == nil {
		return errors.New("save: mission is nil")
	}

	if m.ID == 0 {
		return s.DB.Save(m).Error
	}

	return s.DB.Update(m).Error
}

func (s *MissionSQL) Delete(m *Mission, args MissionArgs) error {
	if m == nil {
		return errors.New("save: mission is nil")
	}

	return s.DB.Delete(m).Error
}

func (s *MissionSQL) Find(args MissionArgs) ([]Mission, error) {
	var missions []Mission

	err := s.DB.Find(&missions).Error

	if err != nil {
		return nil, err
	}

	return missions, nil
}

func (s *MissionSQL) First(args MissionArgs) (*Mission, error) {
	var m Mission

	if err := s.DB.Where(args.Mission).First(&m).Error; err != nil {
		if err == gorm.RecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &m, nil
}
