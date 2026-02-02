package products

import (
	"log/slog"
	"net/http"

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

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request){
	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, products)
}
