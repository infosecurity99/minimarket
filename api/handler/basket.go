package handler

import (
	"connected/api/models"
	"github.com/google/uuid"

	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Basket().CreateBasket(createBasket)
	if err != nil {
		handleResponse(c, "error while creating basket", http.StatusInternalServerError, err)
		return
	}

	basket, err := h.storage.Basket().GetByIdBasket(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting basket by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, basket)
}

// GetByIDBasket godoc
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
func (h Handler) GetByIDBasket(c *gin.Context) {
	var err error

	uid := c.Param("id")

	basket, err := h.storage.Basket().GetByIdBasket(models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "error while getting basket by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, basket)
}

// GetListBasket godoc
// @Router       /baskets [GET]
// @Summary      Get basket list
// @Description  get basket list
// @Tags         basket
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Success      200  {object}  models.BasketResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetListBasket(c *gin.Context) {
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

	resp, err := h.storage.Basket().GetListBasket(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "error while getting basket", http.StatusInternalServerError, err)
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
// @Param        user body models.UpdateBasket true "user"
// @Success      200  {object}  models.Basket
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateBasket(c *gin.Context) {
	updateBasket := models.UpdateBasket{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("invalid uuid"))
		return
	}

	updateBasket.ID = uid

	if err := c.ShouldBindJSON(&updateBasket); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Basket().UpdateBasket(updateBasket)
	if err != nil {
		handleResponse(c, "error while updating basket", http.StatusInternalServerError, err)
		return
	}

	basket, err := h.storage.Basket().GetByIdBasket(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting basket by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, basket)
}

// DeleteBasket godoc
// @Router       /basket/{id} [DELETE]
// @Summary      Delete basket
// @Description  delete basket
// @Tags         basket
// @Accept       json
// @Produce      json
// @Param 		 id path string true "basket_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteBasket(c *gin.Context) {
	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err = h.storage.Basket().DeleteBasket(models.PrimaryKey{
		ID: id.String(),
	}); err != nil {
		handleResponse(c, "error while deleting basket by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data successfully deleted")
}
