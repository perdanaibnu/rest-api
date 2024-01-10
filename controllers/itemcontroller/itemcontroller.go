package itemcontroller

import (
	"belajar-golang-db/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var items []models.Item

	models.DB.Find(&items)
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func Show(c *gin.Context) {
	var item models.Item
	id := c.Param("id")

	if err := models.DB.First(&item, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"item": item})

}

func Create(c *gin.Context) {
	var item models.Item

	if err := c.ShouldBindJSON(&item); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&item)
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil disimpan"})
}

func Update(c *gin.Context) {
	var item models.Item

	id := c.Param("id")
	if err := models.DB.Where("id = ?", id).First(&item).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
		return
	}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&item).Where("id = ?", id).Updates(&item).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diupdate"})
}

func Delete(c *gin.Context) {
	var item models.Item

	var input struct {
		Id json.Number
	}

	id, _ := input.Id.Int64()

	if err := models.DB.Where("id = ?", id).First(&item).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Delete(&item, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Tidak dapat menghapus"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
