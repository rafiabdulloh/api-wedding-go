package dashboardController

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"wedding_be/models"
	"gorm.io/gorm"



)

func Index(ctx *gin.Context){
	fmt.Println("Rafiiiiiiii Ganteng")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Mochamad Rafi Abdulloh",
	})
}

func GetAll(ctx *gin.Context, db *gorm.DB) {
	var data []models.Informasi_mempelai

	db.Find(&data)
	var count int64
	db.Model(&models.Informasi_mempelai{}).Count(&count)
	if count == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": data,
	})

}

func GetInfRek(ctx *gin.Context, db *gorm.DB){
	var data []models.Nomor_rekening

	db.Find(&data)
	var count int64
	db.Model(&models.Nomor_rekening{}).Count(&count)
	if count == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func POSTkomentar(ctx *gin.Context, db *gorm.DB) {
	var data models.Komentar
		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Data yang terikat:", data)
		// Simpan data ke database
		db.Create(&data)

		ctx.JSON(http.StatusOK, gin.H{"message": data})
	
}

func GETkomentar(ctx *gin.Context, db *gorm.DB) {
	var data []models.Komentar
	db.Find(&data)
	var count int64
	fmt.Println(data)
	db.Model(&models.Komentar{}).Count(&count)
	if count == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": data,
	})	
}
