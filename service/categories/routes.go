package categories

import (
	"fmt"
	"net/http"

	"example.com/event-app/types"
	"example.com/event-app/utils"
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
	var createCategoryPayload types.CreateCategoryPayload

	if err := utils.ParseJson(r, &createCategoryPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err := h.store.GetCategoryByName(createCategoryPayload.Name)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("category already exists"))
	}

	err = h.store.CreateCategory(&types.CreateCategoryPayload{
		Name: createCategoryPayload.Name,
	})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]string{"message": "category created successfully"})

}

func (h *Handler) getAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.store.GetAllCategories()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, categories)
}
