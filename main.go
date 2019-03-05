package main

import (
	"github.com/apostrophedottilde/go-forum-api/shared/adapter"
	"github.com/apostrophedottilde/go-forum-api/shared/provider"
	"github.com/apostrophedottilde/go-forum-api/user"
	"os"
	"os/signal"
	"syscall"

	"github.com/apostrophedottilde/go-forum-api/forum"
)

func main() {
	p := provider.New()

	userRepo := p.UserRepository()
	userService := user.NewService(userRepo)
	userController := user.NewController(userService)

	projRepo := p.ForumRepository()
	projService := forum.NewService(projRepo)
	projController := forum.NewController(projService)
	adapter := adapter.New(userController, projController)

	adapter.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	defer close(stop)

	adapter.Start()

	<-stop

	adapter.Stop()
}
