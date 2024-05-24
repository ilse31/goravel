package controllers

import (
	"errors"
	ierr "goravel/app/helpers/constants"
	"goravel/app/helpers/pkg"
	"goravel/app/helpers/request"
	response "goravel/app/helpers/responses"
	"goravel/app/models"
	db "goravel/app/repository/DB"
	"time"

	"github.com/gookit/validate"
	"github.com/goravel/framework/auth"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type UserController struct {
	repoUser    db.UserRepository
	repoDevices db.DevicesRepository
}

func NewUserController(repoUser db.UserRepository, repoDevices db.DevicesRepository) *UserController {
	return &UserController{repoUser, repoDevices}
}

// Show user
// @Summary Show all users
// @Description get users
// @Tags users
// @Accept  json
// @Produce  json
// @Param page query int false "Page"
// @Param per_page query int false "Per Page"
// @Success 200 {object} response.Response{data=[]models.UserRequest} "Success"
// @Router /users [get]
// @Failure 400 {object} response.ErrorResponse{}
func (r *UserController) Show(ctx http.Context) http.Response {

	pages := pkg.NewFromRequest(ctx)

	req := request.Query{
		Limit:  pages.Limit(),
		Offset: pages.Offset(),
	}

	users, total, err := r.repoUser.ListUser(ctx, req.Limit, req.Offset)
	if err != nil {
		return response.ErrInternalServerError(err, ctx)
	}

	pages.SetData(users, int(total))

	return response.SuccessOK(ctx, pages)
}

// Register user
// @Summary Register user
// @Description register user
// @Tags users
// @Accept  json
// @Produce  json
// @Param json body models.UserRequest true "User data"
// @Success 201 {object} response.Response{data=models.UserRequest} "Success"
// @Router /users [post]
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
func (r *UserController) Store(ctx http.Context) http.Response {
	v := validate.New(ctx.Request().All())
	v.StringRule("name", "required")
	v.StringRule("avatar", "required")
	v.StringRule("email", "required|email")
	v.StringRule("phone_number", "required")
	if !v.Validate() {
		return response.ErrBadRequest(errors.New(v.Errors.One()), ctx)
	}

	dateOfBirthStr := ctx.Request().Input("date_of_birth")
	dateOfBirth, err := time.Parse("2006-01-02", dateOfBirthStr)
	if err != nil {
		return response.ErrInternalServerError(err, ctx)
	}

	password, err := facades.Hash().Make(ctx.Request().Input("password"))
	if err != nil {
		return response.ErrInternalServerError(err, ctx)
	}

	user := models.User{}
	if err := ctx.Request().Bind(&user); err != nil {
		return response.ErrInternalServerError(err, ctx)
	}

	isExisEmail, isExistPhone, err := r.repoUser.IsExistEmailandPhoneNumber(ctx, user.PhoneNumber, user.Email)

	if err != nil {
		return response.ErrInternalServerError(err, ctx)
	}

	if isExisEmail {
		return response.ErrBadRequest(ierr.ErrEmailExist, ctx)
	}

	if isExistPhone {
		return response.ErrBadRequest(ierr.ErrPhoneNumberExist, ctx)
	}

	user.DateOfBirth = dateOfBirth
	user.Password = password

	err = r.repoUser.Store(ctx, user)

	if err != nil {
		return response.ErrInternalServerError(err, ctx)
	}

	return response.SuccessCreated(ctx, &models.User{
		Name:        user.Name,
		Avatar:      user.Avatar,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Password:    user.Password,
		Country:     user.Country,
		DateOfBirth: user.DateOfBirth,
		City:        user.City,
		Address:     user.Address,
		Status:      user.Status,
		LastLogin:   user.LastLogin,
	})
}

// Update user
// @Summary Update user
// @Description update user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param json body models.UserRequest true "User data"
// @Success 200 {object} response.Response{data=models.UserRequest} "Success"
// @Router /users/{id} [put]
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
func (r *UserController) Update(ctx http.Context) http.Response {
	v := validate.New(ctx.Request().All())
	if !v.Validate() {
		return response.ErrBadRequest(errors.New(v.Errors.One()), ctx)
	}
	id := ctx.Request().Route("id")
	user := models.User{
		Name:        ctx.Request().Input("name"),
		Avatar:      ctx.Request().Input("avatar"),
		PhoneNumber: ctx.Request().Input("phone_number"),
		Email:       ctx.Request().Input("email"),
		City:        ctx.Request().Input("city"),
		Password:    ctx.Request().Input("password"),
		Address:     ctx.Request().Input("address"),
		Country:     ctx.Request().Input("country"),
	}
	//update user
	_, err := r.repoUser.Update(ctx, user, id)
	if err != nil {
		return response.ErrInternalServerError(err, ctx)
	}
	return nil
}

// Show user by id
// @Summary Show user by id
// @Description get user by id
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} response.Response{data=models.UserRequest} "Success"
// @Router /users/{id} [get]
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
func (r *UserController) ShowById(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	user, err := r.repoUser.GetByID(ctx, id)
	if err != nil {
		return response.ErrInternalServerError(err, ctx)
	}
	return response.SuccessOK(ctx, user)
}

