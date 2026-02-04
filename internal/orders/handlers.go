package orders

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/sorinqu-org/go-backend-api/internal/json"
)

type PlaceOrderBodyParams struct {
	CustomerID int64       `json:"customer_id"`
	Items      []OrderItem `json:"items"`
}

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
