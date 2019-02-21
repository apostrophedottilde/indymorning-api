package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/apostrohedottilde/indymorning/api/adapter"
	pHandler "github.com/apostrohedottilde/indymorning/api/project/handler"
	pService "github.com/apostrohedottilde/indymorning/api/project/service"
	"github.com/apostrohedottilde/indymorning/api/provider"
)

func main() {
	p := provider.New()
	projRepo := p.ProjectRepository()
	projService := pService.New(projRepo)
	projHandler := pHandler.New(projService)
	adapter := adapter.New(projHandler)

	adapter.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	defer close(stop)

	adapter.Start()

	<-stop

	adapter.Stop()
}
