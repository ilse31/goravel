package console

import (
	"fmt"
	"goravel/app/models"

	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/schedule"
	"github.com/goravel/framework/facades"
)

type Kernel struct {
}

func (kernel *Kernel) Schedule() []schedule.Event {
	return []schedule.Event{
		facades.Schedule().Call(func() {
			fmt.Println("This is a scheduled task")
			facades.Log().Info("This is a scheduled task")
			users := make([]models.User, 0)

			facades.Orm().Query().Model(users).Where("status = ?", "nonactive").Find(&users)

			for _, user := range users {
				facades.Orm().Query().Model(&user).Update("status", "active")
			}
		}).EveryMinute(),
	}
}

func (kernel *Kernel) Commands() []console.Command {
	return []console.Command{}
}
