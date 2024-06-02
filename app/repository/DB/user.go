package db

import (
	"context"
	"fmt"
	"goravel/app/models"

	// "github.com/goravel/framework/database/gorm"
	"github.com/goravel/framework/facades"
	"github.com/pkg/errors"

	"gorm.io/gorm"
)

type UserRepository struct {
}

// NewUserRepository creates a new user repository
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// GetByID returns the user with the specified user ID.
func (r *UserRepository) GetByPhoneNumber(ctx context.Context, PhoneNumber string) (models.User, error) {

	user := models.User{PhoneNumber: PhoneNumber}
	q := facades.Orm().Query().Model(&user)
	err := q.Where("phone_number = ?", PhoneNumber).First(&user)
	if err != nil {
		return models.User{}, errors.Wrap(err, "failed to get user by phone number")
	}
	return user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, Email string) (models.User, error) {

	user := models.User{Email: Email}
	q := facades.Orm().Query().Model(&user)
	err := q.Where("email = ?", Email).First(&user)
	if err != nil {
		return models.User{}, errors.Wrap(err, "failed to get user by phone number")
	}
	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, user models.User, id string) (int, error) {

	result, err := facades.Orm().Query().Where("id = ?", id).Update(&user)
	if err != nil {
		return 0, errors.Wrap(err, "failed to update user")
	}
	return int(result.RowsAffected), nil
}

func (r *UserRepository) ListUser(ctx context.Context, limit, offset int) ([]models.User, int, error) {
	users := make([]models.User, 0)
	q := facades.Orm().Query().Model(users)

	var total int64
	if err := q.Count(&total); err != nil {
		return nil, 0, errors.Wrap(err, "failed to count users")
	}

	err := q.Limit(limit).Offset(offset).Find(&users)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to get users")
	}

	return users, int(total), nil
}

func (r *UserRepository) GetByID(ctx context.Context, ID string) (models.User, error) {

	user := models.User{}
	q := facades.Orm().Query().Model(&user)
	err := q.Where("id = ?", ID).First(&user)
	if err != nil {
		return models.User{}, errors.Wrap(err, "failed to get user by phone number")
	}
	return user, nil
}

func (r *UserRepository) IsExistEmailandPhoneNumber(ctx context.Context, PhoneNumber string, Email string) (bool, bool, error) {
	var userExist models.UserExist

	// Check if the email exists
	err := facades.Orm().Query().Raw("SELECT EXISTS (SELECT 1 FROM users WHERE email = ?) AS is_exist_email", Email).Scan(&userExist)
	if err != nil {
		return false, false, errors.Wrap(err, "failed to check email")
	}

	// Check if the phone number exists
	err = facades.Orm().Query().Raw("SELECT EXISTS (SELECT 1 FROM users WHERE phone_number = ?) AS is_exist_phone", PhoneNumber).Scan(&userExist)
	if err != nil {
		return false, false, errors.Wrap(err, "failed to check phone number")
	}

	// Log the results for debugging purposes
	fmt.Printf("isExistEmail: %t, isExistPhone: %t\n", userExist.IsExistEmail, userExist.IsExistPhone)

	return userExist.IsExistEmail, userExist.IsExistPhone, nil
}

// Store creates a new user.
func (r *UserRepository) Store(ctx context.Context, user models.User) error {
	err := facades.Orm().Query().Create(&user)
	if err != nil {
		if gormErr := errors.Unwrap(err); errors.Is(gormErr, gorm.ErrDuplicatedKey) {
			facades.Log().Error("Error: %v", gormErr)
			return errors.New("phone number already exists")
		}
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}
