package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/apostrohedottilde/indymorning/api/adapter"
	"github.com/apostrohedottilde/indymorning/api/project"
	"github.com/apostrohedottilde/indymorning/api/provider"
)

func main() {
	p := provider.New()
	projRepo := p.ProjectRepository()
	projService := project.NewService(projRepo)
	projController := project.NewController(projService)
	adapter := adapter.New(projController)

	adapter.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	defer close(stop)

	adapter.Start()

	<-stop

	adapter.Stop()
}
