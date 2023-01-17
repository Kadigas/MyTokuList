package routers

import (
	"mytokulist/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/categories", controllers.GetAllCategory)
	router.POST("/categories", controllers.InsertCategory)
	router.PUT("/categories/:id", controllers.UpdateCategory)
	router.DELETE("/categories/:id", controllers.DeleteCategory)

	router.GET("/types", controllers.GetAllType)
	router.POST("/types", controllers.InsertType)
	router.PUT("/types/:id", controllers.UpdateType)
	router.DELETE("/types/:id", controllers.DeleteType)

	router.GET("/status", controllers.GetAllStatus)
	router.POST("/status", controllers.InsertStatus)
	router.PUT("/status/:id", controllers.UpdateStatus)
	router.DELETE("/status/:id", controllers.DeleteStatus)

	return router
}
