package models

type Config struct {
	APP_ENV     	string	`validate:"oneof=development production" errormsg:"APP_ENV must be one of development or production"`
  PORT       		string	`validate:"numeric" errormsg:"PORT must be a number"`
  HOST       		string	`validate:"hostname" errormsg:"HOST must be a valid hostname"`
  DATABASE_URI 	string	`validate:"required,postgresql_uri" errormsg:"DATABASE_URI required and must be a valid PostgreSQL URI"`
	JWT_KEY				string	`validate:"required" errormsg:"JWT_KEY required"`
	JWT_EXP				string	`validate:"required" errormsg:"JWT_EXP required"`
}