package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"connected/api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateSale godoc
// @Router       /sale [POST]
// @Summary      Creates a new sale
// @Description  create a new sale
// @Tags         sale
// @Accept       json
// @Produce      json
// @Param        basket body models.CreateSale false "sale"
// @Success      201  {object}  models.Sale
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
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

	sale, err := h.storage.Sale().GetByIdSales(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting sale by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, sale)
}

// GetByIDSales retrieves sale information by ID.
// @Summary Get sale by ID
// @Tags sale
// @Accept json
// @Produce json
// @Param id path string true "Sale ID"
// @Success 200 {object} models.Sale
// @Failure 400 {string} models.Response
// @Failure 500 {string} models.Response
// @Router /sale/{id} [get]
func (h Handler) GetByIDSales(c *gin.Context) {
	var err error

	uid := c.Param("id")

	sale, err := h.storage.Sale().GetByIdSales(models.PrimaryKey{
		ID: uid,
	})
	if err != nil {
		handleResponse(c, "error while getting sale by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, sale)
}

// GetListSales returns a list of sales.
// @Summary Get a list of sales
// @Tags sale
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Param        fromprice    query     uint64  false  "price from for response"
// @Param        toprice    query     uint64  false  "price to for response"
// @Success 200 {object} models.SaleRepos
// @Failure 400 {string} models.Response
// @Failure 500 {string} models.Response
// @Router /sales [get]
func (h Handler) GetListSales(c *gin.Context) {
	var (
		page, limit int
		search      string
		priceFrom   uint64
		priceTo     uint64
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

	priceFromStr := c.Query("price_from")
	if priceFromStr != "" {
		priceFrom, err = strconv.ParseUint(priceFromStr, 64, 64)
		if err != nil {
			handleResponse(c, "error while parsing price_from", http.StatusBadRequest, err.Error())
			return
		}
	}

	priceToStr := c.Query("price_to")
	if priceToStr != "" {
		priceTo, err = strconv.ParseUint(priceToStr, 64, 64)
		if err != nil {
			handleResponse(c, "error while parsing price_to", http.StatusBadRequest, err.Error())
			return
		}
	}

	resp, err := h.storage.Sale().GetListSales(models.GetListRequestSale{
		Page:      page,
		Limit:     limit,
		Search:    search,
		FromPrice: uint64(priceFrom),
		ToPrice:   uint64(priceTo),
	})
	if err != nil {
		handleResponse(c, "error while getting sale", http.StatusInternalServerError, err.Error()) // Adjusted this line
		return
	}

	handleResponse(c, "", http.StatusOK, resp)
}

// UpdateSales godoc
// @Router       /sale/{id} [PUT]
// @Summary      Update sale
// @Description  update sale
// @Tags         sale
// @Accept       json
// @Produce      json
// @Param 		 id path string true "sale"
// @Param        user body models.UpdateSale true "sale"
// @Success      200  {object}  models.Sale
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateSales(c *gin.Context) {
	updateSale := models.UpdateSale{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	updateSale.ID = uid

	if err := c.ShouldBindJSON(&updateSale); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.Sale().UpdateSales(updateSale)
	if err != nil {
		handleResponse(c, "error while updating sale", http.StatusInternalServerError, err.Error())
		return
	}

	sales, err := h.storage.Sale().GetByIdSales(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "error while getting sale by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, sales)
}

// DeleteSales deletes sale by ID.
// @Summary Delete sale by ID
// @Tags sale
// @Accept json
// @Produce json
// @Param id path string true "Sale ID"
// @Success 200 {string} string models.Response
// @Failure 400 {string} string models.Response
// @Failure 500 {string} string models.Response
// @Router /sale/{id} [delete]
func (h Handler) DeleteSales(c *gin.Context) {
	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err = h.storage.Sale().DeleteSales(models.PrimaryKey{
		ID: id.String(),
	}); err != nil {
		handleResponse(c, "error while deleting sale by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "data successfully deleted")
}

// StartSale godoc
// @Router       /star_sale [POST]
// @Summary      Creates a new star_sale
// @Description  create a new star_sale
// @Tags         star_sale
// @Accept       json
// @Produce      json
// @Param        star_sale body models.StartSale false "star_sale"
// @Success      201  {object}  models.StartSale
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) StartSale(c *gin.Context) {

	startSale := models.StartSale{}

	if err := c.ShouldBindJSON(&startSale); err != nil {
		handleResponse(c, "Error: Failed to read request body from client", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Sale().CreateSales(models.CreateSale{
		Clientname:       startSale.Clientname,
		Cashier_id:       startSale.Cashier_id,
		Branch_id:        startSale.Branch_id,
		Shopassistant_id: startSale.Shopassistant_id,
		Status_type:      startSale.Status_type,
		Payment_type:     startSale.Payment_type,
	})
	if err != nil {
		handleResponse(c, "Error: Failed to create sale", http.StatusInternalServerError, err)
		return
	}

	sale, err := h.storage.Sale().GetByIdSales(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(c, "Error: Failed to get sale by ID", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusCreated, sale)
}

// EndSales godoc
// @Router       /end-sale/{id} [PUT]
// @Summary      Update end-sale by ID
// @Description  Update end-sale by ID
// @Tags         end-sale
// @Accept       json
// @Produce      json
// @Param        id path string true "End-sale ID"
// @Param        end-sale body models.EndSales true "End-sale"
// @Success      200 {object} models.Response
// @Failure      400 {object} models.Response "Error: Sale ID is required"
// @Failure      404 {object} models.Response "Error: Basket list not found"
// @Failure      500 {object} models.Response "Error: Internal server error"
func (h Handler) EndSales(c *gin.Context) {
	saleID := c.Param("id")
	if saleID == "" {
		handleResponse(c, "Error: Sale ID is required", http.StatusBadRequest, nil)
		return
	}

	basketResponse, err := h.storage.Basket().GetListBasket(models.GetListRequest{
		Page:   1,
		Limit:  10,
		Search: saleID})
	if err != nil {
		handleResponse(c, "Error: Basket list not found", http.StatusNotFound, err.Error())
		return
	}

	var totalSum uint64 = 0
	for _, basket := range basketResponse.Baskets {
		totalSum += uint64(basket.Price)
	}

	fmt.Println("totalpricerrrrrrrrrrrrrrrrrrrrrrrrrrrr", totalSum)
	id, err := h.storage.Sale().UpdatePrice(context.Background(), uint64(totalSum), saleID)
	if err != nil {
		handleResponse(c, "Error: Internal server error1111111111", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storage.Sale().GetByIdSales(models.PrimaryKey{ID: id})
	if err != nil {
		handleResponse(c, "Error: Internal server error2222222222", http.StatusInternalServerError, err.Error())
		return
	}

	// Xodimlarni va xodim tariflarini olish
	tariffs, err := h.storage.Staff_Tarif().GetListStaff_Tarifs(models.GetListRequest{
		Page:  1,
		Limit: 10,
	})
	if err != nil {
		handleResponse(c, "Error: Internal server errorqqqq", http.StatusInternalServerError, err.Error())
		return
	}

	staffs, err := h.storage.Staff().GetListStaff(models.GetListRequest{
		Page:  1,
		Limit: 100,
	})
	if err != nil {
		handleResponse(c, "Error: Internal server erroraaaaaa", http.StatusInternalServerError, err.Error())
		return
	}

	// Xodimlarni ma'lumotlarini qo'llash
	staffsMap := make(map[string]models.Staff)
	staffBalance := make(map[string]uint64)
	for _, staff := range staffs.Staffs {
		if resp.Cashier_id == staff.ID || resp.Shopassistant_id == staff.ID {
			staffsMap[staff.Tarif_id] = staff
			staffBalance[staff.Tarif_id] = staff.Balance
		}
	}

	// Xodimlar tariflarini o'zlashtirish
	tariffsMap := make(map[string]models.Staff_Tarif)
	for _, tariff := range tariffs.Staff_Tarif_Repos {
		tariffsMap[tariff.ID] = tariff
	}

	// Xodimlar uchun yangilashlar
	for _, staff := range staffsMap {
		tariff, found := tariffsMap[staff.Tarif_id]
		if !found {
			continue
		}

		switch {
		case tariff.Tarif_Type_Enum == "fixed":
			switch resp.Payment_type {
			case "cash":
				staffBalance[staff.ID] += uint64(tariff.Amount_For_Card)
			case "card":
				staffBalance[staff.ID] += uint64(tariff.Amount_For_Cashe)
			}
		case tariff.Tarif_Type_Enum == "percent":
			switch resp.Payment_type {
			case "cash":
				staffBalance[staff.ID] += uint64(totalSum * uint64(tariff.Amount_For_Cashe) / 100)
			case "card":
				staffBalance[staff.ID] += uint64(totalSum * uint64(tariff.Amount_For_Card) / 100)
			}
		}

		// Xodimlarni yangilash
		_, err := h.storage.Staff().UpdateStaffs(models.UpdateStaff{
			ID:        staff.ID,
			BranchID:  staff.Branch_id,
			TariffID:  staff.Tarif_id,
			StaffType: staff.Type_stuff,
			Name:      staff.Name,
			Balance:   staffBalance[staff.ID],
			Login:     staff.Login,
		})
		if err != nil {
			handleResponse(c, "Error: Internal server errorccccc", http.StatusInternalServerError, err.Error())
			return
		}

		// Tranzaktsiyani yaratish
		_, err = h.storage.Transaction().CreateTransaction(models.CreateTransaction{
			Sale_id:          saleID,
			Staff_id:         staff.ID,
			Transaction_type: "topup",
			Sourcetype:       "sales",
			Amount:           float64(totalSum),
			Description:      "staff ",
		})
		if err != nil {
			handleResponse(c, "Error: Internal server errorcccccc", http.StatusInternalServerError, err.Error())
			return
		}
	}

	handleResponse(c, "Success", http.StatusOK, resp)

}
