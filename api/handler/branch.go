package handler

import (
	"errors"
	"net/http"
	"strconv"

	"connected/api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateBranch godoc
// @Router       /branch [POST]
// @Summary      Creates a new branch
// @Description  create a new branch
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param        branch body models.CreateBranch false "branch"
// @Success      201  {object}  models.Branch
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateBranch(c *gin.Context) {
	createBranch := models.CreateBranch{}

	if err := c.ShouldBindJSON(&createBranch); err != nil {
		handleResponse(c, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Branch().Create(createBranch)
	if err != nil {
		handleResponse(c, "error while creating branch", http.StatusInternalServerError, err)
		return
	}

	branch, err := h.storage.Branch().GetByID(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting branch by id1", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, branch)
}

/// GetByID godoc
// @Router       /branch/{id} [GET]
// @Summary      Gets branch
// @Description  get branch by ID
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param        id path string true "branch"
// @Success      200  {object}  models.Branch
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetByID(c *gin.Context) {
	var err error

	uid := c.Param("id")

	branch, err := h.storage.Branch().GetByID(models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "error while getting branch by id2", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, branch)
}

// GetList godoc
// @Router       /branchs [GET]
// @Summary      Get branch list
// @Description  get branch list
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Success      200  {object}  models.BranchResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetList(c *gin.Context) {
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

	resp, err := h.storage.Branch().GetList(models.GetListRequest{
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

// UpdateBranch godoc
// @Router       /branch/{id} [PUT]
// @Summary      Update branch
// @Description  update branch
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param 		 id path string true "branch_id"
// @Param        branch body models.UpdateBranch true "branch"
// @Success      200  {object}  models.Branch
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateBranch(c *gin.Context) {
	updateBranch := models.UpdateBranch{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	updateBranch.ID = uid

	if err := c.ShouldBindJSON(&updateBranch); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.Branch().Update(updateBranch)
	if err != nil {
		handleResponse(c, "error while updating branch", http.StatusInternalServerError, err.Error())
		return
	}

	branch, err := h.storage.Branch().GetByID(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting branch by id3", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, branch)
}

// Delete godoc
// @Router       /branch/{id} [DELETE]
// @Summary      Delete branch
// @Description  delete branch
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param 		 id path string true "branch_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) Delete(c *gin.Context) {
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
