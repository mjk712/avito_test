package service

import (
	"avitotest/internal/models"
	"avitotest/internal/repo"
)

type User interface {
	AddUserIntoSegment(userSegments *models.UserSegments) error
	GetUserSegments(userId string) ([]*models.UserSegment, error)
}

type Segment interface {
	CreateSegment(name string) error
	DeleteSegment(name string) error
}

type Services struct {
	User    User
	Segment Segment
}

type ServicesDependencies struct {
	Repos *repo.Repositories
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		User:    NewUserService(deps.Repos.User),
		Segment: NewSegmentService(deps.Repos.Segment),
	}
}
