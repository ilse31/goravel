package controllers

import (
	_ "goravel/docs"

	"github.com/goravel/framework/contracts/http"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type SwaggerController struct {
	// Dependent services
}

func NewSwaggerController() *SwaggerController {
	return &SwaggerController{
		// Inject services
	}
}

// Index an example for Swagger
//
//	@Summary      Summary
//	@Description  Description
//	@Tags         example
//	@Accept       json
//	@Success      200
//	@Failure      400
//	@Router       /swagger [get]
func (r *SwaggerController) Index(ctx http.Context) http.Response {
	handler := httpSwagger.Handler()
	handler(ctx.Response().Writer(), ctx.Request().Origin())

	return nil
}
