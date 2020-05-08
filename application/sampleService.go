package application

import (
	"awesomeProject/domain/repository"
	"github.com/jinzhu/gorm"
)

type SampleService struct {
	sampleRepository repository.SampleRepository
}

func NewService(sampleRepository repository.SampleRepository) *SampleService {
	return &SampleService{sampleRepository: sampleRepository}
}

func (sampleService SampleService) GetTodoList() *gorm.DB {
	return sampleService.sampleRepository.FindAll()
}
