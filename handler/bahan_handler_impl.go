package handler

import (
	"Recipe_App/models"
	"Recipe_App/usecase"
	"Recipe_App/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BahanHandlerImpl struct {
	bahanUsecase usecase.BahanUsecase
}

func NewBahanHandlerImpl(bahanUsecase usecase.BahanUsecase) BahanHandler {
	return &BahanHandlerImpl{
		bahanUsecase: bahanUsecase,
	}
}

// CreateBahan		godoc
// @Summary			Create bahan
// @Description		Create a new master data bahan.
// @Param			bahan body models.BahanInput true "Create bahan"
// @Produce			application/json
// @Tags			bahan
// @Success			200 {object} utils.Response
// @Router			/api/v1/bahan [post]
func (h *BahanHandlerImpl) Create(c *gin.Context) {
	var input models.BahanInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Create Data Master Bahan Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	bahan, errService := h.bahanUsecase.Create(input)
	if errService != nil {
		response := utils.ApiResponse("Create Data Master Bahan Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Create Data Master Bahan Success", http.StatusOK, "success", bahan)
	c.JSON(http.StatusOK, response)
}

// UpdateBahan		godoc
// @Summary			Update bahan
// @Description		Update data master bahan.
// @Param			bahan body models.BahanInput true "Update bahan"
// @Param			id path integer true "find bahan by id"
// @Produce			application/json
// @Tags			bahan
// @Success			200 {object} utils.Response
// @Router			/api/v1/bahan/{id} [put]
func (h *BahanHandlerImpl) Update(c *gin.Context) {
	id, errParam := strconv.Atoi(c.Param("id"))

	if errParam != nil {
		response := utils.ApiResponse("Invalid id params", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input models.BahanInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Update Data Master Bahan Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	bahan, errService := h.bahanUsecase.Update(uint(id), input)
	if errService != nil {
		response := utils.ApiResponse("Update Data Master Bahan Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Update Data Master Bahan Success", http.StatusOK, "success", bahan)
	c.JSON(http.StatusOK, response)
}

// GetById 				godoc
// @Summary				Get Single bahan by id.
// @Param				id path string true "get bahan by id"
// @Description			Return data master bahan where similar with id.
// @Produce				application/json
// @Tags				bahan
// @Success				200 {object} utils.Response
// @Router				/api/v1/bahan/{id} [get]
func (h *BahanHandlerImpl) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := utils.ApiResponse("Invalid id params", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	bahan, errService := h.bahanUsecase.GetById(uint(id))
	if errService != nil {
		response := utils.ApiResponse("Get data bahan failed", http.StatusBadRequest, "error", errService)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Get data bahan success", http.StatusOK, "success", bahan)
	c.JSON(http.StatusOK, response)
}

// Delete			 	godoc
// @Summary				Delete bahan by id.
// @Param				id path string true "delete bahan by id"
// @Description			Return data boolean.
// @Produce				application/json
// @Tags				bahan
// @Success				200 {object} utils.Response
// @Router				/api/v1/bahan/delete/{id} [delete]
func (h *BahanHandlerImpl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := utils.ApiResponse("Invalid id params", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	bahan, errService := h.bahanUsecase.Delete(uint(id))
	if errService != nil || !bahan {
		response := utils.ApiResponse("Delete data bahan failed", http.StatusBadRequest, "error", errService)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Get data bahan success", http.StatusOK, "success", bahan)
	c.JSON(http.StatusOK, response)
}

// GetAll 				godoc
// @Summary				Get all data bahan.
// @Param				page query int64 true "page number"
// @Param				limit query int64 true "limit number"
// @Description			Return data master bahan with pagination.
// @Produce				application/json
// @Tags				bahan
// @Success				200 {object} utils.Response
// @Router				/api/v1/bahan [get]
func (h *BahanHandlerImpl) GetAll(c *gin.Context) {
	pagination := utils.Pagination(c)

	response, totalRows, err := h.bahanUsecase.GetAll(&pagination)
	if err != nil {
		response := utils.ApiResponseWithPaginate("Get all data bahan failed", http.StatusBadRequest, "error", err, &totalRows)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	responses := utils.ApiResponseWithPaginate("Get data bahan success", http.StatusOK, "success", response, &totalRows)
	c.JSON(http.StatusOK, responses)
}
