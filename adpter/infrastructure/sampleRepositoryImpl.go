package infrastructure

import (
	"awesomeProject/domain/entity"
	"awesomeProject/domain/repository"
	"github.com/jinzhu/gorm"
)

type SampleRepositoryImpl struct {
	db *gorm.DB
}

func NewRepositoryImpl(db *gorm.DB) repository.SampleRepository  {
	return &SampleRepositoryImpl{db: db}
}

func (r SampleRepositoryImpl) FindAll() *gorm.DB {
	var todoList []entity.Todo
	return r.db.Find(&todoList)
}

