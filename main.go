package main

import (
	"github.com/goravel/framework/facades"

	"goravel/bootstrap"
)

func main() {
	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	go facades.Schedule().Run()

	//Start http server by facades.Route().
	go func() {
		if err := facades.Route().Run(); err != nil {
			facades.Log().Errorf("Route run error: %v", err)
		}
	}()

	select {}
}
