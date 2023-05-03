package entities

import (
	"time"
)

const UserTableName = "users"

type User struct {
	ID         string    `json:"id,omitempty"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Profile    Profile   `json:"profile"`
	Verified   bool      `json:"verified"`
	Disabled   bool      `json:"disabled"`
	Mfa        bool      `json:"mfa"`
	LastSeenAt time.Time `json:"lastSeen"`
	Phone      string    `json:"phone"`
	Roles      []Role    `json:"roles"`
	Resets     []Reset   `json:"resets"`
	Sessions   []Session `json:"sessions"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type Profile struct {
	Name   string `json:"name"`   // Full name
	Avatar string `json:"avatar"` // Avatar url
}

type Role struct {
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
}

type Permission struct {
	Scope  string `json:"scope"`
	Action string `json:"action"`
}

type Reset struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type Session struct {
	Type         string    `json:"type"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	ExpiresAt    time.Time `json:"expiresAt"`
	Blocked      bool      `json:"blocked"`
}
