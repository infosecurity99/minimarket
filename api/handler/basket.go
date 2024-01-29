package handler

import (
	"errors"
	"net/http"
	"strconv"

	"connected/api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateBasket godoc
// @Router       /basket [POST]
// @Summary      Creates a new basket
// @Description  create a new basket
// @Tags         basket
// @Accept       json
// @Produce      json
// @Param        basket body models.CreateBasket false "basket"
// @Success      201  {object}  models.Basket
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateBasket(c *gin.Context) {
	createBasket := models.CreateBasket{}

	if err := c.ShouldBindJSON(&createBasket); err != nil {
		handleResponse(c, "Error: Failed to parse request body JSON", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Basket().CreateBasket(createBasket)
	if err != nil {
		handleResponse(c, "Error: Failed to create basket", http.StatusInternalServerError, err)
		return
	}

	basket, err := h.storage.Basket().GetByIdBasket(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error: Failed to find basket by ID", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, basket)
}

// GetByIDBasket godoc
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
// @Router /basket/{id} [get]
func (h Handler) GetByIDBasket(c *gin.Context) {
	var err error

	uid := c.Param("id")

	basket, err := h.storage.Basket().GetByIdBasket(models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "Error: Failed to find basket by ID", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, basket)
}

// GetListBasket gets a list of baskets.
// @Summary Get a list of baskets
// @Tags basket
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Param search query string false "Search term"
// @Success 200 {object} models.BasketResponse
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /baskets [get]
func (h Handler) GetListBasket(c *gin.Context) {
	var (
		page, limit int
		search      string
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, "Error: Failed to parse page parameter", http.StatusBadRequest, err)
		return
	}

	if page < 1 {
		handleResponse(c, "Error: Invalid page number", http.StatusBadRequest, nil)
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, "Error: Failed to parse limit parameter", http.StatusBadRequest, err)
		return
	}

	if limit < 1 {
		handleResponse(c, "Error: Invalid limit value", http.StatusBadRequest, nil)
		return
	}

	search = c.Query("search")

	resp, err := h.storage.Basket().GetListBasket(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "Error: Failed to get basket list", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, resp)
}

// UpdateBasket godoc
// @Router       /basket/{id} [PUT]
// @Summary      Update basket
// @Description  update basket
// @Tags         basket
// @Accept       json
// @Produce      json
// @Param 		 id path string true "basket"
// @Param        user body models.UpdateBasket true "basket"
// @Success      200  {object}  models.Basket
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateBasket(c *gin.Context) {
	updateBasket := models.UpdateBasket{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "Error: Invalid UUID", http.StatusBadRequest, errors.New("UUID is not valid"))
		return
	}

	updateBasket.ID = uid

	if err := c.ShouldBindJSON(&updateBasket); err != nil {
		handleResponse(c, "Error: Failed to parse request body JSON", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Basket().UpdateBasket(updateBasket)
	if err != nil {
		handleResponse(c, "Error: Failed to update basket", http.StatusInternalServerError, err)
		return
	}

	basket, err := h.storage.Basket().GetByIdBasket(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error: Failed to find basket by ID", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, basket)
}

// DeleteBasket deletes basket by ID.
// @Summary Delete basket by ID
// @Tags basket
// @Accept json
// @Produce json
// @Param id path string true "Basket ID"
// @Success 200 {string} string "Data deleted successfully"
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /basket/{id} [delete]
func (h Handler) DeleteBasket(c *gin.Context) {
	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "Error: Invalid UUID", http.StatusBadRequest, err)
		return
	}

	if err = h.storage.Basket().DeleteBasket(models.PrimaryKey{
		ID: id.String(),
	}); err != nil {
		handleResponse(c, "Error: Failed to delete basket by ID", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, "Data deleted successfully")
}
