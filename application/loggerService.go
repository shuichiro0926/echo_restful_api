package application

import (
	"awesomeProject/domain/entity"
	"awesomeProject/domain/repository"
)

type LoggerService struct {
	deployHistoryRepository repository.DeployHistoryRepository
}

func NewLoggerService(deployHistoryRepository repository.DeployHistoryRepository) LoggerService {
	return LoggerService{
		deployHistoryRepository: deployHistoryRepository,
	}
}

func (l LoggerService) SearchDeployHistory(settingType string, target string, status bool, limit int) []entity.DeployHistory {
	return l.deployHistoryRepository.Search(settingType, target, status, limit)
}