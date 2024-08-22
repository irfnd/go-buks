package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID					uuid.UUID				`json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey" example:"021534ff-7ae1-4ff0-a109-9198c07c4aee"`
	Title				string					`json:"title" gorm:"not null" example:"The Go Programming Language"`
	Author			string					`json:"author" gorm:"not null" example:"Alan A. A. Donovan, Brian W. Kernighan"`
	Publisher		string					`json:"publisher" gorm:"default:-" example:"Addison-Wesley Professional"`
	Year				int							`json:"year" gorm:"not null" example:"2015"`
	UserID   		uuid.UUID				`json:"-" gorm:"not null"`
	CreatedAt		time.Time				`json:"createdAt" gorm:"autoCreateTime" example:"2024-08-20T01:39:03.304451Z"`
  UpdatedAt		time.Time				`json:"updatedAt" gorm:"autoUpdateTime" example:"2024-08-20T01:39:03.304451Z"`
  DeletedAt		gorm.DeletedAt	`json:"-" gorm:"index"`
	User   			User      			`json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ParamsBook struct {
	ID					string					`json:"id" params:"id" validate:"uuid4" errormsg:"ID is not valid UUID"`
}

type QueryBook struct {
	Search			string					`json:"search"`
	Page				int							`json:"page" validate:"gt=0" errormsg:"Page must be greater than 0" default:"1"`
	Limit				int							`json:"limit" validate:"gt=0" errormsg:"Limit must be greater than 0" default:"10"`
}
type QueryBookResponse struct {
	Statuscode		int						`json:"statusCode" default:"200"`
	Message				string				`json:"message" default:"Retrieved Books Successfully"`
	Data					struct {
		Books 			*[]Book				`json:"books"`
		Pagination	struct {
			Page			int						`json:"page" default:"1"`
			Limit			int						`json:"limit" default:"10"`
		}													`json:"pagination"`
	}														`json:"data"`
}
type QueryBookResponseBadRequest struct {
	Statuscode		int						`json:"statusCode" default:"400"`
	Message				string				`json:"message" default:"Validate Queries Failed"`
	Reason				[]struct {
		Field				string				`json:"field" default:"page"`
		Message			string				`json:"message" default:"Page must be greater than 0"`
		Value				string				`json:"value" default:"-1"`
	}														`json:"reason"`
}

type GetBookResponse struct {
	Statuscode		int						`json:"statusCode" default:"200"`
	Message				string				`json:"message" default:"Retrieved Book Successfully"`
	Data					*Book					`json:"data"`
}
type GetBookResponseBadRequest struct {
	Statuscode		int						`json:"statusCode" default:"400"`
	Message				string				`json:"message" default:"Validate Params Failed"`
	Reason				[]struct {
		Field				string				`json:"field" default:"id"`
		Message			string				`json:"message" default:"ID is not valid UUID"`
		Value				string				`json:"value" default:"adsfasdf"`
	}														`json:"reason"`
}

type CreateBook struct {
	Title				string					`json:"title" validate:"required" errormsg:"Title is required" example:"The Go Programming Language"`
	Author			string					`json:"author" validate:"required" errormsg:"Author is required" example:"Alan A. A. Donovan, Brian W. Kernighan"`
	Publisher		string					`json:"publisher" example:"Addison-Wesley Professional"`
	Year				int							`json:"year" validate:"required,gt=0,lt=2100" errormsg:"Year is required and must be greater than 0 and less than 2100" example:"2015"`
}
type CreateBookResponse struct {
	Statuscode		int						`json:"statusCode" default:"201"`
	Message				string				`json:"message" default:"Created Book Successfully"`
	Data					*Book
}
type CreateBookResponseBadRequest struct {
	Statuscode		int						`json:"statusCode" default:"400"`
	Message				string				`json:"message" default:"Validate Body Failed"`
	Reason				[]struct {
		Field				string				`json:"field" default:"year"`
		Message			string				`json:"message" default:"Year is required and must be greater than 0 and less than 2100"`
		Value				string				`json:"value" default:"2200"`
	}
}

type UpdateBook struct {
	Title				string					`json:"title" example:"The Go Programming Language"`
	Author			string					`json:"author" example:"Alan A. A. Donovan, Brian W. Kernighan"`
	Publisher		string					`json:"publisher" example:"Alan A. A. Donovan, Brian W. Kernighan"`
	Year				int							`json:"year" validate:"gt=0,lt=2100" errormsg:"Year must be greater than 0 and less than 2100" example:"2015"`
}
type UpdateBookResponse struct {
	Statuscode		int						`json:"statusCode" default:"200"`
	Message				string				`json:"message" default:"Updated Book Successfully"`
	Data					*Book					`json:"data"`
}
type UpdateBookResponseBadRequest struct {
	Statuscode		int						`json:"statusCode" default:"400"`
	Message				string				`json:"message" default:"Validate Body Failed"`
	Reason				[]struct {
		Field				string				`json:"field" default:"year"`
		Message			string				`json:"message" default:"Year must be greater than 0 and less than 2100"`
		Value				string				`json:"value" default:"2200"`
	}														`json:"reason"`
}

type DeleteBookResponse struct {
	Statuscode		int						`json:"statusCode" default:"200"`
	Message				string				`json:"message" default:"Deleted Book Successfully"`
	Data					*Book					`json:"data"`
}
type DeleteBookResponseBadRequest struct {
	Statuscode		int						`json:"statusCode" default:"400"`
	Message				string				`json:"message" default:"Validate Params Failed"`
	Reason				[]struct {
		Field				string				`json:"field" default:"id"`
		Message			string				`json:"message" default:"ID is not valid UUID"`
		Value				string				`json:"value" default:"adsfasdf"`
	}														`json:"reason"`
}