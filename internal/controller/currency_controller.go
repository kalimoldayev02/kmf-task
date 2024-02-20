package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kalimoldayev02/kmf-task/pkg/utils"
)

func (c *Controller) SaveCurrency(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date := vars["date"]

	if err := c.Validator.Var(date, "required,currency_date"); err != nil {
		responseField := map[string]string{}
		responseField["message"] = "invalid format date"

		http.Error(w, utils.ToJson(responseField), http.StatusBadRequest)
		return
	}

	fmt.Println(c.Service.Currency.Save(date))

	// TODO: call method from service
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Date: " + date))
}
