package provider

import (
	"github.com/apostrohedottilde/indymorning/api/project"
	"github.com/apostrohedottilde/indymorning/api/user"
)

type Provider struct {
}

func (provider *Provider) UserRepository() *user.UserRepository {
	return user.NewRepository()
}

func (provider *Provider) ProjectRepository() *project.ProjectRepository {
	return project.NewRepository()
}

func (provider *Provider) Close() {

}

func New() *Provider {
	return &Provider{}
}
