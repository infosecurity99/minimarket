package handler

import (
	"errors"
	"net/http"
	"strconv"

	"connected/api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// create transaction
func (h Handler) CreateTransaction(c *gin.Context) {
	createTransaction := models.CreateTransaction{}

	if err := c.ShouldBindJSON(&createTransaction); err != nil {
		handleResponse(c, "Error: Failed to parse request body JSON", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Transaction().CreateTransaction(createTransaction)
	if err != nil {
		handleResponse(c, "Error: Failed to create transaction", http.StatusInternalServerError, err)
		return
	}

	transaction, err := h.storage.Transaction().GetByIdTransaction(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error: Failed to find transaction by ID", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, transaction)
}

// get by id  transaction
func (h Handler) GetByIdTransaction(c *gin.Context) {
	var err error

	uid := c.Param("id")

	transaction, err := h.storage.Transaction().GetByIdTransaction(models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "Error: Failed to find transaction by ID", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, transaction)
}

// getlist transaction
func (h Handler) GetListTransaction(c *gin.Context) {
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

	resp, err := h.storage.Transaction().GetListTransaction(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "Error: Failed to get transaction list", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, resp)
}

// create transaction
func (h Handler) UpdateTransaction(c *gin.Context) {
	updateTransaction := models.UpdateTransaction{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "Invalid UUID", http.StatusBadRequest, errors.New("UUID is not valid"))
		return
	}

	if err := c.ShouldBindJSON(&updateTransaction); err != nil {
		handleResponse(c, "Error: Failed to parse request body JSON", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Transaction().UpdateTransaction(updateTransaction)
	if err != nil {
		handleResponse(c, "Error: Failed to update transaction", http.StatusInternalServerError, err)
		return
	}

	transaction, err := h.storage.Transaction().GetByIdTransaction(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error: Failed to find transaction by ID", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, transaction)
}

// delete transaction
func (h Handler) DeleteTransaction(c *gin.Context) {
	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "Invalid UUID", http.StatusBadRequest, err)
		return
	}

	if err = h.storage.Transaction().DeleteTransaction(models.PrimaryKey{
		ID: id.String(),
	}); err != nil {
		handleResponse(c, "Error: Failed to delete transaction by ID", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, "Data deleted successfully")
}
