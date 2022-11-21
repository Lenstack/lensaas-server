package services

import (
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/core/entities"
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/core/repositories"
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/util"
	"github.com/Masterminds/squirrel"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type IAuthenticationService interface {
	SignIn(email string, password string) (token string, expiration string, err error)
	SignUp(email string, password string) (userId string, err error)
	SignOut(token string) (message string, err error)
	RefreshToken(token string) (newToken string, err error)
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

func (as *AuthenticationService) SignIn(email string, password string) (token string, expiration string, err error) {
	user, err := as.authenticationRepository.GetUserByEmail(email)
	if err != nil {
		return "", "", status.Errorf(codes.Internal, "error while getting user", err)
	}

	isPasswordMatched, _ := as.bcryptManager.CompareHashedPassword(user.Password, password)
	if !isPasswordMatched {
		return "", "", status.Errorf(codes.Internal, "wrong password")
	}

	expirationTime, err := time.ParseDuration(viper.Get("JWT_EXPIRATION_TIME").(string))
	generateToken, err := as.jwtManager.GenerateToken(user.Id, expirationTime)
	if err != nil {
		return "", "", status.Errorf(codes.Internal, "error while generating token")
	}

	err = as.redisRepository.SetHashValue(user.Id, uuid.New().String(), generateToken, expirationTime)
	if err != nil {
		return "", "", err
	}

	expiration = time.Now().Add(expirationTime).Format(time.RFC3339)
	return generateToken, expiration, nil
}

func (as *AuthenticationService) SignUp(email string, password string) (userId string, err error) {
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

	userId, err = as.authenticationRepository.CreateUser(user)
	if err != nil {
		return "", status.Errorf(codes.Internal, "error while creating user")
	}

	return userId, nil
}

func (as *AuthenticationService) SignOut(token string) (message string, err error) {
	return "SignOut", nil
}

func (as *AuthenticationService) RefreshToken(token string) (newToken string, err error) {
	userId, err := as.jwtManager.ValidateToken(token)
	if err != nil {
		return "", status.Errorf(codes.Internal, "error while parsing token")
	}

	expirationTime, err := time.ParseDuration(viper.Get("JWT_EXPIRATION_TIME").(string))
	generateToken, err := as.jwtManager.GenerateToken(userId, expirationTime)
	if err != nil {
		return "", status.Errorf(codes.Internal, "error while generating token")
	}
	return generateToken, nil
}
