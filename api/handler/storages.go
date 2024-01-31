package handler

import (
	"errors"
	"net/http"
	"strconv"

	"connected/api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateStorages godoc
// @Router       /storage [POST]
// @Summary      Creates a new storage
// @Description  create a new storage
// @Tags         storage
// @Accept       json
// @Produce      json
// @Param        basket body models.CreateStorage false "storage"
// @Success      201  {object}  models.Storage
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateStorages(c *gin.Context) {
	createStorage := models.CreateStorage{}

	if err := c.ShouldBindJSON(&createStorage); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Storag().CreateStorages(createStorage)
	if err != nil {
		handleResponse(c, "error while creating storage", http.StatusInternalServerError, err)
		return
	}

	storage, err := h.storage.Storag().GetByIdStorages(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting storage by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, storage)
}

// GetByIdSorages retrieves storage information by ID.
// @Summary Get storage by ID
// @Tags storage
// @Accept json
// @Produce json
// @Param id path string true "Storage ID"
// @Success 200 {object} models.Storage
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /storage/{id} [get]
func (h Handler) GetByIdSorages(c *gin.Context) {
	var err error

	uid := c.Param("id")

	storage, err := h.storage.Storag().GetByIdStorages(models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "error while getting storage by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, storage)
}

// GetListStorages returns a list of storage.
// @Summary Get a list of storage
// @Tags storage
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} models.Storage
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /storages [get]
func (h Handler) GetListStorages(c *gin.Context) {
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

	resp, err := h.storage.Storag().GetListStorages(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "error while getting branch", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, resp)
}

// UpdateSorages godoc
// @Router       /storage/{id} [PUT]
// @Summary      Update storage
// @Description  update storage
// @Tags         storage
// @Accept       json
// @Produce      json
// @Param 		 id path string true "storage"
// @Param        storage body models.UpdateStorage true "storage"
// @Success      200  {object}  models.Storage
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateSorages(c *gin.Context) {
	updateStorage := models.UpdateStorage{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	updateStorage.ID = uid

	if err := c.ShouldBindJSON(&updateStorage); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.Storag().UpdateStorages(updateStorage)
	if err != nil {
		handleResponse(c, "error while updating storage", http.StatusInternalServerError, err.Error())
		return
	}

	storage, err := h.storage.Storag().GetByIdStorages(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting storage by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, storage)
}

// DeleteStorages deletes storage information by ID.
// @Summary Delete storage by ID
// @Tags storage
// @Accept json
// @Produce json
// @Param id path string true "Storage ID"
// @Success 200 {string} string "Data successfully deleted"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /storage/{id} [delete]
func (h Handler) DeleteStorages(c *gin.Context) {
	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err = h.storage.Storag().DeleteStorages(models.PrimaryKey{
		ID: id.String(),
	}); err != nil {
		handleResponse(c, "error while deleting storage by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data successfully deleted")
}
