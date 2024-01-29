package handler

import (
	"connected/api/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//create storage
func (h Handler) CreateStorages(c *gin.Context) {
	createstorage := models.CreateStorage{}

	if err := c.ShouldBindJSON(&createstorage); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Storag().CreateStorages(createstorage)
	if err != nil {
		handleResponse(c, "error while creating storage", http.StatusInternalServerError, err)
		return
	}

	storage, err := h.storage.Storag().GetByIdStorages(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting getbyid by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, storage)
}

//get by id  storage
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

//getlist storage
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

//update   storage
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

//delete   storage
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
