package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kalimoldayev02/kmf-task/internal/controller"
)

func newCurrencyRoutes(r *mux.Router, c *controller.Controller) {
	// сurrency Роуты
	currencyApi := r.PathPrefix("/currency").Subrouter()
	currencyApi.HandleFunc("/save/{date}", c.SaveCurrency).Methods(http.MethodGet)
}
