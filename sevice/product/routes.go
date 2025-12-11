package product

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shalshcode08/GO-APIs/types"
	"github.com/shalshcode08/GO-APIs/utils"
)

type Handler struct {
	store types.ProductStorte
}

func NewHandler(store types.ProductStorte) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodGet)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ps)
}
