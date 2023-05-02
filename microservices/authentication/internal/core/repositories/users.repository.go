package repositories

import (
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/entities"
	"github.com/surrealdb/surrealdb.go"
)

type IUserRepository interface {
	FindById(userId string) (user entities.User, err error)
	FindByEmail(email string) (user entities.User, err error)
	FindRefreshToken(refreshToken string) (user entities.User, err error)
	Create(user entities.User) (userData entities.User, err error)
}

type UserRepository struct {
	Database *surrealdb.DB
}

func (r *UserRepository) FindById(userId string) (user entities.User, err error) {
	return entities.User{}, nil
}

func (r *UserRepository) FindByEmail(email string) (user entities.User, err error) {
	return entities.User{}, nil
}

func (r *UserRepository) Create(user entities.User) (err error) {
	return nil
}
