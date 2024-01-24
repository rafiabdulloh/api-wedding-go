package main

import (
	dashboardController "wedding_be/controllers"

	"github.com/gin-gonic/gin"
	"wedding_be/models"
	"github.com/gin-contrib/cors"
)

func main() {
	models.GetConnect()
	models.Seeder()

	
	r := gin.Default()
	// Konfigurasi CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	r.Use(cors.New(config))

	r.GET("/", func(ctx *gin.Context) {
		dashboardController.Index(ctx)
	})

	r.GET("/informasi_mempelai", func(ctx *gin.Context) {
		dashboardController.GetAll(ctx, models.DB)
	})

	r.GET("/nomor_rekening", func(ctx *gin.Context){
		dashboardController.GetInfRek(ctx, models.DB)
	})

	r.POST("/komentar", func(ctx *gin.Context){
		dashboardController.POSTkomentar(ctx, models.DB)
	})

	r.GET("/komentar", func(ctx *gin.Context){
		dashboardController.GETkomentar(ctx, models.DB)
	})

	r.Run(":8081")
}
