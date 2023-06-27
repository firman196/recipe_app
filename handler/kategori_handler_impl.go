package handler

import (
	"Recipe_App/models"
	"Recipe_App/usecase"
	"Recipe_App/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KategoriHandlerImpl struct {
	kategoriUsecase usecase.KategoriUsecase
}

func NewKategoriHandlerImpl(kategoriUsecase usecase.KategoriUsecase) KategoriHandler {
	return &KategoriHandlerImpl{
		kategoriUsecase: kategoriUsecase,
	}
}

// CreateKategori	godoc
// @Summary			Create kategori
// @Description		Create a new master data kategori.
// @Param			kategori body models.KategoriInput true "Create kategori"
// @Produce			application/json
// @Tags			kategori
// @Success			200 {object} utils.Response
// @Router			/api/v1/kategori [post]
func (h *KategoriHandlerImpl) Create(c *gin.Context) {
	var input models.KategoriInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Create Data Master Kategori Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	kategori, errService := h.kategoriUsecase.Create(input)
	if errService != nil {
		response := utils.ApiResponse("Create Data Master Kategori Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Create Data Master Kategori Success", http.StatusOK, "success", kategori)
	c.JSON(http.StatusOK, response)
}

// UpdateKategori	godoc
// @Summary			Update kategori
// @Description		Update data master kategori.
// @Param			bahan body models.KategoriInput true "Update kategori"
// @Param			id path integer true "find kategori by id"
// @Produce			application/json
// @Tags			kategori
// @Success			200 {object} utils.Response
// @Router			/api/v1/kategori/{id} [put]
func (h *KategoriHandlerImpl) Update(c *gin.Context) {
	id, errParam := strconv.Atoi(c.Param("id"))

	if errParam != nil {
		response := utils.ApiResponse("Invalid id params", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input models.KategoriInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.ApiResponse("Update Data Master Kategori Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	kategori, errService := h.kategoriUsecase.Update(uint(id), input)
	if errService != nil {
		response := utils.ApiResponse("Update Data Master Kategori Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Update Data Master Kategori Success", http.StatusOK, "success", kategori)
	c.JSON(http.StatusOK, response)
}

// GetById 				godoc
// @Summary				Get Single kategori by id.
// @Param				id path string true "get kategori by id"
// @Description			Return data master kategori where similar with id.
// @Produce				application/json
// @Tags				kategori
// @Success				200 {object} utils.Response
// @Router				/api/v1/kategori/{id} [get]
func (h *KategoriHandlerImpl) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := utils.ApiResponse("Invalid id params", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	kategori, errService := h.kategoriUsecase.GetById(uint(id))
	if errService != nil {
		response := utils.ApiResponse("Get data kategori failed", http.StatusBadRequest, "error", errService)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Get data kategori success", http.StatusOK, "success", kategori)
	c.JSON(http.StatusOK, response)
}

// Delete			 	godoc
// @Summary				Delete kategori by id.
// @Param				id path string true "delete kategori by id"
// @Description			Return data boolean.
// @Produce				application/json
// @Tags				kategori
// @Success				200 {object} utils.Response
// @Router				/api/v1/kategori/delete/{id} [delete]
func (h *KategoriHandlerImpl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := utils.ApiResponse("Invalid id params", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	kategori, errService := h.kategoriUsecase.Delete(uint(id))
	if errService != nil || !kategori {
		response := utils.ApiResponse("Delete data kategori failed", http.StatusBadRequest, "error", errService)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.ApiResponse("Get data kategori success", http.StatusOK, "success", kategori)
	c.JSON(http.StatusOK, response)
}

// GetAll 				godoc
// @Summary				Get all data kategori.
// @Param				page query int64 true "page number"
// @Param				limit query int64 true "limit number"
// @Description			Return data master kategori with pagination.
// @Produce				application/json
// @Tags				kategori
// @Success				200 {object} utils.Response
// @Router				/api/v1/kategori [get]
func (h *KategoriHandlerImpl) GetAll(c *gin.Context) {
	pagination := utils.Pagination(c)

	response, totalRows, err := h.kategoriUsecase.GetAll(&pagination)
	if err != nil {
		response := utils.ApiResponseWithPaginate("Get all data kategori failed", http.StatusBadRequest, "error", err, &totalRows)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	responses := utils.ApiResponseWithPaginate("Get data kategori success", http.StatusOK, "success", response, &totalRows)
	c.JSON(http.StatusOK, responses)
}
