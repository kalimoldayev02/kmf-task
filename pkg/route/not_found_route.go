package route

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Not found Route
func NotFoundRoute(r *mux.Router) {
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		response := map[string]string{"message": "page not found"}

		json.NewEncoder(w).Encode(response)
	})
}
