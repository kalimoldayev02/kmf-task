package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kalimoldayev02/kmf-task/pkg/utils"
)

func newCurrencyHandler(r *mux.Router, h *Handler) {
	r.HandleFunc("/save/{date}", h.createCurrency).Methods(http.MethodGet)
}

func (h *Handler) createCurrency(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date := vars["date"]

	if err := h.validator.Var(date, "required,currency_date"); err != nil {
		responseField := map[string]string{}
		responseField["message"] = "invalid format date"

		http.Error(w, utils.ToJson(responseField), http.StatusBadRequest)
		return
	}

	result := h.service.Currency.Create(date)
	data := map[string]bool{"status": result}
	response, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !result {
		http.Error(w, string(response), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
