package repository_impl

import (
	"awesomeProject/domain/entity"
	"awesomeProject/domain/repository"
	"github.com/jinzhu/gorm"
)

type DeployHistoryRepositoryImpl struct {
	db *gorm.DB
}

func NewDeployHistoryRepositoryImpl(db *gorm.DB) repository.DeployHistoryRepository {
	return DeployHistoryRepositoryImpl{
		db: db,
	}
}

func (d DeployHistoryRepositoryImpl) Search(settingType string, target string, status bool, limit int) []entity.DeployHistory {
	var deployHistories []entity.DeployHistory
	d.db.Where(&entity.DeployHistory{
		SettingType: settingType,
		Target: target,
		Status: status,
	}).Limit(limit).Find(&deployHistories)
	return deployHistories
}

func (d DeployHistoryRepositoryImpl) Create(deployHistory entity.DeployHistory) entity.DeployHistory {
	d.db.Create(&deployHistory)
	return deployHistory
}