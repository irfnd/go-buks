package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID				uuid.UUID				`json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey" example:"021534ff-7ae1-4ff0-a109-9198c07c4aee"`
	Email			string					`json:"email" gorm:"uniqueIndex;not null" example:"user1@mail.com"`
	Fullname	string					`json:"fullname" gorm:"not null" example:"User One"`
	Password	string					`json:"-" gorm:"not null" example:"hashed-password"`
	Books			[]Book					`json:"-" gorm:"foreignKey:UserID"`
	CreatedAt	time.Time				`json:"createdAt" gorm:"autoCreateTime" example:"2024-08-20T01:39:03.304451Z"`
  UpdatedAt	time.Time				`json:"updatedAt" gorm:"autoUpdateTime" example:"2024-08-20T01:39:03.304451Z"`
  DeletedAt	gorm.DeletedAt	`json:"-" gorm:"index"`
}

type Register struct {
	Fullname	string	`json:"fullname" validate:"required" errormsg:"Fullname required" example:"User One"`
	Email			string	`json:"email" validate:"required,email" errormsg:"Email required and must be a valid email" example:"user1@mail.com"`
	Password	string	`json:"password" validate:"required,strong_password" errormsg:"Password required and must be a strong password" example:"hashed-password"`
}
type RegisterResponse struct {
	Statuscode	int			`json:"statusCode" default:"201"`
	Message			string	`json:"message" default:"Registered Successfully"`
	Data				*User		`json:"data"`
}
type RegisterResponseBadRequest struct {
	Statuscode		int						`json:"statusCode" default:"400"`
	Message				string				`json:"message" default:"Validate Body Failed"`
	Reason				[]struct {
		Field				string				`json:"field" default:"password"`
		Message			string				`json:"message" default:"Password required and must be a strong password"`
		Value				string				`json:"value" default:"password123."`
	}														`json:"reason"`
}

type Login struct {
	Email			string	`json:"email" validate:"required" errormsg:"Email required" example:"user1@mail.com"`
	Password	string	`json:"password" validate:"required" errormsg:"Password required" example:"Password123."`
}
type LoginResponse struct {
	Statuscode	int				`json:"statusCode" default:"200"`
	Message			string		`json:"message" default:"Logged in Successfully"`
	Data				struct {
		Token			string		`json:"token"`
	}											`json:"data"`
}
type LoginResponseBadRequest struct {
	Statuscode		int						`json:"statusCode" default:"400"`
	Message				string				`json:"message" default:"Validate Body Failed"`
	Reason				[]struct {
		Field				string				`json:"field" default:"password"`
		Message			string				`json:"message" default:"Password required"`
		Value				string				`json:"value" default:""`
	}														`json:"reason"`
}