// Login user
// @Summary Login user
// @Description login user
// @Tags users
// @Accept  json
// @Produce  json
// @Param json body models.UserLogin true "User data"
// @Success 200 {object} response.Response{data=models.ResponseUserLogin} "Success"
// @Router /users/login [post]
// @Failure 400 {object} response.ErrorResponse{}
// @Failure 500 {object} response.ErrorResponse{}
func (r *UserController) Login(ctx http.Context) http.Response {
	v := validate.New(ctx.Request().All())
	v.StringRule("email", "required|email")
	v.StringRule("password", "required")
	v.StringRule("device_id", "required")
	v.StringRule("device_ip", "required")
	if !v.Validate() {
		return response.ErrBadRequest(errors.New(v.Errors.One()), ctx)
	}

	userLogin := models.UserLogin{}
	if err := ctx.Request().Bind(&userLogin); err != nil {
		return response.ErrInternalServerError(err, ctx)
	}

	user, err := r.repoUser.GetByEmail(ctx, userLogin.Email)
	if err != nil {
		return response.ErrInternalServerError(err, ctx)
	}
	if !facades.Hash().Check(userLogin.Password, user.Password) {
		return response.ErrBadRequest(ierr.ErrInvalidPassword, ctx)
	}

	token, err := facades.Auth().Login(ctx, &user)
	if err != nil {
		return response.ErrInternalServerError(err, ctx)
	}

	//get token by user id
	tokenByUserID, err := r.repoDevices.GetTokenByUserID(ctx, user.ID)

	if err != nil {
		return response.ErrInternalServerError(err, ctx)
	}

	//if exist token and device id. return error. dont allow user login with same device
	if tokenByUserID.Token != "" && tokenByUserID.DeviceID == ctx.Request().Header("DeviceID") {
		payload, err := facades.Auth().Parse(ctx, tokenByUserID.Token)
		if err != nil {
			if errors.Is(err, auth.ErrorTokenExpired) {
				facades.Log().Info("Token expired at", payload.ExpireAt)
			}
		}

		if payload.ExpireAt.After(time.Now()) {
			// Token masih valid, kirimkan pesan error
			facades.Log().Info("Token masih valid", payload.ExpireAt.Unix(), time.Now().Unix())
			return response.ErrBadRequest(ierr.ErrUserAlreadyLogin, ctx)
		}
	}

	// Jika token tidak ada atau telah kedaluwarsa, lanjutkan dengan update token

	result, err := r.repoDevices.UpdateDevice(ctx, tokenByUserID, user, token, ctx.Request().Header("DeviceID"), ctx.Request().Header("DeviceIP"))

	if err != nil {
		return response.ErrInternalServerError(err, ctx)
	}

	if result == 0 {
		// Jika tidak ada perangkat yang diupdate, maka insert data baru
		devices := models.Devices{
			UserID:   user.ID,
			Token:    token,
			DeviceID: ctx.Request().Header("DeviceID"),
			DeviceIP: ctx.Request().Header("DeviceIP"),
		}

		_, err := r.repoDevices.CreateDevice(ctx, devices)
		if err != nil {
			return response.ErrInternalServerError(errors.New("failed to create device"), ctx)
		}

	}

	loginResponse := models.ResponseUserLogin{
		Token: token,
	}

	return response.SuccessOK(ctx, loginResponse)
}
