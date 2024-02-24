package http

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) NotFoundRoute() {
	h.router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		response := map[string]string{"message": "page not found"}

		json.NewEncoder(w).Encode(response)
	})
}
