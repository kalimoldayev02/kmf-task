package route

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kalimoldayev02/kmf-task/app/controller"
)

func PublicRoutes(r *mux.Router, c *controller.Controller) {
	// Currency Routes
	currencyApi := r.PathPrefix("/currency").Subrouter()
	currencyApi.HandleFunc("/save/{date}", c.SaveCurrency).Methods(http.MethodGet)
}
