package models

import "github.com/goravel/framework/database/orm"

type Devices struct {
	orm.Model
	Token    string
	DeviceID string
	DeviceIP string
	UserID   uint
	orm.SoftDeletes
}
