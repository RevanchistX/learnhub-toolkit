package services

import (
	"gorm.io/gorm"
	"learnhub-toolkit/database/models"
)

type Rooms struct {
	Service
}

func (r *Rooms) ById(id int) (*gorm.DB, models.Room, error) {
	var model models.Room
	result := r.Db().First(&model, id)
	return result, model, nil
}

func (r *Rooms) All(room interface{}) (*gorm.DB, []*models.Room) {
	var list []*models.Room
	result := r.Db().Where(room).Find(&list)
	return result, list
}

func (r *Rooms) Create(room *models.Room) (*gorm.DB, *models.Room) {
	result := r.Db().Create(&room)
	return result, room
}

func (r *Rooms) Update(room *models.Room) (*gorm.DB, *models.Room) {
	result := r.Db().Save(&room)
	return result, room
}

func (r *Rooms) Delete(id int) (*gorm.DB, models.Room) {
	var model models.Room
	r.Db().First(&model, id)
	result := r.Db().Delete(&model)
	return result, model
}
