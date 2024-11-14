package routes

import (
	"ignbmd/blogging-platform-api/internal/app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(router *gin.Engine, postController *controllers.PostController) {
	router.GET("/posts", postController.FindAll)
	router.GET("/posts/:id", postController.FindByID)
	router.POST("/posts", postController.Create)
	router.PUT("/posts/:id", postController.Update)
	router.DELETE("/posts/:id", postController.Delete)
}
