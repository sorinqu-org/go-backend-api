package products

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sorinqu-org/go-backend-api/internal/json"
)

type handler struct {
	service service
}

func NewHandler(s service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, products)
}

func (h *handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productIDQuery := chi.URLParam(r, "id")

	productID, err := strconv.ParseInt(productIDQuery, 10, 64)
	if err != nil {
		http.Error(w, "Invalid query product ID", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetProduct(r.Context(), productID)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	json.Write(w, http.StatusOK, product)
}
