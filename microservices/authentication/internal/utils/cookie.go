package utils

import (
	"github.com/gorilla/securecookie"
	"net/http"
	"time"
)

var secureCookie = securecookie.New(securecookie.GenerateRandomKey(32), securecookie.GenerateRandomKey(32))

type ICookieStore interface {
	SetCookieAccessToken(wr http.ResponseWriter, accessToken string)
	SetCookieRefreshToken(wr http.ResponseWriter, refreshToken string)
	SetCookieCsrfToken(wr http.ResponseWriter, csrfToken string)
	SetCookieOAuth2State(wr http.ResponseWriter, state string)
	GetCookieDecrypted(req *http.Request, cookieName string) (string, error)
	DestroyCookie(wr http.ResponseWriter, cookieName string)
}

type CookieStore struct {
	AccessTokenCookieName  string
	RefreshTokenCookieName string
	CsrfTokenCookieName    string
	OauthStateCookieName   string
	DefaultExpiration      time.Time
}

func NewCookieStore() *CookieStore {
	return &CookieStore{
		AccessTokenCookieName:  "session.access_token",
		RefreshTokenCookieName: "session.refresh_token",
		CsrfTokenCookieName:    "session.csrf_token",
		OauthStateCookieName:   "session.oauth_state",
		DefaultExpiration:      time.Now().Add(24 * time.Hour),
	}
}

func (c *CookieStore) SetCookieAccessToken(wr http.ResponseWriter, accessToken string) {
	// TODO: encrypt access token
	encoded, err := secureCookie.Encode(c.AccessTokenCookieName, accessToken)
	if err != nil {
		panic(err)
	}
	cookie := http.Cookie{
		Name:     c.AccessTokenCookieName,
		Value:    encoded,
		Expires:  time.Now().Add(5 * time.Minute),
		HttpOnly: true,
	}
	http.SetCookie(wr, &cookie)
}

func (c *CookieStore) SetCookieRefreshToken(wr http.ResponseWriter, refreshToken string) {
	// TODO: encrypt refresh token
	encoded, err := secureCookie.Encode(c.RefreshTokenCookieName, refreshToken)
	if err != nil {
		panic(err)
	}
	cookie := http.Cookie{
		Name:     c.RefreshTokenCookieName,
		Value:    encoded,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(wr, &cookie)
}

func (c *CookieStore) SetCookieCsrfToken(wr http.ResponseWriter, csrfToken string) {
	// TODO: encrypt csrf token
	encoded, err := secureCookie.Encode(c.CsrfTokenCookieName, csrfToken)
	if err != nil {
		panic(err)
	}
	cookie := http.Cookie{
		Name:     c.CsrfTokenCookieName,
		Value:    encoded,
		Expires:  c.DefaultExpiration,
		HttpOnly: true,
	}
	http.SetCookie(wr, &cookie)
}

func (c *CookieStore) SetCookieOAuth2State(wr http.ResponseWriter, state string) {
	// TODO: encrypt state
	encoded, err := secureCookie.Encode(c.OauthStateCookieName, state)
	if err != nil {
		panic(err)
	}
	cookie := http.Cookie{
		Name:     c.OauthStateCookieName,
		Value:    encoded,
		Expires:  time.Now().Add(5 * time.Minute),
		HttpOnly: true,
	}
	http.SetCookie(wr, &cookie)
}

func (c *CookieStore) GetCookieDecrypted(req *http.Request, cookieName string) (string, error) {
	cookie, err := req.Cookie(cookieName)
	if err != nil {
		return "", err
	}
	// TODO: decrypt cookie value
	var value string
	err = secureCookie.Decode(cookieName, cookie.Value, &value)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (c *CookieStore) DestroyCookie(wr http.ResponseWriter, cookieName string) {
	cookie := http.Cookie{
		Name:     cookieName,
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}
	http.SetCookie(wr, &cookie)
}
