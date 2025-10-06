package user

import (
	"fmt"
	"net/http"

	"example.com/event-app/config"
	"example.com/event-app/service/auth"
	"example.com/event-app/service/auth/middlewares"

	"example.com/event-app/types"
	"example.com/event-app/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRouters(router *mux.Router) {
	router.HandleFunc("/login", h.login).Methods("POST")
	router.HandleFunc("/register", h.registerUser).Methods("POST")
	router.Handle("/user-profile", middlewares.AuthMiddleware(http.HandlerFunc(h.getProfile))).Methods("GET")

}

func (h *Handler) getProfile(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middlewares.UserIDKey).(int)
	user, err := h.store.GetUserByID(userId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, user)
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	var loginPayload types.LoginUser
	if err := utils.ParseJson(r, &loginPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	//check if user exits
	user, err := h.store.GetUserByEmail(loginPayload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with this email does not exist"))
		return

	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, user.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if !auth.CompareHashedPasswords(user.Password, []byte(loginPayload.Password)) {
		fmt.Println("Hashed password from DB:", user.Password)
		fmt.Println("Login password:", loginPayload.Password)

		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{"token": token, "message": "login successful"})
}
func (h *Handler) registerUser(w http.ResponseWriter, r *http.Request) {
	//get json payload
	var registerPayload types.RegisterUser
	if err := utils.ParseJson(r, &registerPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//check if user exits
	_, err := h.store.GetUserByEmail(registerPayload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with this email"))
		return

	}
	hashedPassword, err := auth.HashPassword(registerPayload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
	//if not create a new user
	err = h.store.CreateUser(&types.RegisterUser{
		Email:     registerPayload.Email,
		FirstName: registerPayload.FirstName,
		LastName:  registerPayload.LastName,
		Role:      registerPayload.Role,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, nil)

}
