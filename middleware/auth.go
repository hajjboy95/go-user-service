package middleware

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/hajjboy95/go-user-service/models"
	"github.com/hajjboy95/go-user-service/utils"
	"net/http"
	"os"
	"strings"
)

type AuthenticationMiddleware struct {
	safePaths []string
}

func (as *AuthenticationMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{"/v1/api/user/new", "/v1/api/user/login"}
		requestPath := r.URL.Path

		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string] interface{})
		tokenHeader := r.Header.Get("x-token-header")

		if tokenHeader == "" {
			response = utils.Message(false, "Missing middleware Token")
			utils.Respond(http.StatusForbidden, w, response)
			return
		}

		splitTokens := strings.Split(tokenHeader, " ")

		if len(splitTokens) != 2 {
			response = utils.Message(false, "Invalid/Malformed middleware token")
			utils.Respond(http.StatusForbidden, w, response)
			return
		}

		jwtToken := splitTokens[1]
		tokenModel := &models.Token{}

		token, err := jwt.ParseWithClaims(jwtToken, tokenModel, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			response = utils.Message(false, "Malformed Auth Token")
			utils.Respond(http.StatusForbidden, w, response)
			return
		}

		if !token.Valid {
			response = utils.Message(false, "Invalid Token")
			utils.Respond(http.StatusForbidden, w, response)
		}

		fmt.Printf("User %d", tokenModel.UserId)
		ctx := context.WithValue(r.Context(), "user", tokenModel.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed along the middleware chain
	})
}

func NewAuthenticationMiddleware(safePaths []string) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		safePaths: safePaths,
	}
}