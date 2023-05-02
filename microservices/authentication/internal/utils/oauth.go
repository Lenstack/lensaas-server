package utils

import (
	"errors"
	"github.com/spf13/viper"
)

type Oauth struct {
	Providers []OauthProvider
}

type OauthProvider struct {
	Provider string
	Config   OauthProviderConfig
}

type OauthProviderConfig struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
	Scopes       []string
	AuthUrl      string
	TokenUrl     string
	UserInfoUrl  string
}

type OauthUserInformation struct {
	Id     string
	Name   string
	Email  string
	Avatar string
}

func NewOauth() *Oauth {
	return &Oauth{
		// Implemented Providers: google, facebook, github, discord, apple etc.
		Providers: []OauthProvider{
			{
				Provider: "google",
				Config: OauthProviderConfig{
					ClientId:     viper.GetString("OAUTH2_GOOGLE_CLIENT_ID"),
					ClientSecret: viper.GetString("OAUTH2_GOOGLE_CLIENT_SECRET"),
					RedirectUrl:  "http://localhost:8080/v1/authentication/sign_in_callback?provider=google",
					Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
					AuthUrl:      "https://accounts.google.com/o/oauth2/auth",
					TokenUrl:     "https://oauth2.googleapis.com/token",
					UserInfoUrl:  "https://www.googleapis.com/oauth2/v3/userinfo",
				},
			},
			{
				Provider: "facebook",
				Config: OauthProviderConfig{
					ClientId:     viper.GetString("OAUTH2_FACEBOOK_CLIENT_ID"),
					ClientSecret: viper.GetString("OAUTH2_FACEBOOK_CLIENT_SECRET"),
					RedirectUrl:  "http://localhost:8080/v1/authentication/sign_in_callback?provider=facebook",
					Scopes:       []string{"email", "public_profile"},
					AuthUrl:      "https://www.facebook.com/v3.2/dialog/oauth",
					TokenUrl:     "https://graph.facebook.com/v3.2/oauth/access_token",
					UserInfoUrl:  "https://graph.facebook.com/me",
				},
			},
			{
				Provider: "discord",
				Config: OauthProviderConfig{
					ClientId:     viper.GetString("OAUTH2_DISCORD_CLIENT_ID"),
					ClientSecret: viper.GetString("OAUTH2_DISCORD_CLIENT_SECRET"),
					RedirectUrl:  "http://localhost:8080/v1/authentication/sign_in_callback?provider=discord",
					Scopes:       []string{"identify", "email"},
					AuthUrl:      "https://discord.com/api/oauth2/authorize",
					TokenUrl:     "https://discord.com/api/oauth2/token",
					UserInfoUrl:  "https://discord.com/api/users/@me",
				},
			},
			{
				Provider: "github",
				Config: OauthProviderConfig{
					ClientId:     viper.GetString("OAUTH2_GITHUB_CLIENT_ID"),
					ClientSecret: viper.GetString("OAUTH2_GITHUB_CLIENT_SECRET"),
					RedirectUrl:  "http://localhost:8080/v1/authentication/sign_in_callback?provider=github",
					Scopes:       []string{"user:email"},
					AuthUrl:      "https://github.com/login/oauth/authorize",
					TokenUrl:     "https://github.com/login/oauth/access_token",
					UserInfoUrl:  "https://api.github.com/user",
				},
			},
			{
				Provider: "apple",
				Config: OauthProviderConfig{
					ClientId:     viper.GetString("OAUTH2_APPLE_CLIENT_ID"),
					ClientSecret: viper.GetString("OAUTH2_APPLE_CLIENT_SECRET"),
					RedirectUrl:  "http://localhost:8080/v1/authentication/sign_in_callback?provider=apple",
					Scopes:       []string{"name", "email"},
					AuthUrl:      "https://appleid.apple.com/auth/authorize",
					TokenUrl:     "https://appleid.apple.com/auth/token",
					UserInfoUrl:  "https://appleid.apple.com/auth/userinfo",
				},
			},
		},
	}
}

func (o *Oauth) GetProvider(provider string) (OauthProvider, error) {
	for _, p := range o.Providers {
		if p.Provider == provider {
			return p, nil
		}
	}
	return OauthProvider{}, errors.New("provider not found")
}

func (o *Oauth) GetUserInformation(provider string, userInformation map[string]interface{}) (OauthUserInformation, error) {
	switch provider {
	case "google":
		return OauthUserInformation{
			Id:     userInformation["sub"].(string),
			Name:   userInformation["name"].(string),
			Email:  userInformation["email"].(string),
			Avatar: userInformation["picture"].(string),
		}, nil
	case "facebook":
		return OauthUserInformation{
			Id:     userInformation["id"].(string),
			Name:   userInformation["name"].(string),
			Email:  userInformation["email"].(string),
			Avatar: "https://graph.facebook.com/" + userInformation["id"].(string) + "/picture?type=large",
		}, nil
	case "discord":
		return OauthUserInformation{
			Id:     userInformation["id"].(string),
			Name:   userInformation["username"].(string),
			Email:  userInformation["email"].(string),
			Avatar: "https://cdn.discordapp.com/avatars/" + userInformation["id"].(string) + "/" + userInformation["avatar"].(string) + ".png",
		}, nil
	case "github":
		return OauthUserInformation{
			Id:     userInformation["id"].(string),
			Name:   userInformation["name"].(string),
			Email:  userInformation["email"].(string),
			Avatar: userInformation["avatar_url"].(string),
		}, nil
	case "apple":
		return OauthUserInformation{
			Id:     userInformation["sub"].(string),
			Name:   userInformation["name"].(string),
			Email:  userInformation["email"].(string),
			Avatar: "",
		}, nil
	default:
		return OauthUserInformation{}, errors.New("provider not found")
	}
}
