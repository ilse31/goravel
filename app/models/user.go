package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	Name        string
	Avatar      string
	PhoneNumber string
	Email       string
	Password    string
	DateOfBirth time.Time
	LastLogin   time.Time
	Country     string
	City        string
	Address     string
	Status      string
	orm.SoftDeletes
}

func (u *User) TableName() string {
	return "users"
}

func (r *User) Connection() string {
	return "postgresql"
}

type UserRequest struct {
	Name        string    `json:"name" validate:"required"`
	Avatar      string    `json:"avatar"`
	PhoneNumber string    `json:"phone_number" validate:"required"`
	Email       string    `json:"email" validate:"required|email"`
	Password    string    `json:"password" validate:"required"`
	DateOfBirth time.Time `json:"date_of_birth"`
	LastLogin   time.Time `json:"last_login"`
	Country     string    `json:"country"`
	City        string    `json:"city"`
	Address     string    `json:"address"`
	Status      string    `json:"status"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required|email"`
	Password string `json:"password" validate:"required"`
}

type ResponseUserLogin struct {
	Token string `json:"token"`
}
