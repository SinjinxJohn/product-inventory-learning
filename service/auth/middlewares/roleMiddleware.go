package middlewares

import (
	"net/http"

	"example.com/event-app/types"
	"example.com/event-app/utils"
	"github.com/gorilla/mux"
)

func RoleMiddleware(store types.UserStore, role string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userId := r.Context().Value("userID").(int)
			user, err := store.GetUserByID(userId)
			if err != nil {
				utils.WriteError(w, http.StatusUnauthorized, err)
				return
			}

			if user.Role != role {
				utils.WriteError(w, http.StatusUnauthorized, err)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
