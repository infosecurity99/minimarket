package handler

import (
	"connected/api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"errors"
	"net/http"
	"strconv"
)

// CreateProduct godoc
// @Router       /product [POST]
// @Summary      Creates a new product
// @Description  create a new product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        basket body models.CreateProduct false "product"
// @Success      201  {object}  models.Product
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
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

// GetByIDProduct retrieves product information by ID.
// @Summary Get product by ID
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /product/{id} [get]
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

// GetListProduct returns a list of products.
// @Summary Get a list of products
// @Tags product
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} models.Product
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /products [get]
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

// UpdateProduct godoc
// @Router       /product/{id} [PUT]
// @Summary      Update product
// @Description  update product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 id path string true "product"
// @Param        user body models.UpdateProduct true "product"
// @Success      200  {object}  models.Product
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
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

// DeleteProduct deletes product by ID.
// @Summary Delete product by ID
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {string} string "Data successfully deleted"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /product/{id} [delete]
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

// BarcodeQrCode godoc
// @Router       /basket/{id} [GET]
// @Summary      Gets basket
// @Description  get basket by ID
// @Tags         basket
// @Accept       json
// @Produce      json
// @Param        id path string true "basket"
// @Success      200  {object}  models.Basket
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) BarcodeQrCode(c *gin.Context) {
       
}
