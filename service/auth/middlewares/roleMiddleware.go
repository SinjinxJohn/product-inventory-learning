package middlewares

import (
	"fmt"
	"net/http"

	"example.com/event-app/types"
	"example.com/event-app/utils"
	"github.com/gorilla/mux"
)

func RoleMiddleware(store types.UserStore, role string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			val := r.Context().Value(UserIDKey)
			userId, ok := val.(int)
			if !ok {
				utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("userID missing in context"))
				return
			}

			user, err := store.GetUserByID(userId)
			if err != nil {
				utils.WriteError(w, http.StatusUnauthorized, err)
				return
			}

			if user.Role != role {
				utils.WriteError(w, http.StatusForbidden, fmt.Errorf("unauthorized role"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
