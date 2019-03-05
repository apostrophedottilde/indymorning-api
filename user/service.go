package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type service interface {
	Register(model User) (user, error)
	FindOne(loggedInUser string, id string) (user, error)
	FindAll(loggedInUser string) ([]user, error)
	Update(loggedInUser string, id string, update User) (user, error)
	Delete(loggedInUser string, s string) error
}

// Service struct.
type Service struct {
	repository *Repository
}

func (ps *Service) Register(model User) (user, error) {
	fmt.Println("wont fuckin work")
	hash, err := bcrypt.GenerateFromPassword([]byte("users-password"), bcrypt.MinCost)
	model.Password = string(hash)
	fmt.Println("bcrypted password", model.Password)
	fmt.Println("wont fuckin work")
	dummy, err := ps.repository.Create(model)
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

// FindOne User by ID.
func (ps *Service) FindOne(loggedInUser string, id string) (user, error) {
	dummy, err := ps.repository.FindOne(id)
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

// FindAll projects.
func (ps *Service) FindAll(loggedInUser string) ([]user, error) {
	dummy, err := ps.repository.FindAll()
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

// Update an existing project by ID with the data in 'update' struct.
func (ps *Service) Update(loggedInUser string, id string, update User) (user, error) {
	dummy, err := ps.repository.Update(id, update)
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

// Delete a project by ID.
func (ps *Service) Delete(loggedInUser string, id string) error {
	err := ps.repository.Delete(id)
	if err != nil {
		panic(err)
	}
	return err
}

// NewService factory function. Takes Repository then constructs and returns Service.
func NewService(repo *Repository) *Service {
	return &Service{
		repository: repo,
	}
}
