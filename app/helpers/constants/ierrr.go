package ierr

import "fmt"

type Error struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

var (
	ErrInternal         = Error{Code: "00000", Message: "we encountered an error while processing your request (internal server error)"}
	ErrResourceNotFound = Error{Code: "00001", Message: "the requested resource was not found"}
	ErrBadRequest       = Error{Code: "00002", Message: "your request is in a bad format"}
	ErrUnauthorized     = Error{Code: "00003", Message: "you are not authorized to perform the requested action"}
	ErrForbidden        = Error{Code: "00004", Message: "you don't have access to this resource"}
)

var (
	ErrPhoneNumberExist = Error{Code: "00005", Message: "phone number already exist"}
	ErrEmailExist       = Error{Code: "00006", Message: "email already exist"}
	ErrNoRowsAffected   = Error{Code: "00006", Message: "no rows affected"}
	ErrInvalidPassword  = Error{Code: "00007", Message: "invalid password"}
	ErrUserAlreadyLogin = Error{Code: "00008", Message: "user already login"}
)
