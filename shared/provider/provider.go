package provider

import (
	"github.com/apostrophedottilde/go-forum-api/forum"
	"github.com/apostrophedottilde/go-forum-api/user"
)

type Provider struct {
}

func (provider *Provider) UserRepository() *user.Repository {
	return user.NewRepository()
}

func (provider *Provider) ForumRepository() *forum.Repository {
	return forum.NewRepository()
}

func (provider *Provider) Close() {

}

func New() *Provider {
	return &Provider{}
}
