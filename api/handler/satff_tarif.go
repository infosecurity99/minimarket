package handler

import (
	"connected/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//create CreateStaff_Tarif
func (h Handler) CreateStaff_Tarif(c *gin.Context) {
	createStaffTarif := models.CreateStaff_Tarif{}

	if err := c.ShouldBindJSON(&createStaffTarif); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Staff_Tarif().CreateStaff_Tarifs(createStaffTarif)
	if err != nil {
		handleResponse(c, "error while creating user", http.StatusInternalServerError, err)
		return
	}

	satfftarif, err := h.storage.Staff_Tarif().GetByIdStaff_Tarifs(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting stafftarif by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, satfftarif)
}

//get by id  CreateStaff_Tarif
func (h Handler) GetByIdStaff_Tarif(c *gin.Context) {

}

//getlist CreateStaff_Tarif
func (h Handler) GetListStaff_Tarif(c *gin.Context) {

}

//create CreateStaff_Tarif
func (h Handler) UpdateStaff_Tarif(c *gin.Context) {

}

//delete CreateStaff_Tarif
func (h Handler) DeleteStaff_Tarif(c *gin.Context) {

}
