package provider

import (
	"github.com/apostrohedottilde/indymorning/api/project"
)

type Provider struct {
}

func (provider *Provider) ProjectRepository() *project.ProjectRepository {
	return project.NewRepository()
}

func (provider *Provider) Close() {

}

func New() *Provider {
	return &Provider{}
}
