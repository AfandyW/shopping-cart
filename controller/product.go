package controller

import (
	"encoding/json"
	"net/http"

	"github.com/AfandyW/shopping-cart/domain"
	"github.com/gorilla/mux"
)

type Controller struct {
	service domain.Service
}

func NewController(s domain.Service) *Controller {
	return &Controller{
		service: s,
	}
}

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func send(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	resp := response{
		Code:    code,
		Message: message,
		Data:    data,
	}

	w.WriteHeader(code)
	byData, _ := json.Marshal(resp)
	w.Write(byData)
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.ProductRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		send(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	if req.ProductName == "" {
		send(w, http.StatusBadRequest, "Nama Produk tidak boleh kosong", nil)
		return
	}

	if req.Quantity < 0 {
		send(w, http.StatusBadRequest, "Kuantitas Produk minimal 1", nil)
		return
	}

	err = c.service.Create(r.Context(), req)
	if err != nil {
		send(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	send(w, http.StatusCreated, "success create new products", nil)
}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	var req domain.ProductRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		send(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	params := mux.Vars(r)
	productCode := params["productCode"]

	if req.Quantity < 0 {
		send(w, http.StatusBadRequest, "Kuantitas Produk minimal 1", nil)
		return
	}

	req.ProductCode = productCode

	err = c.service.Update(r.Context(), req)
	if err != nil {
		send(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	send(w, http.StatusCreated, "success update quantity products", nil)
}

func (c *Controller) List(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	productName := query.Get("product_name")

	products, err := c.service.List(r.Context(), domain.Filter{
		ProductName: productName,
	})
	if err != nil {
		send(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	send(w, http.StatusCreated, "success get list product", products)
}

func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productCode := params["productCode"]
	product, err := c.service.Get(r.Context(), productCode)
	if err != nil {
		send(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	send(w, http.StatusCreated, "success get product", product)
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productCode := params["productCode"]
	err := c.service.Delete(r.Context(), productCode)
	if err != nil {
		send(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	send(w, http.StatusCreated, "success delete product", "")
}
