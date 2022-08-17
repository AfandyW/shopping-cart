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
	router.HandleFunc("/products", c.Create).Methods(http.MethodPost)
	router.HandleFunc("/products", c.List).Methods(http.MethodGet)
	router.HandleFunc("/products/{productCode}", c.Update).Methods(http.MethodPut)
	router.HandleFunc("/products/{productCode}", c.Get).Methods(http.MethodGet)
	router.HandleFunc("/products/{productCode}", c.Delete).Methods(http.MethodDelete)

	return router
}
