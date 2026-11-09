package main

import (
	"context"

	service "github.com/ArpitSandal/SkillManagerGo/Service"
	"github.com/ArpitSandal/SkillManagerGo/controller"
	"github.com/ArpitSandal/SkillManagerGo/repository"
	"github.com/ArpitSandal/SkillManagerGo/utility"
)

// this will handle the initialization of the router and controllers and services
func main() {

	ctx := context.Background()
	utility.InitializeLogger()

	// initialize the db
	pgxPool := utility.InitializeDB(ctx)

	// initialize the repository
	repo := repository.NewSkillRepository(pgxPool)
	defer pgxPool.Close()

	// initiazlie the service
	service := service.NewSkillService(repo)

	// initialize the controller
	controller := controller.NewSkillController(service)

	router := Router{
		controller: controller,
	}
	// always initialize the router at the end
	router.InitializeRouter()

}
