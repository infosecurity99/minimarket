package handler

import (
	"errors"
	"net/http"
	"strconv"

	"connected/api/models"
	"connected/pkg/check"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateStaff godoc
// @Router       /staff [POST]
// @Summary      Creates a new staff
// @Description  create a new staff
// @Tags         staff
// @Accept       json
// @Produce      json
// @Param        staff body models.CreateStaff false "staff"
// @Success      201  {object}  models.Staff
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateStaff(c *gin.Context) {
	createStaff := models.CreateStaff{}

	if err := c.ShouldBindJSON(&createStaff); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Staff().CreateStaff(createStaff)
	if err != nil {
		handleResponse(c, "error while creating staff", http.StatusInternalServerError, err)
		return
	}

	staff, err := h.storage.Staff().GetByIdStaff(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting staff by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, staff)
}

// GetByIdStaff retrieves staff information by ID.
// @Summary Get staff by ID
// @Tags staff
// @Accept json
// @Produce json
// @Param id path string true "Staff ID"
// @Success 200 {object} models.Staff
// @Failure 400 {string} models.Response
// @Failure 500 {string} models.Response
// @Router /staff/{id} [get]
func (h Handler) GetByIdStaff(c *gin.Context) {
	var err error

	uid := c.Param("id")

	staff, err := h.storage.Staff().GetByIdStaff(models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "error while getting staff by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, staff)
}

// GetListStaff returns a list of staff.
// @Summary Get a list of staff
// @Tags staff
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} models.Staff
// @Failure 400 {string} models.Response
// @Failure 500 {string} models.Response
// @Router /staffs [get]
func (h Handler) GetListStaff(c *gin.Context) {
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

	resp, err := h.storage.Staff().GetListStaff(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "error while getting staff", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, resp)
}

// UpdateStaff godoc
// @Router       /staff/{id} [PUT]
// @Summary      Update staff
// @Description  update staff
// @Tags         staff
// @Accept       json
// @Produce      json
// @Param 		 id path string true "staff"
// @Param        user body models.UpdateStaff true "staff"
// @Success      200  {object}  models.Staff
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateStaff(c *gin.Context) {
	updateStaff := models.UpdateStaff{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	updateStaff.ID = uid

	if err := c.ShouldBindJSON(&updateStaff); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.Staff().UpdateStaffs(updateStaff)
	if err != nil {
		handleResponse(c, "error while updating staff", http.StatusInternalServerError, err.Error())
		return
	}

	branch, err := h.storage.Staff().GetByIdStaff(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting staff by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, branch)
}

// DeleteStaff deletes staff information by ID.
// @Summary Delete staff by ID
// @Tags staff
// @Accept json
// @Produce json
// @Param id path string true "Staff ID"
// @Success 200 {string} models.Response
// @Failure 400 {string} models.Response
// @Failure 500 {string} models.Response
// @Router /staff/{id} [delete]
func (h Handler) DeleteStaff(c *gin.Context) {
	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err = h.storage.Staff().DeleteStaff(models.PrimaryKey{
		ID: id.String(),
	}); err != nil {
		handleResponse(c, "error while deleting staff by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data successfully deleted")
}

// UpdateStaffPassword godoc
// @Router       /staff/{id} [PATCH]
// @Summary      Update staff password
// @Description  update staff password
// @Tags         staff
// @Accept       json
// @Produce      json
// @Param 		 id path string true "staff_id"
// @Param        staff body models.UpdateStaffPassword true "staff"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateStaffPassword(c *gin.Context) {
	updateStaffPassword := models.UpdateStaffPassword{}

	if err := c.ShouldBindJSON(&updateStaffPassword); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	uid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		handleResponse(c, "error while parsing uuid", http.StatusBadRequest, err.Error())
		return
	}

	updateStaffPassword.ID = uid.String()

	oldPassword, err := h.storage.Staff().GetPassword(updateStaffPassword.ID)
	if err != nil {
		handleResponse(c, "error while getting password by id", http.StatusInternalServerError, err.Error())
		return
	}

	if oldPassword != updateStaffPassword.OldPassword {
		handleResponse(c, "old password is not correct", http.StatusBadRequest, "old password is not correct")
		return
	}

	if err = check.ValidatePassword(updateStaffPassword.NewPassword); err != nil {
		handleResponse(c, "new password is weak", http.StatusBadRequest, err.Error())
		return
	}

	if err = h.storage.Staff().UpdatePassword(updateStaffPassword); err != nil {
		handleResponse(c, "error while updating user password by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "password successfully updated")
}
