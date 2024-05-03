package models

import "github.com/goravel/framework/database/orm"

type Roles struct {
	orm.Model
	RoleName    string
	Permissions string 	
	orm.SoftDeletes
}
