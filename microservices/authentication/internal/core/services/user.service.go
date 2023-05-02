package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/entities"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/repositories"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/utils"
	"github.com/surrealdb/surrealdb.go"
	"golang.org/x/oauth2"
)

type IUserService interface {
	SignInWithCredentials(email string, password string) (accessToken string, refreshToken string, err error)
	SignInWithOAuth(provider string, state string) (url string, err error)
	SignInWithOAuthCallback(provider string, code string) (accessToken string, refreshToken string, err error)
	SignUp(name string, email string, password string) (err error)
	SignOut(userId string) (err error)
	RefreshToken(userId string) (accessToken string, err error)
	ForgotPassword(email string) (err error)
	ResetPassword(userId string, password string) (err error)
	Me(userId string) (user entities.User, err error)
	VerifyEmail(userId string) (err error)
}

type UserService struct {
	OauthClientConfig oauth2.Config
	Jwt               utils.Jwt
	Bcrypt            utils.Bcrypt
	Email             utils.Email
	UserRepository    repositories.UserRepository
}

func NewUserService(database *surrealdb.DB, jwt utils.Jwt, bcrypt utils.Bcrypt, email utils.Email) *UserService {
	return &UserService{
		UserRepository: repositories.UserRepository{
			Database: database,
		},
		Jwt:    jwt,
		Bcrypt: bcrypt,
		Email:  email,
	}
}

func (s *UserService) SignInWithCredentials(email string, password string) (accessToken string, refreshToken string, err error) {
	// If user doesn't exist then return error user not found
	userByEmail, err := s.UserRepository.FindByEmail(email)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	// If user is disabled or not verified then return error user is disabled or not verified
	if userByEmail.Disabled || userByEmail.Verified {
		return "", "", errors.New("user is disabled or not verified")
	}

	// Compare password with hash
	err = s.Bcrypt.ComparePassword(userByEmail.Password, password)
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}

	// Generate access token and refresh token
	accessToken, err = s.Jwt.GenerateToken(userByEmail.ID, s.Jwt.ExpirationTimeAccess)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = s.Jwt.GenerateToken(userByEmail.ID, s.Jwt.ExpirationTimeRefresh)
	if err != nil {
		return "", "", err
	}

	// Update user last sign in date and ip address and user agent and device type
	err = s.UserRepository.UpdateLastSignIn(userByEmail.ID)
	if err != nil {
		return "", "", errors.New("cannot update user last sign in")
	}

	return accessToken, refreshToken, nil
}

func (s *UserService) SignInWithOAuth(provider string, state string) (url string, err error) {
	// Get oauth providers from config
	oauthProviders := utils.NewOauth()
	// Search provider in database and get client id and client secret
	oauthProvider, err := oauthProviders.GetProvider(provider)
	if err != nil {
		return "", err
	}
	// Set oauth client config
	s.OauthClientConfig = oauth2.Config{
		ClientID:     oauthProvider.Config.ClientId,
		ClientSecret: oauthProvider.Config.ClientSecret,
		RedirectURL:  oauthProvider.Config.RedirectUrl,
		Scopes:       oauthProvider.Config.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:   oauthProvider.Config.AuthUrl,
			TokenURL:  oauthProvider.Config.TokenUrl,
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}
	// Generate auth code url with state
	url = s.OauthClientConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return url, nil
}

