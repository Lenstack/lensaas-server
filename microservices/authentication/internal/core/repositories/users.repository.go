package repositories

import (
	"fmt"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/entities"
	"github.com/surrealdb/surrealdb.go"
)

type IUserRepository interface {
	Find() (users []entities.User, err error)
	FindById(userId string) (user entities.User, err error)
	FindByEmail(email string) (user entities.User, err error)
	FindRefreshToken(refreshToken string) (user entities.User, err error)
	Create(user entities.User) (userData entities.User, err error)
	UpdateLastSignIn(userId string) (err error)
}

type UserRepository struct {
	Database *surrealdb.DB
}

func (r *UserRepository) Find() (users []entities.User, err error) {
	response, err := r.Database.Select("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	fmt.Println(response)
	return users, nil
}

func (r *UserRepository) FindById(userId string) (user entities.User, err error) {
	return entities.User{}, nil
}

func (r *UserRepository) FindByEmail(email string) (user entities.User, err error) {
	return entities.User{}, nil
}

func (r *UserRepository) Create(user entities.User) (err error) {
	data, err := r.Database.Create(entities.UserTableName, user)
	if err != nil {
		return err
	}

	fmt.Println(data)
	return nil
}

func (r *UserRepository) UpdateLastSignIn(userId string) (err error) {
	return nil
}
