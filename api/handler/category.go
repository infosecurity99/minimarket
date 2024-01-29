package handler

import (
	"errors"
	"net/http"
	"strconv"

	"connected/api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateCategory godoc
// @Router       /category [POST]
// @Summary      Creates a new category
// @Description  create a new category
// @Tags         category
// @Accept       json
// @Produce      json
// @Param        basket body models.CreateCategory false "category"
// @Success      201  {object}  models.Category
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateCategory(c *gin.Context) {
	createCategory := models.CreateCategory{}

	if err := c.ShouldBindJSON(&createCategory); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Category().CreateCategory(createCategory)
	if err != nil {
		handleResponse(c, "error while creating category", http.StatusInternalServerError, err)
		return
	}

	category, err := h.storage.Category().GetByIdCategory(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting category by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, category)
}

// GetByIDCategory retrieves category information by ID.
// @Summary Get category by ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} models.Category
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /category/{id} [get]
func (h Handler) GetByIDCategory(c *gin.Context) {
	var err error

	uid := c.Param("id")

	category, err := h.storage.Category().GetByIdCategory(models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "error while getting category by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, category)
}

// GetListCategory returns a list of categories.
// @Summary Get a list of categories
// @Tags category
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} models.Category
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /categories [get]
func (h Handler) GetListCategory(c *gin.Context) {
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

	resp, err := h.storage.Category().GetListCategory(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "error while getting category", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, resp)
}

// UpdateCategory godoc
// @Router       /category/{id} [PUT]
// @Summary      Update category
// @Description  update category
// @Tags         category
// @Accept       json
// @Produce      json
// @Param 		 id path string true "category"
// @Param        user body models.UpdateCategory true "category"
// @Success      200  {object}  models.Category
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateCategory(c *gin.Context) {
	updateCategory := models.UpdateCategory{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	updateCategory.ID = uid

	if err := c.ShouldBindJSON(&updateCategory); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.Category().UpdateCategory(updateCategory)
	if err != nil {
		handleResponse(c, "error while updating category", http.StatusInternalServerError, err.Error())
		return
	}

	category, err := h.storage.Category().GetByIdCategory(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting category by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, category)
}

// DeleteCategory deletes category by ID.
// @Summary Delete category by ID
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {string} string "Data successfully deleted"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /category/{id} [delete]
func (h Handler) DeleteCategory(c *gin.Context) {
	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err = h.storage.Category().DeleteCategory(models.PrimaryKey{
		ID: id.String(),
	}); err != nil {
		handleResponse(c, "error while deleting category by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data successfully deleted")
}
