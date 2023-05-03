package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/entities"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/core/repositories"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/utils"
	"github.com/surrealdb/surrealdb.go"
	"golang.org/x/oauth2"
	"io"
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
	ResendEmailVerification(email string) (err error)
	GenerateMFAQRCode(userId string) (qrCode string, err error)
	EnableMFA(userId string, code string) (err error)
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
	token, err := s.OauthClientConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", "", errors.New("cannot exchange auth code for token")
	}

	// Get user information from oauth provider
	response, err := s.OauthClientConfig.Client(context.Background(), token).Get(oauthProvider.Config.UserInfoUrl) // get url from oauth provider
	if err != nil {
		return "", "", errors.New("cannot get user info from oauth provider")
	}

	// Close response body
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

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
		return "", "", errors.New("user not found")
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

	// Create user entity
	user := entities.User{
		Profile:  entities.Profile{Name: name},
		Email:    email,
		Password: hashedPassword,
	}

	// Create user in database
	userCreated, err := s.UserRepository.Create(user)
	if err != nil {
		return errors.New("cannot create user")
	}

	// Generate email verification token with user id
	emailVerificationToken, err := s.Jwt.GenerateToken(userCreated.ID, s.Jwt.ExpirationTimeEmailVerification)
	if err != nil {
		return err
	}

	// Create email verification token in database
	err = s.UserRepository.CreateEmailVerificationToken(userCreated.ID, emailVerificationToken)
	if err != nil {
		return err
	}

	// Create email verification and add to queue
	err = s.Email.Add(utils.Email{
		From:    s.Email.From,
		To:      []string{email},
		Subject: "Verify email",
		Body:    "<h1>Verify email</h1>" + emailVerificationToken,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) SignOut(userId string, refreshToken string) (err error) {
	// Search session in database by user id and refresh token
	err = s.UserRepository.FindByIdAndRefreshToken(userId, refreshToken)
	if err != nil {
		return errors.New("session not found")
	}
	// Revoked refresh token in database
	err = s.UserRepository.RevokeRefreshToken(userId, refreshToken)
	if err != nil {
		return errors.New("cannot revoke refresh token")
	}
	return nil
}

func (s *UserService) RefreshToken(userId string, refreshToken string) (accessToken string, err error) {
	// Search session in database by user id and refresh token
	err = s.UserRepository.FindByIdAndRefreshToken(userId, refreshToken)
	if err != nil {
		return "", errors.New("session not found")
	}

	// Generate access token with user id
	accessToken, err = s.Jwt.GenerateToken(userId, s.Jwt.ExpirationTimeAccess)
	if err != nil {
		return "", errors.New("cannot generate access token")
	}

	// Update session in database with access token

	return accessToken, nil
}

func (s *UserService) ForgotPassword(email string) (err error) {
	// Search user in database by email
	userByEmail, err := s.UserRepository.FindByEmail(email)
	if err != nil {
		return err
	}

	// Generate reset password token
	resetPasswordToken, err := s.Jwt.GenerateToken(userByEmail.ID, s.Jwt.ExpirationTimeResetPassword)
	if err != nil {
		return err
	}

	// Create reset password token in database
	err = s.UserRepository.CreateResetPasswordToken(userByEmail.ID, resetPasswordToken)
	if err != nil {
		return err
	}

	// Send email with reset password token
	err = s.Email.Add(utils.Email{
		From:    s.Email.From,
		To:      []string{email},
		Subject: "Reset password",
		Body:    "<h1>Reset password</h1>" + resetPasswordToken,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) ResetPassword(userId string, resetPasswordToken string, password string) (err error) {
	// Search user in database by userId and reset password token
	err = s.UserRepository.FindByIdAndRefreshToken(userId, resetPasswordToken)
	if err != nil {
		return errors.New("user not found")
	}

	// Encrypt password with bcrypt
	hashedPassword, err := s.Bcrypt.HashPassword(password)
	if err != nil {
		return errors.New("cannot encrypt password")
	}

	// Update password in database
	err = s.UserRepository.UpdatePassword(userId, hashedPassword)
	if err != nil {
		return errors.New("cannot update password")
	}

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

	// If user is already verified then return error
	if user.Verified {
		return errors.New("user already verified")
	}

	// Update user verified in database
	err = s.UserRepository.VerifyEmail(userId)
	if err != nil {
		return errors.New("cannot update user verified")
	}

	// Send email with email verification confirmation
	err = s.Email.Add(utils.Email{
		From:    s.Email.From,
		To:      []string{""},
		Subject: "Email verification confirmation",
		Body:    "<h1>Welcome to the club buddy</h1>",
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) ResendEmailVerification(email string) (err error) {
	// Search user in database by userId
	user, err := s.UserRepository.FindById(email)
	if err != nil {
		return errors.New("user not found")
	}

	// If user is verified then return error
	if user.Verified {
		return errors.New("user is already verified")
	}

	// Generate email verification token with user id
	emailVerificationToken, err := s.Jwt.GenerateToken(user.ID, s.Jwt.ExpirationTimeEmailVerification)
	if err != nil {
		return err
	}

	// Save email verification token in database
	err = s.UserRepository.CreateEmailVerificationToken(user.ID, emailVerificationToken)
	if err != nil {
		return err
	}

	// Create email verification and add to queue
	err = s.Email.Add(utils.Email{
		From:    s.Email.From,
		To:      []string{user.Email},
		Subject: "Verify email",
		Body:    "<h1>Verify email</h1>" + emailVerificationToken,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GenerateMFAQRCode(userId string) (qrCode string, err error) {
	// Search user in database by userId
	user, err := s.UserRepository.FindById(userId)
	if err != nil {
		return "", errors.New("user not found")
	}

	fmt.Println(user.Email)

	// Generate MFA QR Code
	/*
		qrCode, err = s.MFA.GenerateQRCode(user.Email, user.MFASecret)
		if err != nil {
			return "", errors.New("cannot generate MFA QR Code")
		}
	*/
	return qrCode, nil
}

func (s *UserService) EnableMFA(userId string, code string) (err error) {
	return nil
}
