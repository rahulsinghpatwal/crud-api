package model

import (
	"crud/internal/entity"
	"crud/utils"
)

type Repository interface {
	GetUserById(string) (*utils.UserResponse, error)
	ListUsers(int, int) (*[]utils.UserResponse, error)
	CreateUser(entity.User) (*utils.UserResponse, error)
	DeleteUser(string, utils.UpdateUserRequest) (*utils.DeleteResponse, error)
	UpdateUser(string) (*utils.UpdateResponse, error)
	GetUserByEmail(string, string) error
}

type Service interface {
	GetUserById(string) (*utils.UserResponse, error)
	ListUsers(int, int) (*[]utils.UserResponse, error)
	CreateUser(entity.User) (*utils.UserResponse, error)
	DeleteUser(string) (*utils.DeleteResponse, error)
	UpdateUser(string, utils.UpdateUserRequest) (*utils.UpdateResponse, error)
	GetUserByEmail(string, string) error
}
