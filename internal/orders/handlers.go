package orders

import (
	"database/sql"
	"errors"
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

func (h *handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	requestData := PlaceOrderBodyParams{}
	if json.Read(r, w, &requestData) != nil {
		return
	}

	orderID, err := h.service.PlaceOrder(
		r.Context(),
		requestData.CustomerID,
		requestData.Items,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Out of Stock", http.StatusBadRequest)
			return
		}

		http.Error(
			w,
			"Failed to place order: "+err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	json.Write(w, http.StatusOK, map[string]int64{"order_id": orderID})
}

func (h *handler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	orderIDQuery := chi.URLParam(r, "id")
	orderID, err := strconv.ParseInt(orderIDQuery, 10, 64)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	order, err := h.service.GetOrderByID(r.Context(), orderID)
	if err != nil {
		http.Error(
			w,
			"Failed to get order: "+err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	json.Write(w, http.StatusOK, order)
}

func (h *handler) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.service.ListOrders(r.Context())
	if err != nil {
		http.Error(
			w,
			"Failed to list orders: "+err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	json.Write(w, http.StatusOK, orders)
}
