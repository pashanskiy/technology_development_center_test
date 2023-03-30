package storage

import (
	apiUserService "technology_development_center_test/api/user"

	"gorm.io/gorm"
)

type UserService struct {
	apiUserService.UserServiceServer

	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}
