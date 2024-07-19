package http

import (
	"context"
	"encoding/json"
	"net/http"
	"orderService/internal/delivery/http/request"
	"orderService/internal/delivery/http/response"
	"orderService/internal/domain"
	"orderService/internal/usecase"
)

type OrderHandler struct {
	UseCase usecase.OrderUseCase
}

func (h *OrderHandler) SendOrder(w http.ResponseWriter, r *http.Request) {

	requestBody := &request.Request{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&response.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	data := domain.New(requestBody)
	if err := h.UseCase.CreateOrder(context.Background(), data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response.Response{
		Status:  http.StatusOK,
		Message: "OK",
	})
	return

}
