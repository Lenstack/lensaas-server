package repositories

import (
	"fmt"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/entities"
	"github.com/surrealdb/surrealdb.go"
)

type IUserRepository interface {
	FindById(userId string) (user entities.User, err error)
	FindByEmail(email string) (user entities.User, err error)
	Create(user entities.User) (userData entities.User, err error)
}

type UserRepository struct {
	Database *surrealdb.DB
}

func (r *UserRepository) FindById(userId string) (user entities.User, err error) {
	query, err := r.Database.Query(entities.UserTableName, userId)
	if err != nil {
		return entities.User{}, err
	}
	fmt.Println(query)
	return
}

func (r *UserRepository) FindByEmail(email string) (user entities.User, err error) {
	return
}

func (r *UserRepository) Create(user entities.User) (err error) {
	userCreated, err := r.Database.Create(entities.UserTableName, user)
	if err != nil {
		return err
	}
	fmt.Println(userCreated)
	return nil
}
