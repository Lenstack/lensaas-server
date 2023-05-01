package utils

import (
	"net/http/httptest"
	"testing"
)

func TestSetCookieAccessToken(t *testing.T) {
	c := NewCookieStore()

	rr := httptest.NewRecorder()
	accessToken := "testAccessToken"

	c.SetCookieAccessToken(rr, accessToken)

	if len(rr.Header().Get("Set-Cookie")) == 0 {
		t.Errorf("Cookie not set")
	}

	cookies := rr.Result().Cookies()
	if len(cookies) != 1 {
		t.Errorf("Expected 1 cookie, got %d", len(cookies))
	}

	cookie := cookies[0]
	if cookie.Name != c.AccessTokenCookieName {
		t.Errorf("Expected cookie name %q, got %q", c.AccessTokenCookieName, cookie.Name)
	}

	if cookie.Value == "" {
		t.Errorf("Cookie value is empty")
	}

	t.Log(cookie.Value)
}

func TestSetCookieRefreshToken(t *testing.T) {
	c := NewCookieStore()

	rr := httptest.NewRecorder()
	refreshToken := "testRefreshToken"

	c.SetCookieRefreshToken(rr, refreshToken)

	if len(rr.Header().Get("Set-Cookie")) == 0 {
		t.Errorf("Cookie not set")
	}

	cookies := rr.Result().Cookies()
	if len(cookies) != 1 {
		t.Errorf("Expected 1 cookie, got %d", len(cookies))
	}

	cookie := cookies[0]
	if cookie.Name != c.RefreshTokenCookieName {
		t.Errorf("Expected cookie name %q, got %q", c.RefreshTokenCookieName, cookie.Name)
	}

	if cookie.Value == "" {
		t.Errorf("Cookie value is empty")
	}

	t.Log(cookie.Value)
}

func TestSetCookieCsrfToken(t *testing.T) {
	c := NewCookieStore()

	rr := httptest.NewRecorder()
	csrfToken := "testCsrfToken"

	c.SetCookieCsrfToken(rr, csrfToken)

	if len(rr.Header().Get("Set-Cookie")) == 0 {
		t.Errorf("Cookie not set")
	}

	cookies := rr.Result().Cookies()
	if len(cookies) != 1 {
		t.Errorf("Expected 1 cookie, got %d", len(cookies))
	}

	cookie := cookies[0]
	if cookie.Name != c.CsrfTokenCookieName {
		t.Errorf("Expected cookie name %q, got %q", c.CsrfTokenCookieName, cookie.Name)
	}

	if cookie.Value == "" {
		t.Errorf("Cookie value is empty")
	}

	t.Log(cookie.Value)

}

func TestSetCookieOAuthState(t *testing.T) {
	c := NewCookieStore()

	rr := httptest.NewRecorder()
	oauthState := "testOAuthState"

	c.SetCookieOAuth2State(rr, oauthState)

	if len(rr.Header().Get("Set-Cookie")) == 0 {
		t.Errorf("Cookie not set")
	}

	cookies := rr.Result().Cookies()
	if len(cookies) != 1 {
		t.Errorf("Expected 1 cookie, got %d", len(cookies))
	}

	cookie := cookies[0]
	if cookie.Name != c.OauthStateCookieName {
		t.Errorf("Expected cookie name %q, got %q", c.OauthStateCookieName, cookie.Name)
	}

	if cookie.Value == "" {
		t.Errorf("Cookie value is empty")
	}

	t.Log(cookie.Value)

}

func TestGetCookieDecrypted(t *testing.T) {
	c := NewCookieStore()

	rr := httptest.NewRecorder()
	csrfToken := "testCsrfToken"

	c.SetCookieCsrfToken(rr, csrfToken)

	if len(rr.Header().Get("Set-Cookie")) == 0 {
		t.Errorf("Cookie not set")
	}

	cookies := rr.Result().Cookies()
	if len(cookies) != 1 {
		t.Errorf("Expected 1 cookie, got %d", len(cookies))
	}

	cookie := cookies[0]
	if cookie.Name != c.CsrfTokenCookieName {
		t.Errorf("Expected cookie name %q, got %q", c.CsrfTokenCookieName, cookie.Name)
	}

	decryptedCookie, err := c.GetCookieDecrypted(rr.Result().Request, cookie.Name)
	if err != nil {
		t.Error(err)
	}

	if decryptedCookie != csrfToken {
		t.Errorf("Expected cookie value %q, got %q", csrfToken, decryptedCookie)
	}

	t.Log(decryptedCookie)
}
