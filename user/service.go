package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(model User) (user, error)
	FindOne(loggedInUser string, id string) (user, error)
	FindAll(loggedInUser string) ([]user, error)
	Update(loggedInUser string, id string, update User) (user, error)
	Delete(loggedInUser string, s string) error
}

// UserService struct.
type UserService struct {
	repository *UserRepository
}

func (ps *UserService) Register(model User) (user, error) {
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
func (ps *UserService) FindOne(loggedInUser string, id string) (user, error) {
	dummy, err := ps.repository.FindOne(id)
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

// FindAll projects.
func (ps *UserService) FindAll(loggedInUser string) ([]user, error) {
	dummy, err := ps.repository.FindAll()
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

// Update an existing project by ID with the data in 'update' struct.
func (ps *UserService) Update(loggedInUser string, id string, update User) (user, error) {
	dummy, err := ps.repository.Update(id, update)
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

// Delete a project by ID.
func (ps *UserService) Delete(loggedInUser string, id string) error {
	err := ps.repository.Delete(id)
	if err != nil {
		panic(err)
	}
	return err
}

// NewService factory function. Takes UserRepository then constructs and returns UserService.
func NewService(repo *UserRepository) *UserService {
	return &UserService{
		repository: repo,
	}
}
