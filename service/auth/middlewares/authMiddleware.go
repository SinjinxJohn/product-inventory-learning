package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"example.com/event-app/service/auth"
	"example.com/event-app/utils"
)

type contextKey string

const UserIDKey contextKey = "userID"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) == 0 {
			utils.WriteError(w, http.StatusUnauthorized,
				fmt.Errorf("authorization header missing, cannot access route"))
			return
		}

		// authHeaderStr =

		tokenStr, err := utils.ValidateTokenFormat(authHeader)
		if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, err)
			return
		}

		userID, err := auth.ParseAndValidateToken(tokenStr)
		if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, err)
			return
		}

		// Add userID to request context
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
