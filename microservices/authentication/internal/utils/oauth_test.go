package utils

import (
	"testing"
)

func TestOauthGetProvider(t *testing.T) {
	oauthProviders := NewOauth()
	provider, err := oauthProviders.GetProvider("google")
	if err != nil {
		t.Error(err)
	}
	if provider.Provider != "google" {
		t.Error("Expected provider to be google")
	}

	provider, err = oauthProviders.GetProvider("facebook")
	if err != nil {
		t.Error(err)
	}

	if provider.Provider != "facebook" {
		t.Error("Expected provider to be facebook")
	}

	t.Log("TestOauthGetProvider passed")
}

func TestOauthGetUserInformation(t *testing.T) {
	oauthProviders := NewOauth()
	userInformation, err := oauthProviders.GetUserInformation("google", map[string]interface{}{
		"sub":     "123",
		"name":    "test",
		"email":   "test@hotmail.com",
		"picture": "avatar.png",
	})

	if err != nil {
		t.Error(err)
	}

	if userInformation.Id != "123" || userInformation.Name != "test" || userInformation.Email != "test@hotmail.com" || userInformation.Avatar != "avatar.png" {
		t.Error("Expected user information to be correct")
	}
	t.Log(userInformation)
	t.Log("TestOauthGetUserInformation passed")
}
