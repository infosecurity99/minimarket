package handler

import (
	"errors"
	"net/http"
	"strconv"

	"connected/api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateTransaction godoc
// @Router       /transaction [POST]
// @Summary      Creates a new storage
// @Description  create a new storage
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param        transaction body models.CreateTransaction false "transaction"
// @Success      201  {object}  models.Transaction
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
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

// GetByIdTransaction retrieves transaction information by ID.
// @Summary Get transaction by ID
// @Tags transaction
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} models.Transaction
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /transaction/{id} [get]
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

// GetListTransaction returns a list of transactions.
// @Summary Get a list of transactions
// @Tags transaction
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Param        fromamount    query     float64  false  "amount from for response"
// @Param        toamount    query     float64  false  "amount to for response"
// @Success 200 {object} models.Transaction
// @Failure 400 {string} models.Response
// @Failure 500 {string} models.Response
// @Router /transactions [get]
func (h Handler) GetListTransaction(c *gin.Context) {
	var (
		page, limit          int
		search               string
		fromAmount, toAmount uint64
		err                  error
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

	fromAmountStr := c.DefaultQuery("fromamount", "0")
	fromAmount, err = strconv.ParseUint(fromAmountStr, 64, 64)
	if err != nil {
		handleResponse(c, "Error: Failed to parse fromamount parameter", http.StatusBadRequest, err)
		return
	}

	toAmountStr := c.DefaultQuery("toamount", "0")
	toAmount, err = strconv.ParseUint(toAmountStr, 64, 64)
	if err != nil {
		handleResponse(c, "Error: Failed to parse toamount parameter", http.StatusBadRequest, err)
		return
	}

	resp, err := h.storage.Transaction().GetListTransaction(models.GetListRequestTransaction{
		Page:       page,
		Limit:      limit,
		Search:     search,
		FromAmount: fromAmount,
		ToAmount:   toAmount,
	})
	if err != nil {

		handleResponse(c, "Error: Failed to get transaction list", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, resp)
}

// UpdateTransaction godoc
// @Router       /transaction/{id} [PUT]
// @Summary      Update sale
// @Description  update sale
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param 		 id path string true "transaction"
// @Param        transaction body models.UpdateTransaction true "transaction"
// @Success      200  {object}  models.Transaction
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
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

// DeleteTransaction deletes transaction information by ID.
// @Summary Delete transaction by ID
// @Tags transaction
// @Accept json
// @Produce json
// @Param id path string true "transaction"
// @Success 200 {string} string "Data deleted successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /transaction/{id} [delete]
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
