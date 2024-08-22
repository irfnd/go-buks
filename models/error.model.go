package models

import (
	"fmt"
)

type CustomError struct {
	Statuscode	int						`json:"statusCode" default:"500"`
	Message			string				`json:"message" default:"Internal Server Error"`
	Reason			interface{}		`json:"reason" default:"Details Error"`
}
type CustomErrorResponse struct {
	Statuscode	int						`json:"statusCode" default:"500"`
	Message			string				`json:"message" default:"Internal Server Error"`
	Reason			string				`json:"reason" default:"Details Error"`
}
func (e *CustomError) Error() string {
	return fmt.Sprintf("Statuscode: %d, Message: %s, Reason: %s", e.Statuscode, e.Message, e.Reason)
}

type ValidationDetail struct {
	Field			string		`json:"field"`
	Value			string		`json:"value"`
	Message		string		`json:"message"`
}

type UnauthorizedResponse struct {
	Statuscode	int						`json:"statusCode" default:"401"`
	Message			string				`json:"message" default:"Unauthorized"`
	Reason			string				`json:"reason" default:"Authorization Error"`
}