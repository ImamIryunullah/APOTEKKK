package controllers

import (
	"apotek-management/config"
	"apotek-management/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateObat(c *gin.Context) {
	var obat models.Obat
	if err := c.ShouldBindJSON(&obat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&obat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, obat)
}

// CreateBatchObat untuk menambahkan banyak obat sekaligus
func CreateBatchObat(c *gin.Context) {
	var obatList []models.Obat

	// Bind data JSON array ke struct Obat
	if err := c.ShouldBindJSON(&obatList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Menyimpan banyak obat ke database
	if err := config.DB.Create(&obatList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, obatList)
}


func GetAllObat(c *gin.Context) {
	var obatList []models.Obat
	if err := config.DB.Find(&obatList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, obatList)
}



func GetObatByID(c *gin.Context) {
	id := c.Param("id")
	var obat models.Obat
	if err := config.DB.First(&obat, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Obat not found"})
		return
	}

	c.JSON(http.StatusOK, obat)
}

func UpdateObat(c *gin.Context) {
	id := c.Param("id")
	var obat models.Obat
	if err := config.DB.First(&obat, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Obat not found"})
		return
	}

	if err := c.ShouldBindJSON(&obat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&obat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, obat)
}

// UpdateBatchObat untuk memperbarui banyak obat sekaligus
func UpdateBatchObat(c *gin.Context) {
	var obatList []models.Obat

	// Bind data JSON array ke struct Obat
	if err := c.ShouldBindJSON(&obatList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Memperbarui banyak obat
	for _, obat := range obatList {
		if err := config.DB.Save(&obat).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, obatList)
}


func DeleteObat(c *gin.Context) {
	id := c.Param("id")
	var obat models.Obat
	if err := config.DB.First(&obat, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Obat not found"})
		return
	}

	if err := config.DB.Delete(&obat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Obat deleted successfully"})
}

// DeleteBatchObat untuk menghapus banyak obat sekaligus
func DeleteBatchObat(c *gin.Context) {
	var ids []uint

	// Bind data JSON array ke array IDs
	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Menghapus obat berdasarkan ID
	if err := config.DB.Delete(&models.Obat{}, ids).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Obat(s) deleted successfully"})
}
