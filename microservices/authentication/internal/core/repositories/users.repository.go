package repositories

import (
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/entities"
	"github.com/surrealdb/surrealdb.go"
)

type IUserRepository interface {
	Find() (users []entities.User, err error)
	FindById(userId string) (user entities.User, err error)
	FindByEmail(email string) (user entities.User, err error)
	FindByIdAndRefreshToken(userId string, refreshToken string) (err error)
	Create(user entities.User) (userCreated entities.User, err error)
	CreateEmailVerificationToken(userId string, token string) (err error)
	CreateResetPasswordToken(userId string, token string) (err error)
	UpdateLastSignIn(userId string) (err error)
	UpdatePassword(userId string, password string) (err error)
	VerifyEmail(userId string) (err error)
	RevokeRefreshToken(userId string, refreshToken string) (err error)
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

func (r *UserRepository) FindByIdAndRefreshToken(userId string, refreshToken string) (err error) {
	return nil
}

func (r *UserRepository) Create(user entities.User) (userCreated entities.User, err error) {
	return userCreated, nil
}

func (r *UserRepository) CreateEmailVerificationToken(userId string, token string) (err error) {
	return nil
}

func (r *UserRepository) CreateResetPasswordToken(userId string, token string) (err error) {
	return nil
}

func (r *UserRepository) UpdateLastSignIn(userId string) (err error) {
	return nil
}

func (r *UserRepository) UpdatePassword(userId string, password string) (err error) {
	return nil
}

func (r *UserRepository) VerifyEmail(userId string) (err error) {
	return nil
}

func (r *UserRepository) RevokeRefreshToken(userId string, refreshToken string) (err error) {
	return nil
}
