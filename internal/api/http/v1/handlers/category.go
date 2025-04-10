package handlers

import (
	"inventory-service/internal/domain"
	"inventory-service/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	uc *usecase.CategoryUsecase
}

func NewCategoryHandler(uc *usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{uc}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var category domain.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.Create(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, category)
}

func (h *CategoryHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	category, err := h.uc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "некорректный ID"})
		return
	}

	var category domain.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "некорректные данные"})
		return
	}

	if err := h.uc.Update(uint(id), &category); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Категория успешно обновлена"})
}


func (h *CategoryHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "некорректный ID"})
		return
	}

	if err := h.uc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Категория успешно удалена"})
}


func (h *CategoryHandler) List(c *gin.Context) {
	name := c.Query("name")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	categories, err := h.uc.List(name, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}
