package app

import (
	"net/http"

	"github.com/AfandyW/shopping-cart/controller"
	"github.com/gorilla/mux"
)

func NewRouter(c controller.Controller) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("ping sukses"))
	})
	r := router.PathPrefix("/api/v1").Subrouter()
	r.HandleFunc("/products", c.Create).Methods(http.MethodPost)
	r.HandleFunc("/products", c.List).Methods(http.MethodGet)
	r.HandleFunc("/products/{productCode}", c.Update).Methods(http.MethodPut)
	r.HandleFunc("/products/{productCode}", c.Get).Methods(http.MethodGet)
	r.HandleFunc("/products/{productCode}", c.Delete).Methods(http.MethodDelete)

	return router
}
