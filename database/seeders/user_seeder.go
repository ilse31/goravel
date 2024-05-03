package seeders

import (
	"goravel/app/models"

	"github.com/goravel/framework/facades"
)

type UserSeeder struct {
	
}

// Signature The name and signature of the seeder.
func (s *UserSeeder) Signature() string {
	return "UserSeeder"
}

// Run executes the seeder logic.
func (s *UserSeeder) Run() error {
	user := &models.User{
		Name:   "Goravel",
		Avatar: "https://goravel.com/avatar.png",
	}
	return facades.Orm().Query().Create(&user)
}
