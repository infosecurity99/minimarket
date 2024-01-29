package handler

import (
	"errors"
	"net/http"
	"strconv"

	"connected/api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateSale godoc
// @Router       /sale [POST]
// @Summary      Creates a new sale
// @Description  create a new sale
// @Tags         sale
// @Accept       json
// @Produce      json
// @Param        basket body models.CreateSale false "sale"
// @Success      201  {object}  models.Sale
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateSale(c *gin.Context) {
	createSale := models.CreateSale{}

	if err := c.ShouldBindJSON(&createSale); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Sale().CreateSales(createSale)
	if err != nil {
		handleResponse(c, "error while creating sale", http.StatusInternalServerError, err)
		return
	}

	sale, err := h.storage.Sale().GetByIdSales(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting sale by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, sale)
}

// GetByIDSales retrieves sale information by ID.
// @Summary Get sale by ID
// @Tags sale
// @Accept json
// @Produce json
// @Param id path string true "Sale ID"
// @Success 200 {object} models.Sale
// @Failure 400 {string} models.Response
// @Failure 500 {string} models.Response
// @Router /sale/{id} [get]
func (h Handler) GetByIDSales(c *gin.Context) {
	var err error

	uid := c.Param("id")

	sale, err := h.storage.Sale().GetByIdSales(models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "error while getting sale by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, sale)
}

// GetListSales returns a list of sales.
// @Summary Get a list of sales
// @Tags sale
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} models.Sale
// @Failure 400 {string} models.Response
// @Failure 500 {string} models.Response
// @Router /sales [get]
func (h Handler) GetListSales(c *gin.Context) {
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

	resp, err := h.storage.Sale().GetListSales(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "error while getting sale", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, resp)
}

// UpdateSales godoc
// @Router       /sale/{id} [PUT]
// @Summary      Update sale
// @Description  update sale
// @Tags         sale
// @Accept       json
// @Produce      json
// @Param 		 id path string true "sale"
// @Param        user body models.UpdateSale true "sale"
// @Success      200  {object}  models.Sale
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateSales(c *gin.Context) {
	updateSale := models.UpdateSale{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	updateSale.ID = uid

	if err := c.ShouldBindJSON(&updateSale); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.Sale().UpdateSales(updateSale)
	if err != nil {
		handleResponse(c, "error while updating sale", http.StatusInternalServerError, err.Error())
		return
	}

	sales, err := h.storage.Sale().GetByIdSales(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting sale by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, sales)
}

// DeleteSales deletes sale by ID.
// @Summary Delete sale by ID
// @Tags sale
// @Accept json
// @Produce json
// @Param id path string true "Sale ID"
// @Success 200 {string} string models.Response
// @Failure 400 {string} string models.Response
// @Failure 500 {string} string models.Response
// @Router /sale/{id} [delete]
func (h Handler) DeleteSales(c *gin.Context) {
	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err = h.storage.Sale().DeleteSales(models.PrimaryKey{
		ID: id.String(),
	}); err != nil {
		handleResponse(c, "error while deleting sale by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data successfully deleted")
}
