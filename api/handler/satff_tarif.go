package handler

import (
	"errors"
	"net/http"
	"strconv"

	"connected/api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// create CreateStaff_Tarif
func (h Handler) CreateStaff_Tarif(c *gin.Context) {
	createStaffTarif := models.CreateStaff_Tarif{}

	if err := c.ShouldBindJSON(&createStaffTarif); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Staff_Tarif().CreateStaff_Tarifs(createStaffTarif)
	if err != nil {
		handleResponse(c, "error while creating stafftarif", http.StatusInternalServerError, err)
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

// get by id  CreateStaff_Tarif
func (h Handler) GetByIdStaff_Tarif(c *gin.Context) {
	var err error

	uid := c.Param("id")

	staftarif2, err := h.storage.Staff_Tarif().GetByIdStaff_Tarifs(models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "error while getting staff tarif by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, staftarif2)
}

// getlist CreateStaff_Tarif
func (h Handler) GetListStaff_Tarif(c *gin.Context) {
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

	resp, err := h.storage.Staff_Tarif().GetListStaff_Tarifs(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "error while getting stafftarif", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, resp)
}

// create CreateStaff_Tarif
func (h Handler) UpdateStaff_Tarif(c *gin.Context) {
	updatestafTarif := models.UpdateStaff_Tarif{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	updatestafTarif.ID = uid

	if err := c.ShouldBindJSON(&updatestafTarif); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.Staff_Tarif().UpdateStaff_Tarifs(updatestafTarif)
	if err != nil {
		handleResponse(c, "error while updating user", http.StatusInternalServerError, err.Error())
		return
	}

	branch, err := h.storage.Staff_Tarif().GetByIdStaff_Tarifs(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting branch by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, branch)
}

// delete CreateStaff_Tarif
func (h Handler) DeleteStaff_Tarif(c *gin.Context) {
	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err = h.storage.Staff_Tarif().DeleteStaff_Tarifs(models.PrimaryKey{
		ID: id.String(),
	}); err != nil {
		handleResponse(c, "error while deleting staf tarif by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data successfully deleted")
}
