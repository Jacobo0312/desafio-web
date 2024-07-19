package handler

import (
	"app/internal/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type TicketHandler struct {
	service service.ServiceTicketDefault
}

func NewTicketHandler(service service.ServiceTicketDefault) *TicketHandler {
	return &TicketHandler{
		service: service,
	}
}

func (h *TicketHandler) GetTotalTickets(w http.ResponseWriter, r *http.Request) {
	total, err := h.service.GetTotalTickets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(total)
}

func (h *TicketHandler) GetTotalAmountTickets(w http.ResponseWriter, r *http.Request) {
	total, err := h.service.GetTotalAmountTickets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(total)
}

func (h *TicketHandler) GetAverageByCountry(w http.ResponseWriter, r *http.Request) {
	destination := chi.URLParam(r, "dest")
	average, err := h.service.AverageDestination(destination)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(average)
}

func (h *TicketHandler) GetByCountry(w http.ResponseWriter, r *http.Request) {
	country := chi.URLParam(r, "dest")
	total, err := h.service.GetTotalAmountTicketsByCountry(country)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(total)
}
