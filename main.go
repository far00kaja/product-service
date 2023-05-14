package main

import (
	"net/http"
	"sribuu/products/common/httpservice"
	"sribuu/products/controller"
	"sribuu/products/middleware"
	"sribuu/products/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	categoryService    service.CategoryService       = service.New()
	categoryController controller.CategoryController = controller.New(categoryService)
)

func main() {
	godotenv.Load()
	middleware.SetupLogger()
	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger())
	server.Use(gin.Logger())

	apiGroup := server.Group("/api")
	apiGroup.GET("/category", func(ctx *gin.Context) {
		ctx.JSON(200, categoryController.FindAll())
	})
	apiGroup.POST("/category", middleware.BasicAuth(), func(ctx *gin.Context) {
		result, err := categoryController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": httpservice.ErrBadRequestPayload.Error(),
				"status":  1,
			})
			return
		}
		ctx.JSON(http.StatusCreated, httpservice.ResponseData(result, err))
	})
	server.Run("localhost:9999")
}
