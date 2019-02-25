package main

import (
	"github.com/apostrophedottilde/indymorning-api/shared/adapter"
	"github.com/apostrophedottilde/indymorning-api/shared/provider"
	"github.com/apostrophedottilde/indymorning-api/user"
	"os"
	"os/signal"
	"syscall"

	"github.com/apostrophedottilde/indymorning-api/project"
)

func main() {
	p := provider.New()

	userRepo := p.UserRepository()
	userService := user.NewService(userRepo)
	userController := user.NewController(userService)

	projRepo := p.ProjectRepository()
	projService := project.NewService(projRepo)
	projController := project.NewController(projService)
	adapter := adapter.New(userController, projController)

	adapter.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	defer close(stop)

	adapter.Start()

	<-stop

	adapter.Stop()
}
