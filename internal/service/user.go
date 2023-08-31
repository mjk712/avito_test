package service

import (
	"avitotest/internal/models"
	"avitotest/internal/repo"
)

type UserService struct {
	userRepo repo.User
}

func NewUserService(userRepo repo.User) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) AddUserIntoSegment(userSegments *models.UserSegments) error {
	return s.userRepo.AddUserIntoSegment(userSegments)
}

func (s *UserService) GetUserSegments(userId string) ([]*models.UserSegment, error) {
	userSegments, err := s.userRepo.GetUserSegments(userId)
	if err != nil {
		return nil, err
	}
	return userSegments, nil
}
