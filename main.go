package main

import (
	"net/http"
	"sribuu/products/controller"
	"sribuu/products/middleware"
	"sribuu/products/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	categoryService    service.CategoryService       = service.New()
	categoryController controller.CategoryController = controller.New(categoryService)
)

func main() {
	godotenv.Load()
	middleware.SetupLogger()
	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger(), gindump.Dump())
	server.Use(gin.Logger())
	server.GET("/category", func(ctx *gin.Context) {
		ctx.JSON(200, categoryController.FindAll())
	})
	server.POST("/category", middleware.BasicAuth(), func(ctx *gin.Context) {
		err := categoryController.Save(ctx)
		if err != nil {
			// ctx.BindHeader("content-type:application/json")
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"status":  1,
			})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "success created category",
			"status":  0,
		})

	})

	// server.GET("/categories", getCategories)
	// server.POST("/categories", postCategories)

	server.Run("localhost:9999")
}
