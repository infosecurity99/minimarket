package handler

import (
	"errors"
	"net/http"
	"strconv"

	"connected/api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateStaff_Tarif godoc
// @Router       /staff_tarif [POST]
// @Summary      Creates a new staff_tarif
// @Description  create a new staff_tarif
// @Tags         staff_tarif
// @Accept       json
// @Produce      json
// @Param        basket body models.CreateStaff_Tarif false "staff_tarif"
// @Success      201  {object}  models.Staff_Tarif
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
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

	staffTarif, err := h.storage.Staff_Tarif().GetByIdStaff_Tarifs(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting stafftarif by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, staffTarif)
}

// GetByIdStaff_Tarif retrieves staff tarif information by ID.
// @Summary Get staff tarif by ID
// @Tags staff_tarif
// @Accept json
// @Produce json
// @Param id path string true "Staff_Tarif ID"
// @Success 200 {object} models.Staff_Tarif
// @Failure 400 {string} models.Response
// @Failure 500 {string} models.Response
// @Router /staff_tarif/{id} [get]
func (h Handler) GetByIdStaff_Tarif(c *gin.Context) {
	var err error

	uid := c.Param("id")

	staffTarif, err := h.storage.Staff_Tarif().GetByIdStaff_Tarifs(models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "error while getting staff tarif by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, staffTarif)
}

// GetListStaff_Tarif returns a list of staff tarifs.
// @Summary Get a list of staff tarifs
// @Tags staff_tarif
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} models.Staff_Tarif
// @Failure 400 {string} models.Response
// @Failure 500 {string} models.Response
// @Router /staff_tarifs [get]
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

// UpdateStaff_Tarif godoc
// @Router       /staff_tarif/{id} [PUT]
// @Summary      Update staff_tarif
// @Description  update staff_tarif
// @Tags         staff_tarif
// @Accept       json
// @Produce      json
// @Param 		 id path string true "staff_tarif"
// @Param        user body models.UpdateStaff_Tarif true "staff_tarif"
// @Success      200  {object}  models.Staff_Tarif
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateStaff_Tarif(c *gin.Context) {
	updateStaffTarif := models.UpdateStaff_Tarif{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	updateStaffTarif.ID = uid

	if err := c.ShouldBindJSON(&updateStaffTarif); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.Staff_Tarif().UpdateStaff_Tarifs(updateStaffTarif)
	if err != nil {
		handleResponse(c, "error while updating stafftarif", http.StatusInternalServerError, err.Error())
		return
	}

	staffTarif, err := h.storage.Staff_Tarif().GetByIdStaff_Tarifs(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting stafftarif by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, staffTarif)
}

// DeleteStaff_Tarif deletes staff tarif information by ID.
// @Summary Delete staff tarif by ID
// @Tags staff_tarif
// @Accept json
// @Produce json
// @Param id path string true "Staff_Tarif ID"
// @Success 200 {string} models.Response
// @Failure 400 {string} models.Response
// @Failure 500 {string} models.Response
// @Router /staff_tarif/{id} [delete]
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
		handleResponse(c, "error while deleting stafftarif by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data successfully deleted")
}
