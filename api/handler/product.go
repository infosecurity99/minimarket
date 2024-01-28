package handler

import (
	"connected/api/models"
	"github.com/google/uuid"

	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateProduct(c *gin.Context) {
	createProduct := models.CreateProduct{}

	if err := c.ShouldBindJSON(&createProduct); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Product().CreateProduct(createProduct)
	if err != nil {
		handleResponse(c, "error while creating product", http.StatusInternalServerError, err)
		return
	}

	product, err := h.storage.Product().GetByIdProduct(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting product by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, product)
}

func (h Handler) GetByIDProduct(c *gin.Context) {
	var err error

	uid := c.Param("id")

	product, err := h.storage.Product().GetByIdProduct(models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "error while getting product by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, product)
}

func (h Handler) GetListProduct(c *gin.Context) {
	var (
		page, limit int
		search      string
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)

	if err != nil {
		handleResponse(c, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	search = c.Query("search")

	resp, err := h.storage.Product().GetListProduct(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "error while getting product", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, resp)
}

func (h Handler) UpdateProduct(c *gin.Context) {
	updateProduct := models.UpdateProduct{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("invalid uuid"))
		return
	}

	updateProduct.ID = uid

	if err := c.ShouldBindJSON(&updateProduct); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Product().UpdateProduct(updateProduct)
	if err != nil {
		handleResponse(c, "error while updating product", http.StatusInternalServerError, err)
		return
	}

	product, err := h.storage.Product().GetByIdProduct(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting product by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, product)
}

func (h Handler) DeleteProduct(c *gin.Context) {
	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err = h.storage.Product().DeleteProduct(models.PrimaryKey{
		ID: id.String(),
	}); err != nil {
		handleResponse(c, "error while deleting product by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data successfully deleted")
}
