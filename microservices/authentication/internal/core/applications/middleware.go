package applications

import (
	"context"
	"github.com/Lenstack/lensaas-server/microservices/authentication/internal/utils"
	"github.com/gorilla/csrf"
	"github.com/gorilla/securecookie"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"net/http"
	"strings"
	"time"
)

func (m *Microservice) MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		// Create cookie store
		cookieStore := utils.NewCookieStore()

		// Get access token from cookie
		accessToken, err := cookieStore.GetCookieDecrypted(req, cookieStore.AccessTokenCookieName)
		if err != nil {
			m.Log.Error("Error getting access token from cookie", zap.Error(err))
			wr.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Validate access token and get userId
		userId, err := m.Jwt.ValidateToken(accessToken)
		if err != nil {
			m.Log.Error("Unauthorized", zap.Error(err))
			wr.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Set the userId in the request context
		ctx := req.Context()
		ctx = context.WithValue(ctx, "userId", userId)
		req = req.WithContext(ctx)

		next.ServeHTTP(wr, req)
	})
}

func (m *Microservice) MiddlewarePermission(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		next.ServeHTTP(w, r)
	})
}

func (m *Microservice) MiddlewareLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a new ResponseWriter that wraps the original
		// and intercepts the status code
		responseWriter := &responseWriter{w, http.StatusOK}

		// Call the next middleware or handler with the wrapped ResponseWriter
		next.ServeHTTP(responseWriter, r)

		// Log information about the incoming request and the response
		m.Log.Info("Incoming request",
			zap.String("Method", r.Method),
			zap.String("URL", r.URL.Path),
			zap.String("RemoteAddr", r.RemoteAddr),
			zap.String("UserAgent", r.UserAgent()),
		)

		// If the response code is in the 200 range, log an info
		if responseWriter.statusCode >= 200 && responseWriter.statusCode < 300 {
			m.Log.Info("Successful request",
				zap.String("Method", r.Method),
				zap.String("URL", r.URL.Path),
				zap.String("RemoteAddr", r.RemoteAddr),
				zap.String("UserAgent", r.UserAgent()),
				zap.Int("StatusCode", responseWriter.statusCode),
			)
		}

		// If the response code is in the 400 range, log a warning
		if responseWriter.statusCode >= 400 && responseWriter.statusCode < 500 {
			m.Log.Warn("Unsuccessful request",
				zap.String("Method", r.Method),
				zap.String("URL", r.URL.Path),
				zap.String("RemoteAddr", r.RemoteAddr),
				zap.String("UserAgent", r.UserAgent()),
				zap.Int("StatusCode", responseWriter.statusCode),
			)
		}

		// If the response code is in the 500 range, log an error
		if responseWriter.statusCode >= 500 {
			m.Log.Error("Internal server error",
				zap.String("Method", r.Method),
				zap.String("URL", r.URL.Path),
				zap.String("RemoteAddr", r.RemoteAddr),
				zap.String("UserAgent", r.UserAgent()),
				zap.Int("StatusCode", responseWriter.statusCode),
			)
		}
	})
}

// Define a custom ResponseWriter that wraps the original and intercepts the status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (m *Microservice) MiddlewareCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "application/json")
		// Stop here for a Preflighted OPTIONS request.
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (m *Microservice) MiddlewareRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (m *Microservice) MiddlewareRateLimit(next http.Handler) http.Handler {
	// Create a rate limiter that allows up to 10 requests per second
	limiter := rate.NewLimiter(rate.Every(time.Second), 10)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the rate limiter allows the request
		if !limiter.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		// Call the next middleware or handler
		next.ServeHTTP(w, r)
	})
}

func (m *Microservice) MiddlewareCsrf(next http.Handler) http.Handler {
	return csrf.Protect(securecookie.GenerateRandomKey(32),
		csrf.Secure(false), csrf.CookieName("session.csrf_token"),
		csrf.RequestHeader("X-CSRF-Token"),
		csrf.TrustedOrigins([]string{"*"}))(next)
}

func (m *Microservice) MiddlewareCleanPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" && r.URL.Path[len(r.URL.Path)-1] == '/' {
			http.Redirect(w, r, strings.TrimSuffix(r.URL.Path, "/"), http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (m *Microservice) MiddlewareAllowContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost || r.Method == http.MethodPut {
			if r.Header.Get("Content-Type") != "application/json" {
				http.Error(w, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
