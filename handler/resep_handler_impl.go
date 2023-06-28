package handler

import (
	"Recipe_App/models"
	"Recipe_App/usecase"
	"Recipe_App/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResepHandlerImpl struct {
	resepUsecase usecase.ResepUsecase
}

func NewResepHandlerImpl(resepUsecase usecase.ResepUsecase) ResepHandler {
	return &ResepHandlerImpl{
		resepUsecase: resepUsecase,
	}
}

// CreateResep	godoc
// @Summary			Create resep
// @Description		Create a new data resep.
// @Param			resep body models.ResepInput true "Create resep"
// @Produce			application/json
// @Tags			resep
// @Success			200 {object} utils.Response
// @Router			/api/v1/resep [post]
func (h *ResepHandlerImpl) Create(c *gin.Context) {
	var input models.ResepInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Create Data Resep Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	resep, errService := h.resepUsecase.Create(input)
	if errService != nil {
		response := utils.ApiResponse("Create Data Resep Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Create Data Resep Success", http.StatusOK, "success", resep)
	c.JSON(http.StatusOK, response)
}

// UpdateResep		godoc
// @Summary			Update resep
// @Description		Update data master resep.
// @Param			bahan body models.ResepInput true "Update resep"
// @Param			id path integer true "find resep by id"
// @Produce			application/json
// @Tags			resep
// @Success			200 {object} utils.Response
// @Router			/api/v1/resep/{id} [put]
func (h *ResepHandlerImpl) Update(c *gin.Context) {
	id, errParam := strconv.Atoi(c.Param("id"))

	if errParam != nil {
		response := utils.ApiResponse("Invalid id params", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input models.ResepInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Update Data Resep Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	resep, errService := h.resepUsecase.Update(uint(id), input)
	if errService != nil {
		response := utils.ApiResponse("Update Data Resep Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Update Data Resep Success", http.StatusOK, "success", resep)
	c.JSON(http.StatusOK, response)
}

// GetById 				godoc
// @Summary				Get Single resep by id.
// @Param				id path string true "get resep by id"
// @Description			Return data resep where similar with id.
// @Produce				application/json
// @Tags				resep
// @Success				200 {object} utils.Response
// @Router				/api/v1/resep/{id} [get]
func (h *ResepHandlerImpl) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := utils.ApiResponse("Invalid id params", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	resep, errService := h.resepUsecase.GetById(uint(id))
	if errService != nil {
		response := utils.ApiResponse("Get data resep failed", http.StatusBadRequest, "error", errService)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Get data resep success", http.StatusOK, "success", resep)
	c.JSON(http.StatusOK, response)
}

// Delete			 	godoc
// @Summary				Delete resep by id.
// @Param				id path string true "delete resep by id"
// @Description			Return data boolean.
// @Produce				application/json
// @Tags				resep
// @Success				200 {object} utils.Response
// @Router				/api/v1/resep/delete/{id} [delete]
func (h *ResepHandlerImpl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := utils.ApiResponse("Invalid id params", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	resep, errService := h.resepUsecase.Delete(uint(id))
	if errService != nil || !resep {
		response := utils.ApiResponse("Delete data resep failed", http.StatusBadRequest, "error", errService)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Get data resep success", http.StatusOK, "success", resep)
	c.JSON(http.StatusOK, response)
}
