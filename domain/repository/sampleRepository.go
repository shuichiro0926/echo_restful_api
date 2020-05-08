package repository

import (
	"github.com/jinzhu/gorm"
)

type SampleRepository interface {
	FindAll() *gorm.DB
}
