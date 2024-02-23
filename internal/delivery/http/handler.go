package http

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/kalimoldayev02/kmf-task/internal/service"
)

type Handler struct {
	service   *service.Service
	validator *validator.Validate
	router    *mux.Router
}

func NewHandler(s *service.Service, v *validator.Validate) *Handler {
	r := mux.NewRouter()
	r.Use(jsonResponse)

	return &Handler{
		service:   s,
		validator: v,
		router:    r,
	}
}

func (h *Handler) PublicHandler() {
	currencyApi := h.router.PathPrefix("/currency").Subrouter()
	{
		newCurrencyHandler(currencyApi, h)
	}
}

func jsonResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) GetRouter() *mux.Router {
	return h.router
}
