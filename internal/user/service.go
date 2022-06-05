package user

import (
	"crud/internal/entity"
	"crud/internal/user/repo"
	"crud/utils"

	"go.uber.org/zap"
)

type Service struct {
	repo repo.Repository
}

func NewService(repo repo.Repository) Service {
	return Service{repo}
}

func (s Service) GetUserById(userId string) (*utils.UserResponse, error) {
	response, err := s.repo.GetUserById(userId)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}

	return response, nil
}

func (s Service) ListUsers(l int, offset int) (*[]utils.UserResponse, error) {
	response, err := s.repo.ListUsers(l, offset)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}

	return response, nil
}
func (s Service) CreateUser(user entity.User) (*utils.UserResponse, error) {
	responesUser, err := s.repo.CreateUser(user)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	return responesUser, nil
}

func (s Service) DeleteUser(userID string) (*utils.DeleteResponse, error) {
	response, err := s.repo.DeleteUser(userID)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	return response, nil
}
func (s Service) UpdateUser(userID string, user utils.UpdateUserRequest) (*utils.UpdateResponse, error) {
	response, err := s.repo.UpdateUser(userID, user)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	return response, nil
}

func (s Service) GetUserByEmail(email string, password string) error {
	err := s.repo.GetUserByEmail(email, password)
	if err != nil {
		zap.S().Error(err)
		return err
	}

	return nil
}
