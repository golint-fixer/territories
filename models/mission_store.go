package models

import "github.com/jinzhu/gorm"

type MissionDS interface {
	SaveMission(*Mission) error
	DeleteMission(*Mission) error
	FindMissionById(*Mission) error
	FindMissions(Mission) ([]Mission, error)
}

func MissionStore(db *gorm.DB) MissionDS {
	return &MissionSQL{DB: db}
}
