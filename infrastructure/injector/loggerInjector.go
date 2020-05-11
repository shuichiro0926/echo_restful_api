package injector

import (
	"awesomeProject/application"
	"awesomeProject/domain/repository"
	"awesomeProject/infrastructure"
	"awesomeProject/interface/controller"
	"awesomeProject/interface/repository_impl"
)

func InjectRepository() repository.DeployHistoryRepository {
	return repository_impl.NewDeployHistoryRepositoryImpl(infrastructure.GetConnection())
}

func InjectService() application.LoggerService {
	return application.NewLoggerService(InjectRepository())
}

func InjectController() controller.LoggerController {
	return controller.NewLoggerController(InjectService())
}
