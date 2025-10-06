package api

import (
	"database/sql"
	"log"
	"net/http"

	"example.com/event-app/service/auth/middlewares"
	"example.com/event-app/service/categories"
	"example.com/event-app/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRouters(subrouter)

	categoryStore := categories.NewStore(s.db)
	categoryHandler := categories.NewHandler(categoryStore)

	adminRouter := router.PathPrefix("/api/v1/admin").Subrouter()
	adminRouter.Use(middlewares.AuthMiddleware)

	adminRouter.Use(middlewares.RoleMiddleware(userStore, "admin"))
	categoryHandler.RegisterRouters(adminRouter)

	log.Println("Starting API server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
