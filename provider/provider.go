package provider

import (
	"github.com/apostrohedottilde/indymorning/api/database"
)

type Provider struct {
}

func (provider *Provider) ProjectRepository() *database.ProjectRepository {
	return database.New()
}

func (provider *Provider) Close() {

}

func New() *Provider {
	return &Provider{}
}
