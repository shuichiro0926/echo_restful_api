package entity

import (
	"time"
)

type DeployHistory struct {
	DeployHistoryID uint `gorm:"primary_key"`
	SettingType     string
	Target          string
	SettingFile     string
	Status          bool
	ProcessLog      string
	InsertTm        time.Time
	UpdateTm        time.Time
}
