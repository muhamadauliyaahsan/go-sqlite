package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ahsan/go-sqlite-crud/internal/model"
	"github.com/ahsan/go-sqlite-crud/internal/service"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with name, description and price
// @Tags products
// @Accept json
// @Produce json
// @Param product body model.Product true "Product object"
// @Success 201 {object} model.Product
// @Failure 400 {string} string "Bad Request"
// @Router /products [post]
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateProduct(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Get list of all products
// @Tags products
// @Produce json
// @Success 200 {array} model.Product
// @Router /products [get]
func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

// GetProductByID godoc
// @Summary Get a product by ID
// @Description Get product details by ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} model.Product
// @Failure 404 {string} string "Not Found"
// @Router /products/{id} [get]
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	product, err := h.service.GetProductByID(uint(id))
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update product details by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body model.Product true "Product object"
// @Success 200 {object} model.Product
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product.ID = uint(id)

	if err := h.service.UpdateProduct(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete product by ID
// @Tags products
// @Param id path int true "Product ID"
// @Success 204 "No Content"
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	if err := h.service.DeleteProduct(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
