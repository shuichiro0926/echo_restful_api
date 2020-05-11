package repository

import (
	"awesomeProject/domain/entity"
)

type DeployHistoryRepository interface {
	Search(settingType string, target string, status bool, limit int) []entity.DeployHistory
	Create(deployHistory entity.DeployHistory) entity.DeployHistory
}