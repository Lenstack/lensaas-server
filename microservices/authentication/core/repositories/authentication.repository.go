package repositories

import (
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/core/entities"
	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
)

type IAuthenticationRepository interface {
	CreateUser(user entities.User) (userId string, err error)
	GetUserByEmail(email string) (user entities.User, err error)
	GetUserIdByEmail(email string) (userId string, err error)
}

type AuthenticationRepository struct {
	Database squirrel.StatementBuilderType
}

func (ar *AuthenticationRepository) CreateUser(user entities.User) (userId string, err error) {
	bq := ar.Database.
		Insert(entities.UserTableName).
		Columns("Id", "Name", "Avatar", "Email", "Password",
			"Verified", "Code", "Token", "Roles").
		Values(&user.Id, &user.Name, &user.Avatar, &user.Email,
			&user.Password, &user.Verified, &user.Code, &user.Token, pq.Array(&user.Roles)).
		Suffix("RETURNING Id")

	err = bq.QueryRow().Scan(&userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (ar *AuthenticationRepository) GetUserByEmail(email string) (user entities.User, err error) {
	bq := ar.Database.Select("Id", "Name", "Avatar", "Email", "Password",
		"Verified", "Code", "Token", "Roles", "CreatedAt", "UpdatedAt").
		From(entities.UserTableName).
		Where(squirrel.Eq{"email": email})

	err = bq.QueryRow().Scan(&user.Id, &user.Name, &user.Avatar, &user.Email,
		&user.Password, &user.Verified, &user.Code, &user.Token, pq.Array(&user.Roles),
		&user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (ar *AuthenticationRepository) GetUserIdByEmail(email string) (userId string, err error) {
	bq := ar.Database.Select("Id").
		From(entities.UserTableName).
		Where(squirrel.Eq{"email": email})

	err = bq.QueryRow().Scan(&userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}
