package models

import "github.com/jinzhu/gorm"

type MissionDS interface {
	Save(*Mission, MissionArgs) error
	Delete(*Mission, MissionArgs) error
	First(MissionArgs) (*Mission, error)
	Find(MissionArgs) ([]Mission, error)
}

func MissionStore(db *gorm.DB) MissionDS {
	return &MissionSQL{DB: db}
}
