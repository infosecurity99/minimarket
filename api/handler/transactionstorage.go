package handler

import (
	"errors"
	"net/http"
	"strconv"

	"connected/api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateTransactionStorage godoc
// @Router       /transaction_storage [POST]
// @Summary      Creates a new storage
// @Description  create a new storage
// @Tags         transaction_storage
// @Accept       json
// @Produce      json
// @Param        basket body models.CreateTransactionStorage false "transaction_storage"
// @Success      201  {object}  models.TransactionStorage
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateTransactionStorage(c *gin.Context) {
	createTransactionStorage := models.CreateTransactionStorage{}

	if err := c.ShouldBindJSON(&createTransactionStorage); err != nil {
		handleResponse(c, "Error: Failed to parse request body JSON", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.TransactionStorage().CreateTransactionStorage(createTransactionStorage)
	if err != nil {
		handleResponse(c, "Error: Failed to create transaction storage", http.StatusInternalServerError, err)
		return
	}

	transactionstorage, err := h.storage.TransactionStorage().GetByIdTranasactionStorage(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error: Failed to find transaction storage by ID", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, transactionstorage)
}

// GetByIdTranasactionStorage retrieves transaction storage information by ID.
// @Summary Get transaction storage by ID
// @Tags transaction_storage
// @Accept json
// @Produce json
// @Param id path string true "Transaction Storage ID"
// @Success 200 {object} models.TransactionStorage
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /transaction_storage/{id} [get]
func (h Handler) GetByIdTranasactionStorage(c *gin.Context) {
	var err error

	uid := c.Param("id")

	transactionstorage, err := h.storage.TransactionStorage().GetByIdTranasactionStorage(models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "Error: Failed to find transaction storage by ID", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, transactionstorage)
}

// GetListTransactionStorage returns a list of transaction storage.
// @Summary Get a list of transaction storage
// @Tags transaction_storage
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {object} models.Transaction
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /transaction_storages [get]
func (h Handler) GetListTransactionStorage(c *gin.Context) {
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

	resp, err := h.storage.TransactionStorage().GetListTransactionStorage(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "Error: Failed to get transaction storage list", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, resp)
}

// UpdateTransactionStorage godoc
// @Router       /transaction_storage/{id} [PUT]
// @Summary      Update transaction_storage
// @Description  update transaction_storage
// @Tags         transaction_storage
// @Accept       json
// @Produce      json
// @Param 		 id path string true "transaction_storage"
// @Param        user body models.UpdateTransactionStorage true "transaction_storage"
// @Success      200  {object}  models.TransactionStorage
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateTransactionStorage(c *gin.Context) {
	updateTransactionStorage := models.UpdateTransactionStorage{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "Invalid UUID", http.StatusBadRequest, errors.New("UUID is not valid"))
		return
	}

	if err := c.ShouldBindJSON(&updateTransactionStorage); err != nil {
		handleResponse(c, "Error: Failed to parse request body JSON", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.TransactionStorage().UpdateTransactionStorage(updateTransactionStorage)
	if err != nil {
		handleResponse(c, "Error: Failed to update transaction storage", http.StatusInternalServerError, err)
		return
	}

	transactionstorage, err := h.storage.TransactionStorage().GetByIdTranasactionStorage(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error: Failed to find transaction storage by ID", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, transactionstorage)
}

// DeleteTransactionStorage deletes transaction storage by ID.
// @Summary Delete transaction storage by ID
// @Tags transaction_storage
// @Accept json
// @Produce json
// @Param id path string true "Transaction Storage ID"
// @Success 200 {string} string "Data deleted successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /transaction_storage/{id} [delete]
func (h Handler) DeleteTransactionStorage(c *gin.Context) {
	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "Invalid UUID", http.StatusBadRequest, err)
		return
	}

	if err = h.storage.TransactionStorage().DeleteTransactionStorage(models.PrimaryKey{
		ID: id.String(),
	}); err != nil {
		handleResponse(c, "Error: Failed to delete transaction storage by ID", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, "Data deleted successfully")
}
