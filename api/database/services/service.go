package services

import (
	"gorm.io/gorm"
	"learnhub-toolkit/api/database"
)

type Service struct {
}

func (s *Service) Db() *gorm.DB {
	return database.Db()
}

func (s *Service) Find(i interface{}) {

}
