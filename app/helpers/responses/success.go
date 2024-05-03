package response

import (
	"github.com/goravel/framework/contracts/http"
)

// Response struct
type Response struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data"`
}

// Success responses with JSON formatresponseMsg
func Success(c http.Context, code int, data interface{}, msg ...string) http.Response {

	responseMsg := buildResponseMsg("Success", msg...)

	if data == nil {
		data = map[string]interface{}{}
	}

	res := Response{
		Success: true,
		Message: responseMsg,
		Data:    data,
	}
	return c.Response().Success().Json(res)
}

// SuccessOK returns code 200
func SuccessOK(c http.Context, data interface{}, msg ...string) http.Response {
	return Success(c, http.StatusOK, data, msg...)
}

// SuccessCreated returns code 201
func SuccessCreated(c http.Context, data interface{}, msg ...string) http.Response {
	return Success(c, http.StatusCreated, data, msg...)
}
