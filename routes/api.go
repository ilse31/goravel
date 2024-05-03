package routes

import (
	db "goravel/app/repository/DB"

	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
	"net/http"
)

func Api() {
	userController := controllers.NewUserController(*db.NewUserRepository(), *db.NewDevicesRepository(), http.Request{})
	facades.Route().Get("/users", userController.Show)
	facades.Route().Get("/users/{id}", userController.ShowById)
	facades.Route().Post("/users", userController.Store)
	facades.Route().Put("/users/{id}", userController.Update)
	facades.Route().Post("/users/login", userController.Login)
}
