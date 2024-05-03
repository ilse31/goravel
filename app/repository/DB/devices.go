package db

import (
	"context"
	"goravel/app/models"

	// "github.com/goravel/framework/database/gorm"
	"github.com/goravel/framework/facades"
	"github.com/pkg/errors"
)

type DevicesRepository struct {
}

// NewDevicesRepository creates a new user repository
func NewDevicesRepository() *DevicesRepository {
	return &DevicesRepository{}
}

func (r *DevicesRepository) GetTokenByUserID(ctx context.Context, ID uint) (models.Devices, error) {

	device := models.Devices{}
	q := facades.Orm().Query().Model(&device)
	err := q.Where("user_id = ?", ID).First(&device)
	if err != nil {
		return models.Devices{}, errors.Wrap(err, "failed to get devices by userid")
	}
	return device, nil
}

func (r *DevicesRepository) CreateDevice(ctx context.Context, device models.Devices) (models.Devices, error) {
	q := facades.Orm().Query().Model(&device)
	err := q.Create(&device)
	if err != nil {
		return models.Devices{}, errors.Wrap(err, "failed to create device")
	}
	return device, nil
}

func (r *DevicesRepository) UpdateDevice(ctx context.Context, device models.Devices, user models.User, token string, deviceID, networkIp string) (int, error) {

	result, err := facades.Orm().Query().Where("user_id = ?", user.ID).Update(&models.Devices{
		UserID:   user.ID,
		Token:    token,
		DeviceID: deviceID,
		DeviceIP: networkIp,
	})

	if err != nil {
		return 0, errors.Wrap(err, "failed to update device")
	}

	return int(result.RowsAffected), nil
}
