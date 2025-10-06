package categories

import (
	"net/http"

	"example.com/event-app/types"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.CategoryStore
}

func NewHandler(store types.CategoryStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRouters(router *mux.Router) {

	router.HandleFunc("/create-category", h.createCategory).Methods("POST")
	router.HandleFunc("/get-all-categories", h.getAllCategories).Methods("GET")

}

func (h *Handler) createCategory(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getAllCategories(w http.ResponseWriter, r *http.Request) {

}
