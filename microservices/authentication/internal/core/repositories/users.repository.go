package repositories

import (
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/entities"
	"github.com/surrealdb/surrealdb.go"
)

type IUserRepository interface {
	Find() (users []entities.User, err error)
	FindById(userId string) (user entities.User, err error)
	FindByEmail(email string) (user entities.User, err error)
	Create(user entities.User) (err error)
	CreateVerificationToken(userId string, token string) (err error)
	CreateResetPasswordToken(userId string, token string) (err error)
	UpdateLastSignIn(userId string) (err error)
}

type UserRepository struct {
	Database *surrealdb.DB
}

func (r *UserRepository) Find() (users []entities.User, err error) {
	return users, nil
}

func (r *UserRepository) FindById(userId string) (user entities.User, err error) {
	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (user entities.User, err error) {
	return user, nil
}

func (r *UserRepository) Create(user entities.User) (err error) {
	return nil
}

func (r *UserRepository) CreateVerificationToken(userId string, token string) (err error) {
	return nil
}

func (r *UserRepository) CreateResetPasswordToken(userId string, token string) (err error) {
	return nil
}

func (r *UserRepository) UpdateLastSignIn(userId string) (err error) {
	return nil
}
