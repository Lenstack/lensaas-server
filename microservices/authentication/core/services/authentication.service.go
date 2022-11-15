package services

import (
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/core/entities"
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/core/repositories"
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/util"
	"github.com/Masterminds/squirrel"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type IAuthenticationService interface {
	SignIn(email string, password string) (token string, err error)
	SignUp(email string, password string) (token string, err error)
	SignOut(token string) (message string, err error)
}

type AuthenticationService struct {
	authenticationRepository repositories.AuthenticationRepository
	redisRepository          repositories.RedisRepository
	jwtManager               util.JwtManager
	bcryptManager            util.BcryptManager
}

func NewAuthenticationService(database squirrel.StatementBuilderType, redisManager *redis.Client, jwtManager util.JwtManager) *AuthenticationService {
	return &AuthenticationService{
		authenticationRepository: repositories.AuthenticationRepository{
			Database: database,
		},
		redisRepository: repositories.RedisRepository{
			Client: redisManager,
		},
		jwtManager: jwtManager,
	}
}

func (as *AuthenticationService) SignIn(email string, password string) (token string, err error) {
	user, err := as.authenticationRepository.GetUserByEmail(email)
	if err != nil {
		return "", status.Errorf(codes.Internal, "error while getting user", err)
	}

	isPasswordMatched, _ := as.bcryptManager.CompareHashedPassword(user.Password, password)
	if !isPasswordMatched {
		return "", status.Errorf(codes.Internal, "wrong password")
	}

	generateToken, err := as.jwtManager.GenerateToken(user.Id)
	if err != nil {
		return "", status.Errorf(codes.Internal, "error while generating token")
	}

	expirationTime, err := time.ParseDuration("1h")
	err = as.redisRepository.SetHashValue(user.Id, uuid.New().String(), generateToken, expirationTime)
	if err != nil {
		return "", err
	}

	return generateToken, nil
}

func (as *AuthenticationService) SignUp(email string, password string) (token string, err error) {
	isUserExist, err := as.authenticationRepository.GetUserIdByEmail(email)
	if isUserExist != "" {
		return "", status.Errorf(codes.Internal, "this email is already registered")
	}

	hashedPassword, err := as.bcryptManager.HashPassword(password)
	if err != nil {
		return "", status.Errorf(codes.Internal, "error while hashing password")
	}

	user := entities.User{
		Id:       uuid.New().String(),
		Email:    email,
		Password: hashedPassword,
	}

	userId, err := as.authenticationRepository.CreateUser(user)
	if err != nil {
		return "", status.Errorf(codes.Internal, "error while creating user")
	}

	generateToken, err := as.jwtManager.GenerateToken(userId)
	if err != nil {
		return "", status.Errorf(codes.Internal, "error while generating token")
	}

	expirationTime, err := time.ParseDuration("1h")
	err = as.redisRepository.SetHashValue(userId, uuid.New().String(), generateToken, expirationTime)
	if err != nil {
		return "", err
	}

	return generateToken, nil
}

func (as *AuthenticationService) SignOut(token string) (message string, err error) {
	return "SignOut", nil
}
