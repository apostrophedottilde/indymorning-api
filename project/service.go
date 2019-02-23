package project

type Service interface {
	Create(loggedInUser string, model GameProject) (GameProject, error)
	FindOne(loggedInUser string, id string) (GameProject, error)
	FindAll(loggedInUser string) ([]GameProject, error)
	Update(loggedInUser string, id string, update GameProject) (GameProject, error)
	Delete(loggedInUser string, s string) error
}

// ProjectService struct.
type ProjectService struct {
	repository *ProjectRepository
}

// Create a new Project.
func (ps *ProjectService) Create(loggedInUser string, model GameProject) (GameProject, error) {
	model.Creator = loggedInUser
	dummy, err := ps.repository.Create(model)
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

// FindOne Project by ID.
func (ps *ProjectService) FindOne(loggedInUser string, id string) (GameProject, error) {
	dummy, err := ps.repository.FindOne(id)
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

// FindAll projects.
func (ps *ProjectService) FindAll(loggedInUser string) ([]GameProject, error) {
	dummy, err := ps.repository.FindAll()
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

// Update an existing project by ID with the data in 'update' struct.
func (ps *ProjectService) Update(loggedInUser string, id string, update GameProject) (GameProject, error) {
	dummy, err := ps.repository.Update(id, update)
	if err != nil {
		panic(err)
	}
	return dummy, nil
}

// Delete a project by ID.
func (ps *ProjectService) Delete(loggedInUser string, id string) error {
	err := ps.repository.Delete(id)
	if err != nil {
		panic(err)
	}
	return err
}

// NewService factory function. Takes ProjectRepository then constructs and returns ProjectService.
func NewService(repo *ProjectRepository) *ProjectService {
	return &ProjectService{
		repository: repo,
	}
}
