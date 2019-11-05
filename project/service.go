package project

type ProjectService struct {
	Repository IProjectRepository
}

type IProjectService interface {
	GetAllProjects(string) ([]Project, error)
}

func (service *ProjectService) GetAllProjects(accountID string) ([]Project, error) {
	return service.Repository.GetAllProjectsForAccount(accountID)
}
