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

func (h *BahanHandlerImpl) Create(c *gin.Context) {
	var input models.Bahan
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

func (h *BahanHandlerImpl) Update(c *gin.Context) {
	id, errParam := strconv.Atoi(c.Param("id"))

	if errParam != nil {
		response := utils.ApiResponse("Invalid id params", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input models.Bahan
	input.Id = uint(id)
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Update Data Master Bahan Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	bahan, errService := h.bahanUsecase.Update(input)
	if errService != nil {
		response := utils.ApiResponse("Update Data Master Bahan Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Update Data Master Bahan Success", http.StatusOK, "success", bahan)
	c.JSON(http.StatusOK, response)
}

func (h *BahanHandlerImpl) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := utils.ApiResponse("Get data bahan failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inventory, errService := h.bahanUsecase.GetById(uint(id))
	if errService != nil {
		response := utils.ApiResponse("Get data bahan failed", http.StatusBadRequest, "error", errService)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Get data bahan success", http.StatusOK, "success", inventory)
	c.JSON(http.StatusOK, response)
}
