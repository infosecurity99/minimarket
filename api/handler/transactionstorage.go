package handler

import (
	"connected/api/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//create transactionstorage
func (h Handler) CreateTransactionStorage(c *gin.Context) {
	createTransactionStorage := models.CreateTransactionStorage{}

	if err := c.ShouldBindJSON(&createTransactionStorage); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.TransactionStorage().CreateTransactionStorage(createTransactionStorage)
	if err != nil {
		handleResponse(c, "error while creating transactionstorage", http.StatusInternalServerError, err)
		return
	}

	transactionstorage, err := h.storage.TransactionStorage().GetByIdTranasactionStorage(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting transaction by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, transactionstorage)
}

//getbyid transactionstorage
func (h Handler) GetByIdTranasactionStorage(c *gin.Context) {
	var err error

	uid := c.Param("id")

	transactionstorage, err := h.storage.TransactionStorage().GetByIdTranasactionStorage(models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "error while getting transactionstorage by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, transactionstorage)
}

//getlist transactionstorage
func (h Handler) GetListTransactionStorage(c *gin.Context) {
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

	resp, err := h.storage.TransactionStorage().GetListTransactionStorage(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "error while getting transactionstorage", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, resp)
}

//update transactionstorage
func (h Handler) UpdateTransactionStorage(c *gin.Context) {
	updateTranasactionStorage := models.UpdateTransactionStorage{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	updateTranasactionStorage.ID = uid

	if err := c.ShouldBindJSON(&updateTranasactionStorage); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.TransactionStorage().UpdateTransactionStorage(updateTranasactionStorage)
	if err != nil {
		handleResponse(c, "error while updating transactionstorage", http.StatusInternalServerError, err.Error())
		return
	}

	tranasactionstorage, err := h.storage.TransactionStorage().GetByIdTranasactionStorage(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting tranasactionstorage by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, tranasactionstorage)
}

//delete transactionstorage
func (h Handler) DeleteTransactionStorage(c *gin.Context) {
	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err = h.storage.Branch().Delete(models.PrimaryKey{
		ID: id.String(),
	}); err != nil {
		handleResponse(c, "error while deleting branch by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data successfully deleted")
}
