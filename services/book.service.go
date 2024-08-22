package services

import (
	"go-buks/database"
	"go-buks/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func BrowseBooks(data *[]models.Book, queries models.QueryBook, userId uuid.UUID) *gorm.DB {
	offset := (queries.Page - 1) * queries.Limit
	db := database.DB.Offset(offset).Limit(queries.Limit)
	if queries.Search != "" {
		db = db.Where("title ILIKE ?", "%"+queries.Search+"%")
	}
	return db.Where("user_id = ?", userId).Preload("User").Find(data)
}

func GetBook(data *models.Book, id string, userId uuid.UUID) *gorm.DB {
	return database.DB.First(data, "id = ? and user_id = ?", id, userId)
}

func CreateBook(data *models.Book) *gorm.DB {
	return database.DB.Create(data).First(data)
}

func UpdateBook(data *models.Book, id string, userId uuid.UUID) *gorm.DB {
	return database.DB.Clauses(clause.Returning{}).Where("id = ? and user_id = ?", id, userId).Updates(data)
}

func DeleteBook(data *models.Book, id string, userId uuid.UUID) *gorm.DB {
	return database.DB.Clauses(clause.Returning{}).Delete(data, "id = ? and user_id = ?", id, userId)
}