func (s *UserService) SignInWithOAuthCallback(provider string, code string) (accessToken string, refreshToken string, err error) {
	// Get oauth providers from config
	oauthProviders := utils.NewOauth()

	// Search provider in database and get client id and client secret
	oauthProvider, err := oauthProviders.GetProvider(provider)
	if err != nil {
		return "", "", err
	}

	// Exchange auth code for token
	token, err := s.OauthClientConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return "", "", errors.New("cannot exchange auth code for token")
	}

	// Get user information from oauth provider
	response, err := s.OauthClientConfig.Client(oauth2.NoContext, token).Get(oauthProvider.Config.UserInfoUrl) // get url from oauth provider
	if err != nil {
		return "", "", errors.New("cannot get user info from oauth provider")
	}

	// Close response body
	defer response.Body.Close()

	// Decode user information
	var userInformation map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&userInformation)
	if err != nil {
		return "", "", errors.New("cannot decode user information")
	}

	// Get user information from oauth provider
	information, err := oauthProviders.GetUserInformation(provider, userInformation)
	if err != nil {
		return "", "", err
	}

	// Search user in database by email
	userByEmail, err := s.UserRepository.FindByEmail(information.Email)
	if err != nil {
		return "", "", err
	}

	fmt.Println(userByEmail)

	// If user doesn't exist then create user in database

	// Generate access token and refresh token with user id
	accessToken, err = s.Jwt.GenerateToken(userByEmail.ID, s.Jwt.ExpirationTimeAccess)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = s.Jwt.GenerateToken(userByEmail.ID, s.Jwt.ExpirationTimeRefresh)
	if err != nil {
		return "", "", err
	}

	// Create session in database with user id and access token and refresh token

	// Update user last login date and ip address and user agent and device type
	err = s.UserRepository.UpdateLastSignIn(userByEmail.ID)
	if err != nil {
		return "", "", errors.New("cannot update user last sign in")
	}

	return accessToken, refreshToken, nil
}

func (s *UserService) SignUp(name string, email string, password string) (err error) {
	// Search user in database by email
	userByEmail, _ := s.UserRepository.FindByEmail(email)
	if userByEmail.ID != "" {
		return errors.New("user already exists")
	}

	// Hash password with bcrypt
	hashedPassword, err := s.Bcrypt.HashPassword(password)
	if err != nil {
		return err
	}

	// Create user in database
	err = s.UserRepository.Create(entities.User{
		Profile:  entities.Profile{Name: name},
		Email:    email,
		Password: hashedPassword,
		Roles: []entities.Role{
			{
				Name: "user",
				Permissions: []entities.Permission{
					{Scope: "users", Action: "update"},
					{Scope: "users", Action: "delete"},
					{Scope: "users", Action: "create"},
					{Scope: "users", Action: "list"},
				},
			},
		},
	})

	// Create email verification and add to queue
	err = s.Email.Add(utils.Email{
		From:    s.Email.From,
		To:      []string{email},
		Subject: "Verify email",
		Body:    "<h1>Verify email</h1>",
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) SignOut(userId string) (err error) {
	// Search and revoke access, refresh token from database
	return nil
}

func (s *UserService) RefreshToken(userId string) (accessToken string, err error) {
	// Generate access token
	accessToken, err = s.Jwt.GenerateToken(userId, s.Jwt.ExpirationTimeAccess)
	if err != nil {
		return "", err
	}

	// Search and revoke access, refresh token from database

	return accessToken, nil
}

func (s *UserService) ForgotPassword(email string) (err error) {
	// Search user in database by email

	// Generate reset password token

	// Send email with reset password token
	err = s.Email.Add(utils.Email{
		From:    s.Email.From,
		To:      []string{email},
		Subject: "Reset password",
		Body:    "<h1>Reset password</h1>",
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) ResetPassword(userId string, password string) (err error) {
	// Search user in database by userId

	// Update password in database

	// Send email with reset password confirmation
	err = s.Email.Add(utils.Email{
		From:    s.Email.From,
		To:      []string{""},
		Subject: "Reset password confirmation",
		Body:    "<h1>Reset password confirmation</h1>",
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Me(userId string) (user entities.User, err error) {
	return s.UserRepository.FindById(userId)
}

func (s *UserService) VerifyEmail(userId string) (err error) {
	// Search user in database by userId
	user, err := s.UserRepository.FindById(userId)
	if err != nil {
		return errors.New("user not found")
	}
	// Update user verified in database
	fmt.Println(user)

	// Send email with email verification confirmation
	err = s.Email.Add(utils.Email{
		From:    s.Email.From,
		To:      []string{""},
		Subject: "Email verification confirmation",
		Body:    "<h1>Email verification confirmation</h1>",
	})
	if err != nil {
		return err
	}

	return nil
}