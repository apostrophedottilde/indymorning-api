package forum

type service interface {
	Create(loggedInUser string, model Forum) (Forum, error)
	FindOne(loggedInUser string, id string) (Forum, error)
	FindAll(loggedInUser string) ([]Forum, error)
	Update(loggedInUser string, id string, update Forum) (Forum, error)
	Delete(loggedInUser string, s string) error
}

// Service struct.
type Service struct {
	repository *Repository
}

// Create a new forum.
func (ps *Service) Create(loggedInUser string, model Forum) (Forum, error) {
	model.Creator = loggedInUser
	dummy, err := ps.repository.Create(model)
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

// FindOne forum by ID.
func (ps *Service) FindOne(loggedInUser string, id string) (Forum, error) {
	dummy, err := ps.repository.FindOne(id)
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

// FindAll projects.
func (ps *Service) FindAll(loggedInUser string) ([]Forum, error) {
	dummy, err := ps.repository.FindAll()
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

// Update an existing project by ID with the data in 'update' struct.
func (ps *Service) Update(loggedInUser string, id string, update Forum) (Forum, error) {
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
