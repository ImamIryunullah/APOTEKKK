package controllers

import (
	"apotek-management/config"
	"apotek-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTransaksi untuk menambahkan transaksi baru
func CreateTransaksi(c *gin.Context) {
	var transaksi models.Transaksi
	if err := c.ShouldBindJSON(&transaksi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&transaksi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, transaksi)
}



func GetAllTransaksi(c *gin.Context) {
	var transaksiList []models.Transaksi
	if err := config.DB.Find(&transaksiList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaksiList)
}

func GetTransaksiByID(c *gin.Context) {
	id := c.Param("id")
	var transaksi models.Transaksi
	if err := config.DB.First(&transaksi, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaksi not found"})
		return
	}

	c.JSON(http.StatusOK, transaksi)
}

// UpdateTransaksi untuk memperbarui transaksi berdasarkan ID
func UpdateTransaksi(c *gin.Context) {
	id := c.Param("id")
	var transaksi models.Transaksi
	if err := config.DB.First(&transaksi, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaksi not found"})
		return
	}

	if err := c.ShouldBindJSON(&transaksi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&transaksi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaksi)
}

// DeleteTransaksi untuk menghapus transaksi berdasarkan ID
func DeleteTransaksi(c *gin.Context) {
	id := c.Param("id")
	var transaksi models.Transaksi
	if err := config.DB.First(&transaksi, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaksi not found"})
		return
	}

	if err := config.DB.Delete(&transaksi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaksi deleted successfully"})
}

func CreateBatchTransaksi(c *gin.Context) {
	var transaksiList []models.Transaksi
	if err := c.ShouldBindJSON(&transaksiList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if err := config.DB.Create(&transaksiList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create transactions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Batch transactions created successfully",
		"data":    transaksiList,
	})
}

// UpdateBatchTransaksi untuk memperbarui beberapa transaksi sekaligus
func UpdateBatchTransaksi(c *gin.Context) {
	var transaksiList []models.Transaksi
	if err := c.ShouldBindJSON(&transaksiList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	for _, transaksi := range transaksiList {
		if err := config.DB.Save(&transaksi).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update some transactions"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Batch transactions updated successfully",
		"data":    transaksiList,
	})
}

// DeleteBatchTransaksi untuk menghapus beberapa transaksi sekaligus berdasarkan ID
func DeleteBatchTransaksi(c *gin.Context) {
	var ids []uint
	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if err := config.DB.Delete(&models.Transaksi{}, "id IN ?", ids).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete transactions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Batch transactions deleted successfully",
	})
}
