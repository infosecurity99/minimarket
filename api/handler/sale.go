package handler

import (
	"connected/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//create sale  handler
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

	branch, err := h.storage.Sale().GetByIdSales(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting branch by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, branch)
}

//getbyid ale handler

func (h Handler) GetByIDSales(c *gin.Context) {
}

//getlist saler  handler
func (h Handler) GetListSales(c *gin.Context) {

}

//update slae handler
func (h Handler) UpdateSales(c *gin.Context) {

}

//delete sale handler
func (h Handler) DeleteSales(c *gin.Context) {

}
