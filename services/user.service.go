package services

import (
	"go-buks/database"
	"go-buks/models"

	"gorm.io/gorm"
)

func GetUser(data *models.User, conds ...interface{}) *gorm.DB {
	return database.DB.Model(&models.User{}).Take(data, conds...)
}

func CreateUser(data *models.User) *gorm.DB {
	return database.DB.Create(data)
}
